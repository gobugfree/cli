# ReportFinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetIdentifier** | Pointer to **string** |  | [optional] 
**Cwe** | Pointer to **string** |  | [optional] 
**Decided** | Pointer to [**ReportAssessment**](ReportAssessment.md) |  | [optional] 
**Feedback** | Pointer to **string** |  | [optional] 
**ScopeEntryId** | Pointer to **int32** |  | [optional] 
**Submitted** | [**ReportAssessment**](ReportAssessment.md) |  | 

## Methods

### NewReportFinding

`func NewReportFinding(submitted ReportAssessment, ) *ReportFinding`

NewReportFinding instantiates a new ReportFinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportFindingWithDefaults

`func NewReportFindingWithDefaults() *ReportFinding`

NewReportFindingWithDefaults instantiates a new ReportFinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetIdentifier

`func (o *ReportFinding) GetAssetIdentifier() string`

GetAssetIdentifier returns the AssetIdentifier field if non-nil, zero value otherwise.

### GetAssetIdentifierOk

`func (o *ReportFinding) GetAssetIdentifierOk() (*string, bool)`

GetAssetIdentifierOk returns a tuple with the AssetIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetIdentifier

`func (o *ReportFinding) SetAssetIdentifier(v string)`

SetAssetIdentifier sets AssetIdentifier field to given value.

### HasAssetIdentifier

`func (o *ReportFinding) HasAssetIdentifier() bool`

HasAssetIdentifier returns a boolean if a field has been set.

### GetCwe

`func (o *ReportFinding) GetCwe() string`

GetCwe returns the Cwe field if non-nil, zero value otherwise.

### GetCweOk

`func (o *ReportFinding) GetCweOk() (*string, bool)`

GetCweOk returns a tuple with the Cwe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCwe

`func (o *ReportFinding) SetCwe(v string)`

SetCwe sets Cwe field to given value.

### HasCwe

`func (o *ReportFinding) HasCwe() bool`

HasCwe returns a boolean if a field has been set.

### GetDecided

`func (o *ReportFinding) GetDecided() ReportAssessment`

GetDecided returns the Decided field if non-nil, zero value otherwise.

### GetDecidedOk

`func (o *ReportFinding) GetDecidedOk() (*ReportAssessment, bool)`

GetDecidedOk returns a tuple with the Decided field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecided

`func (o *ReportFinding) SetDecided(v ReportAssessment)`

SetDecided sets Decided field to given value.

### HasDecided

`func (o *ReportFinding) HasDecided() bool`

HasDecided returns a boolean if a field has been set.

### GetFeedback

`func (o *ReportFinding) GetFeedback() string`

GetFeedback returns the Feedback field if non-nil, zero value otherwise.

### GetFeedbackOk

`func (o *ReportFinding) GetFeedbackOk() (*string, bool)`

GetFeedbackOk returns a tuple with the Feedback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedback

`func (o *ReportFinding) SetFeedback(v string)`

SetFeedback sets Feedback field to given value.

### HasFeedback

`func (o *ReportFinding) HasFeedback() bool`

HasFeedback returns a boolean if a field has been set.

### GetScopeEntryId

`func (o *ReportFinding) GetScopeEntryId() int32`

GetScopeEntryId returns the ScopeEntryId field if non-nil, zero value otherwise.

### GetScopeEntryIdOk

`func (o *ReportFinding) GetScopeEntryIdOk() (*int32, bool)`

GetScopeEntryIdOk returns a tuple with the ScopeEntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeEntryId

`func (o *ReportFinding) SetScopeEntryId(v int32)`

SetScopeEntryId sets ScopeEntryId field to given value.

### HasScopeEntryId

`func (o *ReportFinding) HasScopeEntryId() bool`

HasScopeEntryId returns a boolean if a field has been set.

### GetSubmitted

`func (o *ReportFinding) GetSubmitted() ReportAssessment`

GetSubmitted returns the Submitted field if non-nil, zero value otherwise.

### GetSubmittedOk

`func (o *ReportFinding) GetSubmittedOk() (*ReportAssessment, bool)`

GetSubmittedOk returns a tuple with the Submitted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmitted

`func (o *ReportFinding) SetSubmitted(v ReportAssessment)`

SetSubmitted sets Submitted field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


