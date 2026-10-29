# EngagementDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessData** | Pointer to **string** |  | [optional] 
**Attachments** | Pointer to [**[]EngagementFile**](EngagementFile.md) |  | [optional] 
**Badges** | Pointer to [**[]EngagementBadge**](EngagementBadge.md) |  | [optional] 
**BriefingAccessible** | **bool** |  | 
**Cleared** | **bool** |  | 
**CreatedAt** | **time.Time** |  | 
**DisclosurePolicy** | **string** |  | 
**EndsOn** | Pointer to **time.Time** |  | [optional] 
**Id** | **string** |  | 
**Kind** | **string** |  | 
**Leaderboard** | Pointer to [**[]EngagementLeaderboard**](EngagementLeaderboard.md) |  | [optional] 
**Locked** | **bool** |  | 
**MessageProxyEnabled** | **bool** |  | 
**NdaAccepted** | **bool** |  | 
**NdaRequired** | **bool** |  | 
**NdaUrl** | Pointer to **string** |  | [optional] 
**OrgAbout** | Pointer to **string** |  | [optional] 
**RequiredVerificationLevel** | **string** |  | 
**RulesAccepted** | **bool** |  | 
**Slug** | **string** |  | 
**StartsOn** | Pointer to **time.Time** |  | [optional] 
**Status** | **string** |  | 
**Title** | **string** |  | 
**UpdatedAt** | **time.Time** |  | 
**Visibility** | **string** |  | 

## Methods

### NewEngagementDetail

`func NewEngagementDetail(briefingAccessible bool, cleared bool, createdAt time.Time, disclosurePolicy string, id string, kind string, locked bool, messageProxyEnabled bool, ndaAccepted bool, ndaRequired bool, requiredVerificationLevel string, rulesAccepted bool, slug string, status string, title string, updatedAt time.Time, visibility string, ) *EngagementDetail`

NewEngagementDetail instantiates a new EngagementDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngagementDetailWithDefaults

`func NewEngagementDetailWithDefaults() *EngagementDetail`

NewEngagementDetailWithDefaults instantiates a new EngagementDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessData

`func (o *EngagementDetail) GetAccessData() string`

GetAccessData returns the AccessData field if non-nil, zero value otherwise.

### GetAccessDataOk

`func (o *EngagementDetail) GetAccessDataOk() (*string, bool)`

GetAccessDataOk returns a tuple with the AccessData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessData

`func (o *EngagementDetail) SetAccessData(v string)`

SetAccessData sets AccessData field to given value.

### HasAccessData

`func (o *EngagementDetail) HasAccessData() bool`

HasAccessData returns a boolean if a field has been set.

### GetAttachments

`func (o *EngagementDetail) GetAttachments() []EngagementFile`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *EngagementDetail) GetAttachmentsOk() (*[]EngagementFile, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *EngagementDetail) SetAttachments(v []EngagementFile)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *EngagementDetail) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetBadges

`func (o *EngagementDetail) GetBadges() []EngagementBadge`

GetBadges returns the Badges field if non-nil, zero value otherwise.

### GetBadgesOk

`func (o *EngagementDetail) GetBadgesOk() (*[]EngagementBadge, bool)`

GetBadgesOk returns a tuple with the Badges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadges

`func (o *EngagementDetail) SetBadges(v []EngagementBadge)`

SetBadges sets Badges field to given value.

### HasBadges

`func (o *EngagementDetail) HasBadges() bool`

HasBadges returns a boolean if a field has been set.

### GetBriefingAccessible

`func (o *EngagementDetail) GetBriefingAccessible() bool`

GetBriefingAccessible returns the BriefingAccessible field if non-nil, zero value otherwise.

### GetBriefingAccessibleOk

`func (o *EngagementDetail) GetBriefingAccessibleOk() (*bool, bool)`

GetBriefingAccessibleOk returns a tuple with the BriefingAccessible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBriefingAccessible

`func (o *EngagementDetail) SetBriefingAccessible(v bool)`

SetBriefingAccessible sets BriefingAccessible field to given value.


