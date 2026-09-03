# Disclosure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**PlannedOn** | Pointer to **time.Time** |  | [optional] 
**PublishedOn** | Pointer to **time.Time** |  | [optional] 
**ReportId** | **string** |  | 
**Status** | **string** |  | 
**Summary** | Pointer to **string** |  | [optional] 

## Methods

### NewDisclosure

`func NewDisclosure(id string, reportId string, status string, ) *Disclosure`

NewDisclosure instantiates a new Disclosure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDisclosureWithDefaults

`func NewDisclosureWithDefaults() *Disclosure`

NewDisclosureWithDefaults instantiates a new Disclosure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Disclosure) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Disclosure) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Disclosure) SetId(v string)`

SetId sets Id field to given value.


### GetPlannedOn

`func (o *Disclosure) GetPlannedOn() time.Time`

GetPlannedOn returns the PlannedOn field if non-nil, zero value otherwise.

### GetPlannedOnOk

`func (o *Disclosure) GetPlannedOnOk() (*time.Time, bool)`

GetPlannedOnOk returns a tuple with the PlannedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedOn

`func (o *Disclosure) SetPlannedOn(v time.Time)`

SetPlannedOn sets PlannedOn field to given value.

### HasPlannedOn

`func (o *Disclosure) HasPlannedOn() bool`

HasPlannedOn returns a boolean if a field has been set.

### GetPublishedOn

`func (o *Disclosure) GetPublishedOn() time.Time`

GetPublishedOn returns the PublishedOn field if non-nil, zero value otherwise.

### GetPublishedOnOk

`func (o *Disclosure) GetPublishedOnOk() (*time.Time, bool)`

GetPublishedOnOk returns a tuple with the PublishedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedOn

`func (o *Disclosure) SetPublishedOn(v time.Time)`

SetPublishedOn sets PublishedOn field to given value.

### HasPublishedOn

`func (o *Disclosure) HasPublishedOn() bool`

HasPublishedOn returns a boolean if a field has been set.

### GetReportId

`func (o *Disclosure) GetReportId() string`

GetReportId returns the ReportId field if non-nil, zero value otherwise.

### GetReportIdOk

`func (o *Disclosure) GetReportIdOk() (*string, bool)`

GetReportIdOk returns a tuple with the ReportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportId

`func (o *Disclosure) SetReportId(v string)`

SetReportId sets ReportId field to given value.


### GetStatus

`func (o *Disclosure) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Disclosure) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Disclosure) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSummary

`func (o *Disclosure) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *Disclosure) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *Disclosure) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *Disclosure) HasSummary() bool`

HasSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


