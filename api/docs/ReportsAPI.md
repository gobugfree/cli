# \ReportsAPI

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetReports**](ReportsAPI.md#GetReports) | **Get** /reports | The researcher&#39;s own reports: filterable by status, engagement, severity, and search, sortable, paginated
[**GetReportsByReportId**](ReportsAPI.md#GetReportsByReportId) | **Get** /reports/{report_id} | One of the researcher&#39;s own reports with its status history and researcher-visible discussion; 404 for other researchers&#39; reports (ADR-0009)
[**GetReportsCollaborations**](ReportsAPI.md#GetReportsCollaborations) | **Get** /reports/collaborations | The researcher&#39;s bounty-split collaboration invitations
[**GetReportsEdit**](ReportsAPI.md#GetReportsEdit) | **Get** /reports/{report_id}/edit | The researcher&#39;s draft report open in the editor; 404 for foreign or already-submitted reports
[**GetReportsEvidenceByEvidenceId**](ReportsAPI.md#GetReportsEvidenceByEvidenceId) | **Get** /reports/{report_id}/evidence/{evidence_id} | Download an evidence file (attachment); the owner always, the other parties to the report after submission; 404 otherwise
[**GetReportsExport**](ReportsAPI.md#GetReportsExport) | **Get** /reports/export | Download the researcher&#39;s filtered reports as CSV or JSON (every page, at most 10000 rows)
[**GetReportsOpen**](ReportsAPI.md#GetReportsOpen) | **Get** /reports/{report_id}/open | Redirect to the surface the caller may open the report on; 404 when there is none
[**PostReportsCollaborators**](ReportsAPI.md#PostReportsCollaborators) | **Post** /reports/{report_id}/collaborators | Invite a researcher to an own draft&#39;s bounty split with a proposed share
[**PostReportsCollaboratorsRemove**](ReportsAPI.md#PostReportsCollaboratorsRemove) | **Post** /reports/{report_id}/collaborators/{collaborator_id}/remove | Remove a collaborator from an own draft
[**PostReportsCollaboratorsRespond**](ReportsAPI.md#PostReportsCollaboratorsRespond) | **Post** /reports/{report_id}/collaborators/respond | Accept or decline a bounty-split invitation (invitee only)
[**PostReportsComments**](ReportsAPI.md#PostReportsComments) | **Post** /reports/{report_id}/comments | Post a message on an own report&#39;s thread; answers a pending information request
[**PostReportsDelete**](ReportsAPI.md#PostReportsDelete) | **Post** /reports/{report_id}/delete | Delete an own draft report
[**PostReportsDisclosure**](ReportsAPI.md#PostReportsDisclosure) | **Post** /reports/{report_id}/disclosure | Request coordinated disclosure on an own resolved report; owner-only, engagement policy must allow it
[**PostReportsEdit**](ReportsAPI.md#PostReportsEdit) | **Post** /reports/{report_id}/edit | Re-save an own draft report
[**PostReportsEvidence**](ReportsAPI.md#PostReportsEvidence) | **Post** /reports/{report_id}/evidence | Attach a multipart evidence file to an own draft report; scanned before it can download
[**PostReportsEvidenceDelete**](ReportsAPI.md#PostReportsEvidenceDelete) | **Post** /reports/{report_id}/evidence/{evidence_id}/delete | Delete an evidence file from an own draft report
[**PostReportsSubmit**](ReportsAPI.md#PostReportsSubmit) | **Post** /reports/{report_id}/submit | Submit an own draft report; runs full completeness validation (400 names the missing fields) and re-checks engagement access; 409 pending_budget_exceeded when the pending-report budget is used up



## GetReports

> ReportMyReportList GetReports(ctx).Status(status).Severity(severity).Q(q).Sort(sort).Dir(dir).Page(page).Size(size).Engagement(engagement).Execute()

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
	sort := "sort_example" // string | Sort key (optional)
	dir := "dir_example" // string | Sort direction (optional)
	page := "page_example" // string | Page number, from 1 (optional)
	size := "size_example" // string | Rows per page (optional)
	engagement := "engagement_example" // string | Engagement as org_slug/engagement_slug (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.GetReports(context.Background()).Status(status).Severity(severity).Q(q).Sort(sort).Dir(dir).Page(page).Size(size).Engagement(engagement).Execute()
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

> GetReportsExport(ctx).Status(status).Severity(severity).Q(q).Sort(sort).Dir(dir).Page(page).Size(size).Engagement(engagement).Format(format).Execute()

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
	sort := "sort_example" // string | Sort key (optional)
	dir := "dir_example" // string | Sort direction (optional)
	page := "page_example" // string | Page number, from 1 (optional)
	size := "size_example" // string | Rows per page (optional)
	engagement := "engagement_example" // string | Engagement as org_slug/engagement_slug (optional)
	format := "format_example" // string | Download format (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ReportsAPI.GetReportsExport(context.Background()).Status(status).Severity(severity).Q(q).Sort(sort).Dir(dir).Page(page).Size(size).Engagement(engagement).Format(format).Execute()
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

Post a message on an own report's thread; answers a pending information request

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


## PostReportsDisclosure

> Disclosure PostReportsDisclosure(ctx, reportId).Execute()

Request coordinated disclosure on an own resolved report; owner-only, engagement policy must allow it

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
	resp, r, err := apiClient.ReportsAPI.PostReportsDisclosure(context.Background(), reportId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.PostReportsDisclosure``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostReportsDisclosure`: Disclosure
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.PostReportsDisclosure`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostReportsDisclosureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Disclosure**](Disclosure.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostReportsEdit

> PostReportsEdit(ctx, reportId).Title(title).Ac(ac).Av(av).Avail(avail).Conf(conf).Description(description).Endpoint(endpoint).Impact(impact).Integ(integ).Origin(origin).Poc(poc).Pr(pr).Scope(scope).ScopeEntryId(scopeEntryId).Severity(severity).Steps(steps).Ui(ui).VulnerabilityType(vulnerabilityType).Execute()

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
	title := "title_example" // string | Vulnerability title
	ac := "ac_example" // string | CVSS attack complexity: L or H (optional)
	av := "av_example" // string | CVSS attack vector: N, A, L, or P (optional)
	avail := "avail_example" // string | CVSS availability impact: N, L, or H (optional)
	conf := "conf_example" // string | CVSS confidentiality impact: N, L, or H (optional)
	description := "description_example" // string | Vulnerability description and root cause (markdown) (optional)
	endpoint := "endpoint_example" // string | Affected endpoint or URL, refines the selected asset (optional)
	impact := "impact_example" // string | Security impact description (optional)
	integ := "integ_example" // string | CVSS integrity impact: N, L, or H (optional)
	origin := "origin_example" // string | AI-use disclosure; empty keeps the default until submission (optional)
	poc := "poc_example" // string | Proof-of-concept payload, exploit code, or evidence reference. Required to submit, optional on a draft (optional)
	pr := "pr_example" // string | CVSS privileges required: N, L, or H (optional)
	scope := "scope_example" // string | CVSS scope: U or C (optional)
	scopeEntryId := "scopeEntryId_example" // string | In-scope asset id from the engagement's live scope (optional)
	severity := "severity_example" // string | Claimed severity; empty derives it from the CVSS metrics (optional)
	steps := "steps_example" // string | Reproduction steps, appended to the description (optional)
	ui := "ui_example" // string | CVSS user interaction: N or R (optional)
	vulnerabilityType := "vulnerabilityType_example" // string | Vulnerability classification, e.g. XSS (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ReportsAPI.PostReportsEdit(context.Background(), reportId).Title(title).Ac(ac).Av(av).Avail(avail).Conf(conf).Description(description).Endpoint(endpoint).Impact(impact).Integ(integ).Origin(origin).Poc(poc).Pr(pr).Scope(scope).ScopeEntryId(scopeEntryId).Severity(severity).Steps(steps).Ui(ui).VulnerabilityType(vulnerabilityType).Execute()
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

 **title** | **string** | Vulnerability title | 
 **ac** | **string** | CVSS attack complexity: L or H | 
 **av** | **string** | CVSS attack vector: N, A, L, or P | 
 **avail** | **string** | CVSS availability impact: N, L, or H | 
 **conf** | **string** | CVSS confidentiality impact: N, L, or H | 
 **description** | **string** | Vulnerability description and root cause (markdown) | 
 **endpoint** | **string** | Affected endpoint or URL, refines the selected asset | 
 **impact** | **string** | Security impact description | 
 **integ** | **string** | CVSS integrity impact: N, L, or H | 
 **origin** | **string** | AI-use disclosure; empty keeps the default until submission | 
 **poc** | **string** | Proof-of-concept payload, exploit code, or evidence reference. Required to submit, optional on a draft | 
 **pr** | **string** | CVSS privileges required: N, L, or H | 
 **scope** | **string** | CVSS scope: U or C | 
 **scopeEntryId** | **string** | In-scope asset id from the engagement&#39;s live scope | 
 **severity** | **string** | Claimed severity; empty derives it from the CVSS metrics | 
 **steps** | **string** | Reproduction steps, appended to the description | 
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

> ReportEvidence PostReportsEvidence(ctx, reportId).Execute()

Attach a multipart evidence file to an own draft report; scanned before it can download

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
	resp, r, err := apiClient.ReportsAPI.PostReportsEvidence(context.Background(), reportId).Execute()
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


### Return type

[**ReportEvidence**](ReportEvidence.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
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

> Report PostReportsSubmit(ctx, reportId).Description(description).Origin(origin).Title(title).Ac(ac).Av(av).Avail(avail).Conf(conf).Endpoint(endpoint).Impact(impact).Integ(integ).Poc(poc).Pr(pr).Scope(scope).ScopeEntryId(scopeEntryId).Severity(severity).Steps(steps).Ui(ui).VulnerabilityType(vulnerabilityType).Execute()

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
	resp, r, err := apiClient.ReportsAPI.PostReportsSubmit(context.Background(), reportId).Description(description).Origin(origin).Title(title).Ac(ac).Av(av).Avail(avail).Conf(conf).Endpoint(endpoint).Impact(impact).Integ(integ).Poc(poc).Pr(pr).Scope(scope).ScopeEntryId(scopeEntryId).Severity(severity).Steps(steps).Ui(ui).VulnerabilityType(vulnerabilityType).Execute()
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

