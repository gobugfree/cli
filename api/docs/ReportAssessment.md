# ReportAssessment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CvssScore** | Pointer to **float32** |  | [optional] 
**CvssVector** | Pointer to **string** |  | [optional] 
**Justification** | Pointer to **string** |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 

## Methods

### NewReportAssessment

`func NewReportAssessment() *ReportAssessment`

NewReportAssessment instantiates a new ReportAssessment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportAssessmentWithDefaults

`func NewReportAssessmentWithDefaults() *ReportAssessment`

NewReportAssessmentWithDefaults instantiates a new ReportAssessment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCvssScore

`func (o *ReportAssessment) GetCvssScore() float32`

GetCvssScore returns the CvssScore field if non-nil, zero value otherwise.

### GetCvssScoreOk

`func (o *ReportAssessment) GetCvssScoreOk() (*float32, bool)`

GetCvssScoreOk returns a tuple with the CvssScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssScore

`func (o *ReportAssessment) SetCvssScore(v float32)`

SetCvssScore sets CvssScore field to given value.

### HasCvssScore

`func (o *ReportAssessment) HasCvssScore() bool`

HasCvssScore returns a boolean if a field has been set.

### GetCvssVector

`func (o *ReportAssessment) GetCvssVector() string`

GetCvssVector returns the CvssVector field if non-nil, zero value otherwise.

### GetCvssVectorOk

`func (o *ReportAssessment) GetCvssVectorOk() (*string, bool)`

GetCvssVectorOk returns a tuple with the CvssVector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssVector

`func (o *ReportAssessment) SetCvssVector(v string)`

SetCvssVector sets CvssVector field to given value.

### HasCvssVector

`func (o *ReportAssessment) HasCvssVector() bool`

HasCvssVector returns a boolean if a field has been set.

### GetJustification

`func (o *ReportAssessment) GetJustification() string`

GetJustification returns the Justification field if non-nil, zero value otherwise.

### GetJustificationOk

`func (o *ReportAssessment) GetJustificationOk() (*string, bool)`

GetJustificationOk returns a tuple with the Justification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJustification

`func (o *ReportAssessment) SetJustification(v string)`

SetJustification sets Justification field to given value.

### HasJustification

`func (o *ReportAssessment) HasJustification() bool`

HasJustification returns a boolean if a field has been set.

### GetSeverity

`func (o *ReportAssessment) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ReportAssessment) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ReportAssessment) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *ReportAssessment) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


