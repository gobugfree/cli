# ReportEvidence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Caption** | Pointer to **string** |  | [optional] 
**Id** | **int32** |  | 
**ScanStatus** | **string** |  | 

## Methods

### NewReportEvidence

`func NewReportEvidence(id int32, scanStatus string, ) *ReportEvidence`

NewReportEvidence instantiates a new ReportEvidence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportEvidenceWithDefaults

`func NewReportEvidenceWithDefaults() *ReportEvidence`

NewReportEvidenceWithDefaults instantiates a new ReportEvidence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaption

`func (o *ReportEvidence) GetCaption() string`

GetCaption returns the Caption field if non-nil, zero value otherwise.

### GetCaptionOk

`func (o *ReportEvidence) GetCaptionOk() (*string, bool)`

GetCaptionOk returns a tuple with the Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaption

`func (o *ReportEvidence) SetCaption(v string)`

SetCaption sets Caption field to given value.

### HasCaption

`func (o *ReportEvidence) HasCaption() bool`

HasCaption returns a boolean if a field has been set.

### GetId

`func (o *ReportEvidence) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReportEvidence) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReportEvidence) SetId(v int32)`

SetId sets Id field to given value.


### GetScanStatus

`func (o *ReportEvidence) GetScanStatus() string`

GetScanStatus returns the ScanStatus field if non-nil, zero value otherwise.

### GetScanStatusOk

`func (o *ReportEvidence) GetScanStatusOk() (*string, bool)`

GetScanStatusOk returns a tuple with the ScanStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanStatus

`func (o *ReportEvidence) SetScanStatus(v string)`

SetScanStatus sets ScanStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


