# \LegacyAPI

All URIs are relative to *https://app.gobugfree.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLegacy**](LegacyAPI.md#GetLegacy) | **Get** /legacy | Where athena&#39;s home led: a redirect to the program list (JSON: the target)
[**GetLegacyByRest**](LegacyAPI.md#GetLegacyByRest) | **Get** /legacy/{rest} | Where an athena link leads: a redirect to the report its id maps to, or to the page that replaced athena&#39;s (JSON: the target, 404 for an unknown id or path)
[**GetLegacyEmbeddedSubmit**](LegacyAPI.md#GetLegacyEmbeddedSubmit) | **Get** /legacy/embedded/{program}/submit | Where athena&#39;s embedded submission form leads, without a session: the embedded form of the program when it is open; 404 for any other program (JSON: the target)
[**GetLegacyProgramsByProgram**](LegacyAPI.md#GetLegacyProgramsByProgram) | **Get** /legacy/programs/{program} | Where athena&#39;s public program link leads: for a signed-in viewer the full program page, for anyone else the public program page of a public engagement; 404 without a session for any other program (JSON: the target)
[**GetLegacyProgramsEmbeddedSubmit**](LegacyAPI.md#GetLegacyProgramsEmbeddedSubmit) | **Get** /legacy/programs/embedded/{program}/submit | Where athena&#39;s embedded submission form on its own site leads, without a session: the embedded form of the program when it is open; 404 for any other program (JSON: the target)



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


## GetLegacyEmbeddedSubmit

> LegacylinkLink GetLegacyEmbeddedSubmit(ctx, program).Execute()

Where athena's embedded submission form leads, without a session: the embedded form of the program when it is open; 404 for any other program (JSON: the target)

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LegacyAPI.GetLegacyEmbeddedSubmit(context.Background(), program).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LegacyAPI.GetLegacyEmbeddedSubmit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLegacyEmbeddedSubmit`: LegacylinkLink
	fmt.Fprintf(os.Stdout, "Response from `LegacyAPI.GetLegacyEmbeddedSubmit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**program** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLegacyEmbeddedSubmitRequest struct via the builder pattern


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


## GetLegacyProgramsByProgram

> LegacylinkLink GetLegacyProgramsByProgram(ctx, program).Execute()

Where athena's public program link leads: for a signed-in viewer the full program page, for anyone else the public program page of a public engagement; 404 without a session for any other program (JSON: the target)

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LegacyAPI.GetLegacyProgramsByProgram(context.Background(), program).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LegacyAPI.GetLegacyProgramsByProgram``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLegacyProgramsByProgram`: LegacylinkLink
	fmt.Fprintf(os.Stdout, "Response from `LegacyAPI.GetLegacyProgramsByProgram`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**program** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLegacyProgramsByProgramRequest struct via the builder pattern


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


## GetLegacyProgramsEmbeddedSubmit

> LegacylinkLink GetLegacyProgramsEmbeddedSubmit(ctx, program).Execute()

Where athena's embedded submission form on its own site leads, without a session: the embedded form of the program when it is open; 404 for any other program (JSON: the target)

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LegacyAPI.GetLegacyProgramsEmbeddedSubmit(context.Background(), program).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LegacyAPI.GetLegacyProgramsEmbeddedSubmit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLegacyProgramsEmbeddedSubmit`: LegacylinkLink
	fmt.Fprintf(os.Stdout, "Response from `LegacyAPI.GetLegacyProgramsEmbeddedSubmit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**program** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLegacyProgramsEmbeddedSubmitRequest struct via the builder pattern


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

