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

// checks if the BillableMetricGroupValuesInnerOneOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillableMetricGroupValuesInnerOneOf{}

// BillableMetricGroupValuesInnerOneOf Second dimension of group.
type BillableMetricGroupValuesInnerOneOf struct {
	// Name of the event property used to group values.
	Key string `json:"key"`
	// Array of strings representing all possible values.
	Values []string `json:"values"`
}

type _BillableMetricGroupValuesInnerOneOf BillableMetricGroupValuesInnerOneOf

// NewBillableMetricGroupValuesInnerOneOf instantiates a new BillableMetricGroupValuesInnerOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillableMetricGroupValuesInnerOneOf(key string, values []string) *BillableMetricGroupValuesInnerOneOf {
	this := BillableMetricGroupValuesInnerOneOf{}
	this.Key = key
	this.Values = values
	return &this
}

// NewBillableMetricGroupValuesInnerOneOfWithDefaults instantiates a new BillableMetricGroupValuesInnerOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillableMetricGroupValuesInnerOneOfWithDefaults() *BillableMetricGroupValuesInnerOneOf {
	this := BillableMetricGroupValuesInnerOneOf{}
	return &this
}

// GetKey returns the Key field value
func (o *BillableMetricGroupValuesInnerOneOf) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *BillableMetricGroupValuesInnerOneOf) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *BillableMetricGroupValuesInnerOneOf) SetKey(v string) {
	o.Key = v
}

// GetValues returns the Values field value
func (o *BillableMetricGroupValuesInnerOneOf) GetValues() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *BillableMetricGroupValuesInnerOneOf) GetValuesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *BillableMetricGroupValuesInnerOneOf) SetValues(v []string) {
	o.Values = v
}

func (o BillableMetricGroupValuesInnerOneOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillableMetricGroupValuesInnerOneOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	toSerialize["values"] = o.Values
	return toSerialize, nil
}

func (o *BillableMetricGroupValuesInnerOneOf) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"key",
		"values",
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

	varBillableMetricGroupValuesInnerOneOf := _BillableMetricGroupValuesInnerOneOf{}

	err = json.Unmarshal(bytes, &varBillableMetricGroupValuesInnerOneOf)

	if err != nil {
		return err
	}

	*o = BillableMetricGroupValuesInnerOneOf(varBillableMetricGroupValuesInnerOneOf)

	return err
}

type NullableBillableMetricGroupValuesInnerOneOf struct {
	value *BillableMetricGroupValuesInnerOneOf
	isSet bool
}

func (v NullableBillableMetricGroupValuesInnerOneOf) Get() *BillableMetricGroupValuesInnerOneOf {
	return v.value
}

func (v *NullableBillableMetricGroupValuesInnerOneOf) Set(val *BillableMetricGroupValuesInnerOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBillableMetricGroupValuesInnerOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBillableMetricGroupValuesInnerOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillableMetricGroupValuesInnerOneOf(val *BillableMetricGroupValuesInnerOneOf) *NullableBillableMetricGroupValuesInnerOneOf {
	return &NullableBillableMetricGroupValuesInnerOneOf{value: val, isSet: true}
}

func (v NullableBillableMetricGroupValuesInnerOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillableMetricGroupValuesInnerOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


