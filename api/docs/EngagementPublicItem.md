# EngagementPublicItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Badges** | Pointer to [**[]EngagementBadge**](EngagementBadge.md) |  | [optional] 
**Kind** | **string** |  | 
**Locked** | **bool** |  | 
**LogoUrl** | Pointer to **string** |  | [optional] 
**OrgName** | **string** |  | 
**OrgSlug** | **string** |  | 
**Slug** | **string** |  | 
**Summary** | Pointer to **string** |  | [optional] 
**Title** | **string** |  | 

## Methods

### NewEngagementPublicItem

`func NewEngagementPublicItem(kind string, locked bool, orgName string, orgSlug string, slug string, title string, ) *EngagementPublicItem`

NewEngagementPublicItem instantiates a new EngagementPublicItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngagementPublicItemWithDefaults

`func NewEngagementPublicItemWithDefaults() *EngagementPublicItem`

NewEngagementPublicItemWithDefaults instantiates a new EngagementPublicItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBadges

`func (o *EngagementPublicItem) GetBadges() []EngagementBadge`

GetBadges returns the Badges field if non-nil, zero value otherwise.

### GetBadgesOk

`func (o *EngagementPublicItem) GetBadgesOk() (*[]EngagementBadge, bool)`

GetBadgesOk returns a tuple with the Badges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadges

`func (o *EngagementPublicItem) SetBadges(v []EngagementBadge)`

SetBadges sets Badges field to given value.

### HasBadges

`func (o *EngagementPublicItem) HasBadges() bool`

HasBadges returns a boolean if a field has been set.

### GetKind

`func (o *EngagementPublicItem) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *EngagementPublicItem) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *EngagementPublicItem) SetKind(v string)`

SetKind sets Kind field to given value.


### GetLocked

`func (o *EngagementPublicItem) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *EngagementPublicItem) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *EngagementPublicItem) SetLocked(v bool)`

SetLocked sets Locked field to given value.


### GetLogoUrl

`func (o *EngagementPublicItem) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *EngagementPublicItem) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *EngagementPublicItem) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *EngagementPublicItem) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### GetOrgName

`func (o *EngagementPublicItem) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *EngagementPublicItem) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *EngagementPublicItem) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.


### GetOrgSlug

`func (o *EngagementPublicItem) GetOrgSlug() string`

GetOrgSlug returns the OrgSlug field if non-nil, zero value otherwise.

### GetOrgSlugOk

`func (o *EngagementPublicItem) GetOrgSlugOk() (*string, bool)`

GetOrgSlugOk returns a tuple with the OrgSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgSlug

`func (o *EngagementPublicItem) SetOrgSlug(v string)`

SetOrgSlug sets OrgSlug field to given value.


### GetSlug

`func (o *EngagementPublicItem) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *EngagementPublicItem) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *EngagementPublicItem) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetSummary

`func (o *EngagementPublicItem) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *EngagementPublicItem) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *EngagementPublicItem) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *EngagementPublicItem) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetTitle

`func (o *EngagementPublicItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *EngagementPublicItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *EngagementPublicItem) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


