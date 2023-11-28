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

// checks if the OrganizationBillingConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationBillingConfiguration{}

// OrganizationBillingConfiguration The custom billing settings for your organization.
type OrganizationBillingConfiguration struct {
	// The customer invoice message that appears at the bottom of each billing documents.
	InvoiceFooter NullableString `json:"invoice_footer,omitempty"`
	// The grace period, expressed in days, for finalizing the invoice. This period refers to the additional time granted to your customers beyond the invoice due date to adjust usage and line items. Can be overwritten by the customer’s grace period.
	InvoiceGracePeriod *int32 `json:"invoice_grace_period,omitempty"`
	// The locale of the billing documents, expressed in the ISO 639-1 format. This field indicates the language or regional variant used for the documents content issued or the embeddable customer portal.
	DocumentLocale *string `json:"document_locale,omitempty"`
}

// NewOrganizationBillingConfiguration instantiates a new OrganizationBillingConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationBillingConfiguration() *OrganizationBillingConfiguration {
	this := OrganizationBillingConfiguration{}
	return &this
}

// NewOrganizationBillingConfigurationWithDefaults instantiates a new OrganizationBillingConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationBillingConfigurationWithDefaults() *OrganizationBillingConfiguration {
	this := OrganizationBillingConfiguration{}
	return &this
}

// GetInvoiceFooter returns the InvoiceFooter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationBillingConfiguration) GetInvoiceFooter() string {
	if o == nil || IsNil(o.InvoiceFooter.Get()) {
		var ret string
		return ret
	}
	return *o.InvoiceFooter.Get()
}

// GetInvoiceFooterOk returns a tuple with the InvoiceFooter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationBillingConfiguration) GetInvoiceFooterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InvoiceFooter.Get(), o.InvoiceFooter.IsSet()
}

// HasInvoiceFooter returns a boolean if a field has been set.
func (o *OrganizationBillingConfiguration) HasInvoiceFooter() bool {
	if o != nil && o.InvoiceFooter.IsSet() {
		return true
	}

	return false
}

// SetInvoiceFooter gets a reference to the given NullableString and assigns it to the InvoiceFooter field.
func (o *OrganizationBillingConfiguration) SetInvoiceFooter(v string) {
	o.InvoiceFooter.Set(&v)
}
// SetInvoiceFooterNil sets the value for InvoiceFooter to be an explicit nil
func (o *OrganizationBillingConfiguration) SetInvoiceFooterNil() {
	o.InvoiceFooter.Set(nil)
}

// UnsetInvoiceFooter ensures that no value is present for InvoiceFooter, not even an explicit nil
func (o *OrganizationBillingConfiguration) UnsetInvoiceFooter() {
	o.InvoiceFooter.Unset()
}

// GetInvoiceGracePeriod returns the InvoiceGracePeriod field value if set, zero value otherwise.
func (o *OrganizationBillingConfiguration) GetInvoiceGracePeriod() int32 {
	if o == nil || IsNil(o.InvoiceGracePeriod) {
		var ret int32
		return ret
	}
	return *o.InvoiceGracePeriod
}

// GetInvoiceGracePeriodOk returns a tuple with the InvoiceGracePeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationBillingConfiguration) GetInvoiceGracePeriodOk() (*int32, bool) {
	if o == nil || IsNil(o.InvoiceGracePeriod) {
		return nil, false
	}
	return o.InvoiceGracePeriod, true
}

// HasInvoiceGracePeriod returns a boolean if a field has been set.
func (o *OrganizationBillingConfiguration) HasInvoiceGracePeriod() bool {
	if o != nil && !IsNil(o.InvoiceGracePeriod) {
		return true
	}

	return false
}

// SetInvoiceGracePeriod gets a reference to the given int32 and assigns it to the InvoiceGracePeriod field.
func (o *OrganizationBillingConfiguration) SetInvoiceGracePeriod(v int32) {
	o.InvoiceGracePeriod = &v
}

// GetDocumentLocale returns the DocumentLocale field value if set, zero value otherwise.
func (o *OrganizationBillingConfiguration) GetDocumentLocale() string {
	if o == nil || IsNil(o.DocumentLocale) {
		var ret string
		return ret
	}
	return *o.DocumentLocale
}

// GetDocumentLocaleOk returns a tuple with the DocumentLocale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationBillingConfiguration) GetDocumentLocaleOk() (*string, bool) {
	if o == nil || IsNil(o.DocumentLocale) {
		return nil, false
	}
	return o.DocumentLocale, true
}

// HasDocumentLocale returns a boolean if a field has been set.
func (o *OrganizationBillingConfiguration) HasDocumentLocale() bool {
	if o != nil && !IsNil(o.DocumentLocale) {
		return true
	}

	return false
}

// SetDocumentLocale gets a reference to the given string and assigns it to the DocumentLocale field.
func (o *OrganizationBillingConfiguration) SetDocumentLocale(v string) {
	o.DocumentLocale = &v
}

func (o OrganizationBillingConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationBillingConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.InvoiceFooter.IsSet() {
		toSerialize["invoice_footer"] = o.InvoiceFooter.Get()
	}
	if !IsNil(o.InvoiceGracePeriod) {
		toSerialize["invoice_grace_period"] = o.InvoiceGracePeriod
	}
	if !IsNil(o.DocumentLocale) {
		toSerialize["document_locale"] = o.DocumentLocale
	}
	return toSerialize, nil
}

type NullableOrganizationBillingConfiguration struct {
	value *OrganizationBillingConfiguration
	isSet bool
}

func (v NullableOrganizationBillingConfiguration) Get() *OrganizationBillingConfiguration {
	return v.value
}

func (v *NullableOrganizationBillingConfiguration) Set(val *OrganizationBillingConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationBillingConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationBillingConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationBillingConfiguration(val *OrganizationBillingConfiguration) *NullableOrganizationBillingConfiguration {
	return &NullableOrganizationBillingConfiguration{value: val, isSet: true}
}

func (v NullableOrganizationBillingConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationBillingConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


