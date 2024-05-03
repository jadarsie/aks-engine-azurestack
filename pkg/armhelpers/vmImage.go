// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package armhelpers

import (
	"context"

	compute "github.com/Azure/azure-sdk-for-go/profile/p20200901/resourcemanager/compute/armcompute"
	"github.com/pkg/errors"
)

// ListVirtualMachineImages returns the list of images available in the current environment
func (az *AzureClient) ListVirtualMachineImages(ctx context.Context, location, publisherName, offer, skus string) ([]*compute.VirtualMachineImageResource, error) {
	list, err := az.virtualMachineImagesClient.List(ctx, location, publisherName, offer, skus, nil)
	if err != nil {
		return nil, errors.Wrap(err, "listing virtual machine images")
	}
	return list.VirtualMachineImageResourceArray, nil
}

// GetVirtualMachineImage returns an image or an error where there is no image
func (az *AzureClient) GetVirtualMachineImage(ctx context.Context, location, publisherName, offer, skus, version string) (*compute.VirtualMachineImage, error) {
	image, err := az.virtualMachineImagesClient.Get(ctx, location, publisherName, offer, skus, version, nil)
	if err != nil {
		return nil, errors.Wrap(err, "fetching virtual machine image")
	}
	return &image.VirtualMachineImage, err
}
