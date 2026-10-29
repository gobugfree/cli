# \MarkdownAPI

All URIs are relative to *https://app.gobugfree.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostMarkdownPreview**](MarkdownAPI.md#PostMarkdownPreview) | **Post** /markdown/preview | Render a Markdown document as the readers see it; nothing is stored



## PostMarkdownPreview

> map[string]interface{} PostMarkdownPreview(ctx).Content(content).Execute()

Render a Markdown document as the readers see it; nothing is stored

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
	content := "content_example" // string | The Markdown document to render

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarkdownAPI.PostMarkdownPreview(context.Background()).Content(content).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarkdownAPI.PostMarkdownPreview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMarkdownPreview`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `MarkdownAPI.PostMarkdownPreview`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMarkdownPreviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **content** | **string** | The Markdown document to render | 

### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

