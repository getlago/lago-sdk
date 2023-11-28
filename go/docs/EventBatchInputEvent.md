# EventBatchInputEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionId** | **string** | This field represents a unique identifier for the event. It is crucial for ensuring idempotency, meaning that each event can be uniquely identified and processed without causing any unintended side effects. | 
**ExternalCustomerId** | Pointer to **string** | The customer external unique identifier (provided by your own application). This field is optional if you send the &#x60;external_subscription_ids&#x60;, targeting a specific subscription. | [optional] 
**ExternalSubscriptionIds** | **[]string** | Array of unique identifiers of the targeted subscriptions within your application. | 
**Code** | **string** | The code that identifies a targeted billable metric. It is essential that this code matches the &#x60;code&#x60; property of one of your active billable metrics. If the provided code does not correspond to any active billable metric, it will be ignored during the process. | 
**Timestamp** | Pointer to **int32** | This field captures the Unix timestamp in seconds indicating the occurrence of the event in Coordinated Universal Time (UTC). If this timestamp is not provided, the API will automatically set it to the time of event reception. | [optional] 
**Properties** | Pointer to [**EventBatchInputEventProperties**](EventBatchInputEventProperties.md) |  | [optional] 

## Methods

### NewEventBatchInputEvent

`func NewEventBatchInputEvent(transactionId string, externalSubscriptionIds []string, code string, ) *EventBatchInputEvent`

NewEventBatchInputEvent instantiates a new EventBatchInputEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventBatchInputEventWithDefaults

`func NewEventBatchInputEventWithDefaults() *EventBatchInputEvent`

NewEventBatchInputEventWithDefaults instantiates a new EventBatchInputEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionId

`func (o *EventBatchInputEvent) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *EventBatchInputEvent) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *EventBatchInputEvent) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetExternalCustomerId

`func (o *EventBatchInputEvent) GetExternalCustomerId() string`

GetExternalCustomerId returns the ExternalCustomerId field if non-nil, zero value otherwise.

### GetExternalCustomerIdOk

`func (o *EventBatchInputEvent) GetExternalCustomerIdOk() (*string, bool)`

GetExternalCustomerIdOk returns a tuple with the ExternalCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalCustomerId

`func (o *EventBatchInputEvent) SetExternalCustomerId(v string)`

SetExternalCustomerId sets ExternalCustomerId field to given value.

### HasExternalCustomerId

`func (o *EventBatchInputEvent) HasExternalCustomerId() bool`

HasExternalCustomerId returns a boolean if a field has been set.

### GetExternalSubscriptionIds

`func (o *EventBatchInputEvent) GetExternalSubscriptionIds() []string`

GetExternalSubscriptionIds returns the ExternalSubscriptionIds field if non-nil, zero value otherwise.

### GetExternalSubscriptionIdsOk

`func (o *EventBatchInputEvent) GetExternalSubscriptionIdsOk() (*[]string, bool)`

GetExternalSubscriptionIdsOk returns a tuple with the ExternalSubscriptionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSubscriptionIds

`func (o *EventBatchInputEvent) SetExternalSubscriptionIds(v []string)`

SetExternalSubscriptionIds sets ExternalSubscriptionIds field to given value.


### GetCode

`func (o *EventBatchInputEvent) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *EventBatchInputEvent) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *EventBatchInputEvent) SetCode(v string)`

SetCode sets Code field to given value.


### GetTimestamp

`func (o *EventBatchInputEvent) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *EventBatchInputEvent) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *EventBatchInputEvent) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *EventBatchInputEvent) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetProperties

`func (o *EventBatchInputEvent) GetProperties() EventBatchInputEventProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *EventBatchInputEvent) GetPropertiesOk() (*EventBatchInputEventProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *EventBatchInputEvent) SetProperties(v EventBatchInputEventProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *EventBatchInputEvent) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


