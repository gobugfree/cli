# UserPublicExternal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Handle** | **string** |  | 
**LastChecked** | Pointer to **time.Time** |  | [optional] 
**Provider** | **string** |  | 
**Url** | **string** |  | 
**VerifiedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewUserPublicExternal

`func NewUserPublicExternal(handle string, provider string, url string, ) *UserPublicExternal`

NewUserPublicExternal instantiates a new UserPublicExternal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserPublicExternalWithDefaults

`func NewUserPublicExternalWithDefaults() *UserPublicExternal`

NewUserPublicExternalWithDefaults instantiates a new UserPublicExternal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHandle

`func (o *UserPublicExternal) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *UserPublicExternal) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *UserPublicExternal) SetHandle(v string)`

SetHandle sets Handle field to given value.


### GetLastChecked

`func (o *UserPublicExternal) GetLastChecked() time.Time`

GetLastChecked returns the LastChecked field if non-nil, zero value otherwise.

### GetLastCheckedOk

`func (o *UserPublicExternal) GetLastCheckedOk() (*time.Time, bool)`

GetLastCheckedOk returns a tuple with the LastChecked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastChecked

`func (o *UserPublicExternal) SetLastChecked(v time.Time)`

SetLastChecked sets LastChecked field to given value.

### HasLastChecked

`func (o *UserPublicExternal) HasLastChecked() bool`

HasLastChecked returns a boolean if a field has been set.

### GetProvider

`func (o *UserPublicExternal) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *UserPublicExternal) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *UserPublicExternal) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetUrl

`func (o *UserPublicExternal) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UserPublicExternal) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UserPublicExternal) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetVerifiedAt

`func (o *UserPublicExternal) GetVerifiedAt() time.Time`

GetVerifiedAt returns the VerifiedAt field if non-nil, zero value otherwise.

### GetVerifiedAtOk

`func (o *UserPublicExternal) GetVerifiedAtOk() (*time.Time, bool)`

GetVerifiedAtOk returns a tuple with the VerifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedAt

`func (o *UserPublicExternal) SetVerifiedAt(v time.Time)`

SetVerifiedAt sets VerifiedAt field to given value.

### HasVerifiedAt

`func (o *UserPublicExternal) HasVerifiedAt() bool`

HasVerifiedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


