// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package armhelpers

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/profile/p20200901/resourcemanager/resources/armresources"
)

// EnsureResourceGroup ensures the named resource group exists in the given location.
func (az *AzureClient) EnsureResourceGroup(ctx context.Context, name, location string, managedBy *string) (*armresources.ResourceGroup, error) {
	var tags map[string]*string
	group, err := az.groupsClient.Get(ctx, name, nil)
	if err == nil {
		tags = group.Tags
	}
	response, err := az.groupsClient.CreateOrUpdate(ctx, name, armresources.ResourceGroup{
		Name:      &name,
		Location:  &location,
		ManagedBy: managedBy,
		Tags:      tags,
	}, nil)

	if err != nil {
		return nil, err
	}

	return &response.ResourceGroup, nil
}

// CheckResourceGroupExistence return if the resource group exists
func (az *AzureClient) CheckResourceGroupExistence(ctx context.Context, name string) (armresources.ResourceGroupsClientCheckExistenceResponse, error) {
	return az.groupsClient.CheckExistence(ctx, name, nil)
}

// DeleteResourceGroup delete the named resource group
func (az *AzureClient) DeleteResourceGroup(ctx context.Context, name string) error {
	poller, err := az.groupsClient.BeginDelete(ctx, name, nil)
	if err != nil {
		return err
	}
	if _, err = poller.PollUntilDone(ctx, nil); err != nil {
		return err
	}
	return err
}
