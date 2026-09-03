# PaymentTopupQueueItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgName** | **string** |  | 
**OrgSlug** | **string** |  | 
**WalletTopup** | [**PaymentWalletTopup**](PaymentWalletTopup.md) |  | 

## Methods

### NewPaymentTopupQueueItem

`func NewPaymentTopupQueueItem(orgName string, orgSlug string, walletTopup PaymentWalletTopup, ) *PaymentTopupQueueItem`

NewPaymentTopupQueueItem instantiates a new PaymentTopupQueueItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentTopupQueueItemWithDefaults

`func NewPaymentTopupQueueItemWithDefaults() *PaymentTopupQueueItem`

NewPaymentTopupQueueItemWithDefaults instantiates a new PaymentTopupQueueItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgName

`func (o *PaymentTopupQueueItem) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *PaymentTopupQueueItem) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *PaymentTopupQueueItem) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.


### GetOrgSlug

`func (o *PaymentTopupQueueItem) GetOrgSlug() string`

GetOrgSlug returns the OrgSlug field if non-nil, zero value otherwise.

### GetOrgSlugOk

`func (o *PaymentTopupQueueItem) GetOrgSlugOk() (*string, bool)`

GetOrgSlugOk returns a tuple with the OrgSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgSlug

`func (o *PaymentTopupQueueItem) SetOrgSlug(v string)`

SetOrgSlug sets OrgSlug field to given value.


### GetWalletTopup

`func (o *PaymentTopupQueueItem) GetWalletTopup() PaymentWalletTopup`

GetWalletTopup returns the WalletTopup field if non-nil, zero value otherwise.

### GetWalletTopupOk

`func (o *PaymentTopupQueueItem) GetWalletTopupOk() (*PaymentWalletTopup, bool)`

GetWalletTopupOk returns a tuple with the WalletTopup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletTopup

`func (o *PaymentTopupQueueItem) SetWalletTopup(v PaymentWalletTopup)`

SetWalletTopup sets WalletTopup field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


