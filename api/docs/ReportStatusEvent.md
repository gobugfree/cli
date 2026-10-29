# ReportStatusEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**At** | **time.Time** |  | 
**From** | **string** |  | 
**To** | **string** |  | 

## Methods

### NewReportStatusEvent

`func NewReportStatusEvent(at time.Time, from string, to string, ) *ReportStatusEvent`

NewReportStatusEvent instantiates a new ReportStatusEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportStatusEventWithDefaults

`func NewReportStatusEventWithDefaults() *ReportStatusEvent`

NewReportStatusEventWithDefaults instantiates a new ReportStatusEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAt

`func (o *ReportStatusEvent) GetAt() time.Time`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *ReportStatusEvent) GetAtOk() (*time.Time, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *ReportStatusEvent) SetAt(v time.Time)`

SetAt sets At field to given value.


### GetFrom

`func (o *ReportStatusEvent) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *ReportStatusEvent) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *ReportStatusEvent) SetFrom(v string)`

SetFrom sets From field to given value.


### GetTo

`func (o *ReportStatusEvent) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *ReportStatusEvent) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *ReportStatusEvent) SetTo(v string)`

SetTo sets To field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


