# MessageNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **string** |  | [optional] 
**Created** | **time.Time** |  | 
**EntityId** | Pointer to **string** |  | [optional] 
**EntityKind** | Pointer to **string** |  | [optional] 
**Id** | **int32** |  | 
**IsRead** | **bool** |  | 
**Kind** | **string** |  | 
**Title** | **string** |  | 

## Methods

### NewMessageNotification

`func NewMessageNotification(created time.Time, id int32, isRead bool, kind string, title string, ) *MessageNotification`

NewMessageNotification instantiates a new MessageNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageNotificationWithDefaults

`func NewMessageNotificationWithDefaults() *MessageNotification`

NewMessageNotificationWithDefaults instantiates a new MessageNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *MessageNotification) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *MessageNotification) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *MessageNotification) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *MessageNotification) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetCreated

`func (o *MessageNotification) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *MessageNotification) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *MessageNotification) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetEntityId

`func (o *MessageNotification) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *MessageNotification) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *MessageNotification) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *MessageNotification) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetEntityKind

`func (o *MessageNotification) GetEntityKind() string`

GetEntityKind returns the EntityKind field if non-nil, zero value otherwise.

### GetEntityKindOk

`func (o *MessageNotification) GetEntityKindOk() (*string, bool)`

GetEntityKindOk returns a tuple with the EntityKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityKind

`func (o *MessageNotification) SetEntityKind(v string)`

SetEntityKind sets EntityKind field to given value.

### HasEntityKind

`func (o *MessageNotification) HasEntityKind() bool`

HasEntityKind returns a boolean if a field has been set.

### GetId

`func (o *MessageNotification) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MessageNotification) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MessageNotification) SetId(v int32)`

SetId sets Id field to given value.


### GetIsRead

`func (o *MessageNotification) GetIsRead() bool`

GetIsRead returns the IsRead field if non-nil, zero value otherwise.

### GetIsReadOk

`func (o *MessageNotification) GetIsReadOk() (*bool, bool)`

GetIsReadOk returns a tuple with the IsRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRead

`func (o *MessageNotification) SetIsRead(v bool)`

SetIsRead sets IsRead field to given value.


### GetKind

`func (o *MessageNotification) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *MessageNotification) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *MessageNotification) SetKind(v string)`

SetKind sets Kind field to given value.


### GetTitle

`func (o *MessageNotification) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MessageNotification) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MessageNotification) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


