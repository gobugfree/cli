# \OrganizationInvitationsAPI

All URIs are relative to *https://app.gobugfree.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationInvitations**](OrganizationInvitationsAPI.md#GetOrganizationInvitations) | **Get** /organization-invitations | The caller&#39;s pending organization invitations, with what each would grant
[**PostOrganizationInvitationsAccept**](OrganizationInvitationsAPI.md#PostOrganizationInvitationsAccept) | **Post** /organization-invitations/{member_id}/accept | Accept one of the caller&#39;s pending invitations: the membership becomes active and joins the organization&#39;s groups; 404 for another user&#39;s or a decided invitation
[**PostOrganizationInvitationsDecline**](OrganizationInvitationsAPI.md#PostOrganizationInvitationsDecline) | **Post** /organization-invitations/{member_id}/decline | Decline one of the caller&#39;s pending invitations: the invitation is deleted; 404 for another user&#39;s or a decided invitation



## GetOrganizationInvitations

> OrganizationInvitationList GetOrganizationInvitations(ctx).Execute()

The caller's pending organization invitations, with what each would grant

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
	resp, r, err := apiClient.OrganizationInvitationsAPI.GetOrganizationInvitations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationInvitationsAPI.GetOrganizationInvitations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationInvitations`: OrganizationInvitationList
	fmt.Fprintf(os.Stdout, "Response from `OrganizationInvitationsAPI.GetOrganizationInvitations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationInvitationsRequest struct via the builder pattern


### Return type

[**OrganizationInvitationList**](OrganizationInvitationList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOrganizationInvitationsAccept

> OrganizationAcceptedInvitation PostOrganizationInvitationsAccept(ctx, memberId).Execute()

Accept one of the caller's pending invitations: the membership becomes active and joins the organization's groups; 404 for another user's or a decided invitation

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
	memberId := "memberId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationInvitationsAPI.PostOrganizationInvitationsAccept(context.Background(), memberId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationInvitationsAPI.PostOrganizationInvitationsAccept``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostOrganizationInvitationsAccept`: OrganizationAcceptedInvitation
	fmt.Fprintf(os.Stdout, "Response from `OrganizationInvitationsAPI.PostOrganizationInvitationsAccept`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memberId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrganizationInvitationsAcceptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrganizationAcceptedInvitation**](OrganizationAcceptedInvitation.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOrganizationInvitationsDecline

> PostOrganizationInvitationsDecline(ctx, memberId).Execute()

Decline one of the caller's pending invitations: the invitation is deleted; 404 for another user's or a decided invitation

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
	memberId := "memberId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationInvitationsAPI.PostOrganizationInvitationsDecline(context.Background(), memberId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationInvitationsAPI.PostOrganizationInvitationsDecline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memberId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrganizationInvitationsDeclineRequest struct via the builder pattern


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

