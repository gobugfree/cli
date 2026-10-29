# \ListsAPI

All URIs are relative to *https://app.gobugfree.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostListsColumns**](ListsAPI.md#PostListsColumns) | **Post** /lists/columns | Save the caller&#39;s visible columns and their order for a list; reset&#x3D;1 restores the default
[**PostListsLayout**](ListsAPI.md#PostListsLayout) | **Post** /lists/layout | Save the caller&#39;s layout of a split-view list: single (the list alone), split (the list beside the open record), or board (the kanban)
[**PostListsViews**](ListsAPI.md#PostListsViews) | **Post** /lists/views | Save the current filters of a list as a named view of the caller (at most 20 per list, names unique per list)
[**PostListsViewsDelete**](ListsAPI.md#PostListsViewsDelete) | **Post** /lists/views/{id}/delete | Delete one of the caller&#39;s saved views; 404 when it is not theirs



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


## PostListsLayout

> PostListsLayout(ctx).Layout(layout).List(list).Return_(return_).Execute()

Save the caller's layout of a split-view list: single (the list alone), split (the list beside the open record), or board (the kanban)

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
	layout := "layout_example" // string | The layout: single (the list alone), split (the list beside the open record), or board (the kanban)
	list := "list_example" // string | The list key a feature registered its layouts under (for example triage-inbox or org-reports)
	return_ := "return__example" // string | The list URL to return to (browsers) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ListsAPI.PostListsLayout(context.Background()).Layout(layout).List(list).Return_(return_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.PostListsLayout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostListsLayoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **layout** | **string** | The layout: single (the list alone), split (the list beside the open record), or board (the kanban) | 
 **list** | **string** | The list key a feature registered its layouts under (for example triage-inbox or org-reports) | 
 **return_** | **string** | The list URL to return to (browsers) | 

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


## PostListsViews

> ListviewView PostListsViews(ctx).List(list).Name(name).Query(query).Return_(return_).Execute()

Save the current filters of a list as a named view of the caller (at most 20 per list, names unique per list)

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
	list := "list_example" // string | The list key a feature registered (for example triage-inbox)
	name := "name_example" // string | The view's name, unique per list
	query := "query_example" // string | The filter query string to save, without page or markers (optional)
	return_ := "return__example" // string | The list URL to return to (browsers) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.PostListsViews(context.Background()).List(list).Name(name).Query(query).Return_(return_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.PostListsViews``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostListsViews`: ListviewView
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.PostListsViews`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostListsViewsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | The list key a feature registered (for example triage-inbox) | 
 **name** | **string** | The view&#39;s name, unique per list | 
 **query** | **string** | The filter query string to save, without page or markers | 
 **return_** | **string** | The list URL to return to (browsers) | 

### Return type

[**ListviewView**](ListviewView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostListsViewsDelete

> PostListsViewsDelete(ctx, id).Execute()

Delete one of the caller's saved views; 404 when it is not theirs

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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ListsAPI.PostListsViewsDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.PostListsViewsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostListsViewsDeleteRequest struct via the builder pattern


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

