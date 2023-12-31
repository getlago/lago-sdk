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

// checks if the BillableMetricsPaginated type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillableMetricsPaginated{}

// BillableMetricsPaginated struct for BillableMetricsPaginated
type BillableMetricsPaginated struct {
	BillableMetrics []BillableMetricObject `json:"billable_metrics"`
	Meta PaginationMeta `json:"meta"`
}

type _BillableMetricsPaginated BillableMetricsPaginated

// NewBillableMetricsPaginated instantiates a new BillableMetricsPaginated object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillableMetricsPaginated(billableMetrics []BillableMetricObject, meta PaginationMeta) *BillableMetricsPaginated {
	this := BillableMetricsPaginated{}
	this.BillableMetrics = billableMetrics
	this.Meta = meta
	return &this
}

// NewBillableMetricsPaginatedWithDefaults instantiates a new BillableMetricsPaginated object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillableMetricsPaginatedWithDefaults() *BillableMetricsPaginated {
	this := BillableMetricsPaginated{}
	return &this
}

// GetBillableMetrics returns the BillableMetrics field value
func (o *BillableMetricsPaginated) GetBillableMetrics() []BillableMetricObject {
	if o == nil {
		var ret []BillableMetricObject
		return ret
	}

	return o.BillableMetrics
}

// GetBillableMetricsOk returns a tuple with the BillableMetrics field value
// and a boolean to check if the value has been set.
func (o *BillableMetricsPaginated) GetBillableMetricsOk() ([]BillableMetricObject, bool) {
	if o == nil {
		return nil, false
	}
	return o.BillableMetrics, true
}

// SetBillableMetrics sets field value
func (o *BillableMetricsPaginated) SetBillableMetrics(v []BillableMetricObject) {
	o.BillableMetrics = v
}

// GetMeta returns the Meta field value
func (o *BillableMetricsPaginated) GetMeta() PaginationMeta {
	if o == nil {
		var ret PaginationMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *BillableMetricsPaginated) GetMetaOk() (*PaginationMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *BillableMetricsPaginated) SetMeta(v PaginationMeta) {
	o.Meta = v
}

func (o BillableMetricsPaginated) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillableMetricsPaginated) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["billable_metrics"] = o.BillableMetrics
	toSerialize["meta"] = o.Meta
	return toSerialize, nil
}

func (o *BillableMetricsPaginated) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"billable_metrics",
		"meta",
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

	varBillableMetricsPaginated := _BillableMetricsPaginated{}

	err = json.Unmarshal(bytes, &varBillableMetricsPaginated)

	if err != nil {
		return err
	}

	*o = BillableMetricsPaginated(varBillableMetricsPaginated)

	return err
}

type NullableBillableMetricsPaginated struct {
	value *BillableMetricsPaginated
	isSet bool
}

func (v NullableBillableMetricsPaginated) Get() *BillableMetricsPaginated {
	return v.value
}

func (v *NullableBillableMetricsPaginated) Set(val *BillableMetricsPaginated) {
	v.value = val
	v.isSet = true
}

func (v NullableBillableMetricsPaginated) IsSet() bool {
	return v.isSet
}

func (v *NullableBillableMetricsPaginated) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillableMetricsPaginated(val *BillableMetricsPaginated) *NullableBillableMetricsPaginated {
	return &NullableBillableMetricsPaginated{value: val, isSet: true}
}

func (v NullableBillableMetricsPaginated) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillableMetricsPaginated) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


