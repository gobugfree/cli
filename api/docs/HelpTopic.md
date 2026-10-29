# HelpTopic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Links** | **map[string]string** |  | 
**Operations** | [**[]HelpOperation**](HelpOperation.md) |  | 
**Page** | **string** |  | 
**Pattern** | **string** |  | 
**Summary** | **string** |  | 

## Methods

### NewHelpTopic

`func NewHelpTopic(links map[string]string, operations []HelpOperation, page string, pattern string, summary string, ) *HelpTopic`

NewHelpTopic instantiates a new HelpTopic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelpTopicWithDefaults

`func NewHelpTopicWithDefaults() *HelpTopic`

NewHelpTopicWithDefaults instantiates a new HelpTopic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *HelpTopic) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HelpTopic) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HelpTopic) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HelpTopic) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLinks

`func (o *HelpTopic) GetLinks() map[string]string`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *HelpTopic) GetLinksOk() (*map[string]string, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *HelpTopic) SetLinks(v map[string]string)`

SetLinks sets Links field to given value.


### GetOperations

`func (o *HelpTopic) GetOperations() []HelpOperation`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *HelpTopic) GetOperationsOk() (*[]HelpOperation, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *HelpTopic) SetOperations(v []HelpOperation)`

SetOperations sets Operations field to given value.


### GetPage

`func (o *HelpTopic) GetPage() string`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *HelpTopic) GetPageOk() (*string, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *HelpTopic) SetPage(v string)`

SetPage sets Page field to given value.


### GetPattern

`func (o *HelpTopic) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *HelpTopic) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *HelpTopic) SetPattern(v string)`

SetPattern sets Pattern field to given value.


### GetSummary

`func (o *HelpTopic) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *HelpTopic) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *HelpTopic) SetSummary(v string)`

SetSummary sets Summary field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


