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
	"time"
	"fmt"
)

// checks if the AddOnObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddOnObject{}

// AddOnObject struct for AddOnObject
type AddOnObject struct {
	// Unique identifier of the add-on, created by Lago.
	LagoId string `json:"lago_id"`
	// The name of the add-on.
	Name string `json:"name"`
	// Specifies the name that will be displayed on an invoice. If no value is set for this field, the name of the actual charge will be used as the default display name.
	InvoiceDisplayName *string `json:"invoice_display_name,omitempty"`
	// Unique code used to identify the add-on.
	Code string `json:"code"`
	// The cost of the add-on in cents, excluding any applicable taxes, that is billed to a customer. By creating a one-off invoice, you will be able to override this value.
	AmountCents int32 `json:"amount_cents"`
	AmountCurrency Currency `json:"amount_currency"`
	// The description of the add-on.
	Description NullableString `json:"description,omitempty"`
	// The date and time when the add-on was created. It is expressed in UTC format according to the ISO 8601 datetime standard. This field provides the timestamp for the exact moment when the add-on was initially created.
	CreatedAt time.Time `json:"created_at"`
	// All taxes applied to the add-on.
	Taxes []TaxObject `json:"taxes,omitempty"`
}

type _AddOnObject AddOnObject

// NewAddOnObject instantiates a new AddOnObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddOnObject(lagoId string, name string, code string, amountCents int32, amountCurrency Currency, createdAt time.Time) *AddOnObject {
	this := AddOnObject{}
	this.LagoId = lagoId
	this.Name = name
	this.Code = code
	this.AmountCents = amountCents
	this.AmountCurrency = amountCurrency
	this.CreatedAt = createdAt
	return &this
}

// NewAddOnObjectWithDefaults instantiates a new AddOnObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddOnObjectWithDefaults() *AddOnObject {
	this := AddOnObject{}
	return &this
}

// GetLagoId returns the LagoId field value
func (o *AddOnObject) GetLagoId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LagoId
}

// GetLagoIdOk returns a tuple with the LagoId field value
// and a boolean to check if the value has been set.
func (o *AddOnObject) GetLagoIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LagoId, true
}

// SetLagoId sets field value
func (o *AddOnObject) SetLagoId(v string) {
	o.LagoId = v
}

// GetName returns the Name field value
func (o *AddOnObject) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AddOnObject) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AddOnObject) SetName(v string) {
	o.Name = v
}

// GetInvoiceDisplayName returns the InvoiceDisplayName field value if set, zero value otherwise.
func (o *AddOnObject) GetInvoiceDisplayName() string {
	if o == nil || IsNil(o.InvoiceDisplayName) {
		var ret string
		return ret
	}
	return *o.InvoiceDisplayName
}

// GetInvoiceDisplayNameOk returns a tuple with the InvoiceDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddOnObject) GetInvoiceDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceDisplayName) {
		return nil, false
	}
	return o.InvoiceDisplayName, true
}

// HasInvoiceDisplayName returns a boolean if a field has been set.
func (o *AddOnObject) HasInvoiceDisplayName() bool {
	if o != nil && !IsNil(o.InvoiceDisplayName) {
		return true
	}

	return false
}

// SetInvoiceDisplayName gets a reference to the given string and assigns it to the InvoiceDisplayName field.
func (o *AddOnObject) SetInvoiceDisplayName(v string) {
	o.InvoiceDisplayName = &v
}

// GetCode returns the Code field value
func (o *AddOnObject) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *AddOnObject) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *AddOnObject) SetCode(v string) {
	o.Code = v
}

// GetAmountCents returns the AmountCents field value
func (o *AddOnObject) GetAmountCents() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AmountCents
}

// GetAmountCentsOk returns a tuple with the AmountCents field value
// and a boolean to check if the value has been set.
func (o *AddOnObject) GetAmountCentsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountCents, true
}

// SetAmountCents sets field value
func (o *AddOnObject) SetAmountCents(v int32) {
	o.AmountCents = v
}

// GetAmountCurrency returns the AmountCurrency field value
func (o *AddOnObject) GetAmountCurrency() Currency {
	if o == nil {
		var ret Currency
		return ret
	}

	return o.AmountCurrency
}

// GetAmountCurrencyOk returns a tuple with the AmountCurrency field value
// and a boolean to check if the value has been set.
func (o *AddOnObject) GetAmountCurrencyOk() (*Currency, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountCurrency, true
}

// SetAmountCurrency sets field value
func (o *AddOnObject) SetAmountCurrency(v Currency) {
	o.AmountCurrency = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AddOnObject) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddOnObject) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *AddOnObject) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *AddOnObject) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *AddOnObject) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *AddOnObject) UnsetDescription() {
	o.Description.Unset()
}

// GetCreatedAt returns the CreatedAt field value
func (o *AddOnObject) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AddOnObject) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *AddOnObject) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetTaxes returns the Taxes field value if set, zero value otherwise.
func (o *AddOnObject) GetTaxes() []TaxObject {
	if o == nil || IsNil(o.Taxes) {
		var ret []TaxObject
		return ret
	}
	return o.Taxes
}

// GetTaxesOk returns a tuple with the Taxes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddOnObject) GetTaxesOk() ([]TaxObject, bool) {
	if o == nil || IsNil(o.Taxes) {
		return nil, false
	}
	return o.Taxes, true
}

// HasTaxes returns a boolean if a field has been set.
func (o *AddOnObject) HasTaxes() bool {
	if o != nil && !IsNil(o.Taxes) {
		return true
	}

	return false
}

// SetTaxes gets a reference to the given []TaxObject and assigns it to the Taxes field.
func (o *AddOnObject) SetTaxes(v []TaxObject) {
	o.Taxes = v
}

func (o AddOnObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddOnObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["lago_id"] = o.LagoId
	toSerialize["name"] = o.Name
	if !IsNil(o.InvoiceDisplayName) {
		toSerialize["invoice_display_name"] = o.InvoiceDisplayName
	}
	toSerialize["code"] = o.Code
	toSerialize["amount_cents"] = o.AmountCents
	toSerialize["amount_currency"] = o.AmountCurrency
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	toSerialize["created_at"] = o.CreatedAt
	if !IsNil(o.Taxes) {
		toSerialize["taxes"] = o.Taxes
	}
	return toSerialize, nil
}

func (o *AddOnObject) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"lago_id",
		"name",
		"code",
		"amount_cents",
		"amount_currency",
		"created_at",
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

	varAddOnObject := _AddOnObject{}

	err = json.Unmarshal(bytes, &varAddOnObject)

	if err != nil {
		return err
	}

	*o = AddOnObject(varAddOnObject)

	return err
}

type NullableAddOnObject struct {
	value *AddOnObject
	isSet bool
}

func (v NullableAddOnObject) Get() *AddOnObject {
	return v.value
}

func (v *NullableAddOnObject) Set(val *AddOnObject) {
	v.value = val
	v.isSet = true
}

func (v NullableAddOnObject) IsSet() bool {
	return v.isSet
}

func (v *NullableAddOnObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddOnObject(val *AddOnObject) *NullableAddOnObject {
	return &NullableAddOnObject{value: val, isSet: true}
}

func (v NullableAddOnObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddOnObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


