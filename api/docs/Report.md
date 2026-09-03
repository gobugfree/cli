# Report

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetIdentifier** | Pointer to **string** |  | [optional] 
**CvssScore** | Pointer to **float32** |  | [optional] 
**CvssVector** | Pointer to **string** |  | [optional] 
**EngagementId** | **string** |  | 
**Id** | **string** |  | 
**Origin** | **string** |  | 
**SeveritySubmitted** | Pointer to **string** |  | [optional] 
**Status** | **string** |  | 
**SubmittedAt** | Pointer to **time.Time** |  | [optional] 
**Title** | **string** |  | 

## Methods

### NewReport

`func NewReport(engagementId string, id string, origin string, status string, title string, ) *Report`

NewReport instantiates a new Report object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportWithDefaults

`func NewReportWithDefaults() *Report`

NewReportWithDefaults instantiates a new Report object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetIdentifier

`func (o *Report) GetAssetIdentifier() string`

GetAssetIdentifier returns the AssetIdentifier field if non-nil, zero value otherwise.

### GetAssetIdentifierOk

`func (o *Report) GetAssetIdentifierOk() (*string, bool)`

GetAssetIdentifierOk returns a tuple with the AssetIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetIdentifier

`func (o *Report) SetAssetIdentifier(v string)`

SetAssetIdentifier sets AssetIdentifier field to given value.

### HasAssetIdentifier

`func (o *Report) HasAssetIdentifier() bool`

HasAssetIdentifier returns a boolean if a field has been set.

### GetCvssScore

`func (o *Report) GetCvssScore() float32`

GetCvssScore returns the CvssScore field if non-nil, zero value otherwise.

### GetCvssScoreOk

`func (o *Report) GetCvssScoreOk() (*float32, bool)`

GetCvssScoreOk returns a tuple with the CvssScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssScore

`func (o *Report) SetCvssScore(v float32)`

SetCvssScore sets CvssScore field to given value.

### HasCvssScore

`func (o *Report) HasCvssScore() bool`

HasCvssScore returns a boolean if a field has been set.

### GetCvssVector

`func (o *Report) GetCvssVector() string`

GetCvssVector returns the CvssVector field if non-nil, zero value otherwise.

### GetCvssVectorOk

`func (o *Report) GetCvssVectorOk() (*string, bool)`

GetCvssVectorOk returns a tuple with the CvssVector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssVector

`func (o *Report) SetCvssVector(v string)`

SetCvssVector sets CvssVector field to given value.

### HasCvssVector

`func (o *Report) HasCvssVector() bool`

HasCvssVector returns a boolean if a field has been set.

### GetEngagementId

`func (o *Report) GetEngagementId() string`

GetEngagementId returns the EngagementId field if non-nil, zero value otherwise.

### GetEngagementIdOk

`func (o *Report) GetEngagementIdOk() (*string, bool)`

GetEngagementIdOk returns a tuple with the EngagementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngagementId

`func (o *Report) SetEngagementId(v string)`

SetEngagementId sets EngagementId field to given value.


### GetId

`func (o *Report) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Report) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Report) SetId(v string)`

SetId sets Id field to given value.


### GetOrigin

`func (o *Report) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *Report) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *Report) SetOrigin(v string)`

SetOrigin sets Origin field to given value.


### GetSeveritySubmitted

`func (o *Report) GetSeveritySubmitted() string`

GetSeveritySubmitted returns the SeveritySubmitted field if non-nil, zero value otherwise.

### GetSeveritySubmittedOk

`func (o *Report) GetSeveritySubmittedOk() (*string, bool)`

GetSeveritySubmittedOk returns a tuple with the SeveritySubmitted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeveritySubmitted

`func (o *Report) SetSeveritySubmitted(v string)`

SetSeveritySubmitted sets SeveritySubmitted field to given value.

### HasSeveritySubmitted

`func (o *Report) HasSeveritySubmitted() bool`

HasSeveritySubmitted returns a boolean if a field has been set.

### GetStatus

`func (o *Report) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Report) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Report) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSubmittedAt

`func (o *Report) GetSubmittedAt() time.Time`

GetSubmittedAt returns the SubmittedAt field if non-nil, zero value otherwise.

### GetSubmittedAtOk

`func (o *Report) GetSubmittedAtOk() (*time.Time, bool)`

GetSubmittedAtOk returns a tuple with the SubmittedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedAt

`func (o *Report) SetSubmittedAt(v time.Time)`

SetSubmittedAt sets SubmittedAt field to given value.

### HasSubmittedAt

`func (o *Report) HasSubmittedAt() bool`

HasSubmittedAt returns a boolean if a field has been set.

### GetTitle

`func (o *Report) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Report) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Report) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


