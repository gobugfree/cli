# OrganizationMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** |  | 
**Email** | **string** |  | 
**Grants** | [**[]OrganizationMemberGrant**](OrganizationMemberGrant.md) |  | 
**HasFinance** | **bool** |  | 
**Id** | **int32** |  | 
**InvitedBy** | Pointer to **string** |  | [optional] 
**IsAdmin** | **bool** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Status** | **string** |  | 
**UserId** | **int32** |  | 

## Methods

### NewOrganizationMember

`func NewOrganizationMember(createdAt time.Time, email string, grants []OrganizationMemberGrant, hasFinance bool, id int32, isAdmin bool, status string, userId int32, ) *OrganizationMember`

NewOrganizationMember instantiates a new OrganizationMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationMemberWithDefaults

`func NewOrganizationMemberWithDefaults() *OrganizationMember`

NewOrganizationMemberWithDefaults instantiates a new OrganizationMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *OrganizationMember) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrganizationMember) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrganizationMember) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetEmail

`func (o *OrganizationMember) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OrganizationMember) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OrganizationMember) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetGrants

`func (o *OrganizationMember) GetGrants() []OrganizationMemberGrant`

GetGrants returns the Grants field if non-nil, zero value otherwise.

### GetGrantsOk

`func (o *OrganizationMember) GetGrantsOk() (*[]OrganizationMemberGrant, bool)`

GetGrantsOk returns a tuple with the Grants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrants

`func (o *OrganizationMember) SetGrants(v []OrganizationMemberGrant)`

SetGrants sets Grants field to given value.


### GetHasFinance

`func (o *OrganizationMember) GetHasFinance() bool`

GetHasFinance returns the HasFinance field if non-nil, zero value otherwise.

### GetHasFinanceOk

`func (o *OrganizationMember) GetHasFinanceOk() (*bool, bool)`

GetHasFinanceOk returns a tuple with the HasFinance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFinance

`func (o *OrganizationMember) SetHasFinance(v bool)`

SetHasFinance sets HasFinance field to given value.


### GetId

`func (o *OrganizationMember) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationMember) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationMember) SetId(v int32)`

SetId sets Id field to given value.


### GetInvitedBy

`func (o *OrganizationMember) GetInvitedBy() string`

GetInvitedBy returns the InvitedBy field if non-nil, zero value otherwise.

### GetInvitedByOk

`func (o *OrganizationMember) GetInvitedByOk() (*string, bool)`

GetInvitedByOk returns a tuple with the InvitedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitedBy

`func (o *OrganizationMember) SetInvitedBy(v string)`

SetInvitedBy sets InvitedBy field to given value.

### HasInvitedBy

`func (o *OrganizationMember) HasInvitedBy() bool`

HasInvitedBy returns a boolean if a field has been set.

### GetIsAdmin

`func (o *OrganizationMember) GetIsAdmin() bool`

GetIsAdmin returns the IsAdmin field if non-nil, zero value otherwise.

### GetIsAdminOk

`func (o *OrganizationMember) GetIsAdminOk() (*bool, bool)`

GetIsAdminOk returns a tuple with the IsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdmin

`func (o *OrganizationMember) SetIsAdmin(v bool)`

SetIsAdmin sets IsAdmin field to given value.


### GetName

`func (o *OrganizationMember) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationMember) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationMember) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrganizationMember) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *OrganizationMember) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrganizationMember) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrganizationMember) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUserId

`func (o *OrganizationMember) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *OrganizationMember) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *OrganizationMember) SetUserId(v int32)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


