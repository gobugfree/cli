# BadgeCatalogItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** |  | 
**Color** | **string** |  | 
**Criteria** | **string** |  | 
**Current** | Pointer to **int32** |  | [optional] 
**Description** | **string** |  | 
**Earned** | **bool** |  | 
**EarnedAt** | Pointer to **time.Time** |  | [optional] 
**Hidden** | Pointer to **bool** |  | [optional] 
**Icon** | **string** |  | 
**Name** | **string** |  | 
**Stars** | **int32** |  | 
**Threshold** | **int32** |  | 

## Methods

### NewBadgeCatalogItem

`func NewBadgeCatalogItem(code string, color string, criteria string, description string, earned bool, icon string, name string, stars int32, threshold int32, ) *BadgeCatalogItem`

NewBadgeCatalogItem instantiates a new BadgeCatalogItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBadgeCatalogItemWithDefaults

`func NewBadgeCatalogItemWithDefaults() *BadgeCatalogItem`

NewBadgeCatalogItemWithDefaults instantiates a new BadgeCatalogItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *BadgeCatalogItem) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *BadgeCatalogItem) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *BadgeCatalogItem) SetCode(v string)`

SetCode sets Code field to given value.


### GetColor

`func (o *BadgeCatalogItem) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *BadgeCatalogItem) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *BadgeCatalogItem) SetColor(v string)`

SetColor sets Color field to given value.


### GetCriteria

`func (o *BadgeCatalogItem) GetCriteria() string`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *BadgeCatalogItem) GetCriteriaOk() (*string, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *BadgeCatalogItem) SetCriteria(v string)`

SetCriteria sets Criteria field to given value.


### GetCurrent

`func (o *BadgeCatalogItem) GetCurrent() int32`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *BadgeCatalogItem) GetCurrentOk() (*int32, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *BadgeCatalogItem) SetCurrent(v int32)`

SetCurrent sets Current field to given value.

### HasCurrent

`func (o *BadgeCatalogItem) HasCurrent() bool`

HasCurrent returns a boolean if a field has been set.

### GetDescription

`func (o *BadgeCatalogItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BadgeCatalogItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BadgeCatalogItem) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEarned

`func (o *BadgeCatalogItem) GetEarned() bool`

GetEarned returns the Earned field if non-nil, zero value otherwise.

### GetEarnedOk

`func (o *BadgeCatalogItem) GetEarnedOk() (*bool, bool)`

GetEarnedOk returns a tuple with the Earned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarned

`func (o *BadgeCatalogItem) SetEarned(v bool)`

SetEarned sets Earned field to given value.


### GetEarnedAt

`func (o *BadgeCatalogItem) GetEarnedAt() time.Time`

GetEarnedAt returns the EarnedAt field if non-nil, zero value otherwise.

### GetEarnedAtOk

`func (o *BadgeCatalogItem) GetEarnedAtOk() (*time.Time, bool)`

GetEarnedAtOk returns a tuple with the EarnedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarnedAt

`func (o *BadgeCatalogItem) SetEarnedAt(v time.Time)`

SetEarnedAt sets EarnedAt field to given value.

### HasEarnedAt

`func (o *BadgeCatalogItem) HasEarnedAt() bool`

HasEarnedAt returns a boolean if a field has been set.

### GetHidden

`func (o *BadgeCatalogItem) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *BadgeCatalogItem) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *BadgeCatalogItem) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *BadgeCatalogItem) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetIcon

`func (o *BadgeCatalogItem) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *BadgeCatalogItem) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *BadgeCatalogItem) SetIcon(v string)`

SetIcon sets Icon field to given value.


### GetName

`func (o *BadgeCatalogItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BadgeCatalogItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BadgeCatalogItem) SetName(v string)`

SetName sets Name field to given value.


### GetStars

`func (o *BadgeCatalogItem) GetStars() int32`

GetStars returns the Stars field if non-nil, zero value otherwise.

### GetStarsOk

`func (o *BadgeCatalogItem) GetStarsOk() (*int32, bool)`

GetStarsOk returns a tuple with the Stars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStars

`func (o *BadgeCatalogItem) SetStars(v int32)`

SetStars sets Stars field to given value.


### GetThreshold

`func (o *BadgeCatalogItem) GetThreshold() int32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *BadgeCatalogItem) GetThresholdOk() (*int32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *BadgeCatalogItem) SetThreshold(v int32)`

SetThreshold sets Threshold field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


