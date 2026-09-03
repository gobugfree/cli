# ReportOrgReportDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignmentOptions** | [**[]ReportAssignee**](ReportAssignee.md) |  | 
**Comments** | [**[]ReportScopedComment**](ReportScopedComment.md) |  | 
**Contributors** | Pointer to [**[]ReportContributor**](ReportContributor.md) |  | [optional] 
**CvssVector** | Pointer to **string** |  | [optional] 
**Cwe** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Evidence** | Pointer to [**[]ReportEvidence**](ReportEvidence.md) |  | [optional] 
**Impact** | Pointer to **string** |  | [optional] 
**InfoRequestedAt** | Pointer to **time.Time** |  | [optional] 
**Origin** | **string** |  | 
**Poc** | Pointer to **string** |  | [optional] 
**RejectionReason** | Pointer to **string** |  | [optional] 
**Retest** | Pointer to [**ReportRetest**](ReportRetest.md) |  | [optional] 
**RetestHistory** | Pointer to [**[]ReportRetest**](ReportRetest.md) |  | [optional] 
**SeverityJustification** | Pointer to **string** |  | [optional] 
**StatusHistory** | [**[]ReportStatusEvent**](ReportStatusEvent.md) |  | 
**Steps** | Pointer to **string** |  | [optional] 

## Methods

### NewReportOrgReportDetail

`func NewReportOrgReportDetail(assignmentOptions []ReportAssignee, comments []ReportScopedComment, origin string, statusHistory []ReportStatusEvent, ) *ReportOrgReportDetail`

NewReportOrgReportDetail instantiates a new ReportOrgReportDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportOrgReportDetailWithDefaults

`func NewReportOrgReportDetailWithDefaults() *ReportOrgReportDetail`

NewReportOrgReportDetailWithDefaults instantiates a new ReportOrgReportDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignmentOptions

`func (o *ReportOrgReportDetail) GetAssignmentOptions() []ReportAssignee`

GetAssignmentOptions returns the AssignmentOptions field if non-nil, zero value otherwise.

### GetAssignmentOptionsOk

`func (o *ReportOrgReportDetail) GetAssignmentOptionsOk() (*[]ReportAssignee, bool)`

GetAssignmentOptionsOk returns a tuple with the AssignmentOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentOptions

`func (o *ReportOrgReportDetail) SetAssignmentOptions(v []ReportAssignee)`

SetAssignmentOptions sets AssignmentOptions field to given value.


### GetComments

`func (o *ReportOrgReportDetail) GetComments() []ReportScopedComment`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *ReportOrgReportDetail) GetCommentsOk() (*[]ReportScopedComment, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *ReportOrgReportDetail) SetComments(v []ReportScopedComment)`

SetComments sets Comments field to given value.


### GetContributors

`func (o *ReportOrgReportDetail) GetContributors() []ReportContributor`

GetContributors returns the Contributors field if non-nil, zero value otherwise.

### GetContributorsOk

`func (o *ReportOrgReportDetail) GetContributorsOk() (*[]ReportContributor, bool)`

GetContributorsOk returns a tuple with the Contributors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributors

`func (o *ReportOrgReportDetail) SetContributors(v []ReportContributor)`

SetContributors sets Contributors field to given value.

### HasContributors

`func (o *ReportOrgReportDetail) HasContributors() bool`

HasContributors returns a boolean if a field has been set.

### GetCvssVector

`func (o *ReportOrgReportDetail) GetCvssVector() string`

GetCvssVector returns the CvssVector field if non-nil, zero value otherwise.

### GetCvssVectorOk

`func (o *ReportOrgReportDetail) GetCvssVectorOk() (*string, bool)`

GetCvssVectorOk returns a tuple with the CvssVector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssVector

`func (o *ReportOrgReportDetail) SetCvssVector(v string)`

SetCvssVector sets CvssVector field to given value.

### HasCvssVector

`func (o *ReportOrgReportDetail) HasCvssVector() bool`

HasCvssVector returns a boolean if a field has been set.

### GetCwe

`func (o *ReportOrgReportDetail) GetCwe() string`

GetCwe returns the Cwe field if non-nil, zero value otherwise.

### GetCweOk

`func (o *ReportOrgReportDetail) GetCweOk() (*string, bool)`

GetCweOk returns a tuple with the Cwe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCwe

`func (o *ReportOrgReportDetail) SetCwe(v string)`

SetCwe sets Cwe field to given value.

### HasCwe

`func (o *ReportOrgReportDetail) HasCwe() bool`

HasCwe returns a boolean if a field has been set.

### GetDescription

`func (o *ReportOrgReportDetail) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ReportOrgReportDetail) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ReportOrgReportDetail) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ReportOrgReportDetail) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEvidence

`func (o *ReportOrgReportDetail) GetEvidence() []ReportEvidence`

GetEvidence returns the Evidence field if non-nil, zero value otherwise.

### GetEvidenceOk

`func (o *ReportOrgReportDetail) GetEvidenceOk() (*[]ReportEvidence, bool)`

GetEvidenceOk returns a tuple with the Evidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidence

`func (o *ReportOrgReportDetail) SetEvidence(v []ReportEvidence)`

SetEvidence sets Evidence field to given value.

### HasEvidence

`func (o *ReportOrgReportDetail) HasEvidence() bool`

HasEvidence returns a boolean if a field has been set.

### GetImpact

`func (o *ReportOrgReportDetail) GetImpact() string`

GetImpact returns the Impact field if non-nil, zero value otherwise.

