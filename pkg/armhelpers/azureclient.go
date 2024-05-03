// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package armhelpers

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Azure/aks-engine-azurestack/pkg/kubernetes"
	authorization "github.com/Azure/azure-sdk-for-go/profile/p20200901/resourcemanager/authorization/armauthorization"
	compute "github.com/Azure/azure-sdk-for-go/profile/p20200901/resourcemanager/compute/armcompute"
	network "github.com/Azure/azure-sdk-for-go/profile/p20200901/resourcemanager/network/armnetwork"
	resources "github.com/Azure/azure-sdk-for-go/profile/p20200901/resourcemanager/resources/armresources"
	subscriptions "github.com/Azure/azure-sdk-for-go/profile/p20200901/resourcemanager/resources/armsubscriptions"
	storage "github.com/Azure/azure-sdk-for-go/profile/p20200901/resourcemanager/storage/armstorage"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

const (
	// ApplicationDir is the name of the dir where the token is cached
	ApplicationDir = ".acsengine"
)

var (
	// RequiredResourceProviders is the list of Azure Resource Providers needed for AKS Engine to function
	RequiredResourceProviders = []string{"Microsoft.Compute", "Microsoft.Storage", "Microsoft.Network"}
)

// AzureClient implements the `AKSEngineClient` interface.
// This client is backed by real Azure clients talking to an ARM endpoint.
type AzureClient struct {
	acceptLanguages []string
	auxiliaryTokens []string
	environment     azure.Environment
	subscriptionID  string

	authorizationClient        *authorization.RoleAssignmentsClient
	deploymentsClient          *resources.DeploymentsClient
	deploymentOperationsClient *resources.DeploymentOperationsClient
	storageAccountsClient      *storage.AccountsClient
	storageBlobClientFactory   func(key, blobURI string) (*azblob.Client, error)
	interfacesClient           *network.InterfacesClient
	groupsClient               *resources.ResourceGroupsClient
	subscriptionsClient        *subscriptions.Client
	providersClient            *resources.ProvidersClient
	virtualMachinesClient      *compute.VirtualMachinesClient
	disksClient                *compute.DisksClient
	availabilitySetsClient     *compute.AvailabilitySetsClient
	virtualMachineImagesClient *compute.VirtualMachineImagesClient
}

// GetKubernetesClient returns a KubernetesClient hooked up to the api server at the apiserverURL.
func (az *AzureClient) GetKubernetesClient(apiserverURL, kubeConfig string, interval, timeout time.Duration) (kubernetes.Client, error) {
	return kubernetes.NewClient(apiserverURL, kubeConfig, interval, timeout)
}

// NewAzureClientWithClientSecret returns an AzureClient via client_id and client_secret
func NewAzureClientWithClientSecret(env azure.Environment, subscriptionID, clientID, clientSecret string) (*AzureClient, error) {
	tenantID, err := getOAuthConfig(env, subscriptionID)
	if err != nil {
		return nil, err
	}
	creds, err := azidentity.NewClientSecretCredential(tenantID, clientID, clientSecret, nil)
	if err != nil {
		return nil, err
	}
	client := getClient(env, subscriptionID, creds)
	client.storageBlobClientFactory = keysBlobClient()
	return client, nil
}

// NewAzureClientWithClientSecretExternalTenant returns an AzureClient via client_id and client_secret from a tenant
func NewAzureClientWithClientSecretExternalTenant(env azure.Environment, subscriptionID, tenantID, clientID, clientSecret string) (*AzureClient, error) {
	creds, err := azidentity.NewClientSecretCredential(tenantID, clientID, clientSecret, nil)
	if err != nil {
		return nil, err
	}
	client := getClient(env, subscriptionID, creds)
	client.storageBlobClientFactory = keysBlobClient()
	return client, nil
}

