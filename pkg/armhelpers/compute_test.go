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

func TestGetAvailabilitySet(t *testing.T) {
	mc, err := NewHTTPMockClient()
	if err != nil {
		t.Fatalf("failed to create HttpMockClient - %s", err)
	}

	mc.RegisterLogin()
	mc.RegisterGetAvailabilitySet()

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

	vmas, err := azureClient.GetAvailabilitySet(context.Background(), resourceGroup, virtualMachineAvailabilitySetName)
	if err != nil {
		t.Fatalf("can't get availability set: %s", err)
	}

	var expected int32 = 3
	if *vmas.Properties.PlatformFaultDomainCount != expected {
		t.Fatalf("expected PlatformFaultDomainCount of %d but got %v", expected, *vmas.Properties.PlatformFaultDomainCount)
	}
	if *vmas.Properties.PlatformUpdateDomainCount != expected {
		t.Fatalf("expected PlatformUpdateDomainCount of %d but got %v", expected, *vmas.Properties.PlatformUpdateDomainCount)
	}

	if vmas.Properties.ProximityPlacementGroup != nil && vmas.Properties.ProximityPlacementGroup.ID != nil {
		t.Fatalf("expected ProximityPlacementGroup of %q but got %v", "", *vmas.Properties.ProximityPlacementGroup.ID)
	}

	l := "eastus"
	if *vmas.Location != l {
		t.Fatalf("expected Location of %s but got %v", l, *vmas.Location)
	}
}
