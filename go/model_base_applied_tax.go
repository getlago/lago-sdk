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
)

// checks if the BaseAppliedTax type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseAppliedTax{}

// BaseAppliedTax struct for BaseAppliedTax
type BaseAppliedTax struct {
	// Unique identifier of the applied tax, created by Lago.
	LagoId *string `json:"lago_id,omitempty"`
	// Unique identifier of the tax, created by Lago.
	LagoTaxId *string `json:"lago_tax_id,omitempty"`
	// Name of the tax.
	TaxName *string `json:"tax_name,omitempty"`
	// Unique code used to identify the tax associated with the API request.
	TaxCode *string `json:"tax_code,omitempty"`
	// The percentage rate of the tax
	TaxRate *float32 `json:"tax_rate,omitempty"`
	// Internal description of the taxe
	TaxDescription *string `json:"tax_description,omitempty"`
	// Amount of the tax
	AmountCents *int32 `json:"amount_cents,omitempty"`
	AmountCurrency *Currency `json:"amount_currency,omitempty"`
	// The date and time when the applied tax was created. It is expressed in UTC format according to the ISO 8601 datetime standard. This field provides the timestamp for the exact moment when the applied tax was initially created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

// NewBaseAppliedTax instantiates a new BaseAppliedTax object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseAppliedTax() *BaseAppliedTax {
	this := BaseAppliedTax{}
	return &this
}

// NewBaseAppliedTaxWithDefaults instantiates a new BaseAppliedTax object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseAppliedTaxWithDefaults() *BaseAppliedTax {
	this := BaseAppliedTax{}
	return &this
}

// GetLagoId returns the LagoId field value if set, zero value otherwise.
func (o *BaseAppliedTax) GetLagoId() string {
	if o == nil || IsNil(o.LagoId) {
		var ret string
		return ret
	}
	return *o.LagoId
}

// GetLagoIdOk returns a tuple with the LagoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseAppliedTax) GetLagoIdOk() (*string, bool) {
	if o == nil || IsNil(o.LagoId) {
		return nil, false
	}
	return o.LagoId, true
}

// HasLagoId returns a boolean if a field has been set.
func (o *BaseAppliedTax) HasLagoId() bool {
	if o != nil && !IsNil(o.LagoId) {
		return true
	}

	return false
}

// SetLagoId gets a reference to the given string and assigns it to the LagoId field.
func (o *BaseAppliedTax) SetLagoId(v string) {
	o.LagoId = &v
}

// GetLagoTaxId returns the LagoTaxId field value if set, zero value otherwise.
func (o *BaseAppliedTax) GetLagoTaxId() string {
	if o == nil || IsNil(o.LagoTaxId) {
		var ret string
		return ret
	}
	return *o.LagoTaxId
}

// GetLagoTaxIdOk returns a tuple with the LagoTaxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseAppliedTax) GetLagoTaxIdOk() (*string, bool) {
	if o == nil || IsNil(o.LagoTaxId) {
		return nil, false
	}
	return o.LagoTaxId, true
}

// HasLagoTaxId returns a boolean if a field has been set.
func (o *BaseAppliedTax) HasLagoTaxId() bool {
	if o != nil && !IsNil(o.LagoTaxId) {
		return true
	}

	return false
}

// SetLagoTaxId gets a reference to the given string and assigns it to the LagoTaxId field.
func (o *BaseAppliedTax) SetLagoTaxId(v string) {
	o.LagoTaxId = &v
}

// GetTaxName returns the TaxName field value if set, zero value otherwise.
func (o *BaseAppliedTax) GetTaxName() string {
	if o == nil || IsNil(o.TaxName) {
		var ret string
		return ret
	}
	return *o.TaxName
}

// GetTaxNameOk returns a tuple with the TaxName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseAppliedTax) GetTaxNameOk() (*string, bool) {
	if o == nil || IsNil(o.TaxName) {
		return nil, false
	}
	return o.TaxName, true
}

// HasTaxName returns a boolean if a field has been set.
func (o *BaseAppliedTax) HasTaxName() bool {
	if o != nil && !IsNil(o.TaxName) {
		return true
	}

	return false
}

// SetTaxName gets a reference to the given string and assigns it to the TaxName field.
func (o *BaseAppliedTax) SetTaxName(v string) {
	o.TaxName = &v
}

// GetTaxCode returns the TaxCode field value if set, zero value otherwise.
func (o *BaseAppliedTax) GetTaxCode() string {
	if o == nil || IsNil(o.TaxCode) {
		var ret string
		return ret
	}
	return *o.TaxCode
}

// GetTaxCodeOk returns a tuple with the TaxCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseAppliedTax) GetTaxCodeOk() (*string, bool) {
	if o == nil || IsNil(o.TaxCode) {
		return nil, false
	}
	return o.TaxCode, true
}

// HasTaxCode returns a boolean if a field has been set.
func (o *BaseAppliedTax) HasTaxCode() bool {
	if o != nil && !IsNil(o.TaxCode) {
		return true
	}

	return false
}

// SetTaxCode gets a reference to the given string and assigns it to the TaxCode field.
func (o *BaseAppliedTax) SetTaxCode(v string) {
	o.TaxCode = &v
}

