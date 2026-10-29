# \SearchAPI

All URIs are relative to *https://app.gobugfree.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSearch**](SearchAPI.md#GetSearch) | **Get** /search | Search engagements, reports, documents, members, and organizations the user may open; q is the term, area selects the workspace (researcher, org:{org_slug}, or backoffice)



## GetSearch

> SearchResultList GetSearch(ctx).Execute()

Search engagements, reports, documents, members, and organizations the user may open; q is the term, area selects the workspace (researcher, org:{org_slug}, or backoffice)

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
	resp, r, err := apiClient.SearchAPI.GetSearch(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.GetSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSearch`: SearchResultList
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.GetSearch`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSearchRequest struct via the builder pattern


### Return type

[**SearchResultList**](SearchResultList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

