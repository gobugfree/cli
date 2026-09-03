# DashboardOrgDashboard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivePrograms** | **int32** |  | 
**BountySpendCents** | **int32** |  | 
**BySeverity** | **map[string]int32** |  | 
**Currency** | **string** |  | 
**MedianTimeToResolution** | Pointer to **string** |  | [optional] 
**MedianTimeToTriage** | Pointer to **string** |  | [optional] 
**NewLast30Days** | **int32** |  | 
**OpenCriticalHigh** | **int32** |  | 
**OpenTotal** | **int32** |  | 
**RecentlyReported** | [**[]DashboardRecent**](DashboardRecent.md) |  | 
**ReopenedOpen** | **int32** |  | 
**Researchers** | **int32** |  | 
**ResolvedTotal** | **int32** |  | 
**SlaBreached** | **int32** |  | 
**SlaCompliant** | **int32** |  | 
**SlaTracked** | **int32** |  | 

## Methods

### NewDashboardOrgDashboard

`func NewDashboardOrgDashboard(activePrograms int32, bountySpendCents int32, bySeverity map[string]int32, currency string, newLast30Days int32, openCriticalHigh int32, openTotal int32, recentlyReported []DashboardRecent, reopenedOpen int32, researchers int32, resolvedTotal int32, slaBreached int32, slaCompliant int32, slaTracked int32, ) *DashboardOrgDashboard`

NewDashboardOrgDashboard instantiates a new DashboardOrgDashboard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardOrgDashboardWithDefaults

`func NewDashboardOrgDashboardWithDefaults() *DashboardOrgDashboard`

NewDashboardOrgDashboardWithDefaults instantiates a new DashboardOrgDashboard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivePrograms

`func (o *DashboardOrgDashboard) GetActivePrograms() int32`

GetActivePrograms returns the ActivePrograms field if non-nil, zero value otherwise.

### GetActiveProgramsOk

`func (o *DashboardOrgDashboard) GetActiveProgramsOk() (*int32, bool)`

GetActiveProgramsOk returns a tuple with the ActivePrograms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivePrograms

`func (o *DashboardOrgDashboard) SetActivePrograms(v int32)`

SetActivePrograms sets ActivePrograms field to given value.


### GetBountySpendCents

`func (o *DashboardOrgDashboard) GetBountySpendCents() int32`

GetBountySpendCents returns the BountySpendCents field if non-nil, zero value otherwise.

### GetBountySpendCentsOk

`func (o *DashboardOrgDashboard) GetBountySpendCentsOk() (*int32, bool)`

GetBountySpendCentsOk returns a tuple with the BountySpendCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBountySpendCents

`func (o *DashboardOrgDashboard) SetBountySpendCents(v int32)`

SetBountySpendCents sets BountySpendCents field to given value.


### GetBySeverity

`func (o *DashboardOrgDashboard) GetBySeverity() map[string]int32`

GetBySeverity returns the BySeverity field if non-nil, zero value otherwise.

### GetBySeverityOk

`func (o *DashboardOrgDashboard) GetBySeverityOk() (*map[string]int32, bool)`

GetBySeverityOk returns a tuple with the BySeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBySeverity

`func (o *DashboardOrgDashboard) SetBySeverity(v map[string]int32)`

SetBySeverity sets BySeverity field to given value.


### GetCurrency

