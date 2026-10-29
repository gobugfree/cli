# OrganizationInvitation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** |  | 
**HasFinance** | **bool** |  | 
**HasScanReports** | **bool** |  | 
**Id** | **int32** |  | 
**InvitedBy** | **string** |  | 
**IsAdmin** | **bool** |  | 
**Message** | **string** |  | 
**Organization** | **string** |  | 
**OrganizationName** | **string** |  | 
**Programs** | **[]string** |  | 

## Methods

### NewOrganizationInvitation

`func NewOrganizationInvitation(createdAt time.Time, hasFinance bool, hasScanReports bool, id int32, invitedBy string, isAdmin bool, message string, organization string, organizationName string, programs []string, ) *OrganizationInvitation`

NewOrganizationInvitation instantiates a new OrganizationInvitation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationInvitationWithDefaults

`func NewOrganizationInvitationWithDefaults() *OrganizationInvitation`

NewOrganizationInvitationWithDefaults instantiates a new OrganizationInvitation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *OrganizationInvitation) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrganizationInvitation) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrganizationInvitation) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetHasFinance

`func (o *OrganizationInvitation) GetHasFinance() bool`

GetHasFinance returns the HasFinance field if non-nil, zero value otherwise.

### GetHasFinanceOk

`func (o *OrganizationInvitation) GetHasFinanceOk() (*bool, bool)`

GetHasFinanceOk returns a tuple with the HasFinance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFinance

`func (o *OrganizationInvitation) SetHasFinance(v bool)`

SetHasFinance sets HasFinance field to given value.


### GetHasScanReports

`func (o *OrganizationInvitation) GetHasScanReports() bool`

GetHasScanReports returns the HasScanReports field if non-nil, zero value otherwise.

### GetHasScanReportsOk

`func (o *OrganizationInvitation) GetHasScanReportsOk() (*bool, bool)`

GetHasScanReportsOk returns a tuple with the HasScanReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasScanReports

`func (o *OrganizationInvitation) SetHasScanReports(v bool)`

SetHasScanReports sets HasScanReports field to given value.


### GetId

`func (o *OrganizationInvitation) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationInvitation) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationInvitation) SetId(v int32)`

SetId sets Id field to given value.


### GetInvitedBy

`func (o *OrganizationInvitation) GetInvitedBy() string`

GetInvitedBy returns the InvitedBy field if non-nil, zero value otherwise.

### GetInvitedByOk

`func (o *OrganizationInvitation) GetInvitedByOk() (*string, bool)`

GetInvitedByOk returns a tuple with the InvitedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitedBy

`func (o *OrganizationInvitation) SetInvitedBy(v string)`

SetInvitedBy sets InvitedBy field to given value.


### GetIsAdmin

`func (o *OrganizationInvitation) GetIsAdmin() bool`

GetIsAdmin returns the IsAdmin field if non-nil, zero value otherwise.

### GetIsAdminOk

`func (o *OrganizationInvitation) GetIsAdminOk() (*bool, bool)`

GetIsAdminOk returns a tuple with the IsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdmin

`func (o *OrganizationInvitation) SetIsAdmin(v bool)`

SetIsAdmin sets IsAdmin field to given value.


### GetMessage

`func (o *OrganizationInvitation) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *OrganizationInvitation) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *OrganizationInvitation) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetOrganization

`func (o *OrganizationInvitation) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *OrganizationInvitation) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *OrganizationInvitation) SetOrganization(v string)`

SetOrganization sets Organization field to given value.


### GetOrganizationName

`func (o *OrganizationInvitation) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *OrganizationInvitation) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *OrganizationInvitation) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.


### GetPrograms

`func (o *OrganizationInvitation) GetPrograms() []string`

GetPrograms returns the Programs field if non-nil, zero value otherwise.

### GetProgramsOk

`func (o *OrganizationInvitation) GetProgramsOk() (*[]string, bool)`

GetProgramsOk returns a tuple with the Programs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrograms

`func (o *OrganizationInvitation) SetPrograms(v []string)`

SetPrograms sets Programs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


