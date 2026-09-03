# ScopeEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asset** | Pointer to **string** |  | [optional] 
**AssetId** | **int32** |  | 
**AssetVersion** | **int32** |  | 
**CreatedAt** | **time.Time** |  | 
**Id** | **int32** |  | 
**InScope** | **bool** |  | 
**Instructions** | Pointer to **string** |  | [optional] 
**Kind** | **string** |  | 
**Pattern** | **string** |  | 
**Tier** | Pointer to **string** |  | [optional] 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewScopeEntry

`func NewScopeEntry(assetId int32, assetVersion int32, createdAt time.Time, id int32, inScope bool, kind string, pattern string, updatedAt time.Time, ) *ScopeEntry`

NewScopeEntry instantiates a new ScopeEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScopeEntryWithDefaults

`func NewScopeEntryWithDefaults() *ScopeEntry`

NewScopeEntryWithDefaults instantiates a new ScopeEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsset

`func (o *ScopeEntry) GetAsset() string`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *ScopeEntry) GetAssetOk() (*string, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *ScopeEntry) SetAsset(v string)`

SetAsset sets Asset field to given value.

### HasAsset

`func (o *ScopeEntry) HasAsset() bool`

HasAsset returns a boolean if a field has been set.

### GetAssetId

`func (o *ScopeEntry) GetAssetId() int32`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *ScopeEntry) GetAssetIdOk() (*int32, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *ScopeEntry) SetAssetId(v int32)`

SetAssetId sets AssetId field to given value.


### GetAssetVersion

`func (o *ScopeEntry) GetAssetVersion() int32`

GetAssetVersion returns the AssetVersion field if non-nil, zero value otherwise.

### GetAssetVersionOk

`func (o *ScopeEntry) GetAssetVersionOk() (*int32, bool)`

GetAssetVersionOk returns a tuple with the AssetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetVersion

`func (o *ScopeEntry) SetAssetVersion(v int32)`

SetAssetVersion sets AssetVersion field to given value.


### GetCreatedAt

`func (o *ScopeEntry) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ScopeEntry) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ScopeEntry) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetId

`func (o *ScopeEntry) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScopeEntry) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScopeEntry) SetId(v int32)`

SetId sets Id field to given value.


### GetInScope

`func (o *ScopeEntry) GetInScope() bool`

GetInScope returns the InScope field if non-nil, zero value otherwise.

### GetInScopeOk

`func (o *ScopeEntry) GetInScopeOk() (*bool, bool)`

GetInScopeOk returns a tuple with the InScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInScope

`func (o *ScopeEntry) SetInScope(v bool)`

SetInScope sets InScope field to given value.


### GetInstructions

`func (o *ScopeEntry) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *ScopeEntry) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *ScopeEntry) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *ScopeEntry) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### GetKind

`func (o *ScopeEntry) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ScopeEntry) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ScopeEntry) SetKind(v string)`

SetKind sets Kind field to given value.


### GetPattern

`func (o *ScopeEntry) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *ScopeEntry) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *ScopeEntry) SetPattern(v string)`

SetPattern sets Pattern field to given value.


### GetTier

`func (o *ScopeEntry) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *ScopeEntry) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *ScopeEntry) SetTier(v string)`

SetTier sets Tier field to given value.

### HasTier

`func (o *ScopeEntry) HasTier() bool`

HasTier returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ScopeEntry) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ScopeEntry) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ScopeEntry) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


