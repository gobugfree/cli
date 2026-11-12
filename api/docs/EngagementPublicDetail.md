# EngagementPublicDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Badges** | Pointer to [**[]EngagementBadge**](EngagementBadge.md) |  | [optional] 
**Bounty** | Pointer to [**[]EngagementPublicBounty**](EngagementPublicBounty.md) |  | [optional] 
**BriefingAccessible** | **bool** |  | 
**EndsOn** | Pointer to **time.Time** |  | [optional] 
**Kind** | **string** |  | 
**Locked** | **bool** |  | 
**LogoUrl** | Pointer to **string** |  | [optional] 
**NdaRequired** | **bool** |  | 
**OrgAbout** | Pointer to **string** |  | [optional] 
**OrgName** | **string** |  | 
**OrgSlug** | **string** |  | 
**Program** | Pointer to [**EngagementProgram**](EngagementProgram.md) |  | [optional] 
**RequiredVerificationLevel** | **string** |  | 
**Scope** | Pointer to [**[]EngagementScopeItem**](EngagementScopeItem.md) |  | [optional] 
**SignInUrl** | **string** |  | 
**Slug** | **string** |  | 
**StartsOn** | Pointer to **time.Time** |  | [optional] 
**Status** | **string** |  | 
**Summary** | Pointer to **string** |  | [optional] 
**Title** | **string** |  | 

## Methods

### NewEngagementPublicDetail

`func NewEngagementPublicDetail(briefingAccessible bool, kind string, locked bool, ndaRequired bool, orgName string, orgSlug string, requiredVerificationLevel string, signInUrl string, slug string, status string, title string, ) *EngagementPublicDetail`

NewEngagementPublicDetail instantiates a new EngagementPublicDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngagementPublicDetailWithDefaults

`func NewEngagementPublicDetailWithDefaults() *EngagementPublicDetail`

NewEngagementPublicDetailWithDefaults instantiates a new EngagementPublicDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBadges

`func (o *EngagementPublicDetail) GetBadges() []EngagementBadge`

GetBadges returns the Badges field if non-nil, zero value otherwise.

### GetBadgesOk

`func (o *EngagementPublicDetail) GetBadgesOk() (*[]EngagementBadge, bool)`

GetBadgesOk returns a tuple with the Badges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadges

`func (o *EngagementPublicDetail) SetBadges(v []EngagementBadge)`

SetBadges sets Badges field to given value.

### HasBadges

`func (o *EngagementPublicDetail) HasBadges() bool`

HasBadges returns a boolean if a field has been set.

### GetBounty

`func (o *EngagementPublicDetail) GetBounty() []EngagementPublicBounty`

GetBounty returns the Bounty field if non-nil, zero value otherwise.

### GetBountyOk

`func (o *EngagementPublicDetail) GetBountyOk() (*[]EngagementPublicBounty, bool)`

GetBountyOk returns a tuple with the Bounty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBounty

`func (o *EngagementPublicDetail) SetBounty(v []EngagementPublicBounty)`

SetBounty sets Bounty field to given value.

### HasBounty

`func (o *EngagementPublicDetail) HasBounty() bool`

HasBounty returns a boolean if a field has been set.

### GetBriefingAccessible

`func (o *EngagementPublicDetail) GetBriefingAccessible() bool`

GetBriefingAccessible returns the BriefingAccessible field if non-nil, zero value otherwise.

### GetBriefingAccessibleOk

`func (o *EngagementPublicDetail) GetBriefingAccessibleOk() (*bool, bool)`

GetBriefingAccessibleOk returns a tuple with the BriefingAccessible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBriefingAccessible

`func (o *EngagementPublicDetail) SetBriefingAccessible(v bool)`

SetBriefingAccessible sets BriefingAccessible field to given value.


### GetEndsOn

`func (o *EngagementPublicDetail) GetEndsOn() time.Time`

GetEndsOn returns the EndsOn field if non-nil, zero value otherwise.

### GetEndsOnOk

`func (o *EngagementPublicDetail) GetEndsOnOk() (*time.Time, bool)`

GetEndsOnOk returns a tuple with the EndsOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndsOn

`func (o *EngagementPublicDetail) SetEndsOn(v time.Time)`

SetEndsOn sets EndsOn field to given value.

### HasEndsOn

`func (o *EngagementPublicDetail) HasEndsOn() bool`

HasEndsOn returns a boolean if a field has been set.

### GetKind

`func (o *EngagementPublicDetail) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *EngagementPublicDetail) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *EngagementPublicDetail) SetKind(v string)`

SetKind sets Kind field to given value.


### GetLocked

`func (o *EngagementPublicDetail) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *EngagementPublicDetail) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *EngagementPublicDetail) SetLocked(v bool)`

SetLocked sets Locked field to given value.


### GetLogoUrl

