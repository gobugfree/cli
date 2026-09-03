# OrganizationInvitesList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dir** | **string** |  | 
**Items** | [**[]OrganizationMember**](OrganizationMember.md) |  | 
**Page** | **int32** |  | 
**Size** | **int32** |  | 
**Sort** | **string** |  | 
**Total** | **int32** |  | 

## Methods

### NewOrganizationInvitesList

`func NewOrganizationInvitesList(dir string, items []OrganizationMember, page int32, size int32, sort string, total int32, ) *OrganizationInvitesList`

NewOrganizationInvitesList instantiates a new OrganizationInvitesList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationInvitesListWithDefaults

`func NewOrganizationInvitesListWithDefaults() *OrganizationInvitesList`

NewOrganizationInvitesListWithDefaults instantiates a new OrganizationInvitesList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDir

`func (o *OrganizationInvitesList) GetDir() string`

GetDir returns the Dir field if non-nil, zero value otherwise.

### GetDirOk

`func (o *OrganizationInvitesList) GetDirOk() (*string, bool)`

GetDirOk returns a tuple with the Dir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDir

`func (o *OrganizationInvitesList) SetDir(v string)`

SetDir sets Dir field to given value.


### GetItems

`func (o *OrganizationInvitesList) GetItems() []OrganizationMember`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *OrganizationInvitesList) GetItemsOk() (*[]OrganizationMember, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *OrganizationInvitesList) SetItems(v []OrganizationMember)`

SetItems sets Items field to given value.


### GetPage

`func (o *OrganizationInvitesList) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *OrganizationInvitesList) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *OrganizationInvitesList) SetPage(v int32)`

SetPage sets Page field to given value.


### GetSize

`func (o *OrganizationInvitesList) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *OrganizationInvitesList) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *OrganizationInvitesList) SetSize(v int32)`

SetSize sets Size field to given value.


### GetSort

`func (o *OrganizationInvitesList) GetSort() string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *OrganizationInvitesList) GetSortOk() (*string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *OrganizationInvitesList) SetSort(v string)`

SetSort sets Sort field to given value.


### GetTotal

`func (o *OrganizationInvitesList) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *OrganizationInvitesList) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *OrganizationInvitesList) SetTotal(v int32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


