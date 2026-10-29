# \AccountAPI

All URIs are relative to *https://app.gobugfree.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccountByRest**](AccountAPI.md#GetAccountByRest) | **Get** /account/{rest} | Every other old account link opens the program list (JSON: 404)
[**GetAccountProgramsSubmissionsByLegacyId**](AccountAPI.md#GetAccountProgramsSubmissionsByLegacyId) | **Get** /account/programs/{program}/submissions/{legacy_id} | Where an old report link that names its program leads: a redirect to the report, or the report list for an unknown id (JSON: the target, 404 for an unknown id)
[**GetAccountSubmissionsByLegacyId**](AccountAPI.md#GetAccountSubmissionsByLegacyId) | **Get** /account/submissions/{legacy_id} | Where an old report link leads: a redirect to the report, or the report list for an unknown id (JSON: the target, 404 for an unknown id)



## GetAccountByRest

> GetAccountByRest(ctx, rest).Execute()

Every other old account link opens the program list (JSON: 404)

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
	r, err := apiClient.AccountAPI.GetAccountByRest(context.Background(), rest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.GetAccountByRest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rest** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountByRestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetAccountProgramsSubmissionsByLegacyId

> LegacylinkLink GetAccountProgramsSubmissionsByLegacyId(ctx, program, legacyId).Execute()

Where an old report link that names its program leads: a redirect to the report, or the report list for an unknown id (JSON: the target, 404 for an unknown id)

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
	program := "program_example" // string | 
	legacyId := "legacyId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.GetAccountProgramsSubmissionsByLegacyId(context.Background(), program, legacyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.GetAccountProgramsSubmissionsByLegacyId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountProgramsSubmissionsByLegacyId`: LegacylinkLink
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.GetAccountProgramsSubmissionsByLegacyId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**program** | **string** |  | 
**legacyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountProgramsSubmissionsByLegacyIdRequest struct via the builder pattern


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


## GetAccountSubmissionsByLegacyId

> LegacylinkLink GetAccountSubmissionsByLegacyId(ctx, legacyId).Execute()

Where an old report link leads: a redirect to the report, or the report list for an unknown id (JSON: the target, 404 for an unknown id)

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
	legacyId := "legacyId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.GetAccountSubmissionsByLegacyId(context.Background(), legacyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.GetAccountSubmissionsByLegacyId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountSubmissionsByLegacyId`: LegacylinkLink
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.GetAccountSubmissionsByLegacyId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**legacyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountSubmissionsByLegacyIdRequest struct via the builder pattern


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

