# \EngagementsAPI

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEngagements**](EngagementsAPI.md#GetEngagements) | **Get** /engagements | Engagement directory: active and paused engagements visible to the authenticated user, filterable by kind, industry, scope type (scope), and minimum bounty (bounty), paginated with page
[**GetEngagementsByOrgSlugAndEngagementSlug**](EngagementsAPI.md#GetEngagementsByOrgSlugAndEngagementSlug) | **Get** /engagements/{org_slug}/{engagement_slug} | Engagement detail; public engagements for any user, private ones members-only with 404 for non-members (ADR-0009)
[**GetEngagementsMessageProxy**](EngagementsAPI.md#GetEngagementsMessageProxy) | **Get** /engagements/{org_slug}/{engagement_slug}/message-proxy | Authentication codes forwarded from the engagement&#39;s 2FA-protected test accounts within the last 15 minutes; requires testing access and accepted rules
[**PostEngagementsRulesAccept**](EngagementsAPI.md#PostEngagementsRulesAccept) | **Post** /engagements/{org_slug}/{engagement_slug}/rules/accept | Accept the engagement&#39;s published rules of engagement; required before submitting reports, re-required after a new version publishes
[**PostEngagementsSubmit**](EngagementsAPI.md#PostEngagementsSubmit) | **Post** /engagements/{org_slug}/{engagement_slug}/submit | Create and submit a vulnerability report in one call; members and, for active public engagements, fully onboarded researchers. 400 names the missing fields; 409 pending_budget_exceeded keeps the content as the draft named by draftId. To submit an existing draft use postReportsSubmit
[**PostEngagementsSubmitDraft**](EngagementsAPI.md#PostEngagementsSubmitDraft) | **Post** /engagements/{org_slug}/{engagement_slug}/submit/draft | Save a private draft report; no completeness validation, only a title is required. Submit it later with postReportsSubmit



## GetEngagements

> EngagementListItemList GetEngagements(ctx).Execute()

Engagement directory: active and paused engagements visible to the authenticated user, filterable by kind, industry, scope type (scope), and minimum bounty (bounty), paginated with page

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

Engagement detail; public engagements for any user, private ones members-only with 404 for non-members (ADR-0009)

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

> Report PostEngagementsSubmit(ctx, orgSlug, engagementSlug).Description(description).Origin(origin).Title(title).Ac(ac).Av(av).Avail(avail).Conf(conf).Endpoint(endpoint).Impact(impact).Integ(integ).Poc(poc).Pr(pr).Scope(scope).ScopeEntryId(scopeEntryId).Severity(severity).Steps(steps).Ui(ui).VulnerabilityType(vulnerabilityType).Execute()

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
	description := "description_example" // string | Vulnerability description and root cause (markdown)
	origin := "origin_example" // string | AI-use disclosure
	title := "title_example" // string | Vulnerability title
	ac := "ac_example" // string | CVSS attack complexity: L or H (optional)
	av := "av_example" // string | CVSS attack vector: N, A, L, or P (optional)
	avail := "avail_example" // string | CVSS availability impact: N, L, or H (optional)
	conf := "conf_example" // string | CVSS confidentiality impact: N, L, or H (optional)
	endpoint := "endpoint_example" // string | Affected endpoint or URL, refines the selected asset. Submitting needs this or a scopeEntryId (optional)
	impact := "impact_example" // string | Security impact description (optional)
	integ := "integ_example" // string | CVSS integrity impact: N, L, or H (optional)
	poc := "poc_example" // string | Proof-of-concept payload, exploit code, or evidence reference. Required to submit, optional on a draft (optional)
	pr := "pr_example" // string | CVSS privileges required: N, L, or H (optional)
	scope := "scope_example" // string | CVSS scope: U or C (optional)
	scopeEntryId := "scopeEntryId_example" // string | In-scope asset id from the engagement's live scope. Submitting needs this or an endpoint (optional)
	severity := "severity_example" // string | Claimed severity; empty derives it from the CVSS metrics. Submitting needs this or all eight CVSS metrics (optional)
	steps := "steps_example" // string | Reproduction steps, appended to the description (optional)
	ui := "ui_example" // string | CVSS user interaction: N or R (optional)
	vulnerabilityType := "vulnerabilityType_example" // string | Vulnerability classification, e.g. XSS (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EngagementsAPI.PostEngagementsSubmit(context.Background(), orgSlug, engagementSlug).Description(description).Origin(origin).Title(title).Ac(ac).Av(av).Avail(avail).Conf(conf).Endpoint(endpoint).Impact(impact).Integ(integ).Poc(poc).Pr(pr).Scope(scope).ScopeEntryId(scopeEntryId).Severity(severity).Steps(steps).Ui(ui).VulnerabilityType(vulnerabilityType).Execute()
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


 **description** | **string** | Vulnerability description and root cause (markdown) | 
 **origin** | **string** | AI-use disclosure | 
 **title** | **string** | Vulnerability title | 
 **ac** | **string** | CVSS attack complexity: L or H | 
 **av** | **string** | CVSS attack vector: N, A, L, or P | 
 **avail** | **string** | CVSS availability impact: N, L, or H | 
 **conf** | **string** | CVSS confidentiality impact: N, L, or H | 
 **endpoint** | **string** | Affected endpoint or URL, refines the selected asset. Submitting needs this or a scopeEntryId | 
 **impact** | **string** | Security impact description | 
 **integ** | **string** | CVSS integrity impact: N, L, or H | 
 **poc** | **string** | Proof-of-concept payload, exploit code, or evidence reference. Required to submit, optional on a draft | 
 **pr** | **string** | CVSS privileges required: N, L, or H | 
 **scope** | **string** | CVSS scope: U or C | 
 **scopeEntryId** | **string** | In-scope asset id from the engagement&#39;s live scope. Submitting needs this or an endpoint | 
 **severity** | **string** | Claimed severity; empty derives it from the CVSS metrics. Submitting needs this or all eight CVSS metrics | 
 **steps** | **string** | Reproduction steps, appended to the description | 
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


## PostEngagementsSubmitDraft

> Report PostEngagementsSubmitDraft(ctx, orgSlug, engagementSlug).Description(description).Origin(origin).Title(title).Ac(ac).Av(av).Avail(avail).Conf(conf).Endpoint(endpoint).Impact(impact).Integ(integ).Poc(poc).Pr(pr).Scope(scope).ScopeEntryId(scopeEntryId).Severity(severity).Steps(steps).Ui(ui).VulnerabilityType(vulnerabilityType).Execute()

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
	description := "description_example" // string | Vulnerability description and root cause (markdown)
	origin := "origin_example" // string | AI-use disclosure
	title := "title_example" // string | Vulnerability title
	ac := "ac_example" // string | CVSS attack complexity: L or H (optional)
	av := "av_example" // string | CVSS attack vector: N, A, L, or P (optional)
	avail := "avail_example" // string | CVSS availability impact: N, L, or H (optional)
	conf := "conf_example" // string | CVSS confidentiality impact: N, L, or H (optional)
	endpoint := "endpoint_example" // string | Affected endpoint or URL, refines the selected asset. Submitting needs this or a scopeEntryId (optional)
	impact := "impact_example" // string | Security impact description (optional)
	integ := "integ_example" // string | CVSS integrity impact: N, L, or H (optional)
	poc := "poc_example" // string | Proof-of-concept payload, exploit code, or evidence reference. Required to submit, optional on a draft (optional)
	pr := "pr_example" // string | CVSS privileges required: N, L, or H (optional)
	scope := "scope_example" // string | CVSS scope: U or C (optional)
	scopeEntryId := "scopeEntryId_example" // string | In-scope asset id from the engagement's live scope. Submitting needs this or an endpoint (optional)
	severity := "severity_example" // string | Claimed severity; empty derives it from the CVSS metrics. Submitting needs this or all eight CVSS metrics (optional)
	steps := "steps_example" // string | Reproduction steps, appended to the description (optional)
	ui := "ui_example" // string | CVSS user interaction: N or R (optional)
	vulnerabilityType := "vulnerabilityType_example" // string | Vulnerability classification, e.g. XSS (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EngagementsAPI.PostEngagementsSubmitDraft(context.Background(), orgSlug, engagementSlug).Description(description).Origin(origin).Title(title).Ac(ac).Av(av).Avail(avail).Conf(conf).Endpoint(endpoint).Impact(impact).Integ(integ).Poc(poc).Pr(pr).Scope(scope).ScopeEntryId(scopeEntryId).Severity(severity).Steps(steps).Ui(ui).VulnerabilityType(vulnerabilityType).Execute()
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


 **description** | **string** | Vulnerability description and root cause (markdown) | 
 **origin** | **string** | AI-use disclosure | 
 **title** | **string** | Vulnerability title | 
 **ac** | **string** | CVSS attack complexity: L or H | 
 **av** | **string** | CVSS attack vector: N, A, L, or P | 
 **avail** | **string** | CVSS availability impact: N, L, or H | 
 **conf** | **string** | CVSS confidentiality impact: N, L, or H | 
 **endpoint** | **string** | Affected endpoint or URL, refines the selected asset. Submitting needs this or a scopeEntryId | 
 **impact** | **string** | Security impact description | 
 **integ** | **string** | CVSS integrity impact: N, L, or H | 
 **poc** | **string** | Proof-of-concept payload, exploit code, or evidence reference. Required to submit, optional on a draft | 
 **pr** | **string** | CVSS privileges required: N, L, or H | 
 **scope** | **string** | CVSS scope: U or C | 
 **scopeEntryId** | **string** | In-scope asset id from the engagement&#39;s live scope. Submitting needs this or an endpoint | 
 **severity** | **string** | Claimed severity; empty derives it from the CVSS metrics. Submitting needs this or all eight CVSS metrics | 
 **steps** | **string** | Reproduction steps, appended to the description | 
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

