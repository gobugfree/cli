# EventInvite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**EndsAt** | **time.Time** |  | 
**Joined** | **bool** |  | 
**Programs** | [**[]EventInviteProgram**](EventInviteProgram.md) |  | 
**StartsAt** | **time.Time** |  | 
**Title** | **string** |  | 

## Methods

### NewEventInvite

`func NewEventInvite(endsAt time.Time, joined bool, programs []EventInviteProgram, startsAt time.Time, title string, ) *EventInvite`

NewEventInvite instantiates a new EventInvite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventInviteWithDefaults

`func NewEventInviteWithDefaults() *EventInvite`

NewEventInviteWithDefaults instantiates a new EventInvite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *EventInvite) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EventInvite) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EventInvite) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EventInvite) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEndsAt

`func (o *EventInvite) GetEndsAt() time.Time`

GetEndsAt returns the EndsAt field if non-nil, zero value otherwise.

### GetEndsAtOk

`func (o *EventInvite) GetEndsAtOk() (*time.Time, bool)`

GetEndsAtOk returns a tuple with the EndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndsAt

`func (o *EventInvite) SetEndsAt(v time.Time)`

SetEndsAt sets EndsAt field to given value.


### GetJoined

`func (o *EventInvite) GetJoined() bool`

GetJoined returns the Joined field if non-nil, zero value otherwise.

### GetJoinedOk

`func (o *EventInvite) GetJoinedOk() (*bool, bool)`

GetJoinedOk returns a tuple with the Joined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoined

`func (o *EventInvite) SetJoined(v bool)`

SetJoined sets Joined field to given value.


### GetPrograms

`func (o *EventInvite) GetPrograms() []EventInviteProgram`

GetPrograms returns the Programs field if non-nil, zero value otherwise.

### GetProgramsOk

`func (o *EventInvite) GetProgramsOk() (*[]EventInviteProgram, bool)`

GetProgramsOk returns a tuple with the Programs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrograms

`func (o *EventInvite) SetPrograms(v []EventInviteProgram)`

SetPrograms sets Programs field to given value.


### GetStartsAt

`func (o *EventInvite) GetStartsAt() time.Time`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *EventInvite) GetStartsAtOk() (*time.Time, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *EventInvite) SetStartsAt(v time.Time)`

SetStartsAt sets StartsAt field to given value.


### GetTitle

`func (o *EventInvite) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *EventInvite) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *EventInvite) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


