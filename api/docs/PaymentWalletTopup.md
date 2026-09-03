# PaymentWalletTopup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountCents** | **int32** |  | 
**Created** | **time.Time** |  | 
**Currency** | **string** |  | 
**FailureReason** | **string** |  | 
**Id** | **int32** |  | 
**Reference** | **string** |  | 
**RequestedBy** | **int32** |  | 
**Status** | **string** |  | 
**TransactionId** | Pointer to **int32** |  | [optional] 
**Updated** | **time.Time** |  | 
**WalletId** | **int32** |  | 

## Methods

### NewPaymentWalletTopup

`func NewPaymentWalletTopup(amountCents int32, created time.Time, currency string, failureReason string, id int32, reference string, requestedBy int32, status string, updated time.Time, walletId int32, ) *PaymentWalletTopup`

NewPaymentWalletTopup instantiates a new PaymentWalletTopup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentWalletTopupWithDefaults

`func NewPaymentWalletTopupWithDefaults() *PaymentWalletTopup`

NewPaymentWalletTopupWithDefaults instantiates a new PaymentWalletTopup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountCents

`func (o *PaymentWalletTopup) GetAmountCents() int32`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *PaymentWalletTopup) GetAmountCentsOk() (*int32, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *PaymentWalletTopup) SetAmountCents(v int32)`

SetAmountCents sets AmountCents field to given value.


### GetCreated

`func (o *PaymentWalletTopup) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PaymentWalletTopup) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PaymentWalletTopup) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetCurrency

`func (o *PaymentWalletTopup) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PaymentWalletTopup) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PaymentWalletTopup) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetFailureReason

`func (o *PaymentWalletTopup) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *PaymentWalletTopup) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *PaymentWalletTopup) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.


### GetId

`func (o *PaymentWalletTopup) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentWalletTopup) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentWalletTopup) SetId(v int32)`

SetId sets Id field to given value.


### GetReference

`func (o *PaymentWalletTopup) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PaymentWalletTopup) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PaymentWalletTopup) SetReference(v string)`

SetReference sets Reference field to given value.


### GetRequestedBy

`func (o *PaymentWalletTopup) GetRequestedBy() int32`

GetRequestedBy returns the RequestedBy field if non-nil, zero value otherwise.

### GetRequestedByOk

`func (o *PaymentWalletTopup) GetRequestedByOk() (*int32, bool)`

GetRequestedByOk returns a tuple with the RequestedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedBy

`func (o *PaymentWalletTopup) SetRequestedBy(v int32)`

SetRequestedBy sets RequestedBy field to given value.


### GetStatus

`func (o *PaymentWalletTopup) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentWalletTopup) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentWalletTopup) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTransactionId

`func (o *PaymentWalletTopup) GetTransactionId() int32`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *PaymentWalletTopup) GetTransactionIdOk() (*int32, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *PaymentWalletTopup) SetTransactionId(v int32)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *PaymentWalletTopup) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetUpdated

`func (o *PaymentWalletTopup) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *PaymentWalletTopup) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *PaymentWalletTopup) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.


### GetWalletId

`func (o *PaymentWalletTopup) GetWalletId() int32`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *PaymentWalletTopup) GetWalletIdOk() (*int32, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *PaymentWalletTopup) SetWalletId(v int32)`

SetWalletId sets WalletId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


