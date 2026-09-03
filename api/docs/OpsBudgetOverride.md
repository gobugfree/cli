# OpsBudgetOverride

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresAt** | Pointer to **time.Time** |  | [optional] 
**Id** | **int32** |  | 
**PendingLimit** | Pointer to **int32** |  | [optional] 
**Reason** | **string** |  | 
**Researcher** | **string** |  | 

## Methods

### NewOpsBudgetOverride

`func NewOpsBudgetOverride(id int32, reason string, researcher string, ) *OpsBudgetOverride`

NewOpsBudgetOverride instantiates a new OpsBudgetOverride object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpsBudgetOverrideWithDefaults

`func NewOpsBudgetOverrideWithDefaults() *OpsBudgetOverride`

NewOpsBudgetOverrideWithDefaults instantiates a new OpsBudgetOverride object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresAt

`func (o *OpsBudgetOverride) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *OpsBudgetOverride) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *OpsBudgetOverride) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *OpsBudgetOverride) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetId

`func (o *OpsBudgetOverride) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OpsBudgetOverride) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OpsBudgetOverride) SetId(v int32)`

SetId sets Id field to given value.


### GetPendingLimit

`func (o *OpsBudgetOverride) GetPendingLimit() int32`

GetPendingLimit returns the PendingLimit field if non-nil, zero value otherwise.

### GetPendingLimitOk

`func (o *OpsBudgetOverride) GetPendingLimitOk() (*int32, bool)`

GetPendingLimitOk returns a tuple with the PendingLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingLimit

`func (o *OpsBudgetOverride) SetPendingLimit(v int32)`

SetPendingLimit sets PendingLimit field to given value.

### HasPendingLimit

`func (o *OpsBudgetOverride) HasPendingLimit() bool`

HasPendingLimit returns a boolean if a field has been set.

### GetReason

`func (o *OpsBudgetOverride) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *OpsBudgetOverride) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *OpsBudgetOverride) SetReason(v string)`

SetReason sets Reason field to given value.


### GetResearcher

`func (o *OpsBudgetOverride) GetResearcher() string`

GetResearcher returns the Researcher field if non-nil, zero value otherwise.

### GetResearcherOk

`func (o *OpsBudgetOverride) GetResearcherOk() (*string, bool)`

GetResearcherOk returns a tuple with the Researcher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResearcher

`func (o *OpsBudgetOverride) SetResearcher(v string)`

SetResearcher sets Researcher field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


