// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package armhelpers

import (
	"context"
	"testing"
)

func TestDeleteVirtualMachine(t *testing.T) {
	mc, err := NewHTTPMockClient()
	if err != nil {
		t.Fatalf("failed to create HttpMockClient - %s", err)
	}

	mc.RegisterLogin()
	mc.RegisterVirtualMachineEndpoint()
	mc.RegisterDeleteOperation()

	err = mc.Activate()
	if err != nil {
		t.Fatalf("failed to activate HttpMockClient - %s", err)
	}
	defer mc.DeactivateAndReset()

	env := mc.GetEnvironment()
	azureClient, err := NewAzureClientWithClientSecret(env, subscriptionID, "clientID", "secret")
	if err != nil {
		t.Fatalf("can not get client %s", err)
	}

	err = azureClient.DeleteVirtualMachine(context.Background(), resourceGroup, virtualMachineName)
	if err != nil {
		t.Error(err)
	}
}
