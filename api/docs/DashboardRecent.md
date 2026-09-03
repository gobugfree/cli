# DashboardRecent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Engagement** | **string** |  | 
**Id** | **string** |  | 
**Severity** | **string** |  | 
**SubmittedAt** | Pointer to **time.Time** |  | [optional] 
**Title** | **string** |  | 

## Methods

### NewDashboardRecent

`func NewDashboardRecent(engagement string, id string, severity string, title string, ) *DashboardRecent`

NewDashboardRecent instantiates a new DashboardRecent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardRecentWithDefaults

`func NewDashboardRecentWithDefaults() *DashboardRecent`

NewDashboardRecentWithDefaults instantiates a new DashboardRecent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEngagement

`func (o *DashboardRecent) GetEngagement() string`

GetEngagement returns the Engagement field if non-nil, zero value otherwise.

### GetEngagementOk

`func (o *DashboardRecent) GetEngagementOk() (*string, bool)`

GetEngagementOk returns a tuple with the Engagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngagement

`func (o *DashboardRecent) SetEngagement(v string)`

SetEngagement sets Engagement field to given value.


### GetId

`func (o *DashboardRecent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DashboardRecent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DashboardRecent) SetId(v string)`

SetId sets Id field to given value.


### GetSeverity

`func (o *DashboardRecent) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *DashboardRecent) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *DashboardRecent) SetSeverity(v string)`

SetSeverity sets Severity field to given value.


### GetSubmittedAt

`func (o *DashboardRecent) GetSubmittedAt() time.Time`

GetSubmittedAt returns the SubmittedAt field if non-nil, zero value otherwise.

### GetSubmittedAtOk

`func (o *DashboardRecent) GetSubmittedAtOk() (*time.Time, bool)`

GetSubmittedAtOk returns a tuple with the SubmittedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedAt

`func (o *DashboardRecent) SetSubmittedAt(v time.Time)`

SetSubmittedAt sets SubmittedAt field to given value.

### HasSubmittedAt

`func (o *DashboardRecent) HasSubmittedAt() bool`

HasSubmittedAt returns a boolean if a field has been set.

### GetTitle

`func (o *DashboardRecent) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DashboardRecent) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DashboardRecent) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