`func (o *EngagementPublicDetail) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *EngagementPublicDetail) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *EngagementPublicDetail) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *EngagementPublicDetail) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### GetNdaRequired

`func (o *EngagementPublicDetail) GetNdaRequired() bool`

GetNdaRequired returns the NdaRequired field if non-nil, zero value otherwise.

### GetNdaRequiredOk

`func (o *EngagementPublicDetail) GetNdaRequiredOk() (*bool, bool)`

GetNdaRequiredOk returns a tuple with the NdaRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNdaRequired

`func (o *EngagementPublicDetail) SetNdaRequired(v bool)`

SetNdaRequired sets NdaRequired field to given value.


### GetOrgAbout

`func (o *EngagementPublicDetail) GetOrgAbout() string`

GetOrgAbout returns the OrgAbout field if non-nil, zero value otherwise.

### GetOrgAboutOk

`func (o *EngagementPublicDetail) GetOrgAboutOk() (*string, bool)`

GetOrgAboutOk returns a tuple with the OrgAbout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgAbout

`func (o *EngagementPublicDetail) SetOrgAbout(v string)`

SetOrgAbout sets OrgAbout field to given value.

### HasOrgAbout

`func (o *EngagementPublicDetail) HasOrgAbout() bool`

HasOrgAbout returns a boolean if a field has been set.

### GetOrgName

`func (o *EngagementPublicDetail) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *EngagementPublicDetail) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *EngagementPublicDetail) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.


### GetOrgSlug

`func (o *EngagementPublicDetail) GetOrgSlug() string`

GetOrgSlug returns the OrgSlug field if non-nil, zero value otherwise.

### GetOrgSlugOk

`func (o *EngagementPublicDetail) GetOrgSlugOk() (*string, bool)`

GetOrgSlugOk returns a tuple with the OrgSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgSlug

`func (o *EngagementPublicDetail) SetOrgSlug(v string)`

SetOrgSlug sets OrgSlug field to given value.


### GetProgram

`func (o *EngagementPublicDetail) GetProgram() EngagementProgram`

GetProgram returns the Program field if non-nil, zero value otherwise.

### GetProgramOk

`func (o *EngagementPublicDetail) GetProgramOk() (*EngagementProgram, bool)`

GetProgramOk returns a tuple with the Program field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgram

`func (o *EngagementPublicDetail) SetProgram(v EngagementProgram)`

SetProgram sets Program field to given value.

### HasProgram

`func (o *EngagementPublicDetail) HasProgram() bool`

HasProgram returns a boolean if a field has been set.

### GetRequiredVerificationLevel

`func (o *EngagementPublicDetail) GetRequiredVerificationLevel() string`

GetRequiredVerificationLevel returns the RequiredVerificationLevel field if non-nil, zero value otherwise.

### GetRequiredVerificationLevelOk

`func (o *EngagementPublicDetail) GetRequiredVerificationLevelOk() (*string, bool)`

GetRequiredVerificationLevelOk returns a tuple with the RequiredVerificationLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredVerificationLevel

`func (o *EngagementPublicDetail) SetRequiredVerificationLevel(v string)`

SetRequiredVerificationLevel sets RequiredVerificationLevel field to given value.


### GetScope

`func (o *EngagementPublicDetail) GetScope() []EngagementScopeItem`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *EngagementPublicDetail) GetScopeOk() (*[]EngagementScopeItem, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *EngagementPublicDetail) SetScope(v []EngagementScopeItem)`

SetScope sets Scope field to given value.

### HasScope

`func (o *EngagementPublicDetail) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetSignInUrl

`func (o *EngagementPublicDetail) GetSignInUrl() string`

GetSignInUrl returns the SignInUrl field if non-nil, zero value otherwise.

### GetSignInUrlOk

`func (o *EngagementPublicDetail) GetSignInUrlOk() (*string, bool)`

GetSignInUrlOk returns a tuple with the SignInUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignInUrl

`func (o *EngagementPublicDetail) SetSignInUrl(v string)`

SetSignInUrl sets SignInUrl field to given value.


### GetSlug

`func (o *EngagementPublicDetail) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *EngagementPublicDetail) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *EngagementPublicDetail) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetStartsOn

`func (o *EngagementPublicDetail) GetStartsOn() time.Time`

GetStartsOn returns the StartsOn field if non-nil, zero value otherwise.

### GetStartsOnOk

`func (o *EngagementPublicDetail) GetStartsOnOk() (*time.Time, bool)`

GetStartsOnOk returns a tuple with the StartsOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsOn

`func (o *EngagementPublicDetail) SetStartsOn(v time.Time)`

SetStartsOn sets StartsOn field to given value.

### HasStartsOn

`func (o *EngagementPublicDetail) HasStartsOn() bool`

HasStartsOn returns a boolean if a field has been set.

### GetStatus

`func (o *EngagementPublicDetail) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EngagementPublicDetail) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EngagementPublicDetail) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSummary

`func (o *EngagementPublicDetail) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *EngagementPublicDetail) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *EngagementPublicDetail) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *EngagementPublicDetail) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetTitle

`func (o *EngagementPublicDetail) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *EngagementPublicDetail) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *EngagementPublicDetail) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


