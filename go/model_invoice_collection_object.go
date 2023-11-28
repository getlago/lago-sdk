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

// checks if the InvoiceCollectionObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvoiceCollectionObject{}

// InvoiceCollectionObject struct for InvoiceCollectionObject
type InvoiceCollectionObject struct {
	// Identifies the month to analyze revenue.
	Month string `json:"month"`
	// The payment status of the invoices.
	PaymentStatus *string `json:"payment_status,omitempty"`
	// Contains invoices count.
	InvoicesCount int32 `json:"invoices_count"`
	// The total amount of revenue for a period, expressed in cents.
	AmountCents int32 `json:"amount_cents"`
	Currency *Currency `json:"currency,omitempty"`
}

type _InvoiceCollectionObject InvoiceCollectionObject

// NewInvoiceCollectionObject instantiates a new InvoiceCollectionObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceCollectionObject(month string, invoicesCount int32, amountCents int32) *InvoiceCollectionObject {
	this := InvoiceCollectionObject{}
	this.Month = month
	this.InvoicesCount = invoicesCount
	this.AmountCents = amountCents
	return &this
}

// NewInvoiceCollectionObjectWithDefaults instantiates a new InvoiceCollectionObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceCollectionObjectWithDefaults() *InvoiceCollectionObject {
	this := InvoiceCollectionObject{}
	return &this
}

// GetMonth returns the Month field value
func (o *InvoiceCollectionObject) GetMonth() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Month
}

// GetMonthOk returns a tuple with the Month field value
// and a boolean to check if the value has been set.
func (o *InvoiceCollectionObject) GetMonthOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Month, true
}

// SetMonth sets field value
func (o *InvoiceCollectionObject) SetMonth(v string) {
	o.Month = v
}

// GetPaymentStatus returns the PaymentStatus field value if set, zero value otherwise.
func (o *InvoiceCollectionObject) GetPaymentStatus() string {
	if o == nil || IsNil(o.PaymentStatus) {
		var ret string
		return ret
	}
	return *o.PaymentStatus
}

// GetPaymentStatusOk returns a tuple with the PaymentStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceCollectionObject) GetPaymentStatusOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentStatus) {
		return nil, false
	}
	return o.PaymentStatus, true
}

// HasPaymentStatus returns a boolean if a field has been set.
func (o *InvoiceCollectionObject) HasPaymentStatus() bool {
	if o != nil && !IsNil(o.PaymentStatus) {
		return true
	}

	return false
}

// SetPaymentStatus gets a reference to the given string and assigns it to the PaymentStatus field.
func (o *InvoiceCollectionObject) SetPaymentStatus(v string) {
	o.PaymentStatus = &v
}

// GetInvoicesCount returns the InvoicesCount field value
func (o *InvoiceCollectionObject) GetInvoicesCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.InvoicesCount
}

// GetInvoicesCountOk returns a tuple with the InvoicesCount field value
// and a boolean to check if the value has been set.
func (o *InvoiceCollectionObject) GetInvoicesCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvoicesCount, true
}

// SetInvoicesCount sets field value
func (o *InvoiceCollectionObject) SetInvoicesCount(v int32) {
	o.InvoicesCount = v
}

// GetAmountCents returns the AmountCents field value
func (o *InvoiceCollectionObject) GetAmountCents() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AmountCents
}

// GetAmountCentsOk returns a tuple with the AmountCents field value
// and a boolean to check if the value has been set.
func (o *InvoiceCollectionObject) GetAmountCentsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountCents, true
}

// SetAmountCents sets field value
func (o *InvoiceCollectionObject) SetAmountCents(v int32) {
	o.AmountCents = v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *InvoiceCollectionObject) GetCurrency() Currency {
	if o == nil || IsNil(o.Currency) {
		var ret Currency
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceCollectionObject) GetCurrencyOk() (*Currency, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *InvoiceCollectionObject) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given Currency and assigns it to the Currency field.
func (o *InvoiceCollectionObject) SetCurrency(v Currency) {
	o.Currency = &v
}

func (o InvoiceCollectionObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvoiceCollectionObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["month"] = o.Month
	if !IsNil(o.PaymentStatus) {
		toSerialize["payment_status"] = o.PaymentStatus
	}
	toSerialize["invoices_count"] = o.InvoicesCount
	toSerialize["amount_cents"] = o.AmountCents
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	return toSerialize, nil
}

func (o *InvoiceCollectionObject) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"month",
		"invoices_count",
		"amount_cents",
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

	varInvoiceCollectionObject := _InvoiceCollectionObject{}

	err = json.Unmarshal(bytes, &varInvoiceCollectionObject)

	if err != nil {
		return err
	}

	*o = InvoiceCollectionObject(varInvoiceCollectionObject)

	return err
}

type NullableInvoiceCollectionObject struct {
	value *InvoiceCollectionObject
	isSet bool
}

func (v NullableInvoiceCollectionObject) Get() *InvoiceCollectionObject {
	return v.value
}

func (v *NullableInvoiceCollectionObject) Set(val *InvoiceCollectionObject) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceCollectionObject) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceCollectionObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceCollectionObject(val *InvoiceCollectionObject) *NullableInvoiceCollectionObject {
	return &NullableInvoiceCollectionObject{value: val, isSet: true}
}

func (v NullableInvoiceCollectionObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceCollectionObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


