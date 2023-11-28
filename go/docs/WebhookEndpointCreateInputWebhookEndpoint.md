# WebhookEndpointCreateInputWebhookEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebhookUrl** | **string** | The URL of the webhook endpoint. | 
**SignatureAlgo** | Pointer to **NullableString** | The signature used for the webhook. If no value is passed, | [optional] 

## Methods

### NewWebhookEndpointCreateInputWebhookEndpoint

`func NewWebhookEndpointCreateInputWebhookEndpoint(webhookUrl string, ) *WebhookEndpointCreateInputWebhookEndpoint`

NewWebhookEndpointCreateInputWebhookEndpoint instantiates a new WebhookEndpointCreateInputWebhookEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookEndpointCreateInputWebhookEndpointWithDefaults

`func NewWebhookEndpointCreateInputWebhookEndpointWithDefaults() *WebhookEndpointCreateInputWebhookEndpoint`

NewWebhookEndpointCreateInputWebhookEndpointWithDefaults instantiates a new WebhookEndpointCreateInputWebhookEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebhookUrl

`func (o *WebhookEndpointCreateInputWebhookEndpoint) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *WebhookEndpointCreateInputWebhookEndpoint) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *WebhookEndpointCreateInputWebhookEndpoint) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.


### GetSignatureAlgo

`func (o *WebhookEndpointCreateInputWebhookEndpoint) GetSignatureAlgo() string`

GetSignatureAlgo returns the SignatureAlgo field if non-nil, zero value otherwise.

### GetSignatureAlgoOk

`func (o *WebhookEndpointCreateInputWebhookEndpoint) GetSignatureAlgoOk() (*string, bool)`

GetSignatureAlgoOk returns a tuple with the SignatureAlgo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureAlgo

`func (o *WebhookEndpointCreateInputWebhookEndpoint) SetSignatureAlgo(v string)`

SetSignatureAlgo sets SignatureAlgo field to given value.

### HasSignatureAlgo

`func (o *WebhookEndpointCreateInputWebhookEndpoint) HasSignatureAlgo() bool`

HasSignatureAlgo returns a boolean if a field has been set.

### SetSignatureAlgoNil

`func (o *WebhookEndpointCreateInputWebhookEndpoint) SetSignatureAlgoNil(b bool)`

 SetSignatureAlgoNil sets the value for SignatureAlgo to be an explicit nil

### UnsetSignatureAlgo
`func (o *WebhookEndpointCreateInputWebhookEndpoint) UnsetSignatureAlgo()`

UnsetSignatureAlgo ensures that no value is present for SignatureAlgo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


