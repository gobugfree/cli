# ReportFeed

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**At** | **time.Time** |  | 
**Author** | Pointer to **string** |  | [optional] 
**AuthorParty** | Pointer to **string** |  | [optional] 
**Body** | Pointer to **string** |  | [optional] 
**From** | Pointer to **string** |  | [optional] 
**Kind** | **string** |  | 
**To** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 

## Methods

### NewReportFeed

`func NewReportFeed(at time.Time, kind string, ) *ReportFeed`

NewReportFeed instantiates a new ReportFeed object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportFeedWithDefaults

`func NewReportFeedWithDefaults() *ReportFeed`

NewReportFeedWithDefaults instantiates a new ReportFeed object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAt

`func (o *ReportFeed) GetAt() time.Time`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *ReportFeed) GetAtOk() (*time.Time, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *ReportFeed) SetAt(v time.Time)`

SetAt sets At field to given value.


### GetAuthor

`func (o *ReportFeed) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *ReportFeed) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *ReportFeed) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *ReportFeed) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetAuthorParty

`func (o *ReportFeed) GetAuthorParty() string`

GetAuthorParty returns the AuthorParty field if non-nil, zero value otherwise.

### GetAuthorPartyOk

`func (o *ReportFeed) GetAuthorPartyOk() (*string, bool)`

GetAuthorPartyOk returns a tuple with the AuthorParty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorParty

`func (o *ReportFeed) SetAuthorParty(v string)`

SetAuthorParty sets AuthorParty field to given value.

### HasAuthorParty

`func (o *ReportFeed) HasAuthorParty() bool`

HasAuthorParty returns a boolean if a field has been set.

### GetBody

`func (o *ReportFeed) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *ReportFeed) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *ReportFeed) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *ReportFeed) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetFrom

`func (o *ReportFeed) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *ReportFeed) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *ReportFeed) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *ReportFeed) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetKind

`func (o *ReportFeed) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ReportFeed) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ReportFeed) SetKind(v string)`

SetKind sets Kind field to given value.


### GetTo

`func (o *ReportFeed) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *ReportFeed) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *ReportFeed) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *ReportFeed) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetVisibility

`func (o *ReportFeed) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ReportFeed) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ReportFeed) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ReportFeed) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


