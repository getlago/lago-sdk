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

// checks if the EventEstimateFeesInputEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventEstimateFeesInputEvent{}

// EventEstimateFeesInputEvent struct for EventEstimateFeesInputEvent
type EventEstimateFeesInputEvent struct {
	// The code that identifies a targeted billable metric. It is essential that this code matches the `code` property of one of your active billable metrics. If the provided code does not correspond to any active billable metric, it will be ignored during the process.
	Code string `json:"code"`
	// The customer external unique identifier (provided by your own application). This field is optional if you send the `external_subscription_id`, targeting a specific subscription.
	ExternalCustomerId *string `json:"external_customer_id,omitempty"`
	// The unique identifier of the subscription within your application. It is a mandatory field when the customer possesses multiple subscriptions or when the `external_customer_id` is not provided.
	ExternalSubscriptionId *string `json:"external_subscription_id,omitempty"`
	// This field represents additional properties associated with the event, which are utilized in the calculation of the final fee. This object becomes mandatory when the targeted billable metric employs a `sum_agg`, `max_agg`, or `unique_count_agg` aggregation method. However, when using a simple `count_agg`, this object is not required.
	Properties map[string]interface{} `json:"properties,omitempty"`
}

type _EventEstimateFeesInputEvent EventEstimateFeesInputEvent

// NewEventEstimateFeesInputEvent instantiates a new EventEstimateFeesInputEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventEstimateFeesInputEvent(code string) *EventEstimateFeesInputEvent {
	this := EventEstimateFeesInputEvent{}
	this.Code = code
	return &this
}

// NewEventEstimateFeesInputEventWithDefaults instantiates a new EventEstimateFeesInputEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventEstimateFeesInputEventWithDefaults() *EventEstimateFeesInputEvent {
	this := EventEstimateFeesInputEvent{}
	return &this
}

// GetCode returns the Code field value
func (o *EventEstimateFeesInputEvent) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *EventEstimateFeesInputEvent) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *EventEstimateFeesInputEvent) SetCode(v string) {
	o.Code = v
}

// GetExternalCustomerId returns the ExternalCustomerId field value if set, zero value otherwise.
func (o *EventEstimateFeesInputEvent) GetExternalCustomerId() string {
	if o == nil || IsNil(o.ExternalCustomerId) {
		var ret string
		return ret
	}
	return *o.ExternalCustomerId
}

// GetExternalCustomerIdOk returns a tuple with the ExternalCustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventEstimateFeesInputEvent) GetExternalCustomerIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalCustomerId) {
		return nil, false
	}
	return o.ExternalCustomerId, true
}

// HasExternalCustomerId returns a boolean if a field has been set.
func (o *EventEstimateFeesInputEvent) HasExternalCustomerId() bool {
	if o != nil && !IsNil(o.ExternalCustomerId) {
		return true
	}

	return false
}

// SetExternalCustomerId gets a reference to the given string and assigns it to the ExternalCustomerId field.
func (o *EventEstimateFeesInputEvent) SetExternalCustomerId(v string) {
	o.ExternalCustomerId = &v
}

// GetExternalSubscriptionId returns the ExternalSubscriptionId field value if set, zero value otherwise.
func (o *EventEstimateFeesInputEvent) GetExternalSubscriptionId() string {
	if o == nil || IsNil(o.ExternalSubscriptionId) {
		var ret string
		return ret
	}
	return *o.ExternalSubscriptionId
}

// GetExternalSubscriptionIdOk returns a tuple with the ExternalSubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventEstimateFeesInputEvent) GetExternalSubscriptionIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalSubscriptionId) {
		return nil, false
	}
	return o.ExternalSubscriptionId, true
}

// HasExternalSubscriptionId returns a boolean if a field has been set.
func (o *EventEstimateFeesInputEvent) HasExternalSubscriptionId() bool {
	if o != nil && !IsNil(o.ExternalSubscriptionId) {
		return true
	}

	return false
}

// SetExternalSubscriptionId gets a reference to the given string and assigns it to the ExternalSubscriptionId field.
func (o *EventEstimateFeesInputEvent) SetExternalSubscriptionId(v string) {
	o.ExternalSubscriptionId = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *EventEstimateFeesInputEvent) GetProperties() map[string]interface{} {
	if o == nil || IsNil(o.Properties) {
		var ret map[string]interface{}
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventEstimateFeesInputEvent) GetPropertiesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Properties) {
		return map[string]interface{}{}, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *EventEstimateFeesInputEvent) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given map[string]interface{} and assigns it to the Properties field.
func (o *EventEstimateFeesInputEvent) SetProperties(v map[string]interface{}) {
	o.Properties = v
}

func (o EventEstimateFeesInputEvent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventEstimateFeesInputEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	if !IsNil(o.ExternalCustomerId) {
		toSerialize["external_customer_id"] = o.ExternalCustomerId
	}
	if !IsNil(o.ExternalSubscriptionId) {
		toSerialize["external_subscription_id"] = o.ExternalSubscriptionId
	}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	return toSerialize, nil
}

func (o *EventEstimateFeesInputEvent) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varEventEstimateFeesInputEvent := _EventEstimateFeesInputEvent{}

	err = json.Unmarshal(bytes, &varEventEstimateFeesInputEvent)

	if err != nil {
		return err
	}

	*o = EventEstimateFeesInputEvent(varEventEstimateFeesInputEvent)

	return err
}

type NullableEventEstimateFeesInputEvent struct {
	value *EventEstimateFeesInputEvent
	isSet bool
}

func (v NullableEventEstimateFeesInputEvent) Get() *EventEstimateFeesInputEvent {
	return v.value
}

func (v *NullableEventEstimateFeesInputEvent) Set(val *EventEstimateFeesInputEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableEventEstimateFeesInputEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableEventEstimateFeesInputEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventEstimateFeesInputEvent(val *EventEstimateFeesInputEvent) *NullableEventEstimateFeesInputEvent {
	return &NullableEventEstimateFeesInputEvent{value: val, isSet: true}
}

func (v NullableEventEstimateFeesInputEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventEstimateFeesInputEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


