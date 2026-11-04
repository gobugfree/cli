# \LegacyAPI

All URIs are relative to *https://app.gobugfree.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLegacy**](LegacyAPI.md#GetLegacy) | **Get** /legacy | Where athena&#39;s home led: a redirect to the program list (JSON: the target)
[**GetLegacyByRest**](LegacyAPI.md#GetLegacyByRest) | **Get** /legacy/{rest} | Where an athena link leads: a redirect to the report its id maps to, or to the page that replaced athena&#39;s (JSON: the target, 404 for an unknown id or path)



## GetLegacy

> GetLegacy(ctx).Execute()

Where athena's home led: a redirect to the program list (JSON: the target)

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LegacyAPI.GetLegacy(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LegacyAPI.GetLegacy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLegacyRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLegacyByRest

> LegacylinkLink GetLegacyByRest(ctx, rest).Execute()

Where an athena link leads: a redirect to the report its id maps to, or to the page that replaced athena's (JSON: the target, 404 for an unknown id or path)

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
	rest := "rest_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LegacyAPI.GetLegacyByRest(context.Background(), rest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LegacyAPI.GetLegacyByRest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLegacyByRest`: LegacylinkLink
	fmt.Fprintf(os.Stdout, "Response from `LegacyAPI.GetLegacyByRest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rest** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLegacyByRestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LegacylinkLink**](LegacylinkLink.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