// GetTaxRate returns the TaxRate field value if set, zero value otherwise.
func (o *BaseAppliedTax) GetTaxRate() float32 {
	if o == nil || IsNil(o.TaxRate) {
		var ret float32
		return ret
	}
	return *o.TaxRate
}

// GetTaxRateOk returns a tuple with the TaxRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseAppliedTax) GetTaxRateOk() (*float32, bool) {
	if o == nil || IsNil(o.TaxRate) {
		return nil, false
	}
	return o.TaxRate, true
}

// HasTaxRate returns a boolean if a field has been set.
func (o *BaseAppliedTax) HasTaxRate() bool {
	if o != nil && !IsNil(o.TaxRate) {
		return true
	}

	return false
}

// SetTaxRate gets a reference to the given float32 and assigns it to the TaxRate field.
func (o *BaseAppliedTax) SetTaxRate(v float32) {
	o.TaxRate = &v
}

// GetTaxDescription returns the TaxDescription field value if set, zero value otherwise.
func (o *BaseAppliedTax) GetTaxDescription() string {
	if o == nil || IsNil(o.TaxDescription) {
		var ret string
		return ret
	}
	return *o.TaxDescription
}

// GetTaxDescriptionOk returns a tuple with the TaxDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseAppliedTax) GetTaxDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.TaxDescription) {
		return nil, false
	}
	return o.TaxDescription, true
}

// HasTaxDescription returns a boolean if a field has been set.
func (o *BaseAppliedTax) HasTaxDescription() bool {
	if o != nil && !IsNil(o.TaxDescription) {
		return true
	}

	return false
}

// SetTaxDescription gets a reference to the given string and assigns it to the TaxDescription field.
func (o *BaseAppliedTax) SetTaxDescription(v string) {
	o.TaxDescription = &v
}

// GetAmountCents returns the AmountCents field value if set, zero value otherwise.
func (o *BaseAppliedTax) GetAmountCents() int32 {
	if o == nil || IsNil(o.AmountCents) {
		var ret int32
		return ret
	}
	return *o.AmountCents
}

// GetAmountCentsOk returns a tuple with the AmountCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseAppliedTax) GetAmountCentsOk() (*int32, bool) {
	if o == nil || IsNil(o.AmountCents) {
		return nil, false
	}
	return o.AmountCents, true
}

// HasAmountCents returns a boolean if a field has been set.
func (o *BaseAppliedTax) HasAmountCents() bool {
	if o != nil && !IsNil(o.AmountCents) {
		return true
	}

	return false
}

// SetAmountCents gets a reference to the given int32 and assigns it to the AmountCents field.
func (o *BaseAppliedTax) SetAmountCents(v int32) {
	o.AmountCents = &v
}

// GetAmountCurrency returns the AmountCurrency field value if set, zero value otherwise.
func (o *BaseAppliedTax) GetAmountCurrency() Currency {
	if o == nil || IsNil(o.AmountCurrency) {
		var ret Currency
		return ret
	}
	return *o.AmountCurrency
}

// GetAmountCurrencyOk returns a tuple with the AmountCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseAppliedTax) GetAmountCurrencyOk() (*Currency, bool) {
	if o == nil || IsNil(o.AmountCurrency) {
		return nil, false
	}
	return o.AmountCurrency, true
}

// HasAmountCurrency returns a boolean if a field has been set.
func (o *BaseAppliedTax) HasAmountCurrency() bool {
	if o != nil && !IsNil(o.AmountCurrency) {
		return true
	}

	return false
}

// SetAmountCurrency gets a reference to the given Currency and assigns it to the AmountCurrency field.
func (o *BaseAppliedTax) SetAmountCurrency(v Currency) {
	o.AmountCurrency = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *BaseAppliedTax) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseAppliedTax) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *BaseAppliedTax) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *BaseAppliedTax) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

func (o BaseAppliedTax) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaseAppliedTax) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LagoId) {
		toSerialize["lago_id"] = o.LagoId
	}
	if !IsNil(o.LagoTaxId) {
		toSerialize["lago_tax_id"] = o.LagoTaxId
	}
	if !IsNil(o.TaxName) {
		toSerialize["tax_name"] = o.TaxName
	}
	if !IsNil(o.TaxCode) {
		toSerialize["tax_code"] = o.TaxCode
	}
	if !IsNil(o.TaxRate) {
		toSerialize["tax_rate"] = o.TaxRate
	}
	if !IsNil(o.TaxDescription) {
		toSerialize["tax_description"] = o.TaxDescription
	}
	if !IsNil(o.AmountCents) {
		toSerialize["amount_cents"] = o.AmountCents
	}
	if !IsNil(o.AmountCurrency) {
		toSerialize["amount_currency"] = o.AmountCurrency
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	return toSerialize, nil
}

type NullableBaseAppliedTax struct {
	value *BaseAppliedTax
	isSet bool
}

func (v NullableBaseAppliedTax) Get() *BaseAppliedTax {
	return v.value
}

func (v *NullableBaseAppliedTax) Set(val *BaseAppliedTax) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseAppliedTax) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseAppliedTax) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseAppliedTax(val *BaseAppliedTax) *NullableBaseAppliedTax {
	return &NullableBaseAppliedTax{value: val, isSet: true}
}

func (v NullableBaseAppliedTax) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseAppliedTax) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


