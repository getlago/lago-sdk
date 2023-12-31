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
)

// checks if the WalletCreateInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WalletCreateInput{}

// WalletCreateInput struct for WalletCreateInput
type WalletCreateInput struct {
	Wallet *WalletCreateInputWallet `json:"wallet,omitempty"`
}

// NewWalletCreateInput instantiates a new WalletCreateInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWalletCreateInput() *WalletCreateInput {
	this := WalletCreateInput{}
	return &this
}

// NewWalletCreateInputWithDefaults instantiates a new WalletCreateInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWalletCreateInputWithDefaults() *WalletCreateInput {
	this := WalletCreateInput{}
	return &this
}

// GetWallet returns the Wallet field value if set, zero value otherwise.
func (o *WalletCreateInput) GetWallet() WalletCreateInputWallet {
	if o == nil || IsNil(o.Wallet) {
		var ret WalletCreateInputWallet
		return ret
	}
	return *o.Wallet
}

// GetWalletOk returns a tuple with the Wallet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletCreateInput) GetWalletOk() (*WalletCreateInputWallet, bool) {
	if o == nil || IsNil(o.Wallet) {
		return nil, false
	}
	return o.Wallet, true
}

// HasWallet returns a boolean if a field has been set.
func (o *WalletCreateInput) HasWallet() bool {
	if o != nil && !IsNil(o.Wallet) {
		return true
	}

	return false
}

// SetWallet gets a reference to the given WalletCreateInputWallet and assigns it to the Wallet field.
func (o *WalletCreateInput) SetWallet(v WalletCreateInputWallet) {
	o.Wallet = &v
}

func (o WalletCreateInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WalletCreateInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Wallet) {
		toSerialize["wallet"] = o.Wallet
	}
	return toSerialize, nil
}

type NullableWalletCreateInput struct {
	value *WalletCreateInput
	isSet bool
}

func (v NullableWalletCreateInput) Get() *WalletCreateInput {
	return v.value
}

func (v *NullableWalletCreateInput) Set(val *WalletCreateInput) {
	v.value = val
	v.isSet = true
}

func (v NullableWalletCreateInput) IsSet() bool {
	return v.isSet
}

func (v *NullableWalletCreateInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWalletCreateInput(val *WalletCreateInput) *NullableWalletCreateInput {
	return &NullableWalletCreateInput{value: val, isSet: true}
}

func (v NullableWalletCreateInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWalletCreateInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


