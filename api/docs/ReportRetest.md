# ReportRetest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DecidedAt** | Pointer to **time.Time** |  | [optional] 
**Due** | Pointer to **time.Time** |  | [optional] 
**Id** | **string** |  | 
**OutcomeNotes** | Pointer to **string** |  | [optional] 
**RemediationNotes** | **string** |  | 
**RequestedAt** | **time.Time** |  | 
**Status** | **string** |  | 

## Methods

### NewReportRetest

`func NewReportRetest(id string, remediationNotes string, requestedAt time.Time, status string, ) *ReportRetest`

NewReportRetest instantiates a new ReportRetest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportRetestWithDefaults

`func NewReportRetestWithDefaults() *ReportRetest`

NewReportRetestWithDefaults instantiates a new ReportRetest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDecidedAt

`func (o *ReportRetest) GetDecidedAt() time.Time`

GetDecidedAt returns the DecidedAt field if non-nil, zero value otherwise.

### GetDecidedAtOk

`func (o *ReportRetest) GetDecidedAtOk() (*time.Time, bool)`

GetDecidedAtOk returns a tuple with the DecidedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecidedAt

`func (o *ReportRetest) SetDecidedAt(v time.Time)`

SetDecidedAt sets DecidedAt field to given value.

### HasDecidedAt

`func (o *ReportRetest) HasDecidedAt() bool`

HasDecidedAt returns a boolean if a field has been set.

### GetDue

`func (o *ReportRetest) GetDue() time.Time`

GetDue returns the Due field if non-nil, zero value otherwise.

### GetDueOk

`func (o *ReportRetest) GetDueOk() (*time.Time, bool)`

GetDueOk returns a tuple with the Due field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDue

`func (o *ReportRetest) SetDue(v time.Time)`

SetDue sets Due field to given value.

### HasDue

`func (o *ReportRetest) HasDue() bool`

HasDue returns a boolean if a field has been set.

### GetId

`func (o *ReportRetest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReportRetest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReportRetest) SetId(v string)`

SetId sets Id field to given value.


### GetOutcomeNotes

`func (o *ReportRetest) GetOutcomeNotes() string`

GetOutcomeNotes returns the OutcomeNotes field if non-nil, zero value otherwise.

### GetOutcomeNotesOk

`func (o *ReportRetest) GetOutcomeNotesOk() (*string, bool)`

GetOutcomeNotesOk returns a tuple with the OutcomeNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcomeNotes

`func (o *ReportRetest) SetOutcomeNotes(v string)`

SetOutcomeNotes sets OutcomeNotes field to given value.

### HasOutcomeNotes

`func (o *ReportRetest) HasOutcomeNotes() bool`

HasOutcomeNotes returns a boolean if a field has been set.

### GetRemediationNotes

`func (o *ReportRetest) GetRemediationNotes() string`

GetRemediationNotes returns the RemediationNotes field if non-nil, zero value otherwise.

### GetRemediationNotesOk

`func (o *ReportRetest) GetRemediationNotesOk() (*string, bool)`

GetRemediationNotesOk returns a tuple with the RemediationNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediationNotes

`func (o *ReportRetest) SetRemediationNotes(v string)`

SetRemediationNotes sets RemediationNotes field to given value.


### GetRequestedAt

`func (o *ReportRetest) GetRequestedAt() time.Time`

GetRequestedAt returns the RequestedAt field if non-nil, zero value otherwise.

### GetRequestedAtOk

`func (o *ReportRetest) GetRequestedAtOk() (*time.Time, bool)`

GetRequestedAtOk returns a tuple with the RequestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedAt

`func (o *ReportRetest) SetRequestedAt(v time.Time)`

SetRequestedAt sets RequestedAt field to given value.


### GetStatus

`func (o *ReportRetest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ReportRetest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ReportRetest) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


