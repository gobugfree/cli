# AuditEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** |  | 
**ActorEmail** | Pointer to **string** |  | [optional] 
**ActorId** | Pointer to **int32** |  | [optional] 
**ActorName** | Pointer to **string** |  | [optional] 
**Created** | **time.Time** |  | 
**Digest** | **string** |  | 
**EntityId** | **string** |  | 
**EntityKind** | **string** |  | 
**Id** | **int32** |  | 
**Ip** | Pointer to **string** |  | [optional] 
**PrevDigest** | Pointer to **string** |  | [optional] 
**Target** | Pointer to **string** |  | [optional] 

## Methods

### NewAuditEvent

`func NewAuditEvent(action string, created time.Time, digest string, entityId string, entityKind string, id int32, ) *AuditEvent`

NewAuditEvent instantiates a new AuditEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditEventWithDefaults

`func NewAuditEventWithDefaults() *AuditEvent`

NewAuditEventWithDefaults instantiates a new AuditEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *AuditEvent) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *AuditEvent) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *AuditEvent) SetAction(v string)`

SetAction sets Action field to given value.


### GetActorEmail

`func (o *AuditEvent) GetActorEmail() string`

GetActorEmail returns the ActorEmail field if non-nil, zero value otherwise.

### GetActorEmailOk

`func (o *AuditEvent) GetActorEmailOk() (*string, bool)`

GetActorEmailOk returns a tuple with the ActorEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorEmail

`func (o *AuditEvent) SetActorEmail(v string)`

SetActorEmail sets ActorEmail field to given value.

### HasActorEmail

`func (o *AuditEvent) HasActorEmail() bool`

HasActorEmail returns a boolean if a field has been set.

### GetActorId

`func (o *AuditEvent) GetActorId() int32`

GetActorId returns the ActorId field if non-nil, zero value otherwise.

### GetActorIdOk

`func (o *AuditEvent) GetActorIdOk() (*int32, bool)`

GetActorIdOk returns a tuple with the ActorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorId

`func (o *AuditEvent) SetActorId(v int32)`

SetActorId sets ActorId field to given value.

### HasActorId

`func (o *AuditEvent) HasActorId() bool`

HasActorId returns a boolean if a field has been set.

### GetActorName

`func (o *AuditEvent) GetActorName() string`

GetActorName returns the ActorName field if non-nil, zero value otherwise.

### GetActorNameOk

`func (o *AuditEvent) GetActorNameOk() (*string, bool)`

GetActorNameOk returns a tuple with the ActorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorName

`func (o *AuditEvent) SetActorName(v string)`

SetActorName sets ActorName field to given value.

### HasActorName

`func (o *AuditEvent) HasActorName() bool`

HasActorName returns a boolean if a field has been set.

### GetCreated

`func (o *AuditEvent) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AuditEvent) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AuditEvent) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetDigest

`func (o *AuditEvent) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *AuditEvent) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *AuditEvent) SetDigest(v string)`

SetDigest sets Digest field to given value.


### GetEntityId

`func (o *AuditEvent) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *AuditEvent) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *AuditEvent) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.


### GetEntityKind

`func (o *AuditEvent) GetEntityKind() string`

GetEntityKind returns the EntityKind field if non-nil, zero value otherwise.

### GetEntityKindOk

`func (o *AuditEvent) GetEntityKindOk() (*string, bool)`

GetEntityKindOk returns a tuple with the EntityKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityKind

`func (o *AuditEvent) SetEntityKind(v string)`

SetEntityKind sets EntityKind field to given value.


### GetId

`func (o *AuditEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuditEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuditEvent) SetId(v int32)`

SetId sets Id field to given value.


### GetIp

`func (o *AuditEvent) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *AuditEvent) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *AuditEvent) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *AuditEvent) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetPrevDigest

`func (o *AuditEvent) GetPrevDigest() string`

GetPrevDigest returns the PrevDigest field if non-nil, zero value otherwise.

### GetPrevDigestOk

`func (o *AuditEvent) GetPrevDigestOk() (*string, bool)`

GetPrevDigestOk returns a tuple with the PrevDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevDigest

`func (o *AuditEvent) SetPrevDigest(v string)`

SetPrevDigest sets PrevDigest field to given value.

### HasPrevDigest

`func (o *AuditEvent) HasPrevDigest() bool`

HasPrevDigest returns a boolean if a field has been set.

### GetTarget

`func (o *AuditEvent) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *AuditEvent) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *AuditEvent) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *AuditEvent) HasTarget() bool`

HasTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


