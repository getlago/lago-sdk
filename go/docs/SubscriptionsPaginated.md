# SubscriptionsPaginated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subscriptions** | [**[]SubscriptionObject**](SubscriptionObject.md) |  | 
**Meta** | [**PaginationMeta**](PaginationMeta.md) |  | 

## Methods

### NewSubscriptionsPaginated

`func NewSubscriptionsPaginated(subscriptions []SubscriptionObject, meta PaginationMeta, ) *SubscriptionsPaginated`

NewSubscriptionsPaginated instantiates a new SubscriptionsPaginated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionsPaginatedWithDefaults

`func NewSubscriptionsPaginatedWithDefaults() *SubscriptionsPaginated`

NewSubscriptionsPaginatedWithDefaults instantiates a new SubscriptionsPaginated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptions

`func (o *SubscriptionsPaginated) GetSubscriptions() []SubscriptionObject`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *SubscriptionsPaginated) GetSubscriptionsOk() (*[]SubscriptionObject, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *SubscriptionsPaginated) SetSubscriptions(v []SubscriptionObject)`

SetSubscriptions sets Subscriptions field to given value.


### GetMeta

`func (o *SubscriptionsPaginated) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *SubscriptionsPaginated) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *SubscriptionsPaginated) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


