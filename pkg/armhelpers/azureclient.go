// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package armhelpers

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/Azure/aks-engine-azurestack/pkg/kubernetes"
	authorization "github.com/Azure/azure-sdk-for-go/profile/p20200901/resourcemanager/authorization/armauthorization"
	compute "github.com/Azure/azure-sdk-for-go/profile/p20200901/resourcemanager/compute/armcompute"
	network "github.com/Azure/azure-sdk-for-go/profile/p20200901/resourcemanager/network/armnetwork"
	resources "github.com/Azure/azure-sdk-for-go/profile/p20200901/resourcemanager/resources/armresources"
	storage "github.com/Azure/azure-sdk-for-go/profile/p20200901/resourcemanager/storage/armstorage"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

var (
	// RequiredResourceProviders is the list of Azure Resource Providers needed for AKS Engine to function
	RequiredResourceProviders = []string{"Microsoft.Compute", "Microsoft.Storage", "Microsoft.Network"}
)

// AzureClient implements the `AKSEngineClient` interface.
// This client is backed by real Azure clients talking to an ARM endpoint.
type AzureClient struct {
	environment    azure.Environment
	subscriptionID string

	authorizationClient        *authorization.RoleAssignmentsClient
	deploymentsClient          *resources.DeploymentsClient
	deploymentOperationsClient *resources.DeploymentOperationsClient
	storageAccountsClient      *storage.AccountsClient
	storageBlobClientFactory   func(key, blobURI string) (*azblob.Client, error)
	interfacesClient           *network.InterfacesClient
	groupsClient               *resources.ResourceGroupsClient
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
func NewAzureClient(env azure.Environment, subscriptionID string, creds azcore.TokenCredential, options *arm.ClientOptions) (*AzureClient, error) {
	c := &AzureClient{}
	c.environment = env
	c.subscriptionID = subscriptionID
	c.authorizationClient, _ = authorization.NewRoleAssignmentsClient(subscriptionID, creds, options)
	c.deploymentsClient, _ = resources.NewDeploymentsClient(subscriptionID, creds, options)
	c.deploymentOperationsClient, _ = resources.NewDeploymentOperationsClient(subscriptionID, creds, options)
	c.storageAccountsClient, _ = storage.NewAccountsClient(subscriptionID, creds, options)
	c.interfacesClient, _ = network.NewInterfacesClient(subscriptionID, creds, options)
	c.groupsClient, _ = resources.NewResourceGroupsClient(subscriptionID, creds, options)
	c.providersClient, _ = resources.NewProvidersClient(subscriptionID, creds, options)
	c.virtualMachinesClient, _ = compute.NewVirtualMachinesClient(subscriptionID, creds, options)
	c.disksClient, _ = compute.NewDisksClient(subscriptionID, creds, options)
	c.availabilitySetsClient, _ = compute.NewAvailabilitySetsClient(subscriptionID, creds, options)
	c.virtualMachineImagesClient, _ = compute.NewVirtualMachineImagesClient(subscriptionID, creds, options)
	c.storageBlobClientFactory = keysBlobClient()
	return c, nil
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
