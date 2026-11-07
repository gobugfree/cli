# EngagementProgram

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CvssBands** | [**[]EngagementCVSSBand**](EngagementCVSSBand.md) |  | 
**NonQualifiedText** | Pointer to **string** |  | [optional] 
**NonQualifiedVulns** | Pointer to **[]string** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**Procedure** | Pointer to **[]string** |  | [optional] 
**QualifiedText** | Pointer to **string** |  | [optional] 
**QualifiedVulns** | Pointer to **[]string** |  | [optional] 
**RulesText** | Pointer to **string** |  | [optional] 
**SafeHarbor** | Pointer to **string** |  | [optional] 
**ScopeNotes** | Pointer to **string** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**TestingAllowed** | Pointer to **[]string** |  | [optional] 
**TestingDisallowed** | Pointer to **[]string** |  | [optional] 

## Methods

### NewEngagementProgram

`func NewEngagementProgram(cvssBands []EngagementCVSSBand, ) *EngagementProgram`

NewEngagementProgram instantiates a new EngagementProgram object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngagementProgramWithDefaults

`func NewEngagementProgramWithDefaults() *EngagementProgram`

NewEngagementProgramWithDefaults instantiates a new EngagementProgram object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCvssBands

`func (o *EngagementProgram) GetCvssBands() []EngagementCVSSBand`

GetCvssBands returns the CvssBands field if non-nil, zero value otherwise.

### GetCvssBandsOk

`func (o *EngagementProgram) GetCvssBandsOk() (*[]EngagementCVSSBand, bool)`

GetCvssBandsOk returns a tuple with the CvssBands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssBands

`func (o *EngagementProgram) SetCvssBands(v []EngagementCVSSBand)`

SetCvssBands sets CvssBands field to given value.


### GetNonQualifiedText

`func (o *EngagementProgram) GetNonQualifiedText() string`

GetNonQualifiedText returns the NonQualifiedText field if non-nil, zero value otherwise.

### GetNonQualifiedTextOk

`func (o *EngagementProgram) GetNonQualifiedTextOk() (*string, bool)`

GetNonQualifiedTextOk returns a tuple with the NonQualifiedText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonQualifiedText

`func (o *EngagementProgram) SetNonQualifiedText(v string)`

SetNonQualifiedText sets NonQualifiedText field to given value.

### HasNonQualifiedText

`func (o *EngagementProgram) HasNonQualifiedText() bool`

HasNonQualifiedText returns a boolean if a field has been set.

### GetNonQualifiedVulns

`func (o *EngagementProgram) GetNonQualifiedVulns() []string`

GetNonQualifiedVulns returns the NonQualifiedVulns field if non-nil, zero value otherwise.

### GetNonQualifiedVulnsOk

`func (o *EngagementProgram) GetNonQualifiedVulnsOk() (*[]string, bool)`

GetNonQualifiedVulnsOk returns a tuple with the NonQualifiedVulns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonQualifiedVulns

`func (o *EngagementProgram) SetNonQualifiedVulns(v []string)`

SetNonQualifiedVulns sets NonQualifiedVulns field to given value.

### HasNonQualifiedVulns

`func (o *EngagementProgram) HasNonQualifiedVulns() bool`

HasNonQualifiedVulns returns a boolean if a field has been set.

### GetNotes

