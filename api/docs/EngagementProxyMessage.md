# EngagementProxyMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** |  | 
**Text** | **string** |  | 

## Methods

### NewEngagementProxyMessage

`func NewEngagementProxyMessage(createdAt time.Time, text string, ) *EngagementProxyMessage`

NewEngagementProxyMessage instantiates a new EngagementProxyMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngagementProxyMessageWithDefaults

`func NewEngagementProxyMessageWithDefaults() *EngagementProxyMessage`

NewEngagementProxyMessageWithDefaults instantiates a new EngagementProxyMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *EngagementProxyMessage) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EngagementProxyMessage) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EngagementProxyMessage) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetText

`func (o *EngagementProxyMessage) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *EngagementProxyMessage) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *EngagementProxyMessage) SetText(v string)`

SetText sets Text field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


