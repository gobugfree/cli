# EngagementVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | [**EngagementVersionConfig**](EngagementVersionConfig.md) |  | 
**CreatedAt** | **time.Time** |  | 
**Id** | **string** |  | 
**Number** | **int32** |  | 
**Status** | **string** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewEngagementVersion

`func NewEngagementVersion(config EngagementVersionConfig, createdAt time.Time, id string, number int32, status string, updatedAt time.Time, ) *EngagementVersion`

NewEngagementVersion instantiates a new EngagementVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngagementVersionWithDefaults

`func NewEngagementVersionWithDefaults() *EngagementVersion`

NewEngagementVersionWithDefaults instantiates a new EngagementVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *EngagementVersion) GetConfig() EngagementVersionConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *EngagementVersion) GetConfigOk() (*EngagementVersionConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *EngagementVersion) SetConfig(v EngagementVersionConfig)`

SetConfig sets Config field to given value.


### GetCreatedAt

`func (o *EngagementVersion) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EngagementVersion) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EngagementVersion) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetId

`func (o *EngagementVersion) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EngagementVersion) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EngagementVersion) SetId(v string)`

SetId sets Id field to given value.


### GetNumber

`func (o *EngagementVersion) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *EngagementVersion) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *EngagementVersion) SetNumber(v int32)`

SetNumber sets Number field to given value.


### GetStatus

`func (o *EngagementVersion) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EngagementVersion) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EngagementVersion) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUpdatedAt

`func (o *EngagementVersion) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EngagementVersion) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EngagementVersion) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


