/*
Lago API documentation

Lago API allows your application to push customer information and metrics (events) from your application to the billing application.

API version: 0.52.0-beta
Contact: tech@getlago.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lagoapi

import (
	"encoding/json"
	"fmt"
)

// checks if the EventBatchInputEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventBatchInputEvent{}

// EventBatchInputEvent struct for EventBatchInputEvent
type EventBatchInputEvent struct {
	// This field represents a unique identifier for the event. It is crucial for ensuring idempotency, meaning that each event can be uniquely identified and processed without causing any unintended side effects.
	TransactionId string `json:"transaction_id"`
	// The customer external unique identifier (provided by your own application). This field is optional if you send the `external_subscription_ids`, targeting a specific subscription.
	ExternalCustomerId *string `json:"external_customer_id,omitempty"`
	// Array of unique identifiers of the targeted subscriptions within your application.
	ExternalSubscriptionIds []string `json:"external_subscription_ids"`
	// The code that identifies a targeted billable metric. It is essential that this code matches the `code` property of one of your active billable metrics. If the provided code does not correspond to any active billable metric, it will be ignored during the process.
	Code string `json:"code"`
	// This field captures the Unix timestamp in seconds indicating the occurrence of the event in Coordinated Universal Time (UTC). If this timestamp is not provided, the API will automatically set it to the time of event reception.
	Timestamp *int32 `json:"timestamp,omitempty"`
	Properties *EventBatchInputEventProperties `json:"properties,omitempty"`
}

type _EventBatchInputEvent EventBatchInputEvent

// NewEventBatchInputEvent instantiates a new EventBatchInputEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventBatchInputEvent(transactionId string, externalSubscriptionIds []string, code string) *EventBatchInputEvent {
	this := EventBatchInputEvent{}
	this.TransactionId = transactionId
	this.ExternalSubscriptionIds = externalSubscriptionIds
	this.Code = code
	return &this
}

// NewEventBatchInputEventWithDefaults instantiates a new EventBatchInputEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventBatchInputEventWithDefaults() *EventBatchInputEvent {
	this := EventBatchInputEvent{}
	return &this
}

// GetTransactionId returns the TransactionId field value
func (o *EventBatchInputEvent) GetTransactionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value
// and a boolean to check if the value has been set.
func (o *EventBatchInputEvent) GetTransactionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionId, true
}

// SetTransactionId sets field value
func (o *EventBatchInputEvent) SetTransactionId(v string) {
	o.TransactionId = v
}

// GetExternalCustomerId returns the ExternalCustomerId field value if set, zero value otherwise.
func (o *EventBatchInputEvent) GetExternalCustomerId() string {
	if o == nil || IsNil(o.ExternalCustomerId) {
		var ret string
		return ret
	}
	return *o.ExternalCustomerId
}

// GetExternalCustomerIdOk returns a tuple with the ExternalCustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventBatchInputEvent) GetExternalCustomerIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalCustomerId) {
		return nil, false
	}
	return o.ExternalCustomerId, true
}

// HasExternalCustomerId returns a boolean if a field has been set.
func (o *EventBatchInputEvent) HasExternalCustomerId() bool {
	if o != nil && !IsNil(o.ExternalCustomerId) {
		return true
	}

	return false
}

// SetExternalCustomerId gets a reference to the given string and assigns it to the ExternalCustomerId field.
func (o *EventBatchInputEvent) SetExternalCustomerId(v string) {
	o.ExternalCustomerId = &v
}

// GetExternalSubscriptionIds returns the ExternalSubscriptionIds field value
func (o *EventBatchInputEvent) GetExternalSubscriptionIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ExternalSubscriptionIds
}

// GetExternalSubscriptionIdsOk returns a tuple with the ExternalSubscriptionIds field value
// and a boolean to check if the value has been set.
func (o *EventBatchInputEvent) GetExternalSubscriptionIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExternalSubscriptionIds, true
}

// SetExternalSubscriptionIds sets field value
func (o *EventBatchInputEvent) SetExternalSubscriptionIds(v []string) {
	o.ExternalSubscriptionIds = v
}

// GetCode returns the Code field value
func (o *EventBatchInputEvent) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *EventBatchInputEvent) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *EventBatchInputEvent) SetCode(v string) {
	o.Code = v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *EventBatchInputEvent) GetTimestamp() int32 {
	if o == nil || IsNil(o.Timestamp) {
		var ret int32
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventBatchInputEvent) GetTimestampOk() (*int32, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *EventBatchInputEvent) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int32 and assigns it to the Timestamp field.
func (o *EventBatchInputEvent) SetTimestamp(v int32) {
	o.Timestamp = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *EventBatchInputEvent) GetProperties() EventBatchInputEventProperties {
	if o == nil || IsNil(o.Properties) {
		var ret EventBatchInputEventProperties
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventBatchInputEvent) GetPropertiesOk() (*EventBatchInputEventProperties, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *EventBatchInputEvent) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given EventBatchInputEventProperties and assigns it to the Properties field.
func (o *EventBatchInputEvent) SetProperties(v EventBatchInputEventProperties) {
	o.Properties = &v
}

func (o EventBatchInputEvent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventBatchInputEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["transaction_id"] = o.TransactionId
	if !IsNil(o.ExternalCustomerId) {
		toSerialize["external_customer_id"] = o.ExternalCustomerId
	}
	toSerialize["external_subscription_ids"] = o.ExternalSubscriptionIds
	toSerialize["code"] = o.Code
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	return toSerialize, nil
}

func (o *EventBatchInputEvent) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"transaction_id",
		"external_subscription_ids",
		"code",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varEventBatchInputEvent := _EventBatchInputEvent{}

	err = json.Unmarshal(bytes, &varEventBatchInputEvent)

	if err != nil {
		return err
	}

	*o = EventBatchInputEvent(varEventBatchInputEvent)

	return err
}

type NullableEventBatchInputEvent struct {
	value *EventBatchInputEvent
	isSet bool
}

func (v NullableEventBatchInputEvent) Get() *EventBatchInputEvent {
	return v.value
}

func (v *NullableEventBatchInputEvent) Set(val *EventBatchInputEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableEventBatchInputEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableEventBatchInputEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventBatchInputEvent(val *EventBatchInputEvent) *NullableEventBatchInputEvent {
	return &NullableEventBatchInputEvent{value: val, isSet: true}
}

func (v NullableEventBatchInputEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventBatchInputEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


