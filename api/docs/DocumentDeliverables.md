# DocumentDeliverables

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Engagement** | **string** |  | 
**Files** | [**[]DocumentDeliverable**](DocumentDeliverable.md) |  | 
**Statement** | [**DocumentStatementState**](DocumentStatementState.md) |  | 

## Methods

### NewDocumentDeliverables

`func NewDocumentDeliverables(engagement string, files []DocumentDeliverable, statement DocumentStatementState, ) *DocumentDeliverables`

NewDocumentDeliverables instantiates a new DocumentDeliverables object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentDeliverablesWithDefaults

`func NewDocumentDeliverablesWithDefaults() *DocumentDeliverables`

NewDocumentDeliverablesWithDefaults instantiates a new DocumentDeliverables object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEngagement

`func (o *DocumentDeliverables) GetEngagement() string`

GetEngagement returns the Engagement field if non-nil, zero value otherwise.

### GetEngagementOk

`func (o *DocumentDeliverables) GetEngagementOk() (*string, bool)`

GetEngagementOk returns a tuple with the Engagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngagement

`func (o *DocumentDeliverables) SetEngagement(v string)`

SetEngagement sets Engagement field to given value.


### GetFiles

`func (o *DocumentDeliverables) GetFiles() []DocumentDeliverable`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *DocumentDeliverables) GetFilesOk() (*[]DocumentDeliverable, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *DocumentDeliverables) SetFiles(v []DocumentDeliverable)`

SetFiles sets Files field to given value.


### GetStatement

`func (o *DocumentDeliverables) GetStatement() DocumentStatementState`

GetStatement returns the Statement field if non-nil, zero value otherwise.

### GetStatementOk

`func (o *DocumentDeliverables) GetStatementOk() (*DocumentStatementState, bool)`

GetStatementOk returns a tuple with the Statement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatement

`func (o *DocumentDeliverables) SetStatement(v DocumentStatementState)`

SetStatement sets Statement field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


