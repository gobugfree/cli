# ReportOrgReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetIdentifier** | Pointer to **string** |  | [optional] 
**Assignee** | Pointer to [**ReportAssignee**](ReportAssignee.md) |  | [optional] 
**CvssScore** | Pointer to **float32** |  | [optional] 
**EngagementSlug** | Pointer to **string** |  | [optional] 
**EngagementTitle** | Pointer to **string** |  | [optional] 
**Id** | **string** |  | 
**Severity** | Pointer to **string** |  | [optional] 
**SeveritySubmitted** | Pointer to **string** |  | [optional] 
**Status** | **string** |  | 
**SubmittedAt** | Pointer to **time.Time** |  | [optional] 
**Title** | **string** |  | 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewReportOrgReport

`func NewReportOrgReport(id string, status string, title string, ) *ReportOrgReport`

NewReportOrgReport instantiates a new ReportOrgReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportOrgReportWithDefaults

`func NewReportOrgReportWithDefaults() *ReportOrgReport`

NewReportOrgReportWithDefaults instantiates a new ReportOrgReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetIdentifier

`func (o *ReportOrgReport) GetAssetIdentifier() string`

GetAssetIdentifier returns the AssetIdentifier field if non-nil, zero value otherwise.

### GetAssetIdentifierOk

`func (o *ReportOrgReport) GetAssetIdentifierOk() (*string, bool)`

GetAssetIdentifierOk returns a tuple with the AssetIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetIdentifier

`func (o *ReportOrgReport) SetAssetIdentifier(v string)`

SetAssetIdentifier sets AssetIdentifier field to given value.

### HasAssetIdentifier

`func (o *ReportOrgReport) HasAssetIdentifier() bool`

HasAssetIdentifier returns a boolean if a field has been set.

### GetAssignee

`func (o *ReportOrgReport) GetAssignee() ReportAssignee`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *ReportOrgReport) GetAssigneeOk() (*ReportAssignee, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *ReportOrgReport) SetAssignee(v ReportAssignee)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *ReportOrgReport) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### GetCvssScore

`func (o *ReportOrgReport) GetCvssScore() float32`

GetCvssScore returns the CvssScore field if non-nil, zero value otherwise.

### GetCvssScoreOk

`func (o *ReportOrgReport) GetCvssScoreOk() (*float32, bool)`

GetCvssScoreOk returns a tuple with the CvssScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssScore

`func (o *ReportOrgReport) SetCvssScore(v float32)`

SetCvssScore sets CvssScore field to given value.

### HasCvssScore

`func (o *ReportOrgReport) HasCvssScore() bool`

HasCvssScore returns a boolean if a field has been set.

### GetEngagementSlug

`func (o *ReportOrgReport) GetEngagementSlug() string`

GetEngagementSlug returns the EngagementSlug field if non-nil, zero value otherwise.

### GetEngagementSlugOk

`func (o *ReportOrgReport) GetEngagementSlugOk() (*string, bool)`

GetEngagementSlugOk returns a tuple with the EngagementSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngagementSlug

`func (o *ReportOrgReport) SetEngagementSlug(v string)`

SetEngagementSlug sets EngagementSlug field to given value.

### HasEngagementSlug

`func (o *ReportOrgReport) HasEngagementSlug() bool`

HasEngagementSlug returns a boolean if a field has been set.

### GetEngagementTitle

`func (o *ReportOrgReport) GetEngagementTitle() string`

GetEngagementTitle returns the EngagementTitle field if non-nil, zero value otherwise.

### GetEngagementTitleOk

`func (o *ReportOrgReport) GetEngagementTitleOk() (*string, bool)`

GetEngagementTitleOk returns a tuple with the EngagementTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngagementTitle

`func (o *ReportOrgReport) SetEngagementTitle(v string)`

SetEngagementTitle sets EngagementTitle field to given value.

### HasEngagementTitle

`func (o *ReportOrgReport) HasEngagementTitle() bool`

HasEngagementTitle returns a boolean if a field has been set.

### GetId

`func (o *ReportOrgReport) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReportOrgReport) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReportOrgReport) SetId(v string)`

SetId sets Id field to given value.


### GetSeverity

`func (o *ReportOrgReport) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ReportOrgReport) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ReportOrgReport) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *ReportOrgReport) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetSeveritySubmitted

`func (o *ReportOrgReport) GetSeveritySubmitted() string`

GetSeveritySubmitted returns the SeveritySubmitted field if non-nil, zero value otherwise.

### GetSeveritySubmittedOk

`func (o *ReportOrgReport) GetSeveritySubmittedOk() (*string, bool)`

GetSeveritySubmittedOk returns a tuple with the SeveritySubmitted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeveritySubmitted

`func (o *ReportOrgReport) SetSeveritySubmitted(v string)`

SetSeveritySubmitted sets SeveritySubmitted field to given value.

### HasSeveritySubmitted

`func (o *ReportOrgReport) HasSeveritySubmitted() bool`

HasSeveritySubmitted returns a boolean if a field has been set.

### GetStatus

`func (o *ReportOrgReport) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ReportOrgReport) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ReportOrgReport) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSubmittedAt

`func (o *ReportOrgReport) GetSubmittedAt() time.Time`

GetSubmittedAt returns the SubmittedAt field if non-nil, zero value otherwise.

### GetSubmittedAtOk

`func (o *ReportOrgReport) GetSubmittedAtOk() (*time.Time, bool)`

GetSubmittedAtOk returns a tuple with the SubmittedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedAt

`func (o *ReportOrgReport) SetSubmittedAt(v time.Time)`

SetSubmittedAt sets SubmittedAt field to given value.

### HasSubmittedAt

`func (o *ReportOrgReport) HasSubmittedAt() bool`

HasSubmittedAt returns a boolean if a field has been set.

### GetTitle

`func (o *ReportOrgReport) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ReportOrgReport) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ReportOrgReport) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetUpdatedAt

`func (o *ReportOrgReport) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ReportOrgReport) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ReportOrgReport) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ReportOrgReport) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


