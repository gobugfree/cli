# Scopes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assets** | [**[]ScopeAsset**](ScopeAsset.md) |  | 
**Engagements** | [**[]ScopeEngagementScopeInfo**](ScopeEngagementScopeInfo.md) |  | 

## Methods

### NewScopes

`func NewScopes(assets []ScopeAsset, engagements []ScopeEngagementScopeInfo, ) *Scopes`

NewScopes instantiates a new Scopes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScopesWithDefaults

`func NewScopesWithDefaults() *Scopes`

NewScopesWithDefaults instantiates a new Scopes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssets

`func (o *Scopes) GetAssets() []ScopeAsset`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *Scopes) GetAssetsOk() (*[]ScopeAsset, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *Scopes) SetAssets(v []ScopeAsset)`

SetAssets sets Assets field to given value.


### GetEngagements

`func (o *Scopes) GetEngagements() []ScopeEngagementScopeInfo`

GetEngagements returns the Engagements field if non-nil, zero value otherwise.

### GetEngagementsOk

`func (o *Scopes) GetEngagementsOk() (*[]ScopeEngagementScopeInfo, bool)`

GetEngagementsOk returns a tuple with the Engagements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngagements

`func (o *Scopes) SetEngagements(v []ScopeEngagementScopeInfo)`

SetEngagements sets Engagements field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


