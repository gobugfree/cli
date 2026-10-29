# UserSecurity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]UserSignInMethod**](UserSignInMethod.md) |  | 
**LinkedProviders** | **[]string** |  | 

## Methods

### NewUserSecurity

`func NewUserSecurity(items []UserSignInMethod, linkedProviders []string, ) *UserSecurity`

NewUserSecurity instantiates a new UserSecurity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSecurityWithDefaults

`func NewUserSecurityWithDefaults() *UserSecurity`

NewUserSecurityWithDefaults instantiates a new UserSecurity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *UserSecurity) GetItems() []UserSignInMethod`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *UserSecurity) GetItemsOk() (*[]UserSignInMethod, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *UserSecurity) SetItems(v []UserSignInMethod)`

SetItems sets Items field to given value.


### GetLinkedProviders

`func (o *UserSecurity) GetLinkedProviders() []string`

GetLinkedProviders returns the LinkedProviders field if non-nil, zero value otherwise.

### GetLinkedProvidersOk

`func (o *UserSecurity) GetLinkedProvidersOk() (*[]string, bool)`

GetLinkedProvidersOk returns a tuple with the LinkedProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedProviders

`func (o *UserSecurity) SetLinkedProviders(v []string)`

SetLinkedProviders sets LinkedProviders field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


