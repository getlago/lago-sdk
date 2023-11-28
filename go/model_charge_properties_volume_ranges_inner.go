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

// checks if the ChargePropertiesVolumeRangesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChargePropertiesVolumeRangesInner{}

// ChargePropertiesVolumeRangesInner struct for ChargePropertiesVolumeRangesInner
type ChargePropertiesVolumeRangesInner struct {
	// Specifies the lower value of a tier for a `volume` charge model. It must be either 0 or the previous range's `to_value + 1` to maintain the proper sequence of values.
	FromValue int32 `json:"from_value"`
	// Specifies the highest value of a tier for a `volume` charge model. - This value must be higher than the `from_value` of the same tier. - This value must be `null` for the last tier.
	ToValue NullableInt32 `json:"to_value"`
	// The unit price, excluding tax, for a specific tier of a `volume` charge model. It is expressed as a decimal value.
	FlatAmount string `json:"flat_amount"`
	// The flat amount for a whole tier, excluding tax, for a `volume` charge model. It is expressed as a decimal value.
	PerUnitAmount string `json:"per_unit_amount"`
}

type _ChargePropertiesVolumeRangesInner ChargePropertiesVolumeRangesInner

// NewChargePropertiesVolumeRangesInner instantiates a new ChargePropertiesVolumeRangesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChargePropertiesVolumeRangesInner(fromValue int32, toValue NullableInt32, flatAmount string, perUnitAmount string) *ChargePropertiesVolumeRangesInner {
	this := ChargePropertiesVolumeRangesInner{}
	this.FromValue = fromValue
	this.ToValue = toValue
	this.FlatAmount = flatAmount
	this.PerUnitAmount = perUnitAmount
	return &this
}

// NewChargePropertiesVolumeRangesInnerWithDefaults instantiates a new ChargePropertiesVolumeRangesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChargePropertiesVolumeRangesInnerWithDefaults() *ChargePropertiesVolumeRangesInner {
	this := ChargePropertiesVolumeRangesInner{}
	return &this
}

// GetFromValue returns the FromValue field value
func (o *ChargePropertiesVolumeRangesInner) GetFromValue() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FromValue
}

// GetFromValueOk returns a tuple with the FromValue field value
// and a boolean to check if the value has been set.
func (o *ChargePropertiesVolumeRangesInner) GetFromValueOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromValue, true
}

// SetFromValue sets field value
func (o *ChargePropertiesVolumeRangesInner) SetFromValue(v int32) {
	o.FromValue = v
}

// GetToValue returns the ToValue field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *ChargePropertiesVolumeRangesInner) GetToValue() int32 {
	if o == nil || o.ToValue.Get() == nil {
		var ret int32
		return ret
	}

	return *o.ToValue.Get()
}

// GetToValueOk returns a tuple with the ToValue field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChargePropertiesVolumeRangesInner) GetToValueOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ToValue.Get(), o.ToValue.IsSet()
}

// SetToValue sets field value
func (o *ChargePropertiesVolumeRangesInner) SetToValue(v int32) {
	o.ToValue.Set(&v)
}

// GetFlatAmount returns the FlatAmount field value
func (o *ChargePropertiesVolumeRangesInner) GetFlatAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FlatAmount
}

// GetFlatAmountOk returns a tuple with the FlatAmount field value
// and a boolean to check if the value has been set.
func (o *ChargePropertiesVolumeRangesInner) GetFlatAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FlatAmount, true
}

// SetFlatAmount sets field value
func (o *ChargePropertiesVolumeRangesInner) SetFlatAmount(v string) {
	o.FlatAmount = v
}

// GetPerUnitAmount returns the PerUnitAmount field value
func (o *ChargePropertiesVolumeRangesInner) GetPerUnitAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PerUnitAmount
}

// GetPerUnitAmountOk returns a tuple with the PerUnitAmount field value
// and a boolean to check if the value has been set.
func (o *ChargePropertiesVolumeRangesInner) GetPerUnitAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PerUnitAmount, true
}

// SetPerUnitAmount sets field value
func (o *ChargePropertiesVolumeRangesInner) SetPerUnitAmount(v string) {
	o.PerUnitAmount = v
}

func (o ChargePropertiesVolumeRangesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChargePropertiesVolumeRangesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["from_value"] = o.FromValue
	toSerialize["to_value"] = o.ToValue.Get()
	toSerialize["flat_amount"] = o.FlatAmount
	toSerialize["per_unit_amount"] = o.PerUnitAmount
	return toSerialize, nil
}

func (o *ChargePropertiesVolumeRangesInner) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"from_value",
		"to_value",
		"flat_amount",
		"per_unit_amount",
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

	varChargePropertiesVolumeRangesInner := _ChargePropertiesVolumeRangesInner{}

	err = json.Unmarshal(bytes, &varChargePropertiesVolumeRangesInner)

	if err != nil {
		return err
	}

	*o = ChargePropertiesVolumeRangesInner(varChargePropertiesVolumeRangesInner)

	return err
}

type NullableChargePropertiesVolumeRangesInner struct {
	value *ChargePropertiesVolumeRangesInner
	isSet bool
}

func (v NullableChargePropertiesVolumeRangesInner) Get() *ChargePropertiesVolumeRangesInner {
	return v.value
}

func (v *NullableChargePropertiesVolumeRangesInner) Set(val *ChargePropertiesVolumeRangesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableChargePropertiesVolumeRangesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableChargePropertiesVolumeRangesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChargePropertiesVolumeRangesInner(val *ChargePropertiesVolumeRangesInner) *NullableChargePropertiesVolumeRangesInner {
	return &NullableChargePropertiesVolumeRangesInner{value: val, isSet: true}
}

func (v NullableChargePropertiesVolumeRangesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChargePropertiesVolumeRangesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

