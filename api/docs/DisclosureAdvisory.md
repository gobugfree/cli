# DisclosureAdvisory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffectedAsset** | Pointer to **string** |  | [optional] 
**Cve** | Pointer to **string** |  | [optional] 
**DisclosureId** | **string** |  | 
**Organization** | **string** |  | 
**PublishedOn** | Pointer to **time.Time** |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**Summary** | **string** |  | 
**Title** | **string** |  | 

## Methods

### NewDisclosureAdvisory

`func NewDisclosureAdvisory(disclosureId string, organization string, summary string, title string, ) *DisclosureAdvisory`

NewDisclosureAdvisory instantiates a new DisclosureAdvisory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDisclosureAdvisoryWithDefaults

`func NewDisclosureAdvisoryWithDefaults() *DisclosureAdvisory`

NewDisclosureAdvisoryWithDefaults instantiates a new DisclosureAdvisory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffectedAsset

`func (o *DisclosureAdvisory) GetAffectedAsset() string`

GetAffectedAsset returns the AffectedAsset field if non-nil, zero value otherwise.

### GetAffectedAssetOk

`func (o *DisclosureAdvisory) GetAffectedAssetOk() (*string, bool)`

GetAffectedAssetOk returns a tuple with the AffectedAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedAsset

`func (o *DisclosureAdvisory) SetAffectedAsset(v string)`

SetAffectedAsset sets AffectedAsset field to given value.

### HasAffectedAsset

`func (o *DisclosureAdvisory) HasAffectedAsset() bool`

HasAffectedAsset returns a boolean if a field has been set.

### GetCve

`func (o *DisclosureAdvisory) GetCve() string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *DisclosureAdvisory) GetCveOk() (*string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *DisclosureAdvisory) SetCve(v string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *DisclosureAdvisory) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDisclosureId

`func (o *DisclosureAdvisory) GetDisclosureId() string`

GetDisclosureId returns the DisclosureId field if non-nil, zero value otherwise.

### GetDisclosureIdOk

`func (o *DisclosureAdvisory) GetDisclosureIdOk() (*string, bool)`

GetDisclosureIdOk returns a tuple with the DisclosureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclosureId

`func (o *DisclosureAdvisory) SetDisclosureId(v string)`

SetDisclosureId sets DisclosureId field to given value.


### GetOrganization

`func (o *DisclosureAdvisory) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *DisclosureAdvisory) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *DisclosureAdvisory) SetOrganization(v string)`

SetOrganization sets Organization field to given value.


### GetPublishedOn

`func (o *DisclosureAdvisory) GetPublishedOn() time.Time`

GetPublishedOn returns the PublishedOn field if non-nil, zero value otherwise.

### GetPublishedOnOk

`func (o *DisclosureAdvisory) GetPublishedOnOk() (*time.Time, bool)`

GetPublishedOnOk returns a tuple with the PublishedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedOn

`func (o *DisclosureAdvisory) SetPublishedOn(v time.Time)`

SetPublishedOn sets PublishedOn field to given value.

### HasPublishedOn

`func (o *DisclosureAdvisory) HasPublishedOn() bool`

HasPublishedOn returns a boolean if a field has been set.

### GetSeverity

`func (o *DisclosureAdvisory) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *DisclosureAdvisory) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *DisclosureAdvisory) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *DisclosureAdvisory) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetSummary

`func (o *DisclosureAdvisory) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *DisclosureAdvisory) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *DisclosureAdvisory) SetSummary(v string)`

SetSummary sets Summary field to given value.


### GetTitle

`func (o *DisclosureAdvisory) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DisclosureAdvisory) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DisclosureAdvisory) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


