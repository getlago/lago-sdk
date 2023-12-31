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

// checks if the FeesPaginated type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeesPaginated{}

// FeesPaginated struct for FeesPaginated
type FeesPaginated struct {
	Fees []FeeObject `json:"fees"`
	Meta PaginationMeta `json:"meta"`
}

type _FeesPaginated FeesPaginated

// NewFeesPaginated instantiates a new FeesPaginated object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeesPaginated(fees []FeeObject, meta PaginationMeta) *FeesPaginated {
	this := FeesPaginated{}
	this.Fees = fees
	this.Meta = meta
	return &this
}

// NewFeesPaginatedWithDefaults instantiates a new FeesPaginated object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeesPaginatedWithDefaults() *FeesPaginated {
	this := FeesPaginated{}
	return &this
}

// GetFees returns the Fees field value
func (o *FeesPaginated) GetFees() []FeeObject {
	if o == nil {
		var ret []FeeObject
		return ret
	}

	return o.Fees
}

// GetFeesOk returns a tuple with the Fees field value
// and a boolean to check if the value has been set.
func (o *FeesPaginated) GetFeesOk() ([]FeeObject, bool) {
	if o == nil {
		return nil, false
	}
	return o.Fees, true
}

// SetFees sets field value
func (o *FeesPaginated) SetFees(v []FeeObject) {
	o.Fees = v
}

// GetMeta returns the Meta field value
func (o *FeesPaginated) GetMeta() PaginationMeta {
	if o == nil {
		var ret PaginationMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *FeesPaginated) GetMetaOk() (*PaginationMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *FeesPaginated) SetMeta(v PaginationMeta) {
	o.Meta = v
}

func (o FeesPaginated) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FeesPaginated) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fees"] = o.Fees
	toSerialize["meta"] = o.Meta
	return toSerialize, nil
}

func (o *FeesPaginated) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fees",
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

	varFeesPaginated := _FeesPaginated{}

	err = json.Unmarshal(bytes, &varFeesPaginated)

	if err != nil {
		return err
	}

	*o = FeesPaginated(varFeesPaginated)

	return err
}

type NullableFeesPaginated struct {
	value *FeesPaginated
	isSet bool
}

func (v NullableFeesPaginated) Get() *FeesPaginated {
	return v.value
}

func (v *NullableFeesPaginated) Set(val *FeesPaginated) {
	v.value = val
	v.isSet = true
}

func (v NullableFeesPaginated) IsSet() bool {
	return v.isSet
}

func (v *NullableFeesPaginated) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeesPaginated(val *FeesPaginated) *NullableFeesPaginated {
	return &NullableFeesPaginated{value: val, isSet: true}
}

func (v NullableFeesPaginated) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeesPaginated) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


