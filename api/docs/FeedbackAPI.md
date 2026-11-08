# \FeedbackAPI

All URIs are relative to *https://app.gobugfree.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFeedback**](FeedbackAPI.md#GetFeedback) | **Get** /feedback | The signed-in user&#39;s feedback threads, and the feedback overlay for htmx
[**GetFeedbackByFeedbackId**](FeedbackAPI.md#GetFeedbackByFeedbackId) | **Get** /feedback/{feedback_id} | One of the signed-in user&#39;s feedback threads with its replies
[**GetFeedbackFilesByFileId**](FeedbackAPI.md#GetFeedbackFilesByFileId) | **Get** /feedback/{feedback_id}/files/{file_id} | One file of the signed-in user&#39;s feedback thread: an image inline, anything else as a download
[**PostFeedback**](FeedbackAPI.md#PostFeedback) | **Post** /feedback | Send product feedback
[**PostFeedbackDelete**](FeedbackAPI.md#PostFeedbackDelete) | **Post** /feedback/{feedback_id}/delete | Delete one of the signed-in user&#39;s feedback threads with its replies
[**PostFeedbackReply**](FeedbackAPI.md#PostFeedbackReply) | **Post** /feedback/{feedback_id}/reply | Reply in one of the signed-in user&#39;s feedback threads



## GetFeedback

> FeedbackList GetFeedback(ctx).Execute()

The signed-in user's feedback threads, and the feedback overlay for htmx

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
	resp, r, err := apiClient.FeedbackAPI.GetFeedback(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeedbackAPI.GetFeedback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFeedback`: FeedbackList
	fmt.Fprintf(os.Stdout, "Response from `FeedbackAPI.GetFeedback`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFeedbackRequest struct via the builder pattern


### Return type

[**FeedbackList**](FeedbackList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeedbackByFeedbackId

> FeedbackThread GetFeedbackByFeedbackId(ctx, feedbackId).Execute()

One of the signed-in user's feedback threads with its replies

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
	feedbackId := "feedbackId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeedbackAPI.GetFeedbackByFeedbackId(context.Background(), feedbackId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeedbackAPI.GetFeedbackByFeedbackId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFeedbackByFeedbackId`: FeedbackThread
	fmt.Fprintf(os.Stdout, "Response from `FeedbackAPI.GetFeedbackByFeedbackId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**feedbackId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFeedbackByFeedbackIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FeedbackThread**](FeedbackThread.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeedbackFilesByFileId

> GetFeedbackFilesByFileId(ctx, feedbackId, fileId).Execute()

One file of the signed-in user's feedback thread: an image inline, anything else as a download

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
	feedbackId := "feedbackId_example" // string | 
	fileId := "fileId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FeedbackAPI.GetFeedbackFilesByFileId(context.Background(), feedbackId, fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeedbackAPI.GetFeedbackFilesByFileId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**feedbackId** | **string** |  | 
**fileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFeedbackFilesByFileIdRequest struct via the builder pattern


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


## PostFeedback

> Feedback PostFeedback(ctx).Kind(kind).Message(message).Page(page).Rating(rating).Execute()

Send product feedback

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
	kind := "kind_example" // string | bug, idea, question, praise, or other
	message := "message_example" // string | The feedback text
	page := "page_example" // string | The app path the feedback refers to (optional)
	rating := "rating_example" // string | Optional rating from 1 to 5 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeedbackAPI.PostFeedback(context.Background()).Kind(kind).Message(message).Page(page).Rating(rating).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeedbackAPI.PostFeedback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostFeedback`: Feedback
	fmt.Fprintf(os.Stdout, "Response from `FeedbackAPI.PostFeedback`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostFeedbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kind** | **string** | bug, idea, question, praise, or other | 
 **message** | **string** | The feedback text | 
 **page** | **string** | The app path the feedback refers to | 
 **rating** | **string** | Optional rating from 1 to 5 | 

### Return type

[**Feedback**](Feedback.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFeedbackDelete

> PostFeedbackDelete(ctx, feedbackId).Execute()

Delete one of the signed-in user's feedback threads with its replies

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
	feedbackId := "feedbackId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FeedbackAPI.PostFeedbackDelete(context.Background(), feedbackId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeedbackAPI.PostFeedbackDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**feedbackId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostFeedbackDeleteRequest struct via the builder pattern


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


## PostFeedbackReply

> FeedbackThread PostFeedbackReply(ctx, feedbackId).Body(body).Execute()

Reply in one of the signed-in user's feedback threads

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
	feedbackId := "feedbackId_example" // string | 
	body := "body_example" // string | The reply text

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeedbackAPI.PostFeedbackReply(context.Background(), feedbackId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeedbackAPI.PostFeedbackReply``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostFeedbackReply`: FeedbackThread
	fmt.Fprintf(os.Stdout, "Response from `FeedbackAPI.PostFeedbackReply`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**feedbackId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostFeedbackReplyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** | The reply text | 

### Return type

[**FeedbackThread**](FeedbackThread.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

