# \EngagementsAPI

All URIs are relative to *https://app.gobugfree.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEngagements**](EngagementsAPI.md#GetEngagements) | **Get** /engagements | Engagement directory: active and paused engagements visible to the authenticated user, the ones a running event bundles first, filterable by kind, running event (event&#x3D;live), industry, scope type (scope), and minimum bounty (bounty), paginated with page
[**GetEngagementsByOrgSlugAndEngagementSlug**](EngagementsAPI.md#GetEngagementsByOrgSlugAndEngagementSlug) | **Get** /engagements/{org_slug}/{engagement_slug} | Engagement detail; public engagements for any user, full details of other engagements for members only (ADR-0009); non-members of a private or invite_only engagement, and pending invitees, get the locked teaser with locked true and access, everyone else 404
[**GetEngagementsFilesByName**](EngagementsAPI.md#GetEngagementsFilesByName) | **Get** /engagements/{org_slug}/{engagement_slug}/files/{name} | Download a program file of the published version: the NDA once the verification gate is passed, an attachment once cleared to test (verification, NDA, and rules accepted); 404 otherwise
[**GetEngagementsMessageProxy**](EngagementsAPI.md#GetEngagementsMessageProxy) | **Get** /engagements/{org_slug}/{engagement_slug}/message-proxy | Authentication codes forwarded from the engagement&#39;s 2FA-protected test accounts within the last 15 minutes; requires testing access and accepted rules
[**GetEngagementsSubmit**](EngagementsAPI.md#GetEngagementsSubmit) | **Get** /engagements/{org_slug}/{engagement_slug}/submit | The report submission form of one engagement: the caller&#39;s pending-report budget, the weakness catalog, the severity presets with their CVSS vectors, the reward bands, and the caller&#39;s recent accuracy
[**PostEngagementsAccessRequest**](EngagementsAPI.md#PostEngagementsAccessRequest) | **Post** /engagements/{org_slug}/{engagement_slug}/access-request | Request access to a listed private engagement; the test account fields are required when the engagement asks for test accounts; 404 unless the engagement takes requests from the caller, 409 on a second request or a pending invitation
[**PostEngagementsInvitationAccept**](EngagementsAPI.md#PostEngagementsInvitationAccept) | **Post** /engagements/{org_slug}/{engagement_slug}/invitation/accept | Accept the caller&#39;s pending invitation and join the engagement; 404 without one
[**PostEngagementsInvitationDecline**](EngagementsAPI.md#PostEngagementsInvitationDecline) | **Post** /engagements/{org_slug}/{engagement_slug}/invitation/decline | Decline the caller&#39;s pending invitation; 404 without one
[**PostEngagementsNdaAccept**](EngagementsAPI.md#PostEngagementsNdaAccept) | **Post** /engagements/{org_slug}/{engagement_slug}/nda/accept | Accept the NDA of the engagement&#39;s published version; required before the rules of engagement; 409 without an NDA or while the verification gate applies; a new NDA file asks again
[**PostEngagementsRulesAccept**](EngagementsAPI.md#PostEngagementsRulesAccept) | **Post** /engagements/{org_slug}/{engagement_slug}/rules/accept | Accept the engagement&#39;s published rules of engagement; required before submitting reports, re-required after a new version publishes
[**PostEngagementsSubmit**](EngagementsAPI.md#PostEngagementsSubmit) | **Post** /engagements/{org_slug}/{engagement_slug}/submit | Create and submit a vulnerability report in one call; members and, for active public engagements, fully onboarded researchers. 400 names the missing fields; 409 pending_budget_exceeded keeps the content as the draft named by draftId. To submit an existing draft use postReportsSubmit
[**PostEngagementsSubmitDraft**](EngagementsAPI.md#PostEngagementsSubmitDraft) | **Post** /engagements/{org_slug}/{engagement_slug}/submit/draft | Save a private draft report; no completeness validation, only a title is required. Submit it later with postReportsSubmit



## GetEngagements

> EngagementListItemList GetEngagements(ctx).Execute()

Engagement directory: active and paused engagements visible to the authenticated user, the ones a running event bundles first, filterable by kind, running event (event=live), industry, scope type (scope), and minimum bounty (bounty), paginated with page

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
	resp, r, err := apiClient.EngagementsAPI.GetEngagements(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngagementsAPI.GetEngagements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEngagements`: EngagementListItemList
	fmt.Fprintf(os.Stdout, "Response from `EngagementsAPI.GetEngagements`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEngagementsRequest struct via the builder pattern


### Return type

[**EngagementListItemList**](EngagementListItemList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEngagementsByOrgSlugAndEngagementSlug

> EngagementDetail GetEngagementsByOrgSlugAndEngagementSlug(ctx, orgSlug, engagementSlug).Execute()

Engagement detail; public engagements for any user, full details of other engagements for members only (ADR-0009); non-members of a private or invite_only engagement, and pending invitees, get the locked teaser with locked true and access, everyone else 404

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
	resp, r, err := apiClient.EngagementsAPI.GetEngagementsByOrgSlugAndEngagementSlug(context.Background(), orgSlug, engagementSlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngagementsAPI.GetEngagementsByOrgSlugAndEngagementSlug``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEngagementsByOrgSlugAndEngagementSlug`: EngagementDetail
	fmt.Fprintf(os.Stdout, "Response from `EngagementsAPI.GetEngagementsByOrgSlugAndEngagementSlug`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 
**engagementSlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEngagementsByOrgSlugAndEngagementSlugRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EngagementDetail**](EngagementDetail.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEngagementsFilesByName

> GetEngagementsFilesByName(ctx, orgSlug, engagementSlug, name).Execute()

Download a program file of the published version: the NDA once the verification gate is passed, an attachment once cleared to test (verification, NDA, and rules accepted); 404 otherwise

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
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EngagementsAPI.GetEngagementsFilesByName(context.Background(), orgSlug, engagementSlug, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngagementsAPI.GetEngagementsFilesByName``: %v\n", err)
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
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEngagementsFilesByNameRequest struct via the builder pattern


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


## GetEngagementsMessageProxy

> EngagementProxyMessageList GetEngagementsMessageProxy(ctx, orgSlug, engagementSlug).Execute()

Authentication codes forwarded from the engagement's 2FA-protected test accounts within the last 15 minutes; requires testing access and accepted rules

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
	resp, r, err := apiClient.EngagementsAPI.GetEngagementsMessageProxy(context.Background(), orgSlug, engagementSlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngagementsAPI.GetEngagementsMessageProxy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEngagementsMessageProxy`: EngagementProxyMessageList
	fmt.Fprintf(os.Stdout, "Response from `EngagementsAPI.GetEngagementsMessageProxy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 
**engagementSlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEngagementsMessageProxyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EngagementProxyMessageList**](EngagementProxyMessageList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEngagementsSubmit

> GetEngagementsSubmit(ctx, orgSlug, engagementSlug).Submitted(submitted).Execute()

The report submission form of one engagement: the caller's pending-report budget, the weakness catalog, the severity presets with their CVSS vectors, the reward bands, and the caller's recent accuracy



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
	submitted := "submitted_example" // string | The id of a report the caller just submitted to this engagement, shown as a confirmation (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EngagementsAPI.GetEngagementsSubmit(context.Background(), orgSlug, engagementSlug).Submitted(submitted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngagementsAPI.GetEngagementsSubmit``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetEngagementsSubmitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **submitted** | **string** | The id of a report the caller just submitted to this engagement, shown as a confirmation | 

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


## PostEngagementsAccessRequest

> AdmissionState PostEngagementsAccessRequest(ctx, orgSlug, engagementSlug).MobileNumber(mobileNumber).TestEmail1(testEmail1).TestEmail2(testEmail2).TestEmail3(testEmail3).Execute()

Request access to a listed private engagement; the test account fields are required when the engagement asks for test accounts; 404 unless the engagement takes requests from the caller, 409 on a second request or a pending invitation

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
	mobileNumber := "mobileNumber_example" // string | Mobile number for the test accounts in international format; required when the engagement asks for test accounts (optional)
	testEmail1 := "testEmail1_example" // string | First test account email; required when the engagement asks for test accounts (optional)
	testEmail2 := "testEmail2_example" // string | Second test account email; required when the engagement asks for test accounts (optional)
	testEmail3 := "testEmail3_example" // string | Third test account email; required when the engagement asks for test accounts (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EngagementsAPI.PostEngagementsAccessRequest(context.Background(), orgSlug, engagementSlug).MobileNumber(mobileNumber).TestEmail1(testEmail1).TestEmail2(testEmail2).TestEmail3(testEmail3).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngagementsAPI.PostEngagementsAccessRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostEngagementsAccessRequest`: AdmissionState
	fmt.Fprintf(os.Stdout, "Response from `EngagementsAPI.PostEngagementsAccessRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 
**engagementSlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostEngagementsAccessRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **mobileNumber** | **string** | Mobile number for the test accounts in international format; required when the engagement asks for test accounts | 
 **testEmail1** | **string** | First test account email; required when the engagement asks for test accounts | 
 **testEmail2** | **string** | Second test account email; required when the engagement asks for test accounts | 
 **testEmail3** | **string** | Third test account email; required when the engagement asks for test accounts | 

### Return type

[**AdmissionState**](AdmissionState.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostEngagementsInvitationAccept

> AdmissionState PostEngagementsInvitationAccept(ctx, orgSlug, engagementSlug).Execute()

Accept the caller's pending invitation and join the engagement; 404 without one

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
	resp, r, err := apiClient.EngagementsAPI.PostEngagementsInvitationAccept(context.Background(), orgSlug, engagementSlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngagementsAPI.PostEngagementsInvitationAccept``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostEngagementsInvitationAccept`: AdmissionState
	fmt.Fprintf(os.Stdout, "Response from `EngagementsAPI.PostEngagementsInvitationAccept`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 
**engagementSlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostEngagementsInvitationAcceptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AdmissionState**](AdmissionState.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostEngagementsInvitationDecline

> AdmissionState PostEngagementsInvitationDecline(ctx, orgSlug, engagementSlug).Execute()

Decline the caller's pending invitation; 404 without one

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
	resp, r, err := apiClient.EngagementsAPI.PostEngagementsInvitationDecline(context.Background(), orgSlug, engagementSlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngagementsAPI.PostEngagementsInvitationDecline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostEngagementsInvitationDecline`: AdmissionState
	fmt.Fprintf(os.Stdout, "Response from `EngagementsAPI.PostEngagementsInvitationDecline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 
**engagementSlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostEngagementsInvitationDeclineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AdmissionState**](AdmissionState.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostEngagementsNdaAccept

> PostEngagementsNdaAccept(ctx, orgSlug, engagementSlug).Accept(accept).Execute()

Accept the NDA of the engagement's published version; required before the rules of engagement; 409 without an NDA or while the verification gate applies; a new NDA file asks again

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
	accept := "accept_example" // string | on to accept the NDA of the published version

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EngagementsAPI.PostEngagementsNdaAccept(context.Background(), orgSlug, engagementSlug).Accept(accept).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngagementsAPI.PostEngagementsNdaAccept``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostEngagementsNdaAcceptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accept** | **string** | on to accept the NDA of the published version | 

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


## PostEngagementsRulesAccept

> PostEngagementsRulesAccept(ctx, orgSlug, engagementSlug).Execute()

Accept the engagement's published rules of engagement; required before submitting reports, re-required after a new version publishes

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
	r, err := apiClient.EngagementsAPI.PostEngagementsRulesAccept(context.Background(), orgSlug, engagementSlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngagementsAPI.PostEngagementsRulesAccept``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostEngagementsRulesAcceptRequest struct via the builder pattern


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


## PostEngagementsSubmit

> Report PostEngagementsSubmit(ctx, orgSlug, engagementSlug).Content(content).Ac(ac).Av(av).Avail(avail).Conf(conf).E(e).Integ(integ).Pr(pr).Rc(rc).Reason(reason).Rl(rl).Scope(scope).ScopeEntryId(scopeEntryId).Ui(ui).VulnerabilityType(vulnerabilityType).Execute()

Create and submit a vulnerability report in one call; members and, for active public engagements, fully onboarded researchers. 400 names the missing fields; 409 pending_budget_exceeded keeps the content as the draft named by draftId. To submit an existing draft use postReportsSubmit

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
	content := "content_example" // string | The report body as one Markdown document. Its single H1 is the report title; the sections are Description, Steps to reproduce, Impact, and Proof of concept. Description and Proof of concept are required to submit
	ac := "ac_example" // string | CVSS attack complexity: L or H (optional)
	av := "av_example" // string | CVSS attack vector: N, A, L, or P (optional)
	avail := "avail_example" // string | CVSS availability impact: N, L, or H (optional)
	conf := "conf_example" // string | CVSS confidentiality impact: N, L, or H (optional)
	e := "e_example" // string | CVSS exploit code maturity: X, U, P, F, or H; X or empty is not defined (optional)
	integ := "integ_example" // string | CVSS integrity impact: N, L, or H (optional)
	pr := "pr_example" // string | CVSS privileges required: N, L, or H (optional)
	rc := "rc_example" // string | CVSS report confidence: X, U, R, or C; X or empty is not defined (optional)
	reason := "reason_example" // string | Why this revision exists, recorded on the revision (post-submission edits) (optional)
	rl := "rl_example" // string | CVSS remediation level: X, O, T, W, or U; X or empty is not defined (optional)
	scope := "scope_example" // string | CVSS scope: U or C (optional)
	scopeEntryId := "scopeEntryId_example" // string | In-scope asset id from the engagement's live scope. Submitting needs one (optional)
	ui := "ui_example" // string | CVSS user interaction: N or R (optional)
	vulnerabilityType := "vulnerabilityType_example" // string | Weakness key from the catalog, e.g. CWE-79 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EngagementsAPI.PostEngagementsSubmit(context.Background(), orgSlug, engagementSlug).Content(content).Ac(ac).Av(av).Avail(avail).Conf(conf).E(e).Integ(integ).Pr(pr).Rc(rc).Reason(reason).Rl(rl).Scope(scope).ScopeEntryId(scopeEntryId).Ui(ui).VulnerabilityType(vulnerabilityType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngagementsAPI.PostEngagementsSubmit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostEngagementsSubmit`: Report
	fmt.Fprintf(os.Stdout, "Response from `EngagementsAPI.PostEngagementsSubmit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 
**engagementSlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostEngagementsSubmitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **content** | **string** | The report body as one Markdown document. Its single H1 is the report title; the sections are Description, Steps to reproduce, Impact, and Proof of concept. Description and Proof of concept are required to submit | 
 **ac** | **string** | CVSS attack complexity: L or H | 
 **av** | **string** | CVSS attack vector: N, A, L, or P | 
 **avail** | **string** | CVSS availability impact: N, L, or H | 
 **conf** | **string** | CVSS confidentiality impact: N, L, or H | 
 **e** | **string** | CVSS exploit code maturity: X, U, P, F, or H; X or empty is not defined | 
 **integ** | **string** | CVSS integrity impact: N, L, or H | 
 **pr** | **string** | CVSS privileges required: N, L, or H | 
 **rc** | **string** | CVSS report confidence: X, U, R, or C; X or empty is not defined | 
 **reason** | **string** | Why this revision exists, recorded on the revision (post-submission edits) | 
 **rl** | **string** | CVSS remediation level: X, O, T, W, or U; X or empty is not defined | 
 **scope** | **string** | CVSS scope: U or C | 
 **scopeEntryId** | **string** | In-scope asset id from the engagement&#39;s live scope. Submitting needs one | 
 **ui** | **string** | CVSS user interaction: N or R | 
 **vulnerabilityType** | **string** | Weakness key from the catalog, e.g. CWE-79 | 

### Return type

[**Report**](Report.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostEngagementsSubmitDraft

> Report PostEngagementsSubmitDraft(ctx, orgSlug, engagementSlug).Content(content).Ac(ac).Av(av).Avail(avail).Conf(conf).E(e).Integ(integ).Pr(pr).Rc(rc).Rl(rl).Scope(scope).ScopeEntryId(scopeEntryId).Ui(ui).VulnerabilityType(vulnerabilityType).Execute()

Save a private draft report; no completeness validation, only a title is required. Submit it later with postReportsSubmit

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
	content := "content_example" // string | The report body as one Markdown document. Its single H1 is the report title; the sections are Description, Steps to reproduce, Impact, and Proof of concept. A draft may leave sections empty
	ac := "ac_example" // string | CVSS attack complexity: L or H (optional)
	av := "av_example" // string | CVSS attack vector: N, A, L, or P (optional)
	avail := "avail_example" // string | CVSS availability impact: N, L, or H (optional)
	conf := "conf_example" // string | CVSS confidentiality impact: N, L, or H (optional)
	e := "e_example" // string | CVSS exploit code maturity: X, U, P, F, or H; X or empty is not defined (optional)
	integ := "integ_example" // string | CVSS integrity impact: N, L, or H (optional)
	pr := "pr_example" // string | CVSS privileges required: N, L, or H (optional)
	rc := "rc_example" // string | CVSS report confidence: X, U, R, or C; X or empty is not defined (optional)
	rl := "rl_example" // string | CVSS remediation level: X, O, T, W, or U; X or empty is not defined (optional)
	scope := "scope_example" // string | CVSS scope: U or C (optional)
	scopeEntryId := "scopeEntryId_example" // string | In-scope asset id from the engagement's live scope (optional)
	ui := "ui_example" // string | CVSS user interaction: N or R (optional)
	vulnerabilityType := "vulnerabilityType_example" // string | Vulnerability classification, e.g. XSS (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EngagementsAPI.PostEngagementsSubmitDraft(context.Background(), orgSlug, engagementSlug).Content(content).Ac(ac).Av(av).Avail(avail).Conf(conf).E(e).Integ(integ).Pr(pr).Rc(rc).Rl(rl).Scope(scope).ScopeEntryId(scopeEntryId).Ui(ui).VulnerabilityType(vulnerabilityType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngagementsAPI.PostEngagementsSubmitDraft``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostEngagementsSubmitDraft`: Report
	fmt.Fprintf(os.Stdout, "Response from `EngagementsAPI.PostEngagementsSubmitDraft`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 
**engagementSlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostEngagementsSubmitDraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **content** | **string** | The report body as one Markdown document. Its single H1 is the report title; the sections are Description, Steps to reproduce, Impact, and Proof of concept. A draft may leave sections empty | 
 **ac** | **string** | CVSS attack complexity: L or H | 
 **av** | **string** | CVSS attack vector: N, A, L, or P | 
 **avail** | **string** | CVSS availability impact: N, L, or H | 
 **conf** | **string** | CVSS confidentiality impact: N, L, or H | 
 **e** | **string** | CVSS exploit code maturity: X, U, P, F, or H; X or empty is not defined | 
 **integ** | **string** | CVSS integrity impact: N, L, or H | 
 **pr** | **string** | CVSS privileges required: N, L, or H | 
 **rc** | **string** | CVSS report confidence: X, U, R, or C; X or empty is not defined | 
 **rl** | **string** | CVSS remediation level: X, O, T, W, or U; X or empty is not defined | 
 **scope** | **string** | CVSS scope: U or C | 
 **scopeEntryId** | **string** | In-scope asset id from the engagement&#39;s live scope | 
 **ui** | **string** | CVSS user interaction: N or R | 
 **vulnerabilityType** | **string** | Vulnerability classification, e.g. XSS | 

### Return type

[**Report**](Report.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

