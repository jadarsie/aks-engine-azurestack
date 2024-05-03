// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package armhelpers

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/profile/p20200901/resourcemanager/resources/armresources"
	log "github.com/sirupsen/logrus"
)

// DeployTemplate implements the TemplateDeployer interface for the AzureClient client
func (az *AzureClient) DeployTemplate(ctx context.Context, resourceGroupName, deploymentName string, template map[string]interface{}, parameters map[string]interface{}) (*armresources.DeploymentExtended, error) {
	mode := armresources.DeploymentModeIncremental
	deployment := armresources.Deployment{
		Properties: &armresources.DeploymentProperties{
			Template:   &template,
			Parameters: &parameters,
			Mode:       &mode,
		},
	}
	log.Infof("Starting ARM Deployment %s in resource group %s. This will take some time...", deploymentName, resourceGroupName)
	poller, err := az.deploymentsClient.BeginCreateOrUpdate(ctx, resourceGroupName, deploymentName, deployment, nil)
	if err != nil {
		return nil, err
	}
	outcomeText := "Succeeded"
	de, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		outcomeText = fmt.Sprintf("Error: %v", err)
		log.Infof("Finished ARM Deployment (%s). %s", deploymentName, outcomeText)
		return nil, err
	}
	log.Infof("Finished ARM Deployment (%s). %s", deploymentName, outcomeText)
	return &de.DeploymentExtended, err
}

// ValidateTemplate validate the template and parameters
func (az *AzureClient) ValidateTemplate(ctx context.Context, resourceGroupName, deploymentName string, template map[string]interface{}, parameters map[string]interface{}) (*armresources.DeploymentsClientValidateResponse, error) {
	mode := armresources.DeploymentModeIncremental
	deployment := armresources.Deployment{
		Properties: &armresources.DeploymentProperties{
			Template:   &template,
			Parameters: &parameters,
			Mode:       &mode,
		},
	}
	poller, err := az.deploymentsClient.BeginValidate(ctx, resourceGroupName, deploymentName, deployment, nil)
	if err != nil {
		return nil, err
	}
	response, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		return nil, err
	}
	return &response, err
}

// GetDeployment returns the template deployment
func (az *AzureClient) GetDeployment(ctx context.Context, resourceGroupName, deploymentName string) (armresources.DeploymentsClientGetResponse, error) {
	return az.deploymentsClient.Get(ctx, resourceGroupName, deploymentName, nil)
}

// CheckDeploymentExistence returns if the deployment already exists
func (az *AzureClient) CheckDeploymentExistence(ctx context.Context, resourceGroupName string, deploymentName string) (armresources.DeploymentsClientCheckExistenceResponse, error) {
	return az.deploymentsClient.CheckExistence(ctx, resourceGroupName, deploymentName, nil)
}
