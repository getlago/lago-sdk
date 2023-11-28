# EventInputEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionId** | **string** | This field represents a unique identifier for the event. It is crucial for ensuring idempotency, meaning that each event can be uniquely identified and processed without causing any unintended side effects. | 
**ExternalCustomerId** | Pointer to **string** | The customer external unique identifier (provided by your own application). This field is optional if you send the &#x60;external_subscription_id&#x60;, targeting a specific subscription. | [optional] 
**ExternalSubscriptionId** | Pointer to **string** | The unique identifier of the subscription within your application. It is a mandatory field when the customer possesses multiple subscriptions or when the &#x60;external_customer_id&#x60; is not provided. | [optional] 
**Code** | **string** | The code that identifies a targeted billable metric. It is essential that this code matches the &#x60;code&#x60; property of one of your active billable metrics. If the provided code does not correspond to any active billable metric, it will be ignored during the process. | 
**Timestamp** | Pointer to [**EventInputEventTimestamp**](EventInputEventTimestamp.md) |  | [optional] 
**Properties** | Pointer to **map[string]string** | This field represents additional properties associated with the event, which are utilized in the calculation of the final fee. This object becomes mandatory when the targeted billable metric employs a &#x60;sum_agg&#x60;, &#x60;max_agg&#x60;, or &#x60;unique_count_agg&#x60; aggregation method. However, when using a simple &#x60;count_agg&#x60;, this object is not required. | [optional] 

## Methods

### NewEventInputEvent

`func NewEventInputEvent(transactionId string, code string, ) *EventInputEvent`

NewEventInputEvent instantiates a new EventInputEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventInputEventWithDefaults

`func NewEventInputEventWithDefaults() *EventInputEvent`

NewEventInputEventWithDefaults instantiates a new EventInputEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionId

`func (o *EventInputEvent) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *EventInputEvent) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *EventInputEvent) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetExternalCustomerId

`func (o *EventInputEvent) GetExternalCustomerId() string`

GetExternalCustomerId returns the ExternalCustomerId field if non-nil, zero value otherwise.

### GetExternalCustomerIdOk

`func (o *EventInputEvent) GetExternalCustomerIdOk() (*string, bool)`

GetExternalCustomerIdOk returns a tuple with the ExternalCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalCustomerId

`func (o *EventInputEvent) SetExternalCustomerId(v string)`

SetExternalCustomerId sets ExternalCustomerId field to given value.

### HasExternalCustomerId

`func (o *EventInputEvent) HasExternalCustomerId() bool`

HasExternalCustomerId returns a boolean if a field has been set.

### GetExternalSubscriptionId

`func (o *EventInputEvent) GetExternalSubscriptionId() string`

GetExternalSubscriptionId returns the ExternalSubscriptionId field if non-nil, zero value otherwise.

### GetExternalSubscriptionIdOk

`func (o *EventInputEvent) GetExternalSubscriptionIdOk() (*string, bool)`

GetExternalSubscriptionIdOk returns a tuple with the ExternalSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSubscriptionId

`func (o *EventInputEvent) SetExternalSubscriptionId(v string)`

SetExternalSubscriptionId sets ExternalSubscriptionId field to given value.

### HasExternalSubscriptionId

`func (o *EventInputEvent) HasExternalSubscriptionId() bool`

HasExternalSubscriptionId returns a boolean if a field has been set.

### GetCode

`func (o *EventInputEvent) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *EventInputEvent) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *EventInputEvent) SetCode(v string)`

SetCode sets Code field to given value.


### GetTimestamp

`func (o *EventInputEvent) GetTimestamp() EventInputEventTimestamp`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *EventInputEvent) GetTimestampOk() (*EventInputEventTimestamp, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *EventInputEvent) SetTimestamp(v EventInputEventTimestamp)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *EventInputEvent) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetProperties

`func (o *EventInputEvent) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *EventInputEvent) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *EventInputEvent) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *EventInputEvent) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


