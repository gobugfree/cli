# ScopeAsset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Availability** | **string** |  | 
**Confidentiality** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**Criticality** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Environment** | **string** |  | 
**Id** | **int32** |  | 
**Identifier** | **string** |  | 
**Integrity** | **string** |  | 
**Kind** | **string** |  | 
**Tier** | Pointer to **string** |  | [optional] 
**UpdatedAt** | **time.Time** |  | 
**Version** | **int32** |  | 

## Methods

### NewScopeAsset

`func NewScopeAsset(availability string, confidentiality string, createdAt time.Time, criticality string, environment string, id int32, identifier string, integrity string, kind string, updatedAt time.Time, version int32, ) *ScopeAsset`

NewScopeAsset instantiates a new ScopeAsset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScopeAssetWithDefaults

`func NewScopeAssetWithDefaults() *ScopeAsset`

NewScopeAssetWithDefaults instantiates a new ScopeAsset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailability

`func (o *ScopeAsset) GetAvailability() string`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *ScopeAsset) GetAvailabilityOk() (*string, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *ScopeAsset) SetAvailability(v string)`

SetAvailability sets Availability field to given value.


### GetConfidentiality

`func (o *ScopeAsset) GetConfidentiality() string`

GetConfidentiality returns the Confidentiality field if non-nil, zero value otherwise.

### GetConfidentialityOk

`func (o *ScopeAsset) GetConfidentialityOk() (*string, bool)`

GetConfidentialityOk returns a tuple with the Confidentiality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidentiality

`func (o *ScopeAsset) SetConfidentiality(v string)`

SetConfidentiality sets Confidentiality field to given value.


### GetCreatedAt

`func (o *ScopeAsset) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ScopeAsset) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ScopeAsset) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCriticality

`func (o *ScopeAsset) GetCriticality() string`

GetCriticality returns the Criticality field if non-nil, zero value otherwise.

### GetCriticalityOk

`func (o *ScopeAsset) GetCriticalityOk() (*string, bool)`

GetCriticalityOk returns a tuple with the Criticality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticality

`func (o *ScopeAsset) SetCriticality(v string)`

SetCriticality sets Criticality field to given value.


### GetDescription

`func (o *ScopeAsset) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ScopeAsset) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ScopeAsset) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ScopeAsset) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnvironment

`func (o *ScopeAsset) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ScopeAsset) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ScopeAsset) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### GetId

`func (o *ScopeAsset) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScopeAsset) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScopeAsset) SetId(v int32)`

SetId sets Id field to given value.


### GetIdentifier

`func (o *ScopeAsset) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *ScopeAsset) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *ScopeAsset) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetIntegrity

`func (o *ScopeAsset) GetIntegrity() string`

GetIntegrity returns the Integrity field if non-nil, zero value otherwise.

### GetIntegrityOk

`func (o *ScopeAsset) GetIntegrityOk() (*string, bool)`

GetIntegrityOk returns a tuple with the Integrity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrity

`func (o *ScopeAsset) SetIntegrity(v string)`

SetIntegrity sets Integrity field to given value.


### GetKind

`func (o *ScopeAsset) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ScopeAsset) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ScopeAsset) SetKind(v string)`

SetKind sets Kind field to given value.


### GetTier

`func (o *ScopeAsset) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *ScopeAsset) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *ScopeAsset) SetTier(v string)`

SetTier sets Tier field to given value.

### HasTier

`func (o *ScopeAsset) HasTier() bool`

HasTier returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ScopeAsset) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ScopeAsset) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ScopeAsset) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetVersion

`func (o *ScopeAsset) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ScopeAsset) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ScopeAsset) SetVersion(v int32)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


