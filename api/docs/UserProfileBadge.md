# UserProfileBadge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** |  | 
**EarnedAt** | **time.Time** |  | 
**Hidden** | **bool** |  | 
**Name** | **string** |  | 

## Methods

### NewUserProfileBadge

`func NewUserProfileBadge(code string, earnedAt time.Time, hidden bool, name string, ) *UserProfileBadge`

NewUserProfileBadge instantiates a new UserProfileBadge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserProfileBadgeWithDefaults

`func NewUserProfileBadgeWithDefaults() *UserProfileBadge`

NewUserProfileBadgeWithDefaults instantiates a new UserProfileBadge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *UserProfileBadge) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UserProfileBadge) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UserProfileBadge) SetCode(v string)`

SetCode sets Code field to given value.


### GetEarnedAt

`func (o *UserProfileBadge) GetEarnedAt() time.Time`

GetEarnedAt returns the EarnedAt field if non-nil, zero value otherwise.

### GetEarnedAtOk

`func (o *UserProfileBadge) GetEarnedAtOk() (*time.Time, bool)`

GetEarnedAtOk returns a tuple with the EarnedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarnedAt

`func (o *UserProfileBadge) SetEarnedAt(v time.Time)`

SetEarnedAt sets EarnedAt field to given value.


### GetHidden

`func (o *UserProfileBadge) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *UserProfileBadge) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *UserProfileBadge) SetHidden(v bool)`

SetHidden sets Hidden field to given value.


### GetName

`func (o *UserProfileBadge) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserProfileBadge) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserProfileBadge) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


