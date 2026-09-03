# \WebhookAPI

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostWebhookMessageProxyBySecret**](WebhookAPI.md#PostWebhookMessageProxyBySecret) | **Post** /webhook/message-proxy/{secret} | SMS gateway webhook: store one forwarded SMS ({from, content: {text}}) for the engagement matching the path secret and sender number



## PostWebhookMessageProxyBySecret

> EngagementProxyAccepted PostWebhookMessageProxyBySecret(ctx, secret).Execute()

SMS gateway webhook: store one forwarded SMS ({from, content: {text}}) for the engagement matching the path secret and sender number

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gobugfree/cli/v3/api"
)

func main() {
	secret := "secret_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookAPI.PostWebhookMessageProxyBySecret(context.Background(), secret).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.PostWebhookMessageProxyBySecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostWebhookMessageProxyBySecret`: EngagementProxyAccepted
	fmt.Fprintf(os.Stdout, "Response from `WebhookAPI.PostWebhookMessageProxyBySecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**secret** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostWebhookMessageProxyBySecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EngagementProxyAccepted**](EngagementProxyAccepted.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

