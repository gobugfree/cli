# Retest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DecidedAt** | Pointer to **time.Time** |  | [optional] 
**Due** | Pointer to **time.Time** |  | [optional] 
**Engagement** | Pointer to **string** |  | [optional] 
**Id** | **string** |  | 
**OutcomeNotes** | Pointer to **string** |  | [optional] 
**RemediationNotes** | **string** |  | 
**ReportId** | **string** |  | 
**ReportTitle** | Pointer to **string** |  | [optional] 
**RequestedAt** | **time.Time** |  | 
**Status** | **string** |  | 

## Methods

### NewRetest

`func NewRetest(id string, remediationNotes string, reportId string, requestedAt time.Time, status string, ) *Retest`

NewRetest instantiates a new Retest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRetestWithDefaults

`func NewRetestWithDefaults() *Retest`

NewRetestWithDefaults instantiates a new Retest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDecidedAt

`func (o *Retest) GetDecidedAt() time.Time`

GetDecidedAt returns the DecidedAt field if non-nil, zero value otherwise.

### GetDecidedAtOk

`func (o *Retest) GetDecidedAtOk() (*time.Time, bool)`

GetDecidedAtOk returns a tuple with the DecidedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecidedAt

`func (o *Retest) SetDecidedAt(v time.Time)`

SetDecidedAt sets DecidedAt field to given value.

### HasDecidedAt

`func (o *Retest) HasDecidedAt() bool`

HasDecidedAt returns a boolean if a field has been set.

### GetDue

`func (o *Retest) GetDue() time.Time`

GetDue returns the Due field if non-nil, zero value otherwise.

### GetDueOk

`func (o *Retest) GetDueOk() (*time.Time, bool)`

GetDueOk returns a tuple with the Due field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDue

`func (o *Retest) SetDue(v time.Time)`

SetDue sets Due field to given value.

### HasDue

`func (o *Retest) HasDue() bool`

HasDue returns a boolean if a field has been set.

### GetEngagement

`func (o *Retest) GetEngagement() string`

GetEngagement returns the Engagement field if non-nil, zero value otherwise.

### GetEngagementOk

`func (o *Retest) GetEngagementOk() (*string, bool)`

GetEngagementOk returns a tuple with the Engagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngagement

`func (o *Retest) SetEngagement(v string)`

SetEngagement sets Engagement field to given value.

### HasEngagement

`func (o *Retest) HasEngagement() bool`

HasEngagement returns a boolean if a field has been set.

### GetId

`func (o *Retest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Retest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Retest) SetId(v string)`

SetId sets Id field to given value.


### GetOutcomeNotes

`func (o *Retest) GetOutcomeNotes() string`

GetOutcomeNotes returns the OutcomeNotes field if non-nil, zero value otherwise.

### GetOutcomeNotesOk

`func (o *Retest) GetOutcomeNotesOk() (*string, bool)`

GetOutcomeNotesOk returns a tuple with the OutcomeNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcomeNotes

`func (o *Retest) SetOutcomeNotes(v string)`

SetOutcomeNotes sets OutcomeNotes field to given value.

### HasOutcomeNotes

`func (o *Retest) HasOutcomeNotes() bool`

HasOutcomeNotes returns a boolean if a field has been set.

### GetRemediationNotes

`func (o *Retest) GetRemediationNotes() string`

GetRemediationNotes returns the RemediationNotes field if non-nil, zero value otherwise.

### GetRemediationNotesOk

`func (o *Retest) GetRemediationNotesOk() (*string, bool)`

GetRemediationNotesOk returns a tuple with the RemediationNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediationNotes

`func (o *Retest) SetRemediationNotes(v string)`

SetRemediationNotes sets RemediationNotes field to given value.


### GetReportId

`func (o *Retest) GetReportId() string`

GetReportId returns the ReportId field if non-nil, zero value otherwise.

### GetReportIdOk

`func (o *Retest) GetReportIdOk() (*string, bool)`

GetReportIdOk returns a tuple with the ReportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportId

`func (o *Retest) SetReportId(v string)`

SetReportId sets ReportId field to given value.


### GetReportTitle

`func (o *Retest) GetReportTitle() string`

GetReportTitle returns the ReportTitle field if non-nil, zero value otherwise.

### GetReportTitleOk

`func (o *Retest) GetReportTitleOk() (*string, bool)`

GetReportTitleOk returns a tuple with the ReportTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportTitle

`func (o *Retest) SetReportTitle(v string)`

SetReportTitle sets ReportTitle field to given value.

### HasReportTitle

`func (o *Retest) HasReportTitle() bool`

HasReportTitle returns a boolean if a field has been set.

### GetRequestedAt

`func (o *Retest) GetRequestedAt() time.Time`

GetRequestedAt returns the RequestedAt field if non-nil, zero value otherwise.

### GetRequestedAtOk

`func (o *Retest) GetRequestedAtOk() (*time.Time, bool)`

GetRequestedAtOk returns a tuple with the RequestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedAt

`func (o *Retest) SetRequestedAt(v time.Time)`

SetRequestedAt sets RequestedAt field to given value.


### GetStatus

`func (o *Retest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Retest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Retest) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


