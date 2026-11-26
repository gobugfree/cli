# DocumentDeliverable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DownloadPath** | Pointer to **string** |  | [optional] 
**Findings** | Pointer to [**DocumentDeliverableFindings**](DocumentDeliverableFindings.md) |  | [optional] 
**Id** | **int32** |  | 
**Name** | **string** |  | 
**ReleasedAt** | Pointer to **time.Time** |  | [optional] 
**ScanStatus** | **string** |  | 
**Size** | **int32** |  | 
**Summary** | Pointer to **string** |  | [optional] 
**UploadedAt** | **time.Time** |  | 

## Methods

### NewDocumentDeliverable

`func NewDocumentDeliverable(id int32, name string, scanStatus string, size int32, uploadedAt time.Time, ) *DocumentDeliverable`

NewDocumentDeliverable instantiates a new DocumentDeliverable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentDeliverableWithDefaults

`func NewDocumentDeliverableWithDefaults() *DocumentDeliverable`

NewDocumentDeliverableWithDefaults instantiates a new DocumentDeliverable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDownloadPath

`func (o *DocumentDeliverable) GetDownloadPath() string`

GetDownloadPath returns the DownloadPath field if non-nil, zero value otherwise.

### GetDownloadPathOk

`func (o *DocumentDeliverable) GetDownloadPathOk() (*string, bool)`

GetDownloadPathOk returns a tuple with the DownloadPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadPath

`func (o *DocumentDeliverable) SetDownloadPath(v string)`

SetDownloadPath sets DownloadPath field to given value.

### HasDownloadPath

`func (o *DocumentDeliverable) HasDownloadPath() bool`

HasDownloadPath returns a boolean if a field has been set.

### GetFindings

`func (o *DocumentDeliverable) GetFindings() DocumentDeliverableFindings`

GetFindings returns the Findings field if non-nil, zero value otherwise.

### GetFindingsOk

`func (o *DocumentDeliverable) GetFindingsOk() (*DocumentDeliverableFindings, bool)`

GetFindingsOk returns a tuple with the Findings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFindings

`func (o *DocumentDeliverable) SetFindings(v DocumentDeliverableFindings)`

SetFindings sets Findings field to given value.

### HasFindings

`func (o *DocumentDeliverable) HasFindings() bool`

HasFindings returns a boolean if a field has been set.

### GetId

`func (o *DocumentDeliverable) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DocumentDeliverable) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DocumentDeliverable) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *DocumentDeliverable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DocumentDeliverable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DocumentDeliverable) SetName(v string)`

SetName sets Name field to given value.


### GetReleasedAt

`func (o *DocumentDeliverable) GetReleasedAt() time.Time`

GetReleasedAt returns the ReleasedAt field if non-nil, zero value otherwise.

### GetReleasedAtOk

`func (o *DocumentDeliverable) GetReleasedAtOk() (*time.Time, bool)`

GetReleasedAtOk returns a tuple with the ReleasedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleasedAt

`func (o *DocumentDeliverable) SetReleasedAt(v time.Time)`

SetReleasedAt sets ReleasedAt field to given value.

### HasReleasedAt

`func (o *DocumentDeliverable) HasReleasedAt() bool`

HasReleasedAt returns a boolean if a field has been set.

### GetScanStatus

`func (o *DocumentDeliverable) GetScanStatus() string`

GetScanStatus returns the ScanStatus field if non-nil, zero value otherwise.

### GetScanStatusOk

`func (o *DocumentDeliverable) GetScanStatusOk() (*string, bool)`

GetScanStatusOk returns a tuple with the ScanStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanStatus

`func (o *DocumentDeliverable) SetScanStatus(v string)`

SetScanStatus sets ScanStatus field to given value.


### GetSize

`func (o *DocumentDeliverable) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *DocumentDeliverable) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *DocumentDeliverable) SetSize(v int32)`

SetSize sets Size field to given value.


### GetSummary

`func (o *DocumentDeliverable) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *DocumentDeliverable) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *DocumentDeliverable) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *DocumentDeliverable) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetUploadedAt

`func (o *DocumentDeliverable) GetUploadedAt() time.Time`

GetUploadedAt returns the UploadedAt field if non-nil, zero value otherwise.

### GetUploadedAtOk

`func (o *DocumentDeliverable) GetUploadedAtOk() (*time.Time, bool)`

GetUploadedAtOk returns a tuple with the UploadedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadedAt

`func (o *DocumentDeliverable) SetUploadedAt(v time.Time)`

SetUploadedAt sets UploadedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


