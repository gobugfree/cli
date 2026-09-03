# DashboardEngagementStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BountySpendCents** | **int32** |  | 
**BySeverity** | **map[string]int32** |  | 
**CostLimitCents** | **int32** |  | 
**Currency** | **string** |  | 
**MedianTimeToTriage** | Pointer to **string** |  | [optional] 
**ReportsLast30Days** | **int32** |  | 
**ReportsTotal** | **int32** |  | 
**Title** | **string** |  | 

## Methods

### NewDashboardEngagementStats

`func NewDashboardEngagementStats(bountySpendCents int32, bySeverity map[string]int32, costLimitCents int32, currency string, reportsLast30Days int32, reportsTotal int32, title string, ) *DashboardEngagementStats`

NewDashboardEngagementStats instantiates a new DashboardEngagementStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardEngagementStatsWithDefaults

`func NewDashboardEngagementStatsWithDefaults() *DashboardEngagementStats`

NewDashboardEngagementStatsWithDefaults instantiates a new DashboardEngagementStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBountySpendCents

`func (o *DashboardEngagementStats) GetBountySpendCents() int32`

GetBountySpendCents returns the BountySpendCents field if non-nil, zero value otherwise.

### GetBountySpendCentsOk

`func (o *DashboardEngagementStats) GetBountySpendCentsOk() (*int32, bool)`

GetBountySpendCentsOk returns a tuple with the BountySpendCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBountySpendCents

`func (o *DashboardEngagementStats) SetBountySpendCents(v int32)`

SetBountySpendCents sets BountySpendCents field to given value.


### GetBySeverity

`func (o *DashboardEngagementStats) GetBySeverity() map[string]int32`

GetBySeverity returns the BySeverity field if non-nil, zero value otherwise.

### GetBySeverityOk

`func (o *DashboardEngagementStats) GetBySeverityOk() (*map[string]int32, bool)`

GetBySeverityOk returns a tuple with the BySeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBySeverity

`func (o *DashboardEngagementStats) SetBySeverity(v map[string]int32)`

SetBySeverity sets BySeverity field to given value.


### GetCostLimitCents

`func (o *DashboardEngagementStats) GetCostLimitCents() int32`

GetCostLimitCents returns the CostLimitCents field if non-nil, zero value otherwise.

### GetCostLimitCentsOk

`func (o *DashboardEngagementStats) GetCostLimitCentsOk() (*int32, bool)`

GetCostLimitCentsOk returns a tuple with the CostLimitCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostLimitCents

`func (o *DashboardEngagementStats) SetCostLimitCents(v int32)`

SetCostLimitCents sets CostLimitCents field to given value.


### GetCurrency

`func (o *DashboardEngagementStats) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *DashboardEngagementStats) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *DashboardEngagementStats) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetMedianTimeToTriage

`func (o *DashboardEngagementStats) GetMedianTimeToTriage() string`

GetMedianTimeToTriage returns the MedianTimeToTriage field if non-nil, zero value otherwise.

### GetMedianTimeToTriageOk

`func (o *DashboardEngagementStats) GetMedianTimeToTriageOk() (*string, bool)`

GetMedianTimeToTriageOk returns a tuple with the MedianTimeToTriage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedianTimeToTriage

`func (o *DashboardEngagementStats) SetMedianTimeToTriage(v string)`

SetMedianTimeToTriage sets MedianTimeToTriage field to given value.

### HasMedianTimeToTriage

`func (o *DashboardEngagementStats) HasMedianTimeToTriage() bool`

HasMedianTimeToTriage returns a boolean if a field has been set.

### GetReportsLast30Days

`func (o *DashboardEngagementStats) GetReportsLast30Days() int32`

GetReportsLast30Days returns the ReportsLast30Days field if non-nil, zero value otherwise.

### GetReportsLast30DaysOk

`func (o *DashboardEngagementStats) GetReportsLast30DaysOk() (*int32, bool)`

GetReportsLast30DaysOk returns a tuple with the ReportsLast30Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportsLast30Days

`func (o *DashboardEngagementStats) SetReportsLast30Days(v int32)`

SetReportsLast30Days sets ReportsLast30Days field to given value.


### GetReportsTotal

`func (o *DashboardEngagementStats) GetReportsTotal() int32`

GetReportsTotal returns the ReportsTotal field if non-nil, zero value otherwise.

### GetReportsTotalOk

`func (o *DashboardEngagementStats) GetReportsTotalOk() (*int32, bool)`

GetReportsTotalOk returns a tuple with the ReportsTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportsTotal

`func (o *DashboardEngagementStats) SetReportsTotal(v int32)`

SetReportsTotal sets ReportsTotal field to given value.


### GetTitle

`func (o *DashboardEngagementStats) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DashboardEngagementStats) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DashboardEngagementStats) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


