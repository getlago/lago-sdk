# WebhookEndpointsPaginated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebhookEndpoints** | [**[]WebhookEndpointObject**](WebhookEndpointObject.md) |  | 
**Meta** | [**PaginationMeta**](PaginationMeta.md) |  | 

## Methods

### NewWebhookEndpointsPaginated

`func NewWebhookEndpointsPaginated(webhookEndpoints []WebhookEndpointObject, meta PaginationMeta, ) *WebhookEndpointsPaginated`

NewWebhookEndpointsPaginated instantiates a new WebhookEndpointsPaginated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookEndpointsPaginatedWithDefaults

`func NewWebhookEndpointsPaginatedWithDefaults() *WebhookEndpointsPaginated`

NewWebhookEndpointsPaginatedWithDefaults instantiates a new WebhookEndpointsPaginated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebhookEndpoints

`func (o *WebhookEndpointsPaginated) GetWebhookEndpoints() []WebhookEndpointObject`

GetWebhookEndpoints returns the WebhookEndpoints field if non-nil, zero value otherwise.

### GetWebhookEndpointsOk

`func (o *WebhookEndpointsPaginated) GetWebhookEndpointsOk() (*[]WebhookEndpointObject, bool)`

GetWebhookEndpointsOk returns a tuple with the WebhookEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookEndpoints

`func (o *WebhookEndpointsPaginated) SetWebhookEndpoints(v []WebhookEndpointObject)`

SetWebhookEndpoints sets WebhookEndpoints field to given value.


### GetMeta

`func (o *WebhookEndpointsPaginated) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *WebhookEndpointsPaginated) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *WebhookEndpointsPaginated) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