### GetCleared

`func (o *EngagementDetail) GetCleared() bool`

GetCleared returns the Cleared field if non-nil, zero value otherwise.

### GetClearedOk

`func (o *EngagementDetail) GetClearedOk() (*bool, bool)`

GetClearedOk returns a tuple with the Cleared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleared

`func (o *EngagementDetail) SetCleared(v bool)`

SetCleared sets Cleared field to given value.


### GetCreatedAt

`func (o *EngagementDetail) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EngagementDetail) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EngagementDetail) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDisclosurePolicy

`func (o *EngagementDetail) GetDisclosurePolicy() string`

GetDisclosurePolicy returns the DisclosurePolicy field if non-nil, zero value otherwise.

### GetDisclosurePolicyOk

`func (o *EngagementDetail) GetDisclosurePolicyOk() (*string, bool)`

GetDisclosurePolicyOk returns a tuple with the DisclosurePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclosurePolicy

`func (o *EngagementDetail) SetDisclosurePolicy(v string)`

SetDisclosurePolicy sets DisclosurePolicy field to given value.


### GetEndsOn

`func (o *EngagementDetail) GetEndsOn() time.Time`

GetEndsOn returns the EndsOn field if non-nil, zero value otherwise.

### GetEndsOnOk

`func (o *EngagementDetail) GetEndsOnOk() (*time.Time, bool)`

GetEndsOnOk returns a tuple with the EndsOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndsOn

`func (o *EngagementDetail) SetEndsOn(v time.Time)`

SetEndsOn sets EndsOn field to given value.

### HasEndsOn

`func (o *EngagementDetail) HasEndsOn() bool`

HasEndsOn returns a boolean if a field has been set.

### GetId

`func (o *EngagementDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EngagementDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EngagementDetail) SetId(v string)`

SetId sets Id field to given value.


### GetKind

`func (o *EngagementDetail) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *EngagementDetail) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *EngagementDetail) SetKind(v string)`

SetKind sets Kind field to given value.


### GetLeaderboard

`func (o *EngagementDetail) GetLeaderboard() []EngagementLeaderboard`

GetLeaderboard returns the Leaderboard field if non-nil, zero value otherwise.

### GetLeaderboardOk

`func (o *EngagementDetail) GetLeaderboardOk() (*[]EngagementLeaderboard, bool)`

GetLeaderboardOk returns a tuple with the Leaderboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaderboard

`func (o *EngagementDetail) SetLeaderboard(v []EngagementLeaderboard)`

SetLeaderboard sets Leaderboard field to given value.

### HasLeaderboard

`func (o *EngagementDetail) HasLeaderboard() bool`

HasLeaderboard returns a boolean if a field has been set.

### GetLocked

`func (o *EngagementDetail) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *EngagementDetail) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *EngagementDetail) SetLocked(v bool)`

SetLocked sets Locked field to given value.


### GetMessageProxyEnabled

`func (o *EngagementDetail) GetMessageProxyEnabled() bool`

GetMessageProxyEnabled returns the MessageProxyEnabled field if non-nil, zero value otherwise.

### GetMessageProxyEnabledOk

`func (o *EngagementDetail) GetMessageProxyEnabledOk() (*bool, bool)`

GetMessageProxyEnabledOk returns a tuple with the MessageProxyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageProxyEnabled

`func (o *EngagementDetail) SetMessageProxyEnabled(v bool)`

SetMessageProxyEnabled sets MessageProxyEnabled field to given value.


### GetNdaAccepted

`func (o *EngagementDetail) GetNdaAccepted() bool`

GetNdaAccepted returns the NdaAccepted field if non-nil, zero value otherwise.

### GetNdaAcceptedOk

`func (o *EngagementDetail) GetNdaAcceptedOk() (*bool, bool)`

GetNdaAcceptedOk returns a tuple with the NdaAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNdaAccepted

`func (o *EngagementDetail) SetNdaAccepted(v bool)`

SetNdaAccepted sets NdaAccepted field to given value.


### GetNdaRequired

`func (o *EngagementDetail) GetNdaRequired() bool`

GetNdaRequired returns the NdaRequired field if non-nil, zero value otherwise.

### GetNdaRequiredOk

`func (o *EngagementDetail) GetNdaRequiredOk() (*bool, bool)`

GetNdaRequiredOk returns a tuple with the NdaRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNdaRequired

`func (o *EngagementDetail) SetNdaRequired(v bool)`

SetNdaRequired sets NdaRequired field to given value.


### GetNdaUrl

`func (o *EngagementDetail) GetNdaUrl() string`

GetNdaUrl returns the NdaUrl field if non-nil, zero value otherwise.

### GetNdaUrlOk

`func (o *EngagementDetail) GetNdaUrlOk() (*string, bool)`

GetNdaUrlOk returns a tuple with the NdaUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNdaUrl

`func (o *EngagementDetail) SetNdaUrl(v string)`

SetNdaUrl sets NdaUrl field to given value.

### HasNdaUrl

`func (o *EngagementDetail) HasNdaUrl() bool`

HasNdaUrl returns a boolean if a field has been set.

### GetOrgAbout

`func (o *EngagementDetail) GetOrgAbout() string`

GetOrgAbout returns the OrgAbout field if non-nil, zero value otherwise.

### GetOrgAboutOk

`func (o *EngagementDetail) GetOrgAboutOk() (*string, bool)`

GetOrgAboutOk returns a tuple with the OrgAbout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgAbout

`func (o *EngagementDetail) SetOrgAbout(v string)`

SetOrgAbout sets OrgAbout field to given value.

### HasOrgAbout

`func (o *EngagementDetail) HasOrgAbout() bool`

HasOrgAbout returns a boolean if a field has been set.

### GetRequiredVerificationLevel

`func (o *EngagementDetail) GetRequiredVerificationLevel() string`

GetRequiredVerificationLevel returns the RequiredVerificationLevel field if non-nil, zero value otherwise.

### GetRequiredVerificationLevelOk

`func (o *EngagementDetail) GetRequiredVerificationLevelOk() (*string, bool)`

GetRequiredVerificationLevelOk returns a tuple with the RequiredVerificationLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredVerificationLevel

`func (o *EngagementDetail) SetRequiredVerificationLevel(v string)`

SetRequiredVerificationLevel sets RequiredVerificationLevel field to given value.


### GetRulesAccepted

`func (o *EngagementDetail) GetRulesAccepted() bool`

GetRulesAccepted returns the RulesAccepted field if non-nil, zero value otherwise.

### GetRulesAcceptedOk

`func (o *EngagementDetail) GetRulesAcceptedOk() (*bool, bool)`

GetRulesAcceptedOk returns a tuple with the RulesAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRulesAccepted

`func (o *EngagementDetail) SetRulesAccepted(v bool)`

SetRulesAccepted sets RulesAccepted field to given value.


### GetSlug

`func (o *EngagementDetail) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *EngagementDetail) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *EngagementDetail) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetStartsOn

`func (o *EngagementDetail) GetStartsOn() time.Time`

GetStartsOn returns the StartsOn field if non-nil, zero value otherwise.

### GetStartsOnOk

`func (o *EngagementDetail) GetStartsOnOk() (*time.Time, bool)`

GetStartsOnOk returns a tuple with the StartsOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsOn

`func (o *EngagementDetail) SetStartsOn(v time.Time)`

SetStartsOn sets StartsOn field to given value.

### HasStartsOn

`func (o *EngagementDetail) HasStartsOn() bool`

HasStartsOn returns a boolean if a field has been set.

### GetStatus

`func (o *EngagementDetail) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EngagementDetail) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EngagementDetail) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTitle

`func (o *EngagementDetail) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *EngagementDetail) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *EngagementDetail) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetUpdatedAt

`func (o *EngagementDetail) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EngagementDetail) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EngagementDetail) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetVisibility

`func (o *EngagementDetail) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *EngagementDetail) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *EngagementDetail) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


