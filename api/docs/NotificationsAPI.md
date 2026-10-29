# \NotificationsAPI

All URIs are relative to *https://app.gobugfree.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNotifications**](NotificationsAPI.md#GetNotifications) | **Get** /notifications | The user&#39;s notifications, newest first, a page at a time
[**GetNotificationsByNotificationId**](NotificationsAPI.md#GetNotificationsByNotificationId) | **Get** /notifications/{notification_id} | Open one notification: marks it read and forwards to its entity; 404 for other users&#39; notifications (ADR-0009)
[**GetNotificationsPreferences**](NotificationsAPI.md#GetNotificationsPreferences) | **Get** /notifications/preferences | The user&#39;s email notification switches, one per notification kind
[**PostNotificationsPreferences**](NotificationsAPI.md#PostNotificationsPreferences) | **Post** /notifications/preferences | Switch the user&#39;s email notifications per kind
[**PostNotificationsRead**](NotificationsAPI.md#PostNotificationsRead) | **Post** /notifications/{notification_id}/read | Mark one notification of the user as read; 404 for other users&#39; notifications
[**PostNotificationsReadAll**](NotificationsAPI.md#PostNotificationsReadAll) | **Post** /notifications/read-all | Mark every notification of the user as read



## GetNotifications

> MessageNotificationPage GetNotifications(ctx).Before(before).Execute()

The user's notifications, newest first, a page at a time



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
	before := "before_example" // string | Notification id: list the page older than it (the previous answer's next) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.GetNotifications(context.Background()).Before(before).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.GetNotifications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNotifications`: MessageNotificationPage
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.GetNotifications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **before** | **string** | Notification id: list the page older than it (the previous answer&#39;s next) | 

### Return type

[**MessageNotificationPage**](MessageNotificationPage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotificationsByNotificationId

> GetNotificationsByNotificationId(ctx, notificationId).Execute()

Open one notification: marks it read and forwards to its entity; 404 for other users' notifications (ADR-0009)

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
	notificationId := "notificationId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NotificationsAPI.GetNotificationsByNotificationId(context.Background(), notificationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.GetNotificationsByNotificationId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationsByNotificationIdRequest struct via the builder pattern


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


## GetNotificationsPreferences

> MessageEmailPreferenceList GetNotificationsPreferences(ctx).Execute()

The user's email notification switches, one per notification kind

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
	resp, r, err := apiClient.NotificationsAPI.GetNotificationsPreferences(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.GetNotificationsPreferences``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNotificationsPreferences`: MessageEmailPreferenceList
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.GetNotificationsPreferences`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationsPreferencesRequest struct via the builder pattern


### Return type

[**MessageEmailPreferenceList**](MessageEmailPreferenceList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostNotificationsPreferences

> PostNotificationsPreferences(ctx).EngagementAccessRequested(engagementAccessRequested).EngagementInvitationResponded(engagementInvitationResponded).EngagementRulesChanged(engagementRulesChanged).EngagementStatusChanged(engagementStatusChanged).OrgReportAwaitingFeedback(orgReportAwaitingFeedback).OrgReportCommentAdded(orgReportCommentAdded).OrgReportStatusChanged(orgReportStatusChanged).OrgReportSubmitted(orgReportSubmitted).OrganizationJoinDecided(organizationJoinDecided).OrganizationJoinRequested(organizationJoinRequested).ReportCommentAdded(reportCommentAdded).ReportInfoRequested(reportInfoRequested).ReportStatusChanged(reportStatusChanged).ScanReportPublished(scanReportPublished).VersionDecided(versionDecided).Execute()

Switch the user's email notifications per kind

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
	engagementAccessRequested := "engagementAccessRequested_example" // string | Organization members: send \\\"on\\\" for email on program access requests, omit to switch it off (optional)
	engagementInvitationResponded := "engagementInvitationResponded_example" // string | Organization members: send \\\"on\\\" for email on answers to program invitations, omit to switch it off (optional)
	engagementRulesChanged := "engagementRulesChanged_example" // string | Send \\\"on\\\" for email on program rule changes, omit to switch it off (optional)
	engagementStatusChanged := "engagementStatusChanged_example" // string | Send \\\"on\\\" for email on program status changes, omit to switch it off (optional)
	orgReportAwaitingFeedback := "orgReportAwaitingFeedback_example" // string | Organization members: send \\\"on\\\" for email on reports awaiting your decision, omit to switch it off (optional)
	orgReportCommentAdded := "orgReportCommentAdded_example" // string | Organization members: send \\\"on\\\" for email on messages to you on program reports, omit to switch it off (optional)
	orgReportStatusChanged := "orgReportStatusChanged_example" // string | Organization members: send \\\"on\\\" for email on status changes of program reports, omit to switch it off (optional)
	orgReportSubmitted := "orgReportSubmitted_example" // string | Organization members: send \\\"on\\\" for email on new reports to your programs, omit to switch it off (optional)
	organizationJoinDecided := "organizationJoinDecided_example" // string | Send \\\"on\\\" for email on decisions on your requests to join an organization, omit to switch it off (optional)
	organizationJoinRequested := "organizationJoinRequested_example" // string | Organization members: send \\\"on\\\" for email on join requests from your verified email domains, omit to switch it off (optional)
	reportCommentAdded := "reportCommentAdded_example" // string | Send \\\"on\\\" for email on new report comments, omit to switch it off (optional)
	reportInfoRequested := "reportInfoRequested_example" // string | Send \\\"on\\\" for email on information requests, omit to switch it off (optional)
	reportStatusChanged := "reportStatusChanged_example" // string | Send \\\"on\\\" for email on report status changes, omit to switch it off (optional)
	scanReportPublished := "scanReportPublished_example" // string | Organization members: send \\\"on\\\" for email on new scan reports, omit to switch it off (optional)
	versionDecided := "versionDecided_example" // string | Send \\\"on\\\" for email on content review decisions (a program, the organization profile, an asset), omit to switch it off (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NotificationsAPI.PostNotificationsPreferences(context.Background()).EngagementAccessRequested(engagementAccessRequested).EngagementInvitationResponded(engagementInvitationResponded).EngagementRulesChanged(engagementRulesChanged).EngagementStatusChanged(engagementStatusChanged).OrgReportAwaitingFeedback(orgReportAwaitingFeedback).OrgReportCommentAdded(orgReportCommentAdded).OrgReportStatusChanged(orgReportStatusChanged).OrgReportSubmitted(orgReportSubmitted).OrganizationJoinDecided(organizationJoinDecided).OrganizationJoinRequested(organizationJoinRequested).ReportCommentAdded(reportCommentAdded).ReportInfoRequested(reportInfoRequested).ReportStatusChanged(reportStatusChanged).ScanReportPublished(scanReportPublished).VersionDecided(versionDecided).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.PostNotificationsPreferences``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostNotificationsPreferencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **engagementAccessRequested** | **string** | Organization members: send \\\&quot;on\\\&quot; for email on program access requests, omit to switch it off | 
 **engagementInvitationResponded** | **string** | Organization members: send \\\&quot;on\\\&quot; for email on answers to program invitations, omit to switch it off | 
 **engagementRulesChanged** | **string** | Send \\\&quot;on\\\&quot; for email on program rule changes, omit to switch it off | 
 **engagementStatusChanged** | **string** | Send \\\&quot;on\\\&quot; for email on program status changes, omit to switch it off | 
 **orgReportAwaitingFeedback** | **string** | Organization members: send \\\&quot;on\\\&quot; for email on reports awaiting your decision, omit to switch it off | 
 **orgReportCommentAdded** | **string** | Organization members: send \\\&quot;on\\\&quot; for email on messages to you on program reports, omit to switch it off | 
 **orgReportStatusChanged** | **string** | Organization members: send \\\&quot;on\\\&quot; for email on status changes of program reports, omit to switch it off | 
 **orgReportSubmitted** | **string** | Organization members: send \\\&quot;on\\\&quot; for email on new reports to your programs, omit to switch it off | 
 **organizationJoinDecided** | **string** | Send \\\&quot;on\\\&quot; for email on decisions on your requests to join an organization, omit to switch it off | 
 **organizationJoinRequested** | **string** | Organization members: send \\\&quot;on\\\&quot; for email on join requests from your verified email domains, omit to switch it off | 
 **reportCommentAdded** | **string** | Send \\\&quot;on\\\&quot; for email on new report comments, omit to switch it off | 
 **reportInfoRequested** | **string** | Send \\\&quot;on\\\&quot; for email on information requests, omit to switch it off | 
 **reportStatusChanged** | **string** | Send \\\&quot;on\\\&quot; for email on report status changes, omit to switch it off | 
 **scanReportPublished** | **string** | Organization members: send \\\&quot;on\\\&quot; for email on new scan reports, omit to switch it off | 
 **versionDecided** | **string** | Send \\\&quot;on\\\&quot; for email on content review decisions (a program, the organization profile, an asset), omit to switch it off | 

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


## PostNotificationsRead

> PostNotificationsRead(ctx, notificationId).Execute()

Mark one notification of the user as read; 404 for other users' notifications

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
	notificationId := "notificationId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NotificationsAPI.PostNotificationsRead(context.Background(), notificationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.PostNotificationsRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostNotificationsReadRequest struct via the builder pattern


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


## PostNotificationsReadAll

> PostNotificationsReadAll(ctx).Execute()

Mark every notification of the user as read

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
	r, err := apiClient.NotificationsAPI.PostNotificationsReadAll(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.PostNotificationsReadAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostNotificationsReadAllRequest struct via the builder pattern


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

