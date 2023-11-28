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

// checks if the EventEstimateFeesInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventEstimateFeesInput{}

// EventEstimateFeesInput struct for EventEstimateFeesInput
type EventEstimateFeesInput struct {
	Event EventEstimateFeesInputEvent `json:"event"`
}

type _EventEstimateFeesInput EventEstimateFeesInput

// NewEventEstimateFeesInput instantiates a new EventEstimateFeesInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventEstimateFeesInput(event EventEstimateFeesInputEvent) *EventEstimateFeesInput {
	this := EventEstimateFeesInput{}
	this.Event = event
	return &this
}

// NewEventEstimateFeesInputWithDefaults instantiates a new EventEstimateFeesInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventEstimateFeesInputWithDefaults() *EventEstimateFeesInput {
	this := EventEstimateFeesInput{}
	return &this
}

// GetEvent returns the Event field value
func (o *EventEstimateFeesInput) GetEvent() EventEstimateFeesInputEvent {
	if o == nil {
		var ret EventEstimateFeesInputEvent
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *EventEstimateFeesInput) GetEventOk() (*EventEstimateFeesInputEvent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *EventEstimateFeesInput) SetEvent(v EventEstimateFeesInputEvent) {
	o.Event = v
}

func (o EventEstimateFeesInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventEstimateFeesInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["event"] = o.Event
	return toSerialize, nil
}

func (o *EventEstimateFeesInput) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"event",
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

	varEventEstimateFeesInput := _EventEstimateFeesInput{}

	err = json.Unmarshal(bytes, &varEventEstimateFeesInput)

	if err != nil {
		return err
	}

	*o = EventEstimateFeesInput(varEventEstimateFeesInput)

	return err
}

type NullableEventEstimateFeesInput struct {
	value *EventEstimateFeesInput
	isSet bool
}

func (v NullableEventEstimateFeesInput) Get() *EventEstimateFeesInput {
	return v.value
}

func (v *NullableEventEstimateFeesInput) Set(val *EventEstimateFeesInput) {
	v.value = val
	v.isSet = true
}

func (v NullableEventEstimateFeesInput) IsSet() bool {
	return v.isSet
}

func (v *NullableEventEstimateFeesInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventEstimateFeesInput(val *EventEstimateFeesInput) *NullableEventEstimateFeesInput {
	return &NullableEventEstimateFeesInput{value: val, isSet: true}
}

func (v NullableEventEstimateFeesInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventEstimateFeesInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


