# OpsTriageWorkspace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetIdentifier** | Pointer to **string** |  | [optional] 
**Comments** | [**[]OpsWorkspaceComment**](OpsWorkspaceComment.md) |  | 
**CvssVector** | Pointer to **string** |  | [optional] 
**Cwe** | Pointer to **string** |  | [optional] 
**Description** | **string** |  | 
**FirstRespondedAt** | Pointer to **time.Time** |  | [optional] 
**Impact** | Pointer to **string** |  | [optional] 
**InfoRequestedAt** | Pointer to **time.Time** |  | [optional] 
**Origin** | **string** |  | 
**Poc** | Pointer to **string** |  | [optional] 

## Methods

### NewOpsTriageWorkspace

`func NewOpsTriageWorkspace(comments []OpsWorkspaceComment, description string, origin string, ) *OpsTriageWorkspace`

NewOpsTriageWorkspace instantiates a new OpsTriageWorkspace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpsTriageWorkspaceWithDefaults

`func NewOpsTriageWorkspaceWithDefaults() *OpsTriageWorkspace`

NewOpsTriageWorkspaceWithDefaults instantiates a new OpsTriageWorkspace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetIdentifier

`func (o *OpsTriageWorkspace) GetAssetIdentifier() string`

GetAssetIdentifier returns the AssetIdentifier field if non-nil, zero value otherwise.

### GetAssetIdentifierOk

`func (o *OpsTriageWorkspace) GetAssetIdentifierOk() (*string, bool)`

GetAssetIdentifierOk returns a tuple with the AssetIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetIdentifier

`func (o *OpsTriageWorkspace) SetAssetIdentifier(v string)`

SetAssetIdentifier sets AssetIdentifier field to given value.

### HasAssetIdentifier

`func (o *OpsTriageWorkspace) HasAssetIdentifier() bool`

HasAssetIdentifier returns a boolean if a field has been set.

### GetComments

`func (o *OpsTriageWorkspace) GetComments() []OpsWorkspaceComment`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *OpsTriageWorkspace) GetCommentsOk() (*[]OpsWorkspaceComment, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *OpsTriageWorkspace) SetComments(v []OpsWorkspaceComment)`

SetComments sets Comments field to given value.


### GetCvssVector

`func (o *OpsTriageWorkspace) GetCvssVector() string`

GetCvssVector returns the CvssVector field if non-nil, zero value otherwise.

### GetCvssVectorOk

`func (o *OpsTriageWorkspace) GetCvssVectorOk() (*string, bool)`

GetCvssVectorOk returns a tuple with the CvssVector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssVector

`func (o *OpsTriageWorkspace) SetCvssVector(v string)`

SetCvssVector sets CvssVector field to given value.

### HasCvssVector

`func (o *OpsTriageWorkspace) HasCvssVector() bool`

HasCvssVector returns a boolean if a field has been set.

### GetCwe

`func (o *OpsTriageWorkspace) GetCwe() string`

GetCwe returns the Cwe field if non-nil, zero value otherwise.

### GetCweOk

`func (o *OpsTriageWorkspace) GetCweOk() (*string, bool)`

GetCweOk returns a tuple with the Cwe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCwe

`func (o *OpsTriageWorkspace) SetCwe(v string)`

SetCwe sets Cwe field to given value.

### HasCwe

`func (o *OpsTriageWorkspace) HasCwe() bool`

HasCwe returns a boolean if a field has been set.

### GetDescription

`func (o *OpsTriageWorkspace) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OpsTriageWorkspace) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OpsTriageWorkspace) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetFirstRespondedAt

`func (o *OpsTriageWorkspace) GetFirstRespondedAt() time.Time`

GetFirstRespondedAt returns the FirstRespondedAt field if non-nil, zero value otherwise.

### GetFirstRespondedAtOk

`func (o *OpsTriageWorkspace) GetFirstRespondedAtOk() (*time.Time, bool)`

GetFirstRespondedAtOk returns a tuple with the FirstRespondedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstRespondedAt

`func (o *OpsTriageWorkspace) SetFirstRespondedAt(v time.Time)`

SetFirstRespondedAt sets FirstRespondedAt field to given value.

### HasFirstRespondedAt

`func (o *OpsTriageWorkspace) HasFirstRespondedAt() bool`

HasFirstRespondedAt returns a boolean if a field has been set.

### GetImpact

`func (o *OpsTriageWorkspace) GetImpact() string`

GetImpact returns the Impact field if non-nil, zero value otherwise.

### GetImpactOk

`func (o *OpsTriageWorkspace) GetImpactOk() (*string, bool)`

GetImpactOk returns a tuple with the Impact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpact

`func (o *OpsTriageWorkspace) SetImpact(v string)`

SetImpact sets Impact field to given value.

### HasImpact

`func (o *OpsTriageWorkspace) HasImpact() bool`

HasImpact returns a boolean if a field has been set.

### GetInfoRequestedAt

`func (o *OpsTriageWorkspace) GetInfoRequestedAt() time.Time`

GetInfoRequestedAt returns the InfoRequestedAt field if non-nil, zero value otherwise.

### GetInfoRequestedAtOk

`func (o *OpsTriageWorkspace) GetInfoRequestedAtOk() (*time.Time, bool)`

GetInfoRequestedAtOk returns a tuple with the InfoRequestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfoRequestedAt

`func (o *OpsTriageWorkspace) SetInfoRequestedAt(v time.Time)`

SetInfoRequestedAt sets InfoRequestedAt field to given value.

### HasInfoRequestedAt

`func (o *OpsTriageWorkspace) HasInfoRequestedAt() bool`

HasInfoRequestedAt returns a boolean if a field has been set.

### GetOrigin

`func (o *OpsTriageWorkspace) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *OpsTriageWorkspace) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *OpsTriageWorkspace) SetOrigin(v string)`

SetOrigin sets Origin field to given value.


### GetPoc

`func (o *OpsTriageWorkspace) GetPoc() string`

GetPoc returns the Poc field if non-nil, zero value otherwise.

### GetPocOk

`func (o *OpsTriageWorkspace) GetPocOk() (*string, bool)`

GetPocOk returns a tuple with the Poc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoc

`func (o *OpsTriageWorkspace) SetPoc(v string)`

SetPoc sets Poc field to given value.

### HasPoc

`func (o *OpsTriageWorkspace) HasPoc() bool`

HasPoc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


