# ScopeEngagementScope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Editable** | **bool** |  | 
**Entries** | [**[]ScopeEntry**](ScopeEntry.md) |  | 
**Slug** | **string** |  | 
**Title** | **string** |  | 
**VersionNumber** | Pointer to **int32** |  | [optional] 
**VersionStatus** | Pointer to **string** |  | [optional] 

## Methods

### NewScopeEngagementScope

`func NewScopeEngagementScope(editable bool, entries []ScopeEntry, slug string, title string, ) *ScopeEngagementScope`

NewScopeEngagementScope instantiates a new ScopeEngagementScope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScopeEngagementScopeWithDefaults

`func NewScopeEngagementScopeWithDefaults() *ScopeEngagementScope`

NewScopeEngagementScopeWithDefaults instantiates a new ScopeEngagementScope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEditable

`func (o *ScopeEngagementScope) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *ScopeEngagementScope) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *ScopeEngagementScope) SetEditable(v bool)`

SetEditable sets Editable field to given value.


### GetEntries

`func (o *ScopeEngagementScope) GetEntries() []ScopeEntry`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *ScopeEngagementScope) GetEntriesOk() (*[]ScopeEntry, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *ScopeEngagementScope) SetEntries(v []ScopeEntry)`

SetEntries sets Entries field to given value.


### GetSlug

`func (o *ScopeEngagementScope) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *ScopeEngagementScope) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *ScopeEngagementScope) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetTitle

`func (o *ScopeEngagementScope) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ScopeEngagementScope) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ScopeEngagementScope) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetVersionNumber

`func (o *ScopeEngagementScope) GetVersionNumber() int32`

GetVersionNumber returns the VersionNumber field if non-nil, zero value otherwise.

### GetVersionNumberOk

`func (o *ScopeEngagementScope) GetVersionNumberOk() (*int32, bool)`

GetVersionNumberOk returns a tuple with the VersionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionNumber

`func (o *ScopeEngagementScope) SetVersionNumber(v int32)`

SetVersionNumber sets VersionNumber field to given value.

### HasVersionNumber

`func (o *ScopeEngagementScope) HasVersionNumber() bool`

HasVersionNumber returns a boolean if a field has been set.

### GetVersionStatus

`func (o *ScopeEngagementScope) GetVersionStatus() string`

GetVersionStatus returns the VersionStatus field if non-nil, zero value otherwise.

### GetVersionStatusOk

`func (o *ScopeEngagementScope) GetVersionStatusOk() (*string, bool)`

GetVersionStatusOk returns a tuple with the VersionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionStatus

`func (o *ScopeEngagementScope) SetVersionStatus(v string)`

SetVersionStatus sets VersionStatus field to given value.

### HasVersionStatus

`func (o *ScopeEngagementScope) HasVersionStatus() bool`

HasVersionStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


