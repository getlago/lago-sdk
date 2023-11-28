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

// checks if the WalletTransactionCreateInputWalletTransaction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WalletTransactionCreateInputWalletTransaction{}

// WalletTransactionCreateInputWalletTransaction struct for WalletTransactionCreateInputWalletTransaction
type WalletTransactionCreateInputWalletTransaction struct {
	// Unique identifier assigned to the wallet within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the wallet’s record within the Lago system.
	WalletId string `json:"wallet_id"`
	// The number of paid credits. Required only if there is no granted credits.
	PaidCredits *string `json:"paid_credits,omitempty"`
	// The number of free granted credits. Required only if there is no paid credits.
	GrantedCredits *string `json:"granted_credits,omitempty"`
}

type _WalletTransactionCreateInputWalletTransaction WalletTransactionCreateInputWalletTransaction

// NewWalletTransactionCreateInputWalletTransaction instantiates a new WalletTransactionCreateInputWalletTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWalletTransactionCreateInputWalletTransaction(walletId string) *WalletTransactionCreateInputWalletTransaction {
	this := WalletTransactionCreateInputWalletTransaction{}
	this.WalletId = walletId
	return &this
}

// NewWalletTransactionCreateInputWalletTransactionWithDefaults instantiates a new WalletTransactionCreateInputWalletTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWalletTransactionCreateInputWalletTransactionWithDefaults() *WalletTransactionCreateInputWalletTransaction {
	this := WalletTransactionCreateInputWalletTransaction{}
	return &this
}

// GetWalletId returns the WalletId field value
func (o *WalletTransactionCreateInputWalletTransaction) GetWalletId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WalletId
}

// GetWalletIdOk returns a tuple with the WalletId field value
// and a boolean to check if the value has been set.
func (o *WalletTransactionCreateInputWalletTransaction) GetWalletIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WalletId, true
}

// SetWalletId sets field value
func (o *WalletTransactionCreateInputWalletTransaction) SetWalletId(v string) {
	o.WalletId = v
}

// GetPaidCredits returns the PaidCredits field value if set, zero value otherwise.
func (o *WalletTransactionCreateInputWalletTransaction) GetPaidCredits() string {
	if o == nil || IsNil(o.PaidCredits) {
		var ret string
		return ret
	}
	return *o.PaidCredits
}

// GetPaidCreditsOk returns a tuple with the PaidCredits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletTransactionCreateInputWalletTransaction) GetPaidCreditsOk() (*string, bool) {
	if o == nil || IsNil(o.PaidCredits) {
		return nil, false
	}
	return o.PaidCredits, true
}

// HasPaidCredits returns a boolean if a field has been set.
func (o *WalletTransactionCreateInputWalletTransaction) HasPaidCredits() bool {
	if o != nil && !IsNil(o.PaidCredits) {
		return true
	}

	return false
}

// SetPaidCredits gets a reference to the given string and assigns it to the PaidCredits field.
func (o *WalletTransactionCreateInputWalletTransaction) SetPaidCredits(v string) {
	o.PaidCredits = &v
}

// GetGrantedCredits returns the GrantedCredits field value if set, zero value otherwise.
func (o *WalletTransactionCreateInputWalletTransaction) GetGrantedCredits() string {
	if o == nil || IsNil(o.GrantedCredits) {
		var ret string
		return ret
	}
	return *o.GrantedCredits
}

// GetGrantedCreditsOk returns a tuple with the GrantedCredits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletTransactionCreateInputWalletTransaction) GetGrantedCreditsOk() (*string, bool) {
	if o == nil || IsNil(o.GrantedCredits) {
		return nil, false
	}
	return o.GrantedCredits, true
}

// HasGrantedCredits returns a boolean if a field has been set.
func (o *WalletTransactionCreateInputWalletTransaction) HasGrantedCredits() bool {
	if o != nil && !IsNil(o.GrantedCredits) {
		return true
	}

	return false
}

// SetGrantedCredits gets a reference to the given string and assigns it to the GrantedCredits field.
func (o *WalletTransactionCreateInputWalletTransaction) SetGrantedCredits(v string) {
	o.GrantedCredits = &v
}

func (o WalletTransactionCreateInputWalletTransaction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WalletTransactionCreateInputWalletTransaction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["wallet_id"] = o.WalletId
	if !IsNil(o.PaidCredits) {
		toSerialize["paid_credits"] = o.PaidCredits
	}
	if !IsNil(o.GrantedCredits) {
		toSerialize["granted_credits"] = o.GrantedCredits
	}
	return toSerialize, nil
}

func (o *WalletTransactionCreateInputWalletTransaction) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"wallet_id",
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

	varWalletTransactionCreateInputWalletTransaction := _WalletTransactionCreateInputWalletTransaction{}

	err = json.Unmarshal(bytes, &varWalletTransactionCreateInputWalletTransaction)

	if err != nil {
		return err
	}

	*o = WalletTransactionCreateInputWalletTransaction(varWalletTransactionCreateInputWalletTransaction)

	return err
}

type NullableWalletTransactionCreateInputWalletTransaction struct {
	value *WalletTransactionCreateInputWalletTransaction
	isSet bool
}

func (v NullableWalletTransactionCreateInputWalletTransaction) Get() *WalletTransactionCreateInputWalletTransaction {
	return v.value
}

func (v *NullableWalletTransactionCreateInputWalletTransaction) Set(val *WalletTransactionCreateInputWalletTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableWalletTransactionCreateInputWalletTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableWalletTransactionCreateInputWalletTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWalletTransactionCreateInputWalletTransaction(val *WalletTransactionCreateInputWalletTransaction) *NullableWalletTransactionCreateInputWalletTransaction {
	return &NullableWalletTransactionCreateInputWalletTransaction{value: val, isSet: true}
}

func (v NullableWalletTransactionCreateInputWalletTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWalletTransactionCreateInputWalletTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


