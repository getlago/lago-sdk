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

// checks if the AppliedCoupon type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppliedCoupon{}

// AppliedCoupon struct for AppliedCoupon
type AppliedCoupon struct {
	AppliedCoupon AppliedCouponObject `json:"applied_coupon"`
}

type _AppliedCoupon AppliedCoupon

// NewAppliedCoupon instantiates a new AppliedCoupon object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppliedCoupon(appliedCoupon AppliedCouponObject) *AppliedCoupon {
	this := AppliedCoupon{}
	this.AppliedCoupon = appliedCoupon
	return &this
}

// NewAppliedCouponWithDefaults instantiates a new AppliedCoupon object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppliedCouponWithDefaults() *AppliedCoupon {
	this := AppliedCoupon{}
	return &this
}

// GetAppliedCoupon returns the AppliedCoupon field value
func (o *AppliedCoupon) GetAppliedCoupon() AppliedCouponObject {
	if o == nil {
		var ret AppliedCouponObject
		return ret
	}

	return o.AppliedCoupon
}

// GetAppliedCouponOk returns a tuple with the AppliedCoupon field value
// and a boolean to check if the value has been set.
func (o *AppliedCoupon) GetAppliedCouponOk() (*AppliedCouponObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppliedCoupon, true
}

// SetAppliedCoupon sets field value
func (o *AppliedCoupon) SetAppliedCoupon(v AppliedCouponObject) {
	o.AppliedCoupon = v
}

func (o AppliedCoupon) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppliedCoupon) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["applied_coupon"] = o.AppliedCoupon
	return toSerialize, nil
}

func (o *AppliedCoupon) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"applied_coupon",
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

	varAppliedCoupon := _AppliedCoupon{}

	err = json.Unmarshal(bytes, &varAppliedCoupon)

	if err != nil {
		return err
	}

	*o = AppliedCoupon(varAppliedCoupon)

	return err
}

type NullableAppliedCoupon struct {
	value *AppliedCoupon
	isSet bool
}

func (v NullableAppliedCoupon) Get() *AppliedCoupon {
	return v.value
}

func (v *NullableAppliedCoupon) Set(val *AppliedCoupon) {
	v.value = val
	v.isSet = true
}

func (v NullableAppliedCoupon) IsSet() bool {
	return v.isSet
}

func (v *NullableAppliedCoupon) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppliedCoupon(val *AppliedCoupon) *NullableAppliedCoupon {
	return &NullableAppliedCoupon{value: val, isSet: true}
}

func (v NullableAppliedCoupon) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppliedCoupon) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

