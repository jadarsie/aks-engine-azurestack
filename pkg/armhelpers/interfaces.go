// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package armhelpers

import (
	"context"
	"time"

	"github.com/Azure/aks-engine-azurestack/pkg/kubernetes"
	"github.com/Azure/azure-sdk-for-go/profile/p20200901/resourcemanager/authorization/armauthorization"
	"github.com/Azure/azure-sdk-for-go/profile/p20200901/resourcemanager/compute/armcompute"
	"github.com/Azure/azure-sdk-for-go/profile/p20200901/resourcemanager/resources/armresources"
)

// VMImageFetcher is an extension of AKSEngine client allows us to operate on the virtual machine images in the environment
type VMImageFetcher interface {
	// ListVirtualMachineImages return a list of images
	ListVirtualMachineImages(ctx context.Context, location, publisherName, offer, skus string) ([]*armcompute.VirtualMachineImageResource, error)
	// GetVirtualMachineImage return a virtual machine image
	GetVirtualMachineImage(ctx context.Context, location, publisherName, offer, skus, version string) (*armcompute.VirtualMachineImage, error)
}

// AKSEngineClient is the interface used to talk to an Azure environment.
// This interface exposes just the subset of Azure APIs and clients needed for AKS Engine.
type AKSEngineClient interface {

	//AddAcceptLanguages sets the list of languages to accept on this request
	// AddAcceptLanguages(languages []string)

	// RESOURCES
	// DeployTemplate can deploy a template into Azure ARM
	DeployTemplate(ctx context.Context, resourceGroup, name string, template, parameters map[string]interface{}) (*armresources.DeploymentExtended, error)
	// EnsureResourceGroup ensures the specified resource group exists in the specified location
	EnsureResourceGroup(ctx context.Context, resourceGroup, location string, managedBy *string) (*armresources.ResourceGroup, error)

	// COMPUTE
	// ListVirtualMachines lists VM resources
	ListVirtualMachines(ctx context.Context, resourceGroup string) ([]*armcompute.VirtualMachine, error)
	// GetVirtualMachine retrieves the specified virtual machine.
	GetVirtualMachine(ctx context.Context, resourceGroup, name string) (*armcompute.VirtualMachine, error)
	// RestartVirtualMachine restarts the specified virtual machine.
	RestartVirtualMachine(ctx context.Context, resourceGroup, name string) error
	// DeleteVirtualMachine deletes the specified virtual machine.
	DeleteVirtualMachine(ctx context.Context, resourceGroup, name string) error
	// GetVirtualMachinePowerState returns the virtual machine's PowerState status code
	GetVirtualMachinePowerState(ctx context.Context, resourceGroup, name string) (string, error)

	//
	// STORAGE
	DeleteVirtualHardDisk(ctx context.Context, resourceGroup string, vhd *armcompute.VirtualHardDisk) error

	//
	// NETWORK

	// DeleteNetworkInterface deletes the specified network interface.
	DeleteNetworkInterface(ctx context.Context, resourceGroup, nicName string) error

	// RBAC
	DeleteRoleAssignmentByID(ctx context.Context, roleAssignmentNameID string) (*armauthorization.RoleAssignment, error)
	ListRoleAssignmentsForPrincipal(ctx context.Context, scope string, principalID string) ([]*armauthorization.RoleAssignment, error)

	// MANAGED DISKS
	DeleteManagedDisk(ctx context.Context, resourceGroupName string, diskName string) error
	ListManagedDisksByResourceGroup(ctx context.Context, resourceGroupName string) ([]*armcompute.Disk, error)

	GetKubernetesClient(apiserverURL, kubeConfig string, interval, timeout time.Duration) (kubernetes.Client, error)

	ListProviders(ctx context.Context) ([]*armresources.Provider, error)

	// DEPLOYMENTS

	// ListDeploymentOperations gets all deployments operations for a deployment.
	ListDeploymentOperations(ctx context.Context, resourceGroupName string, deploymentName string) ([]*armresources.DeploymentOperation, error)
}
