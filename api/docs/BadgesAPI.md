# \BadgesAPI

All URIs are relative to *https://app.gobugfree.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBadges**](BadgesAPI.md#GetBadges) | **Get** /badges | Achievement badge catalog of the authenticated researcher: earned badges with their award date and visibility, locked badges with progress toward the threshold
[**PostBadgesVisibility**](BadgesAPI.md#PostBadgesVisibility) | **Post** /badges/{code}/visibility | Hide or show one of the researcher&#39;s own earned badges on the public profile



## GetBadges

> BadgeCatalogItemList GetBadges(ctx).Execute()

Achievement badge catalog of the authenticated researcher: earned badges with their award date and visibility, locked badges with progress toward the threshold

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
	resp, r, err := apiClient.BadgesAPI.GetBadges(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BadgesAPI.GetBadges``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBadges`: BadgeCatalogItemList
	fmt.Fprintf(os.Stdout, "Response from `BadgesAPI.GetBadges`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBadgesRequest struct via the builder pattern


### Return type

[**BadgeCatalogItemList**](BadgeCatalogItemList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostBadgesVisibility

> map[string]interface{} PostBadgesVisibility(ctx, code).Hidden(hidden).Execute()

Hide or show one of the researcher's own earned badges on the public profile

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
	code := "code_example" // string | 
	hidden := "hidden_example" // string | Whether the badge is hidden on the public profile

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BadgesAPI.PostBadgesVisibility(context.Background(), code).Hidden(hidden).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BadgesAPI.PostBadgesVisibility``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostBadgesVisibility`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `BadgesAPI.PostBadgesVisibility`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostBadgesVisibilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hidden** | **string** | Whether the badge is hidden on the public profile | 

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

