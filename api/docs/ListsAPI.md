# \ListsAPI

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostListsColumns**](ListsAPI.md#PostListsColumns) | **Post** /lists/columns | Save the caller&#39;s visible columns and their order for a list; reset&#x3D;1 restores the default



## PostListsColumns

> PostListsColumns(ctx).List(list).Return_(return_).Move(move).Reset(reset).Execute()

Save the caller's visible columns and their order for a list; reset=1 restores the default

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
	list := "list_example" // string | The list key a feature registered its column catalog under (for example my-reports or audit-log)
	return_ := "return__example" // string | The list URL to return to
	move := "move_example" // string | Nudge one column: <key>:up or <key>:down (optional)
	reset := "reset_example" // string | 1 restores the default set (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ListsAPI.PostListsColumns(context.Background()).List(list).Return_(return_).Move(move).Reset(reset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.PostListsColumns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostListsColumnsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | The list key a feature registered its column catalog under (for example my-reports or audit-log) | 
 **return_** | **string** | The list URL to return to | 
 **move** | **string** | Nudge one column: &lt;key&gt;:up or &lt;key&gt;:down | 
 **reset** | **string** | 1 restores the default set | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

