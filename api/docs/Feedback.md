# Feedback

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwaitingStaff** | **bool** |  | 
**CreatedAt** | **string** |  | 
**Id** | **int32** |  | 
**Kind** | **string** |  | 
**LastMessageAt** | **string** |  | 
**Message** | **string** |  | 
**Page** | Pointer to **string** |  | [optional] 
**Rating** | Pointer to **int32** |  | [optional] 
**ReplyCount** | **int32** |  | 
**Status** | **string** |  | 

## Methods

### NewFeedback

`func NewFeedback(awaitingStaff bool, createdAt string, id int32, kind string, lastMessageAt string, message string, replyCount int32, status string, ) *Feedback`

NewFeedback instantiates a new Feedback object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeedbackWithDefaults

`func NewFeedbackWithDefaults() *Feedback`

NewFeedbackWithDefaults instantiates a new Feedback object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwaitingStaff

`func (o *Feedback) GetAwaitingStaff() bool`

GetAwaitingStaff returns the AwaitingStaff field if non-nil, zero value otherwise.

### GetAwaitingStaffOk

`func (o *Feedback) GetAwaitingStaffOk() (*bool, bool)`

GetAwaitingStaffOk returns a tuple with the AwaitingStaff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwaitingStaff

`func (o *Feedback) SetAwaitingStaff(v bool)`

SetAwaitingStaff sets AwaitingStaff field to given value.


### GetCreatedAt

`func (o *Feedback) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Feedback) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Feedback) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetId

`func (o *Feedback) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Feedback) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Feedback) SetId(v int32)`

SetId sets Id field to given value.


### GetKind

`func (o *Feedback) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Feedback) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Feedback) SetKind(v string)`

SetKind sets Kind field to given value.


### GetLastMessageAt

`func (o *Feedback) GetLastMessageAt() string`

GetLastMessageAt returns the LastMessageAt field if non-nil, zero value otherwise.

### GetLastMessageAtOk

`func (o *Feedback) GetLastMessageAtOk() (*string, bool)`

GetLastMessageAtOk returns a tuple with the LastMessageAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMessageAt

`func (o *Feedback) SetLastMessageAt(v string)`

SetLastMessageAt sets LastMessageAt field to given value.


### GetMessage

`func (o *Feedback) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Feedback) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Feedback) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetPage

`func (o *Feedback) GetPage() string`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *Feedback) GetPageOk() (*string, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *Feedback) SetPage(v string)`

SetPage sets Page field to given value.

### HasPage

`func (o *Feedback) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetRating

`func (o *Feedback) GetRating() int32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *Feedback) GetRatingOk() (*int32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *Feedback) SetRating(v int32)`

SetRating sets Rating field to given value.

### HasRating

`func (o *Feedback) HasRating() bool`

HasRating returns a boolean if a field has been set.

### GetReplyCount

`func (o *Feedback) GetReplyCount() int32`

GetReplyCount returns the ReplyCount field if non-nil, zero value otherwise.

### GetReplyCountOk

`func (o *Feedback) GetReplyCountOk() (*int32, bool)`

GetReplyCountOk returns a tuple with the ReplyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyCount

`func (o *Feedback) SetReplyCount(v int32)`

SetReplyCount sets ReplyCount field to given value.


### GetStatus

`func (o *Feedback) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Feedback) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Feedback) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


