# EventPublicLeaderboard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndsAt** | **time.Time** |  | 
**Items** | [**[]EventPublicEntry**](EventPublicEntry.md) |  | 
**Partners** | [**[]EventPublicPartner**](EventPublicPartner.md) |  | 
**RefreshedAt** | **time.Time** |  | 
**Slug** | **string** |  | 
**StartsAt** | **time.Time** |  | 
**Stats** | [**EventPublicStats**](EventPublicStats.md) |  | 
**Status** | **string** |  | 
**Title** | **string** |  | 

## Methods

### NewEventPublicLeaderboard

`func NewEventPublicLeaderboard(endsAt time.Time, items []EventPublicEntry, partners []EventPublicPartner, refreshedAt time.Time, slug string, startsAt time.Time, stats EventPublicStats, status string, title string, ) *EventPublicLeaderboard`

NewEventPublicLeaderboard instantiates a new EventPublicLeaderboard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventPublicLeaderboardWithDefaults

`func NewEventPublicLeaderboardWithDefaults() *EventPublicLeaderboard`

NewEventPublicLeaderboardWithDefaults instantiates a new EventPublicLeaderboard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndsAt

`func (o *EventPublicLeaderboard) GetEndsAt() time.Time`

GetEndsAt returns the EndsAt field if non-nil, zero value otherwise.

### GetEndsAtOk

`func (o *EventPublicLeaderboard) GetEndsAtOk() (*time.Time, bool)`

GetEndsAtOk returns a tuple with the EndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndsAt

`func (o *EventPublicLeaderboard) SetEndsAt(v time.Time)`

SetEndsAt sets EndsAt field to given value.


### GetItems

`func (o *EventPublicLeaderboard) GetItems() []EventPublicEntry`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *EventPublicLeaderboard) GetItemsOk() (*[]EventPublicEntry, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *EventPublicLeaderboard) SetItems(v []EventPublicEntry)`

SetItems sets Items field to given value.


### GetPartners

`func (o *EventPublicLeaderboard) GetPartners() []EventPublicPartner`

GetPartners returns the Partners field if non-nil, zero value otherwise.

### GetPartnersOk

`func (o *EventPublicLeaderboard) GetPartnersOk() (*[]EventPublicPartner, bool)`

GetPartnersOk returns a tuple with the Partners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartners

`func (o *EventPublicLeaderboard) SetPartners(v []EventPublicPartner)`

SetPartners sets Partners field to given value.


### GetRefreshedAt

`func (o *EventPublicLeaderboard) GetRefreshedAt() time.Time`

GetRefreshedAt returns the RefreshedAt field if non-nil, zero value otherwise.

### GetRefreshedAtOk

`func (o *EventPublicLeaderboard) GetRefreshedAtOk() (*time.Time, bool)`

GetRefreshedAtOk returns a tuple with the RefreshedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshedAt

`func (o *EventPublicLeaderboard) SetRefreshedAt(v time.Time)`

SetRefreshedAt sets RefreshedAt field to given value.


### GetSlug

`func (o *EventPublicLeaderboard) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *EventPublicLeaderboard) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *EventPublicLeaderboard) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetStartsAt

`func (o *EventPublicLeaderboard) GetStartsAt() time.Time`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *EventPublicLeaderboard) GetStartsAtOk() (*time.Time, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *EventPublicLeaderboard) SetStartsAt(v time.Time)`

SetStartsAt sets StartsAt field to given value.


### GetStats

`func (o *EventPublicLeaderboard) GetStats() EventPublicStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *EventPublicLeaderboard) GetStatsOk() (*EventPublicStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *EventPublicLeaderboard) SetStats(v EventPublicStats)`

SetStats sets Stats field to given value.


### GetStatus

`func (o *EventPublicLeaderboard) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EventPublicLeaderboard) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EventPublicLeaderboard) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTitle

`func (o *EventPublicLeaderboard) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *EventPublicLeaderboard) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *EventPublicLeaderboard) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


