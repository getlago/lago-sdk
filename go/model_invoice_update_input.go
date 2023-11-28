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

// checks if the InvoiceUpdateInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvoiceUpdateInput{}

// InvoiceUpdateInput struct for InvoiceUpdateInput
type InvoiceUpdateInput struct {
	Invoice InvoiceUpdateInputInvoice `json:"invoice"`
}

type _InvoiceUpdateInput InvoiceUpdateInput

// NewInvoiceUpdateInput instantiates a new InvoiceUpdateInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceUpdateInput(invoice InvoiceUpdateInputInvoice) *InvoiceUpdateInput {
	this := InvoiceUpdateInput{}
	this.Invoice = invoice
	return &this
}

// NewInvoiceUpdateInputWithDefaults instantiates a new InvoiceUpdateInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceUpdateInputWithDefaults() *InvoiceUpdateInput {
	this := InvoiceUpdateInput{}
	return &this
}

// GetInvoice returns the Invoice field value
func (o *InvoiceUpdateInput) GetInvoice() InvoiceUpdateInputInvoice {
	if o == nil {
		var ret InvoiceUpdateInputInvoice
		return ret
	}

	return o.Invoice
}

// GetInvoiceOk returns a tuple with the Invoice field value
// and a boolean to check if the value has been set.
func (o *InvoiceUpdateInput) GetInvoiceOk() (*InvoiceUpdateInputInvoice, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Invoice, true
}

// SetInvoice sets field value
func (o *InvoiceUpdateInput) SetInvoice(v InvoiceUpdateInputInvoice) {
	o.Invoice = v
}

func (o InvoiceUpdateInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvoiceUpdateInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["invoice"] = o.Invoice
	return toSerialize, nil
}

func (o *InvoiceUpdateInput) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"invoice",
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

	varInvoiceUpdateInput := _InvoiceUpdateInput{}

	err = json.Unmarshal(bytes, &varInvoiceUpdateInput)

	if err != nil {
		return err
	}

	*o = InvoiceUpdateInput(varInvoiceUpdateInput)

	return err
}

type NullableInvoiceUpdateInput struct {
	value *InvoiceUpdateInput
	isSet bool
}

func (v NullableInvoiceUpdateInput) Get() *InvoiceUpdateInput {
	return v.value
}

func (v *NullableInvoiceUpdateInput) Set(val *InvoiceUpdateInput) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceUpdateInput) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceUpdateInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceUpdateInput(val *InvoiceUpdateInput) *NullableInvoiceUpdateInput {
	return &NullableInvoiceUpdateInput{value: val, isSet: true}
}

func (v NullableInvoiceUpdateInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceUpdateInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


