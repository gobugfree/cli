# UserCertification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CredentialId** | **string** |  | 
**ExpiresOn** | Pointer to **time.Time** |  | [optional] 
**Id** | **int32** |  | 
**IssuedOn** | Pointer to **time.Time** |  | [optional] 
**Issuer** | **string** |  | 
**IssuerName** | **string** |  | 
**Name** | **string** |  | 
**Note** | Pointer to **string** |  | [optional] 
**Status** | **string** |  | 
**Url** | Pointer to **string** |  | [optional] 
**VerifiedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewUserCertification

`func NewUserCertification(credentialId string, id int32, issuer string, issuerName string, name string, status string, ) *UserCertification`

NewUserCertification instantiates a new UserCertification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserCertificationWithDefaults

`func NewUserCertificationWithDefaults() *UserCertification`

NewUserCertificationWithDefaults instantiates a new UserCertification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentialId

`func (o *UserCertification) GetCredentialId() string`

GetCredentialId returns the CredentialId field if non-nil, zero value otherwise.

### GetCredentialIdOk

`func (o *UserCertification) GetCredentialIdOk() (*string, bool)`

GetCredentialIdOk returns a tuple with the CredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialId

`func (o *UserCertification) SetCredentialId(v string)`

SetCredentialId sets CredentialId field to given value.


### GetExpiresOn

`func (o *UserCertification) GetExpiresOn() time.Time`

GetExpiresOn returns the ExpiresOn field if non-nil, zero value otherwise.

### GetExpiresOnOk

`func (o *UserCertification) GetExpiresOnOk() (*time.Time, bool)`

GetExpiresOnOk returns a tuple with the ExpiresOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresOn

`func (o *UserCertification) SetExpiresOn(v time.Time)`

SetExpiresOn sets ExpiresOn field to given value.

### HasExpiresOn

`func (o *UserCertification) HasExpiresOn() bool`

HasExpiresOn returns a boolean if a field has been set.

### GetId

`func (o *UserCertification) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserCertification) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserCertification) SetId(v int32)`

SetId sets Id field to given value.


### GetIssuedOn

`func (o *UserCertification) GetIssuedOn() time.Time`

GetIssuedOn returns the IssuedOn field if non-nil, zero value otherwise.

### GetIssuedOnOk

`func (o *UserCertification) GetIssuedOnOk() (*time.Time, bool)`

GetIssuedOnOk returns a tuple with the IssuedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedOn

`func (o *UserCertification) SetIssuedOn(v time.Time)`

SetIssuedOn sets IssuedOn field to given value.

### HasIssuedOn

`func (o *UserCertification) HasIssuedOn() bool`

HasIssuedOn returns a boolean if a field has been set.

### GetIssuer

`func (o *UserCertification) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *UserCertification) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *UserCertification) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.


### GetIssuerName

`func (o *UserCertification) GetIssuerName() string`

GetIssuerName returns the IssuerName field if non-nil, zero value otherwise.

### GetIssuerNameOk

`func (o *UserCertification) GetIssuerNameOk() (*string, bool)`

GetIssuerNameOk returns a tuple with the IssuerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerName

`func (o *UserCertification) SetIssuerName(v string)`

SetIssuerName sets IssuerName field to given value.


### GetName

`func (o *UserCertification) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserCertification) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserCertification) SetName(v string)`

SetName sets Name field to given value.


### GetNote

`func (o *UserCertification) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *UserCertification) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *UserCertification) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *UserCertification) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetStatus

`func (o *UserCertification) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UserCertification) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UserCertification) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUrl

`func (o *UserCertification) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UserCertification) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UserCertification) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *UserCertification) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetVerifiedAt

`func (o *UserCertification) GetVerifiedAt() time.Time`

GetVerifiedAt returns the VerifiedAt field if non-nil, zero value otherwise.

### GetVerifiedAtOk

`func (o *UserCertification) GetVerifiedAtOk() (*time.Time, bool)`

GetVerifiedAtOk returns a tuple with the VerifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedAt

`func (o *UserCertification) SetVerifiedAt(v time.Time)`

SetVerifiedAt sets VerifiedAt field to given value.

### HasVerifiedAt

`func (o *UserCertification) HasVerifiedAt() bool`

HasVerifiedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


