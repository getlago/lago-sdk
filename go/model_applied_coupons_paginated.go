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

// checks if the AppliedCouponsPaginated type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppliedCouponsPaginated{}

// AppliedCouponsPaginated struct for AppliedCouponsPaginated
type AppliedCouponsPaginated struct {
	AppliedCoupons []AppliedCouponObjectExtended `json:"applied_coupons"`
	Meta PaginationMeta `json:"meta"`
}

type _AppliedCouponsPaginated AppliedCouponsPaginated

// NewAppliedCouponsPaginated instantiates a new AppliedCouponsPaginated object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppliedCouponsPaginated(appliedCoupons []AppliedCouponObjectExtended, meta PaginationMeta) *AppliedCouponsPaginated {
	this := AppliedCouponsPaginated{}
	this.AppliedCoupons = appliedCoupons
	this.Meta = meta
	return &this
}

// NewAppliedCouponsPaginatedWithDefaults instantiates a new AppliedCouponsPaginated object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppliedCouponsPaginatedWithDefaults() *AppliedCouponsPaginated {
	this := AppliedCouponsPaginated{}
	return &this
}

// GetAppliedCoupons returns the AppliedCoupons field value
func (o *AppliedCouponsPaginated) GetAppliedCoupons() []AppliedCouponObjectExtended {
	if o == nil {
		var ret []AppliedCouponObjectExtended
		return ret
	}

	return o.AppliedCoupons
}

// GetAppliedCouponsOk returns a tuple with the AppliedCoupons field value
// and a boolean to check if the value has been set.
func (o *AppliedCouponsPaginated) GetAppliedCouponsOk() ([]AppliedCouponObjectExtended, bool) {
	if o == nil {
		return nil, false
	}
	return o.AppliedCoupons, true
}

// SetAppliedCoupons sets field value
func (o *AppliedCouponsPaginated) SetAppliedCoupons(v []AppliedCouponObjectExtended) {
	o.AppliedCoupons = v
}

// GetMeta returns the Meta field value
func (o *AppliedCouponsPaginated) GetMeta() PaginationMeta {
	if o == nil {
		var ret PaginationMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *AppliedCouponsPaginated) GetMetaOk() (*PaginationMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *AppliedCouponsPaginated) SetMeta(v PaginationMeta) {
	o.Meta = v
}

func (o AppliedCouponsPaginated) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppliedCouponsPaginated) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["applied_coupons"] = o.AppliedCoupons
	toSerialize["meta"] = o.Meta
	return toSerialize, nil
}

func (o *AppliedCouponsPaginated) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"applied_coupons",
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

	varAppliedCouponsPaginated := _AppliedCouponsPaginated{}

	err = json.Unmarshal(bytes, &varAppliedCouponsPaginated)

	if err != nil {
		return err
	}

	*o = AppliedCouponsPaginated(varAppliedCouponsPaginated)

	return err
}

type NullableAppliedCouponsPaginated struct {
	value *AppliedCouponsPaginated
	isSet bool
}

func (v NullableAppliedCouponsPaginated) Get() *AppliedCouponsPaginated {
	return v.value
}

func (v *NullableAppliedCouponsPaginated) Set(val *AppliedCouponsPaginated) {
	v.value = val
	v.isSet = true
}

func (v NullableAppliedCouponsPaginated) IsSet() bool {
	return v.isSet
}

func (v *NullableAppliedCouponsPaginated) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppliedCouponsPaginated(val *AppliedCouponsPaginated) *NullableAppliedCouponsPaginated {
	return &NullableAppliedCouponsPaginated{value: val, isSet: true}
}

func (v NullableAppliedCouponsPaginated) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppliedCouponsPaginated) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

