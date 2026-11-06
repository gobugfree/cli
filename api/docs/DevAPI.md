# \DevAPI

All URIs are relative to *https://app.gobugfree.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDevBrandSiteBySite**](DevAPI.md#GetDevBrandSiteBySite) | **Get** /dev/brand-site/{site} | Development-only stub brand site homepage (acme, plain, ico-only) the website brand extraction can fetch locally
[**GetDevBrandSiteBySiteAndFile**](DevAPI.md#GetDevBrandSiteBySiteAndFile) | **Get** /dev/brand-site/{site}/{file} | Development-only stub brand site asset, addressed by its file stem
[**GetDevLlmV1Models**](DevAPI.md#GetDevLlmV1Models) | **Get** /dev/llm/v1/models | Development-only stub worker: the models it answers to, in the OpenAI format
[**PostDevLlmV1ChatCompletions**](DevAPI.md#PostDevLlmV1ChatCompletions) | **Post** /dev/llm/v1/chat/completions | Development-only stub worker: streams the start of the last user message back in the OpenAI format
[**PostDevLlmV1Embeddings**](DevAPI.md#PostDevLlmV1Embeddings) | **Post** /dev/llm/v1/embeddings | Development-only stub worker: hashed bag-of-words embeddings in the OpenAI format



## GetDevBrandSiteBySite

> GetDevBrandSiteBySite(ctx, site).Execute()

Development-only stub brand site homepage (acme, plain, ico-only) the website brand extraction can fetch locally

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
	site := "site_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DevAPI.GetDevBrandSiteBySite(context.Background(), site).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevAPI.GetDevBrandSiteBySite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**site** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDevBrandSiteBySiteRequest struct via the builder pattern


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


## GetDevBrandSiteBySiteAndFile

> GetDevBrandSiteBySiteAndFile(ctx, site, file).Execute()

Development-only stub brand site asset, addressed by its file stem

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
	site := "site_example" // string | 
	file := "file_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DevAPI.GetDevBrandSiteBySiteAndFile(context.Background(), site, file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevAPI.GetDevBrandSiteBySiteAndFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**site** | **string** |  | 
**file** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDevBrandSiteBySiteAndFileRequest struct via the builder pattern


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


## GetDevLlmV1Models

> GetDevLlmV1Models(ctx).Execute()

Development-only stub worker: the models it answers to, in the OpenAI format

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
	r, err := apiClient.DevAPI.GetDevLlmV1Models(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevAPI.GetDevLlmV1Models``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDevLlmV1ModelsRequest struct via the builder pattern


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


## PostDevLlmV1ChatCompletions

> PostDevLlmV1ChatCompletions(ctx).Execute()

Development-only stub worker: streams the start of the last user message back in the OpenAI format

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
	r, err := apiClient.DevAPI.PostDevLlmV1ChatCompletions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevAPI.PostDevLlmV1ChatCompletions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostDevLlmV1ChatCompletionsRequest struct via the builder pattern


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


## PostDevLlmV1Embeddings

> PostDevLlmV1Embeddings(ctx).Execute()

Development-only stub worker: hashed bag-of-words embeddings in the OpenAI format

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
	r, err := apiClient.DevAPI.PostDevLlmV1Embeddings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevAPI.PostDevLlmV1Embeddings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostDevLlmV1EmbeddingsRequest struct via the builder pattern


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