`func (o *DashboardOrgDashboard) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *DashboardOrgDashboard) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *DashboardOrgDashboard) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetMedianTimeToResolution

`func (o *DashboardOrgDashboard) GetMedianTimeToResolution() string`

GetMedianTimeToResolution returns the MedianTimeToResolution field if non-nil, zero value otherwise.

### GetMedianTimeToResolutionOk

`func (o *DashboardOrgDashboard) GetMedianTimeToResolutionOk() (*string, bool)`

GetMedianTimeToResolutionOk returns a tuple with the MedianTimeToResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedianTimeToResolution

`func (o *DashboardOrgDashboard) SetMedianTimeToResolution(v string)`

SetMedianTimeToResolution sets MedianTimeToResolution field to given value.

### HasMedianTimeToResolution

`func (o *DashboardOrgDashboard) HasMedianTimeToResolution() bool`

HasMedianTimeToResolution returns a boolean if a field has been set.

### GetMedianTimeToTriage

`func (o *DashboardOrgDashboard) GetMedianTimeToTriage() string`

GetMedianTimeToTriage returns the MedianTimeToTriage field if non-nil, zero value otherwise.

### GetMedianTimeToTriageOk

`func (o *DashboardOrgDashboard) GetMedianTimeToTriageOk() (*string, bool)`

GetMedianTimeToTriageOk returns a tuple with the MedianTimeToTriage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedianTimeToTriage

`func (o *DashboardOrgDashboard) SetMedianTimeToTriage(v string)`

SetMedianTimeToTriage sets MedianTimeToTriage field to given value.

### HasMedianTimeToTriage

`func (o *DashboardOrgDashboard) HasMedianTimeToTriage() bool`

HasMedianTimeToTriage returns a boolean if a field has been set.

### GetNewLast30Days

`func (o *DashboardOrgDashboard) GetNewLast30Days() int32`

GetNewLast30Days returns the NewLast30Days field if non-nil, zero value otherwise.

### GetNewLast30DaysOk

`func (o *DashboardOrgDashboard) GetNewLast30DaysOk() (*int32, bool)`

GetNewLast30DaysOk returns a tuple with the NewLast30Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewLast30Days

`func (o *DashboardOrgDashboard) SetNewLast30Days(v int32)`

SetNewLast30Days sets NewLast30Days field to given value.


### GetOpenCriticalHigh

`func (o *DashboardOrgDashboard) GetOpenCriticalHigh() int32`

GetOpenCriticalHigh returns the OpenCriticalHigh field if non-nil, zero value otherwise.

### GetOpenCriticalHighOk

`func (o *DashboardOrgDashboard) GetOpenCriticalHighOk() (*int32, bool)`

GetOpenCriticalHighOk returns a tuple with the OpenCriticalHigh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenCriticalHigh

`func (o *DashboardOrgDashboard) SetOpenCriticalHigh(v int32)`

SetOpenCriticalHigh sets OpenCriticalHigh field to given value.


### GetOpenTotal

`func (o *DashboardOrgDashboard) GetOpenTotal() int32`

GetOpenTotal returns the OpenTotal field if non-nil, zero value otherwise.

### GetOpenTotalOk

`func (o *DashboardOrgDashboard) GetOpenTotalOk() (*int32, bool)`

GetOpenTotalOk returns a tuple with the OpenTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenTotal

`func (o *DashboardOrgDashboard) SetOpenTotal(v int32)`

SetOpenTotal sets OpenTotal field to given value.


### GetRecentlyReported

`func (o *DashboardOrgDashboard) GetRecentlyReported() []DashboardRecent`

GetRecentlyReported returns the RecentlyReported field if non-nil, zero value otherwise.

### GetRecentlyReportedOk

`func (o *DashboardOrgDashboard) GetRecentlyReportedOk() (*[]DashboardRecent, bool)`

GetRecentlyReportedOk returns a tuple with the RecentlyReported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentlyReported

`func (o *DashboardOrgDashboard) SetRecentlyReported(v []DashboardRecent)`

SetRecentlyReported sets RecentlyReported field to given value.


### GetReopenedOpen

`func (o *DashboardOrgDashboard) GetReopenedOpen() int32`

GetReopenedOpen returns the ReopenedOpen field if non-nil, zero value otherwise.

### GetReopenedOpenOk

`func (o *DashboardOrgDashboard) GetReopenedOpenOk() (*int32, bool)`

GetReopenedOpenOk returns a tuple with the ReopenedOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReopenedOpen

`func (o *DashboardOrgDashboard) SetReopenedOpen(v int32)`

SetReopenedOpen sets ReopenedOpen field to given value.


### GetResearchers

`func (o *DashboardOrgDashboard) GetResearchers() int32`

GetResearchers returns the Researchers field if non-nil, zero value otherwise.

### GetResearchersOk

`func (o *DashboardOrgDashboard) GetResearchersOk() (*int32, bool)`

GetResearchersOk returns a tuple with the Researchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResearchers

`func (o *DashboardOrgDashboard) SetResearchers(v int32)`

SetResearchers sets Researchers field to given value.


### GetResolvedTotal

`func (o *DashboardOrgDashboard) GetResolvedTotal() int32`

GetResolvedTotal returns the ResolvedTotal field if non-nil, zero value otherwise.

### GetResolvedTotalOk

`func (o *DashboardOrgDashboard) GetResolvedTotalOk() (*int32, bool)`

GetResolvedTotalOk returns a tuple with the ResolvedTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedTotal

`func (o *DashboardOrgDashboard) SetResolvedTotal(v int32)`

SetResolvedTotal sets ResolvedTotal field to given value.


### GetSlaBreached

`func (o *DashboardOrgDashboard) GetSlaBreached() int32`

GetSlaBreached returns the SlaBreached field if non-nil, zero value otherwise.

### GetSlaBreachedOk

`func (o *DashboardOrgDashboard) GetSlaBreachedOk() (*int32, bool)`

GetSlaBreachedOk returns a tuple with the SlaBreached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaBreached

`func (o *DashboardOrgDashboard) SetSlaBreached(v int32)`

SetSlaBreached sets SlaBreached field to given value.


### GetSlaCompliant

`func (o *DashboardOrgDashboard) GetSlaCompliant() int32`

GetSlaCompliant returns the SlaCompliant field if non-nil, zero value otherwise.

### GetSlaCompliantOk

`func (o *DashboardOrgDashboard) GetSlaCompliantOk() (*int32, bool)`

GetSlaCompliantOk returns a tuple with the SlaCompliant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaCompliant

`func (o *DashboardOrgDashboard) SetSlaCompliant(v int32)`

SetSlaCompliant sets SlaCompliant field to given value.


### GetSlaTracked

`func (o *DashboardOrgDashboard) GetSlaTracked() int32`

GetSlaTracked returns the SlaTracked field if non-nil, zero value otherwise.

### GetSlaTrackedOk

`func (o *DashboardOrgDashboard) GetSlaTrackedOk() (*int32, bool)`

GetSlaTrackedOk returns a tuple with the SlaTracked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaTracked

`func (o *DashboardOrgDashboard) SetSlaTracked(v int32)`

SetSlaTracked sets SlaTracked field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


