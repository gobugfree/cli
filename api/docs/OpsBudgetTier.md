# OpsBudgetTier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PendingLimit** | Pointer to **int32** |  | [optional] 
**Tier** | **string** |  | 

## Methods

### NewOpsBudgetTier

`func NewOpsBudgetTier(tier string, ) *OpsBudgetTier`

NewOpsBudgetTier instantiates a new OpsBudgetTier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpsBudgetTierWithDefaults

`func NewOpsBudgetTierWithDefaults() *OpsBudgetTier`

NewOpsBudgetTierWithDefaults instantiates a new OpsBudgetTier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPendingLimit

`func (o *OpsBudgetTier) GetPendingLimit() int32`

GetPendingLimit returns the PendingLimit field if non-nil, zero value otherwise.

### GetPendingLimitOk

`func (o *OpsBudgetTier) GetPendingLimitOk() (*int32, bool)`

GetPendingLimitOk returns a tuple with the PendingLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingLimit

`func (o *OpsBudgetTier) SetPendingLimit(v int32)`

SetPendingLimit sets PendingLimit field to given value.

### HasPendingLimit

`func (o *OpsBudgetTier) HasPendingLimit() bool`

HasPendingLimit returns a boolean if a field has been set.

### GetTier

`func (o *OpsBudgetTier) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *OpsBudgetTier) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *OpsBudgetTier) SetTier(v string)`

SetTier sets Tier field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


