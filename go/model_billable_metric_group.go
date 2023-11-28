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

// checks if the BillableMetricGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillableMetricGroup{}

// BillableMetricGroup Group with one or two dimensions, used to apply differentiated pricing based on additional properties of the billable metric.
type BillableMetricGroup struct {
	// Name of the event property used to group values.
	Key string `json:"key"`
	// Array of strings or objects representing all possible values.
	Values []BillableMetricGroupValuesInner `json:"values"`
}

type _BillableMetricGroup BillableMetricGroup

// NewBillableMetricGroup instantiates a new BillableMetricGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillableMetricGroup(key string, values []BillableMetricGroupValuesInner) *BillableMetricGroup {
	this := BillableMetricGroup{}
	this.Key = key
	this.Values = values
	return &this
}

// NewBillableMetricGroupWithDefaults instantiates a new BillableMetricGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillableMetricGroupWithDefaults() *BillableMetricGroup {
	this := BillableMetricGroup{}
	return &this
}

// GetKey returns the Key field value
func (o *BillableMetricGroup) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *BillableMetricGroup) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *BillableMetricGroup) SetKey(v string) {
	o.Key = v
}

// GetValues returns the Values field value
func (o *BillableMetricGroup) GetValues() []BillableMetricGroupValuesInner {
	if o == nil {
		var ret []BillableMetricGroupValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *BillableMetricGroup) GetValuesOk() ([]BillableMetricGroupValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *BillableMetricGroup) SetValues(v []BillableMetricGroupValuesInner) {
	o.Values = v
}

func (o BillableMetricGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillableMetricGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	toSerialize["values"] = o.Values
	return toSerialize, nil
}

func (o *BillableMetricGroup) UnmarshalJSON(bytes []byte) (err error) {
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

	varBillableMetricGroup := _BillableMetricGroup{}

	err = json.Unmarshal(bytes, &varBillableMetricGroup)

	if err != nil {
		return err
	}

	*o = BillableMetricGroup(varBillableMetricGroup)

	return err
}

type NullableBillableMetricGroup struct {
	value *BillableMetricGroup
	isSet bool
}

func (v NullableBillableMetricGroup) Get() *BillableMetricGroup {
	return v.value
}

func (v *NullableBillableMetricGroup) Set(val *BillableMetricGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableBillableMetricGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableBillableMetricGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillableMetricGroup(val *BillableMetricGroup) *NullableBillableMetricGroup {
	return &NullableBillableMetricGroup{value: val, isSet: true}
}

func (v NullableBillableMetricGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillableMetricGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


