# ProofExternalProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailureReason** | Pointer to **string** |  | [optional] 
**Handle** | **string** |  | 
**Id** | **int32** |  | 
**Instructions** | Pointer to **string** |  | [optional] 
**LastChecked** | Pointer to **time.Time** |  | [optional] 
**Provider** | **string** |  | 
**Secret** | Pointer to **string** |  | [optional] 
**SecretExpires** | Pointer to **time.Time** |  | [optional] 
**Status** | **string** |  | 
**Url** | **string** |  | 
**VerifiedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewProofExternalProfile

`func NewProofExternalProfile(handle string, id int32, provider string, status string, url string, ) *ProofExternalProfile`

NewProofExternalProfile instantiates a new ProofExternalProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProofExternalProfileWithDefaults

`func NewProofExternalProfileWithDefaults() *ProofExternalProfile`

NewProofExternalProfileWithDefaults instantiates a new ProofExternalProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailureReason

`func (o *ProofExternalProfile) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *ProofExternalProfile) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *ProofExternalProfile) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *ProofExternalProfile) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### GetHandle

`func (o *ProofExternalProfile) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *ProofExternalProfile) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *ProofExternalProfile) SetHandle(v string)`

SetHandle sets Handle field to given value.


### GetId

`func (o *ProofExternalProfile) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProofExternalProfile) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProofExternalProfile) SetId(v int32)`

SetId sets Id field to given value.


### GetInstructions

`func (o *ProofExternalProfile) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *ProofExternalProfile) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *ProofExternalProfile) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *ProofExternalProfile) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### GetLastChecked

`func (o *ProofExternalProfile) GetLastChecked() time.Time`

GetLastChecked returns the LastChecked field if non-nil, zero value otherwise.

### GetLastCheckedOk

`func (o *ProofExternalProfile) GetLastCheckedOk() (*time.Time, bool)`

GetLastCheckedOk returns a tuple with the LastChecked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastChecked

`func (o *ProofExternalProfile) SetLastChecked(v time.Time)`

SetLastChecked sets LastChecked field to given value.

### HasLastChecked

`func (o *ProofExternalProfile) HasLastChecked() bool`

HasLastChecked returns a boolean if a field has been set.

### GetProvider

`func (o *ProofExternalProfile) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ProofExternalProfile) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ProofExternalProfile) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetSecret

`func (o *ProofExternalProfile) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *ProofExternalProfile) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *ProofExternalProfile) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *ProofExternalProfile) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetSecretExpires

`func (o *ProofExternalProfile) GetSecretExpires() time.Time`

GetSecretExpires returns the SecretExpires field if non-nil, zero value otherwise.

### GetSecretExpiresOk

`func (o *ProofExternalProfile) GetSecretExpiresOk() (*time.Time, bool)`

GetSecretExpiresOk returns a tuple with the SecretExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretExpires

`func (o *ProofExternalProfile) SetSecretExpires(v time.Time)`

SetSecretExpires sets SecretExpires field to given value.

### HasSecretExpires

`func (o *ProofExternalProfile) HasSecretExpires() bool`

HasSecretExpires returns a boolean if a field has been set.

### GetStatus

`func (o *ProofExternalProfile) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProofExternalProfile) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProofExternalProfile) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUrl

`func (o *ProofExternalProfile) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ProofExternalProfile) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ProofExternalProfile) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetVerifiedAt

`func (o *ProofExternalProfile) GetVerifiedAt() time.Time`

GetVerifiedAt returns the VerifiedAt field if non-nil, zero value otherwise.

### GetVerifiedAtOk

`func (o *ProofExternalProfile) GetVerifiedAtOk() (*time.Time, bool)`

GetVerifiedAtOk returns a tuple with the VerifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedAt

`func (o *ProofExternalProfile) SetVerifiedAt(v time.Time)`

SetVerifiedAt sets VerifiedAt field to given value.

### HasVerifiedAt

`func (o *ProofExternalProfile) HasVerifiedAt() bool`

HasVerifiedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


