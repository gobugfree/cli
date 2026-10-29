# RevisionRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Author** | Pointer to **string** |  | [optional] 
**AuthoredVia** | Pointer to **string** |  | [optional] 
**ChangeReason** | Pointer to **string** |  | [optional] 
**Content** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**Id** | **string** |  | 
**Number** | **int32** |  | 
**ParentVersion** | Pointer to **int32** |  | [optional] 
**PublishedAt** | Pointer to **time.Time** |  | [optional] 
**SchemaVersion** | **int32** |  | 
**Status** | **string** |  | 
**StructuredContent** | **[]int32** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewRevisionRow

`func NewRevisionRow(content string, createdAt time.Time, id string, number int32, schemaVersion int32, status string, structuredContent []int32, updatedAt time.Time, ) *RevisionRow`

NewRevisionRow instantiates a new RevisionRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRevisionRowWithDefaults

`func NewRevisionRowWithDefaults() *RevisionRow`

NewRevisionRowWithDefaults instantiates a new RevisionRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthor

`func (o *RevisionRow) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *RevisionRow) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *RevisionRow) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *RevisionRow) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetAuthoredVia

`func (o *RevisionRow) GetAuthoredVia() string`

GetAuthoredVia returns the AuthoredVia field if non-nil, zero value otherwise.

### GetAuthoredViaOk

`func (o *RevisionRow) GetAuthoredViaOk() (*string, bool)`

GetAuthoredViaOk returns a tuple with the AuthoredVia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthoredVia

`func (o *RevisionRow) SetAuthoredVia(v string)`

SetAuthoredVia sets AuthoredVia field to given value.

### HasAuthoredVia

`func (o *RevisionRow) HasAuthoredVia() bool`

HasAuthoredVia returns a boolean if a field has been set.

### GetChangeReason

`func (o *RevisionRow) GetChangeReason() string`

GetChangeReason returns the ChangeReason field if non-nil, zero value otherwise.

### GetChangeReasonOk

`func (o *RevisionRow) GetChangeReasonOk() (*string, bool)`

GetChangeReasonOk returns a tuple with the ChangeReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeReason

`func (o *RevisionRow) SetChangeReason(v string)`

SetChangeReason sets ChangeReason field to given value.

### HasChangeReason

`func (o *RevisionRow) HasChangeReason() bool`

HasChangeReason returns a boolean if a field has been set.

### GetContent

`func (o *RevisionRow) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *RevisionRow) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *RevisionRow) SetContent(v string)`

SetContent sets Content field to given value.


### GetCreatedAt

`func (o *RevisionRow) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RevisionRow) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RevisionRow) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetId

`func (o *RevisionRow) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RevisionRow) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RevisionRow) SetId(v string)`

SetId sets Id field to given value.


### GetNumber

`func (o *RevisionRow) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *RevisionRow) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *RevisionRow) SetNumber(v int32)`

SetNumber sets Number field to given value.


### GetParentVersion

`func (o *RevisionRow) GetParentVersion() int32`

GetParentVersion returns the ParentVersion field if non-nil, zero value otherwise.

### GetParentVersionOk

`func (o *RevisionRow) GetParentVersionOk() (*int32, bool)`

GetParentVersionOk returns a tuple with the ParentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentVersion

`func (o *RevisionRow) SetParentVersion(v int32)`

SetParentVersion sets ParentVersion field to given value.

### HasParentVersion

`func (o *RevisionRow) HasParentVersion() bool`

HasParentVersion returns a boolean if a field has been set.

### GetPublishedAt

`func (o *RevisionRow) GetPublishedAt() time.Time`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *RevisionRow) GetPublishedAtOk() (*time.Time, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *RevisionRow) SetPublishedAt(v time.Time)`

SetPublishedAt sets PublishedAt field to given value.

### HasPublishedAt

`func (o *RevisionRow) HasPublishedAt() bool`

HasPublishedAt returns a boolean if a field has been set.

### GetSchemaVersion

`func (o *RevisionRow) GetSchemaVersion() int32`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *RevisionRow) GetSchemaVersionOk() (*int32, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *RevisionRow) SetSchemaVersion(v int32)`

SetSchemaVersion sets SchemaVersion field to given value.


### GetStatus

`func (o *RevisionRow) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RevisionRow) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RevisionRow) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStructuredContent

`func (o *RevisionRow) GetStructuredContent() []int32`

GetStructuredContent returns the StructuredContent field if non-nil, zero value otherwise.

### GetStructuredContentOk

`func (o *RevisionRow) GetStructuredContentOk() (*[]int32, bool)`

GetStructuredContentOk returns a tuple with the StructuredContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructuredContent

`func (o *RevisionRow) SetStructuredContent(v []int32)`

SetStructuredContent sets StructuredContent field to given value.


### GetUpdatedAt

`func (o *RevisionRow) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RevisionRow) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RevisionRow) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


