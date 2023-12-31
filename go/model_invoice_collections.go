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

// checks if the InvoiceCollections type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvoiceCollections{}

// InvoiceCollections struct for InvoiceCollections
type InvoiceCollections struct {
	InvoiceCollections []InvoiceCollectionObject `json:"invoice_collections"`
}

type _InvoiceCollections InvoiceCollections

// NewInvoiceCollections instantiates a new InvoiceCollections object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceCollections(invoiceCollections []InvoiceCollectionObject) *InvoiceCollections {
	this := InvoiceCollections{}
	this.InvoiceCollections = invoiceCollections
	return &this
}

// NewInvoiceCollectionsWithDefaults instantiates a new InvoiceCollections object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceCollectionsWithDefaults() *InvoiceCollections {
	this := InvoiceCollections{}
	return &this
}

// GetInvoiceCollections returns the InvoiceCollections field value
func (o *InvoiceCollections) GetInvoiceCollections() []InvoiceCollectionObject {
	if o == nil {
		var ret []InvoiceCollectionObject
		return ret
	}

	return o.InvoiceCollections
}

// GetInvoiceCollectionsOk returns a tuple with the InvoiceCollections field value
// and a boolean to check if the value has been set.
func (o *InvoiceCollections) GetInvoiceCollectionsOk() ([]InvoiceCollectionObject, bool) {
	if o == nil {
		return nil, false
	}
	return o.InvoiceCollections, true
}

// SetInvoiceCollections sets field value
func (o *InvoiceCollections) SetInvoiceCollections(v []InvoiceCollectionObject) {
	o.InvoiceCollections = v
}

func (o InvoiceCollections) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvoiceCollections) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["invoice_collections"] = o.InvoiceCollections
	return toSerialize, nil
}

func (o *InvoiceCollections) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"invoice_collections",
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

	varInvoiceCollections := _InvoiceCollections{}

	err = json.Unmarshal(bytes, &varInvoiceCollections)

	if err != nil {
		return err
	}

	*o = InvoiceCollections(varInvoiceCollections)

	return err
}

type NullableInvoiceCollections struct {
	value *InvoiceCollections
	isSet bool
}

func (v NullableInvoiceCollections) Get() *InvoiceCollections {
	return v.value
}

func (v *NullableInvoiceCollections) Set(val *InvoiceCollections) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceCollections) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceCollections) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceCollections(val *InvoiceCollections) *NullableInvoiceCollections {
	return &NullableInvoiceCollections{value: val, isSet: true}
}

func (v NullableInvoiceCollections) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceCollections) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


