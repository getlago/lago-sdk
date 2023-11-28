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

// checks if the SubscriptionCreateInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionCreateInput{}

// SubscriptionCreateInput struct for SubscriptionCreateInput
type SubscriptionCreateInput struct {
	Subscription SubscriptionCreateInputSubscription `json:"subscription"`
}

type _SubscriptionCreateInput SubscriptionCreateInput

// NewSubscriptionCreateInput instantiates a new SubscriptionCreateInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionCreateInput(subscription SubscriptionCreateInputSubscription) *SubscriptionCreateInput {
	this := SubscriptionCreateInput{}
	this.Subscription = subscription
	return &this
}

// NewSubscriptionCreateInputWithDefaults instantiates a new SubscriptionCreateInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionCreateInputWithDefaults() *SubscriptionCreateInput {
	this := SubscriptionCreateInput{}
	return &this
}

// GetSubscription returns the Subscription field value
func (o *SubscriptionCreateInput) GetSubscription() SubscriptionCreateInputSubscription {
	if o == nil {
		var ret SubscriptionCreateInputSubscription
		return ret
	}

	return o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value
// and a boolean to check if the value has been set.
func (o *SubscriptionCreateInput) GetSubscriptionOk() (*SubscriptionCreateInputSubscription, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subscription, true
}

// SetSubscription sets field value
func (o *SubscriptionCreateInput) SetSubscription(v SubscriptionCreateInputSubscription) {
	o.Subscription = v
}

func (o SubscriptionCreateInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionCreateInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["subscription"] = o.Subscription
	return toSerialize, nil
}

func (o *SubscriptionCreateInput) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"subscription",
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

	varSubscriptionCreateInput := _SubscriptionCreateInput{}

	err = json.Unmarshal(bytes, &varSubscriptionCreateInput)

	if err != nil {
		return err
	}

	*o = SubscriptionCreateInput(varSubscriptionCreateInput)

	return err
}

type NullableSubscriptionCreateInput struct {
	value *SubscriptionCreateInput
	isSet bool
}

func (v NullableSubscriptionCreateInput) Get() *SubscriptionCreateInput {
	return v.value
}

func (v *NullableSubscriptionCreateInput) Set(val *SubscriptionCreateInput) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionCreateInput) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionCreateInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionCreateInput(val *SubscriptionCreateInput) *NullableSubscriptionCreateInput {
	return &NullableSubscriptionCreateInput{value: val, isSet: true}
}

func (v NullableSubscriptionCreateInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionCreateInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

