// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package armhelpers

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
)

func TestDeleteManagedDisk(t *testing.T) {
	mc, err := NewHTTPMockClient()
	if err != nil {
		t.Fatalf("failed to create HttpMockClient - %s", err)
	}

	mc.RegisterLogin()
	mc.RegisterDeleteManagedDisk()

	err = mc.Activate()
	if err != nil {
		t.Fatalf("failed to activate HttpMockClient - %s", err)
	}
	defer mc.DeactivateAndReset()

	azureClient, err := NewAzureClient(subscriptionID, &fake.TokenCredential{}, cloud.AzurePublic)
	if err != nil {
		t.Fatalf("can not get client %s", err)
	}

	err = azureClient.DeleteManagedDisk(context.Background(), resourceGroup, virutalDiskName)
	if err != nil {
		t.Error(err)
	}
}
