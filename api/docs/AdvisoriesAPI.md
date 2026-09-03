# \AdvisoriesAPI

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAdvisoriesByDisclosureId**](AdvisoriesAPI.md#GetAdvisoriesByDisclosureId) | **Get** /advisories/{disclosure_id} | The public advisory of a published disclosure; 404 until published



## GetAdvisoriesByDisclosureId

> DisclosureAdvisory GetAdvisoriesByDisclosureId(ctx, disclosureId).Execute()

The public advisory of a published disclosure; 404 until published

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
	disclosureId := "disclosureId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdvisoriesAPI.GetAdvisoriesByDisclosureId(context.Background(), disclosureId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdvisoriesAPI.GetAdvisoriesByDisclosureId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAdvisoriesByDisclosureId`: DisclosureAdvisory
	fmt.Fprintf(os.Stdout, "Response from `AdvisoriesAPI.GetAdvisoriesByDisclosureId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**disclosureId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdvisoriesByDisclosureIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DisclosureAdvisory**](DisclosureAdvisory.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

