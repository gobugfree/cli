# Integrations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventCatalog** | **[]string** |  | 
**Webhooks** | [**[]IntegrationWebhook**](IntegrationWebhook.md) |  | 

## Methods

### NewIntegrations

`func NewIntegrations(eventCatalog []string, webhooks []IntegrationWebhook, ) *Integrations`

NewIntegrations instantiates a new Integrations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationsWithDefaults

`func NewIntegrationsWithDefaults() *Integrations`

NewIntegrationsWithDefaults instantiates a new Integrations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventCatalog

`func (o *Integrations) GetEventCatalog() []string`

GetEventCatalog returns the EventCatalog field if non-nil, zero value otherwise.

### GetEventCatalogOk

`func (o *Integrations) GetEventCatalogOk() (*[]string, bool)`

GetEventCatalogOk returns a tuple with the EventCatalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCatalog

`func (o *Integrations) SetEventCatalog(v []string)`

SetEventCatalog sets EventCatalog field to given value.


### GetWebhooks

`func (o *Integrations) GetWebhooks() []IntegrationWebhook`

GetWebhooks returns the Webhooks field if non-nil, zero value otherwise.

### GetWebhooksOk

`func (o *Integrations) GetWebhooksOk() (*[]IntegrationWebhook, bool)`

GetWebhooksOk returns a tuple with the Webhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooks

`func (o *Integrations) SetWebhooks(v []IntegrationWebhook)`

SetWebhooks sets Webhooks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


