# MessageNotificationPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]MessageNotification**](MessageNotification.md) |  | 
**Next** | Pointer to **int32** |  | [optional] 

## Methods

### NewMessageNotificationPage

`func NewMessageNotificationPage(items []MessageNotification, ) *MessageNotificationPage`

NewMessageNotificationPage instantiates a new MessageNotificationPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageNotificationPageWithDefaults

`func NewMessageNotificationPageWithDefaults() *MessageNotificationPage`

NewMessageNotificationPageWithDefaults instantiates a new MessageNotificationPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *MessageNotificationPage) GetItems() []MessageNotification`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *MessageNotificationPage) GetItemsOk() (*[]MessageNotification, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *MessageNotificationPage) SetItems(v []MessageNotification)`

SetItems sets Items field to given value.


### GetNext

`func (o *MessageNotificationPage) GetNext() int32`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *MessageNotificationPage) GetNextOk() (*int32, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *MessageNotificationPage) SetNext(v int32)`

SetNext sets Next field to given value.

### HasNext

`func (o *MessageNotificationPage) HasNext() bool`

HasNext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


