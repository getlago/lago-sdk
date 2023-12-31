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

// checks if the BillableMetricUpdateInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillableMetricUpdateInput{}

// BillableMetricUpdateInput struct for BillableMetricUpdateInput
type BillableMetricUpdateInput struct {
	BillableMetric BillableMetricBaseInput `json:"billable_metric"`
}

type _BillableMetricUpdateInput BillableMetricUpdateInput

// NewBillableMetricUpdateInput instantiates a new BillableMetricUpdateInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillableMetricUpdateInput(billableMetric BillableMetricBaseInput) *BillableMetricUpdateInput {
	this := BillableMetricUpdateInput{}
	this.BillableMetric = billableMetric
	return &this
}

// NewBillableMetricUpdateInputWithDefaults instantiates a new BillableMetricUpdateInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillableMetricUpdateInputWithDefaults() *BillableMetricUpdateInput {
	this := BillableMetricUpdateInput{}
	return &this
}

// GetBillableMetric returns the BillableMetric field value
func (o *BillableMetricUpdateInput) GetBillableMetric() BillableMetricBaseInput {
	if o == nil {
		var ret BillableMetricBaseInput
		return ret
	}

	return o.BillableMetric
}

// GetBillableMetricOk returns a tuple with the BillableMetric field value
// and a boolean to check if the value has been set.
func (o *BillableMetricUpdateInput) GetBillableMetricOk() (*BillableMetricBaseInput, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BillableMetric, true
}

// SetBillableMetric sets field value
func (o *BillableMetricUpdateInput) SetBillableMetric(v BillableMetricBaseInput) {
	o.BillableMetric = v
}

func (o BillableMetricUpdateInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillableMetricUpdateInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["billable_metric"] = o.BillableMetric
	return toSerialize, nil
}

func (o *BillableMetricUpdateInput) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"billable_metric",
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

	varBillableMetricUpdateInput := _BillableMetricUpdateInput{}

	err = json.Unmarshal(bytes, &varBillableMetricUpdateInput)

	if err != nil {
		return err
	}

	*o = BillableMetricUpdateInput(varBillableMetricUpdateInput)

	return err
}

type NullableBillableMetricUpdateInput struct {
	value *BillableMetricUpdateInput
	isSet bool
}

func (v NullableBillableMetricUpdateInput) Get() *BillableMetricUpdateInput {
	return v.value
}

func (v *NullableBillableMetricUpdateInput) Set(val *BillableMetricUpdateInput) {
	v.value = val
	v.isSet = true
}

func (v NullableBillableMetricUpdateInput) IsSet() bool {
	return v.isSet
}

func (v *NullableBillableMetricUpdateInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillableMetricUpdateInput(val *BillableMetricUpdateInput) *NullableBillableMetricUpdateInput {
	return &NullableBillableMetricUpdateInput{value: val, isSet: true}
}

func (v NullableBillableMetricUpdateInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillableMetricUpdateInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


