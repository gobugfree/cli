# DocumentVault

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Documents** | [**[]Document**](Document.md) |  | 
**Exports** | [**[]DocumentExport**](DocumentExport.md) |  | 
**Schedules** | [**[]DocumentSchedule**](DocumentSchedule.md) |  | 

## Methods

### NewDocumentVault

`func NewDocumentVault(documents []Document, exports []DocumentExport, schedules []DocumentSchedule, ) *DocumentVault`

NewDocumentVault instantiates a new DocumentVault object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentVaultWithDefaults

`func NewDocumentVaultWithDefaults() *DocumentVault`

NewDocumentVaultWithDefaults instantiates a new DocumentVault object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocuments

`func (o *DocumentVault) GetDocuments() []Document`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *DocumentVault) GetDocumentsOk() (*[]Document, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *DocumentVault) SetDocuments(v []Document)`

SetDocuments sets Documents field to given value.


### GetExports

`func (o *DocumentVault) GetExports() []DocumentExport`

GetExports returns the Exports field if non-nil, zero value otherwise.

### GetExportsOk

`func (o *DocumentVault) GetExportsOk() (*[]DocumentExport, bool)`

GetExportsOk returns a tuple with the Exports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExports

`func (o *DocumentVault) SetExports(v []DocumentExport)`

SetExports sets Exports field to given value.


### GetSchedules

`func (o *DocumentVault) GetSchedules() []DocumentSchedule`

GetSchedules returns the Schedules field if non-nil, zero value otherwise.

### GetSchedulesOk

`func (o *DocumentVault) GetSchedulesOk() (*[]DocumentSchedule, bool)`

GetSchedulesOk returns a tuple with the Schedules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedules

`func (o *DocumentVault) SetSchedules(v []DocumentSchedule)`

SetSchedules sets Schedules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


