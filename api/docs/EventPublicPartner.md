# EventPublicPartner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accepted** | **int32** |  | 
**BudgetCents** | Pointer to **int32** |  | [optional] 
**BudgetCurrency** | **string** |  | 
**CommittedCents** | **int32** |  | 
**CriticalHigh** | **int32** |  | 
**New** | **int32** |  | 
**OrganizationName** | **string** |  | 
**PendingAcceptance** | **int32** |  | 
**Rejected** | **int32** |  | 
**Status** | **string** |  | 
**Title** | **string** |  | 

## Methods

### NewEventPublicPartner

`func NewEventPublicPartner(accepted int32, budgetCurrency string, committedCents int32, criticalHigh int32, new int32, organizationName string, pendingAcceptance int32, rejected int32, status string, title string, ) *EventPublicPartner`

NewEventPublicPartner instantiates a new EventPublicPartner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventPublicPartnerWithDefaults

`func NewEventPublicPartnerWithDefaults() *EventPublicPartner`

NewEventPublicPartnerWithDefaults instantiates a new EventPublicPartner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccepted

`func (o *EventPublicPartner) GetAccepted() int32`

GetAccepted returns the Accepted field if non-nil, zero value otherwise.

### GetAcceptedOk

`func (o *EventPublicPartner) GetAcceptedOk() (*int32, bool)`

GetAcceptedOk returns a tuple with the Accepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccepted

`func (o *EventPublicPartner) SetAccepted(v int32)`

SetAccepted sets Accepted field to given value.


### GetBudgetCents

`func (o *EventPublicPartner) GetBudgetCents() int32`

GetBudgetCents returns the BudgetCents field if non-nil, zero value otherwise.

### GetBudgetCentsOk

`func (o *EventPublicPartner) GetBudgetCentsOk() (*int32, bool)`

GetBudgetCentsOk returns a tuple with the BudgetCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetCents

`func (o *EventPublicPartner) SetBudgetCents(v int32)`

SetBudgetCents sets BudgetCents field to given value.

### HasBudgetCents

`func (o *EventPublicPartner) HasBudgetCents() bool`

HasBudgetCents returns a boolean if a field has been set.

### GetBudgetCurrency

`func (o *EventPublicPartner) GetBudgetCurrency() string`

GetBudgetCurrency returns the BudgetCurrency field if non-nil, zero value otherwise.

### GetBudgetCurrencyOk

`func (o *EventPublicPartner) GetBudgetCurrencyOk() (*string, bool)`

GetBudgetCurrencyOk returns a tuple with the BudgetCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetCurrency

`func (o *EventPublicPartner) SetBudgetCurrency(v string)`

SetBudgetCurrency sets BudgetCurrency field to given value.


### GetCommittedCents

`func (o *EventPublicPartner) GetCommittedCents() int32`

GetCommittedCents returns the CommittedCents field if non-nil, zero value otherwise.

### GetCommittedCentsOk

`func (o *EventPublicPartner) GetCommittedCentsOk() (*int32, bool)`

GetCommittedCentsOk returns a tuple with the CommittedCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommittedCents

`func (o *EventPublicPartner) SetCommittedCents(v int32)`

SetCommittedCents sets CommittedCents field to given value.


### GetCriticalHigh

`func (o *EventPublicPartner) GetCriticalHigh() int32`

GetCriticalHigh returns the CriticalHigh field if non-nil, zero value otherwise.

### GetCriticalHighOk

`func (o *EventPublicPartner) GetCriticalHighOk() (*int32, bool)`

GetCriticalHighOk returns a tuple with the CriticalHigh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalHigh

`func (o *EventPublicPartner) SetCriticalHigh(v int32)`

SetCriticalHigh sets CriticalHigh field to given value.


### GetNew

`func (o *EventPublicPartner) GetNew() int32`

GetNew returns the New field if non-nil, zero value otherwise.

### GetNewOk

`func (o *EventPublicPartner) GetNewOk() (*int32, bool)`

GetNewOk returns a tuple with the New field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNew

`func (o *EventPublicPartner) SetNew(v int32)`

SetNew sets New field to given value.


### GetOrganizationName

`func (o *EventPublicPartner) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *EventPublicPartner) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *EventPublicPartner) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.


### GetPendingAcceptance

`func (o *EventPublicPartner) GetPendingAcceptance() int32`

GetPendingAcceptance returns the PendingAcceptance field if non-nil, zero value otherwise.

### GetPendingAcceptanceOk

`func (o *EventPublicPartner) GetPendingAcceptanceOk() (*int32, bool)`

GetPendingAcceptanceOk returns a tuple with the PendingAcceptance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingAcceptance

`func (o *EventPublicPartner) SetPendingAcceptance(v int32)`

SetPendingAcceptance sets PendingAcceptance field to given value.


### GetRejected

`func (o *EventPublicPartner) GetRejected() int32`

GetRejected returns the Rejected field if non-nil, zero value otherwise.

### GetRejectedOk

`func (o *EventPublicPartner) GetRejectedOk() (*int32, bool)`

GetRejectedOk returns a tuple with the Rejected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejected

`func (o *EventPublicPartner) SetRejected(v int32)`

SetRejected sets Rejected field to given value.


### GetStatus

`func (o *EventPublicPartner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EventPublicPartner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EventPublicPartner) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTitle

`func (o *EventPublicPartner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *EventPublicPartner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *EventPublicPartner) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


