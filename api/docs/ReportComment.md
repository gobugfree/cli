# ReportComment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attachments** | [**[]ReportAttachment**](ReportAttachment.md) |  | 
**Author** | **string** |  | 
**AuthorParty** | **string** |  | 
**Body** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**Id** | **int32** |  | 

## Methods

### NewReportComment

`func NewReportComment(attachments []ReportAttachment, author string, authorParty string, body string, createdAt time.Time, id int32, ) *ReportComment`

NewReportComment instantiates a new ReportComment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportCommentWithDefaults

`func NewReportCommentWithDefaults() *ReportComment`

NewReportCommentWithDefaults instantiates a new ReportComment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachments

`func (o *ReportComment) GetAttachments() []ReportAttachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *ReportComment) GetAttachmentsOk() (*[]ReportAttachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *ReportComment) SetAttachments(v []ReportAttachment)`

SetAttachments sets Attachments field to given value.


### GetAuthor

`func (o *ReportComment) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *ReportComment) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *ReportComment) SetAuthor(v string)`

SetAuthor sets Author field to given value.


### GetAuthorParty

`func (o *ReportComment) GetAuthorParty() string`

GetAuthorParty returns the AuthorParty field if non-nil, zero value otherwise.

### GetAuthorPartyOk

`func (o *ReportComment) GetAuthorPartyOk() (*string, bool)`

GetAuthorPartyOk returns a tuple with the AuthorParty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorParty

`func (o *ReportComment) SetAuthorParty(v string)`

SetAuthorParty sets AuthorParty field to given value.


### GetBody

`func (o *ReportComment) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *ReportComment) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *ReportComment) SetBody(v string)`

SetBody sets Body field to given value.


### GetCreatedAt

`func (o *ReportComment) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ReportComment) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ReportComment) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetId

`func (o *ReportComment) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReportComment) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReportComment) SetId(v int32)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


