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

// checks if the TaxCreateInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaxCreateInput{}

// TaxCreateInput struct for TaxCreateInput
type TaxCreateInput struct {
	Tax TaxCreateInputTax `json:"tax"`
}

type _TaxCreateInput TaxCreateInput

// NewTaxCreateInput instantiates a new TaxCreateInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxCreateInput(tax TaxCreateInputTax) *TaxCreateInput {
	this := TaxCreateInput{}
	this.Tax = tax
	return &this
}

// NewTaxCreateInputWithDefaults instantiates a new TaxCreateInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxCreateInputWithDefaults() *TaxCreateInput {
	this := TaxCreateInput{}
	return &this
}

// GetTax returns the Tax field value
func (o *TaxCreateInput) GetTax() TaxCreateInputTax {
	if o == nil {
		var ret TaxCreateInputTax
		return ret
	}

	return o.Tax
}

// GetTaxOk returns a tuple with the Tax field value
// and a boolean to check if the value has been set.
func (o *TaxCreateInput) GetTaxOk() (*TaxCreateInputTax, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tax, true
}

// SetTax sets field value
func (o *TaxCreateInput) SetTax(v TaxCreateInputTax) {
	o.Tax = v
}

func (o TaxCreateInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaxCreateInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tax"] = o.Tax
	return toSerialize, nil
}

func (o *TaxCreateInput) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tax",
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

	varTaxCreateInput := _TaxCreateInput{}

	err = json.Unmarshal(bytes, &varTaxCreateInput)

	if err != nil {
		return err
	}

	*o = TaxCreateInput(varTaxCreateInput)

	return err
}

type NullableTaxCreateInput struct {
	value *TaxCreateInput
	isSet bool
}

func (v NullableTaxCreateInput) Get() *TaxCreateInput {
	return v.value
}

func (v *NullableTaxCreateInput) Set(val *TaxCreateInput) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxCreateInput) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxCreateInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxCreateInput(val *TaxCreateInput) *NullableTaxCreateInput {
	return &NullableTaxCreateInput{value: val, isSet: true}
}

func (v NullableTaxCreateInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxCreateInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


