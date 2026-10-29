# \EmbedAPI

All URIs are relative to *https://app.gobugfree.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEmbedByOrgSlugAndEngagementSlug**](EmbedAPI.md#GetEmbedByOrgSlugAndEngagementSlug) | **Get** /embed/{org_slug}/{engagement_slug} | The embedded vulnerability submission form of a public VDP, frameable by its allowed origins; 404 for any other engagement
[**GetEmbedResizeJs**](EmbedAPI.md#GetEmbedResizeJs) | **Get** /embed/resize.js | The script a customer page loads with the embed snippet; it sizes the frame to the form
[**PostEmbedByOrgSlugAndEngagementSlug**](EmbedAPI.md#PostEmbedByOrgSlugAndEngagementSlug) | **Post** /embed/{org_slug}/{engagement_slug} | Submit a report anonymously through the embedded form with its proof of work; 429 when the per-IP budget is spent



## GetEmbedByOrgSlugAndEngagementSlug

> GetEmbedByOrgSlugAndEngagementSlug(ctx, orgSlug, engagementSlug).Execute()

The embedded vulnerability submission form of a public VDP, frameable by its allowed origins; 404 for any other engagement

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
	r, err := apiClient.EmbedAPI.GetEmbedByOrgSlugAndEngagementSlug(context.Background(), orgSlug, engagementSlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmbedAPI.GetEmbedByOrgSlugAndEngagementSlug``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetEmbedByOrgSlugAndEngagementSlugRequest struct via the builder pattern


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


## GetEmbedResizeJs

> GetEmbedResizeJs(ctx).Execute()

The script a customer page loads with the embed snippet; it sizes the frame to the form

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
	r, err := apiClient.EmbedAPI.GetEmbedResizeJs(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmbedAPI.GetEmbedResizeJs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEmbedResizeJsRequest struct via the builder pattern


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


## PostEmbedByOrgSlugAndEngagementSlug

> PostEmbedByOrgSlugAndEngagementSlug(ctx, orgSlug, engagementSlug).Challenge(challenge).Description(description).Severity(severity).Signature(signature).Steps(steps).Terms(terms).Title(title).Email(email).Impact(impact).Nonce(nonce).Recommendation(recommendation).ScopeEntryId(scopeEntryId).Weakness(weakness).Website(website).Execute()

Submit a report anonymously through the embedded form with its proof of work; 429 when the per-IP budget is spent

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
	challenge := "challenge_example" // string | 
	description := "description_example" // string | 
	severity := "severity_example" // string | 
	signature := "signature_example" // string | 
	steps := "steps_example" // string | 
	terms := "terms_example" // string | 
	title := "title_example" // string | 
	email := "email_example" // string |  (optional)
	impact := "impact_example" // string |  (optional)
	nonce := "nonce_example" // string |  (optional)
	recommendation := "recommendation_example" // string |  (optional)
	scopeEntryId := "scopeEntryId_example" // string |  (optional)
	weakness := "weakness_example" // string |  (optional)
	website := "website_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EmbedAPI.PostEmbedByOrgSlugAndEngagementSlug(context.Background(), orgSlug, engagementSlug).Challenge(challenge).Description(description).Severity(severity).Signature(signature).Steps(steps).Terms(terms).Title(title).Email(email).Impact(impact).Nonce(nonce).Recommendation(recommendation).ScopeEntryId(scopeEntryId).Weakness(weakness).Website(website).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmbedAPI.PostEmbedByOrgSlugAndEngagementSlug``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostEmbedByOrgSlugAndEngagementSlugRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **challenge** | **string** |  | 
 **description** | **string** |  | 
 **severity** | **string** |  | 
 **signature** | **string** |  | 
 **steps** | **string** |  | 
 **terms** | **string** |  | 
 **title** | **string** |  | 
 **email** | **string** |  | 
 **impact** | **string** |  | 
 **nonce** | **string** |  | 
 **recommendation** | **string** |  | 
 **scopeEntryId** | **string** |  | 
 **weakness** | **string** |  | 
 **website** | **string** |  | 

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

