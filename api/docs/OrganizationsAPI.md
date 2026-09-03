# \OrganizationsAPI

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationsAuditLog**](OrganizationsAPI.md#GetOrganizationsAuditLog) | **Get** /organizations/{org_slug}/audit-log | The organization&#39;s audit log, admin-only; filterable by category, actor, action, q, range or from/to, paginated with page
[**GetOrganizationsAuditLogByEventId**](OrganizationsAPI.md#GetOrganizationsAuditLogByEventId) | **Get** /organizations/{org_slug}/audit-log/{event_id} | One audit event of the organization with its metadata, admin-only
[**GetOrganizationsAuditLogExport**](OrganizationsAPI.md#GetOrganizationsAuditLogExport) | **Get** /organizations/{org_slug}/audit-log/export | Export the filtered audit log as CSV or JSON (format param), admin-only; every export is itself audited
[**GetOrganizationsByOrgSlug**](OrganizationsAPI.md#GetOrganizationsByOrgSlug) | **Get** /organizations/{org_slug} | The organization security-posture dashboard: open/resolved findings, severity and engagement buckets, response medians, SLA compliance, and bounty spend
[**GetOrganizationsDisclosures**](OrganizationsAPI.md#GetOrganizationsDisclosures) | **Get** /organizations/{org_slug}/disclosures | The organization&#39;s coordinated disclosures and their timeline state
[**GetOrganizationsDocuments**](OrganizationsAPI.md#GetOrganizationsDocuments) | **Get** /organizations/{org_slug}/documents | The organization&#39;s document vault and findings exports, members-only
[**GetOrganizationsEngagements**](OrganizationsAPI.md#GetOrganizationsEngagements) | **Get** /organizations/{org_slug}/engagements | All engagements of the organization, members-only
[**GetOrganizationsEngagementsByEngagementSlug**](OrganizationsAPI.md#GetOrganizationsEngagementsByEngagementSlug) | **Get** /organizations/{org_slug}/engagements/{engagement_slug} | Engagement content lifecycle: live and draft versions, changes, and history
[**GetOrganizationsEngagementsReports**](OrganizationsAPI.md#GetOrganizationsEngagementsReports) | **Get** /organizations/{org_slug}/engagements/{engagement_slug}/reports | The engagement&#39;s submitted reports: filterable by status, severity, assignment, and search, sortable, paginated (view&#x3D;board renders the kanban); drafts never appear
[**GetOrganizationsEngagementsReportsByReportId**](OrganizationsAPI.md#GetOrganizationsEngagementsReportsByReportId) | **Get** /organizations/{org_slug}/engagements/{engagement_slug}/reports/{report_id} | One submitted report of the engagement with its submission content, decision context, and status history; 404 for drafts and foreign reports (ADR-0009)
[**GetOrganizationsEngagementsStatistics**](OrganizationsAPI.md#GetOrganizationsEngagementsStatistics) | **Get** /organizations/{org_slug}/engagements/{engagement_slug}/statistics | One engagement&#39;s statistics: report volume, severity distribution, top researchers, and budget against the cost limit
[**GetOrganizationsIntegrations**](OrganizationsAPI.md#GetOrganizationsIntegrations) | **Get** /organizations/{org_slug}/integrations | The organization&#39;s webhooks with recent deliveries; admins only
[**GetOrganizationsMembers**](OrganizationsAPI.md#GetOrganizationsMembers) | **Get** /organizations/{org_slug}/members | The organization&#39;s members and pending invites, admin-only
[**GetOrganizationsMembersExport**](OrganizationsAPI.md#GetOrganizationsMembersExport) | **Get** /organizations/{org_slug}/members/export | Download the organization&#39;s members as CSV or JSON, admin-only
[**GetOrganizationsMembersInvites**](OrganizationsAPI.md#GetOrganizationsMembersInvites) | **Get** /organizations/{org_slug}/members/invites | One page of the organization&#39;s pending invites, admin-only
[**GetOrganizationsReports**](OrganizationsAPI.md#GetOrganizationsReports) | **Get** /organizations/{org_slug}/reports | Every submitted report of the organization across its engagements: filterable by status, engagement, severity, assignment, and search, sortable, paginated
[**GetOrganizationsReportsByReportId**](OrganizationsAPI.md#GetOrganizationsReportsByReportId) | **Get** /organizations/{org_slug}/reports/{report_id} | One submitted report of the organization, opened from the organization-wide list
[**GetOrganizationsScopes**](OrganizationsAPI.md#GetOrganizationsScopes) | **Get** /organizations/{org_slug}/scopes | The organization&#39;s asset inventory and per-engagement scope, members-only
[**GetOrganizationsScopesAssetsByAssetId**](OrganizationsAPI.md#GetOrganizationsScopesAssetsByAssetId) | **Get** /organizations/{org_slug}/scopes/assets/{asset_id} | One inventory asset of the organization
[**GetOrganizationsScopesEngagementsByEngagementSlug**](OrganizationsAPI.md#GetOrganizationsScopesEngagementsByEngagementSlug) | **Get** /organizations/{org_slug}/scopes/engagements/{engagement_slug} | Scope entries of the engagement&#39;s open version, or the live version read-only
[**GetOrganizationsScopesEngagementsEntriesByEntryId**](OrganizationsAPI.md#GetOrganizationsScopesEngagementsEntriesByEntryId) | **Get** /organizations/{org_slug}/scopes/engagements/{engagement_slug}/entries/{entry_id} | One scope entry of the engagement&#39;s editable draft version
[**GetOrganizationsSettings**](OrganizationsAPI.md#GetOrganizationsSettings) | **Get** /organizations/{org_slug}/settings | The organization&#39;s general settings: live profile values plus the open draft, admin-only
[**GetOrganizationsVerification**](OrganizationsAPI.md#GetOrganizationsVerification) | **Get** /organizations/{org_slug}/verification | The organization&#39;s verification and terms status; admins answer information requests here
[**GetOrganizationsWallet**](OrganizationsAPI.md#GetOrganizationsWallet) | **Get** /organizations/{org_slug}/wallet | The organization&#39;s prepaid wallet: balance, lifetime stats, and the filtered ledger; finance role only
[**GetOrganizationsWalletExport**](OrganizationsAPI.md#GetOrganizationsWalletExport) | **Get** /organizations/{org_slug}/wallet/export | Export the filtered wallet ledger as CSV or JSON; finance role only
[**PostOrganizationsConductReportFlag**](OrganizationsAPI.md#PostOrganizationsConductReportFlag) | **Post** /organizations/{org_slug}/conduct/report/{report_id}/flag | Report a researcher to trust &amp; safety from a report of the organization
[**PostOrganizationsDisclosuresAccept**](OrganizationsAPI.md#PostOrganizationsDisclosuresAccept) | **Post** /organizations/{org_slug}/disclosures/{disclosure_id}/accept | Accept a disclosure request; moves it into coordination
[**PostOrganizationsDisclosuresDecline**](OrganizationsAPI.md#PostOrganizationsDisclosuresDecline) | **Post** /organizations/{org_slug}/disclosures/{disclosure_id}/decline | Decline a disclosure request or withdraw a scheduled one
[**PostOrganizationsDisclosuresPublish**](OrganizationsAPI.md#PostOrganizationsDisclosuresPublish) | **Post** /organizations/{org_slug}/disclosures/{disclosure_id}/publish | Publish the coordinated advisory; moves it to published
[**PostOrganizationsDisclosuresSchedule**](OrganizationsAPI.md#PostOrganizationsDisclosuresSchedule) | **Post** /organizations/{org_slug}/disclosures/{disclosure_id}/schedule | Agree the disclosure date; moves it to scheduled
[**PostOrganizationsDocuments**](OrganizationsAPI.md#PostOrganizationsDocuments) | **Post** /organizations/{org_slug}/documents | Upload a document (multipart field &#39;file&#39;); re-uploading a name appends a version
[**PostOrganizationsDocumentsExports**](OrganizationsAPI.md#PostOrganizationsDocumentsExports) | **Post** /organizations/{org_slug}/documents/exports | Queue a findings export; generation is asynchronous
[**PostOrganizationsDocumentsExportsRetry**](OrganizationsAPI.md#PostOrganizationsDocumentsExportsRetry) | **Post** /organizations/{org_slug}/documents/exports/{export_id}/retry | Re-queue a failed findings export (S0351); only exports in the failed state can be retried
[**PostOrganizationsDocumentsSchedules**](OrganizationsAPI.md#PostOrganizationsDocumentsSchedules) | **Post** /organizations/{org_slug}/documents/schedules | Create a recurring findings-export schedule (daily, weekly, or monthly); the first run happens on the next sweep
[**PostOrganizationsDocumentsSchedulesDelete**](OrganizationsAPI.md#PostOrganizationsDocumentsSchedulesDelete) | **Post** /organizations/{org_slug}/documents/schedules/{schedule_id}/delete | Delete an export schedule; past runs keep their history
[**PostOrganizationsDocumentsSchedulesPause**](OrganizationsAPI.md#PostOrganizationsDocumentsSchedulesPause) | **Post** /organizations/{org_slug}/documents/schedules/{schedule_id}/pause | Pause an export schedule
[**PostOrganizationsDocumentsSchedulesResume**](OrganizationsAPI.md#PostOrganizationsDocumentsSchedulesResume) | **Post** /organizations/{org_slug}/documents/schedules/{schedule_id}/resume | Resume a paused export schedule
[**PostOrganizationsDocumentsSchedulesRun**](OrganizationsAPI.md#PostOrganizationsDocumentsSchedulesRun) | **Post** /organizations/{org_slug}/documents/schedules/{schedule_id}/run | Run an export schedule immediately; also restarts its cadence interval
[**PostOrganizationsEngagements**](OrganizationsAPI.md#PostOrganizationsEngagements) | **Post** /organizations/{org_slug}/engagements | Create a draft bug-bounty engagement; admins only, the slug derives from the title
[**PostOrganizationsEngagementsClose**](OrganizationsAPI.md#PostOrganizationsEngagementsClose) | **Post** /organizations/{org_slug}/engagements/{engagement_slug}/close | Close the engagement permanently: submissions stop, history stays readable; 409 unless active or paused
[**PostOrganizationsEngagementsContent**](OrganizationsAPI.md#PostOrganizationsEngagementsContent) | **Post** /organizations/{org_slug}/engagements/{engagement_slug}/content | Save the engagement&#39;s draft content (summary, testing rules, qualifying vulnerabilities, procedure, safe harbor); opens a draft version when none is open; 409 while the open version is in review
[**PostOrganizationsEngagementsMessageProxy**](OrganizationsAPI.md#PostOrganizationsEngagementsMessageProxy) | **Post** /organizations/{org_slug}/engagements/{engagement_slug}/message-proxy | Configure the SMS message proxy: set the gateway sender number, generating the webhook secret on first enable
[**PostOrganizationsEngagementsMessageProxyDisable**](OrganizationsAPI.md#PostOrganizationsEngagementsMessageProxyDisable) | **Post** /organizations/{org_slug}/engagements/{engagement_slug}/message-proxy/disable | Disable the SMS message proxy and discard the webhook secret
[**PostOrganizationsEngagementsPause**](OrganizationsAPI.md#PostOrganizationsEngagementsPause) | **Post** /organizations/{org_slug}/engagements/{engagement_slug}/pause | Pause a live engagement: new submissions stop; 409 unless active
[**PostOrganizationsEngagementsReopen**](OrganizationsAPI.md#PostOrganizationsEngagementsReopen) | **Post** /organizations/{org_slug}/engagements/{engagement_slug}/reopen | Reopen a changes-requested version as an editable draft; 409 unless changes are requested
[**PostOrganizationsEngagementsReportsAccept**](OrganizationsAPI.md#PostOrganizationsEngagementsReportsAccept) | **Post** /organizations/{org_slug}/engagements/{engagement_slug}/reports/{report_id}/accept | Accept a validated report; 409 for any other workflow state
[**PostOrganizationsEngagementsReportsAcknowledgeEscalation**](OrganizationsAPI.md#PostOrganizationsEngagementsReportsAcknowledgeEscalation) | **Post** /organizations/{org_slug}/engagements/{engagement_slug}/reports/{report_id}/acknowledge-escalation | Acknowledge an escalated report; 409 when it is not escalated or already acknowledged
[**PostOrganizationsEngagementsReportsAssign**](OrganizationsAPI.md#PostOrganizationsEngagementsReportsAssign) | **Post** /organizations/{org_slug}/engagements/{engagement_slug}/reports/{report_id}/assign | Set or clear the organization&#39;s internal owner of the report by user id, one of the detail&#39;s assignmentOptions; carries no permission; 400 names the assignee field when the id is not an active member
[**PostOrganizationsEngagementsReportsComments**](OrganizationsAPI.md#PostOrganizationsEngagementsReportsComments) | **Post** /organizations/{org_slug}/engagements/{engagement_slug}/reports/{report_id}/comments | Post a message on the report&#39;s thread: audience &#39;all&#39; (researcher-visible question) or &#39;organization&#39; (internal note)
[**PostOrganizationsEngagementsReportsReject**](OrganizationsAPI.md#PostOrganizationsEngagementsReportsReject) | **Post** /organizations/{org_slug}/engagements/{engagement_slug}/reports/{report_id}/reject | Reject a validated report with a mandatory reason; 409 for any other workflow state
[**PostOrganizationsEngagementsReportsRetests**](OrganizationsAPI.md#PostOrganizationsEngagementsReportsRetests) | **Post** /organizations/{org_slug}/engagements/{engagement_slug}/reports/{report_id}/retests | Mark the accepted or reopened report fixed and request a retest with remediation notes; 409 when one is already open
[**PostOrganizationsEngagementsReportsRetestsCancel**](OrganizationsAPI.md#PostOrganizationsEngagementsReportsRetestsCancel) | **Post** /organizations/{org_slug}/engagements/{engagement_slug}/reports/{report_id}/retests/{retest_id}/cancel | Cancel an open retest request; 409 once the verifier started, the report stays fixed
[**PostOrganizationsEngagementsResume**](OrganizationsAPI.md#PostOrganizationsEngagementsResume) | **Post** /organizations/{org_slug}/engagements/{engagement_slug}/resume | Resume a paused engagement; 409 unless paused
[**PostOrganizationsEngagementsSubmit**](OrganizationsAPI.md#PostOrganizationsEngagementsSubmit) | **Post** /organizations/{org_slug}/engagements/{engagement_slug}/submit | Send the open draft version to backoffice review with an optional note; 400 until the draft has an in-scope asset; 409 without an open version or when it is not a draft
[**PostOrganizationsEngagementsVerificationLevel**](OrganizationsAPI.md#PostOrganizationsEngagementsVerificationLevel) | **Post** /organizations/{org_slug}/engagements/{engagement_slug}/verification-level | Set the researcher verification level the engagement requires; immediate and audited
[**PostOrganizationsEngagementsVisibility**](OrganizationsAPI.md#PostOrganizationsEngagementsVisibility) | **Post** /organizations/{org_slug}/engagements/{engagement_slug}/visibility | Switch the engagement between private and public; immediate and audited
[**PostOrganizationsIntegrationsWebhooks**](OrganizationsAPI.md#PostOrganizationsIntegrationsWebhooks) | **Post** /organizations/{org_slug}/integrations/webhooks | Register a signed webhook endpoint subscribed to event types
[**PostOrganizationsIntegrationsWebhooksByWebhookId**](OrganizationsAPI.md#PostOrganizationsIntegrationsWebhooksByWebhookId) | **Post** /organizations/{org_slug}/integrations/webhooks/{webhook_id} | Update a webhook&#39;s URL, events, and active flag
[**PostOrganizationsIntegrationsWebhooksDelete**](OrganizationsAPI.md#PostOrganizationsIntegrationsWebhooksDelete) | **Post** /organizations/{org_slug}/integrations/webhooks/{webhook_id}/delete | Delete a webhook
[**PostOrganizationsIntegrationsWebhooksTest**](OrganizationsAPI.md#PostOrganizationsIntegrationsWebhooksTest) | **Post** /organizations/{org_slug}/integrations/webhooks/{webhook_id}/test | Send a signed test (ping) delivery to the webhook
[**PostOrganizationsMembersByMemberId**](OrganizationsAPI.md#PostOrganizationsMembersByMemberId) | **Post** /organizations/{org_slug}/members/{member_id} | Update a member&#39;s category flags and program grants, admin-only; refuses demoting the last admin
[**PostOrganizationsMembersInvite**](OrganizationsAPI.md#PostOrganizationsMembersInvite) | **Post** /organizations/{org_slug}/members/invite | Invite a member by email with category flags and program grants, admin-only
[**PostOrganizationsScopesAssets**](OrganizationsAPI.md#PostOrganizationsScopesAssets) | **Post** /organizations/{org_slug}/scopes/assets | Add an asset to the organization&#39;s inventory
[**PostOrganizationsScopesAssetsByAssetId**](OrganizationsAPI.md#PostOrganizationsScopesAssetsByAssetId) | **Post** /organizations/{org_slug}/scopes/assets/{asset_id} | Update an inventory asset; the identifier must stay unique within the organization
[**PostOrganizationsScopesEngagementsDraft**](OrganizationsAPI.md#PostOrganizationsScopesEngagementsDraft) | **Post** /organizations/{org_slug}/scopes/engagements/{engagement_slug}/draft | Open a draft engagement version so the live scope becomes editable
[**PostOrganizationsScopesEngagementsEntries**](OrganizationsAPI.md#PostOrganizationsScopesEngagementsEntries) | **Post** /organizations/{org_slug}/scopes/engagements/{engagement_slug}/entries | Add a scope entry to the engagement&#39;s draft version; adding the first entry creates the draft
[**PostOrganizationsScopesEngagementsEntriesByEntryId**](OrganizationsAPI.md#PostOrganizationsScopesEngagementsEntriesByEntryId) | **Post** /organizations/{org_slug}/scopes/engagements/{engagement_slug}/entries/{entry_id} | Update a scope entry on the engagement&#39;s editable draft version; 409 when no editable draft is open
[**PostOrganizationsScopesEngagementsEntriesDelete**](OrganizationsAPI.md#PostOrganizationsScopesEngagementsEntriesDelete) | **Post** /organizations/{org_slug}/scopes/engagements/{engagement_slug}/entries/{entry_id}/delete | Delete a scope entry from the engagement&#39;s editable draft version
[**PostOrganizationsSettings**](OrganizationsAPI.md#PostOrganizationsSettings) | **Post** /organizations/{org_slug}/settings | Save the organization profile draft (name, website, industry, brand colors), admin-only; 409 while the draft is frozen in review
[**PostOrganizationsSettingsLogoDelete**](OrganizationsAPI.md#PostOrganizationsSettingsLogoDelete) | **Post** /organizations/{org_slug}/settings/logo/delete | Remove the logo from the profile draft; the pages fall back to the monogram avatar
[**PostOrganizationsSettingsReopen**](OrganizationsAPI.md#PostOrganizationsSettingsReopen) | **Post** /organizations/{org_slug}/settings/reopen | Reopen a changes-requested profile version as a draft; 409 unless changes are requested
[**PostOrganizationsSettingsSubmit**](OrganizationsAPI.md#PostOrganizationsSettingsSubmit) | **Post** /organizations/{org_slug}/settings/submit | Submit the profile draft for backoffice review; requires a non-empty name, 409 without an open draft
[**PostOrganizationsVerification**](OrganizationsAPI.md#PostOrganizationsVerification) | **Post** /organizations/{org_slug}/verification | Answer an information request with updated legal details; returns the registration to the review queue
[**PostOrganizationsWalletRewardsApprove**](OrganizationsAPI.md#PostOrganizationsWalletRewardsApprove) | **Post** /organizations/{org_slug}/wallet/rewards/{id}/approve | Approve a proposed bounty: debit the wallet and mark it approved; finance role only
[**PostOrganizationsWalletRewardsReject**](OrganizationsAPI.md#PostOrganizationsWalletRewardsReject) | **Post** /organizations/{org_slug}/wallet/rewards/{id}/reject | Reject a proposed bounty with a reason; finance role only
[**PostOrganizationsWalletTopups**](OrganizationsAPI.md#PostOrganizationsWalletTopups) | **Post** /organizations/{org_slug}/wallet/topups | Request a bank-transfer wallet top-up; returns the payment reference to quote



## GetOrganizationsAuditLog

> AuditEventList GetOrganizationsAuditLog(ctx, orgSlug).Execute()

The organization's audit log, admin-only; filterable by category, actor, action, q, range or from/to, paginated with page

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.GetOrganizationsAuditLog(context.Background(), orgSlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.GetOrganizationsAuditLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationsAuditLog`: AuditEventList
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.GetOrganizationsAuditLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationsAuditLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuditEventList**](AuditEventList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationsAuditLogByEventId

> AuditEvent GetOrganizationsAuditLogByEventId(ctx, orgSlug, eventId).Execute()

One audit event of the organization with its metadata, admin-only

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
	eventId := "eventId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.GetOrganizationsAuditLogByEventId(context.Background(), orgSlug, eventId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.GetOrganizationsAuditLogByEventId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationsAuditLogByEventId`: AuditEvent
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.GetOrganizationsAuditLogByEventId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 
**eventId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationsAuditLogByEventIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AuditEvent**](AuditEvent.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationsAuditLogExport

> GetOrganizationsAuditLogExport(ctx, orgSlug).Execute()

Export the filtered audit log as CSV or JSON (format param), admin-only; every export is itself audited

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.GetOrganizationsAuditLogExport(context.Background(), orgSlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.GetOrganizationsAuditLogExport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationsAuditLogExportRequest struct via the builder pattern


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


## GetOrganizationsByOrgSlug

> DashboardOrgDashboard GetOrganizationsByOrgSlug(ctx, orgSlug).Execute()

The organization security-posture dashboard: open/resolved findings, severity and engagement buckets, response medians, SLA compliance, and bounty spend

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.GetOrganizationsByOrgSlug(context.Background(), orgSlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.GetOrganizationsByOrgSlug``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationsByOrgSlug`: DashboardOrgDashboard
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.GetOrganizationsByOrgSlug`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationsByOrgSlugRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DashboardOrgDashboard**](DashboardOrgDashboard.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationsDisclosures

> DisclosureOrgDisclosureList GetOrganizationsDisclosures(ctx, orgSlug).Execute()

The organization's coordinated disclosures and their timeline state

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.GetOrganizationsDisclosures(context.Background(), orgSlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.GetOrganizationsDisclosures``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationsDisclosures`: DisclosureOrgDisclosureList
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.GetOrganizationsDisclosures`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationsDisclosuresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DisclosureOrgDisclosureList**](DisclosureOrgDisclosureList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationsDocuments

> DocumentVault GetOrganizationsDocuments(ctx, orgSlug).Execute()

The organization's document vault and findings exports, members-only

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.GetOrganizationsDocuments(context.Background(), orgSlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.GetOrganizationsDocuments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationsDocuments`: DocumentVault
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.GetOrganizationsDocuments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationsDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DocumentVault**](DocumentVault.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationsEngagements

> EngagementListItemList GetOrganizationsEngagements(ctx, orgSlug).Execute()

All engagements of the organization, members-only

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.GetOrganizationsEngagements(context.Background(), orgSlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.GetOrganizationsEngagements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationsEngagements`: EngagementListItemList
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.GetOrganizationsEngagements`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationsEngagementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetOrganizationsEngagementsByEngagementSlug

> EngagementLifecycle GetOrganizationsEngagementsByEngagementSlug(ctx, orgSlug, engagementSlug).Execute()

Engagement content lifecycle: live and draft versions, changes, and history

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
	resp, r, err := apiClient.OrganizationsAPI.GetOrganizationsEngagementsByEngagementSlug(context.Background(), orgSlug, engagementSlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.GetOrganizationsEngagementsByEngagementSlug``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationsEngagementsByEngagementSlug`: EngagementLifecycle
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.GetOrganizationsEngagementsByEngagementSlug`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 
**engagementSlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationsEngagementsByEngagementSlugRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EngagementLifecycle**](EngagementLifecycle.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationsEngagementsReports

> ReportOrgReportList GetOrganizationsEngagementsReports(ctx, orgSlug, engagementSlug).Status(status).Severity(severity).Q(q).Sort(sort).Dir(dir).Page(page).Size(size).Assignment(assignment).View(view).Execute()

The engagement's submitted reports: filterable by status, severity, assignment, and search, sortable, paginated (view=board renders the kanban); drafts never appear

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
	status := []string{"Status_example"} // []string | Workflow states to include (optional)
	severity := []string{"Severity_example"} // []string | Severities (assessed, else submitted) to include (optional)
	q := "q_example" // string | Free-text search: title substring or report id prefix (optional)
	sort := "sort_example" // string | Sort key (optional)
	dir := "dir_example" // string | Sort direction (optional)
	page := "page_example" // string | Page number, from 1 (optional)
	size := "size_example" // string | Rows per page (optional)
	assignment := "assignment_example" // string | Assignee facet: any, unassigned, mine, or user:<id> (optional)
	view := "view_example" // string | HTML view (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.GetOrganizationsEngagementsReports(context.Background(), orgSlug, engagementSlug).Status(status).Severity(severity).Q(q).Sort(sort).Dir(dir).Page(page).Size(size).Assignment(assignment).View(view).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.GetOrganizationsEngagementsReports``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationsEngagementsReports`: ReportOrgReportList
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.GetOrganizationsEngagementsReports`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 
**engagementSlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationsEngagementsReportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **status** | **[]string** | Workflow states to include | 
 **severity** | **[]string** | Severities (assessed, else submitted) to include | 
 **q** | **string** | Free-text search: title substring or report id prefix | 
 **sort** | **string** | Sort key | 
 **dir** | **string** | Sort direction | 
 **page** | **string** | Page number, from 1 | 
 **size** | **string** | Rows per page | 
 **assignment** | **string** | Assignee facet: any, unassigned, mine, or user:&lt;id&gt; | 
 **view** | **string** | HTML view | 

### Return type

[**ReportOrgReportList**](ReportOrgReportList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationsEngagementsReportsByReportId

> ReportOrgReportDetail GetOrganizationsEngagementsReportsByReportId(ctx, orgSlug, engagementSlug, reportId).Execute()

One submitted report of the engagement with its submission content, decision context, and status history; 404 for drafts and foreign reports (ADR-0009)

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
	reportId := "reportId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.GetOrganizationsEngagementsReportsByReportId(context.Background(), orgSlug, engagementSlug, reportId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.GetOrganizationsEngagementsReportsByReportId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationsEngagementsReportsByReportId`: ReportOrgReportDetail
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.GetOrganizationsEngagementsReportsByReportId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 
**engagementSlug** | **string** |  | 
**reportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationsEngagementsReportsByReportIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ReportOrgReportDetail**](ReportOrgReportDetail.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationsEngagementsStatistics

> DashboardEngagementStats GetOrganizationsEngagementsStatistics(ctx, orgSlug, engagementSlug).Execute()

One engagement's statistics: report volume, severity distribution, top researchers, and budget against the cost limit

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
	resp, r, err := apiClient.OrganizationsAPI.GetOrganizationsEngagementsStatistics(context.Background(), orgSlug, engagementSlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.GetOrganizationsEngagementsStatistics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationsEngagementsStatistics`: DashboardEngagementStats
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.GetOrganizationsEngagementsStatistics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 
**engagementSlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationsEngagementsStatisticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DashboardEngagementStats**](DashboardEngagementStats.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationsIntegrations

> Integrations GetOrganizationsIntegrations(ctx, orgSlug).Execute()

The organization's webhooks with recent deliveries; admins only

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.GetOrganizationsIntegrations(context.Background(), orgSlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.GetOrganizationsIntegrations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationsIntegrations`: Integrations
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.GetOrganizationsIntegrations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationsIntegrationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Integrations**](Integrations.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationsMembers

> OrganizationMembersList GetOrganizationsMembers(ctx, orgSlug).Execute()

The organization's members and pending invites, admin-only

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.GetOrganizationsMembers(context.Background(), orgSlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.GetOrganizationsMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationsMembers`: OrganizationMembersList
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.GetOrganizationsMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationsMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrganizationMembersList**](OrganizationMembersList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationsMembersExport

> GetOrganizationsMembersExport(ctx, orgSlug).Execute()

Download the organization's members as CSV or JSON, admin-only

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.GetOrganizationsMembersExport(context.Background(), orgSlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.GetOrganizationsMembersExport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationsMembersExportRequest struct via the builder pattern


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


## GetOrganizationsMembersInvites

> OrganizationInvitesList GetOrganizationsMembersInvites(ctx, orgSlug).Execute()

One page of the organization's pending invites, admin-only

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.GetOrganizationsMembersInvites(context.Background(), orgSlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.GetOrganizationsMembersInvites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationsMembersInvites`: OrganizationInvitesList
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.GetOrganizationsMembersInvites`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationsMembersInvitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrganizationInvitesList**](OrganizationInvitesList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationsReports

> ReportOrgReportList GetOrganizationsReports(ctx, orgSlug).Status(status).Severity(severity).Q(q).Sort(sort).Dir(dir).Page(page).Size(size).Engagement(engagement).Assignment(assignment).View(view).Execute()

Every submitted report of the organization across its engagements: filterable by status, engagement, severity, assignment, and search, sortable, paginated

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
	status := []string{"Status_example"} // []string | Workflow states to include (optional)
	severity := []string{"Severity_example"} // []string | Severities (assessed, else submitted) to include (optional)
	q := "q_example" // string | Free-text search: title substring or report id prefix (optional)
	sort := "sort_example" // string | Sort key (optional)
	dir := "dir_example" // string | Sort direction (optional)
	page := "page_example" // string | Page number, from 1 (optional)
	size := "size_example" // string | Rows per page (optional)
	engagement := "engagement_example" // string | Engagement as org_slug/engagement_slug (optional)
	assignment := "assignment_example" // string | Assignee facet: any, unassigned, mine, or user:<id> (optional)
	view := "view_example" // string | HTML view (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.GetOrganizationsReports(context.Background(), orgSlug).Status(status).Severity(severity).Q(q).Sort(sort).Dir(dir).Page(page).Size(size).Engagement(engagement).Assignment(assignment).View(view).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.GetOrganizationsReports``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationsReports`: ReportOrgReportList
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.GetOrganizationsReports`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationsReportsRequest struct via the builder pattern


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
 **assignment** | **string** | Assignee facet: any, unassigned, mine, or user:&lt;id&gt; | 
 **view** | **string** | HTML view | 

### Return type

[**ReportOrgReportList**](ReportOrgReportList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationsReportsByReportId

> ReportOrgReportDetail GetOrganizationsReportsByReportId(ctx, orgSlug, reportId).Execute()

One submitted report of the organization, opened from the organization-wide list

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
	reportId := "reportId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.GetOrganizationsReportsByReportId(context.Background(), orgSlug, reportId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.GetOrganizationsReportsByReportId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationsReportsByReportId`: ReportOrgReportDetail
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.GetOrganizationsReportsByReportId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 
**reportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationsReportsByReportIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ReportOrgReportDetail**](ReportOrgReportDetail.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationsScopes

> Scopes GetOrganizationsScopes(ctx, orgSlug).Execute()

The organization's asset inventory and per-engagement scope, members-only

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.GetOrganizationsScopes(context.Background(), orgSlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.GetOrganizationsScopes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationsScopes`: Scopes
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.GetOrganizationsScopes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationsScopesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Scopes**](Scopes.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationsScopesAssetsByAssetId

> ScopeAsset GetOrganizationsScopesAssetsByAssetId(ctx, orgSlug, assetId).Execute()

One inventory asset of the organization

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
	assetId := "assetId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.GetOrganizationsScopesAssetsByAssetId(context.Background(), orgSlug, assetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.GetOrganizationsScopesAssetsByAssetId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationsScopesAssetsByAssetId`: ScopeAsset
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.GetOrganizationsScopesAssetsByAssetId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 
**assetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationsScopesAssetsByAssetIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ScopeAsset**](ScopeAsset.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationsScopesEngagementsByEngagementSlug

> ScopeEngagementScope GetOrganizationsScopesEngagementsByEngagementSlug(ctx, orgSlug, engagementSlug).Execute()

Scope entries of the engagement's open version, or the live version read-only

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
	resp, r, err := apiClient.OrganizationsAPI.GetOrganizationsScopesEngagementsByEngagementSlug(context.Background(), orgSlug, engagementSlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.GetOrganizationsScopesEngagementsByEngagementSlug``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationsScopesEngagementsByEngagementSlug`: ScopeEngagementScope
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.GetOrganizationsScopesEngagementsByEngagementSlug`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 
**engagementSlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationsScopesEngagementsByEngagementSlugRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ScopeEngagementScope**](ScopeEngagementScope.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationsScopesEngagementsEntriesByEntryId

> ScopeEntry GetOrganizationsScopesEngagementsEntriesByEntryId(ctx, orgSlug, engagementSlug, entryId).Execute()

One scope entry of the engagement's editable draft version

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
	entryId := "entryId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.GetOrganizationsScopesEngagementsEntriesByEntryId(context.Background(), orgSlug, engagementSlug, entryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.GetOrganizationsScopesEngagementsEntriesByEntryId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationsScopesEngagementsEntriesByEntryId`: ScopeEntry
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.GetOrganizationsScopesEngagementsEntriesByEntryId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 
**engagementSlug** | **string** |  | 
**entryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationsScopesEngagementsEntriesByEntryIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ScopeEntry**](ScopeEntry.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationsSettings

> OrganizationSettings GetOrganizationsSettings(ctx, orgSlug).Execute()

The organization's general settings: live profile values plus the open draft, admin-only

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.GetOrganizationsSettings(context.Background(), orgSlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.GetOrganizationsSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationsSettings`: OrganizationSettings
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.GetOrganizationsSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationsSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrganizationSettings**](OrganizationSettings.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationsVerification

> GetOrganizationsVerification(ctx, orgSlug).Execute()

The organization's verification and terms status; admins answer information requests here

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.GetOrganizationsVerification(context.Background(), orgSlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.GetOrganizationsVerification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationsVerificationRequest struct via the builder pattern


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


## GetOrganizationsWallet

> PaymentTransactionList GetOrganizationsWallet(ctx, orgSlug).Execute()

The organization's prepaid wallet: balance, lifetime stats, and the filtered ledger; finance role only

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.GetOrganizationsWallet(context.Background(), orgSlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.GetOrganizationsWallet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationsWallet`: PaymentTransactionList
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.GetOrganizationsWallet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationsWalletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PaymentTransactionList**](PaymentTransactionList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationsWalletExport

> GetOrganizationsWalletExport(ctx, orgSlug).Execute()

Export the filtered wallet ledger as CSV or JSON; finance role only

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.GetOrganizationsWalletExport(context.Background(), orgSlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.GetOrganizationsWalletExport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationsWalletExportRequest struct via the builder pattern


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


## PostOrganizationsConductReportFlag

> PostOrganizationsConductReportFlag(ctx, orgSlug, reportId).Execute()

Report a researcher to trust & safety from a report of the organization

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
	reportId := "reportId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.PostOrganizationsConductReportFlag(context.Background(), orgSlug, reportId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.PostOrganizationsConductReportFlag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 
**reportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrganizationsConductReportFlagRequest struct via the builder pattern


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


## PostOrganizationsDisclosuresAccept

> PostOrganizationsDisclosuresAccept(ctx, orgSlug, disclosureId).Execute()

Accept a disclosure request; moves it into coordination

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
	disclosureId := "disclosureId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.PostOrganizationsDisclosuresAccept(context.Background(), orgSlug, disclosureId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.PostOrganizationsDisclosuresAccept``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 
**disclosureId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrganizationsDisclosuresAcceptRequest struct via the builder pattern


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


## PostOrganizationsDisclosuresDecline

> PostOrganizationsDisclosuresDecline(ctx, orgSlug, disclosureId).Execute()

Decline a disclosure request or withdraw a scheduled one

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
	disclosureId := "disclosureId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.PostOrganizationsDisclosuresDecline(context.Background(), orgSlug, disclosureId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.PostOrganizationsDisclosuresDecline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 
**disclosureId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrganizationsDisclosuresDeclineRequest struct via the builder pattern


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


## PostOrganizationsDisclosuresPublish

> PostOrganizationsDisclosuresPublish(ctx, orgSlug, disclosureId).Execute()

Publish the coordinated advisory; moves it to published

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
	disclosureId := "disclosureId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.PostOrganizationsDisclosuresPublish(context.Background(), orgSlug, disclosureId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.PostOrganizationsDisclosuresPublish``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 
**disclosureId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrganizationsDisclosuresPublishRequest struct via the builder pattern


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


## PostOrganizationsDisclosuresSchedule

> PostOrganizationsDisclosuresSchedule(ctx, orgSlug, disclosureId).Execute()

Agree the disclosure date; moves it to scheduled

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
	disclosureId := "disclosureId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.PostOrganizationsDisclosuresSchedule(context.Background(), orgSlug, disclosureId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.PostOrganizationsDisclosuresSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 
**disclosureId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrganizationsDisclosuresScheduleRequest struct via the builder pattern


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


## PostOrganizationsDocuments

> Document PostOrganizationsDocuments(ctx, orgSlug).Execute()

Upload a document (multipart field 'file'); re-uploading a name appends a version

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.PostOrganizationsDocuments(context.Background(), orgSlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.PostOrganizationsDocuments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostOrganizationsDocuments`: Document
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.PostOrganizationsDocuments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrganizationsDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Document**](Document.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOrganizationsDocumentsExports

> DocumentExport PostOrganizationsDocumentsExports(ctx, orgSlug).Execute()

Queue a findings export; generation is asynchronous

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.PostOrganizationsDocumentsExports(context.Background(), orgSlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.PostOrganizationsDocumentsExports``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostOrganizationsDocumentsExports`: DocumentExport
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.PostOrganizationsDocumentsExports`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrganizationsDocumentsExportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DocumentExport**](DocumentExport.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOrganizationsDocumentsExportsRetry

> PostOrganizationsDocumentsExportsRetry(ctx, orgSlug, exportId).Execute()

Re-queue a failed findings export (S0351); only exports in the failed state can be retried

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
	exportId := "exportId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.PostOrganizationsDocumentsExportsRetry(context.Background(), orgSlug, exportId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.PostOrganizationsDocumentsExportsRetry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 
**exportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrganizationsDocumentsExportsRetryRequest struct via the builder pattern


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


## PostOrganizationsDocumentsSchedules

> PostOrganizationsDocumentsSchedules(ctx, orgSlug).Execute()

Create a recurring findings-export schedule (daily, weekly, or monthly); the first run happens on the next sweep

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.PostOrganizationsDocumentsSchedules(context.Background(), orgSlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.PostOrganizationsDocumentsSchedules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrganizationsDocumentsSchedulesRequest struct via the builder pattern


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


## PostOrganizationsDocumentsSchedulesDelete

> PostOrganizationsDocumentsSchedulesDelete(ctx, orgSlug, scheduleId).Execute()

Delete an export schedule; past runs keep their history

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
	scheduleId := "scheduleId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.PostOrganizationsDocumentsSchedulesDelete(context.Background(), orgSlug, scheduleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.PostOrganizationsDocumentsSchedulesDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 
**scheduleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrganizationsDocumentsSchedulesDeleteRequest struct via the builder pattern


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


## PostOrganizationsDocumentsSchedulesPause

> PostOrganizationsDocumentsSchedulesPause(ctx, orgSlug, scheduleId).Execute()

Pause an export schedule

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
	scheduleId := "scheduleId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.PostOrganizationsDocumentsSchedulesPause(context.Background(), orgSlug, scheduleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.PostOrganizationsDocumentsSchedulesPause``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 
**scheduleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrganizationsDocumentsSchedulesPauseRequest struct via the builder pattern


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


## PostOrganizationsDocumentsSchedulesResume

> PostOrganizationsDocumentsSchedulesResume(ctx, orgSlug, scheduleId).Execute()

Resume a paused export schedule

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
	scheduleId := "scheduleId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.PostOrganizationsDocumentsSchedulesResume(context.Background(), orgSlug, scheduleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.PostOrganizationsDocumentsSchedulesResume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 
**scheduleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrganizationsDocumentsSchedulesResumeRequest struct via the builder pattern


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


## PostOrganizationsDocumentsSchedulesRun

> PostOrganizationsDocumentsSchedulesRun(ctx, orgSlug, scheduleId).Execute()

Run an export schedule immediately; also restarts its cadence interval

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
	scheduleId := "scheduleId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.PostOrganizationsDocumentsSchedulesRun(context.Background(), orgSlug, scheduleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.PostOrganizationsDocumentsSchedulesRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 
**scheduleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrganizationsDocumentsSchedulesRunRequest struct via the builder pattern


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


## PostOrganizationsEngagements

> EngagementCreated PostOrganizationsEngagements(ctx, orgSlug).Title(title).Visibility(visibility).Summary(summary).Execute()

Create a draft bug-bounty engagement; admins only, the slug derives from the title

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
	title := "title_example" // string | Engagement title; the URL slug derives from it
	visibility := "visibility_example" // string | private engagements are invisible to non-members; public ones appear in the directory once live
	summary := "summary_example" // string | Short summary shown to researchers; editable later (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.PostOrganizationsEngagements(context.Background(), orgSlug).Title(title).Visibility(visibility).Summary(summary).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.PostOrganizationsEngagements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostOrganizationsEngagements`: EngagementCreated
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.PostOrganizationsEngagements`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrganizationsEngagementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **title** | **string** | Engagement title; the URL slug derives from it | 
 **visibility** | **string** | private engagements are invisible to non-members; public ones appear in the directory once live | 
 **summary** | **string** | Short summary shown to researchers; editable later | 

### Return type

[**EngagementCreated**](EngagementCreated.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOrganizationsEngagementsClose

> PostOrganizationsEngagementsClose(ctx, orgSlug, engagementSlug).Execute()

Close the engagement permanently: submissions stop, history stays readable; 409 unless active or paused

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
	r, err := apiClient.OrganizationsAPI.PostOrganizationsEngagementsClose(context.Background(), orgSlug, engagementSlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.PostOrganizationsEngagementsClose``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostOrganizationsEngagementsCloseRequest struct via the builder pattern


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


## PostOrganizationsEngagementsContent

> PostOrganizationsEngagementsContent(ctx, orgSlug, engagementSlug).NonQualifiedVulns(nonQualifiedVulns).Procedure(procedure).QualifiedVulns(qualifiedVulns).SafeHarbor(safeHarbor).Summary(summary).TestingAllowed(testingAllowed).TestingDisallowed(testingDisallowed).Execute()

Save the engagement's draft content (summary, testing rules, qualifying vulnerabilities, procedure, safe harbor); opens a draft version when none is open; 409 while the open version is in review

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
	nonQualifiedVulns := "nonQualifiedVulns_example" // string | Non-qualifying vulnerabilities, one per line (optional)
	procedure := "procedure_example" // string | How-it-works steps, one per line (optional)
	qualifiedVulns := "qualifiedVulns_example" // string | Qualifying vulnerabilities, one per line (optional)
	safeHarbor := "safeHarbor_example" // string | Legal safe-harbor statement (optional)
	summary := "summary_example" // string | Program description blurb (optional)
	testingAllowed := "testingAllowed_example" // string | Allowed testing rules, one per line (optional)
	testingDisallowed := "testingDisallowed_example" // string | Disallowed testing rules, one per line (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.PostOrganizationsEngagementsContent(context.Background(), orgSlug, engagementSlug).NonQualifiedVulns(nonQualifiedVulns).Procedure(procedure).QualifiedVulns(qualifiedVulns).SafeHarbor(safeHarbor).Summary(summary).TestingAllowed(testingAllowed).TestingDisallowed(testingDisallowed).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.PostOrganizationsEngagementsContent``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostOrganizationsEngagementsContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **nonQualifiedVulns** | **string** | Non-qualifying vulnerabilities, one per line | 
 **procedure** | **string** | How-it-works steps, one per line | 
 **qualifiedVulns** | **string** | Qualifying vulnerabilities, one per line | 
 **safeHarbor** | **string** | Legal safe-harbor statement | 
 **summary** | **string** | Program description blurb | 
 **testingAllowed** | **string** | Allowed testing rules, one per line | 
 **testingDisallowed** | **string** | Disallowed testing rules, one per line | 

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


## PostOrganizationsEngagementsMessageProxy

> PostOrganizationsEngagementsMessageProxy(ctx, orgSlug, engagementSlug).FromNumber(fromNumber).Execute()

Configure the SMS message proxy: set the gateway sender number, generating the webhook secret on first enable

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
	fromNumber := "fromNumber_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.PostOrganizationsEngagementsMessageProxy(context.Background(), orgSlug, engagementSlug).FromNumber(fromNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.PostOrganizationsEngagementsMessageProxy``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostOrganizationsEngagementsMessageProxyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fromNumber** | **string** |  | 

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


## PostOrganizationsEngagementsMessageProxyDisable

> PostOrganizationsEngagementsMessageProxyDisable(ctx, orgSlug, engagementSlug).Execute()

Disable the SMS message proxy and discard the webhook secret

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
	r, err := apiClient.OrganizationsAPI.PostOrganizationsEngagementsMessageProxyDisable(context.Background(), orgSlug, engagementSlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.PostOrganizationsEngagementsMessageProxyDisable``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostOrganizationsEngagementsMessageProxyDisableRequest struct via the builder pattern


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


## PostOrganizationsEngagementsPause

> PostOrganizationsEngagementsPause(ctx, orgSlug, engagementSlug).Execute()

Pause a live engagement: new submissions stop; 409 unless active

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
	r, err := apiClient.OrganizationsAPI.PostOrganizationsEngagementsPause(context.Background(), orgSlug, engagementSlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.PostOrganizationsEngagementsPause``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostOrganizationsEngagementsPauseRequest struct via the builder pattern


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


## PostOrganizationsEngagementsReopen

> PostOrganizationsEngagementsReopen(ctx, orgSlug, engagementSlug).Execute()

Reopen a changes-requested version as an editable draft; 409 unless changes are requested

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
	r, err := apiClient.OrganizationsAPI.PostOrganizationsEngagementsReopen(context.Background(), orgSlug, engagementSlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.PostOrganizationsEngagementsReopen``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostOrganizationsEngagementsReopenRequest struct via the builder pattern


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


## PostOrganizationsEngagementsReportsAccept

> ReportOrgReportDetail PostOrganizationsEngagementsReportsAccept(ctx, orgSlug, engagementSlug, reportId).Execute()

Accept a validated report; 409 for any other workflow state

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
	reportId := "reportId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.PostOrganizationsEngagementsReportsAccept(context.Background(), orgSlug, engagementSlug, reportId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.PostOrganizationsEngagementsReportsAccept``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostOrganizationsEngagementsReportsAccept`: ReportOrgReportDetail
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.PostOrganizationsEngagementsReportsAccept`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 
**engagementSlug** | **string** |  | 
**reportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrganizationsEngagementsReportsAcceptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ReportOrgReportDetail**](ReportOrgReportDetail.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOrganizationsEngagementsReportsAcknowledgeEscalation

> PostOrganizationsEngagementsReportsAcknowledgeEscalation(ctx, orgSlug, engagementSlug, reportId).Execute()

Acknowledge an escalated report; 409 when it is not escalated or already acknowledged

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
	reportId := "reportId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.PostOrganizationsEngagementsReportsAcknowledgeEscalation(context.Background(), orgSlug, engagementSlug, reportId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.PostOrganizationsEngagementsReportsAcknowledgeEscalation``: %v\n", err)
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
**reportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrganizationsEngagementsReportsAcknowledgeEscalationRequest struct via the builder pattern


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


## PostOrganizationsEngagementsReportsAssign

> ReportOrgReportDetail PostOrganizationsEngagementsReportsAssign(ctx, orgSlug, engagementSlug, reportId).Assignee(assignee).Execute()

Set or clear the organization's internal owner of the report by user id, one of the detail's assignmentOptions; carries no permission; 400 names the assignee field when the id is not an active member

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
	reportId := "reportId_example" // string | 
	assignee := "assignee_example" // string | Numeric user id of an active member, one of the report detail's assignmentOptions (also the members list userId, never its id). Empty to unassign (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.PostOrganizationsEngagementsReportsAssign(context.Background(), orgSlug, engagementSlug, reportId).Assignee(assignee).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.PostOrganizationsEngagementsReportsAssign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostOrganizationsEngagementsReportsAssign`: ReportOrgReportDetail
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.PostOrganizationsEngagementsReportsAssign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 
**engagementSlug** | **string** |  | 
**reportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrganizationsEngagementsReportsAssignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **assignee** | **string** | Numeric user id of an active member, one of the report detail&#39;s assignmentOptions (also the members list userId, never its id). Empty to unassign | 

### Return type

[**ReportOrgReportDetail**](ReportOrgReportDetail.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOrganizationsEngagementsReportsComments

> ReportPostedComment PostOrganizationsEngagementsReportsComments(ctx, orgSlug, engagementSlug, reportId).Body(body).Audience(audience).Execute()

Post a message on the report's thread: audience 'all' (researcher-visible question) or 'organization' (internal note)

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
	reportId := "reportId_example" // string | 
	body := "body_example" // string | The message text
	audience := "audience_example" // string | Message audience; the server enforces the role's allowlist (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.PostOrganizationsEngagementsReportsComments(context.Background(), orgSlug, engagementSlug, reportId).Body(body).Audience(audience).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.PostOrganizationsEngagementsReportsComments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostOrganizationsEngagementsReportsComments`: ReportPostedComment
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.PostOrganizationsEngagementsReportsComments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 
**engagementSlug** | **string** |  | 
**reportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrganizationsEngagementsReportsCommentsRequest struct via the builder pattern


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


## PostOrganizationsEngagementsReportsReject

> ReportOrgReportDetail PostOrganizationsEngagementsReportsReject(ctx, orgSlug, engagementSlug, reportId).Reason(reason).Execute()

Reject a validated report with a mandatory reason; 409 for any other workflow state

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
	reportId := "reportId_example" // string | 
	reason := "reason_example" // string | Why the organization rejects the validated report; shared with the researcher

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.PostOrganizationsEngagementsReportsReject(context.Background(), orgSlug, engagementSlug, reportId).Reason(reason).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.PostOrganizationsEngagementsReportsReject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostOrganizationsEngagementsReportsReject`: ReportOrgReportDetail
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.PostOrganizationsEngagementsReportsReject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 
**engagementSlug** | **string** |  | 
**reportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrganizationsEngagementsReportsRejectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **reason** | **string** | Why the organization rejects the validated report; shared with the researcher | 

### Return type

[**ReportOrgReportDetail**](ReportOrgReportDetail.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOrganizationsEngagementsReportsRetests

> Retest PostOrganizationsEngagementsReportsRetests(ctx, orgSlug, engagementSlug, reportId).Notes(notes).Execute()

Mark the accepted or reopened report fixed and request a retest with remediation notes; 409 when one is already open

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
	reportId := "reportId_example" // string | 
	notes := "notes_example" // string | What was fixed and how; shown to the verifying researcher

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.PostOrganizationsEngagementsReportsRetests(context.Background(), orgSlug, engagementSlug, reportId).Notes(notes).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.PostOrganizationsEngagementsReportsRetests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostOrganizationsEngagementsReportsRetests`: Retest
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.PostOrganizationsEngagementsReportsRetests`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 
**engagementSlug** | **string** |  | 
**reportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrganizationsEngagementsReportsRetestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **notes** | **string** | What was fixed and how; shown to the verifying researcher | 

### Return type

[**Retest**](Retest.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOrganizationsEngagementsReportsRetestsCancel

> Retest PostOrganizationsEngagementsReportsRetestsCancel(ctx, orgSlug, engagementSlug, reportId, retestId).Execute()

Cancel an open retest request; 409 once the verifier started, the report stays fixed

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
	reportId := "reportId_example" // string | 
	retestId := "retestId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.PostOrganizationsEngagementsReportsRetestsCancel(context.Background(), orgSlug, engagementSlug, reportId, retestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.PostOrganizationsEngagementsReportsRetestsCancel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostOrganizationsEngagementsReportsRetestsCancel`: Retest
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.PostOrganizationsEngagementsReportsRetestsCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 
**engagementSlug** | **string** |  | 
**reportId** | **string** |  | 
**retestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrganizationsEngagementsReportsRetestsCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**Retest**](Retest.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOrganizationsEngagementsResume

> PostOrganizationsEngagementsResume(ctx, orgSlug, engagementSlug).Execute()

Resume a paused engagement; 409 unless paused

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
	r, err := apiClient.OrganizationsAPI.PostOrganizationsEngagementsResume(context.Background(), orgSlug, engagementSlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.PostOrganizationsEngagementsResume``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostOrganizationsEngagementsResumeRequest struct via the builder pattern


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


## PostOrganizationsEngagementsSubmit

> PostOrganizationsEngagementsSubmit(ctx, orgSlug, engagementSlug).Note(note).Execute()

Send the open draft version to backoffice review with an optional note; 400 until the draft has an in-scope asset; 409 without an open version or when it is not a draft

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
	note := "note_example" // string | Note stored on the version transition (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.PostOrganizationsEngagementsSubmit(context.Background(), orgSlug, engagementSlug).Note(note).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.PostOrganizationsEngagementsSubmit``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostOrganizationsEngagementsSubmitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **note** | **string** | Note stored on the version transition | 

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


## PostOrganizationsEngagementsVerificationLevel

> PostOrganizationsEngagementsVerificationLevel(ctx, orgSlug, engagementSlug).Level(level).Execute()

Set the researcher verification level the engagement requires; immediate and audited

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
	level := "level_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.PostOrganizationsEngagementsVerificationLevel(context.Background(), orgSlug, engagementSlug).Level(level).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.PostOrganizationsEngagementsVerificationLevel``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostOrganizationsEngagementsVerificationLevelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **level** | **string** |  | 

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


## PostOrganizationsEngagementsVisibility

> PostOrganizationsEngagementsVisibility(ctx, orgSlug, engagementSlug).Visibility(visibility).Execute()

Switch the engagement between private and public; immediate and audited

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
	visibility := "visibility_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.PostOrganizationsEngagementsVisibility(context.Background(), orgSlug, engagementSlug).Visibility(visibility).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.PostOrganizationsEngagementsVisibility``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostOrganizationsEngagementsVisibilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **visibility** | **string** |  | 

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


## PostOrganizationsIntegrationsWebhooks

> PostOrganizationsIntegrationsWebhooks(ctx, orgSlug).Url(url).Active(active).Events(events).Execute()

Register a signed webhook endpoint subscribed to event types

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
	url := "url_example" // string | HTTPS endpoint that receives signed deliveries
	active := "active_example" // string | Set to keep the webhook active (optional)
	events := []string{"Inner_example"} // []string | Event names to subscribe to (use an asterisk for all) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.PostOrganizationsIntegrationsWebhooks(context.Background(), orgSlug).Url(url).Active(active).Events(events).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.PostOrganizationsIntegrationsWebhooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrganizationsIntegrationsWebhooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **url** | **string** | HTTPS endpoint that receives signed deliveries | 
 **active** | **string** | Set to keep the webhook active | 
 **events** | **[]string** | Event names to subscribe to (use an asterisk for all) | 

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


## PostOrganizationsIntegrationsWebhooksByWebhookId

> PostOrganizationsIntegrationsWebhooksByWebhookId(ctx, orgSlug, webhookId).Url(url).Active(active).Events(events).Execute()

Update a webhook's URL, events, and active flag

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
	webhookId := "webhookId_example" // string | 
	url := "url_example" // string | HTTPS endpoint that receives signed deliveries
	active := "active_example" // string | Set to keep the webhook active (optional)
	events := []string{"Inner_example"} // []string | Event names to subscribe to (use an asterisk for all) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.PostOrganizationsIntegrationsWebhooksByWebhookId(context.Background(), orgSlug, webhookId).Url(url).Active(active).Events(events).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.PostOrganizationsIntegrationsWebhooksByWebhookId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 
**webhookId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrganizationsIntegrationsWebhooksByWebhookIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **url** | **string** | HTTPS endpoint that receives signed deliveries | 
 **active** | **string** | Set to keep the webhook active | 
 **events** | **[]string** | Event names to subscribe to (use an asterisk for all) | 

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


## PostOrganizationsIntegrationsWebhooksDelete

> PostOrganizationsIntegrationsWebhooksDelete(ctx, orgSlug, webhookId).Execute()

Delete a webhook

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
	webhookId := "webhookId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.PostOrganizationsIntegrationsWebhooksDelete(context.Background(), orgSlug, webhookId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.PostOrganizationsIntegrationsWebhooksDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 
**webhookId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrganizationsIntegrationsWebhooksDeleteRequest struct via the builder pattern


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


## PostOrganizationsIntegrationsWebhooksTest

> PostOrganizationsIntegrationsWebhooksTest(ctx, orgSlug, webhookId).Execute()

Send a signed test (ping) delivery to the webhook

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
	webhookId := "webhookId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.PostOrganizationsIntegrationsWebhooksTest(context.Background(), orgSlug, webhookId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.PostOrganizationsIntegrationsWebhooksTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 
**webhookId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrganizationsIntegrationsWebhooksTestRequest struct via the builder pattern


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


## PostOrganizationsMembersByMemberId

> OrganizationMember PostOrganizationsMembersByMemberId(ctx, orgSlug, memberId).Execute()

Update a member's category flags and program grants, admin-only; refuses demoting the last admin

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
	memberId := "memberId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.PostOrganizationsMembersByMemberId(context.Background(), orgSlug, memberId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.PostOrganizationsMembersByMemberId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostOrganizationsMembersByMemberId`: OrganizationMember
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.PostOrganizationsMembersByMemberId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 
**memberId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrganizationsMembersByMemberIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OrganizationMember**](OrganizationMember.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOrganizationsMembersInvite

> OrganizationMember PostOrganizationsMembersInvite(ctx, orgSlug).Execute()

Invite a member by email with category flags and program grants, admin-only

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.PostOrganizationsMembersInvite(context.Background(), orgSlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.PostOrganizationsMembersInvite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostOrganizationsMembersInvite`: OrganizationMember
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.PostOrganizationsMembersInvite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrganizationsMembersInviteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrganizationMember**](OrganizationMember.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOrganizationsScopesAssets

> ScopeAsset PostOrganizationsScopesAssets(ctx, orgSlug).Availability(availability).Confidentiality(confidentiality).Criticality(criticality).Environment(environment).Identifier(identifier).Integrity(integrity).Kind(kind).Description(description).TierId(tierId).Execute()

Add an asset to the organization's inventory

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
	availability := "availability_example" // string | CVSS availability requirement
	confidentiality := "confidentiality_example" // string | CVSS confidentiality requirement
	criticality := "criticality_example" // string | Business criticality
	environment := "environment_example" // string | Deployment environment
	identifier := "identifier_example" // string | Asset identifier, e.g. api.example.com
	integrity := "integrity_example" // string | CVSS integrity requirement
	kind := "kind_example" // string | Target class (E012)
	description := "description_example" // string | What the asset is and why it matters (optional)
	tierId := "tierId_example" // string | Asset tier id, empty for none (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.PostOrganizationsScopesAssets(context.Background(), orgSlug).Availability(availability).Confidentiality(confidentiality).Criticality(criticality).Environment(environment).Identifier(identifier).Integrity(integrity).Kind(kind).Description(description).TierId(tierId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.PostOrganizationsScopesAssets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostOrganizationsScopesAssets`: ScopeAsset
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.PostOrganizationsScopesAssets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrganizationsScopesAssetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **availability** | **string** | CVSS availability requirement | 
 **confidentiality** | **string** | CVSS confidentiality requirement | 
 **criticality** | **string** | Business criticality | 
 **environment** | **string** | Deployment environment | 
 **identifier** | **string** | Asset identifier, e.g. api.example.com | 
 **integrity** | **string** | CVSS integrity requirement | 
 **kind** | **string** | Target class (E012) | 
 **description** | **string** | What the asset is and why it matters | 
 **tierId** | **string** | Asset tier id, empty for none | 

### Return type

[**ScopeAsset**](ScopeAsset.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOrganizationsScopesAssetsByAssetId

> ScopeAsset PostOrganizationsScopesAssetsByAssetId(ctx, orgSlug, assetId).Availability(availability).Confidentiality(confidentiality).Criticality(criticality).Environment(environment).Identifier(identifier).Integrity(integrity).Kind(kind).Description(description).TierId(tierId).Execute()

Update an inventory asset; the identifier must stay unique within the organization

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
	assetId := "assetId_example" // string | 
	availability := "availability_example" // string | CVSS availability requirement
	confidentiality := "confidentiality_example" // string | CVSS confidentiality requirement
	criticality := "criticality_example" // string | Business criticality
	environment := "environment_example" // string | Deployment environment
	identifier := "identifier_example" // string | Asset identifier, e.g. api.example.com
	integrity := "integrity_example" // string | CVSS integrity requirement
	kind := "kind_example" // string | Target class (E012)
	description := "description_example" // string | What the asset is and why it matters (optional)
	tierId := "tierId_example" // string | Asset tier id, empty for none (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.PostOrganizationsScopesAssetsByAssetId(context.Background(), orgSlug, assetId).Availability(availability).Confidentiality(confidentiality).Criticality(criticality).Environment(environment).Identifier(identifier).Integrity(integrity).Kind(kind).Description(description).TierId(tierId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.PostOrganizationsScopesAssetsByAssetId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostOrganizationsScopesAssetsByAssetId`: ScopeAsset
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.PostOrganizationsScopesAssetsByAssetId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 
**assetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrganizationsScopesAssetsByAssetIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **availability** | **string** | CVSS availability requirement | 
 **confidentiality** | **string** | CVSS confidentiality requirement | 
 **criticality** | **string** | Business criticality | 
 **environment** | **string** | Deployment environment | 
 **identifier** | **string** | Asset identifier, e.g. api.example.com | 
 **integrity** | **string** | CVSS integrity requirement | 
 **kind** | **string** | Target class (E012) | 
 **description** | **string** | What the asset is and why it matters | 
 **tierId** | **string** | Asset tier id, empty for none | 

### Return type

[**ScopeAsset**](ScopeAsset.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOrganizationsScopesEngagementsDraft

> ScopeDraft PostOrganizationsScopesEngagementsDraft(ctx, orgSlug, engagementSlug).Execute()

Open a draft engagement version so the live scope becomes editable

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
	resp, r, err := apiClient.OrganizationsAPI.PostOrganizationsScopesEngagementsDraft(context.Background(), orgSlug, engagementSlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.PostOrganizationsScopesEngagementsDraft``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostOrganizationsScopesEngagementsDraft`: ScopeDraft
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.PostOrganizationsScopesEngagementsDraft`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 
**engagementSlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrganizationsScopesEngagementsDraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ScopeDraft**](ScopeDraft.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOrganizationsScopesEngagementsEntries

> ScopeEntry PostOrganizationsScopesEngagementsEntries(ctx, orgSlug, engagementSlug).Scope(scope).AssetId(assetId).Instructions(instructions).TierId(tierId).Execute()

Add a scope entry to the engagement's draft version; adding the first entry creates the draft

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
	scope := "scope_example" // string | Whether the asset is in or out of scope
	assetId := "assetId_example" // string | Inventory asset id to select, required when adding (optional)
	instructions := "instructions_example" // string | Rules of engagement for this selection (optional)
	tierId := "tierId_example" // string | Asset tier override, empty for the asset's own tier (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.PostOrganizationsScopesEngagementsEntries(context.Background(), orgSlug, engagementSlug).Scope(scope).AssetId(assetId).Instructions(instructions).TierId(tierId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.PostOrganizationsScopesEngagementsEntries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostOrganizationsScopesEngagementsEntries`: ScopeEntry
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.PostOrganizationsScopesEngagementsEntries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 
**engagementSlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrganizationsScopesEngagementsEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **scope** | **string** | Whether the asset is in or out of scope | 
 **assetId** | **string** | Inventory asset id to select, required when adding | 
 **instructions** | **string** | Rules of engagement for this selection | 
 **tierId** | **string** | Asset tier override, empty for the asset&#39;s own tier | 

### Return type

[**ScopeEntry**](ScopeEntry.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOrganizationsScopesEngagementsEntriesByEntryId

> ScopeEntry PostOrganizationsScopesEngagementsEntriesByEntryId(ctx, orgSlug, engagementSlug, entryId).Scope(scope).AssetId(assetId).Instructions(instructions).TierId(tierId).Execute()

Update a scope entry on the engagement's editable draft version; 409 when no editable draft is open

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
	entryId := "entryId_example" // string | 
	scope := "scope_example" // string | Whether the asset is in or out of scope
	assetId := "assetId_example" // string | Inventory asset id to select, required when adding (optional)
	instructions := "instructions_example" // string | Rules of engagement for this selection (optional)
	tierId := "tierId_example" // string | Asset tier override, empty for the asset's own tier (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.PostOrganizationsScopesEngagementsEntriesByEntryId(context.Background(), orgSlug, engagementSlug, entryId).Scope(scope).AssetId(assetId).Instructions(instructions).TierId(tierId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.PostOrganizationsScopesEngagementsEntriesByEntryId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostOrganizationsScopesEngagementsEntriesByEntryId`: ScopeEntry
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.PostOrganizationsScopesEngagementsEntriesByEntryId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 
**engagementSlug** | **string** |  | 
**entryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrganizationsScopesEngagementsEntriesByEntryIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **scope** | **string** | Whether the asset is in or out of scope | 
 **assetId** | **string** | Inventory asset id to select, required when adding | 
 **instructions** | **string** | Rules of engagement for this selection | 
 **tierId** | **string** | Asset tier override, empty for the asset&#39;s own tier | 

### Return type

[**ScopeEntry**](ScopeEntry.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOrganizationsScopesEngagementsEntriesDelete

> PostOrganizationsScopesEngagementsEntriesDelete(ctx, orgSlug, engagementSlug, entryId).Execute()

Delete a scope entry from the engagement's editable draft version

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
	entryId := "entryId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.PostOrganizationsScopesEngagementsEntriesDelete(context.Background(), orgSlug, engagementSlug, entryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.PostOrganizationsScopesEngagementsEntriesDelete``: %v\n", err)
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
**entryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrganizationsScopesEngagementsEntriesDeleteRequest struct via the builder pattern


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


## PostOrganizationsSettings

> OrganizationProfileVersion PostOrganizationsSettings(ctx, orgSlug).Execute()

Save the organization profile draft (name, website, industry, brand colors), admin-only; 409 while the draft is frozen in review

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.PostOrganizationsSettings(context.Background(), orgSlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.PostOrganizationsSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostOrganizationsSettings`: OrganizationProfileVersion
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.PostOrganizationsSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrganizationsSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrganizationProfileVersion**](OrganizationProfileVersion.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOrganizationsSettingsLogoDelete

> PostOrganizationsSettingsLogoDelete(ctx, orgSlug).Execute()

Remove the logo from the profile draft; the pages fall back to the monogram avatar

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.PostOrganizationsSettingsLogoDelete(context.Background(), orgSlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.PostOrganizationsSettingsLogoDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrganizationsSettingsLogoDeleteRequest struct via the builder pattern


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


## PostOrganizationsSettingsReopen

> PostOrganizationsSettingsReopen(ctx, orgSlug).Execute()

Reopen a changes-requested profile version as a draft; 409 unless changes are requested

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.PostOrganizationsSettingsReopen(context.Background(), orgSlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.PostOrganizationsSettingsReopen``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrganizationsSettingsReopenRequest struct via the builder pattern


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


## PostOrganizationsSettingsSubmit

> PostOrganizationsSettingsSubmit(ctx, orgSlug).Execute()

Submit the profile draft for backoffice review; requires a non-empty name, 409 without an open draft

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.PostOrganizationsSettingsSubmit(context.Background(), orgSlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.PostOrganizationsSettingsSubmit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrganizationsSettingsSubmitRequest struct via the builder pattern


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


## PostOrganizationsVerification

> PostOrganizationsVerification(ctx, orgSlug).Country(country).LegalName(legalName).RegistryNumber(registryNumber).Execute()

Answer an information request with updated legal details; returns the registration to the review queue

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
	country := "country_example" // string | Country of incorporation
	legalName := "legalName_example" // string | Registered legal entity name
	registryNumber := "registryNumber_example" // string | Commercial-register number (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.PostOrganizationsVerification(context.Background(), orgSlug).Country(country).LegalName(legalName).RegistryNumber(registryNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.PostOrganizationsVerification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrganizationsVerificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **country** | **string** | Country of incorporation | 
 **legalName** | **string** | Registered legal entity name | 
 **registryNumber** | **string** | Commercial-register number | 

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


## PostOrganizationsWalletRewardsApprove

> PostOrganizationsWalletRewardsApprove(ctx, orgSlug, id).Execute()

Approve a proposed bounty: debit the wallet and mark it approved; finance role only

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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.PostOrganizationsWalletRewardsApprove(context.Background(), orgSlug, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.PostOrganizationsWalletRewardsApprove``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrganizationsWalletRewardsApproveRequest struct via the builder pattern


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


## PostOrganizationsWalletRewardsReject

> PostOrganizationsWalletRewardsReject(ctx, orgSlug, id).Reason(reason).Execute()

Reject a proposed bounty with a reason; finance role only

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
	id := "id_example" // string | 
	reason := "reason_example" // string | Why the bounty is rejected

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.PostOrganizationsWalletRewardsReject(context.Background(), orgSlug, id).Reason(reason).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.PostOrganizationsWalletRewardsReject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrganizationsWalletRewardsRejectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **reason** | **string** | Why the bounty is rejected | 

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


## PostOrganizationsWalletTopups

> PostOrganizationsWalletTopups(ctx, orgSlug).Amount(amount).Execute()

Request a bank-transfer wallet top-up; returns the payment reference to quote

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
	amount := "amount_example" // string | Funding amount in whole CHF

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.PostOrganizationsWalletTopups(context.Background(), orgSlug).Amount(amount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.PostOrganizationsWalletTopups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrganizationsWalletTopupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **amount** | **string** | Funding amount in whole CHF | 

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

