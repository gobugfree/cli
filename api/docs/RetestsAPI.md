# \RetestsAPI

All URIs are relative to *https://app.gobugfree.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRetests**](RetestsAPI.md#GetRetests) | **Get** /retests | The researcher&#39;s retest queue: open assignments with due dates, and the decided history
[**PostRetestsStart**](RetestsAPI.md#PostRetestsStart) | **Post** /retests/{retest_id}/start | Start an assigned retest; only the assignee, 409 unless the retest is assigned
[**PostRetestsVerdict**](RetestsAPI.md#PostRetestsVerdict) | **Post** /retests/{retest_id}/verdict | Submit the retest verdict with mandatory reasoning: passed closes the report, failed reopens it; only the assignee, immutable once decided



## GetRetests

> GetRetests(ctx).Execute()

The researcher's retest queue: open assignments with due dates, and the decided history

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
	r, err := apiClient.RetestsAPI.GetRetests(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RetestsAPI.GetRetests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRetestsRequest struct via the builder pattern


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


## PostRetestsStart

> Retest PostRetestsStart(ctx, retestId).Execute()

Start an assigned retest; only the assignee, 409 unless the retest is assigned

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
	retestId := "retestId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RetestsAPI.PostRetestsStart(context.Background(), retestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RetestsAPI.PostRetestsStart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRetestsStart`: Retest
	fmt.Fprintf(os.Stdout, "Response from `RetestsAPI.PostRetestsStart`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**retestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostRetestsStartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Retest**](Retest.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRetestsVerdict

> Retest PostRetestsVerdict(ctx, retestId).Notes(notes).Outcome(outcome).Execute()

Submit the retest verdict with mandatory reasoning: passed closes the report, failed reopens it; only the assignee, immutable once decided

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
	retestId := "retestId_example" // string | 
	notes := "notes_example" // string | How the verifier tested the fix; recorded immutably with the verdict
	outcome := "outcome_example" // string | Whether the fix held: passed closes the report, failed reopens it

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RetestsAPI.PostRetestsVerdict(context.Background(), retestId).Notes(notes).Outcome(outcome).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RetestsAPI.PostRetestsVerdict``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRetestsVerdict`: Retest
	fmt.Fprintf(os.Stdout, "Response from `RetestsAPI.PostRetestsVerdict`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**retestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostRetestsVerdictRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **notes** | **string** | How the verifier tested the fix; recorded immutably with the verdict | 
 **outcome** | **string** | Whether the fix held: passed closes the report, failed reopens it | 

### Return type

[**Retest**](Retest.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

