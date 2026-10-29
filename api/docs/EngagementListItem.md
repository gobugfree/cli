# EngagementListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Badges** | Pointer to [**[]EngagementBadge**](EngagementBadge.md) |  | [optional] 
**Id** | **string** |  | 
**Industry** | Pointer to **string** |  | [optional] 
**Invited** | **bool** |  | 
**Kind** | **string** |  | 
**Locked** | **bool** |  | 
**Member** | **bool** |  | 
**OrgName** | **string** |  | 
**OrgSlug** | **string** |  | 
**RewardCurrency** | Pointer to **string** |  | [optional] 
**RewardMax** | Pointer to **float32** |  | [optional] 
**RewardMin** | Pointer to **float32** |  | [optional] 
**ScopeKinds** | Pointer to **[]string** |  | [optional] 
**Slug** | **string** |  | 
**Status** | **string** |  | 
**Summary** | Pointer to **string** |  | [optional] 
**Title** | **string** |  | 
**Visibility** | **string** |  | 

## Methods

### NewEngagementListItem

`func NewEngagementListItem(id string, invited bool, kind string, locked bool, member bool, orgName string, orgSlug string, slug string, status string, title string, visibility string, ) *EngagementListItem`

NewEngagementListItem instantiates a new EngagementListItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngagementListItemWithDefaults

`func NewEngagementListItemWithDefaults() *EngagementListItem`

NewEngagementListItemWithDefaults instantiates a new EngagementListItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBadges

`func (o *EngagementListItem) GetBadges() []EngagementBadge`

GetBadges returns the Badges field if non-nil, zero value otherwise.

### GetBadgesOk

`func (o *EngagementListItem) GetBadgesOk() (*[]EngagementBadge, bool)`

GetBadgesOk returns a tuple with the Badges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadges

`func (o *EngagementListItem) SetBadges(v []EngagementBadge)`

SetBadges sets Badges field to given value.

### HasBadges

`func (o *EngagementListItem) HasBadges() bool`

HasBadges returns a boolean if a field has been set.

### GetId

`func (o *EngagementListItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EngagementListItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EngagementListItem) SetId(v string)`

SetId sets Id field to given value.


### GetIndustry

`func (o *EngagementListItem) GetIndustry() string`

GetIndustry returns the Industry field if non-nil, zero value otherwise.

### GetIndustryOk

`func (o *EngagementListItem) GetIndustryOk() (*string, bool)`

GetIndustryOk returns a tuple with the Industry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustry

`func (o *EngagementListItem) SetIndustry(v string)`

SetIndustry sets Industry field to given value.

### HasIndustry

`func (o *EngagementListItem) HasIndustry() bool`

HasIndustry returns a boolean if a field has been set.

### GetInvited

`func (o *EngagementListItem) GetInvited() bool`

GetInvited returns the Invited field if non-nil, zero value otherwise.

### GetInvitedOk

`func (o *EngagementListItem) GetInvitedOk() (*bool, bool)`

GetInvitedOk returns a tuple with the Invited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvited

`func (o *EngagementListItem) SetInvited(v bool)`

SetInvited sets Invited field to given value.


### GetKind

`func (o *EngagementListItem) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *EngagementListItem) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *EngagementListItem) SetKind(v string)`

SetKind sets Kind field to given value.


### GetLocked

`func (o *EngagementListItem) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *EngagementListItem) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *EngagementListItem) SetLocked(v bool)`

SetLocked sets Locked field to given value.


### GetMember

`func (o *EngagementListItem) GetMember() bool`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *EngagementListItem) GetMemberOk() (*bool, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *EngagementListItem) SetMember(v bool)`

SetMember sets Member field to given value.


### GetOrgName

`func (o *EngagementListItem) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *EngagementListItem) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *EngagementListItem) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.


### GetOrgSlug

`func (o *EngagementListItem) GetOrgSlug() string`

GetOrgSlug returns the OrgSlug field if non-nil, zero value otherwise.

### GetOrgSlugOk

`func (o *EngagementListItem) GetOrgSlugOk() (*string, bool)`

GetOrgSlugOk returns a tuple with the OrgSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgSlug

`func (o *EngagementListItem) SetOrgSlug(v string)`

SetOrgSlug sets OrgSlug field to given value.


### GetRewardCurrency

`func (o *EngagementListItem) GetRewardCurrency() string`

GetRewardCurrency returns the RewardCurrency field if non-nil, zero value otherwise.

### GetRewardCurrencyOk

`func (o *EngagementListItem) GetRewardCurrencyOk() (*string, bool)`

GetRewardCurrencyOk returns a tuple with the RewardCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardCurrency

`func (o *EngagementListItem) SetRewardCurrency(v string)`

SetRewardCurrency sets RewardCurrency field to given value.

### HasRewardCurrency

`func (o *EngagementListItem) HasRewardCurrency() bool`

HasRewardCurrency returns a boolean if a field has been set.

### GetRewardMax

`func (o *EngagementListItem) GetRewardMax() float32`

GetRewardMax returns the RewardMax field if non-nil, zero value otherwise.

### GetRewardMaxOk

`func (o *EngagementListItem) GetRewardMaxOk() (*float32, bool)`

GetRewardMaxOk returns a tuple with the RewardMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardMax

`func (o *EngagementListItem) SetRewardMax(v float32)`

SetRewardMax sets RewardMax field to given value.

### HasRewardMax

`func (o *EngagementListItem) HasRewardMax() bool`

HasRewardMax returns a boolean if a field has been set.

### GetRewardMin

`func (o *EngagementListItem) GetRewardMin() float32`

GetRewardMin returns the RewardMin field if non-nil, zero value otherwise.

### GetRewardMinOk

`func (o *EngagementListItem) GetRewardMinOk() (*float32, bool)`

GetRewardMinOk returns a tuple with the RewardMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardMin

`func (o *EngagementListItem) SetRewardMin(v float32)`

SetRewardMin sets RewardMin field to given value.

### HasRewardMin

`func (o *EngagementListItem) HasRewardMin() bool`

HasRewardMin returns a boolean if a field has been set.

### GetScopeKinds

`func (o *EngagementListItem) GetScopeKinds() []string`

GetScopeKinds returns the ScopeKinds field if non-nil, zero value otherwise.

### GetScopeKindsOk

`func (o *EngagementListItem) GetScopeKindsOk() (*[]string, bool)`

GetScopeKindsOk returns a tuple with the ScopeKinds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeKinds

`func (o *EngagementListItem) SetScopeKinds(v []string)`

SetScopeKinds sets ScopeKinds field to given value.

### HasScopeKinds

`func (o *EngagementListItem) HasScopeKinds() bool`

HasScopeKinds returns a boolean if a field has been set.

### GetSlug

`func (o *EngagementListItem) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *EngagementListItem) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *EngagementListItem) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetStatus

`func (o *EngagementListItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EngagementListItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EngagementListItem) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSummary

`func (o *EngagementListItem) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *EngagementListItem) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *EngagementListItem) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *EngagementListItem) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetTitle

`func (o *EngagementListItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *EngagementListItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *EngagementListItem) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetVisibility

`func (o *EngagementListItem) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *EngagementListItem) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *EngagementListItem) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


