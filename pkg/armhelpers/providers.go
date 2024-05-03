// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package armhelpers

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/profile/p20200901/resourcemanager/resources/armresources"
	"github.com/pkg/errors"
)

// ListProviders returns all the providers for a given AzureClient
func (az *AzureClient) ListProviders(ctx context.Context) ([]*armresources.Provider, error) {
	pager := az.providersClient.NewListPager(nil)
	list := []*armresources.Provider{}
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "listing providers")
		}
		list = append(list, page.Value...)
	}
	return list, nil
}
