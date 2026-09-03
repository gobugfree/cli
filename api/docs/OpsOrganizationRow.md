# OpsOrganizationRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveEngagements** | **int32** |  | 
**Created** | **time.Time** |  | 
**Engagements** | **int32** |  | 
**Kind** | **string** |  | 
**Name** | **string** |  | 
**Slug** | **string** |  | 
**VerificationStatus** | **string** |  | 

## Methods

### NewOpsOrganizationRow

`func NewOpsOrganizationRow(activeEngagements int32, created time.Time, engagements int32, kind string, name string, slug string, verificationStatus string, ) *OpsOrganizationRow`

NewOpsOrganizationRow instantiates a new OpsOrganizationRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpsOrganizationRowWithDefaults

`func NewOpsOrganizationRowWithDefaults() *OpsOrganizationRow`

NewOpsOrganizationRowWithDefaults instantiates a new OpsOrganizationRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveEngagements

`func (o *OpsOrganizationRow) GetActiveEngagements() int32`

GetActiveEngagements returns the ActiveEngagements field if non-nil, zero value otherwise.

### GetActiveEngagementsOk

`func (o *OpsOrganizationRow) GetActiveEngagementsOk() (*int32, bool)`

GetActiveEngagementsOk returns a tuple with the ActiveEngagements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveEngagements

`func (o *OpsOrganizationRow) SetActiveEngagements(v int32)`

SetActiveEngagements sets ActiveEngagements field to given value.


### GetCreated

`func (o *OpsOrganizationRow) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *OpsOrganizationRow) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *OpsOrganizationRow) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetEngagements

`func (o *OpsOrganizationRow) GetEngagements() int32`

GetEngagements returns the Engagements field if non-nil, zero value otherwise.

### GetEngagementsOk

`func (o *OpsOrganizationRow) GetEngagementsOk() (*int32, bool)`

GetEngagementsOk returns a tuple with the Engagements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngagements

`func (o *OpsOrganizationRow) SetEngagements(v int32)`

SetEngagements sets Engagements field to given value.


### GetKind

`func (o *OpsOrganizationRow) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *OpsOrganizationRow) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *OpsOrganizationRow) SetKind(v string)`

SetKind sets Kind field to given value.


### GetName

`func (o *OpsOrganizationRow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OpsOrganizationRow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OpsOrganizationRow) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *OpsOrganizationRow) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *OpsOrganizationRow) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *OpsOrganizationRow) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetVerificationStatus

`func (o *OpsOrganizationRow) GetVerificationStatus() string`

GetVerificationStatus returns the VerificationStatus field if non-nil, zero value otherwise.

### GetVerificationStatusOk

`func (o *OpsOrganizationRow) GetVerificationStatusOk() (*string, bool)`

GetVerificationStatusOk returns a tuple with the VerificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationStatus

`func (o *OpsOrganizationRow) SetVerificationStatus(v string)`

SetVerificationStatus sets VerificationStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


