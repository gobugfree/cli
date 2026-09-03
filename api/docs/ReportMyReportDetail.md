# ReportMyReportDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetIdentifier** | Pointer to **string** |  | [optional] 
**Comments** | [**[]ReportComment**](ReportComment.md) |  | 
**Contributors** | Pointer to [**[]ReportContributor**](ReportContributor.md) |  | [optional] 
**CvssVector** | Pointer to **string** |  | [optional] 
**Cwe** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Evidence** | Pointer to [**[]ReportEvidence**](ReportEvidence.md) |  | [optional] 
**FirstRespondedAt** | Pointer to **time.Time** |  | [optional] 
**FirstResponseDueAt** | Pointer to **time.Time** |  | [optional] 
**Impact** | Pointer to **string** |  | [optional] 
**InfoRequestedAt** | Pointer to **time.Time** |  | [optional] 
**Poc** | Pointer to **string** |  | [optional] 
**RejectionReason** | Pointer to **string** |  | [optional] 
**ResolutionDueAt** | Pointer to **time.Time** |  | [optional] 
**ResolvedAt** | Pointer to **time.Time** |  | [optional] 
**Retest** | Pointer to [**ReportRetest**](ReportRetest.md) |  | [optional] 
**RetestHistory** | Pointer to [**[]ReportRetest**](ReportRetest.md) |  | [optional] 
**SeverityJustification** | Pointer to **string** |  | [optional] 
**SlaSource** | Pointer to **string** |  | [optional] 
**SlaStatus** | Pointer to **string** |  | [optional] 
**StatusHistory** | [**[]ReportStatusEvent**](ReportStatusEvent.md) |  | 
**Steps** | Pointer to **string** |  | [optional] 

## Methods

### NewReportMyReportDetail

`func NewReportMyReportDetail(comments []ReportComment, statusHistory []ReportStatusEvent, ) *ReportMyReportDetail`

NewReportMyReportDetail instantiates a new ReportMyReportDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportMyReportDetailWithDefaults

`func NewReportMyReportDetailWithDefaults() *ReportMyReportDetail`

NewReportMyReportDetailWithDefaults instantiates a new ReportMyReportDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetIdentifier

`func (o *ReportMyReportDetail) GetAssetIdentifier() string`

GetAssetIdentifier returns the AssetIdentifier field if non-nil, zero value otherwise.

### GetAssetIdentifierOk

`func (o *ReportMyReportDetail) GetAssetIdentifierOk() (*string, bool)`

GetAssetIdentifierOk returns a tuple with the AssetIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetIdentifier

`func (o *ReportMyReportDetail) SetAssetIdentifier(v string)`

SetAssetIdentifier sets AssetIdentifier field to given value.

### HasAssetIdentifier

`func (o *ReportMyReportDetail) HasAssetIdentifier() bool`

HasAssetIdentifier returns a boolean if a field has been set.

### GetComments

`func (o *ReportMyReportDetail) GetComments() []ReportComment`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *ReportMyReportDetail) GetCommentsOk() (*[]ReportComment, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *ReportMyReportDetail) SetComments(v []ReportComment)`

SetComments sets Comments field to given value.


### GetContributors

`func (o *ReportMyReportDetail) GetContributors() []ReportContributor`

GetContributors returns the Contributors field if non-nil, zero value otherwise.

### GetContributorsOk

`func (o *ReportMyReportDetail) GetContributorsOk() (*[]ReportContributor, bool)`

GetContributorsOk returns a tuple with the Contributors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributors

`func (o *ReportMyReportDetail) SetContributors(v []ReportContributor)`

SetContributors sets Contributors field to given value.

### HasContributors

`func (o *ReportMyReportDetail) HasContributors() bool`

HasContributors returns a boolean if a field has been set.

### GetCvssVector

`func (o *ReportMyReportDetail) GetCvssVector() string`

GetCvssVector returns the CvssVector field if non-nil, zero value otherwise.

### GetCvssVectorOk

`func (o *ReportMyReportDetail) GetCvssVectorOk() (*string, bool)`

GetCvssVectorOk returns a tuple with the CvssVector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssVector

`func (o *ReportMyReportDetail) SetCvssVector(v string)`

SetCvssVector sets CvssVector field to given value.

### HasCvssVector

`func (o *ReportMyReportDetail) HasCvssVector() bool`

HasCvssVector returns a boolean if a field has been set.

### GetCwe

`func (o *ReportMyReportDetail) GetCwe() string`

GetCwe returns the Cwe field if non-nil, zero value otherwise.

### GetCweOk

`func (o *ReportMyReportDetail) GetCweOk() (*string, bool)`

GetCweOk returns a tuple with the Cwe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCwe

`func (o *ReportMyReportDetail) SetCwe(v string)`

SetCwe sets Cwe field to given value.

### HasCwe

`func (o *ReportMyReportDetail) HasCwe() bool`

HasCwe returns a boolean if a field has been set.

### GetDescription

`func (o *ReportMyReportDetail) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ReportMyReportDetail) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ReportMyReportDetail) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ReportMyReportDetail) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEvidence

`func (o *ReportMyReportDetail) GetEvidence() []ReportEvidence`

