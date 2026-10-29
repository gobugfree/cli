# HelpOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | **string** |  | 
**Describe** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Id** | **string** |  | 
**Method** | **string** |  | 
**Multipart** | **bool** |  | 
**Path** | **string** |  | 
**Prompt** | **string** |  | 
**Summary** | **string** |  | 

## Methods

### NewHelpOperation

`func NewHelpOperation(command string, describe string, id string, method string, multipart bool, path string, prompt string, summary string, ) *HelpOperation`

NewHelpOperation instantiates a new HelpOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelpOperationWithDefaults

`func NewHelpOperationWithDefaults() *HelpOperation`

NewHelpOperationWithDefaults instantiates a new HelpOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommand

`func (o *HelpOperation) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *HelpOperation) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *HelpOperation) SetCommand(v string)`

SetCommand sets Command field to given value.


### GetDescribe

`func (o *HelpOperation) GetDescribe() string`

GetDescribe returns the Describe field if non-nil, zero value otherwise.

### GetDescribeOk

`func (o *HelpOperation) GetDescribeOk() (*string, bool)`

GetDescribeOk returns a tuple with the Describe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescribe

`func (o *HelpOperation) SetDescribe(v string)`

SetDescribe sets Describe field to given value.


### GetDescription

`func (o *HelpOperation) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HelpOperation) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HelpOperation) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HelpOperation) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *HelpOperation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HelpOperation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HelpOperation) SetId(v string)`

SetId sets Id field to given value.


### GetMethod

`func (o *HelpOperation) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *HelpOperation) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *HelpOperation) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetMultipart

`func (o *HelpOperation) GetMultipart() bool`

GetMultipart returns the Multipart field if non-nil, zero value otherwise.

### GetMultipartOk

`func (o *HelpOperation) GetMultipartOk() (*bool, bool)`

GetMultipartOk returns a tuple with the Multipart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipart

`func (o *HelpOperation) SetMultipart(v bool)`

SetMultipart sets Multipart field to given value.


### GetPath

`func (o *HelpOperation) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *HelpOperation) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *HelpOperation) SetPath(v string)`

SetPath sets Path field to given value.


### GetPrompt

`func (o *HelpOperation) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *HelpOperation) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *HelpOperation) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.


### GetSummary

`func (o *HelpOperation) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *HelpOperation) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *HelpOperation) SetSummary(v string)`

SetSummary sets Summary field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


