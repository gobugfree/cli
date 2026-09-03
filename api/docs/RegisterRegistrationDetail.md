# RegisterRegistrationDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Country** | **string** |  | 
**Domain** | **string** |  | 
**EmailVerified** | **bool** |  | 
**Id** | **int32** |  | 
**LegalName** | **string** |  | 
**ReviewNote** | Pointer to **string** |  | [optional] 
**Status** | **string** |  | 

## Methods

### NewRegisterRegistrationDetail

`func NewRegisterRegistrationDetail(country string, domain string, emailVerified bool, id int32, legalName string, status string, ) *RegisterRegistrationDetail`

NewRegisterRegistrationDetail instantiates a new RegisterRegistrationDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterRegistrationDetailWithDefaults

`func NewRegisterRegistrationDetailWithDefaults() *RegisterRegistrationDetail`

NewRegisterRegistrationDetailWithDefaults instantiates a new RegisterRegistrationDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountry

`func (o *RegisterRegistrationDetail) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *RegisterRegistrationDetail) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *RegisterRegistrationDetail) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetDomain

`func (o *RegisterRegistrationDetail) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *RegisterRegistrationDetail) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *RegisterRegistrationDetail) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetEmailVerified

`func (o *RegisterRegistrationDetail) GetEmailVerified() bool`

GetEmailVerified returns the EmailVerified field if non-nil, zero value otherwise.

### GetEmailVerifiedOk

`func (o *RegisterRegistrationDetail) GetEmailVerifiedOk() (*bool, bool)`

GetEmailVerifiedOk returns a tuple with the EmailVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailVerified

`func (o *RegisterRegistrationDetail) SetEmailVerified(v bool)`

SetEmailVerified sets EmailVerified field to given value.


### GetId

`func (o *RegisterRegistrationDetail) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RegisterRegistrationDetail) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RegisterRegistrationDetail) SetId(v int32)`

SetId sets Id field to given value.


### GetLegalName

`func (o *RegisterRegistrationDetail) GetLegalName() string`

GetLegalName returns the LegalName field if non-nil, zero value otherwise.

### GetLegalNameOk

`func (o *RegisterRegistrationDetail) GetLegalNameOk() (*string, bool)`

GetLegalNameOk returns a tuple with the LegalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalName

`func (o *RegisterRegistrationDetail) SetLegalName(v string)`

SetLegalName sets LegalName field to given value.


### GetReviewNote

`func (o *RegisterRegistrationDetail) GetReviewNote() string`

GetReviewNote returns the ReviewNote field if non-nil, zero value otherwise.

### GetReviewNoteOk

`func (o *RegisterRegistrationDetail) GetReviewNoteOk() (*string, bool)`

GetReviewNoteOk returns a tuple with the ReviewNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewNote

`func (o *RegisterRegistrationDetail) SetReviewNote(v string)`

SetReviewNote sets ReviewNote field to given value.

### HasReviewNote

`func (o *RegisterRegistrationDetail) HasReviewNote() bool`

HasReviewNote returns a boolean if a field has been set.

### GetStatus

`func (o *RegisterRegistrationDetail) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RegisterRegistrationDetail) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RegisterRegistrationDetail) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


