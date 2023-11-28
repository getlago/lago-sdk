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

// checks if the AddOn type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddOn{}

// AddOn struct for AddOn
type AddOn struct {
	AddOn AddOnObject `json:"add_on"`
}

type _AddOn AddOn

// NewAddOn instantiates a new AddOn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddOn(addOn AddOnObject) *AddOn {
	this := AddOn{}
	this.AddOn = addOn
	return &this
}

// NewAddOnWithDefaults instantiates a new AddOn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddOnWithDefaults() *AddOn {
	this := AddOn{}
	return &this
}

// GetAddOn returns the AddOn field value
func (o *AddOn) GetAddOn() AddOnObject {
	if o == nil {
		var ret AddOnObject
		return ret
	}

	return o.AddOn
}

// GetAddOnOk returns a tuple with the AddOn field value
// and a boolean to check if the value has been set.
func (o *AddOn) GetAddOnOk() (*AddOnObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AddOn, true
}

// SetAddOn sets field value
func (o *AddOn) SetAddOn(v AddOnObject) {
	o.AddOn = v
}

func (o AddOn) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddOn) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["add_on"] = o.AddOn
	return toSerialize, nil
}

func (o *AddOn) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"add_on",
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

	varAddOn := _AddOn{}

	err = json.Unmarshal(bytes, &varAddOn)

	if err != nil {
		return err
	}

	*o = AddOn(varAddOn)

	return err
}

type NullableAddOn struct {
	value *AddOn
	isSet bool
}

func (v NullableAddOn) Get() *AddOn {
	return v.value
}

func (v *NullableAddOn) Set(val *AddOn) {
	v.value = val
	v.isSet = true
}

func (v NullableAddOn) IsSet() bool {
	return v.isSet
}

func (v *NullableAddOn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddOn(val *AddOn) *NullableAddOn {
	return &NullableAddOn{value: val, isSet: true}
}

func (v NullableAddOn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddOn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

