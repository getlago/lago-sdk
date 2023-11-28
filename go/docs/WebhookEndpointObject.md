# WebhookEndpointObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LagoId** | **string** | Unique identifier assigned to the wallet within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the webhook endpoint&#39;s record within the Lago system. | 
**LagoOrganizationId** | **string** | Unique identifier assigned to the organization attached to the webhook endpoint within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the organization’s record within the Lago system. | 
**WebhookUrl** | **string** | The name of the wallet. | 
**SignatureAlgo** | Pointer to **string** | The signature algo for the webhook. | [optional] 
**CreatedAt** | **time.Time** | The date of the webhook endpoint creation, represented in ISO 8601 datetime format and expressed in Coordinated Universal Time (UTC). | 

## Methods

### NewWebhookEndpointObject

`func NewWebhookEndpointObject(lagoId string, lagoOrganizationId string, webhookUrl string, createdAt time.Time, ) *WebhookEndpointObject`

NewWebhookEndpointObject instantiates a new WebhookEndpointObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookEndpointObjectWithDefaults

`func NewWebhookEndpointObjectWithDefaults() *WebhookEndpointObject`

NewWebhookEndpointObjectWithDefaults instantiates a new WebhookEndpointObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLagoId

`func (o *WebhookEndpointObject) GetLagoId() string`

GetLagoId returns the LagoId field if non-nil, zero value otherwise.

### GetLagoIdOk

`func (o *WebhookEndpointObject) GetLagoIdOk() (*string, bool)`

GetLagoIdOk returns a tuple with the LagoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagoId

`func (o *WebhookEndpointObject) SetLagoId(v string)`

SetLagoId sets LagoId field to given value.


### GetLagoOrganizationId

`func (o *WebhookEndpointObject) GetLagoOrganizationId() string`

GetLagoOrganizationId returns the LagoOrganizationId field if non-nil, zero value otherwise.

### GetLagoOrganizationIdOk

`func (o *WebhookEndpointObject) GetLagoOrganizationIdOk() (*string, bool)`

GetLagoOrganizationIdOk returns a tuple with the LagoOrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagoOrganizationId

`func (o *WebhookEndpointObject) SetLagoOrganizationId(v string)`

SetLagoOrganizationId sets LagoOrganizationId field to given value.


### GetWebhookUrl

`func (o *WebhookEndpointObject) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *WebhookEndpointObject) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *WebhookEndpointObject) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.


### GetSignatureAlgo

`func (o *WebhookEndpointObject) GetSignatureAlgo() string`

GetSignatureAlgo returns the SignatureAlgo field if non-nil, zero value otherwise.

### GetSignatureAlgoOk

`func (o *WebhookEndpointObject) GetSignatureAlgoOk() (*string, bool)`

GetSignatureAlgoOk returns a tuple with the SignatureAlgo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureAlgo

`func (o *WebhookEndpointObject) SetSignatureAlgo(v string)`

SetSignatureAlgo sets SignatureAlgo field to given value.

### HasSignatureAlgo

`func (o *WebhookEndpointObject) HasSignatureAlgo() bool`

HasSignatureAlgo returns a boolean if a field has been set.

### GetCreatedAt

`func (o *WebhookEndpointObject) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WebhookEndpointObject) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WebhookEndpointObject) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


