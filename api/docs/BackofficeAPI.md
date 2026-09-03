# \BackofficeAPI

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBackofficeBadges**](BackofficeAPI.md#GetBackofficeBadges) | **Get** /backoffice/badges | Staff badge catalog: every badge including inactive ones
[**GetBackofficeBadgesByCode**](BackofficeAPI.md#GetBackofficeBadgesByCode) | **Get** /backoffice/badges/{code} | One badge with its holders
[**GetBackofficeBudgets**](BackofficeAPI.md#GetBackofficeBudgets) | **Get** /backoffice/budgets | The pending-report budget configuration: per-tier limits, slop-rejection penalties, and active per-researcher overrides
[**GetBackofficeCertifications**](BackofficeAPI.md#GetBackofficeCertifications) | **Get** /backoffice/certifications | Staff review queue of claimed researcher certifications
[**GetBackofficeCompliance**](BackofficeAPI.md#GetBackofficeCompliance) | **Get** /backoffice/compliance | Researcher payout compliance (sanctions) for the backoffice stopgap
[**GetBackofficeConduct**](BackofficeAPI.md#GetBackofficeConduct) | **Get** /backoffice/conduct | The trust &amp; safety conduct case queue
[**GetBackofficeConductByCaseId**](BackofficeAPI.md#GetBackofficeConductByCaseId) | **Get** /backoffice/conduct/{case_id} | One conduct case with its enforcement actions and appeals
[**GetBackofficeCustomerOnboarding**](BackofficeAPI.md#GetBackofficeCustomerOnboarding) | **Get** /backoffice/customer-onboarding | The back-office queue of customer onboarding requests, filterable by status and customer email
[**GetBackofficeFeedback**](BackofficeAPI.md#GetBackofficeFeedback) | **Get** /backoffice/feedback | The product feedback queue, optionally filtered by status
[**GetBackofficeOrganizations**](BackofficeAPI.md#GetBackofficeOrganizations) | **Get** /backoffice/organizations | Every customer and partner organization with verification status and engagement counts
[**GetBackofficePayouts**](BackofficeAPI.md#GetBackofficePayouts) | **Get** /backoffice/payouts | Payout run queue: eligible and blocked researchers, plus payouts by state
[**GetBackofficeProfileReviews**](BackofficeAPI.md#GetBackofficeProfileReviews) | **Get** /backoffice/profile-reviews | Organization profile versions awaiting review with their field diffs, plus the recently decided ones
[**GetBackofficeRegistrations**](BackofficeAPI.md#GetBackofficeRegistrations) | **Get** /backoffice/registrations | Registration review queue with domain-verification evidence
[**GetBackofficeSla**](BackofficeAPI.md#GetBackofficeSla) | **Get** /backoffice/sla | The active SLA policy matrix (first-response and resolution hours per engagement type and severity)
[**GetBackofficeTodos**](BackofficeAPI.md#GetBackofficeTodos) | **Get** /backoffice/todos | Staff triage inbox: reports awaiting or in triage, scoped to the caller&#39;s backoffice groups; filterable by status, engagement, severity, assignment, and search, sortable, paginated
[**GetBackofficeTodosByReportId**](BackofficeAPI.md#GetBackofficeTodosByReportId) | **Get** /backoffice/todos/{report_id} | Triage workspace: the full report a triager assesses
[**GetBackofficeTodosDuplicateCandidates**](BackofficeAPI.md#GetBackofficeTodosDuplicateCandidates) | **Get** /backoffice/todos/{report_id}/duplicate-candidates | List candidate original reports in the engagement to mark a duplicate against (staff picker)
[**GetBackofficeTodosExport**](BackofficeAPI.md#GetBackofficeTodosExport) | **Get** /backoffice/todos/export | Download the filtered triage queue as CSV or JSON (every page, at most 10000 rows)
[**GetBackofficeTopups**](BackofficeAPI.md#GetBackofficeTopups) | **Get** /backoffice/topups | Pending wallet top-ups awaiting reconciliation across organizations
[**GetBackofficeVerifications**](BackofficeAPI.md#GetBackofficeVerifications) | **Get** /backoffice/verifications | Pending background checks and researcher verification levels
[**PostBackofficeBadges**](BackofficeAPI.md#PostBackofficeBadges) | **Post** /backoffice/badges | Create a badge; the code is stable, names and descriptions are per-locale with English required, and the criteria kind with threshold and parameters defines the earning logic
[**PostBackofficeBadgesAward**](BackofficeAPI.md#PostBackofficeBadgesAward) | **Post** /backoffice/badges/{code}/award | Manually award a badge to a researcher by handle or email; idempotent, and the researcher is notified only on a new award
[**PostBackofficeBadgesAwardsRevoke**](BackofficeAPI.md#PostBackofficeBadgesAwardsRevoke) | **Post** /backoffice/badges/{code}/awards/{award_id}/revoke | Revoke one badge award; audited
[**PostBackofficeBadgesByCode**](BackofficeAPI.md#PostBackofficeBadgesByCode) | **Post** /backoffice/badges/{code} | Update a badge&#39;s texts, visuals, and criteria; the code cannot change
[**PostBackofficeBadgesDeactivate**](BackofficeAPI.md#PostBackofficeBadgesDeactivate) | **Post** /backoffice/badges/{code}/deactivate | Deactivate a badge: it disappears from the researcher catalog and profiles and stops being awarded; existing awards are kept
[**PostBackofficeBadgesReactivate**](BackofficeAPI.md#PostBackofficeBadgesReactivate) | **Post** /backoffice/badges/{code}/reactivate | Reactivate a previously deactivated badge
[**PostBackofficeBudgetsOverrides**](BackofficeAPI.md#PostBackofficeBudgetsOverrides) | **Post** /backoffice/budgets/overrides | Create a per-researcher budget override with a mandatory reason and optional expiry
[**PostBackofficeBudgetsOverridesDelete**](BackofficeAPI.md#PostBackofficeBudgetsOverridesDelete) | **Post** /backoffice/budgets/overrides/{id}/delete | Remove a budget override; the tier limit applies again
[**PostBackofficeBudgetsPenalties**](BackofficeAPI.md#PostBackofficeBudgetsPenalties) | **Post** /backoffice/budgets/penalties | Update the spam/ai_slop rejection penalty points (always negative)
[**PostBackofficeBudgetsTiers**](BackofficeAPI.md#PostBackofficeBudgetsTiers) | **Post** /backoffice/budgets/tiers | Update the per-tier pending-report limits; empty means unlimited, limits must not decrease on higher tiers
[**PostBackofficeCertificationsReject**](BackofficeAPI.md#PostBackofficeCertificationsReject) | **Post** /backoffice/certifications/{id}/reject | Reject a claimed certification (staff decision, audit-logged)
[**PostBackofficeCertificationsVerify**](BackofficeAPI.md#PostBackofficeCertificationsVerify) | **Post** /backoffice/certifications/{id}/verify | Verify a claimed certification (staff decision, audit-logged)
[**PostBackofficeComplianceByUserId**](BackofficeAPI.md#PostBackofficeComplianceByUserId) | **Post** /backoffice/compliance/{user_id} | Set a researcher&#39;s sanctions status
[**PostBackofficeConduct**](BackofficeAPI.md#PostBackofficeConduct) | **Post** /backoffice/conduct | Open a conduct case against a researcher
[**PostBackofficeConductAppealResolve**](BackofficeAPI.md#PostBackofficeConductAppealResolve) | **Post** /backoffice/conduct/{case_id}/appeal/{appeal_id}/resolve | Uphold or overturn a conduct appeal
[**PostBackofficeConductClose**](BackofficeAPI.md#PostBackofficeConductClose) | **Post** /backoffice/conduct/{case_id}/close | Close a decided conduct case
[**PostBackofficeConductDecide**](BackofficeAPI.md#PostBackofficeConductDecide) | **Post** /backoffice/conduct/{case_id}/decide | Decide a conduct case
[**PostBackofficeConductDismiss**](BackofficeAPI.md#PostBackofficeConductDismiss) | **Post** /backoffice/conduct/{case_id}/dismiss | Dismiss a conduct case
[**PostBackofficeConductEnforce**](BackofficeAPI.md#PostBackofficeConductEnforce) | **Post** /backoffice/conduct/{case_id}/enforce | Record a graduated enforcement action on a decided case
[**PostBackofficeConductEnforceRevoke**](BackofficeAPI.md#PostBackofficeConductEnforceRevoke) | **Post** /backoffice/conduct/{case_id}/enforce/{action_id}/revoke | Lift an active enforcement action
[**PostBackofficeConductNote**](BackofficeAPI.md#PostBackofficeConductNote) | **Post** /backoffice/conduct/{case_id}/note | Add an investigation note to a case
[**PostBackofficeConductReportFlag**](BackofficeAPI.md#PostBackofficeConductReportFlag) | **Post** /backoffice/conduct/report/{report_id}/flag | Open a conduct case pre-linked to a report and its researcher
[**PostBackofficeConductTakeUp**](BackofficeAPI.md#PostBackofficeConductTakeUp) | **Post** /backoffice/conduct/{case_id}/take-up | Take up a conduct case for investigation
[**PostBackofficeCustomerOnboarding**](BackofficeAPI.md#PostBackofficeCustomerOnboarding) | **Post** /backoffice/customer-onboarding | Create a prefilled customer onboarding request and get its share link for the customer
[**PostBackofficeCustomerOnboardingReject**](BackofficeAPI.md#PostBackofficeCustomerOnboardingReject) | **Post** /backoffice/customer-onboarding/{request_id}/reject | Reject a customer onboarding request with an optional note; 409 once it is closed
[**PostBackofficeFeedbackStatus**](BackofficeAPI.md#PostBackofficeFeedbackStatus) | **Post** /backoffice/feedback/{feedback_id}/status | Move a feedback entry to new, reviewed, or resolved
[**PostBackofficePayoutsApprove**](BackofficeAPI.md#PostBackofficePayoutsApprove) | **Post** /backoffice/payouts/{id}/approve | Approve a pending payout
[**PostBackofficePayoutsFail**](BackofficeAPI.md#PostBackofficePayoutsFail) | **Post** /backoffice/payouts/{id}/fail | Mark a processing payout failed with a reason
[**PostBackofficePayoutsPaid**](BackofficeAPI.md#PostBackofficePayoutsPaid) | **Post** /backoffice/payouts/{id}/paid | Mark a processing payout paid; flips its source rewards to paid
[**PostBackofficePayoutsProcess**](BackofficeAPI.md#PostBackofficePayoutsProcess) | **Post** /backoffice/payouts/{id}/process | Start processing an approved payout
[**PostBackofficePayoutsReject**](BackofficeAPI.md#PostBackofficePayoutsReject) | **Post** /backoffice/payouts/{id}/reject | Cancel a pending payout
[**PostBackofficePayoutsRetry**](BackofficeAPI.md#PostBackofficePayoutsRetry) | **Post** /backoffice/payouts/{id}/retry | Retry a failed payout
[**PostBackofficePayoutsRun**](BackofficeAPI.md#PostBackofficePayoutsRun) | **Post** /backoffice/payouts/run | Run payouts: create one payout per eligible researcher (async, idempotent)
[**PostBackofficeProfileReviewsApprove**](BackofficeAPI.md#PostBackofficeProfileReviewsApprove) | **Post** /backoffice/profile-reviews/{version_id}/approve | Approve the profile version and publish it as the organization&#39;s live profile
[**PostBackofficeProfileReviewsChanges**](BackofficeAPI.md#PostBackofficeProfileReviewsChanges) | **Post** /backoffice/profile-reviews/{version_id}/changes | Send the profile version back to its authors with a required note
[**PostBackofficeProfileReviewsReject**](BackofficeAPI.md#PostBackofficeProfileReviewsReject) | **Post** /backoffice/profile-reviews/{version_id}/reject | Reject the profile version with a required note; rejected is final
[**PostBackofficeRegistrationsApprove**](BackofficeAPI.md#PostBackofficeRegistrationsApprove) | **Post** /backoffice/registrations/{id}/approve | Approve the registration: the organization becomes verified and may run engagements
[**PostBackofficeRegistrationsReject**](BackofficeAPI.md#PostBackofficeRegistrationsReject) | **Post** /backoffice/registrations/{id}/reject | Reject the registration with a mandatory reason
[**PostBackofficeRegistrationsReopen**](BackofficeAPI.md#PostBackofficeRegistrationsReopen) | **Post** /backoffice/registrations/{id}/reopen | Re-open an approved registration; the workspace gate re-engages
[**PostBackofficeRegistrationsRequestInfo**](BackofficeAPI.md#PostBackofficeRegistrationsRequestInfo) | **Post** /backoffice/registrations/{id}/request-info | Send the registration back to the organization with a mandatory note
[**PostBackofficeReviewsApprove**](BackofficeAPI.md#PostBackofficeReviewsApprove) | **Post** /backoffice/reviews/{version_id}/approve | Approve a submitted engagement version and publish it live, with an optional note; 409 when another reviewer decided first
[**PostBackofficeReviewsReject**](BackofficeAPI.md#PostBackofficeReviewsReject) | **Post** /backoffice/reviews/{version_id}/reject | Reject a submitted engagement version with a mandatory note; 409 when another reviewer decided first
[**PostBackofficeReviewsRequestChanges**](BackofficeAPI.md#PostBackofficeReviewsRequestChanges) | **Post** /backoffice/reviews/{version_id}/request-changes | Return a submitted engagement version to the organization with a mandatory note; 409 when another reviewer decided first
[**PostBackofficeSla**](BackofficeAPI.md#PostBackofficeSla) | **Post** /backoffice/sla | Publish a new SLA policy version; running reports keep their snapshotted deadlines
[**PostBackofficeTodosAssign**](BackofficeAPI.md#PostBackofficeTodosAssign) | **Post** /backoffice/todos/{report_id}/assign | Assign or reassign the report to a triager (user id, handle, or email); reassigning another triager&#39;s report needs the engagement-manager role or being the current assignee
[**PostBackofficeTodosComments**](BackofficeAPI.md#PostBackofficeTodosComments) | **Post** /backoffice/todos/{report_id}/comments | Post a message on the report&#39;s thread with a chosen audience (all, researcher, organization, or staff-only)
[**PostBackofficeTodosEscalate**](BackofficeAPI.md#PostBackofficeTodosEscalate) | **Post** /backoffice/todos/{report_id}/escalate | Escalate a report as critical with a reason; urgently notifies the organization and top-sorts the report (engagement-manager only)
[**PostBackofficeTodosMarkDuplicate**](BackofficeAPI.md#PostBackofficeTodosMarkDuplicate) | **Post** /backoffice/todos/{report_id}/mark-duplicate | Mark the report under triage as a duplicate of an original report in the same engagement; active assignee only
[**PostBackofficeTodosPickup**](BackofficeAPI.md#PostBackofficeTodosPickup) | **Post** /backoffice/todos/{report_id}/pickup | Pick up a new report for triage: assigns it to the caller and stamps the first response
[**PostBackofficeTodosProposeBounty**](BackofficeAPI.md#PostBackofficeTodosProposeBounty) | **Post** /backoffice/todos/{report_id}/propose-bounty | Propose a bounty on an accepted report, validated against the engagement&#39;s reward range
[**PostBackofficeTodosReject**](BackofficeAPI.md#PostBackofficeTodosReject) | **Post** /backoffice/todos/{report_id}/reject | Reject the report under triage with a mandatory reason; active assignee only
[**PostBackofficeTodosRequestInfo**](BackofficeAPI.md#PostBackofficeTodosRequestInfo) | **Post** /backoffice/todos/{report_id}/request-info | Flag the report as awaiting researcher input with a mandatory message; active assignee only
[**PostBackofficeTodosRequestInfoClear**](BackofficeAPI.md#PostBackofficeTodosRequestInfoClear) | **Post** /backoffice/todos/{report_id}/request-info/clear | Withdraw a pending information request; active assignee only, 409 when none is pending
[**PostBackofficeTodosRevertDuplicate**](BackofficeAPI.md#PostBackofficeTodosRevertDuplicate) | **Post** /backoffice/todos/{report_id}/revert-duplicate | Revert a duplicate marking back to triaging, clearing the cross-link; staff on the report&#39;s scope, 409 when not a duplicate
[**PostBackofficeTodosUnassign**](BackofficeAPI.md#PostBackofficeTodosUnassign) | **Post** /backoffice/todos/{report_id}/unassign | Release the report&#39;s triage assignment; it stays in triaging and shows as unassigned
[**PostBackofficeTodosValidate**](BackofficeAPI.md#PostBackofficeTodosValidate) | **Post** /backoffice/todos/{report_id}/validate | Validate the report under triage with a server-assessed CVSS vector; active assignee only, a justification is required when the assessed severity differs from the submitted one
[**PostBackofficeTopupsFail**](BackofficeAPI.md#PostBackofficeTopupsFail) | **Post** /backoffice/topups/{id}/fail | Fail a pending top-up with a reason; the balance is untouched
[**PostBackofficeTopupsSettle**](BackofficeAPI.md#PostBackofficeTopupsSettle) | **Post** /backoffice/topups/{id}/settle | Settle a pending top-up: credit the wallet and notify the requester
[**PostBackofficeVerificationsApprove**](BackofficeAPI.md#PostBackofficeVerificationsApprove) | **Post** /backoffice/verifications/{user_id}/approve | Approve a requested background check: raises the level to background
[**PostBackofficeVerificationsLevel**](BackofficeAPI.md#PostBackofficeVerificationsLevel) | **Post** /backoffice/verifications/{user_id}/level | Assign a researcher&#39;s verification level with a mandatory note
[**PostBackofficeVerificationsReject**](BackofficeAPI.md#PostBackofficeVerificationsReject) | **Post** /backoffice/verifications/{user_id}/reject | Reject a requested background check with a mandatory note



## GetBackofficeBadges

> BadgeList GetBackofficeBadges(ctx).Execute()

Staff badge catalog: every badge including inactive ones

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
	resp, r, err := apiClient.BackofficeAPI.GetBackofficeBadges(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.GetBackofficeBadges``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBackofficeBadges`: BadgeList
	fmt.Fprintf(os.Stdout, "Response from `BackofficeAPI.GetBackofficeBadges`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBackofficeBadgesRequest struct via the builder pattern


### Return type

[**BadgeList**](BadgeList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBackofficeBadgesByCode

> map[string]interface{} GetBackofficeBadgesByCode(ctx, code).Execute()

One badge with its holders

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
	code := "code_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackofficeAPI.GetBackofficeBadgesByCode(context.Background(), code).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.GetBackofficeBadgesByCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBackofficeBadgesByCode`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `BackofficeAPI.GetBackofficeBadgesByCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBackofficeBadgesByCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBackofficeBudgets

> OpsBudgets GetBackofficeBudgets(ctx).Execute()

The pending-report budget configuration: per-tier limits, slop-rejection penalties, and active per-researcher overrides

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
	resp, r, err := apiClient.BackofficeAPI.GetBackofficeBudgets(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.GetBackofficeBudgets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBackofficeBudgets`: OpsBudgets
	fmt.Fprintf(os.Stdout, "Response from `BackofficeAPI.GetBackofficeBudgets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBackofficeBudgetsRequest struct via the builder pattern


### Return type

[**OpsBudgets**](OpsBudgets.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBackofficeCertifications

> GetBackofficeCertifications(ctx).Execute()

Staff review queue of claimed researcher certifications

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
	r, err := apiClient.BackofficeAPI.GetBackofficeCertifications(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.GetBackofficeCertifications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBackofficeCertificationsRequest struct via the builder pattern


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


## GetBackofficeCompliance

> GetBackofficeCompliance(ctx).Execute()

Researcher payout compliance (sanctions) for the backoffice stopgap

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
	r, err := apiClient.BackofficeAPI.GetBackofficeCompliance(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.GetBackofficeCompliance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBackofficeComplianceRequest struct via the builder pattern


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


## GetBackofficeConduct

> ConductCaseList GetBackofficeConduct(ctx).Execute()

The trust & safety conduct case queue

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
	resp, r, err := apiClient.BackofficeAPI.GetBackofficeConduct(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.GetBackofficeConduct``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBackofficeConduct`: ConductCaseList
	fmt.Fprintf(os.Stdout, "Response from `BackofficeAPI.GetBackofficeConduct`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBackofficeConductRequest struct via the builder pattern


### Return type

[**ConductCaseList**](ConductCaseList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBackofficeConductByCaseId

> ConductCase GetBackofficeConductByCaseId(ctx, caseId).Execute()

One conduct case with its enforcement actions and appeals

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
	caseId := "caseId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackofficeAPI.GetBackofficeConductByCaseId(context.Background(), caseId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.GetBackofficeConductByCaseId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBackofficeConductByCaseId`: ConductCase
	fmt.Fprintf(os.Stdout, "Response from `BackofficeAPI.GetBackofficeConductByCaseId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBackofficeConductByCaseIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConductCase**](ConductCase.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBackofficeCustomerOnboarding

> CustomeronboardingRequestList GetBackofficeCustomerOnboarding(ctx).Execute()

The back-office queue of customer onboarding requests, filterable by status and customer email

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
	resp, r, err := apiClient.BackofficeAPI.GetBackofficeCustomerOnboarding(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.GetBackofficeCustomerOnboarding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBackofficeCustomerOnboarding`: CustomeronboardingRequestList
	fmt.Fprintf(os.Stdout, "Response from `BackofficeAPI.GetBackofficeCustomerOnboarding`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBackofficeCustomerOnboardingRequest struct via the builder pattern


### Return type

[**CustomeronboardingRequestList**](CustomeronboardingRequestList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBackofficeFeedback

> FeedbackList GetBackofficeFeedback(ctx).Execute()

The product feedback queue, optionally filtered by status

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
	resp, r, err := apiClient.BackofficeAPI.GetBackofficeFeedback(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.GetBackofficeFeedback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBackofficeFeedback`: FeedbackList
	fmt.Fprintf(os.Stdout, "Response from `BackofficeAPI.GetBackofficeFeedback`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBackofficeFeedbackRequest struct via the builder pattern


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


## GetBackofficeOrganizations

> OpsOrganizationRowList GetBackofficeOrganizations(ctx).Execute()

Every customer and partner organization with verification status and engagement counts

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
	resp, r, err := apiClient.BackofficeAPI.GetBackofficeOrganizations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.GetBackofficeOrganizations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBackofficeOrganizations`: OpsOrganizationRowList
	fmt.Fprintf(os.Stdout, "Response from `BackofficeAPI.GetBackofficeOrganizations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBackofficeOrganizationsRequest struct via the builder pattern


### Return type

[**OpsOrganizationRowList**](OpsOrganizationRowList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBackofficePayouts

> GetBackofficePayouts(ctx).Execute()

Payout run queue: eligible and blocked researchers, plus payouts by state

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
	r, err := apiClient.BackofficeAPI.GetBackofficePayouts(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.GetBackofficePayouts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBackofficePayoutsRequest struct via the builder pattern


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


## GetBackofficeProfileReviews

> GetBackofficeProfileReviews(ctx).Execute()

Organization profile versions awaiting review with their field diffs, plus the recently decided ones

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
	r, err := apiClient.BackofficeAPI.GetBackofficeProfileReviews(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.GetBackofficeProfileReviews``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBackofficeProfileReviewsRequest struct via the builder pattern


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


## GetBackofficeRegistrations

> GetBackofficeRegistrations(ctx).Execute()

Registration review queue with domain-verification evidence

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
	r, err := apiClient.BackofficeAPI.GetBackofficeRegistrations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.GetBackofficeRegistrations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBackofficeRegistrationsRequest struct via the builder pattern


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


## GetBackofficeSla

> GetBackofficeSla(ctx).Execute()

The active SLA policy matrix (first-response and resolution hours per engagement type and severity)

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
	r, err := apiClient.BackofficeAPI.GetBackofficeSla(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.GetBackofficeSla``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBackofficeSlaRequest struct via the builder pattern


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


## GetBackofficeTodos

> OpsTriageItemList GetBackofficeTodos(ctx).Status(status).Severity(severity).Q(q).Sort(sort).Dir(dir).Page(page).Size(size).Engagement(engagement).Assignment(assignment).Execute()

Staff triage inbox: reports awaiting or in triage, scoped to the caller's backoffice groups; filterable by status, engagement, severity, assignment, and search, sortable, paginated

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
	assignment := "assignment_example" // string | Assignee facet: any, unassigned, mine, or user:<id> (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackofficeAPI.GetBackofficeTodos(context.Background()).Status(status).Severity(severity).Q(q).Sort(sort).Dir(dir).Page(page).Size(size).Engagement(engagement).Assignment(assignment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.GetBackofficeTodos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBackofficeTodos`: OpsTriageItemList
	fmt.Fprintf(os.Stdout, "Response from `BackofficeAPI.GetBackofficeTodos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBackofficeTodosRequest struct via the builder pattern


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

### Return type

[**OpsTriageItemList**](OpsTriageItemList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBackofficeTodosByReportId

> OpsTriageWorkspace GetBackofficeTodosByReportId(ctx, reportId).Execute()

Triage workspace: the full report a triager assesses

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
	resp, r, err := apiClient.BackofficeAPI.GetBackofficeTodosByReportId(context.Background(), reportId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.GetBackofficeTodosByReportId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBackofficeTodosByReportId`: OpsTriageWorkspace
	fmt.Fprintf(os.Stdout, "Response from `BackofficeAPI.GetBackofficeTodosByReportId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBackofficeTodosByReportIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OpsTriageWorkspace**](OpsTriageWorkspace.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBackofficeTodosDuplicateCandidates

> GetBackofficeTodosDuplicateCandidates(ctx, reportId).Execute()

List candidate original reports in the engagement to mark a duplicate against (staff picker)

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
	r, err := apiClient.BackofficeAPI.GetBackofficeTodosDuplicateCandidates(context.Background(), reportId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.GetBackofficeTodosDuplicateCandidates``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetBackofficeTodosDuplicateCandidatesRequest struct via the builder pattern


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


## GetBackofficeTodosExport

> GetBackofficeTodosExport(ctx).Status(status).Severity(severity).Q(q).Sort(sort).Dir(dir).Page(page).Size(size).Engagement(engagement).Assignment(assignment).Format(format).Execute()

Download the filtered triage queue as CSV or JSON (every page, at most 10000 rows)

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
	assignment := "assignment_example" // string | Assignee facet: any, unassigned, mine, or user:<id> (optional)
	format := "format_example" // string | Download format (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackofficeAPI.GetBackofficeTodosExport(context.Background()).Status(status).Severity(severity).Q(q).Sort(sort).Dir(dir).Page(page).Size(size).Engagement(engagement).Assignment(assignment).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.GetBackofficeTodosExport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBackofficeTodosExportRequest struct via the builder pattern


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


## GetBackofficeTopups

> PaymentTopupQueueItemList GetBackofficeTopups(ctx).Execute()

Pending wallet top-ups awaiting reconciliation across organizations

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
	resp, r, err := apiClient.BackofficeAPI.GetBackofficeTopups(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.GetBackofficeTopups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBackofficeTopups`: PaymentTopupQueueItemList
	fmt.Fprintf(os.Stdout, "Response from `BackofficeAPI.GetBackofficeTopups`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBackofficeTopupsRequest struct via the builder pattern


### Return type

[**PaymentTopupQueueItemList**](PaymentTopupQueueItemList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBackofficeVerifications

> GetBackofficeVerifications(ctx).Execute()

Pending background checks and researcher verification levels

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
	r, err := apiClient.BackofficeAPI.GetBackofficeVerifications(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.GetBackofficeVerifications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBackofficeVerificationsRequest struct via the builder pattern


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


## PostBackofficeBadges

> Badge PostBackofficeBadges(ctx).Color(color).CriteriaKind(criteriaKind).DescriptionEn(descriptionEn).Icon(icon).NameEn(nameEn).Stars(stars).Code(code).DescriptionDe(descriptionDe).FromHour(fromHour).MinReports(minReports).NameDe(nameDe).Severity(severity).Threshold(threshold).ToHour(toHour).WindowDays(windowDays).Execute()

Create a badge; the code is stable, names and descriptions are per-locale with English required, and the criteria kind with threshold and parameters defines the earning logic

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
	color := "color_example" // string | Hex-shield color family
	criteriaKind := "criteriaKind_example" // string | Earning criteria kind; manual badges are only awarded by staff
	descriptionEn := "descriptionEn_example" // string | English description
	icon := "icon_example" // string | Glyph from the curated icon set
	nameEn := "nameEn_example" // string | English display name
	stars := "stars_example" // string | Tier stars (1-3)
	code := "code_example" // string | Stable badge code (create only) (optional)
	descriptionDe := "descriptionDe_example" // string | German (de-CH) description (optional)
	fromHour := "fromHour_example" // string | submission_hour_range only: UTC start hour (0-23) (optional)
	minReports := "minReports_example" // string | acceptance_rate only: minimum decided reports (optional)
	nameDe := "nameDe_example" // string | German (de-CH) display name (optional)
	severity := "severity_example" // string | accepted_reports only: restrict to one severity (optional)
	threshold := "threshold_example" // string | Criteria threshold (>= 1); unused for manual badges (optional)
	toHour := "toHour_example" // string | submission_hour_range only: UTC end hour (1-24, exclusive) (optional)
	windowDays := "windowDays_example" // string | accepted_in_window only: rolling window in days (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackofficeAPI.PostBackofficeBadges(context.Background()).Color(color).CriteriaKind(criteriaKind).DescriptionEn(descriptionEn).Icon(icon).NameEn(nameEn).Stars(stars).Code(code).DescriptionDe(descriptionDe).FromHour(fromHour).MinReports(minReports).NameDe(nameDe).Severity(severity).Threshold(threshold).ToHour(toHour).WindowDays(windowDays).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.PostBackofficeBadges``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostBackofficeBadges`: Badge
	fmt.Fprintf(os.Stdout, "Response from `BackofficeAPI.PostBackofficeBadges`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostBackofficeBadgesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **color** | **string** | Hex-shield color family | 
 **criteriaKind** | **string** | Earning criteria kind; manual badges are only awarded by staff | 
 **descriptionEn** | **string** | English description | 
 **icon** | **string** | Glyph from the curated icon set | 
 **nameEn** | **string** | English display name | 
 **stars** | **string** | Tier stars (1-3) | 
 **code** | **string** | Stable badge code (create only) | 
 **descriptionDe** | **string** | German (de-CH) description | 
 **fromHour** | **string** | submission_hour_range only: UTC start hour (0-23) | 
 **minReports** | **string** | acceptance_rate only: minimum decided reports | 
 **nameDe** | **string** | German (de-CH) display name | 
 **severity** | **string** | accepted_reports only: restrict to one severity | 
 **threshold** | **string** | Criteria threshold (&gt;&#x3D; 1); unused for manual badges | 
 **toHour** | **string** | submission_hour_range only: UTC end hour (1-24, exclusive) | 
 **windowDays** | **string** | accepted_in_window only: rolling window in days | 

### Return type

[**Badge**](Badge.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostBackofficeBadgesAward

> Badge PostBackofficeBadgesAward(ctx, code).Researcher(researcher).Note(note).Execute()

Manually award a badge to a researcher by handle or email; idempotent, and the researcher is notified only on a new award

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
	code := "code_example" // string | 
	researcher := "researcher_example" // string | Researcher handle or email
	note := "note_example" // string | Why the badge is granted; audited (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackofficeAPI.PostBackofficeBadgesAward(context.Background(), code).Researcher(researcher).Note(note).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.PostBackofficeBadgesAward``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostBackofficeBadgesAward`: Badge
	fmt.Fprintf(os.Stdout, "Response from `BackofficeAPI.PostBackofficeBadgesAward`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostBackofficeBadgesAwardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **researcher** | **string** | Researcher handle or email | 
 **note** | **string** | Why the badge is granted; audited | 

### Return type

[**Badge**](Badge.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostBackofficeBadgesAwardsRevoke

> PostBackofficeBadgesAwardsRevoke(ctx, code, awardId).Execute()

Revoke one badge award; audited

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
	code := "code_example" // string | 
	awardId := "awardId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackofficeAPI.PostBackofficeBadgesAwardsRevoke(context.Background(), code, awardId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.PostBackofficeBadgesAwardsRevoke``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** |  | 
**awardId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostBackofficeBadgesAwardsRevokeRequest struct via the builder pattern


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


## PostBackofficeBadgesByCode

> Badge PostBackofficeBadgesByCode(ctx, code).Color(color).CriteriaKind(criteriaKind).DescriptionEn(descriptionEn).Icon(icon).NameEn(nameEn).Stars(stars).Code2(code2).DescriptionDe(descriptionDe).FromHour(fromHour).MinReports(minReports).NameDe(nameDe).Severity(severity).Threshold(threshold).ToHour(toHour).WindowDays(windowDays).Execute()

Update a badge's texts, visuals, and criteria; the code cannot change

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
	code := "code_example" // string | 
	color := "color_example" // string | Hex-shield color family
	criteriaKind := "criteriaKind_example" // string | Earning criteria kind; manual badges are only awarded by staff
	descriptionEn := "descriptionEn_example" // string | English description
	icon := "icon_example" // string | Glyph from the curated icon set
	nameEn := "nameEn_example" // string | English display name
	stars := "stars_example" // string | Tier stars (1-3)
	code2 := "code_example" // string | Stable badge code (create only) (optional)
	descriptionDe := "descriptionDe_example" // string | German (de-CH) description (optional)
	fromHour := "fromHour_example" // string | submission_hour_range only: UTC start hour (0-23) (optional)
	minReports := "minReports_example" // string | acceptance_rate only: minimum decided reports (optional)
	nameDe := "nameDe_example" // string | German (de-CH) display name (optional)
	severity := "severity_example" // string | accepted_reports only: restrict to one severity (optional)
	threshold := "threshold_example" // string | Criteria threshold (>= 1); unused for manual badges (optional)
	toHour := "toHour_example" // string | submission_hour_range only: UTC end hour (1-24, exclusive) (optional)
	windowDays := "windowDays_example" // string | accepted_in_window only: rolling window in days (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackofficeAPI.PostBackofficeBadgesByCode(context.Background(), code).Color(color).CriteriaKind(criteriaKind).DescriptionEn(descriptionEn).Icon(icon).NameEn(nameEn).Stars(stars).Code2(code2).DescriptionDe(descriptionDe).FromHour(fromHour).MinReports(minReports).NameDe(nameDe).Severity(severity).Threshold(threshold).ToHour(toHour).WindowDays(windowDays).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.PostBackofficeBadgesByCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostBackofficeBadgesByCode`: Badge
	fmt.Fprintf(os.Stdout, "Response from `BackofficeAPI.PostBackofficeBadgesByCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostBackofficeBadgesByCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **color** | **string** | Hex-shield color family | 
 **criteriaKind** | **string** | Earning criteria kind; manual badges are only awarded by staff | 
 **descriptionEn** | **string** | English description | 
 **icon** | **string** | Glyph from the curated icon set | 
 **nameEn** | **string** | English display name | 
 **stars** | **string** | Tier stars (1-3) | 
 **code2** | **string** | Stable badge code (create only) | 
 **descriptionDe** | **string** | German (de-CH) description | 
 **fromHour** | **string** | submission_hour_range only: UTC start hour (0-23) | 
 **minReports** | **string** | acceptance_rate only: minimum decided reports | 
 **nameDe** | **string** | German (de-CH) display name | 
 **severity** | **string** | accepted_reports only: restrict to one severity | 
 **threshold** | **string** | Criteria threshold (&gt;&#x3D; 1); unused for manual badges | 
 **toHour** | **string** | submission_hour_range only: UTC end hour (1-24, exclusive) | 
 **windowDays** | **string** | accepted_in_window only: rolling window in days | 

### Return type

[**Badge**](Badge.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostBackofficeBadgesDeactivate

> PostBackofficeBadgesDeactivate(ctx, code).Execute()

Deactivate a badge: it disappears from the researcher catalog and profiles and stops being awarded; existing awards are kept

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
	code := "code_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackofficeAPI.PostBackofficeBadgesDeactivate(context.Background(), code).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.PostBackofficeBadgesDeactivate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostBackofficeBadgesDeactivateRequest struct via the builder pattern


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


## PostBackofficeBadgesReactivate

> PostBackofficeBadgesReactivate(ctx, code).Execute()

Reactivate a previously deactivated badge

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
	code := "code_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackofficeAPI.PostBackofficeBadgesReactivate(context.Background(), code).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.PostBackofficeBadgesReactivate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostBackofficeBadgesReactivateRequest struct via the builder pattern


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


## PostBackofficeBudgetsOverrides

> PostBackofficeBudgetsOverrides(ctx).Reason(reason).Researcher(researcher).Expires(expires).Limit(limit).Execute()

Create a per-researcher budget override with a mandatory reason and optional expiry

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
	reason := "reason_example" // string | Why the override exists; audited
	researcher := "researcher_example" // string | Researcher email or handle
	expires := "expires_example" // string | Optional expiry (RFC 3339 or yyyy-mm-dd); empty means permanent (optional)
	limit := "limit_example" // string | Pending-report limit; empty means unlimited (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackofficeAPI.PostBackofficeBudgetsOverrides(context.Background()).Reason(reason).Researcher(researcher).Expires(expires).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.PostBackofficeBudgetsOverrides``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostBackofficeBudgetsOverridesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reason** | **string** | Why the override exists; audited | 
 **researcher** | **string** | Researcher email or handle | 
 **expires** | **string** | Optional expiry (RFC 3339 or yyyy-mm-dd); empty means permanent | 
 **limit** | **string** | Pending-report limit; empty means unlimited | 

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


## PostBackofficeBudgetsOverridesDelete

> PostBackofficeBudgetsOverridesDelete(ctx, id).Execute()

Remove a budget override; the tier limit applies again

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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackofficeAPI.PostBackofficeBudgetsOverridesDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.PostBackofficeBudgetsOverridesDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostBackofficeBudgetsOverridesDeleteRequest struct via the builder pattern


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


## PostBackofficeBudgetsPenalties

> PostBackofficeBudgetsPenalties(ctx).Execute()

Update the spam/ai_slop rejection penalty points (always negative)

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
	r, err := apiClient.BackofficeAPI.PostBackofficeBudgetsPenalties(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.PostBackofficeBudgetsPenalties``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostBackofficeBudgetsPenaltiesRequest struct via the builder pattern


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


## PostBackofficeBudgetsTiers

> PostBackofficeBudgetsTiers(ctx).Execute()

Update the per-tier pending-report limits; empty means unlimited, limits must not decrease on higher tiers

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
	r, err := apiClient.BackofficeAPI.PostBackofficeBudgetsTiers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.PostBackofficeBudgetsTiers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostBackofficeBudgetsTiersRequest struct via the builder pattern


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


## PostBackofficeCertificationsReject

> PostBackofficeCertificationsReject(ctx, id).Note(note).Execute()

Reject a claimed certification (staff decision, audit-logged)

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
	id := "id_example" // string | 
	note := "note_example" // string | Reviewer note stored on the decision (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackofficeAPI.PostBackofficeCertificationsReject(context.Background(), id).Note(note).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.PostBackofficeCertificationsReject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostBackofficeCertificationsRejectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **note** | **string** | Reviewer note stored on the decision | 

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


## PostBackofficeCertificationsVerify

> PostBackofficeCertificationsVerify(ctx, id).Note(note).Execute()

Verify a claimed certification (staff decision, audit-logged)

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
	id := "id_example" // string | 
	note := "note_example" // string | Reviewer note stored on the decision (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackofficeAPI.PostBackofficeCertificationsVerify(context.Background(), id).Note(note).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.PostBackofficeCertificationsVerify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostBackofficeCertificationsVerifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **note** | **string** | Reviewer note stored on the decision | 

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


## PostBackofficeComplianceByUserId

> PostBackofficeComplianceByUserId(ctx, userId).Sanctions(sanctions).Execute()

Set a researcher's sanctions status

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
	userId := "userId_example" // string | 
	sanctions := "sanctions_example" // string | Sanctions status

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackofficeAPI.PostBackofficeComplianceByUserId(context.Background(), userId).Sanctions(sanctions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.PostBackofficeComplianceByUserId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostBackofficeComplianceByUserIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sanctions** | **string** | Sanctions status | 

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


## PostBackofficeConduct

> PostBackofficeConduct(ctx).Execute()

Open a conduct case against a researcher

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
	r, err := apiClient.BackofficeAPI.PostBackofficeConduct(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.PostBackofficeConduct``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostBackofficeConductRequest struct via the builder pattern


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


## PostBackofficeConductAppealResolve

> PostBackofficeConductAppealResolve(ctx, caseId, appealId).Execute()

Uphold or overturn a conduct appeal

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
	caseId := "caseId_example" // string | 
	appealId := "appealId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackofficeAPI.PostBackofficeConductAppealResolve(context.Background(), caseId, appealId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.PostBackofficeConductAppealResolve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseId** | **string** |  | 
**appealId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostBackofficeConductAppealResolveRequest struct via the builder pattern


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


## PostBackofficeConductClose

> PostBackofficeConductClose(ctx, caseId).Execute()

Close a decided conduct case

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
	caseId := "caseId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackofficeAPI.PostBackofficeConductClose(context.Background(), caseId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.PostBackofficeConductClose``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostBackofficeConductCloseRequest struct via the builder pattern


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


## PostBackofficeConductDecide

> PostBackofficeConductDecide(ctx, caseId).Execute()

Decide a conduct case

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
	caseId := "caseId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackofficeAPI.PostBackofficeConductDecide(context.Background(), caseId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.PostBackofficeConductDecide``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostBackofficeConductDecideRequest struct via the builder pattern


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


## PostBackofficeConductDismiss

> PostBackofficeConductDismiss(ctx, caseId).Execute()

Dismiss a conduct case

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
	caseId := "caseId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackofficeAPI.PostBackofficeConductDismiss(context.Background(), caseId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.PostBackofficeConductDismiss``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostBackofficeConductDismissRequest struct via the builder pattern


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


## PostBackofficeConductEnforce

> PostBackofficeConductEnforce(ctx, caseId).Execute()

Record a graduated enforcement action on a decided case

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
	caseId := "caseId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackofficeAPI.PostBackofficeConductEnforce(context.Background(), caseId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.PostBackofficeConductEnforce``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostBackofficeConductEnforceRequest struct via the builder pattern


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


## PostBackofficeConductEnforceRevoke

> PostBackofficeConductEnforceRevoke(ctx, caseId, actionId).Execute()

Lift an active enforcement action

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
	caseId := "caseId_example" // string | 
	actionId := "actionId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackofficeAPI.PostBackofficeConductEnforceRevoke(context.Background(), caseId, actionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.PostBackofficeConductEnforceRevoke``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseId** | **string** |  | 
**actionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostBackofficeConductEnforceRevokeRequest struct via the builder pattern


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


## PostBackofficeConductNote

> PostBackofficeConductNote(ctx, caseId).Execute()

Add an investigation note to a case

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
	caseId := "caseId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackofficeAPI.PostBackofficeConductNote(context.Background(), caseId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.PostBackofficeConductNote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostBackofficeConductNoteRequest struct via the builder pattern


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


## PostBackofficeConductReportFlag

> PostBackofficeConductReportFlag(ctx, reportId).Execute()

Open a conduct case pre-linked to a report and its researcher

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
	r, err := apiClient.BackofficeAPI.PostBackofficeConductReportFlag(context.Background(), reportId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.PostBackofficeConductReportFlag``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostBackofficeConductReportFlagRequest struct via the builder pattern


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


## PostBackofficeConductTakeUp

> PostBackofficeConductTakeUp(ctx, caseId).Execute()

Take up a conduct case for investigation

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
	caseId := "caseId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackofficeAPI.PostBackofficeConductTakeUp(context.Background(), caseId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.PostBackofficeConductTakeUp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostBackofficeConductTakeUpRequest struct via the builder pattern


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


## PostBackofficeCustomerOnboarding

> PostBackofficeCustomerOnboarding(ctx).CustomerEmail(customerEmail).Payload(payload).Execute()

Create a prefilled customer onboarding request and get its share link for the customer

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
	customerEmail := "customerEmail_example" // string | Customer contact email (optional)
	payload := "payload_example" // string | Prefilled wizard state (JSON), including the selected programs (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackofficeAPI.PostBackofficeCustomerOnboarding(context.Background()).CustomerEmail(customerEmail).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.PostBackofficeCustomerOnboarding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostBackofficeCustomerOnboardingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerEmail** | **string** | Customer contact email | 
 **payload** | **string** | Prefilled wizard state (JSON), including the selected programs | 

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


## PostBackofficeCustomerOnboardingReject

> PostBackofficeCustomerOnboardingReject(ctx, requestId).Note(note).Execute()

Reject a customer onboarding request with an optional note; 409 once it is closed

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
	requestId := "requestId_example" // string | 
	note := "note_example" // string | Reviewer rejection note (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackofficeAPI.PostBackofficeCustomerOnboardingReject(context.Background(), requestId).Note(note).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.PostBackofficeCustomerOnboardingReject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostBackofficeCustomerOnboardingRejectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **note** | **string** | Reviewer rejection note | 

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


## PostBackofficeFeedbackStatus

> PostBackofficeFeedbackStatus(ctx, feedbackId).Execute()

Move a feedback entry to new, reviewed, or resolved

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
	r, err := apiClient.BackofficeAPI.PostBackofficeFeedbackStatus(context.Background(), feedbackId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.PostBackofficeFeedbackStatus``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostBackofficeFeedbackStatusRequest struct via the builder pattern


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


## PostBackofficePayoutsApprove

> PostBackofficePayoutsApprove(ctx, id).Execute()

Approve a pending payout

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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackofficeAPI.PostBackofficePayoutsApprove(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.PostBackofficePayoutsApprove``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostBackofficePayoutsApproveRequest struct via the builder pattern


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


## PostBackofficePayoutsFail

> PostBackofficePayoutsFail(ctx, id).Execute()

Mark a processing payout failed with a reason

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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackofficeAPI.PostBackofficePayoutsFail(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.PostBackofficePayoutsFail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostBackofficePayoutsFailRequest struct via the builder pattern


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


## PostBackofficePayoutsPaid

> PostBackofficePayoutsPaid(ctx, id).Execute()

Mark a processing payout paid; flips its source rewards to paid

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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackofficeAPI.PostBackofficePayoutsPaid(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.PostBackofficePayoutsPaid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostBackofficePayoutsPaidRequest struct via the builder pattern


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


## PostBackofficePayoutsProcess

> PostBackofficePayoutsProcess(ctx, id).Execute()

Start processing an approved payout

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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackofficeAPI.PostBackofficePayoutsProcess(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.PostBackofficePayoutsProcess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostBackofficePayoutsProcessRequest struct via the builder pattern


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


## PostBackofficePayoutsReject

> PostBackofficePayoutsReject(ctx, id).Execute()

Cancel a pending payout

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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackofficeAPI.PostBackofficePayoutsReject(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.PostBackofficePayoutsReject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostBackofficePayoutsRejectRequest struct via the builder pattern


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


## PostBackofficePayoutsRetry

> PostBackofficePayoutsRetry(ctx, id).Execute()

Retry a failed payout

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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackofficeAPI.PostBackofficePayoutsRetry(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.PostBackofficePayoutsRetry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostBackofficePayoutsRetryRequest struct via the builder pattern


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


## PostBackofficePayoutsRun

> PostBackofficePayoutsRun(ctx).Execute()

Run payouts: create one payout per eligible researcher (async, idempotent)

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
	r, err := apiClient.BackofficeAPI.PostBackofficePayoutsRun(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.PostBackofficePayoutsRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostBackofficePayoutsRunRequest struct via the builder pattern


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


## PostBackofficeProfileReviewsApprove

> PostBackofficeProfileReviewsApprove(ctx, versionId).Execute()

Approve the profile version and publish it as the organization's live profile

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
	versionId := "versionId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackofficeAPI.PostBackofficeProfileReviewsApprove(context.Background(), versionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.PostBackofficeProfileReviewsApprove``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**versionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostBackofficeProfileReviewsApproveRequest struct via the builder pattern


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


## PostBackofficeProfileReviewsChanges

> PostBackofficeProfileReviewsChanges(ctx, versionId).Execute()

Send the profile version back to its authors with a required note

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
	versionId := "versionId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackofficeAPI.PostBackofficeProfileReviewsChanges(context.Background(), versionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.PostBackofficeProfileReviewsChanges``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**versionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostBackofficeProfileReviewsChangesRequest struct via the builder pattern


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


## PostBackofficeProfileReviewsReject

> PostBackofficeProfileReviewsReject(ctx, versionId).Execute()

Reject the profile version with a required note; rejected is final

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
	versionId := "versionId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackofficeAPI.PostBackofficeProfileReviewsReject(context.Background(), versionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.PostBackofficeProfileReviewsReject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**versionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostBackofficeProfileReviewsRejectRequest struct via the builder pattern


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


## PostBackofficeRegistrationsApprove

> PostBackofficeRegistrationsApprove(ctx, id).Note(note).Execute()

Approve the registration: the organization becomes verified and may run engagements

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
	id := "id_example" // string | 
	note := "note_example" // string | Reviewer note stored on the decision (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackofficeAPI.PostBackofficeRegistrationsApprove(context.Background(), id).Note(note).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.PostBackofficeRegistrationsApprove``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostBackofficeRegistrationsApproveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **note** | **string** | Reviewer note stored on the decision | 

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


## PostBackofficeRegistrationsReject

> PostBackofficeRegistrationsReject(ctx, id).Note(note).Execute()

Reject the registration with a mandatory reason

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
	id := "id_example" // string | 
	note := "note_example" // string | Reviewer note stored on the decision (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackofficeAPI.PostBackofficeRegistrationsReject(context.Background(), id).Note(note).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.PostBackofficeRegistrationsReject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostBackofficeRegistrationsRejectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **note** | **string** | Reviewer note stored on the decision | 

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


## PostBackofficeRegistrationsReopen

> PostBackofficeRegistrationsReopen(ctx, id).Note(note).Execute()

Re-open an approved registration; the workspace gate re-engages

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
	id := "id_example" // string | 
	note := "note_example" // string | Reviewer note stored on the decision (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackofficeAPI.PostBackofficeRegistrationsReopen(context.Background(), id).Note(note).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.PostBackofficeRegistrationsReopen``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostBackofficeRegistrationsReopenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **note** | **string** | Reviewer note stored on the decision | 

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


## PostBackofficeRegistrationsRequestInfo

> PostBackofficeRegistrationsRequestInfo(ctx, id).Note(note).Execute()

Send the registration back to the organization with a mandatory note

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
	id := "id_example" // string | 
	note := "note_example" // string | Reviewer note stored on the decision (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackofficeAPI.PostBackofficeRegistrationsRequestInfo(context.Background(), id).Note(note).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.PostBackofficeRegistrationsRequestInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostBackofficeRegistrationsRequestInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **note** | **string** | Reviewer note stored on the decision | 

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


## PostBackofficeReviewsApprove

> PostBackofficeReviewsApprove(ctx, versionId).Note(note).Execute()

Approve a submitted engagement version and publish it live, with an optional note; 409 when another reviewer decided first

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
	versionId := "versionId_example" // string | 
	note := "note_example" // string | Note stored on the version transition (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackofficeAPI.PostBackofficeReviewsApprove(context.Background(), versionId).Note(note).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.PostBackofficeReviewsApprove``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**versionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostBackofficeReviewsApproveRequest struct via the builder pattern


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


## PostBackofficeReviewsReject

> PostBackofficeReviewsReject(ctx, versionId).Note(note).Execute()

Reject a submitted engagement version with a mandatory note; 409 when another reviewer decided first

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
	versionId := "versionId_example" // string | 
	note := "note_example" // string | Note stored on the version transition

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackofficeAPI.PostBackofficeReviewsReject(context.Background(), versionId).Note(note).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.PostBackofficeReviewsReject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**versionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostBackofficeReviewsRejectRequest struct via the builder pattern


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


## PostBackofficeReviewsRequestChanges

> PostBackofficeReviewsRequestChanges(ctx, versionId).Note(note).Execute()

Return a submitted engagement version to the organization with a mandatory note; 409 when another reviewer decided first

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
	versionId := "versionId_example" // string | 
	note := "note_example" // string | Note stored on the version transition

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackofficeAPI.PostBackofficeReviewsRequestChanges(context.Background(), versionId).Note(note).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.PostBackofficeReviewsRequestChanges``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**versionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostBackofficeReviewsRequestChangesRequest struct via the builder pattern


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


## PostBackofficeSla

> PostBackofficeSla(ctx).Execute()

Publish a new SLA policy version; running reports keep their snapshotted deadlines

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
	r, err := apiClient.BackofficeAPI.PostBackofficeSla(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.PostBackofficeSla``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostBackofficeSlaRequest struct via the builder pattern


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


## PostBackofficeTodosAssign

> OpsTriageWorkspace PostBackofficeTodosAssign(ctx, reportId).Assignee(assignee).Reason(reason).Execute()

Assign or reassign the report to a triager (user id, handle, or email); reassigning another triager's report needs the engagement-manager role or being the current assignee

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
	assignee := "assignee_example" // string | User id, handle, or email of the triager
	reason := "reason_example" // string | Optional reason, recorded in the audit trail (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackofficeAPI.PostBackofficeTodosAssign(context.Background(), reportId).Assignee(assignee).Reason(reason).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.PostBackofficeTodosAssign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostBackofficeTodosAssign`: OpsTriageWorkspace
	fmt.Fprintf(os.Stdout, "Response from `BackofficeAPI.PostBackofficeTodosAssign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostBackofficeTodosAssignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assignee** | **string** | User id, handle, or email of the triager | 
 **reason** | **string** | Optional reason, recorded in the audit trail | 

### Return type

[**OpsTriageWorkspace**](OpsTriageWorkspace.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostBackofficeTodosComments

> PostBackofficeTodosComments(ctx, reportId).Body(body).Audience(audience).Execute()

Post a message on the report's thread with a chosen audience (all, researcher, organization, or staff-only)

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
	audience := "audience_example" // string | Message audience; staff may use every audience, default staff (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackofficeAPI.PostBackofficeTodosComments(context.Background(), reportId).Body(body).Audience(audience).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.PostBackofficeTodosComments``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostBackofficeTodosCommentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** | The message text | 
 **audience** | **string** | Message audience; staff may use every audience, default staff | 

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


## PostBackofficeTodosEscalate

> PostBackofficeTodosEscalate(ctx, reportId).Execute()

Escalate a report as critical with a reason; urgently notifies the organization and top-sorts the report (engagement-manager only)

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
	r, err := apiClient.BackofficeAPI.PostBackofficeTodosEscalate(context.Background(), reportId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.PostBackofficeTodosEscalate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostBackofficeTodosEscalateRequest struct via the builder pattern


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


## PostBackofficeTodosMarkDuplicate

> OpsTriageWorkspace PostBackofficeTodosMarkDuplicate(ctx, reportId).OriginalId(originalId).Execute()

Mark the report under triage as a duplicate of an original report in the same engagement; active assignee only

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
	originalId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The original report's id in the same engagement

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackofficeAPI.PostBackofficeTodosMarkDuplicate(context.Background(), reportId).OriginalId(originalId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.PostBackofficeTodosMarkDuplicate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostBackofficeTodosMarkDuplicate`: OpsTriageWorkspace
	fmt.Fprintf(os.Stdout, "Response from `BackofficeAPI.PostBackofficeTodosMarkDuplicate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostBackofficeTodosMarkDuplicateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **originalId** | **string** | The original report&#39;s id in the same engagement | 

### Return type

[**OpsTriageWorkspace**](OpsTriageWorkspace.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostBackofficeTodosPickup

> OpsTriageWorkspace PostBackofficeTodosPickup(ctx, reportId).Execute()

Pick up a new report for triage: assigns it to the caller and stamps the first response

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
	resp, r, err := apiClient.BackofficeAPI.PostBackofficeTodosPickup(context.Background(), reportId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.PostBackofficeTodosPickup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostBackofficeTodosPickup`: OpsTriageWorkspace
	fmt.Fprintf(os.Stdout, "Response from `BackofficeAPI.PostBackofficeTodosPickup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostBackofficeTodosPickupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OpsTriageWorkspace**](OpsTriageWorkspace.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostBackofficeTodosProposeBounty

> PostBackofficeTodosProposeBounty(ctx, reportId).Amount(amount).Note(note).Execute()

Propose a bounty on an accepted report, validated against the engagement's reward range

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
	amount := "amount_example" // string | Bounty amount in whole CHF
	note := "note_example" // string | Optional note for the finance approver (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackofficeAPI.PostBackofficeTodosProposeBounty(context.Background(), reportId).Amount(amount).Note(note).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.PostBackofficeTodosProposeBounty``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostBackofficeTodosProposeBountyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **amount** | **string** | Bounty amount in whole CHF | 
 **note** | **string** | Optional note for the finance approver | 

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


## PostBackofficeTodosReject

> OpsTriageWorkspace PostBackofficeTodosReject(ctx, reportId).Reason(reason).Category(category).Execute()

Reject the report under triage with a mandatory reason; active assignee only

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
	reason := "reason_example" // string | Why the report is rejected; shared with the researcher
	category := "category_example" // string | Rejection category; spam and ai_slop apply an additional reputation penalty (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackofficeAPI.PostBackofficeTodosReject(context.Background(), reportId).Reason(reason).Category(category).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.PostBackofficeTodosReject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostBackofficeTodosReject`: OpsTriageWorkspace
	fmt.Fprintf(os.Stdout, "Response from `BackofficeAPI.PostBackofficeTodosReject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostBackofficeTodosRejectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reason** | **string** | Why the report is rejected; shared with the researcher | 
 **category** | **string** | Rejection category; spam and ai_slop apply an additional reputation penalty | 

### Return type

[**OpsTriageWorkspace**](OpsTriageWorkspace.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostBackofficeTodosRequestInfo

> PostBackofficeTodosRequestInfo(ctx, reportId).Body(body).Execute()

Flag the report as awaiting researcher input with a mandatory message; active assignee only

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
	body := "body_example" // string | What information the researcher should provide; posted as a researcher-visible message

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackofficeAPI.PostBackofficeTodosRequestInfo(context.Background(), reportId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.PostBackofficeTodosRequestInfo``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostBackofficeTodosRequestInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** | What information the researcher should provide; posted as a researcher-visible message | 

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


## PostBackofficeTodosRequestInfoClear

> PostBackofficeTodosRequestInfoClear(ctx, reportId).Execute()

Withdraw a pending information request; active assignee only, 409 when none is pending

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
	r, err := apiClient.BackofficeAPI.PostBackofficeTodosRequestInfoClear(context.Background(), reportId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.PostBackofficeTodosRequestInfoClear``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostBackofficeTodosRequestInfoClearRequest struct via the builder pattern


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


## PostBackofficeTodosRevertDuplicate

> PostBackofficeTodosRevertDuplicate(ctx, reportId).Execute()

Revert a duplicate marking back to triaging, clearing the cross-link; staff on the report's scope, 409 when not a duplicate

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
	r, err := apiClient.BackofficeAPI.PostBackofficeTodosRevertDuplicate(context.Background(), reportId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.PostBackofficeTodosRevertDuplicate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostBackofficeTodosRevertDuplicateRequest struct via the builder pattern


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


## PostBackofficeTodosUnassign

> PostBackofficeTodosUnassign(ctx, reportId).Execute()

Release the report's triage assignment; it stays in triaging and shows as unassigned

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
	r, err := apiClient.BackofficeAPI.PostBackofficeTodosUnassign(context.Background(), reportId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.PostBackofficeTodosUnassign``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostBackofficeTodosUnassignRequest struct via the builder pattern


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


## PostBackofficeTodosValidate

> OpsTriageWorkspace PostBackofficeTodosValidate(ctx, reportId).Ac(ac).Av(av).Avail(avail).Conf(conf).Integ(integ).Pr(pr).Scope(scope).Ui(ui).Justification(justification).Execute()

Validate the report under triage with a server-assessed CVSS vector; active assignee only, a justification is required when the assessed severity differs from the submitted one

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
	ac := "ac_example" // string | CVSS attack complexity
	av := "av_example" // string | CVSS attack vector
	avail := "avail_example" // string | CVSS availability impact
	conf := "conf_example" // string | CVSS confidentiality impact
	integ := "integ_example" // string | CVSS integrity impact
	pr := "pr_example" // string | CVSS privileges required
	scope := "scope_example" // string | CVSS scope
	ui := "ui_example" // string | CVSS user interaction
	justification := "justification_example" // string | Why the assessed severity differs from the submitted one; required on a severity change (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackofficeAPI.PostBackofficeTodosValidate(context.Background(), reportId).Ac(ac).Av(av).Avail(avail).Conf(conf).Integ(integ).Pr(pr).Scope(scope).Ui(ui).Justification(justification).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.PostBackofficeTodosValidate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostBackofficeTodosValidate`: OpsTriageWorkspace
	fmt.Fprintf(os.Stdout, "Response from `BackofficeAPI.PostBackofficeTodosValidate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostBackofficeTodosValidateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ac** | **string** | CVSS attack complexity | 
 **av** | **string** | CVSS attack vector | 
 **avail** | **string** | CVSS availability impact | 
 **conf** | **string** | CVSS confidentiality impact | 
 **integ** | **string** | CVSS integrity impact | 
 **pr** | **string** | CVSS privileges required | 
 **scope** | **string** | CVSS scope | 
 **ui** | **string** | CVSS user interaction | 
 **justification** | **string** | Why the assessed severity differs from the submitted one; required on a severity change | 

### Return type

[**OpsTriageWorkspace**](OpsTriageWorkspace.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostBackofficeTopupsFail

> PostBackofficeTopupsFail(ctx, id).Reason(reason).Execute()

Fail a pending top-up with a reason; the balance is untouched

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
	id := "id_example" // string | 
	reason := "reason_example" // string | Why the top-up could not be settled

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackofficeAPI.PostBackofficeTopupsFail(context.Background(), id).Reason(reason).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.PostBackofficeTopupsFail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostBackofficeTopupsFailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reason** | **string** | Why the top-up could not be settled | 

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


## PostBackofficeTopupsSettle

> PostBackofficeTopupsSettle(ctx, id).Execute()

Settle a pending top-up: credit the wallet and notify the requester

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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackofficeAPI.PostBackofficeTopupsSettle(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.PostBackofficeTopupsSettle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostBackofficeTopupsSettleRequest struct via the builder pattern


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


## PostBackofficeVerificationsApprove

> PostBackofficeVerificationsApprove(ctx, userId).Execute()

Approve a requested background check: raises the level to background

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
	userId := "userId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackofficeAPI.PostBackofficeVerificationsApprove(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.PostBackofficeVerificationsApprove``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostBackofficeVerificationsApproveRequest struct via the builder pattern


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


## PostBackofficeVerificationsLevel

> PostBackofficeVerificationsLevel(ctx, userId).Level(level).Note(note).Execute()

Assign a researcher's verification level with a mandatory note

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
	userId := "userId_example" // string | 
	level := "level_example" // string | Verification level to assign
	note := "note_example" // string | What was checked

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackofficeAPI.PostBackofficeVerificationsLevel(context.Background(), userId).Level(level).Note(note).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.PostBackofficeVerificationsLevel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostBackofficeVerificationsLevelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **level** | **string** | Verification level to assign | 
 **note** | **string** | What was checked | 

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


## PostBackofficeVerificationsReject

> PostBackofficeVerificationsReject(ctx, userId).Note(note).Execute()

Reject a requested background check with a mandatory note

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
	userId := "userId_example" // string | 
	note := "note_example" // string | Why the background check was rejected

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackofficeAPI.PostBackofficeVerificationsReject(context.Background(), userId).Note(note).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackofficeAPI.PostBackofficeVerificationsReject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostBackofficeVerificationsRejectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **note** | **string** | Why the background check was rejected | 

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

