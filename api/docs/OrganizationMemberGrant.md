# OrganizationMemberGrant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Engagement** | **string** |  | 
**Level** | **string** |  | 
**ScopeEntry** | Pointer to **string** |  | [optional] 

## Methods

### NewOrganizationMemberGrant

`func NewOrganizationMemberGrant(engagement string, level string, ) *OrganizationMemberGrant`

NewOrganizationMemberGrant instantiates a new OrganizationMemberGrant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationMemberGrantWithDefaults

`func NewOrganizationMemberGrantWithDefaults() *OrganizationMemberGrant`

NewOrganizationMemberGrantWithDefaults instantiates a new OrganizationMemberGrant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEngagement

`func (o *OrganizationMemberGrant) GetEngagement() string`

GetEngagement returns the Engagement field if non-nil, zero value otherwise.

### GetEngagementOk

`func (o *OrganizationMemberGrant) GetEngagementOk() (*string, bool)`

GetEngagementOk returns a tuple with the Engagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngagement

`func (o *OrganizationMemberGrant) SetEngagement(v string)`

SetEngagement sets Engagement field to given value.


### GetLevel

`func (o *OrganizationMemberGrant) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *OrganizationMemberGrant) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *OrganizationMemberGrant) SetLevel(v string)`

SetLevel sets Level field to given value.


### GetScopeEntry

`func (o *OrganizationMemberGrant) GetScopeEntry() string`

GetScopeEntry returns the ScopeEntry field if non-nil, zero value otherwise.

### GetScopeEntryOk

`func (o *OrganizationMemberGrant) GetScopeEntryOk() (*string, bool)`

GetScopeEntryOk returns a tuple with the ScopeEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeEntry

`func (o *OrganizationMemberGrant) SetScopeEntry(v string)`

SetScopeEntry sets ScopeEntry field to given value.

### HasScopeEntry

`func (o *OrganizationMemberGrant) HasScopeEntry() bool`

HasScopeEntry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


