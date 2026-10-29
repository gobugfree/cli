# UserPublicCertification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresOn** | Pointer to **time.Time** |  | [optional] 
**IssuedOn** | Pointer to **time.Time** |  | [optional] 
**Issuer** | **string** |  | 
**Name** | **string** |  | 
**Status** | **string** |  | 
**VerifyUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewUserPublicCertification

`func NewUserPublicCertification(issuer string, name string, status string, ) *UserPublicCertification`

NewUserPublicCertification instantiates a new UserPublicCertification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserPublicCertificationWithDefaults

`func NewUserPublicCertificationWithDefaults() *UserPublicCertification`

NewUserPublicCertificationWithDefaults instantiates a new UserPublicCertification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresOn

`func (o *UserPublicCertification) GetExpiresOn() time.Time`

GetExpiresOn returns the ExpiresOn field if non-nil, zero value otherwise.

### GetExpiresOnOk

`func (o *UserPublicCertification) GetExpiresOnOk() (*time.Time, bool)`

GetExpiresOnOk returns a tuple with the ExpiresOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresOn

`func (o *UserPublicCertification) SetExpiresOn(v time.Time)`

SetExpiresOn sets ExpiresOn field to given value.

### HasExpiresOn

`func (o *UserPublicCertification) HasExpiresOn() bool`

HasExpiresOn returns a boolean if a field has been set.

### GetIssuedOn

`func (o *UserPublicCertification) GetIssuedOn() time.Time`

GetIssuedOn returns the IssuedOn field if non-nil, zero value otherwise.

### GetIssuedOnOk

`func (o *UserPublicCertification) GetIssuedOnOk() (*time.Time, bool)`

GetIssuedOnOk returns a tuple with the IssuedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedOn

`func (o *UserPublicCertification) SetIssuedOn(v time.Time)`

SetIssuedOn sets IssuedOn field to given value.

### HasIssuedOn

`func (o *UserPublicCertification) HasIssuedOn() bool`

HasIssuedOn returns a boolean if a field has been set.

### GetIssuer

`func (o *UserPublicCertification) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *UserPublicCertification) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *UserPublicCertification) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.


### GetName

`func (o *UserPublicCertification) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserPublicCertification) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserPublicCertification) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *UserPublicCertification) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UserPublicCertification) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UserPublicCertification) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetVerifyUrl

`func (o *UserPublicCertification) GetVerifyUrl() string`

GetVerifyUrl returns the VerifyUrl field if non-nil, zero value otherwise.

### GetVerifyUrlOk

`func (o *UserPublicCertification) GetVerifyUrlOk() (*string, bool)`

GetVerifyUrlOk returns a tuple with the VerifyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyUrl

`func (o *UserPublicCertification) SetVerifyUrl(v string)`

SetVerifyUrl sets VerifyUrl field to given value.

### HasVerifyUrl

`func (o *UserPublicCertification) HasVerifyUrl() bool`

HasVerifyUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


