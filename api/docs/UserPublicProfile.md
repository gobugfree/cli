# UserPublicProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Badges** | [**[]UserProfileBadge**](UserProfileBadge.md) |  | 
**Bio** | **string** |  | 
**Certifications** | [**[]UserPublicCertification**](UserPublicCertification.md) |  | 
**Country** | Pointer to **string** |  | [optional] 
**DisplayName** | **string** |  | 
**ExternalProfiles** | [**[]UserPublicExternal**](UserPublicExternal.md) |  | 
**Handle** | **string** |  | 
**JoinedAt** | **time.Time** |  | 
**Reputation** | **int32** |  | 
**ReputationHistory** | Pointer to [**[]UserReputationEvent**](UserReputationEvent.md) |  | [optional] 
**Skills** | [**[]UserPublicSkill**](UserPublicSkill.md) |  | 
**Tier** | **string** |  | 
**VerificationLevel** | **string** |  | 
**Website** | Pointer to **string** |  | [optional] 

## Methods

### NewUserPublicProfile

`func NewUserPublicProfile(badges []UserProfileBadge, bio string, certifications []UserPublicCertification, displayName string, externalProfiles []UserPublicExternal, handle string, joinedAt time.Time, reputation int32, skills []UserPublicSkill, tier string, verificationLevel string, ) *UserPublicProfile`

NewUserPublicProfile instantiates a new UserPublicProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserPublicProfileWithDefaults

`func NewUserPublicProfileWithDefaults() *UserPublicProfile`

NewUserPublicProfileWithDefaults instantiates a new UserPublicProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBadges

`func (o *UserPublicProfile) GetBadges() []UserProfileBadge`

GetBadges returns the Badges field if non-nil, zero value otherwise.

### GetBadgesOk

`func (o *UserPublicProfile) GetBadgesOk() (*[]UserProfileBadge, bool)`

GetBadgesOk returns a tuple with the Badges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadges

`func (o *UserPublicProfile) SetBadges(v []UserProfileBadge)`

SetBadges sets Badges field to given value.


### GetBio

`func (o *UserPublicProfile) GetBio() string`

GetBio returns the Bio field if non-nil, zero value otherwise.

### GetBioOk

`func (o *UserPublicProfile) GetBioOk() (*string, bool)`

GetBioOk returns a tuple with the Bio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBio

`func (o *UserPublicProfile) SetBio(v string)`

SetBio sets Bio field to given value.


### GetCertifications

`func (o *UserPublicProfile) GetCertifications() []UserPublicCertification`

GetCertifications returns the Certifications field if non-nil, zero value otherwise.

### GetCertificationsOk

`func (o *UserPublicProfile) GetCertificationsOk() (*[]UserPublicCertification, bool)`

GetCertificationsOk returns a tuple with the Certifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertifications

`func (o *UserPublicProfile) SetCertifications(v []UserPublicCertification)`

SetCertifications sets Certifications field to given value.


### GetCountry

`func (o *UserPublicProfile) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *UserPublicProfile) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *UserPublicProfile) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *UserPublicProfile) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetDisplayName

`func (o *UserPublicProfile) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UserPublicProfile) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UserPublicProfile) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetExternalProfiles

`func (o *UserPublicProfile) GetExternalProfiles() []UserPublicExternal`

GetExternalProfiles returns the ExternalProfiles field if non-nil, zero value otherwise.

### GetExternalProfilesOk

`func (o *UserPublicProfile) GetExternalProfilesOk() (*[]UserPublicExternal, bool)`

GetExternalProfilesOk returns a tuple with the ExternalProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalProfiles

`func (o *UserPublicProfile) SetExternalProfiles(v []UserPublicExternal)`

SetExternalProfiles sets ExternalProfiles field to given value.


### GetHandle

`func (o *UserPublicProfile) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *UserPublicProfile) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *UserPublicProfile) SetHandle(v string)`

SetHandle sets Handle field to given value.


### GetJoinedAt

`func (o *UserPublicProfile) GetJoinedAt() time.Time`

GetJoinedAt returns the JoinedAt field if non-nil, zero value otherwise.

### GetJoinedAtOk

`func (o *UserPublicProfile) GetJoinedAtOk() (*time.Time, bool)`

GetJoinedAtOk returns a tuple with the JoinedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinedAt

`func (o *UserPublicProfile) SetJoinedAt(v time.Time)`

SetJoinedAt sets JoinedAt field to given value.


### GetReputation

`func (o *UserPublicProfile) GetReputation() int32`

GetReputation returns the Reputation field if non-nil, zero value otherwise.

### GetReputationOk

`func (o *UserPublicProfile) GetReputationOk() (*int32, bool)`

GetReputationOk returns a tuple with the Reputation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReputation

`func (o *UserPublicProfile) SetReputation(v int32)`

SetReputation sets Reputation field to given value.


### GetReputationHistory

`func (o *UserPublicProfile) GetReputationHistory() []UserReputationEvent`

GetReputationHistory returns the ReputationHistory field if non-nil, zero value otherwise.

### GetReputationHistoryOk

`func (o *UserPublicProfile) GetReputationHistoryOk() (*[]UserReputationEvent, bool)`

GetReputationHistoryOk returns a tuple with the ReputationHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReputationHistory

`func (o *UserPublicProfile) SetReputationHistory(v []UserReputationEvent)`

SetReputationHistory sets ReputationHistory field to given value.

### HasReputationHistory

`func (o *UserPublicProfile) HasReputationHistory() bool`

HasReputationHistory returns a boolean if a field has been set.

### GetSkills

`func (o *UserPublicProfile) GetSkills() []UserPublicSkill`

GetSkills returns the Skills field if non-nil, zero value otherwise.

### GetSkillsOk

`func (o *UserPublicProfile) GetSkillsOk() (*[]UserPublicSkill, bool)`

GetSkillsOk returns a tuple with the Skills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkills

`func (o *UserPublicProfile) SetSkills(v []UserPublicSkill)`

SetSkills sets Skills field to given value.


### GetTier

`func (o *UserPublicProfile) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *UserPublicProfile) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *UserPublicProfile) SetTier(v string)`

SetTier sets Tier field to given value.


### GetVerificationLevel

`func (o *UserPublicProfile) GetVerificationLevel() string`

GetVerificationLevel returns the VerificationLevel field if non-nil, zero value otherwise.

### GetVerificationLevelOk

`func (o *UserPublicProfile) GetVerificationLevelOk() (*string, bool)`

GetVerificationLevelOk returns a tuple with the VerificationLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationLevel

`func (o *UserPublicProfile) SetVerificationLevel(v string)`

SetVerificationLevel sets VerificationLevel field to given value.


### GetWebsite

`func (o *UserPublicProfile) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *UserPublicProfile) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *UserPublicProfile) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *UserPublicProfile) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


