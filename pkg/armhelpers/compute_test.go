// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package armhelpers

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
)

type LoggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (lrw LoggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
}

// Handler using this response writer
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lrw := &LoggingResponseWriter{ResponseWriter: w}
		next.ServeHTTP(lrw, r)
		fmt.Printf("HTTP %d %s", lrw.statusCode, r.URL.String())
	})
}

func TestDeleteVirtualMachine(t *testing.T) {
	mc, err := NewHTTPMockClient()
	if err != nil {
		t.Fatalf("failed to create HttpMockClient - %s", err)
	}

	mc.RegisterLogin()
	mc.RegisterVirtualMachineEndpoint()
	mc.RegisterDeleteOperation()

	// req := &http.Request{
	// 	Method: http.MethodGet,
	// 	URL: &url.URL{
	// 		Path:     fmt.Sprintf("/subscriptions/%s", subscriptionID),
	// 		RawQuery: "api-version=2016-06-01",
	// 	},
	// 	// Host: "management.azure.com",
	// }
	// handler, _ := mc.mux.Handler(req)
	// handler.ServeHTTP(LoggingResponseWriter{}, req)

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

	err = azureClient.DeleteVirtualMachine(context.Background(), resourceGroup, virtualMachineName)
	if err != nil {
		t.Error(err)
	}
}
