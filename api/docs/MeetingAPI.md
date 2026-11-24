# \MeetingAPI

All URIs are relative to *https://app.gobugfree.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMeetingByToken**](MeetingAPI.md#GetMeetingByToken) | **Get** /meeting/{token} | The page behind a meeting invitation&#39;s link: the meeting, and the free times of one week at which every attendee is free, or the booked time; public, the token in the path is the credential
[**PostMeetingByToken**](MeetingAPI.md#PostMeetingByToken) | **Post** /meeting/{token} | Book one of the free times of a meeting invitation: the meeting is written into the host&#39;s calendar, which invites everyone; 409 when the time was just taken
[**PostMeetingCancel**](MeetingAPI.md#PostMeetingCancel) | **Post** /meeting/{token}/cancel | Cancel the time booked with a meeting invitation&#39;s link before the meeting begins, so another time can be picked with the same link



## GetMeetingByToken

> MeetingBookingPage GetMeetingByToken(ctx, token).Week(week).Execute()

The page behind a meeting invitation's link: the meeting, and the free times of one week at which every attendee is free, or the booked time; public, the token in the path is the credential

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
	week := "week_example" // string | Any day of the week to show (YYYY-MM-DD); the first bookable week when left out (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeetingAPI.GetMeetingByToken(context.Background(), token).Week(week).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeetingAPI.GetMeetingByToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMeetingByToken`: MeetingBookingPage
	fmt.Fprintf(os.Stdout, "Response from `MeetingAPI.GetMeetingByToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMeetingByTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **week** | **string** | Any day of the week to show (YYYY-MM-DD); the first bookable week when left out | 

### Return type

[**MeetingBookingPage**](MeetingBookingPage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMeetingByToken

> PostMeetingByToken(ctx, token).Email(email).Name(name).Start(start).Note(note).Execute()

Book one of the free times of a meeting invitation: the meeting is written into the host's calendar, which invites everyone; 409 when the time was just taken

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
	email := "email_example" // string | Address the calendar invitation is sent to
	name := "name_example" // string | Name of the person who books
	start := "start_example" // string | The picked start time, one of the free times of the page (RFC 3339)
	note := "note_example" // string | Anything the hosts should know (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MeetingAPI.PostMeetingByToken(context.Background(), token).Email(email).Name(name).Start(start).Note(note).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeetingAPI.PostMeetingByToken``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostMeetingByTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **email** | **string** | Address the calendar invitation is sent to | 
 **name** | **string** | Name of the person who books | 
 **start** | **string** | The picked start time, one of the free times of the page (RFC 3339) | 
 **note** | **string** | Anything the hosts should know | 

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


## PostMeetingCancel

> PostMeetingCancel(ctx, token).Execute()

Cancel the time booked with a meeting invitation's link before the meeting begins, so another time can be picked with the same link

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
	r, err := apiClient.MeetingAPI.PostMeetingCancel(context.Background(), token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeetingAPI.PostMeetingCancel``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostMeetingCancelRequest struct via the builder pattern


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

