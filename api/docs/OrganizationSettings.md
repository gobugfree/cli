# OrganizationSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BrandColor** | Pointer to **string** |  | [optional] 
**BrandColorSecondary** | Pointer to **string** |  | [optional] 
**Draft** | Pointer to [**OrganizationProfileVersion**](OrganizationProfileVersion.md) |  | [optional] 
**Industry** | Pointer to **string** |  | [optional] 
**LogoPath** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Slug** | **string** |  | 
**Website** | Pointer to **string** |  | [optional] 

## Methods

### NewOrganizationSettings

`func NewOrganizationSettings(name string, slug string, ) *OrganizationSettings`

NewOrganizationSettings instantiates a new OrganizationSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationSettingsWithDefaults

`func NewOrganizationSettingsWithDefaults() *OrganizationSettings`

NewOrganizationSettingsWithDefaults instantiates a new OrganizationSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrandColor

`func (o *OrganizationSettings) GetBrandColor() string`

GetBrandColor returns the BrandColor field if non-nil, zero value otherwise.

### GetBrandColorOk

`func (o *OrganizationSettings) GetBrandColorOk() (*string, bool)`

GetBrandColorOk returns a tuple with the BrandColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandColor

`func (o *OrganizationSettings) SetBrandColor(v string)`

SetBrandColor sets BrandColor field to given value.

### HasBrandColor

`func (o *OrganizationSettings) HasBrandColor() bool`

HasBrandColor returns a boolean if a field has been set.

### GetBrandColorSecondary

`func (o *OrganizationSettings) GetBrandColorSecondary() string`

GetBrandColorSecondary returns the BrandColorSecondary field if non-nil, zero value otherwise.

### GetBrandColorSecondaryOk

`func (o *OrganizationSettings) GetBrandColorSecondaryOk() (*string, bool)`

GetBrandColorSecondaryOk returns a tuple with the BrandColorSecondary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandColorSecondary

`func (o *OrganizationSettings) SetBrandColorSecondary(v string)`

SetBrandColorSecondary sets BrandColorSecondary field to given value.

### HasBrandColorSecondary

`func (o *OrganizationSettings) HasBrandColorSecondary() bool`

HasBrandColorSecondary returns a boolean if a field has been set.

### GetDraft

`func (o *OrganizationSettings) GetDraft() OrganizationProfileVersion`

GetDraft returns the Draft field if non-nil, zero value otherwise.

### GetDraftOk

`func (o *OrganizationSettings) GetDraftOk() (*OrganizationProfileVersion, bool)`

GetDraftOk returns a tuple with the Draft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraft

`func (o *OrganizationSettings) SetDraft(v OrganizationProfileVersion)`

SetDraft sets Draft field to given value.

### HasDraft

`func (o *OrganizationSettings) HasDraft() bool`

HasDraft returns a boolean if a field has been set.

### GetIndustry

`func (o *OrganizationSettings) GetIndustry() string`

GetIndustry returns the Industry field if non-nil, zero value otherwise.

### GetIndustryOk

`func (o *OrganizationSettings) GetIndustryOk() (*string, bool)`

GetIndustryOk returns a tuple with the Industry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustry

`func (o *OrganizationSettings) SetIndustry(v string)`

SetIndustry sets Industry field to given value.

### HasIndustry

`func (o *OrganizationSettings) HasIndustry() bool`

HasIndustry returns a boolean if a field has been set.

### GetLogoPath

`func (o *OrganizationSettings) GetLogoPath() string`

GetLogoPath returns the LogoPath field if non-nil, zero value otherwise.

### GetLogoPathOk

`func (o *OrganizationSettings) GetLogoPathOk() (*string, bool)`

GetLogoPathOk returns a tuple with the LogoPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoPath

`func (o *OrganizationSettings) SetLogoPath(v string)`

SetLogoPath sets LogoPath field to given value.

### HasLogoPath

`func (o *OrganizationSettings) HasLogoPath() bool`

HasLogoPath returns a boolean if a field has been set.

### GetName

`func (o *OrganizationSettings) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationSettings) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationSettings) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *OrganizationSettings) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *OrganizationSettings) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *OrganizationSettings) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetWebsite

`func (o *OrganizationSettings) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *OrganizationSettings) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *OrganizationSettings) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *OrganizationSettings) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


