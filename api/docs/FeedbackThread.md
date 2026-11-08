# FeedbackThread

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Feedback** | [**Feedback**](Feedback.md) |  | 
**Files** | [**[]FeedbackFile**](FeedbackFile.md) |  | 
**Messages** | [**[]FeedbackMessage**](FeedbackMessage.md) |  | 

## Methods

### NewFeedbackThread

`func NewFeedbackThread(feedback Feedback, files []FeedbackFile, messages []FeedbackMessage, ) *FeedbackThread`

NewFeedbackThread instantiates a new FeedbackThread object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeedbackThreadWithDefaults

`func NewFeedbackThreadWithDefaults() *FeedbackThread`

NewFeedbackThreadWithDefaults instantiates a new FeedbackThread object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeedback

`func (o *FeedbackThread) GetFeedback() Feedback`

GetFeedback returns the Feedback field if non-nil, zero value otherwise.

### GetFeedbackOk

`func (o *FeedbackThread) GetFeedbackOk() (*Feedback, bool)`

GetFeedbackOk returns a tuple with the Feedback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedback

`func (o *FeedbackThread) SetFeedback(v Feedback)`

SetFeedback sets Feedback field to given value.


### GetFiles

`func (o *FeedbackThread) GetFiles() []FeedbackFile`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *FeedbackThread) GetFilesOk() (*[]FeedbackFile, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *FeedbackThread) SetFiles(v []FeedbackFile)`

SetFiles sets Files field to given value.


### GetMessages

`func (o *FeedbackThread) GetMessages() []FeedbackMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *FeedbackThread) GetMessagesOk() (*[]FeedbackMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *FeedbackThread) SetMessages(v []FeedbackMessage)`

SetMessages sets Messages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


