# MeetingBooked

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndsAt** | **time.Time** |  | 
**StartsAt** | **time.Time** |  | 

## Methods

### NewMeetingBooked

`func NewMeetingBooked(endsAt time.Time, startsAt time.Time, ) *MeetingBooked`

NewMeetingBooked instantiates a new MeetingBooked object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeetingBookedWithDefaults

`func NewMeetingBookedWithDefaults() *MeetingBooked`

NewMeetingBookedWithDefaults instantiates a new MeetingBooked object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndsAt

`func (o *MeetingBooked) GetEndsAt() time.Time`

GetEndsAt returns the EndsAt field if non-nil, zero value otherwise.

### GetEndsAtOk

`func (o *MeetingBooked) GetEndsAtOk() (*time.Time, bool)`

GetEndsAtOk returns a tuple with the EndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndsAt

`func (o *MeetingBooked) SetEndsAt(v time.Time)`

SetEndsAt sets EndsAt field to given value.


### GetStartsAt

`func (o *MeetingBooked) GetStartsAt() time.Time`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *MeetingBooked) GetStartsAtOk() (*time.Time, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *MeetingBooked) SetStartsAt(v time.Time)`

SetStartsAt sets StartsAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


