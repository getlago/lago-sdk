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

// checks if the WalletTransactionCreateInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WalletTransactionCreateInput{}

// WalletTransactionCreateInput struct for WalletTransactionCreateInput
type WalletTransactionCreateInput struct {
	WalletTransaction WalletTransactionCreateInputWalletTransaction `json:"wallet_transaction"`
}

type _WalletTransactionCreateInput WalletTransactionCreateInput

// NewWalletTransactionCreateInput instantiates a new WalletTransactionCreateInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWalletTransactionCreateInput(walletTransaction WalletTransactionCreateInputWalletTransaction) *WalletTransactionCreateInput {
	this := WalletTransactionCreateInput{}
	this.WalletTransaction = walletTransaction
	return &this
}

// NewWalletTransactionCreateInputWithDefaults instantiates a new WalletTransactionCreateInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWalletTransactionCreateInputWithDefaults() *WalletTransactionCreateInput {
	this := WalletTransactionCreateInput{}
	return &this
}

// GetWalletTransaction returns the WalletTransaction field value
func (o *WalletTransactionCreateInput) GetWalletTransaction() WalletTransactionCreateInputWalletTransaction {
	if o == nil {
		var ret WalletTransactionCreateInputWalletTransaction
		return ret
	}

	return o.WalletTransaction
}

// GetWalletTransactionOk returns a tuple with the WalletTransaction field value
// and a boolean to check if the value has been set.
func (o *WalletTransactionCreateInput) GetWalletTransactionOk() (*WalletTransactionCreateInputWalletTransaction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WalletTransaction, true
}

// SetWalletTransaction sets field value
func (o *WalletTransactionCreateInput) SetWalletTransaction(v WalletTransactionCreateInputWalletTransaction) {
	o.WalletTransaction = v
}

func (o WalletTransactionCreateInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WalletTransactionCreateInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["wallet_transaction"] = o.WalletTransaction
	return toSerialize, nil
}

func (o *WalletTransactionCreateInput) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"wallet_transaction",
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

	varWalletTransactionCreateInput := _WalletTransactionCreateInput{}

	err = json.Unmarshal(bytes, &varWalletTransactionCreateInput)

	if err != nil {
		return err
	}

	*o = WalletTransactionCreateInput(varWalletTransactionCreateInput)

	return err
}

type NullableWalletTransactionCreateInput struct {
	value *WalletTransactionCreateInput
	isSet bool
}

func (v NullableWalletTransactionCreateInput) Get() *WalletTransactionCreateInput {
	return v.value
}

func (v *NullableWalletTransactionCreateInput) Set(val *WalletTransactionCreateInput) {
	v.value = val
	v.isSet = true
}

func (v NullableWalletTransactionCreateInput) IsSet() bool {
	return v.isSet
}

func (v *NullableWalletTransactionCreateInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWalletTransactionCreateInput(val *WalletTransactionCreateInput) *NullableWalletTransactionCreateInput {
	return &NullableWalletTransactionCreateInput{value: val, isSet: true}
}

func (v NullableWalletTransactionCreateInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWalletTransactionCreateInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