### GetImpactOk

`func (o *ReportOrgReportDetail) GetImpactOk() (*string, bool)`

GetImpactOk returns a tuple with the Impact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpact

`func (o *ReportOrgReportDetail) SetImpact(v string)`

SetImpact sets Impact field to given value.

### HasImpact

`func (o *ReportOrgReportDetail) HasImpact() bool`

HasImpact returns a boolean if a field has been set.

### GetInfoRequestedAt

`func (o *ReportOrgReportDetail) GetInfoRequestedAt() time.Time`

GetInfoRequestedAt returns the InfoRequestedAt field if non-nil, zero value otherwise.

### GetInfoRequestedAtOk

`func (o *ReportOrgReportDetail) GetInfoRequestedAtOk() (*time.Time, bool)`

GetInfoRequestedAtOk returns a tuple with the InfoRequestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfoRequestedAt

`func (o *ReportOrgReportDetail) SetInfoRequestedAt(v time.Time)`

SetInfoRequestedAt sets InfoRequestedAt field to given value.

### HasInfoRequestedAt

`func (o *ReportOrgReportDetail) HasInfoRequestedAt() bool`

HasInfoRequestedAt returns a boolean if a field has been set.

### GetOrigin

`func (o *ReportOrgReportDetail) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *ReportOrgReportDetail) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *ReportOrgReportDetail) SetOrigin(v string)`

SetOrigin sets Origin field to given value.


### GetPoc

`func (o *ReportOrgReportDetail) GetPoc() string`

GetPoc returns the Poc field if non-nil, zero value otherwise.

### GetPocOk

`func (o *ReportOrgReportDetail) GetPocOk() (*string, bool)`

GetPocOk returns a tuple with the Poc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoc

`func (o *ReportOrgReportDetail) SetPoc(v string)`

SetPoc sets Poc field to given value.

### HasPoc

`func (o *ReportOrgReportDetail) HasPoc() bool`

HasPoc returns a boolean if a field has been set.

### GetRejectionReason

`func (o *ReportOrgReportDetail) GetRejectionReason() string`

GetRejectionReason returns the RejectionReason field if non-nil, zero value otherwise.

### GetRejectionReasonOk

`func (o *ReportOrgReportDetail) GetRejectionReasonOk() (*string, bool)`

GetRejectionReasonOk returns a tuple with the RejectionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectionReason

`func (o *ReportOrgReportDetail) SetRejectionReason(v string)`

SetRejectionReason sets RejectionReason field to given value.

### HasRejectionReason

`func (o *ReportOrgReportDetail) HasRejectionReason() bool`

HasRejectionReason returns a boolean if a field has been set.

### GetRetest

`func (o *ReportOrgReportDetail) GetRetest() ReportRetest`

GetRetest returns the Retest field if non-nil, zero value otherwise.

### GetRetestOk

`func (o *ReportOrgReportDetail) GetRetestOk() (*ReportRetest, bool)`

GetRetestOk returns a tuple with the Retest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetest

`func (o *ReportOrgReportDetail) SetRetest(v ReportRetest)`

SetRetest sets Retest field to given value.

### HasRetest

`func (o *ReportOrgReportDetail) HasRetest() bool`

HasRetest returns a boolean if a field has been set.

### GetRetestHistory

`func (o *ReportOrgReportDetail) GetRetestHistory() []ReportRetest`

GetRetestHistory returns the RetestHistory field if non-nil, zero value otherwise.

### GetRetestHistoryOk

`func (o *ReportOrgReportDetail) GetRetestHistoryOk() (*[]ReportRetest, bool)`

GetRetestHistoryOk returns a tuple with the RetestHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetestHistory

`func (o *ReportOrgReportDetail) SetRetestHistory(v []ReportRetest)`

SetRetestHistory sets RetestHistory field to given value.

### HasRetestHistory

`func (o *ReportOrgReportDetail) HasRetestHistory() bool`

HasRetestHistory returns a boolean if a field has been set.

### GetSeverityJustification

`func (o *ReportOrgReportDetail) GetSeverityJustification() string`

GetSeverityJustification returns the SeverityJustification field if non-nil, zero value otherwise.

### GetSeverityJustificationOk

`func (o *ReportOrgReportDetail) GetSeverityJustificationOk() (*string, bool)`

GetSeverityJustificationOk returns a tuple with the SeverityJustification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityJustification

`func (o *ReportOrgReportDetail) SetSeverityJustification(v string)`

SetSeverityJustification sets SeverityJustification field to given value.

### HasSeverityJustification

`func (o *ReportOrgReportDetail) HasSeverityJustification() bool`

HasSeverityJustification returns a boolean if a field has been set.

### GetStatusHistory

`func (o *ReportOrgReportDetail) GetStatusHistory() []ReportStatusEvent`

GetStatusHistory returns the StatusHistory field if non-nil, zero value otherwise.

### GetStatusHistoryOk

`func (o *ReportOrgReportDetail) GetStatusHistoryOk() (*[]ReportStatusEvent, bool)`

GetStatusHistoryOk returns a tuple with the StatusHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusHistory

`func (o *ReportOrgReportDetail) SetStatusHistory(v []ReportStatusEvent)`

SetStatusHistory sets StatusHistory field to given value.


### GetSteps

`func (o *ReportOrgReportDetail) GetSteps() string`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *ReportOrgReportDetail) GetStepsOk() (*string, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *ReportOrgReportDetail) SetSteps(v string)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *ReportOrgReportDetail) HasSteps() bool`

HasSteps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


