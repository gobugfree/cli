# UserOnboarding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConductAccepted** | Pointer to **time.Time** |  | [optional] 
**OnboardedAt** | Pointer to **time.Time** |  | [optional] 
**Step** | **int32** |  | 

## Methods

### NewUserOnboarding

`func NewUserOnboarding(step int32, ) *UserOnboarding`

NewUserOnboarding instantiates a new UserOnboarding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserOnboardingWithDefaults

`func NewUserOnboardingWithDefaults() *UserOnboarding`

NewUserOnboardingWithDefaults instantiates a new UserOnboarding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConductAccepted

`func (o *UserOnboarding) GetConductAccepted() time.Time`

GetConductAccepted returns the ConductAccepted field if non-nil, zero value otherwise.

### GetConductAcceptedOk

`func (o *UserOnboarding) GetConductAcceptedOk() (*time.Time, bool)`

GetConductAcceptedOk returns a tuple with the ConductAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConductAccepted

`func (o *UserOnboarding) SetConductAccepted(v time.Time)`

SetConductAccepted sets ConductAccepted field to given value.

### HasConductAccepted

`func (o *UserOnboarding) HasConductAccepted() bool`

HasConductAccepted returns a boolean if a field has been set.

### GetOnboardedAt

`func (o *UserOnboarding) GetOnboardedAt() time.Time`

GetOnboardedAt returns the OnboardedAt field if non-nil, zero value otherwise.

### GetOnboardedAtOk

`func (o *UserOnboarding) GetOnboardedAtOk() (*time.Time, bool)`

GetOnboardedAtOk returns a tuple with the OnboardedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboardedAt

`func (o *UserOnboarding) SetOnboardedAt(v time.Time)`

SetOnboardedAt sets OnboardedAt field to given value.

### HasOnboardedAt

`func (o *UserOnboarding) HasOnboardedAt() bool`

HasOnboardedAt returns a boolean if a field has been set.

### GetStep

`func (o *UserOnboarding) GetStep() int32`

GetStep returns the Step field if non-nil, zero value otherwise.

### GetStepOk

`func (o *UserOnboarding) GetStepOk() (*int32, bool)`

GetStepOk returns a tuple with the Step field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStep

`func (o *UserOnboarding) SetStep(v int32)`

SetStep sets Step field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


