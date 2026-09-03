# EngagementVersionConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CvssBands** | Pointer to [**[]EngagementCVSSBand**](EngagementCVSSBand.md) |  | [optional] 
**NonQualifiedVulns** | Pointer to **[]string** |  | [optional] 
**Procedure** | Pointer to **[]string** |  | [optional] 
**QualifiedVulns** | Pointer to **[]string** |  | [optional] 
**SafeHarbor** | Pointer to **string** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**TestingRules** | [**EngagementTestingRules**](EngagementTestingRules.md) |  | 

## Methods

### NewEngagementVersionConfig

`func NewEngagementVersionConfig(testingRules EngagementTestingRules, ) *EngagementVersionConfig`

NewEngagementVersionConfig instantiates a new EngagementVersionConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngagementVersionConfigWithDefaults

`func NewEngagementVersionConfigWithDefaults() *EngagementVersionConfig`

NewEngagementVersionConfigWithDefaults instantiates a new EngagementVersionConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCvssBands

`func (o *EngagementVersionConfig) GetCvssBands() []EngagementCVSSBand`

GetCvssBands returns the CvssBands field if non-nil, zero value otherwise.

### GetCvssBandsOk

`func (o *EngagementVersionConfig) GetCvssBandsOk() (*[]EngagementCVSSBand, bool)`

GetCvssBandsOk returns a tuple with the CvssBands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssBands

`func (o *EngagementVersionConfig) SetCvssBands(v []EngagementCVSSBand)`

SetCvssBands sets CvssBands field to given value.

### HasCvssBands

`func (o *EngagementVersionConfig) HasCvssBands() bool`

HasCvssBands returns a boolean if a field has been set.

### GetNonQualifiedVulns

`func (o *EngagementVersionConfig) GetNonQualifiedVulns() []string`

GetNonQualifiedVulns returns the NonQualifiedVulns field if non-nil, zero value otherwise.

### GetNonQualifiedVulnsOk

`func (o *EngagementVersionConfig) GetNonQualifiedVulnsOk() (*[]string, bool)`

GetNonQualifiedVulnsOk returns a tuple with the NonQualifiedVulns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonQualifiedVulns

`func (o *EngagementVersionConfig) SetNonQualifiedVulns(v []string)`

SetNonQualifiedVulns sets NonQualifiedVulns field to given value.

### HasNonQualifiedVulns

`func (o *EngagementVersionConfig) HasNonQualifiedVulns() bool`

HasNonQualifiedVulns returns a boolean if a field has been set.

### GetProcedure

`func (o *EngagementVersionConfig) GetProcedure() []string`

GetProcedure returns the Procedure field if non-nil, zero value otherwise.

### GetProcedureOk

`func (o *EngagementVersionConfig) GetProcedureOk() (*[]string, bool)`

GetProcedureOk returns a tuple with the Procedure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcedure

`func (o *EngagementVersionConfig) SetProcedure(v []string)`

SetProcedure sets Procedure field to given value.

### HasProcedure

`func (o *EngagementVersionConfig) HasProcedure() bool`

HasProcedure returns a boolean if a field has been set.

### GetQualifiedVulns

`func (o *EngagementVersionConfig) GetQualifiedVulns() []string`

GetQualifiedVulns returns the QualifiedVulns field if non-nil, zero value otherwise.

### GetQualifiedVulnsOk

`func (o *EngagementVersionConfig) GetQualifiedVulnsOk() (*[]string, bool)`

GetQualifiedVulnsOk returns a tuple with the QualifiedVulns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualifiedVulns

`func (o *EngagementVersionConfig) SetQualifiedVulns(v []string)`

SetQualifiedVulns sets QualifiedVulns field to given value.

### HasQualifiedVulns

`func (o *EngagementVersionConfig) HasQualifiedVulns() bool`

HasQualifiedVulns returns a boolean if a field has been set.

### GetSafeHarbor

`func (o *EngagementVersionConfig) GetSafeHarbor() string`

GetSafeHarbor returns the SafeHarbor field if non-nil, zero value otherwise.

### GetSafeHarborOk

`func (o *EngagementVersionConfig) GetSafeHarborOk() (*string, bool)`

GetSafeHarborOk returns a tuple with the SafeHarbor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafeHarbor

`func (o *EngagementVersionConfig) SetSafeHarbor(v string)`

SetSafeHarbor sets SafeHarbor field to given value.

### HasSafeHarbor

`func (o *EngagementVersionConfig) HasSafeHarbor() bool`

HasSafeHarbor returns a boolean if a field has been set.

### GetSummary

`func (o *EngagementVersionConfig) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *EngagementVersionConfig) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *EngagementVersionConfig) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *EngagementVersionConfig) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetTestingRules

`func (o *EngagementVersionConfig) GetTestingRules() EngagementTestingRules`

GetTestingRules returns the TestingRules field if non-nil, zero value otherwise.

### GetTestingRulesOk

`func (o *EngagementVersionConfig) GetTestingRulesOk() (*EngagementTestingRules, bool)`

GetTestingRulesOk returns a tuple with the TestingRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestingRules

`func (o *EngagementVersionConfig) SetTestingRules(v EngagementTestingRules)`

SetTestingRules sets TestingRules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