`func (o *EngagementProgram) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *EngagementProgram) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *EngagementProgram) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *EngagementProgram) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetProcedure

`func (o *EngagementProgram) GetProcedure() []string`

GetProcedure returns the Procedure field if non-nil, zero value otherwise.

### GetProcedureOk

`func (o *EngagementProgram) GetProcedureOk() (*[]string, bool)`

GetProcedureOk returns a tuple with the Procedure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcedure

`func (o *EngagementProgram) SetProcedure(v []string)`

SetProcedure sets Procedure field to given value.

### HasProcedure

`func (o *EngagementProgram) HasProcedure() bool`

HasProcedure returns a boolean if a field has been set.

### GetQualifiedText

`func (o *EngagementProgram) GetQualifiedText() string`

GetQualifiedText returns the QualifiedText field if non-nil, zero value otherwise.

### GetQualifiedTextOk

`func (o *EngagementProgram) GetQualifiedTextOk() (*string, bool)`

GetQualifiedTextOk returns a tuple with the QualifiedText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualifiedText

`func (o *EngagementProgram) SetQualifiedText(v string)`

SetQualifiedText sets QualifiedText field to given value.

### HasQualifiedText

`func (o *EngagementProgram) HasQualifiedText() bool`

HasQualifiedText returns a boolean if a field has been set.

### GetQualifiedVulns

`func (o *EngagementProgram) GetQualifiedVulns() []string`

GetQualifiedVulns returns the QualifiedVulns field if non-nil, zero value otherwise.

### GetQualifiedVulnsOk

`func (o *EngagementProgram) GetQualifiedVulnsOk() (*[]string, bool)`

GetQualifiedVulnsOk returns a tuple with the QualifiedVulns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualifiedVulns

`func (o *EngagementProgram) SetQualifiedVulns(v []string)`

SetQualifiedVulns sets QualifiedVulns field to given value.

### HasQualifiedVulns

`func (o *EngagementProgram) HasQualifiedVulns() bool`

HasQualifiedVulns returns a boolean if a field has been set.

### GetRulesText

`func (o *EngagementProgram) GetRulesText() string`

GetRulesText returns the RulesText field if non-nil, zero value otherwise.

### GetRulesTextOk

`func (o *EngagementProgram) GetRulesTextOk() (*string, bool)`

GetRulesTextOk returns a tuple with the RulesText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRulesText

`func (o *EngagementProgram) SetRulesText(v string)`

SetRulesText sets RulesText field to given value.

### HasRulesText

`func (o *EngagementProgram) HasRulesText() bool`

HasRulesText returns a boolean if a field has been set.

### GetSafeHarbor

`func (o *EngagementProgram) GetSafeHarbor() string`

GetSafeHarbor returns the SafeHarbor field if non-nil, zero value otherwise.

### GetSafeHarborOk

`func (o *EngagementProgram) GetSafeHarborOk() (*string, bool)`

GetSafeHarborOk returns a tuple with the SafeHarbor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafeHarbor

`func (o *EngagementProgram) SetSafeHarbor(v string)`

SetSafeHarbor sets SafeHarbor field to given value.

### HasSafeHarbor

`func (o *EngagementProgram) HasSafeHarbor() bool`

HasSafeHarbor returns a boolean if a field has been set.

### GetScopeNotes

`func (o *EngagementProgram) GetScopeNotes() string`

GetScopeNotes returns the ScopeNotes field if non-nil, zero value otherwise.

### GetScopeNotesOk

`func (o *EngagementProgram) GetScopeNotesOk() (*string, bool)`

GetScopeNotesOk returns a tuple with the ScopeNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeNotes

`func (o *EngagementProgram) SetScopeNotes(v string)`

SetScopeNotes sets ScopeNotes field to given value.

### HasScopeNotes

`func (o *EngagementProgram) HasScopeNotes() bool`

HasScopeNotes returns a boolean if a field has been set.

### GetSummary

`func (o *EngagementProgram) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *EngagementProgram) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *EngagementProgram) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *EngagementProgram) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetTestingAllowed

`func (o *EngagementProgram) GetTestingAllowed() []string`

GetTestingAllowed returns the TestingAllowed field if non-nil, zero value otherwise.

### GetTestingAllowedOk

`func (o *EngagementProgram) GetTestingAllowedOk() (*[]string, bool)`

GetTestingAllowedOk returns a tuple with the TestingAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestingAllowed

`func (o *EngagementProgram) SetTestingAllowed(v []string)`

SetTestingAllowed sets TestingAllowed field to given value.

### HasTestingAllowed

`func (o *EngagementProgram) HasTestingAllowed() bool`

HasTestingAllowed returns a boolean if a field has been set.

### GetTestingDisallowed

`func (o *EngagementProgram) GetTestingDisallowed() []string`

GetTestingDisallowed returns the TestingDisallowed field if non-nil, zero value otherwise.

### GetTestingDisallowedOk

`func (o *EngagementProgram) GetTestingDisallowedOk() (*[]string, bool)`

GetTestingDisallowedOk returns a tuple with the TestingDisallowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestingDisallowed

`func (o *EngagementProgram) SetTestingDisallowed(v []string)`

SetTestingDisallowed sets TestingDisallowed field to given value.

### HasTestingDisallowed

`func (o *EngagementProgram) HasTestingDisallowed() bool`

HasTestingDisallowed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


