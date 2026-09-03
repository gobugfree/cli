# \DocsAPI

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDocs**](DocsAPI.md#GetDocs) | **Get** /docs | Interactive API reference: the caller&#39;s OpenAPI document rendered for people, with try-it requests; JSON callers get the document locations



## GetDocs

> ApidocReference GetDocs(ctx).Execute()

Interactive API reference: the caller's OpenAPI document rendered for people, with try-it requests; JSON callers get the document locations

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
	resp, r, err := apiClient.DocsAPI.GetDocs(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocsAPI.GetDocs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDocs`: ApidocReference
	fmt.Fprintf(os.Stdout, "Response from `DocsAPI.GetDocs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDocsRequest struct via the builder pattern


### Return type

[**ApidocReference**](ApidocReference.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

