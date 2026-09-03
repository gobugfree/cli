# EngagementLifecycle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Changes** | Pointer to [**[]EngagementConfigChange**](EngagementConfigChange.md) |  | [optional] 
**Id** | **string** |  | 
**LiveVersion** | Pointer to **int32** |  | [optional] 
**OpenVersion** | Pointer to [**EngagementVersion**](EngagementVersion.md) |  | [optional] 
**Slug** | **string** |  | 
**Title** | **string** |  | 

## Methods

### NewEngagementLifecycle

`func NewEngagementLifecycle(id string, slug string, title string, ) *EngagementLifecycle`

NewEngagementLifecycle instantiates a new EngagementLifecycle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngagementLifecycleWithDefaults

`func NewEngagementLifecycleWithDefaults() *EngagementLifecycle`

NewEngagementLifecycleWithDefaults instantiates a new EngagementLifecycle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChanges

`func (o *EngagementLifecycle) GetChanges() []EngagementConfigChange`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *EngagementLifecycle) GetChangesOk() (*[]EngagementConfigChange, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *EngagementLifecycle) SetChanges(v []EngagementConfigChange)`

SetChanges sets Changes field to given value.

### HasChanges

`func (o *EngagementLifecycle) HasChanges() bool`

HasChanges returns a boolean if a field has been set.

### GetId

`func (o *EngagementLifecycle) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EngagementLifecycle) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EngagementLifecycle) SetId(v string)`

SetId sets Id field to given value.


### GetLiveVersion

`func (o *EngagementLifecycle) GetLiveVersion() int32`

GetLiveVersion returns the LiveVersion field if non-nil, zero value otherwise.

### GetLiveVersionOk

`func (o *EngagementLifecycle) GetLiveVersionOk() (*int32, bool)`

GetLiveVersionOk returns a tuple with the LiveVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveVersion

`func (o *EngagementLifecycle) SetLiveVersion(v int32)`

SetLiveVersion sets LiveVersion field to given value.

### HasLiveVersion

`func (o *EngagementLifecycle) HasLiveVersion() bool`

HasLiveVersion returns a boolean if a field has been set.

### GetOpenVersion

`func (o *EngagementLifecycle) GetOpenVersion() EngagementVersion`

GetOpenVersion returns the OpenVersion field if non-nil, zero value otherwise.

### GetOpenVersionOk

`func (o *EngagementLifecycle) GetOpenVersionOk() (*EngagementVersion, bool)`

GetOpenVersionOk returns a tuple with the OpenVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenVersion

`func (o *EngagementLifecycle) SetOpenVersion(v EngagementVersion)`

SetOpenVersion sets OpenVersion field to given value.

### HasOpenVersion

`func (o *EngagementLifecycle) HasOpenVersion() bool`

HasOpenVersion returns a boolean if a field has been set.

### GetSlug

`func (o *EngagementLifecycle) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *EngagementLifecycle) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *EngagementLifecycle) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetTitle

`func (o *EngagementLifecycle) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *EngagementLifecycle) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *EngagementLifecycle) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


