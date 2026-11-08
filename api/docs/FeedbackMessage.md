# FeedbackMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Author** | **string** |  | 
**Body** | **string** |  | 
**ByStaff** | **bool** |  | 
**CreatedAt** | **string** |  | 
**Files** | [**[]FeedbackFile**](FeedbackFile.md) |  | 
**Id** | **int32** |  | 

## Methods

### NewFeedbackMessage

`func NewFeedbackMessage(author string, body string, byStaff bool, createdAt string, files []FeedbackFile, id int32, ) *FeedbackMessage`

NewFeedbackMessage instantiates a new FeedbackMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeedbackMessageWithDefaults

`func NewFeedbackMessageWithDefaults() *FeedbackMessage`

NewFeedbackMessageWithDefaults instantiates a new FeedbackMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthor

`func (o *FeedbackMessage) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *FeedbackMessage) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *FeedbackMessage) SetAuthor(v string)`

SetAuthor sets Author field to given value.


### GetBody

`func (o *FeedbackMessage) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *FeedbackMessage) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *FeedbackMessage) SetBody(v string)`

SetBody sets Body field to given value.


### GetByStaff

`func (o *FeedbackMessage) GetByStaff() bool`

GetByStaff returns the ByStaff field if non-nil, zero value otherwise.

### GetByStaffOk

`func (o *FeedbackMessage) GetByStaffOk() (*bool, bool)`

GetByStaffOk returns a tuple with the ByStaff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByStaff

`func (o *FeedbackMessage) SetByStaff(v bool)`

SetByStaff sets ByStaff field to given value.


### GetCreatedAt

`func (o *FeedbackMessage) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FeedbackMessage) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FeedbackMessage) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetFiles

`func (o *FeedbackMessage) GetFiles() []FeedbackFile`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *FeedbackMessage) GetFilesOk() (*[]FeedbackFile, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *FeedbackMessage) SetFiles(v []FeedbackFile)`

SetFiles sets Files field to given value.


### GetId

`func (o *FeedbackMessage) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FeedbackMessage) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FeedbackMessage) SetId(v int32)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


