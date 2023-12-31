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

// checks if the WalletsPaginated type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WalletsPaginated{}

// WalletsPaginated struct for WalletsPaginated
type WalletsPaginated struct {
	Wallets []WalletObject `json:"wallets"`
	Meta PaginationMeta `json:"meta"`
}

type _WalletsPaginated WalletsPaginated

// NewWalletsPaginated instantiates a new WalletsPaginated object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWalletsPaginated(wallets []WalletObject, meta PaginationMeta) *WalletsPaginated {
	this := WalletsPaginated{}
	this.Wallets = wallets
	this.Meta = meta
	return &this
}

// NewWalletsPaginatedWithDefaults instantiates a new WalletsPaginated object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWalletsPaginatedWithDefaults() *WalletsPaginated {
	this := WalletsPaginated{}
	return &this
}

// GetWallets returns the Wallets field value
func (o *WalletsPaginated) GetWallets() []WalletObject {
	if o == nil {
		var ret []WalletObject
		return ret
	}

	return o.Wallets
}

// GetWalletsOk returns a tuple with the Wallets field value
// and a boolean to check if the value has been set.
func (o *WalletsPaginated) GetWalletsOk() ([]WalletObject, bool) {
	if o == nil {
		return nil, false
	}
	return o.Wallets, true
}

// SetWallets sets field value
func (o *WalletsPaginated) SetWallets(v []WalletObject) {
	o.Wallets = v
}

// GetMeta returns the Meta field value
func (o *WalletsPaginated) GetMeta() PaginationMeta {
	if o == nil {
		var ret PaginationMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *WalletsPaginated) GetMetaOk() (*PaginationMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *WalletsPaginated) SetMeta(v PaginationMeta) {
	o.Meta = v
}

func (o WalletsPaginated) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WalletsPaginated) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["wallets"] = o.Wallets
	toSerialize["meta"] = o.Meta
	return toSerialize, nil
}

func (o *WalletsPaginated) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"wallets",
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

	varWalletsPaginated := _WalletsPaginated{}

	err = json.Unmarshal(bytes, &varWalletsPaginated)

	if err != nil {
		return err
	}

	*o = WalletsPaginated(varWalletsPaginated)

	return err
}

type NullableWalletsPaginated struct {
	value *WalletsPaginated
	isSet bool
}

func (v NullableWalletsPaginated) Get() *WalletsPaginated {
	return v.value
}

func (v *NullableWalletsPaginated) Set(val *WalletsPaginated) {
	v.value = val
	v.isSet = true
}

func (v NullableWalletsPaginated) IsSet() bool {
	return v.isSet
}

func (v *NullableWalletsPaginated) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWalletsPaginated(val *WalletsPaginated) *NullableWalletsPaginated {
	return &NullableWalletsPaginated{value: val, isSet: true}
}

func (v NullableWalletsPaginated) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWalletsPaginated) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


