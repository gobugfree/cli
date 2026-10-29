# \ReportClaimAPI

All URIs are relative to *https://app.gobugfree.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetReportClaimByToken**](ReportClaimAPI.md#GetReportClaimByToken) | **Get** /report-claim/{token} | The anonymous report a claim link names, for a researcher whose account email matches; 404 for an invalid, used, or expired link, 403 with fields.reason when a program gate or the email refuses
[**PostReportClaimByToken**](ReportClaimAPI.md#PostReportClaimByToken) | **Post** /report-claim/{token} | Claim the anonymous report: the researcher becomes its owner when the email, the program gates, and the pending-report budget allow; 403 with fields.reason otherwise



## GetReportClaimByToken

> ReportClaim GetReportClaimByToken(ctx, token).Execute()

The anonymous report a claim link names, for a researcher whose account email matches; 404 for an invalid, used, or expired link, 403 with fields.reason when a program gate or the email refuses

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
	token := "token_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportClaimAPI.GetReportClaimByToken(context.Background(), token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportClaimAPI.GetReportClaimByToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReportClaimByToken`: ReportClaim
	fmt.Fprintf(os.Stdout, "Response from `ReportClaimAPI.GetReportClaimByToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportClaimByTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReportClaim**](ReportClaim.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostReportClaimByToken

> ReportClaim PostReportClaimByToken(ctx, token).Execute()

Claim the anonymous report: the researcher becomes its owner when the email, the program gates, and the pending-report budget allow; 403 with fields.reason otherwise

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
	token := "token_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportClaimAPI.PostReportClaimByToken(context.Background(), token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportClaimAPI.PostReportClaimByToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostReportClaimByToken`: ReportClaim
	fmt.Fprintf(os.Stdout, "Response from `ReportClaimAPI.PostReportClaimByToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostReportClaimByTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReportClaim**](ReportClaim.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

