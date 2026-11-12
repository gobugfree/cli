# \EventInviteAPI

All URIs are relative to *https://app.gobugfree.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEventInviteByToken**](EventInviteAPI.md#GetEventInviteByToken) | **Get** /event-invite/{token} | The live hacking event an invite link names, with its window and its programs, served without a session too; 404 for an unknown, replaced, or revoked link and once the event ended
[**GetEventInviteLogo**](EventInviteAPI.md#GetEventInviteLogo) | **Get** /event-invite/{token}/logo | The logo of the event an invite link names, served without a session too; 404 without one and for an invalid link
[**PostEventInviteByToken**](EventInviteAPI.md#PostEventInviteByToken) | **Post** /event-invite/{token} | Accept the invite link: the caller joins the event&#39;s roster and is admitted to every bundled engagement; accepting twice changes nothing; without a session a browser is sent to the sign-in and joins on the way back, a JSON caller gets 401



## GetEventInviteByToken

> EventInvite GetEventInviteByToken(ctx, token).Execute()

The live hacking event an invite link names, with its window and its programs, served without a session too; 404 for an unknown, replaced, or revoked link and once the event ended

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
	token := "token_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventInviteAPI.GetEventInviteByToken(context.Background(), token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventInviteAPI.GetEventInviteByToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEventInviteByToken`: EventInvite
	fmt.Fprintf(os.Stdout, "Response from `EventInviteAPI.GetEventInviteByToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventInviteByTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EventInvite**](EventInvite.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventInviteLogo

> GetEventInviteLogo(ctx, token).Execute()

The logo of the event an invite link names, served without a session too; 404 without one and for an invalid link

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
	token := "token_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventInviteAPI.GetEventInviteLogo(context.Background(), token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventInviteAPI.GetEventInviteLogo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventInviteLogoRequest struct via the builder pattern


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


## PostEventInviteByToken

> EventInvite PostEventInviteByToken(ctx, token).Execute()

Accept the invite link: the caller joins the event's roster and is admitted to every bundled engagement; accepting twice changes nothing; without a session a browser is sent to the sign-in and joins on the way back, a JSON caller gets 401

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
	token := "token_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventInviteAPI.PostEventInviteByToken(context.Background(), token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventInviteAPI.PostEventInviteByToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostEventInviteByToken`: EventInvite
	fmt.Fprintf(os.Stdout, "Response from `EventInviteAPI.PostEventInviteByToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostEventInviteByTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EventInvite**](EventInvite.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

