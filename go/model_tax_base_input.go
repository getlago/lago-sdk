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
)

// checks if the TaxBaseInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaxBaseInput{}

// TaxBaseInput struct for TaxBaseInput
type TaxBaseInput struct {
	// Name of the tax.
	Name *string `json:"name,omitempty"`
	// Unique code used to identify the tax associated with the API request.
	Code *string `json:"code,omitempty"`
	// The percentage rate of the tax
	Rate *string `json:"rate,omitempty"`
	// Internal description of the taxe
	Description NullableString `json:"description,omitempty"`
	// Set to `true` if the tax is used as one of the organization's default
	AppliedToOrganization *bool `json:"applied_to_organization,omitempty"`
}

// NewTaxBaseInput instantiates a new TaxBaseInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxBaseInput() *TaxBaseInput {
	this := TaxBaseInput{}
	return &this
}

// NewTaxBaseInputWithDefaults instantiates a new TaxBaseInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxBaseInputWithDefaults() *TaxBaseInput {
	this := TaxBaseInput{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TaxBaseInput) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxBaseInput) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TaxBaseInput) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TaxBaseInput) SetName(v string) {
	o.Name = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *TaxBaseInput) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxBaseInput) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *TaxBaseInput) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *TaxBaseInput) SetCode(v string) {
	o.Code = &v
}

// GetRate returns the Rate field value if set, zero value otherwise.
func (o *TaxBaseInput) GetRate() string {
	if o == nil || IsNil(o.Rate) {
		var ret string
		return ret
	}
	return *o.Rate
}

// GetRateOk returns a tuple with the Rate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxBaseInput) GetRateOk() (*string, bool) {
	if o == nil || IsNil(o.Rate) {
		return nil, false
	}
	return o.Rate, true
}

// HasRate returns a boolean if a field has been set.
func (o *TaxBaseInput) HasRate() bool {
	if o != nil && !IsNil(o.Rate) {
		return true
	}

	return false
}

// SetRate gets a reference to the given string and assigns it to the Rate field.
func (o *TaxBaseInput) SetRate(v string) {
	o.Rate = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaxBaseInput) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaxBaseInput) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *TaxBaseInput) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *TaxBaseInput) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *TaxBaseInput) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *TaxBaseInput) UnsetDescription() {
	o.Description.Unset()
}

// GetAppliedToOrganization returns the AppliedToOrganization field value if set, zero value otherwise.
func (o *TaxBaseInput) GetAppliedToOrganization() bool {
	if o == nil || IsNil(o.AppliedToOrganization) {
		var ret bool
		return ret
	}
	return *o.AppliedToOrganization
}

// GetAppliedToOrganizationOk returns a tuple with the AppliedToOrganization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxBaseInput) GetAppliedToOrganizationOk() (*bool, bool) {
	if o == nil || IsNil(o.AppliedToOrganization) {
		return nil, false
	}
	return o.AppliedToOrganization, true
}

// HasAppliedToOrganization returns a boolean if a field has been set.
func (o *TaxBaseInput) HasAppliedToOrganization() bool {
	if o != nil && !IsNil(o.AppliedToOrganization) {
		return true
	}

	return false
}

// SetAppliedToOrganization gets a reference to the given bool and assigns it to the AppliedToOrganization field.
func (o *TaxBaseInput) SetAppliedToOrganization(v bool) {
	o.AppliedToOrganization = &v
}

func (o TaxBaseInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaxBaseInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Rate) {
		toSerialize["rate"] = o.Rate
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !IsNil(o.AppliedToOrganization) {
		toSerialize["applied_to_organization"] = o.AppliedToOrganization
	}
	return toSerialize, nil
}

type NullableTaxBaseInput struct {
	value *TaxBaseInput
	isSet bool
}

func (v NullableTaxBaseInput) Get() *TaxBaseInput {
	return v.value
}

func (v *NullableTaxBaseInput) Set(val *TaxBaseInput) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxBaseInput) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxBaseInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxBaseInput(val *TaxBaseInput) *NullableTaxBaseInput {
	return &NullableTaxBaseInput{value: val, isSet: true}
}

func (v NullableTaxBaseInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxBaseInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


