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

// checks if the CreditObjectInvoice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreditObjectInvoice{}

// CreditObjectInvoice struct for CreditObjectInvoice
type CreditObjectInvoice struct {
	LagoId string `json:"lago_id"`
	PaymentStatus string `json:"payment_status"`
}

type _CreditObjectInvoice CreditObjectInvoice

// NewCreditObjectInvoice instantiates a new CreditObjectInvoice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditObjectInvoice(lagoId string, paymentStatus string) *CreditObjectInvoice {
	this := CreditObjectInvoice{}
	this.LagoId = lagoId
	this.PaymentStatus = paymentStatus
	return &this
}

// NewCreditObjectInvoiceWithDefaults instantiates a new CreditObjectInvoice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditObjectInvoiceWithDefaults() *CreditObjectInvoice {
	this := CreditObjectInvoice{}
	return &this
}

// GetLagoId returns the LagoId field value
func (o *CreditObjectInvoice) GetLagoId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LagoId
}

// GetLagoIdOk returns a tuple with the LagoId field value
// and a boolean to check if the value has been set.
func (o *CreditObjectInvoice) GetLagoIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LagoId, true
}

// SetLagoId sets field value
func (o *CreditObjectInvoice) SetLagoId(v string) {
	o.LagoId = v
}

// GetPaymentStatus returns the PaymentStatus field value
func (o *CreditObjectInvoice) GetPaymentStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentStatus
}

// GetPaymentStatusOk returns a tuple with the PaymentStatus field value
// and a boolean to check if the value has been set.
func (o *CreditObjectInvoice) GetPaymentStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentStatus, true
}

// SetPaymentStatus sets field value
func (o *CreditObjectInvoice) SetPaymentStatus(v string) {
	o.PaymentStatus = v
}

func (o CreditObjectInvoice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreditObjectInvoice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["lago_id"] = o.LagoId
	toSerialize["payment_status"] = o.PaymentStatus
	return toSerialize, nil
}

func (o *CreditObjectInvoice) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"lago_id",
		"payment_status",
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

	varCreditObjectInvoice := _CreditObjectInvoice{}

	err = json.Unmarshal(bytes, &varCreditObjectInvoice)

	if err != nil {
		return err
	}

	*o = CreditObjectInvoice(varCreditObjectInvoice)

	return err
}

type NullableCreditObjectInvoice struct {
	value *CreditObjectInvoice
	isSet bool
}

func (v NullableCreditObjectInvoice) Get() *CreditObjectInvoice {
	return v.value
}

func (v *NullableCreditObjectInvoice) Set(val *CreditObjectInvoice) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditObjectInvoice) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditObjectInvoice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditObjectInvoice(val *CreditObjectInvoice) *NullableCreditObjectInvoice {
	return &NullableCreditObjectInvoice{value: val, isSet: true}
}

func (v NullableCreditObjectInvoice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditObjectInvoice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

