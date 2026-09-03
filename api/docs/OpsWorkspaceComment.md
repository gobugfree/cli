# OpsWorkspaceComment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Author** | **string** |  | 
**Body** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**Visibility** | **string** |  | 

## Methods

### NewOpsWorkspaceComment

`func NewOpsWorkspaceComment(author string, body string, createdAt time.Time, visibility string, ) *OpsWorkspaceComment`

NewOpsWorkspaceComment instantiates a new OpsWorkspaceComment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpsWorkspaceCommentWithDefaults

`func NewOpsWorkspaceCommentWithDefaults() *OpsWorkspaceComment`

NewOpsWorkspaceCommentWithDefaults instantiates a new OpsWorkspaceComment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthor

`func (o *OpsWorkspaceComment) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *OpsWorkspaceComment) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *OpsWorkspaceComment) SetAuthor(v string)`

SetAuthor sets Author field to given value.


### GetBody

`func (o *OpsWorkspaceComment) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *OpsWorkspaceComment) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *OpsWorkspaceComment) SetBody(v string)`

SetBody sets Body field to given value.


### GetCreatedAt

`func (o *OpsWorkspaceComment) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OpsWorkspaceComment) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OpsWorkspaceComment) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetVisibility

`func (o *OpsWorkspaceComment) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *OpsWorkspaceComment) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *OpsWorkspaceComment) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