GetEvidence returns the Evidence field if non-nil, zero value otherwise.

### GetEvidenceOk

`func (o *ReportMyReportDetail) GetEvidenceOk() (*[]ReportEvidence, bool)`

GetEvidenceOk returns a tuple with the Evidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidence

`func (o *ReportMyReportDetail) SetEvidence(v []ReportEvidence)`

SetEvidence sets Evidence field to given value.

### HasEvidence

`func (o *ReportMyReportDetail) HasEvidence() bool`

HasEvidence returns a boolean if a field has been set.

### GetFirstRespondedAt

`func (o *ReportMyReportDetail) GetFirstRespondedAt() time.Time`

GetFirstRespondedAt returns the FirstRespondedAt field if non-nil, zero value otherwise.

### GetFirstRespondedAtOk

`func (o *ReportMyReportDetail) GetFirstRespondedAtOk() (*time.Time, bool)`

GetFirstRespondedAtOk returns a tuple with the FirstRespondedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstRespondedAt

`func (o *ReportMyReportDetail) SetFirstRespondedAt(v time.Time)`

SetFirstRespondedAt sets FirstRespondedAt field to given value.

### HasFirstRespondedAt

`func (o *ReportMyReportDetail) HasFirstRespondedAt() bool`

HasFirstRespondedAt returns a boolean if a field has been set.

### GetFirstResponseDueAt

`func (o *ReportMyReportDetail) GetFirstResponseDueAt() time.Time`

GetFirstResponseDueAt returns the FirstResponseDueAt field if non-nil, zero value otherwise.

### GetFirstResponseDueAtOk

`func (o *ReportMyReportDetail) GetFirstResponseDueAtOk() (*time.Time, bool)`

GetFirstResponseDueAtOk returns a tuple with the FirstResponseDueAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstResponseDueAt

`func (o *ReportMyReportDetail) SetFirstResponseDueAt(v time.Time)`

SetFirstResponseDueAt sets FirstResponseDueAt field to given value.

### HasFirstResponseDueAt

`func (o *ReportMyReportDetail) HasFirstResponseDueAt() bool`

HasFirstResponseDueAt returns a boolean if a field has been set.

### GetImpact

`func (o *ReportMyReportDetail) GetImpact() string`

GetImpact returns the Impact field if non-nil, zero value otherwise.

### GetImpactOk

`func (o *ReportMyReportDetail) GetImpactOk() (*string, bool)`

GetImpactOk returns a tuple with the Impact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpact

`func (o *ReportMyReportDetail) SetImpact(v string)`

SetImpact sets Impact field to given value.

### HasImpact

`func (o *ReportMyReportDetail) HasImpact() bool`

HasImpact returns a boolean if a field has been set.

### GetInfoRequestedAt

`func (o *ReportMyReportDetail) GetInfoRequestedAt() time.Time`

GetInfoRequestedAt returns the InfoRequestedAt field if non-nil, zero value otherwise.

### GetInfoRequestedAtOk

`func (o *ReportMyReportDetail) GetInfoRequestedAtOk() (*time.Time, bool)`

GetInfoRequestedAtOk returns a tuple with the InfoRequestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfoRequestedAt

`func (o *ReportMyReportDetail) SetInfoRequestedAt(v time.Time)`

SetInfoRequestedAt sets InfoRequestedAt field to given value.

### HasInfoRequestedAt

`func (o *ReportMyReportDetail) HasInfoRequestedAt() bool`

HasInfoRequestedAt returns a boolean if a field has been set.

### GetPoc

`func (o *ReportMyReportDetail) GetPoc() string`

GetPoc returns the Poc field if non-nil, zero value otherwise.

### GetPocOk

`func (o *ReportMyReportDetail) GetPocOk() (*string, bool)`

GetPocOk returns a tuple with the Poc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoc

`func (o *ReportMyReportDetail) SetPoc(v string)`

SetPoc sets Poc field to given value.

### HasPoc

`func (o *ReportMyReportDetail) HasPoc() bool`

HasPoc returns a boolean if a field has been set.

### GetRejectionReason

`func (o *ReportMyReportDetail) GetRejectionReason() string`

GetRejectionReason returns the RejectionReason field if non-nil, zero value otherwise.

### GetRejectionReasonOk

`func (o *ReportMyReportDetail) GetRejectionReasonOk() (*string, bool)`

GetRejectionReasonOk returns a tuple with the RejectionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectionReason

`func (o *ReportMyReportDetail) SetRejectionReason(v string)`

SetRejectionReason sets RejectionReason field to given value.

### HasRejectionReason

`func (o *ReportMyReportDetail) HasRejectionReason() bool`

HasRejectionReason returns a boolean if a field has been set.

### GetResolutionDueAt

`func (o *ReportMyReportDetail) GetResolutionDueAt() time.Time`

GetResolutionDueAt returns the ResolutionDueAt field if non-nil, zero value otherwise.

### GetResolutionDueAtOk

`func (o *ReportMyReportDetail) GetResolutionDueAtOk() (*time.Time, bool)`

GetResolutionDueAtOk returns a tuple with the ResolutionDueAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionDueAt

