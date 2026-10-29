# \ReportsAPI

All URIs are relative to *https://app.gobugfree.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetReports**](ReportsAPI.md#GetReports) | **Get** /reports | The researcher&#39;s own reports: filterable by status, engagement, severity, and search, sortable, paginated
[**GetReportsByReportId**](ReportsAPI.md#GetReportsByReportId) | **Get** /reports/{report_id} | One of the researcher&#39;s own reports with its status history and researcher-visible discussion; 404 for other researchers&#39; reports (ADR-0009)
[**GetReportsCollaborations**](ReportsAPI.md#GetReportsCollaborations) | **Get** /reports/collaborations | The researcher&#39;s bounty-split collaboration invitations
[**GetReportsContentEdit**](ReportsAPI.md#GetReportsContentEdit) | **Get** /reports/{report_id}/content/edit | The report&#39;s body document for editing; 409 once triage has decided
[**GetReportsEdit**](ReportsAPI.md#GetReportsEdit) | **Get** /reports/{report_id}/edit | The researcher&#39;s draft report open in the editor; 404 for foreign or already-submitted reports
[**GetReportsEvidenceByEvidenceId**](ReportsAPI.md#GetReportsEvidenceByEvidenceId) | **Get** /reports/{report_id}/evidence/{evidence_id} | Download an evidence file (attachment); the owner always, the other parties to the report after submission; 404 otherwise
[**GetReportsExport**](ReportsAPI.md#GetReportsExport) | **Get** /reports/export | Download the researcher&#39;s filtered reports as CSV or JSON (every page, at most 10000 rows)
[**GetReportsOpen**](ReportsAPI.md#GetReportsOpen) | **Get** /reports/{report_id}/open | Redirect to the surface the caller may open the report on; 404 when there is none
[**GetReportsVersions**](ReportsAPI.md#GetReportsVersions) | **Get** /reports/{report_id}/versions | The report body&#39;s revisions, newest first; the page is the searchable history with each revision&#39;s diff
[**GetReportsVersionsByNumber**](ReportsAPI.md#GetReportsVersionsByNumber) | **Get** /reports/{report_id}/versions/{number} | One revision of the report body; 404 for an unknown number
[**GetReportsVersionsDiff**](ReportsAPI.md#GetReportsVersionsDiff) | **Get** /reports/{report_id}/versions/{number}/diff | What one revision changed against the revision before it: the document as a unified diff and the finding field by field
[**GetReportsVersionsDocument**](ReportsAPI.md#GetReportsVersionsDocument) | **Get** /reports/{report_id}/versions/{number}/document | One revision as its exchange document: the finding as YAML front matter above the Markdown body (text/markdown)
[**GetReportsVersionsExport**](ReportsAPI.md#GetReportsVersionsExport) | **Get** /reports/{report_id}/versions/export | Download the report&#39;s revision history as CSV or JSON, carrying the history&#39;s own search and sort
[**PostReportsCollaborators**](ReportsAPI.md#PostReportsCollaborators) | **Post** /reports/{report_id}/collaborators | Invite a researcher to an own draft&#39;s bounty split with a proposed share
[**PostReportsCollaboratorsRemove**](ReportsAPI.md#PostReportsCollaboratorsRemove) | **Post** /reports/{report_id}/collaborators/{collaborator_id}/remove | Remove a collaborator from an own draft
[**PostReportsCollaboratorsRespond**](ReportsAPI.md#PostReportsCollaboratorsRespond) | **Post** /reports/{report_id}/collaborators/respond | Accept or decline a bounty-split invitation (invitee only)
[**PostReportsComments**](ReportsAPI.md#PostReportsComments) | **Post** /reports/{report_id}/comments | Post a message on an own report&#39;s thread: audience &#39;researcher&#39; (to the triage team, answers a pending information request), &#39;direct&#39; (to the organization, only while it triages the report), or &#39;researcher_internal&#39; (private note); a multipart body may add up to 5 files as &#39;files&#39; parts
[**PostReportsContent**](ReportsAPI.md#PostReportsContent) | **Post** /reports/{report_id}/content | Rewrite the report&#39;s body as one Markdown document, appending a revision; 409 once triage has decided or for a writer who is neither the reporter nor an accepted collaborator
[**PostReportsDelete**](ReportsAPI.md#PostReportsDelete) | **Post** /reports/{report_id}/delete | Delete an own draft report
[**PostReportsEdit**](ReportsAPI.md#PostReportsEdit) | **Post** /reports/{report_id}/edit | Re-save an own draft report
[**PostReportsEvidence**](ReportsAPI.md#PostReportsEvidence) | **Post** /reports/{report_id}/evidence | Attach a multipart evidence file to an own report while it is a draft or awaits triage; scanned before it can download; files added after submission cannot be removed
[**PostReportsEvidenceDelete**](ReportsAPI.md#PostReportsEvidenceDelete) | **Post** /reports/{report_id}/evidence/{evidence_id}/delete | Delete an evidence file from an own draft report
[**PostReportsSubmit**](ReportsAPI.md#PostReportsSubmit) | **Post** /reports/{report_id}/submit | Submit an own draft report; runs full completeness validation (400 names the missing fields) and re-checks engagement access; 409 pending_budget_exceeded when the pending-report budget is used up
[**PostReportsVersionsRestore**](ReportsAPI.md#PostReportsVersionsRestore) | **Post** /reports/{report_id}/versions/{number}/restore | Make an old revision&#39;s document the report&#39;s next revision; the finding stays as it is; 403 once triage has decided



## GetReports

> ReportMyReportList GetReports(ctx).Status(status).Severity(severity).Q(q).Range_(range_).From(from).To(to).Date(date).Urgent(urgent).Sort(sort).Dir(dir).Page(page).Size(size).Engagement(engagement).Execute()

The researcher's own reports: filterable by status, engagement, severity, and search, sortable, paginated

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
	status := []string{"Status_example"} // []string | Workflow states to include (optional)
	severity := []string{"Severity_example"} // []string | Severities (assessed, else submitted) to include (optional)
	q := "q_example" // string | Free-text search: title substring or report id prefix (optional)
	range_ := "range__example" // string | Time window: a preset counted back from now, or custom with from/to; omitted or all means no window (optional)
	from := "from_example" // string | First day of a custom window (YYYY-MM-DD), open when omitted (optional)
	to := "to_example" // string | Last day of a custom window (YYYY-MM-DD, inclusive), open when omitted (optional)
	date := "date_example" // string | Date the window applies to: submitted (default) or accepted (the acceptance) (optional)
	urgent := "urgent_example" // string | 1 keeps only the reports that need attention now: escalated, SLA breached, or inside the last fifth of an open deadline (optional)
	sort := "sort_example" // string | Sort key (optional)
	dir := "dir_example" // string | Sort direction (optional)
	page := "page_example" // string | Page number, from 1 (optional)
	size := "size_example" // string | Rows per page (optional)
	engagement := "engagement_example" // string | Engagement as org_slug/engagement_slug (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.GetReports(context.Background()).Status(status).Severity(severity).Q(q).Range_(range_).From(from).To(to).Date(date).Urgent(urgent).Sort(sort).Dir(dir).Page(page).Size(size).Engagement(engagement).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetReports``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReports`: ReportMyReportList
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GetReports`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetReportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **[]string** | Workflow states to include | 
 **severity** | **[]string** | Severities (assessed, else submitted) to include | 
 **q** | **string** | Free-text search: title substring or report id prefix | 
 **range_** | **string** | Time window: a preset counted back from now, or custom with from/to; omitted or all means no window | 
 **from** | **string** | First day of a custom window (YYYY-MM-DD), open when omitted | 
 **to** | **string** | Last day of a custom window (YYYY-MM-DD, inclusive), open when omitted | 
 **date** | **string** | Date the window applies to: submitted (default) or accepted (the acceptance) | 
 **urgent** | **string** | 1 keeps only the reports that need attention now: escalated, SLA breached, or inside the last fifth of an open deadline | 
 **sort** | **string** | Sort key | 
 **dir** | **string** | Sort direction | 
 **page** | **string** | Page number, from 1 | 
 **size** | **string** | Rows per page | 
 **engagement** | **string** | Engagement as org_slug/engagement_slug | 

### Return type

[**ReportMyReportList**](ReportMyReportList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportsByReportId

> ReportMyReportDetail GetReportsByReportId(ctx, reportId).Execute()

One of the researcher's own reports with its status history and researcher-visible discussion; 404 for other researchers' reports (ADR-0009)

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
	reportId := "reportId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.GetReportsByReportId(context.Background(), reportId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetReportsByReportId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReportsByReportId`: ReportMyReportDetail
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GetReportsByReportId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportsByReportIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReportMyReportDetail**](ReportMyReportDetail.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportsCollaborations

> ReportInvitationList GetReportsCollaborations(ctx).Execute()

The researcher's bounty-split collaboration invitations

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
	resp, r, err := apiClient.ReportsAPI.GetReportsCollaborations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetReportsCollaborations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReportsCollaborations`: ReportInvitationList
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GetReportsCollaborations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportsCollaborationsRequest struct via the builder pattern


### Return type

[**ReportInvitationList**](ReportInvitationList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportsContentEdit

> GetReportsContentEdit(ctx, reportId).Execute()

The report's body document for editing; 409 once triage has decided

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
	reportId := "reportId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ReportsAPI.GetReportsContentEdit(context.Background(), reportId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetReportsContentEdit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportsContentEditRequest struct via the builder pattern


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


## GetReportsEdit

> GetReportsEdit(ctx, reportId).Execute()

The researcher's draft report open in the editor; 404 for foreign or already-submitted reports

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
	reportId := "reportId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ReportsAPI.GetReportsEdit(context.Background(), reportId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetReportsEdit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportsEditRequest struct via the builder pattern


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


## GetReportsEvidenceByEvidenceId

> GetReportsEvidenceByEvidenceId(ctx, reportId, evidenceId).Execute()

Download an evidence file (attachment); the owner always, the other parties to the report after submission; 404 otherwise

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
	reportId := "reportId_example" // string | 
	evidenceId := "evidenceId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ReportsAPI.GetReportsEvidenceByEvidenceId(context.Background(), reportId, evidenceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetReportsEvidenceByEvidenceId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **string** |  | 
**evidenceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportsEvidenceByEvidenceIdRequest struct via the builder pattern


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


## GetReportsExport

> GetReportsExport(ctx).Status(status).Severity(severity).Q(q).Range_(range_).From(from).To(to).Date(date).Urgent(urgent).Sort(sort).Dir(dir).Page(page).Size(size).Engagement(engagement).Format(format).Execute()

Download the researcher's filtered reports as CSV or JSON (every page, at most 10000 rows)

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
	status := []string{"Status_example"} // []string | Workflow states to include (optional)
	severity := []string{"Severity_example"} // []string | Severities (assessed, else submitted) to include (optional)
	q := "q_example" // string | Free-text search: title substring or report id prefix (optional)
	range_ := "range__example" // string | Time window: a preset counted back from now, or custom with from/to; omitted or all means no window (optional)
	from := "from_example" // string | First day of a custom window (YYYY-MM-DD), open when omitted (optional)
	to := "to_example" // string | Last day of a custom window (YYYY-MM-DD, inclusive), open when omitted (optional)
	date := "date_example" // string | Date the window applies to: submitted (default) or accepted (the acceptance) (optional)
	urgent := "urgent_example" // string | 1 keeps only the reports that need attention now: escalated, SLA breached, or inside the last fifth of an open deadline (optional)
	sort := "sort_example" // string | Sort key (optional)
	dir := "dir_example" // string | Sort direction (optional)
	page := "page_example" // string | Page number, from 1 (optional)
	size := "size_example" // string | Rows per page (optional)
	engagement := "engagement_example" // string | Engagement as org_slug/engagement_slug (optional)
	format := "format_example" // string | Download format (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ReportsAPI.GetReportsExport(context.Background()).Status(status).Severity(severity).Q(q).Range_(range_).From(from).To(to).Date(date).Urgent(urgent).Sort(sort).Dir(dir).Page(page).Size(size).Engagement(engagement).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetReportsExport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetReportsExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **[]string** | Workflow states to include | 
 **severity** | **[]string** | Severities (assessed, else submitted) to include | 
 **q** | **string** | Free-text search: title substring or report id prefix | 
 **range_** | **string** | Time window: a preset counted back from now, or custom with from/to; omitted or all means no window | 
 **from** | **string** | First day of a custom window (YYYY-MM-DD), open when omitted | 
 **to** | **string** | Last day of a custom window (YYYY-MM-DD, inclusive), open when omitted | 
 **date** | **string** | Date the window applies to: submitted (default) or accepted (the acceptance) | 
 **urgent** | **string** | 1 keeps only the reports that need attention now: escalated, SLA breached, or inside the last fifth of an open deadline | 
 **sort** | **string** | Sort key | 
 **dir** | **string** | Sort direction | 
 **page** | **string** | Page number, from 1 | 
 **size** | **string** | Rows per page | 
 **engagement** | **string** | Engagement as org_slug/engagement_slug | 
 **format** | **string** | Download format | 

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


## GetReportsOpen

> GetReportsOpen(ctx, reportId).Execute()

Redirect to the surface the caller may open the report on; 404 when there is none

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
	reportId := "reportId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ReportsAPI.GetReportsOpen(context.Background(), reportId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetReportsOpen``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportsOpenRequest struct via the builder pattern


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


## GetReportsVersions

> ReportContentVersionList GetReportsVersions(ctx, reportId).Execute()

The report body's revisions, newest first; the page is the searchable history with each revision's diff

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
	reportId := "reportId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.GetReportsVersions(context.Background(), reportId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetReportsVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReportsVersions`: ReportContentVersionList
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GetReportsVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportsVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReportContentVersionList**](ReportContentVersionList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportsVersionsByNumber

> ReportContentVersion GetReportsVersionsByNumber(ctx, reportId, number).Execute()

One revision of the report body; 404 for an unknown number

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
	reportId := "reportId_example" // string | 
	number := "number_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.GetReportsVersionsByNumber(context.Background(), reportId, number).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetReportsVersionsByNumber``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReportsVersionsByNumber`: ReportContentVersion
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GetReportsVersionsByNumber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **string** |  | 
**number** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportsVersionsByNumberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ReportContentVersion**](ReportContentVersion.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportsVersionsDiff

> GetReportsVersionsDiff(ctx, reportId, number).Execute()

What one revision changed against the revision before it: the document as a unified diff and the finding field by field

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
	reportId := "reportId_example" // string | 
	number := "number_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ReportsAPI.GetReportsVersionsDiff(context.Background(), reportId, number).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetReportsVersionsDiff``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **string** |  | 
**number** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportsVersionsDiffRequest struct via the builder pattern


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


## GetReportsVersionsDocument

> GetReportsVersionsDocument(ctx, reportId, number).Execute()

One revision as its exchange document: the finding as YAML front matter above the Markdown body (text/markdown)

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
	reportId := "reportId_example" // string | 
	number := "number_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ReportsAPI.GetReportsVersionsDocument(context.Background(), reportId, number).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetReportsVersionsDocument``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **string** |  | 
**number** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportsVersionsDocumentRequest struct via the builder pattern


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


## GetReportsVersionsExport

> GetReportsVersionsExport(ctx, reportId).Execute()

Download the report's revision history as CSV or JSON, carrying the history's own search and sort

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
	reportId := "reportId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ReportsAPI.GetReportsVersionsExport(context.Background(), reportId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetReportsVersionsExport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportsVersionsExportRequest struct via the builder pattern


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


## PostReportsCollaborators

> ReportContributor PostReportsCollaborators(ctx, reportId).Execute()

Invite a researcher to an own draft's bounty split with a proposed share

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
	reportId := "reportId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.PostReportsCollaborators(context.Background(), reportId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.PostReportsCollaborators``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostReportsCollaborators`: ReportContributor
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.PostReportsCollaborators`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostReportsCollaboratorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReportContributor**](ReportContributor.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostReportsCollaboratorsRemove

> PostReportsCollaboratorsRemove(ctx, reportId, collaboratorId).Execute()

Remove a collaborator from an own draft

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
	reportId := "reportId_example" // string | 
	collaboratorId := "collaboratorId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ReportsAPI.PostReportsCollaboratorsRemove(context.Background(), reportId, collaboratorId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.PostReportsCollaboratorsRemove``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **string** |  | 
**collaboratorId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostReportsCollaboratorsRemoveRequest struct via the builder pattern


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


## PostReportsCollaboratorsRespond

> PostReportsCollaboratorsRespond(ctx, reportId).Action(action).Execute()

Accept or decline a bounty-split invitation (invitee only)

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
	reportId := "reportId_example" // string | 
	action := "action_example" // string | accept or decline the invitation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ReportsAPI.PostReportsCollaboratorsRespond(context.Background(), reportId).Action(action).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.PostReportsCollaboratorsRespond``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostReportsCollaboratorsRespondRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **action** | **string** | accept or decline the invitation | 

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


## PostReportsComments

> ReportPostedComment PostReportsComments(ctx, reportId).Body(body).Audience(audience).Execute()

Post a message on an own report's thread: audience 'researcher' (to the triage team, answers a pending information request), 'direct' (to the organization, only while it triages the report), or 'researcher_internal' (private note); a multipart body may add up to 5 files as 'files' parts

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
	reportId := "reportId_example" // string | 
	body := "body_example" // string | The message text
	audience := "audience_example" // string | Message audience; the server enforces the role's allowlist (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.PostReportsComments(context.Background(), reportId).Body(body).Audience(audience).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.PostReportsComments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostReportsComments`: ReportPostedComment
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.PostReportsComments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostReportsCommentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** | The message text | 
 **audience** | **string** | Message audience; the server enforces the role&#39;s allowlist | 

### Return type

[**ReportPostedComment**](ReportPostedComment.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostReportsContent

> Report PostReportsContent(ctx, reportId).Content(content).Reason(reason).Execute()

Rewrite the report's body as one Markdown document, appending a revision; 409 once triage has decided or for a writer who is neither the reporter nor an accepted collaborator

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
	reportId := "reportId_example" // string | 
	content := "content_example" // string | The report body as one Markdown document. Sections: Description, Steps to reproduce, Impact, Proof of concept
	reason := "reason_example" // string | Why this revision exists, recorded on the revision (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.PostReportsContent(context.Background(), reportId).Content(content).Reason(reason).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.PostReportsContent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostReportsContent`: Report
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.PostReportsContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostReportsContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **content** | **string** | The report body as one Markdown document. Sections: Description, Steps to reproduce, Impact, Proof of concept | 
 **reason** | **string** | Why this revision exists, recorded on the revision | 

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


## PostReportsDelete

> PostReportsDelete(ctx, reportId).Execute()

Delete an own draft report

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
	reportId := "reportId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ReportsAPI.PostReportsDelete(context.Background(), reportId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.PostReportsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostReportsDeleteRequest struct via the builder pattern


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


## PostReportsEdit

> PostReportsEdit(ctx, reportId).Content(content).Ac(ac).Av(av).Avail(avail).Conf(conf).E(e).Integ(integ).Pr(pr).Rc(rc).Rl(rl).Scope(scope).ScopeEntryId(scopeEntryId).Ui(ui).VulnerabilityType(vulnerabilityType).Execute()

Re-save an own draft report

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
	reportId := "reportId_example" // string | 
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
	r, err := apiClient.ReportsAPI.PostReportsEdit(context.Background(), reportId).Content(content).Ac(ac).Av(av).Avail(avail).Conf(conf).E(e).Integ(integ).Pr(pr).Rc(rc).Rl(rl).Scope(scope).ScopeEntryId(scopeEntryId).Ui(ui).VulnerabilityType(vulnerabilityType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.PostReportsEdit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostReportsEditRequest struct via the builder pattern


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

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostReportsEvidence

> ReportEvidence PostReportsEvidence(ctx, reportId).Evidence(evidence).Execute()

Attach a multipart evidence file to an own report while it is a draft or awaits triage; scanned before it can download; files added after submission cannot be removed

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
	reportId := "reportId_example" // string | 
	evidence := os.NewFile(1234, "some_file") // *os.File | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.PostReportsEvidence(context.Background(), reportId).Evidence(evidence).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.PostReportsEvidence``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostReportsEvidence`: ReportEvidence
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.PostReportsEvidence`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostReportsEvidenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **evidence** | ***os.File** |  | 

### Return type

[**ReportEvidence**](ReportEvidence.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostReportsEvidenceDelete

> PostReportsEvidenceDelete(ctx, reportId, evidenceId).Execute()

Delete an evidence file from an own draft report

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
	reportId := "reportId_example" // string | 
	evidenceId := "evidenceId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ReportsAPI.PostReportsEvidenceDelete(context.Background(), reportId, evidenceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.PostReportsEvidenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **string** |  | 
**evidenceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostReportsEvidenceDeleteRequest struct via the builder pattern


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


## PostReportsSubmit

> Report PostReportsSubmit(ctx, reportId).Content(content).Ac(ac).Av(av).Avail(avail).Conf(conf).E(e).Integ(integ).Pr(pr).Rc(rc).Reason(reason).Rl(rl).Scope(scope).ScopeEntryId(scopeEntryId).Ui(ui).VulnerabilityType(vulnerabilityType).Execute()

Submit an own draft report; runs full completeness validation (400 names the missing fields) and re-checks engagement access; 409 pending_budget_exceeded when the pending-report budget is used up

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
	reportId := "reportId_example" // string | 
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
	resp, r, err := apiClient.ReportsAPI.PostReportsSubmit(context.Background(), reportId).Content(content).Ac(ac).Av(av).Avail(avail).Conf(conf).E(e).Integ(integ).Pr(pr).Rc(rc).Reason(reason).Rl(rl).Scope(scope).ScopeEntryId(scopeEntryId).Ui(ui).VulnerabilityType(vulnerabilityType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.PostReportsSubmit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostReportsSubmit`: Report
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.PostReportsSubmit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostReportsSubmitRequest struct via the builder pattern


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


## PostReportsVersionsRestore

> PostReportsVersionsRestore(ctx, reportId, number).Execute()

Make an old revision's document the report's next revision; the finding stays as it is; 403 once triage has decided

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
	reportId := "reportId_example" // string | 
	number := "number_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ReportsAPI.PostReportsVersionsRestore(context.Background(), reportId, number).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.PostReportsVersionsRestore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **string** |  | 
**number** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostReportsVersionsRestoreRequest struct via the builder pattern


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

