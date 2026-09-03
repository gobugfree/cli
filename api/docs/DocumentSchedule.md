# DocumentSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cadence** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**Delivery** | **string** |  | 
**Format** | **string** |  | 
**Id** | **int32** |  | 
**IsActive** | **bool** |  | 
**LastRun** | Pointer to **time.Time** |  | [optional] 
**PausedReason** | Pointer to **string** |  | [optional] 

## Methods

### NewDocumentSchedule

`func NewDocumentSchedule(cadence string, createdAt time.Time, delivery string, format string, id int32, isActive bool, ) *DocumentSchedule`

NewDocumentSchedule instantiates a new DocumentSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentScheduleWithDefaults

`func NewDocumentScheduleWithDefaults() *DocumentSchedule`

NewDocumentScheduleWithDefaults instantiates a new DocumentSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCadence

`func (o *DocumentSchedule) GetCadence() string`

GetCadence returns the Cadence field if non-nil, zero value otherwise.

### GetCadenceOk

`func (o *DocumentSchedule) GetCadenceOk() (*string, bool)`

GetCadenceOk returns a tuple with the Cadence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCadence

`func (o *DocumentSchedule) SetCadence(v string)`

SetCadence sets Cadence field to given value.


### GetCreatedAt

`func (o *DocumentSchedule) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DocumentSchedule) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DocumentSchedule) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDelivery

`func (o *DocumentSchedule) GetDelivery() string`

GetDelivery returns the Delivery field if non-nil, zero value otherwise.

### GetDeliveryOk

`func (o *DocumentSchedule) GetDeliveryOk() (*string, bool)`

GetDeliveryOk returns a tuple with the Delivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivery

`func (o *DocumentSchedule) SetDelivery(v string)`

SetDelivery sets Delivery field to given value.


### GetFormat

`func (o *DocumentSchedule) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *DocumentSchedule) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *DocumentSchedule) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetId

`func (o *DocumentSchedule) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DocumentSchedule) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DocumentSchedule) SetId(v int32)`

SetId sets Id field to given value.


### GetIsActive

`func (o *DocumentSchedule) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *DocumentSchedule) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *DocumentSchedule) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetLastRun

`func (o *DocumentSchedule) GetLastRun() time.Time`

GetLastRun returns the LastRun field if non-nil, zero value otherwise.

### GetLastRunOk

`func (o *DocumentSchedule) GetLastRunOk() (*time.Time, bool)`

GetLastRunOk returns a tuple with the LastRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRun

`func (o *DocumentSchedule) SetLastRun(v time.Time)`

SetLastRun sets LastRun field to given value.

### HasLastRun

`func (o *DocumentSchedule) HasLastRun() bool`

HasLastRun returns a boolean if a field has been set.

### GetPausedReason

`func (o *DocumentSchedule) GetPausedReason() string`

GetPausedReason returns the PausedReason field if non-nil, zero value otherwise.

### GetPausedReasonOk

`func (o *DocumentSchedule) GetPausedReasonOk() (*string, bool)`

GetPausedReasonOk returns a tuple with the PausedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPausedReason

`func (o *DocumentSchedule) SetPausedReason(v string)`

SetPausedReason sets PausedReason field to given value.

### HasPausedReason

`func (o *DocumentSchedule) HasPausedReason() bool`

HasPausedReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


