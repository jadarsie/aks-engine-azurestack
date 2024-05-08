// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package armhelpers

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
)

func TestDeleteNetworkInterface(t *testing.T) {
	mc, err := NewHTTPMockClient()
	if err != nil {
		t.Fatalf("failed to create HttpMockClient - %s", err)
	}

	mc.RegisterLogin()
	mc.RegisterDeleteNetworkInterface()

	err = mc.Activate()
	if err != nil {
		t.Fatalf("failed to activate HttpMockClient - %s", err)
	}
	defer mc.DeactivateAndReset()

	env := mc.GetEnvironment()
	azureClient, err := NewAzureClient(env, subscriptionID, &fake.TokenCredential{}, nil)
	if err != nil {
		t.Fatalf("can not get client %s", err)
	}

	err = azureClient.DeleteNetworkInterface(context.Background(), resourceGroup, virtualNicName)
	if err != nil {
		t.Error(err)
	}
}
