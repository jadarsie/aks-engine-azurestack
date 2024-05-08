// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package armhelpers

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/pkg/errors"
)

// NewClientSecretCredential returns an AzureClient via client_id and client_secret
func NewClientSecretCredential(env azure.Environment, subscriptionID, clientID, clientSecret string, options *azidentity.ClientSecretCredentialOptions) (*azidentity.ClientSecretCredential, error) {
	tenantID, err := getOAuthConfig(&fake.TokenCredential{}, env, subscriptionID)
	if err != nil {
		return nil, err
	}
	creds, err := azidentity.NewClientSecretCredential(tenantID, clientID, clientSecret, options)
	if err != nil {
		return nil, err
	}
	return creds, nil
}

// NewClientSecretCredentialExternalTenant returns an AzureClient via client_id and client_secret from a tenant
func NewClientSecretCredentialExternalTenant(env azure.Environment, subscriptionID, clientID, clientSecret string, options *azidentity.ClientSecretCredentialOptions) (*azidentity.ClientSecretCredential, error) {
	tenantID, err := getOAuthConfig(&fake.TokenCredential{}, env, subscriptionID)
	if err != nil {
		return nil, err
	}
	creds, err := azidentity.NewClientSecretCredential(tenantID, clientID, clientSecret, options)
	if err != nil {
		return nil, err
	}
	return creds, nil
}

// NewClientCertificateCredential returns an AzureClient via client_id and jwt certificate assertion
func NewClientCertificateCredential(env azure.Environment, subscriptionID, clientID, certificatePath, privateKeyPath string, options *azidentity.ClientCertificateCredentialOptions) (*azidentity.ClientCertificateCredential, error) {
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
	tenantID, err := getOAuthConfig(&fake.TokenCredential{}, env, subscriptionID)
	if err != nil {
		return nil, err
	}
	return azidentity.NewClientCertificateCredential(tenantID, clientID, []*x509.Certificate{certificate}, privateKey, options)
}

// NewClientCertificateCredentialExternalTenant returns an AzureClient via client_id and jwt certificate assertion a 3rd party tenant
func NewClientCertificateCredentialExternalTenant(env azure.Environment, subscriptionID, clientID, certificatePath, privateKeyPath string, options *azidentity.ClientCertificateCredentialOptions) (*azidentity.ClientCertificateCredential, error) {
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
	tenantID, err := getOAuthConfig(&fake.TokenCredential{}, env, subscriptionID)
	if err != nil {
		return nil, err
	}
	return azidentity.NewClientCertificateCredential(tenantID, clientID, []*x509.Certificate{certificate}, privateKey, options)
}

func getOAuthConfig(credential azcore.TokenCredential, env azure.Environment, subscriptionID string) (string, error) {
	tenantID, err := GetTenantID(credential, subscriptionID, nil)
	if err != nil {
		return env.Name, err
	}
	return tenantID, nil
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
