# MeetingBookingPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Booked** | Pointer to [**MeetingBooked**](MeetingBooked.md) |  | [optional] 
**Days** | Pointer to [**[]MeetingDay**](MeetingDay.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**Meeting** | **string** |  | 
**Minutes** | **int32** |  | 
**Online** | **bool** |  | 
**Organization** | **string** |  | 
**Status** | **string** |  | 
**TimeZone** | **string** |  | 
**Week** | Pointer to **string** |  | [optional] 
**With** | **[]string** |  | 

## Methods

### NewMeetingBookingPage

`func NewMeetingBookingPage(meeting string, minutes int32, online bool, organization string, status string, timeZone string, with []string, ) *MeetingBookingPage`

NewMeetingBookingPage instantiates a new MeetingBookingPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeetingBookingPageWithDefaults

`func NewMeetingBookingPageWithDefaults() *MeetingBookingPage`

NewMeetingBookingPageWithDefaults instantiates a new MeetingBookingPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBooked

`func (o *MeetingBookingPage) GetBooked() MeetingBooked`

GetBooked returns the Booked field if non-nil, zero value otherwise.

### GetBookedOk

`func (o *MeetingBookingPage) GetBookedOk() (*MeetingBooked, bool)`

GetBookedOk returns a tuple with the Booked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBooked

`func (o *MeetingBookingPage) SetBooked(v MeetingBooked)`

SetBooked sets Booked field to given value.

### HasBooked

`func (o *MeetingBookingPage) HasBooked() bool`

HasBooked returns a boolean if a field has been set.

### GetDays

`func (o *MeetingBookingPage) GetDays() []MeetingDay`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *MeetingBookingPage) GetDaysOk() (*[]MeetingDay, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *MeetingBookingPage) SetDays(v []MeetingDay)`

SetDays sets Days field to given value.

### HasDays

`func (o *MeetingBookingPage) HasDays() bool`

HasDays returns a boolean if a field has been set.

### GetDescription

`func (o *MeetingBookingPage) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MeetingBookingPage) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MeetingBookingPage) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MeetingBookingPage) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLocation

`func (o *MeetingBookingPage) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *MeetingBookingPage) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *MeetingBookingPage) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *MeetingBookingPage) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetMeeting

`func (o *MeetingBookingPage) GetMeeting() string`

GetMeeting returns the Meeting field if non-nil, zero value otherwise.

### GetMeetingOk

`func (o *MeetingBookingPage) GetMeetingOk() (*string, bool)`

GetMeetingOk returns a tuple with the Meeting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeeting

`func (o *MeetingBookingPage) SetMeeting(v string)`

SetMeeting sets Meeting field to given value.


### GetMinutes

`func (o *MeetingBookingPage) GetMinutes() int32`

GetMinutes returns the Minutes field if non-nil, zero value otherwise.

### GetMinutesOk

`func (o *MeetingBookingPage) GetMinutesOk() (*int32, bool)`

GetMinutesOk returns a tuple with the Minutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinutes

`func (o *MeetingBookingPage) SetMinutes(v int32)`

SetMinutes sets Minutes field to given value.


### GetOnline

`func (o *MeetingBookingPage) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *MeetingBookingPage) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *MeetingBookingPage) SetOnline(v bool)`

SetOnline sets Online field to given value.


### GetOrganization

`func (o *MeetingBookingPage) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *MeetingBookingPage) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *MeetingBookingPage) SetOrganization(v string)`

SetOrganization sets Organization field to given value.


### GetStatus

`func (o *MeetingBookingPage) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MeetingBookingPage) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MeetingBookingPage) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTimeZone

`func (o *MeetingBookingPage) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *MeetingBookingPage) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *MeetingBookingPage) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.


### GetWeek

`func (o *MeetingBookingPage) GetWeek() string`

GetWeek returns the Week field if non-nil, zero value otherwise.

### GetWeekOk

`func (o *MeetingBookingPage) GetWeekOk() (*string, bool)`

GetWeekOk returns a tuple with the Week field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeek

`func (o *MeetingBookingPage) SetWeek(v string)`

SetWeek sets Week field to given value.

### HasWeek

`func (o *MeetingBookingPage) HasWeek() bool`

HasWeek returns a boolean if a field has been set.

### GetWith

`func (o *MeetingBookingPage) GetWith() []string`

GetWith returns the With field if non-nil, zero value otherwise.

### GetWithOk

`func (o *MeetingBookingPage) GetWithOk() (*[]string, bool)`

GetWithOk returns a tuple with the With field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWith

`func (o *MeetingBookingPage) SetWith(v []string)`

SetWith sets With field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


