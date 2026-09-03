# ReportMyReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CvssScore** | Pointer to **float32** |  | [optional] 
**EngagementSlug** | **string** |  | 
**EngagementTitle** | **string** |  | 
**Id** | **string** |  | 
**InfoRequestedAt** | Pointer to **time.Time** |  | [optional] 
**OrgName** | **string** |  | 
**OrgSlug** | **string** |  | 
**Severity** | Pointer to **string** |  | [optional] 
**SeveritySubmitted** | Pointer to **string** |  | [optional] 
**Status** | **string** |  | 
**SubmittedAt** | Pointer to **time.Time** |  | [optional] 
**Title** | **string** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewReportMyReport

`func NewReportMyReport(engagementSlug string, engagementTitle string, id string, orgName string, orgSlug string, status string, title string, updatedAt time.Time, ) *ReportMyReport`

NewReportMyReport instantiates a new ReportMyReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportMyReportWithDefaults

`func NewReportMyReportWithDefaults() *ReportMyReport`

NewReportMyReportWithDefaults instantiates a new ReportMyReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCvssScore

`func (o *ReportMyReport) GetCvssScore() float32`

GetCvssScore returns the CvssScore field if non-nil, zero value otherwise.

### GetCvssScoreOk

`func (o *ReportMyReport) GetCvssScoreOk() (*float32, bool)`

GetCvssScoreOk returns a tuple with the CvssScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssScore

`func (o *ReportMyReport) SetCvssScore(v float32)`

SetCvssScore sets CvssScore field to given value.

### HasCvssScore

`func (o *ReportMyReport) HasCvssScore() bool`

HasCvssScore returns a boolean if a field has been set.

### GetEngagementSlug

`func (o *ReportMyReport) GetEngagementSlug() string`

GetEngagementSlug returns the EngagementSlug field if non-nil, zero value otherwise.

### GetEngagementSlugOk

`func (o *ReportMyReport) GetEngagementSlugOk() (*string, bool)`

GetEngagementSlugOk returns a tuple with the EngagementSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngagementSlug

`func (o *ReportMyReport) SetEngagementSlug(v string)`

SetEngagementSlug sets EngagementSlug field to given value.


### GetEngagementTitle

`func (o *ReportMyReport) GetEngagementTitle() string`

GetEngagementTitle returns the EngagementTitle field if non-nil, zero value otherwise.

### GetEngagementTitleOk

`func (o *ReportMyReport) GetEngagementTitleOk() (*string, bool)`

GetEngagementTitleOk returns a tuple with the EngagementTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngagementTitle

`func (o *ReportMyReport) SetEngagementTitle(v string)`

SetEngagementTitle sets EngagementTitle field to given value.


### GetId

`func (o *ReportMyReport) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReportMyReport) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReportMyReport) SetId(v string)`

SetId sets Id field to given value.


### GetInfoRequestedAt

`func (o *ReportMyReport) GetInfoRequestedAt() time.Time`

GetInfoRequestedAt returns the InfoRequestedAt field if non-nil, zero value otherwise.

### GetInfoRequestedAtOk

`func (o *ReportMyReport) GetInfoRequestedAtOk() (*time.Time, bool)`

GetInfoRequestedAtOk returns a tuple with the InfoRequestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfoRequestedAt

`func (o *ReportMyReport) SetInfoRequestedAt(v time.Time)`

SetInfoRequestedAt sets InfoRequestedAt field to given value.

### HasInfoRequestedAt

`func (o *ReportMyReport) HasInfoRequestedAt() bool`

HasInfoRequestedAt returns a boolean if a field has been set.

### GetOrgName

`func (o *ReportMyReport) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *ReportMyReport) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *ReportMyReport) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.


### GetOrgSlug

`func (o *ReportMyReport) GetOrgSlug() string`

GetOrgSlug returns the OrgSlug field if non-nil, zero value otherwise.

### GetOrgSlugOk

`func (o *ReportMyReport) GetOrgSlugOk() (*string, bool)`

GetOrgSlugOk returns a tuple with the OrgSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgSlug

`func (o *ReportMyReport) SetOrgSlug(v string)`

SetOrgSlug sets OrgSlug field to given value.


### GetSeverity

`func (o *ReportMyReport) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ReportMyReport) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ReportMyReport) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *ReportMyReport) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetSeveritySubmitted

`func (o *ReportMyReport) GetSeveritySubmitted() string`

GetSeveritySubmitted returns the SeveritySubmitted field if non-nil, zero value otherwise.

### GetSeveritySubmittedOk

`func (o *ReportMyReport) GetSeveritySubmittedOk() (*string, bool)`

GetSeveritySubmittedOk returns a tuple with the SeveritySubmitted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeveritySubmitted

`func (o *ReportMyReport) SetSeveritySubmitted(v string)`

SetSeveritySubmitted sets SeveritySubmitted field to given value.

### HasSeveritySubmitted

`func (o *ReportMyReport) HasSeveritySubmitted() bool`

HasSeveritySubmitted returns a boolean if a field has been set.

### GetStatus

`func (o *ReportMyReport) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ReportMyReport) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ReportMyReport) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSubmittedAt

`func (o *ReportMyReport) GetSubmittedAt() time.Time`

GetSubmittedAt returns the SubmittedAt field if non-nil, zero value otherwise.

### GetSubmittedAtOk

`func (o *ReportMyReport) GetSubmittedAtOk() (*time.Time, bool)`

GetSubmittedAtOk returns a tuple with the SubmittedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedAt

`func (o *ReportMyReport) SetSubmittedAt(v time.Time)`

SetSubmittedAt sets SubmittedAt field to given value.

### HasSubmittedAt

`func (o *ReportMyReport) HasSubmittedAt() bool`

HasSubmittedAt returns a boolean if a field has been set.

### GetTitle

`func (o *ReportMyReport) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ReportMyReport) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ReportMyReport) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetUpdatedAt

`func (o *ReportMyReport) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ReportMyReport) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ReportMyReport) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