`func (o *ReportMyReportDetail) SetResolutionDueAt(v time.Time)`

SetResolutionDueAt sets ResolutionDueAt field to given value.

### HasResolutionDueAt

`func (o *ReportMyReportDetail) HasResolutionDueAt() bool`

HasResolutionDueAt returns a boolean if a field has been set.

### GetResolvedAt

`func (o *ReportMyReportDetail) GetResolvedAt() time.Time`

GetResolvedAt returns the ResolvedAt field if non-nil, zero value otherwise.

### GetResolvedAtOk

`func (o *ReportMyReportDetail) GetResolvedAtOk() (*time.Time, bool)`

GetResolvedAtOk returns a tuple with the ResolvedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedAt

`func (o *ReportMyReportDetail) SetResolvedAt(v time.Time)`

SetResolvedAt sets ResolvedAt field to given value.

### HasResolvedAt

`func (o *ReportMyReportDetail) HasResolvedAt() bool`

HasResolvedAt returns a boolean if a field has been set.

### GetRetest

`func (o *ReportMyReportDetail) GetRetest() ReportRetest`

GetRetest returns the Retest field if non-nil, zero value otherwise.

### GetRetestOk

`func (o *ReportMyReportDetail) GetRetestOk() (*ReportRetest, bool)`

GetRetestOk returns a tuple with the Retest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetest

`func (o *ReportMyReportDetail) SetRetest(v ReportRetest)`

SetRetest sets Retest field to given value.

### HasRetest

`func (o *ReportMyReportDetail) HasRetest() bool`

HasRetest returns a boolean if a field has been set.

### GetRetestHistory

`func (o *ReportMyReportDetail) GetRetestHistory() []ReportRetest`

GetRetestHistory returns the RetestHistory field if non-nil, zero value otherwise.

### GetRetestHistoryOk

`func (o *ReportMyReportDetail) GetRetestHistoryOk() (*[]ReportRetest, bool)`

GetRetestHistoryOk returns a tuple with the RetestHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetestHistory

`func (o *ReportMyReportDetail) SetRetestHistory(v []ReportRetest)`

SetRetestHistory sets RetestHistory field to given value.

### HasRetestHistory

`func (o *ReportMyReportDetail) HasRetestHistory() bool`

HasRetestHistory returns a boolean if a field has been set.

### GetSeverityJustification

`func (o *ReportMyReportDetail) GetSeverityJustification() string`

GetSeverityJustification returns the SeverityJustification field if non-nil, zero value otherwise.

### GetSeverityJustificationOk

`func (o *ReportMyReportDetail) GetSeverityJustificationOk() (*string, bool)`

GetSeverityJustificationOk returns a tuple with the SeverityJustification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityJustification

`func (o *ReportMyReportDetail) SetSeverityJustification(v string)`

SetSeverityJustification sets SeverityJustification field to given value.

### HasSeverityJustification

`func (o *ReportMyReportDetail) HasSeverityJustification() bool`

HasSeverityJustification returns a boolean if a field has been set.

### GetSlaSource

`func (o *ReportMyReportDetail) GetSlaSource() string`

GetSlaSource returns the SlaSource field if non-nil, zero value otherwise.

### GetSlaSourceOk

`func (o *ReportMyReportDetail) GetSlaSourceOk() (*string, bool)`

GetSlaSourceOk returns a tuple with the SlaSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaSource

`func (o *ReportMyReportDetail) SetSlaSource(v string)`

SetSlaSource sets SlaSource field to given value.

### HasSlaSource

`func (o *ReportMyReportDetail) HasSlaSource() bool`

HasSlaSource returns a boolean if a field has been set.

### GetSlaStatus

`func (o *ReportMyReportDetail) GetSlaStatus() string`

GetSlaStatus returns the SlaStatus field if non-nil, zero value otherwise.

### GetSlaStatusOk

`func (o *ReportMyReportDetail) GetSlaStatusOk() (*string, bool)`

GetSlaStatusOk returns a tuple with the SlaStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaStatus

`func (o *ReportMyReportDetail) SetSlaStatus(v string)`

SetSlaStatus sets SlaStatus field to given value.

### HasSlaStatus

`func (o *ReportMyReportDetail) HasSlaStatus() bool`

HasSlaStatus returns a boolean if a field has been set.

### GetStatusHistory

`func (o *ReportMyReportDetail) GetStatusHistory() []ReportStatusEvent`

GetStatusHistory returns the StatusHistory field if non-nil, zero value otherwise.

### GetStatusHistoryOk

`func (o *ReportMyReportDetail) GetStatusHistoryOk() (*[]ReportStatusEvent, bool)`

GetStatusHistoryOk returns a tuple with the StatusHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusHistory

`func (o *ReportMyReportDetail) SetStatusHistory(v []ReportStatusEvent)`

SetStatusHistory sets StatusHistory field to given value.


### GetSteps

`func (o *ReportMyReportDetail) GetSteps() string`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *ReportMyReportDetail) GetStepsOk() (*string, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *ReportMyReportDetail) SetSteps(v string)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *ReportMyReportDetail) HasSteps() bool`

HasSteps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


