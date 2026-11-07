# EngagementScopeItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**InScope** | **bool** |  | 
**Instructions** | Pointer to **string** |  | [optional] 
**Kind** | **string** |  | 
**Pattern** | **string** |  | 

## Methods

### NewEngagementScopeItem

`func NewEngagementScopeItem(inScope bool, kind string, pattern string, ) *EngagementScopeItem`

NewEngagementScopeItem instantiates a new EngagementScopeItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngagementScopeItemWithDefaults

`func NewEngagementScopeItemWithDefaults() *EngagementScopeItem`

NewEngagementScopeItemWithDefaults instantiates a new EngagementScopeItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *EngagementScopeItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EngagementScopeItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EngagementScopeItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EngagementScopeItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInScope

`func (o *EngagementScopeItem) GetInScope() bool`

GetInScope returns the InScope field if non-nil, zero value otherwise.

### GetInScopeOk

`func (o *EngagementScopeItem) GetInScopeOk() (*bool, bool)`

GetInScopeOk returns a tuple with the InScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInScope

`func (o *EngagementScopeItem) SetInScope(v bool)`

SetInScope sets InScope field to given value.


### GetInstructions

`func (o *EngagementScopeItem) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *EngagementScopeItem) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *EngagementScopeItem) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *EngagementScopeItem) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### GetKind

`func (o *EngagementScopeItem) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *EngagementScopeItem) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *EngagementScopeItem) SetKind(v string)`

SetKind sets Kind field to given value.


### GetPattern

`func (o *EngagementScopeItem) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *EngagementScopeItem) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *EngagementScopeItem) SetPattern(v string)`

SetPattern sets Pattern field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