// NewAzureClientWithClientCertificateFile returns an AzureClient via client_id and jwt certificate assertion
func NewAzureClientWithClientCertificateFile(env azure.Environment, subscriptionID, clientID, certificatePath, privateKeyPath string) (*AzureClient, error) {
	certificateData, err := os.ReadFile(certificatePath)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to read certificate")
	}

	block, _ := pem.Decode(certificateData)
	if block == nil {
		return nil, errors.New("Failed to decode pem block from certificate")
	}

	certificate, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to parse certificate")
	}

	privateKey, err := parseRsaPrivateKey(privateKeyPath)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to parse rsa private key")
	}

	return NewAzureClientWithClientCertificate(env, subscriptionID, clientID, certificate, privateKey)
}

// NewAzureClientWithClientCertificateFileExternalTenant returns an AzureClient via client_id and jwt certificate assertion a 3rd party tenant
func NewAzureClientWithClientCertificateFileExternalTenant(env azure.Environment, subscriptionID, tenantID, clientID, certificatePath, privateKeyPath string) (*AzureClient, error) {
	certificate, err := parseCertificate(certificatePath)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to parse certificate")
	}

	privateKey, err := parseRsaPrivateKey(privateKeyPath)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to parse rsa private key")
	}

	return NewAzureClientWithClientCertificateExternalTenant(env, subscriptionID, tenantID, clientID, certificate, privateKey)
}

// NewAzureClientWithClientCertificate returns an AzureClient via client_id and jwt certificate assertion
func NewAzureClientWithClientCertificate(env azure.Environment, subscriptionID, clientID string, certificate *x509.Certificate, privateKey *rsa.PrivateKey) (*AzureClient, error) {
	tenantID, err := getOAuthConfig(env, subscriptionID)
	if err != nil {
		return nil, err
	}
	return newAzureClientWithCertificate(env, subscriptionID, clientID, tenantID, certificate, privateKey)
}

// NewAzureClientWithClientCertificateExternalTenant returns an AzureClient via client_id and jwt certificate assertion against a 3rd party tenant
func NewAzureClientWithClientCertificateExternalTenant(env azure.Environment, subscriptionID, tenantID, clientID string, certificate *x509.Certificate, privateKey *rsa.PrivateKey) (*AzureClient, error) {
	return newAzureClientWithCertificate(env, subscriptionID, clientID, tenantID, certificate, privateKey)
}

func newAzureClientWithCertificate(env azure.Environment, subscriptionID, clientID, tenantID string, certificate *x509.Certificate, privateKey *rsa.PrivateKey) (*AzureClient, error) {
	if certificate == nil {
		return nil, errors.New("certificate should not be nil")
	}
	if privateKey == nil {
		return nil, errors.New("privateKey should not be nil")
	}
	creds, err := azidentity.NewClientCertificateCredential(tenantID, clientID, []*x509.Certificate{certificate}, privateKey, nil)
	if err != nil {
		return nil, err
	}
	client := getClient(env, subscriptionID, creds)
	client.storageBlobClientFactory = keysBlobClient()
	return client, nil
}

func getOAuthConfig(env azure.Environment, subscriptionID string) (string, error) {
	tenantID, err := GetTenantID(subscriptionID)
	if err != nil {
		return "", err
	}
	return tenantID, nil
}

