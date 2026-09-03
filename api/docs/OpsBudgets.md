# OpsBudgets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Overrides** | [**[]OpsBudgetOverride**](OpsBudgetOverride.md) |  | 
**Penalties** | [**[]OpsBudgetPenalty**](OpsBudgetPenalty.md) |  | 
**Tiers** | [**[]OpsBudgetTier**](OpsBudgetTier.md) |  | 

## Methods

### NewOpsBudgets

`func NewOpsBudgets(overrides []OpsBudgetOverride, penalties []OpsBudgetPenalty, tiers []OpsBudgetTier, ) *OpsBudgets`

NewOpsBudgets instantiates a new OpsBudgets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpsBudgetsWithDefaults

`func NewOpsBudgetsWithDefaults() *OpsBudgets`

NewOpsBudgetsWithDefaults instantiates a new OpsBudgets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOverrides

`func (o *OpsBudgets) GetOverrides() []OpsBudgetOverride`

GetOverrides returns the Overrides field if non-nil, zero value otherwise.

### GetOverridesOk

`func (o *OpsBudgets) GetOverridesOk() (*[]OpsBudgetOverride, bool)`

GetOverridesOk returns a tuple with the Overrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrides

`func (o *OpsBudgets) SetOverrides(v []OpsBudgetOverride)`

SetOverrides sets Overrides field to given value.


### GetPenalties

`func (o *OpsBudgets) GetPenalties() []OpsBudgetPenalty`

GetPenalties returns the Penalties field if non-nil, zero value otherwise.

### GetPenaltiesOk

`func (o *OpsBudgets) GetPenaltiesOk() (*[]OpsBudgetPenalty, bool)`

GetPenaltiesOk returns a tuple with the Penalties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPenalties

`func (o *OpsBudgets) SetPenalties(v []OpsBudgetPenalty)`

SetPenalties sets Penalties field to given value.


### GetTiers

`func (o *OpsBudgets) GetTiers() []OpsBudgetTier`

GetTiers returns the Tiers field if non-nil, zero value otherwise.

### GetTiersOk

`func (o *OpsBudgets) GetTiersOk() (*[]OpsBudgetTier, bool)`

GetTiersOk returns a tuple with the Tiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiers

`func (o *OpsBudgets) SetTiers(v []OpsBudgetTier)`

SetTiers sets Tiers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


