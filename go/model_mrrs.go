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

// checks if the Mrrs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Mrrs{}

// Mrrs struct for Mrrs
type Mrrs struct {
	Mrrs []MrrObject `json:"mrrs"`
}

type _Mrrs Mrrs

// NewMrrs instantiates a new Mrrs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMrrs(mrrs []MrrObject) *Mrrs {
	this := Mrrs{}
	this.Mrrs = mrrs
	return &this
}

// NewMrrsWithDefaults instantiates a new Mrrs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMrrsWithDefaults() *Mrrs {
	this := Mrrs{}
	return &this
}

// GetMrrs returns the Mrrs field value
func (o *Mrrs) GetMrrs() []MrrObject {
	if o == nil {
		var ret []MrrObject
		return ret
	}

	return o.Mrrs
}

// GetMrrsOk returns a tuple with the Mrrs field value
// and a boolean to check if the value has been set.
func (o *Mrrs) GetMrrsOk() ([]MrrObject, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mrrs, true
}

// SetMrrs sets field value
func (o *Mrrs) SetMrrs(v []MrrObject) {
	o.Mrrs = v
}

func (o Mrrs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Mrrs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mrrs"] = o.Mrrs
	return toSerialize, nil
}

func (o *Mrrs) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"mrrs",
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

	varMrrs := _Mrrs{}

	err = json.Unmarshal(bytes, &varMrrs)

	if err != nil {
		return err
	}

	*o = Mrrs(varMrrs)

	return err
}

type NullableMrrs struct {
	value *Mrrs
	isSet bool
}

func (v NullableMrrs) Get() *Mrrs {
	return v.value
}

func (v *NullableMrrs) Set(val *Mrrs) {
	v.value = val
	v.isSet = true
}

func (v NullableMrrs) IsSet() bool {
	return v.isSet
}

func (v *NullableMrrs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMrrs(val *Mrrs) *NullableMrrs {
	return &NullableMrrs{value: val, isSet: true}
}

func (v NullableMrrs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMrrs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


