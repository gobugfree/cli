# Badge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** |  | 
**Color** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**Criteria** | **string** |  | 
**Descriptions** | **map[string]string** |  | 
**Icon** | **string** |  | 
**IsActive** | **bool** |  | 
**Names** | **map[string]string** |  | 
**Params** | **map[string]interface{}** |  | 
**Stars** | **int32** |  | 
**Threshold** | **int32** |  | 

## Methods

### NewBadge

`func NewBadge(code string, color string, createdAt time.Time, criteria string, descriptions map[string]string, icon string, isActive bool, names map[string]string, params map[string]interface{}, stars int32, threshold int32, ) *Badge`

NewBadge instantiates a new Badge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBadgeWithDefaults

`func NewBadgeWithDefaults() *Badge`

NewBadgeWithDefaults instantiates a new Badge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *Badge) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Badge) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Badge) SetCode(v string)`

SetCode sets Code field to given value.


### GetColor

`func (o *Badge) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *Badge) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *Badge) SetColor(v string)`

SetColor sets Color field to given value.


### GetCreatedAt

`func (o *Badge) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Badge) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Badge) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCriteria

`func (o *Badge) GetCriteria() string`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *Badge) GetCriteriaOk() (*string, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *Badge) SetCriteria(v string)`

SetCriteria sets Criteria field to given value.


### GetDescriptions

`func (o *Badge) GetDescriptions() map[string]string`

GetDescriptions returns the Descriptions field if non-nil, zero value otherwise.

### GetDescriptionsOk

`func (o *Badge) GetDescriptionsOk() (*map[string]string, bool)`

GetDescriptionsOk returns a tuple with the Descriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptions

`func (o *Badge) SetDescriptions(v map[string]string)`

SetDescriptions sets Descriptions field to given value.


### GetIcon

`func (o *Badge) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *Badge) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *Badge) SetIcon(v string)`

SetIcon sets Icon field to given value.


### GetIsActive

`func (o *Badge) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *Badge) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *Badge) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetNames

`func (o *Badge) GetNames() map[string]string`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *Badge) GetNamesOk() (*map[string]string, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *Badge) SetNames(v map[string]string)`

SetNames sets Names field to given value.


### GetParams

`func (o *Badge) GetParams() map[string]interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *Badge) GetParamsOk() (*map[string]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *Badge) SetParams(v map[string]interface{})`

SetParams sets Params field to given value.


### GetStars

`func (o *Badge) GetStars() int32`

GetStars returns the Stars field if non-nil, zero value otherwise.

### GetStarsOk

`func (o *Badge) GetStarsOk() (*int32, bool)`

GetStarsOk returns a tuple with the Stars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStars

`func (o *Badge) SetStars(v int32)`

SetStars sets Stars field to given value.


### GetThreshold

`func (o *Badge) GetThreshold() int32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *Badge) GetThresholdOk() (*int32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *Badge) SetThreshold(v int32)`

SetThreshold sets Threshold field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