func getClient(env azure.Environment, subscriptionID string, creds azcore.TokenCredential) *AzureClient {
	c := &AzureClient{}
	c.environment = env
	c.subscriptionID = subscriptionID

	c.authorizationClient, _ = authorization.NewRoleAssignmentsClient(subscriptionID, creds, nil)
	c.deploymentsClient, _ = resources.NewDeploymentsClient(subscriptionID, creds, nil)
	c.deploymentOperationsClient, _ = resources.NewDeploymentOperationsClient(subscriptionID, creds, nil)
	c.storageAccountsClient, _ = storage.NewAccountsClient(subscriptionID, creds, nil)
	c.interfacesClient, _ = network.NewInterfacesClient(subscriptionID, creds, nil)
	c.groupsClient, _ = resources.NewResourceGroupsClient(subscriptionID, creds, nil)
	c.subscriptionsClient, _ = subscriptions.NewClient(creds, nil)
	c.providersClient, _ = resources.NewProvidersClient(subscriptionID, creds, nil)
	c.virtualMachinesClient, _ = compute.NewVirtualMachinesClient(subscriptionID, creds, nil)
	c.disksClient, _ = compute.NewDisksClient(subscriptionID, creds, nil)
	c.availabilitySetsClient, _ = compute.NewAvailabilitySetsClient(subscriptionID, creds, nil)
	c.virtualMachineImagesClient, _ = compute.NewVirtualMachineImagesClient(subscriptionID, creds, nil)

	return c
}

// EnsureProvidersRegistered checks if the AzureClient is registered to required resource providers and, if not, register subscription to providers
func (az *AzureClient) EnsureProvidersRegistered(subscriptionID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), DefaultARMOperationTimeout)
	defer cancel()
	pager := az.providersClient.NewListPager(nil)
	providers := []*resources.Provider{}
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return errors.Errorf("Error listing registered providers for subscription %s", subscriptionID)
		}
		providers = append(providers, page.Value...)
	}

	m := make(map[string]bool)
	for _, provider := range providers {
		m[strings.ToLower(to.String(provider.Namespace))] = to.String(provider.RegistrationState) == "Registered"
	}

	for _, provider := range RequiredResourceProviders {
		registered, ok := m[strings.ToLower(provider)]
		if !ok {
			return errors.Errorf("Unknown resource provider %q", provider)
		}
		if registered {
			log.Debugf("Already registered for %q", provider)
		} else {
			log.Infof("Registering subscription to resource provider. provider=%q subscription=%q", provider, subscriptionID)
			if _, err := az.providersClient.Register(ctx, provider, nil); err != nil {
				return err
			}
		}
	}
	return nil
}

func parseCertificate(certificatePath string) (*x509.Certificate, error) {
	certificateData, err := os.ReadFile(certificatePath)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to read certificate")
	}

	block, _ := pem.Decode(certificateData)
	if block == nil {
		return nil, errors.New("Failed to decode pem block from certificate")
	}

	certificate, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to parse certificate")
	}
	return certificate, nil
}

func parseRsaPrivateKey(path string) (*rsa.PrivateKey, error) {
	privateKeyData, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(privateKeyData)
	if block == nil {
		return nil, errors.New("Failed to decode a pem block from private key")
	}

	privatePkcs1Key, errPkcs1 := x509.ParsePKCS1PrivateKey(block.Bytes)
	if errPkcs1 == nil {
		return privatePkcs1Key, nil
	}

	privatePkcs8Key, errPkcs8 := x509.ParsePKCS8PrivateKey(block.Bytes)
	if errPkcs8 == nil {
		privatePkcs8RsaKey, ok := privatePkcs8Key.(*rsa.PrivateKey)
		if !ok {
			return nil, errors.New("pkcs8 contained non-RSA key. Expected RSA key")
		}
		return privatePkcs8RsaKey, nil
	}

	return nil, errors.Errorf("failed to parse private key as Pkcs#1 or Pkcs#8. (%s). (%s)", errPkcs1, errPkcs8)
}

func keysBlobClient() func(key, blobURI string) (*azblob.Client, error) {
	return func(key, blobURI string) (*azblob.Client, error) {
		parts, err := azblob.ParseURL(blobURI)
		if err != nil {
			return nil, err
		}
		name := strings.Split(parts.Host, ".")[0]
		sas, err := azblob.NewSharedKeyCredential(name, key)
		if err != nil {
			return nil, err
		}
		return azblob.NewClientWithSharedKeyCredential(fmt.Sprintf("%s%s", parts.Scheme, parts.Host), sas, nil)
	}
}
