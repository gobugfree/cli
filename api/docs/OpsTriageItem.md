# OpsTriageItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assignee** | Pointer to **string** |  | [optional] 
**CvssScore** | Pointer to **float32** |  | [optional] 
**EngagementSlug** | **string** |  | 
**EngagementTitle** | **string** |  | 
**OrgName** | **string** |  | 
**OrgSlug** | **string** |  | 
**ReportId** | **string** |  | 
**Researcher** | **string** |  | 
**Severity** | Pointer to **string** |  | [optional] 
**SeveritySubmitted** | Pointer to **string** |  | [optional] 
**Status** | **string** |  | 
**SubmittedAt** | Pointer to **time.Time** |  | [optional] 
**Title** | **string** |  | 

## Methods

### NewOpsTriageItem

`func NewOpsTriageItem(engagementSlug string, engagementTitle string, orgName string, orgSlug string, reportId string, researcher string, status string, title string, ) *OpsTriageItem`

NewOpsTriageItem instantiates a new OpsTriageItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpsTriageItemWithDefaults

`func NewOpsTriageItemWithDefaults() *OpsTriageItem`

NewOpsTriageItemWithDefaults instantiates a new OpsTriageItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignee

`func (o *OpsTriageItem) GetAssignee() string`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *OpsTriageItem) GetAssigneeOk() (*string, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *OpsTriageItem) SetAssignee(v string)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *OpsTriageItem) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### GetCvssScore

`func (o *OpsTriageItem) GetCvssScore() float32`

GetCvssScore returns the CvssScore field if non-nil, zero value otherwise.

### GetCvssScoreOk

`func (o *OpsTriageItem) GetCvssScoreOk() (*float32, bool)`

GetCvssScoreOk returns a tuple with the CvssScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssScore

`func (o *OpsTriageItem) SetCvssScore(v float32)`

SetCvssScore sets CvssScore field to given value.

### HasCvssScore

`func (o *OpsTriageItem) HasCvssScore() bool`

HasCvssScore returns a boolean if a field has been set.

### GetEngagementSlug

`func (o *OpsTriageItem) GetEngagementSlug() string`

GetEngagementSlug returns the EngagementSlug field if non-nil, zero value otherwise.

### GetEngagementSlugOk

`func (o *OpsTriageItem) GetEngagementSlugOk() (*string, bool)`

GetEngagementSlugOk returns a tuple with the EngagementSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngagementSlug

`func (o *OpsTriageItem) SetEngagementSlug(v string)`

SetEngagementSlug sets EngagementSlug field to given value.


### GetEngagementTitle

`func (o *OpsTriageItem) GetEngagementTitle() string`

GetEngagementTitle returns the EngagementTitle field if non-nil, zero value otherwise.

### GetEngagementTitleOk

`func (o *OpsTriageItem) GetEngagementTitleOk() (*string, bool)`

GetEngagementTitleOk returns a tuple with the EngagementTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngagementTitle

`func (o *OpsTriageItem) SetEngagementTitle(v string)`

SetEngagementTitle sets EngagementTitle field to given value.


### GetOrgName

`func (o *OpsTriageItem) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *OpsTriageItem) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *OpsTriageItem) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.


### GetOrgSlug

`func (o *OpsTriageItem) GetOrgSlug() string`

GetOrgSlug returns the OrgSlug field if non-nil, zero value otherwise.

### GetOrgSlugOk

`func (o *OpsTriageItem) GetOrgSlugOk() (*string, bool)`

GetOrgSlugOk returns a tuple with the OrgSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgSlug

`func (o *OpsTriageItem) SetOrgSlug(v string)`

SetOrgSlug sets OrgSlug field to given value.


### GetReportId

`func (o *OpsTriageItem) GetReportId() string`

GetReportId returns the ReportId field if non-nil, zero value otherwise.

### GetReportIdOk

`func (o *OpsTriageItem) GetReportIdOk() (*string, bool)`

GetReportIdOk returns a tuple with the ReportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportId

`func (o *OpsTriageItem) SetReportId(v string)`

SetReportId sets ReportId field to given value.


### GetResearcher

`func (o *OpsTriageItem) GetResearcher() string`

GetResearcher returns the Researcher field if non-nil, zero value otherwise.

### GetResearcherOk

`func (o *OpsTriageItem) GetResearcherOk() (*string, bool)`

GetResearcherOk returns a tuple with the Researcher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResearcher

`func (o *OpsTriageItem) SetResearcher(v string)`

SetResearcher sets Researcher field to given value.


### GetSeverity

`func (o *OpsTriageItem) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *OpsTriageItem) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *OpsTriageItem) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *OpsTriageItem) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetSeveritySubmitted

`func (o *OpsTriageItem) GetSeveritySubmitted() string`

GetSeveritySubmitted returns the SeveritySubmitted field if non-nil, zero value otherwise.

### GetSeveritySubmittedOk

`func (o *OpsTriageItem) GetSeveritySubmittedOk() (*string, bool)`

GetSeveritySubmittedOk returns a tuple with the SeveritySubmitted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeveritySubmitted

`func (o *OpsTriageItem) SetSeveritySubmitted(v string)`

SetSeveritySubmitted sets SeveritySubmitted field to given value.

### HasSeveritySubmitted

`func (o *OpsTriageItem) HasSeveritySubmitted() bool`

HasSeveritySubmitted returns a boolean if a field has been set.

### GetStatus

`func (o *OpsTriageItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OpsTriageItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OpsTriageItem) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSubmittedAt

`func (o *OpsTriageItem) GetSubmittedAt() time.Time`

GetSubmittedAt returns the SubmittedAt field if non-nil, zero value otherwise.

### GetSubmittedAtOk

`func (o *OpsTriageItem) GetSubmittedAtOk() (*time.Time, bool)`

GetSubmittedAtOk returns a tuple with the SubmittedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedAt

`func (o *OpsTriageItem) SetSubmittedAt(v time.Time)`

SetSubmittedAt sets SubmittedAt field to given value.

### HasSubmittedAt

`func (o *OpsTriageItem) HasSubmittedAt() bool`

HasSubmittedAt returns a boolean if a field has been set.

### GetTitle

`func (o *OpsTriageItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *OpsTriageItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *OpsTriageItem) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


