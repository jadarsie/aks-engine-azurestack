//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetwork

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/profile/p20200901/internal"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// SecurityRulesClient contains the methods for the SecurityRules group.
// Don't use this type directly, use NewSecurityRulesClient() instead.
type SecurityRulesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSecurityRulesClient creates a new instance of SecurityRulesClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSecurityRulesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SecurityRulesClient, error) {
	cl, err := arm.NewClient(internal.ModuleName+"/armnetwork.SecurityRulesClient", internal.ModuleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SecurityRulesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a security rule in the specified network security group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-11-01
//   - resourceGroupName - The name of the resource group.
//   - networkSecurityGroupName - The name of the network security group.
//   - securityRuleName - The name of the security rule.
//   - securityRuleParameters - Parameters supplied to the create or update network security rule operation.
//   - options - SecurityRulesClientBeginCreateOrUpdateOptions contains the optional parameters for the SecurityRulesClient.BeginCreateOrUpdate
//     method.
func (client *SecurityRulesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, securityRuleName string, securityRuleParameters SecurityRule, options *SecurityRulesClientBeginCreateOrUpdateOptions) (*runtime.Poller[SecurityRulesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, networkSecurityGroupName, securityRuleName, securityRuleParameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[SecurityRulesClientCreateOrUpdateResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[SecurityRulesClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Creates or updates a security rule in the specified network security group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-11-01
func (client *SecurityRulesClient) createOrUpdate(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, securityRuleName string, securityRuleParameters SecurityRule, options *SecurityRulesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, networkSecurityGroupName, securityRuleName, securityRuleParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SecurityRulesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, securityRuleName string, securityRuleParameters SecurityRule, options *SecurityRulesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}/securityRules/{securityRuleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkSecurityGroupName == "" {
		return nil, errors.New("parameter networkSecurityGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkSecurityGroupName}", url.PathEscape(networkSecurityGroupName))
	if securityRuleName == "" {
		return nil, errors.New("parameter securityRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{securityRuleName}", url.PathEscape(securityRuleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, securityRuleParameters)
}

// BeginDelete - Deletes the specified network security rule.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-11-01
//   - resourceGroupName - The name of the resource group.
//   - networkSecurityGroupName - The name of the network security group.
//   - securityRuleName - The name of the security rule.
//   - options - SecurityRulesClientBeginDeleteOptions contains the optional parameters for the SecurityRulesClient.BeginDelete
//     method.
func (client *SecurityRulesClient) BeginDelete(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, securityRuleName string, options *SecurityRulesClientBeginDeleteOptions) (*runtime.Poller[SecurityRulesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, networkSecurityGroupName, securityRuleName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[SecurityRulesClientDeleteResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[SecurityRulesClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes the specified network security rule.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-11-01
func (client *SecurityRulesClient) deleteOperation(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, securityRuleName string, options *SecurityRulesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, networkSecurityGroupName, securityRuleName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SecurityRulesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, securityRuleName string, options *SecurityRulesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}/securityRules/{securityRuleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkSecurityGroupName == "" {
		return nil, errors.New("parameter networkSecurityGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkSecurityGroupName}", url.PathEscape(networkSecurityGroupName))
	if securityRuleName == "" {
		return nil, errors.New("parameter securityRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{securityRuleName}", url.PathEscape(securityRuleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Get the specified network security rule.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-11-01
//   - resourceGroupName - The name of the resource group.
//   - networkSecurityGroupName - The name of the network security group.
//   - securityRuleName - The name of the security rule.
//   - options - SecurityRulesClientGetOptions contains the optional parameters for the SecurityRulesClient.Get method.
func (client *SecurityRulesClient) Get(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, securityRuleName string, options *SecurityRulesClientGetOptions) (SecurityRulesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, networkSecurityGroupName, securityRuleName, options)
	if err != nil {
		return SecurityRulesClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SecurityRulesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SecurityRulesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SecurityRulesClient) getCreateRequest(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, securityRuleName string, options *SecurityRulesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}/securityRules/{securityRuleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkSecurityGroupName == "" {
		return nil, errors.New("parameter networkSecurityGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkSecurityGroupName}", url.PathEscape(networkSecurityGroupName))
	if securityRuleName == "" {
		return nil, errors.New("parameter securityRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{securityRuleName}", url.PathEscape(securityRuleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SecurityRulesClient) getHandleResponse(resp *http.Response) (SecurityRulesClientGetResponse, error) {
	result := SecurityRulesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecurityRule); err != nil {
		return SecurityRulesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets all security rules in a network security group.
//
// Generated from API version 2018-11-01
//   - resourceGroupName - The name of the resource group.
//   - networkSecurityGroupName - The name of the network security group.
//   - options - SecurityRulesClientListOptions contains the optional parameters for the SecurityRulesClient.NewListPager method.
func (client *SecurityRulesClient) NewListPager(resourceGroupName string, networkSecurityGroupName string, options *SecurityRulesClientListOptions) *runtime.Pager[SecurityRulesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[SecurityRulesClientListResponse]{
		More: func(page SecurityRulesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SecurityRulesClientListResponse) (SecurityRulesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, networkSecurityGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SecurityRulesClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return SecurityRulesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SecurityRulesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *SecurityRulesClient) listCreateRequest(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, options *SecurityRulesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}/securityRules"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkSecurityGroupName == "" {
		return nil, errors.New("parameter networkSecurityGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkSecurityGroupName}", url.PathEscape(networkSecurityGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SecurityRulesClient) listHandleResponse(resp *http.Response) (SecurityRulesClientListResponse, error) {
	result := SecurityRulesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecurityRuleListResult); err != nil {
		return SecurityRulesClientListResponse{}, err
	}
	return result, nil
}
