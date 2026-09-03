# OrganizationProfileVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | [**OrganizationProfileConfig**](OrganizationProfileConfig.md) |  | 
**CreatedAt** | **time.Time** |  | 
**Id** | **string** |  | 
**Status** | **string** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewOrganizationProfileVersion

`func NewOrganizationProfileVersion(config OrganizationProfileConfig, createdAt time.Time, id string, status string, updatedAt time.Time, ) *OrganizationProfileVersion`

NewOrganizationProfileVersion instantiates a new OrganizationProfileVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationProfileVersionWithDefaults

`func NewOrganizationProfileVersionWithDefaults() *OrganizationProfileVersion`

NewOrganizationProfileVersionWithDefaults instantiates a new OrganizationProfileVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *OrganizationProfileVersion) GetConfig() OrganizationProfileConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *OrganizationProfileVersion) GetConfigOk() (*OrganizationProfileConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *OrganizationProfileVersion) SetConfig(v OrganizationProfileConfig)`

SetConfig sets Config field to given value.


### GetCreatedAt

`func (o *OrganizationProfileVersion) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrganizationProfileVersion) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrganizationProfileVersion) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetId

`func (o *OrganizationProfileVersion) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationProfileVersion) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationProfileVersion) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *OrganizationProfileVersion) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrganizationProfileVersion) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrganizationProfileVersion) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUpdatedAt

`func (o *OrganizationProfileVersion) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OrganizationProfileVersion) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OrganizationProfileVersion) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


