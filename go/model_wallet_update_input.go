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

// checks if the WalletUpdateInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WalletUpdateInput{}

// WalletUpdateInput struct for WalletUpdateInput
type WalletUpdateInput struct {
	Wallet WalletUpdateInputWallet `json:"wallet"`
}

type _WalletUpdateInput WalletUpdateInput

// NewWalletUpdateInput instantiates a new WalletUpdateInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWalletUpdateInput(wallet WalletUpdateInputWallet) *WalletUpdateInput {
	this := WalletUpdateInput{}
	this.Wallet = wallet
	return &this
}

// NewWalletUpdateInputWithDefaults instantiates a new WalletUpdateInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWalletUpdateInputWithDefaults() *WalletUpdateInput {
	this := WalletUpdateInput{}
	return &this
}

// GetWallet returns the Wallet field value
func (o *WalletUpdateInput) GetWallet() WalletUpdateInputWallet {
	if o == nil {
		var ret WalletUpdateInputWallet
		return ret
	}

	return o.Wallet
}

// GetWalletOk returns a tuple with the Wallet field value
// and a boolean to check if the value has been set.
func (o *WalletUpdateInput) GetWalletOk() (*WalletUpdateInputWallet, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Wallet, true
}

// SetWallet sets field value
func (o *WalletUpdateInput) SetWallet(v WalletUpdateInputWallet) {
	o.Wallet = v
}

func (o WalletUpdateInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WalletUpdateInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["wallet"] = o.Wallet
	return toSerialize, nil
}

func (o *WalletUpdateInput) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"wallet",
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

	varWalletUpdateInput := _WalletUpdateInput{}

	err = json.Unmarshal(bytes, &varWalletUpdateInput)

	if err != nil {
		return err
	}

	*o = WalletUpdateInput(varWalletUpdateInput)

	return err
}

type NullableWalletUpdateInput struct {
	value *WalletUpdateInput
	isSet bool
}

func (v NullableWalletUpdateInput) Get() *WalletUpdateInput {
	return v.value
}

func (v *NullableWalletUpdateInput) Set(val *WalletUpdateInput) {
	v.value = val
	v.isSet = true
}

func (v NullableWalletUpdateInput) IsSet() bool {
	return v.isSet
}

func (v *NullableWalletUpdateInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWalletUpdateInput(val *WalletUpdateInput) *NullableWalletUpdateInput {
	return &NullableWalletUpdateInput{value: val, isSet: true}
}

func (v NullableWalletUpdateInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWalletUpdateInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

