# \PublicAPI

All URIs are relative to *https://app.gobugfree.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPublicEngagements**](PublicAPI.md#GetPublicEngagements) | **Get** /public/engagements | Public engagement directory: every active or paused engagement as an anonymous viewer sees it; locked rows carry no scope and no rewards
[**GetPublicEngagementsLogo**](PublicAPI.md#GetPublicEngagementsLogo) | **Get** /public/engagements/{org_slug}/{engagement_slug}/logo | The organization logo of an engagement the public directory lists; 404 otherwise



## GetPublicEngagements

> EngagementPublicItemList GetPublicEngagements(ctx).Execute()

Public engagement directory: every active or paused engagement as an anonymous viewer sees it; locked rows carry no scope and no rewards

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
	resp, r, err := apiClient.PublicAPI.GetPublicEngagements(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.GetPublicEngagements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicEngagements`: EngagementPublicItemList
	fmt.Fprintf(os.Stdout, "Response from `PublicAPI.GetPublicEngagements`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicEngagementsRequest struct via the builder pattern


### Return type

[**EngagementPublicItemList**](EngagementPublicItemList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublicEngagementsLogo

> GetPublicEngagementsLogo(ctx, orgSlug, engagementSlug).Execute()

The organization logo of an engagement the public directory lists; 404 otherwise

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
	orgSlug := "orgSlug_example" // string | 
	engagementSlug := "engagementSlug_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PublicAPI.GetPublicEngagementsLogo(context.Background(), orgSlug, engagementSlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.GetPublicEngagementsLogo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 
**engagementSlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicEngagementsLogoRequest struct via the builder pattern


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

