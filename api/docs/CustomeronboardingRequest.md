# CustomeronboardingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | **time.Time** |  | 
**CustomerEmail** | **string** |  | 
**EngagementId** | Pointer to **string** |  | [optional] 
**Expires** | **time.Time** |  | 
**Id** | **string** |  | 
**OrganizationId** | Pointer to **int32** |  | [optional] 
**Status** | **string** |  | 
**Submitted** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewCustomeronboardingRequest

`func NewCustomeronboardingRequest(created time.Time, customerEmail string, expires time.Time, id string, status string, ) *CustomeronboardingRequest`

NewCustomeronboardingRequest instantiates a new CustomeronboardingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomeronboardingRequestWithDefaults

`func NewCustomeronboardingRequestWithDefaults() *CustomeronboardingRequest`

NewCustomeronboardingRequestWithDefaults instantiates a new CustomeronboardingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *CustomeronboardingRequest) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CustomeronboardingRequest) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CustomeronboardingRequest) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetCustomerEmail

`func (o *CustomeronboardingRequest) GetCustomerEmail() string`

GetCustomerEmail returns the CustomerEmail field if non-nil, zero value otherwise.

### GetCustomerEmailOk

`func (o *CustomeronboardingRequest) GetCustomerEmailOk() (*string, bool)`

GetCustomerEmailOk returns a tuple with the CustomerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerEmail

`func (o *CustomeronboardingRequest) SetCustomerEmail(v string)`

SetCustomerEmail sets CustomerEmail field to given value.


### GetEngagementId

`func (o *CustomeronboardingRequest) GetEngagementId() string`

GetEngagementId returns the EngagementId field if non-nil, zero value otherwise.

### GetEngagementIdOk

`func (o *CustomeronboardingRequest) GetEngagementIdOk() (*string, bool)`

GetEngagementIdOk returns a tuple with the EngagementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngagementId

`func (o *CustomeronboardingRequest) SetEngagementId(v string)`

SetEngagementId sets EngagementId field to given value.

### HasEngagementId

`func (o *CustomeronboardingRequest) HasEngagementId() bool`

HasEngagementId returns a boolean if a field has been set.

### GetExpires

`func (o *CustomeronboardingRequest) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *CustomeronboardingRequest) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *CustomeronboardingRequest) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.


### GetId

`func (o *CustomeronboardingRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomeronboardingRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomeronboardingRequest) SetId(v string)`

SetId sets Id field to given value.


### GetOrganizationId

`func (o *CustomeronboardingRequest) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *CustomeronboardingRequest) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *CustomeronboardingRequest) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *CustomeronboardingRequest) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetStatus

`func (o *CustomeronboardingRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CustomeronboardingRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CustomeronboardingRequest) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSubmitted

`func (o *CustomeronboardingRequest) GetSubmitted() time.Time`

GetSubmitted returns the Submitted field if non-nil, zero value otherwise.

### GetSubmittedOk

`func (o *CustomeronboardingRequest) GetSubmittedOk() (*time.Time, bool)`

GetSubmittedOk returns a tuple with the Submitted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmitted

`func (o *CustomeronboardingRequest) SetSubmitted(v time.Time)`

SetSubmitted sets Submitted field to given value.

### HasSubmitted

`func (o *CustomeronboardingRequest) HasSubmitted() bool`

HasSubmitted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


