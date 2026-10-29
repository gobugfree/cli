# \FeedbackAPI

All URIs are relative to *https://app.gobugfree.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFeedback**](FeedbackAPI.md#GetFeedback) | **Get** /feedback | The feedback overlay and the signed-in user&#39;s earlier feedback
[**PostFeedback**](FeedbackAPI.md#PostFeedback) | **Post** /feedback | Send product feedback



## GetFeedback

> FeedbackList GetFeedback(ctx).Execute()

The feedback overlay and the signed-in user's earlier feedback

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
	kind := "kind_example" // string | bug, idea, praise, or other
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
 **kind** | **string** | bug, idea, praise, or other | 
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

