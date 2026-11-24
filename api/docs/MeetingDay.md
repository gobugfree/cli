# MeetingDay

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **string** |  | 
**Slots** | [**[]time.Time**](time.Time.md) |  | 

## Methods

### NewMeetingDay

`func NewMeetingDay(date string, slots []time.Time, ) *MeetingDay`

NewMeetingDay instantiates a new MeetingDay object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeetingDayWithDefaults

`func NewMeetingDayWithDefaults() *MeetingDay`

NewMeetingDayWithDefaults instantiates a new MeetingDay object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *MeetingDay) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *MeetingDay) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *MeetingDay) SetDate(v string)`

SetDate sets Date field to given value.


### GetSlots

`func (o *MeetingDay) GetSlots() []time.Time`

GetSlots returns the Slots field if non-nil, zero value otherwise.

### GetSlotsOk

`func (o *MeetingDay) GetSlotsOk() (*[]time.Time, bool)`

GetSlotsOk returns a tuple with the Slots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlots

`func (o *MeetingDay) SetSlots(v []time.Time)`

SetSlots sets Slots field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


