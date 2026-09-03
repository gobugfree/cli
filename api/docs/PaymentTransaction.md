# PaymentTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountCents** | **int32** |  | 
**BalanceAfterCents** | **int32** |  | 
**CreatedAt** | **time.Time** |  | 
**Currency** | **string** |  | 
**Id** | **int32** |  | 
**Kind** | **string** |  | 
**Reference** | Pointer to **string** |  | [optional] 
**ReportTitle** | Pointer to **string** |  | [optional] 

## Methods

### NewPaymentTransaction

`func NewPaymentTransaction(amountCents int32, balanceAfterCents int32, createdAt time.Time, currency string, id int32, kind string, ) *PaymentTransaction`

NewPaymentTransaction instantiates a new PaymentTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentTransactionWithDefaults

`func NewPaymentTransactionWithDefaults() *PaymentTransaction`

NewPaymentTransactionWithDefaults instantiates a new PaymentTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountCents

`func (o *PaymentTransaction) GetAmountCents() int32`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *PaymentTransaction) GetAmountCentsOk() (*int32, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *PaymentTransaction) SetAmountCents(v int32)`

SetAmountCents sets AmountCents field to given value.


### GetBalanceAfterCents

`func (o *PaymentTransaction) GetBalanceAfterCents() int32`

GetBalanceAfterCents returns the BalanceAfterCents field if non-nil, zero value otherwise.

### GetBalanceAfterCentsOk

`func (o *PaymentTransaction) GetBalanceAfterCentsOk() (*int32, bool)`

GetBalanceAfterCentsOk returns a tuple with the BalanceAfterCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceAfterCents

`func (o *PaymentTransaction) SetBalanceAfterCents(v int32)`

SetBalanceAfterCents sets BalanceAfterCents field to given value.


### GetCreatedAt

`func (o *PaymentTransaction) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PaymentTransaction) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PaymentTransaction) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCurrency

`func (o *PaymentTransaction) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PaymentTransaction) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PaymentTransaction) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetId

`func (o *PaymentTransaction) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentTransaction) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentTransaction) SetId(v int32)`

SetId sets Id field to given value.


### GetKind

`func (o *PaymentTransaction) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *PaymentTransaction) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *PaymentTransaction) SetKind(v string)`

SetKind sets Kind field to given value.


### GetReference

`func (o *PaymentTransaction) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PaymentTransaction) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PaymentTransaction) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *PaymentTransaction) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReportTitle

`func (o *PaymentTransaction) GetReportTitle() string`

GetReportTitle returns the ReportTitle field if non-nil, zero value otherwise.

### GetReportTitleOk

`func (o *PaymentTransaction) GetReportTitleOk() (*string, bool)`

GetReportTitleOk returns a tuple with the ReportTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportTitle

`func (o *PaymentTransaction) SetReportTitle(v string)`

SetReportTitle sets ReportTitle field to given value.

### HasReportTitle

`func (o *PaymentTransaction) HasReportTitle() bool`

HasReportTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


