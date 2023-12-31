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

// checks if the CouponUpdateInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CouponUpdateInput{}

// CouponUpdateInput struct for CouponUpdateInput
type CouponUpdateInput struct {
	Coupon CouponBaseInput `json:"coupon"`
}

type _CouponUpdateInput CouponUpdateInput

// NewCouponUpdateInput instantiates a new CouponUpdateInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCouponUpdateInput(coupon CouponBaseInput) *CouponUpdateInput {
	this := CouponUpdateInput{}
	this.Coupon = coupon
	return &this
}

// NewCouponUpdateInputWithDefaults instantiates a new CouponUpdateInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCouponUpdateInputWithDefaults() *CouponUpdateInput {
	this := CouponUpdateInput{}
	return &this
}

// GetCoupon returns the Coupon field value
func (o *CouponUpdateInput) GetCoupon() CouponBaseInput {
	if o == nil {
		var ret CouponBaseInput
		return ret
	}

	return o.Coupon
}

// GetCouponOk returns a tuple with the Coupon field value
// and a boolean to check if the value has been set.
func (o *CouponUpdateInput) GetCouponOk() (*CouponBaseInput, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Coupon, true
}

// SetCoupon sets field value
func (o *CouponUpdateInput) SetCoupon(v CouponBaseInput) {
	o.Coupon = v
}

func (o CouponUpdateInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CouponUpdateInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["coupon"] = o.Coupon
	return toSerialize, nil
}

func (o *CouponUpdateInput) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"coupon",
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

	varCouponUpdateInput := _CouponUpdateInput{}

	err = json.Unmarshal(bytes, &varCouponUpdateInput)

	if err != nil {
		return err
	}

	*o = CouponUpdateInput(varCouponUpdateInput)

	return err
}

type NullableCouponUpdateInput struct {
	value *CouponUpdateInput
	isSet bool
}

func (v NullableCouponUpdateInput) Get() *CouponUpdateInput {
	return v.value
}

func (v *NullableCouponUpdateInput) Set(val *CouponUpdateInput) {
	v.value = val
	v.isSet = true
}

func (v NullableCouponUpdateInput) IsSet() bool {
	return v.isSet
}

func (v *NullableCouponUpdateInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCouponUpdateInput(val *CouponUpdateInput) *NullableCouponUpdateInput {
	return &NullableCouponUpdateInput{value: val, isSet: true}
}

func (v NullableCouponUpdateInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCouponUpdateInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


