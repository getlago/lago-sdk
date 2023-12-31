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

// checks if the CustomerCreateInputCustomer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerCreateInputCustomer{}

// CustomerCreateInputCustomer struct for CustomerCreateInputCustomer
type CustomerCreateInputCustomer struct {
	// The customer external unique identifier (provided by your own application)
	ExternalId string `json:"external_id"`
	// The first line of the billing address
	AddressLine1 NullableString `json:"address_line1,omitempty"`
	// The second line of the billing address
	AddressLine2 NullableString `json:"address_line2,omitempty"`
	// The city of the customer's billing address
	City NullableString `json:"city,omitempty"`
	Country *Country `json:"country,omitempty"`
	Currency *Currency `json:"currency,omitempty"`
	// The email of the customer
	Email NullableString `json:"email,omitempty"`
	// The legal company name of the customer
	LegalName NullableString `json:"legal_name,omitempty"`
	// The legal company number of the customer
	LegalNumber NullableString `json:"legal_number,omitempty"`
	// The logo URL of the customer
	LogoUrl NullableString `json:"logo_url,omitempty"`
	// The full name of the customer
	Name NullableString `json:"name,omitempty"`
	// The phone number of the customer
	Phone NullableString `json:"phone,omitempty"`
	// The state of the customer's billing address
	State NullableString `json:"state,omitempty"`
	// List of unique code used to identify the taxes.
	TaxCodes []string `json:"tax_codes,omitempty"`
	// The tax identification number of the customer
	TaxIdentificationNumber NullableString `json:"tax_identification_number,omitempty"`
	Timezone *Timezone `json:"timezone,omitempty"`
	// The custom website URL of the customer
	Url NullableString `json:"url,omitempty"`
	// The zipcode of the customer's billing address
	Zipcode NullableString `json:"zipcode,omitempty"`
	// The net payment term, expressed in days, specifies the duration within which a customer is expected to remit payment after the invoice is finalized.
	NetPaymentTerm NullableInt32 `json:"net_payment_term,omitempty"`
	BillingConfiguration *CustomerBillingConfiguration `json:"billing_configuration,omitempty"`
	// Set of key-value pairs that you can attach to a customer. This can be useful for storing additional information about the customer in a structured format
	Metadata []CustomerCreateInputCustomerMetadataInner `json:"metadata,omitempty"`
}

type _CustomerCreateInputCustomer CustomerCreateInputCustomer

// NewCustomerCreateInputCustomer instantiates a new CustomerCreateInputCustomer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerCreateInputCustomer(externalId string) *CustomerCreateInputCustomer {
	this := CustomerCreateInputCustomer{}
	this.ExternalId = externalId
	return &this
}

// NewCustomerCreateInputCustomerWithDefaults instantiates a new CustomerCreateInputCustomer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerCreateInputCustomerWithDefaults() *CustomerCreateInputCustomer {
	this := CustomerCreateInputCustomer{}
	return &this
}

// GetExternalId returns the ExternalId field value
func (o *CustomerCreateInputCustomer) GetExternalId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value
// and a boolean to check if the value has been set.
func (o *CustomerCreateInputCustomer) GetExternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalId, true
}

// SetExternalId sets field value
func (o *CustomerCreateInputCustomer) SetExternalId(v string) {
	o.ExternalId = v
}

// GetAddressLine1 returns the AddressLine1 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerCreateInputCustomer) GetAddressLine1() string {
	if o == nil || IsNil(o.AddressLine1.Get()) {
		var ret string
		return ret
	}
	return *o.AddressLine1.Get()
}

// GetAddressLine1Ok returns a tuple with the AddressLine1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerCreateInputCustomer) GetAddressLine1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AddressLine1.Get(), o.AddressLine1.IsSet()
}

// HasAddressLine1 returns a boolean if a field has been set.
func (o *CustomerCreateInputCustomer) HasAddressLine1() bool {
	if o != nil && o.AddressLine1.IsSet() {
		return true
	}

	return false
}

// SetAddressLine1 gets a reference to the given NullableString and assigns it to the AddressLine1 field.
func (o *CustomerCreateInputCustomer) SetAddressLine1(v string) {
	o.AddressLine1.Set(&v)
}
// SetAddressLine1Nil sets the value for AddressLine1 to be an explicit nil
func (o *CustomerCreateInputCustomer) SetAddressLine1Nil() {
	o.AddressLine1.Set(nil)
}

// UnsetAddressLine1 ensures that no value is present for AddressLine1, not even an explicit nil
func (o *CustomerCreateInputCustomer) UnsetAddressLine1() {
	o.AddressLine1.Unset()
}

// GetAddressLine2 returns the AddressLine2 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerCreateInputCustomer) GetAddressLine2() string {
	if o == nil || IsNil(o.AddressLine2.Get()) {
		var ret string
		return ret
	}
	return *o.AddressLine2.Get()
}

// GetAddressLine2Ok returns a tuple with the AddressLine2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerCreateInputCustomer) GetAddressLine2Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AddressLine2.Get(), o.AddressLine2.IsSet()
}

// HasAddressLine2 returns a boolean if a field has been set.
func (o *CustomerCreateInputCustomer) HasAddressLine2() bool {
	if o != nil && o.AddressLine2.IsSet() {
		return true
	}

	return false
}

// SetAddressLine2 gets a reference to the given NullableString and assigns it to the AddressLine2 field.
func (o *CustomerCreateInputCustomer) SetAddressLine2(v string) {
	o.AddressLine2.Set(&v)
}
// SetAddressLine2Nil sets the value for AddressLine2 to be an explicit nil
func (o *CustomerCreateInputCustomer) SetAddressLine2Nil() {
	o.AddressLine2.Set(nil)
}

// UnsetAddressLine2 ensures that no value is present for AddressLine2, not even an explicit nil
func (o *CustomerCreateInputCustomer) UnsetAddressLine2() {
	o.AddressLine2.Unset()
}

// GetCity returns the City field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerCreateInputCustomer) GetCity() string {
	if o == nil || IsNil(o.City.Get()) {
		var ret string
		return ret
	}
	return *o.City.Get()
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerCreateInputCustomer) GetCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.City.Get(), o.City.IsSet()
}

// HasCity returns a boolean if a field has been set.
func (o *CustomerCreateInputCustomer) HasCity() bool {
	if o != nil && o.City.IsSet() {
		return true
	}

	return false
}

// SetCity gets a reference to the given NullableString and assigns it to the City field.
func (o *CustomerCreateInputCustomer) SetCity(v string) {
	o.City.Set(&v)
}
// SetCityNil sets the value for City to be an explicit nil
func (o *CustomerCreateInputCustomer) SetCityNil() {
	o.City.Set(nil)
}

// UnsetCity ensures that no value is present for City, not even an explicit nil
func (o *CustomerCreateInputCustomer) UnsetCity() {
	o.City.Unset()
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *CustomerCreateInputCustomer) GetCountry() Country {
	if o == nil || IsNil(o.Country) {
		var ret Country
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerCreateInputCustomer) GetCountryOk() (*Country, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *CustomerCreateInputCustomer) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given Country and assigns it to the Country field.
func (o *CustomerCreateInputCustomer) SetCountry(v Country) {
	o.Country = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *CustomerCreateInputCustomer) GetCurrency() Currency {
	if o == nil || IsNil(o.Currency) {
		var ret Currency
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerCreateInputCustomer) GetCurrencyOk() (*Currency, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *CustomerCreateInputCustomer) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given Currency and assigns it to the Currency field.
func (o *CustomerCreateInputCustomer) SetCurrency(v Currency) {
	o.Currency = &v
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerCreateInputCustomer) GetEmail() string {
	if o == nil || IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerCreateInputCustomer) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *CustomerCreateInputCustomer) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *CustomerCreateInputCustomer) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *CustomerCreateInputCustomer) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *CustomerCreateInputCustomer) UnsetEmail() {
	o.Email.Unset()
}

// GetLegalName returns the LegalName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerCreateInputCustomer) GetLegalName() string {
	if o == nil || IsNil(o.LegalName.Get()) {
		var ret string
		return ret
	}
	return *o.LegalName.Get()
}

// GetLegalNameOk returns a tuple with the LegalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerCreateInputCustomer) GetLegalNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LegalName.Get(), o.LegalName.IsSet()
}

// HasLegalName returns a boolean if a field has been set.
func (o *CustomerCreateInputCustomer) HasLegalName() bool {
	if o != nil && o.LegalName.IsSet() {
		return true
	}

	return false
}

// SetLegalName gets a reference to the given NullableString and assigns it to the LegalName field.
func (o *CustomerCreateInputCustomer) SetLegalName(v string) {
	o.LegalName.Set(&v)
}
// SetLegalNameNil sets the value for LegalName to be an explicit nil
func (o *CustomerCreateInputCustomer) SetLegalNameNil() {
	o.LegalName.Set(nil)
}

// UnsetLegalName ensures that no value is present for LegalName, not even an explicit nil
func (o *CustomerCreateInputCustomer) UnsetLegalName() {
	o.LegalName.Unset()
}

// GetLegalNumber returns the LegalNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerCreateInputCustomer) GetLegalNumber() string {
	if o == nil || IsNil(o.LegalNumber.Get()) {
		var ret string
		return ret
	}
	return *o.LegalNumber.Get()
}

// GetLegalNumberOk returns a tuple with the LegalNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerCreateInputCustomer) GetLegalNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LegalNumber.Get(), o.LegalNumber.IsSet()
}

// HasLegalNumber returns a boolean if a field has been set.
func (o *CustomerCreateInputCustomer) HasLegalNumber() bool {
	if o != nil && o.LegalNumber.IsSet() {
		return true
	}

	return false
}

// SetLegalNumber gets a reference to the given NullableString and assigns it to the LegalNumber field.
func (o *CustomerCreateInputCustomer) SetLegalNumber(v string) {
	o.LegalNumber.Set(&v)
}
// SetLegalNumberNil sets the value for LegalNumber to be an explicit nil
func (o *CustomerCreateInputCustomer) SetLegalNumberNil() {
	o.LegalNumber.Set(nil)
}

// UnsetLegalNumber ensures that no value is present for LegalNumber, not even an explicit nil
func (o *CustomerCreateInputCustomer) UnsetLegalNumber() {
	o.LegalNumber.Unset()
}

// GetLogoUrl returns the LogoUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerCreateInputCustomer) GetLogoUrl() string {
	if o == nil || IsNil(o.LogoUrl.Get()) {
		var ret string
		return ret
	}
	return *o.LogoUrl.Get()
}

// GetLogoUrlOk returns a tuple with the LogoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerCreateInputCustomer) GetLogoUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LogoUrl.Get(), o.LogoUrl.IsSet()
}

// HasLogoUrl returns a boolean if a field has been set.
func (o *CustomerCreateInputCustomer) HasLogoUrl() bool {
	if o != nil && o.LogoUrl.IsSet() {
		return true
	}

	return false
}

// SetLogoUrl gets a reference to the given NullableString and assigns it to the LogoUrl field.
func (o *CustomerCreateInputCustomer) SetLogoUrl(v string) {
	o.LogoUrl.Set(&v)
}
// SetLogoUrlNil sets the value for LogoUrl to be an explicit nil
func (o *CustomerCreateInputCustomer) SetLogoUrlNil() {
	o.LogoUrl.Set(nil)
}

// UnsetLogoUrl ensures that no value is present for LogoUrl, not even an explicit nil
func (o *CustomerCreateInputCustomer) UnsetLogoUrl() {
	o.LogoUrl.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerCreateInputCustomer) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerCreateInputCustomer) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *CustomerCreateInputCustomer) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *CustomerCreateInputCustomer) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *CustomerCreateInputCustomer) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *CustomerCreateInputCustomer) UnsetName() {
	o.Name.Unset()
}

// GetPhone returns the Phone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerCreateInputCustomer) GetPhone() string {
	if o == nil || IsNil(o.Phone.Get()) {
		var ret string
		return ret
	}
	return *o.Phone.Get()
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerCreateInputCustomer) GetPhoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Phone.Get(), o.Phone.IsSet()
}

// HasPhone returns a boolean if a field has been set.
func (o *CustomerCreateInputCustomer) HasPhone() bool {
	if o != nil && o.Phone.IsSet() {
		return true
	}

	return false
}

// SetPhone gets a reference to the given NullableString and assigns it to the Phone field.
func (o *CustomerCreateInputCustomer) SetPhone(v string) {
	o.Phone.Set(&v)
}
// SetPhoneNil sets the value for Phone to be an explicit nil
func (o *CustomerCreateInputCustomer) SetPhoneNil() {
	o.Phone.Set(nil)
}

// UnsetPhone ensures that no value is present for Phone, not even an explicit nil
func (o *CustomerCreateInputCustomer) UnsetPhone() {
	o.Phone.Unset()
}

// GetState returns the State field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerCreateInputCustomer) GetState() string {
	if o == nil || IsNil(o.State.Get()) {
		var ret string
		return ret
	}
	return *o.State.Get()
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerCreateInputCustomer) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.State.Get(), o.State.IsSet()
}

// HasState returns a boolean if a field has been set.
func (o *CustomerCreateInputCustomer) HasState() bool {
	if o != nil && o.State.IsSet() {
		return true
	}

	return false
}

// SetState gets a reference to the given NullableString and assigns it to the State field.
func (o *CustomerCreateInputCustomer) SetState(v string) {
	o.State.Set(&v)
}
// SetStateNil sets the value for State to be an explicit nil
func (o *CustomerCreateInputCustomer) SetStateNil() {
	o.State.Set(nil)
}

// UnsetState ensures that no value is present for State, not even an explicit nil
func (o *CustomerCreateInputCustomer) UnsetState() {
	o.State.Unset()
}

// GetTaxCodes returns the TaxCodes field value if set, zero value otherwise.
func (o *CustomerCreateInputCustomer) GetTaxCodes() []string {
	if o == nil || IsNil(o.TaxCodes) {
		var ret []string
		return ret
	}
	return o.TaxCodes
}

// GetTaxCodesOk returns a tuple with the TaxCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerCreateInputCustomer) GetTaxCodesOk() ([]string, bool) {
	if o == nil || IsNil(o.TaxCodes) {
		return nil, false
	}
	return o.TaxCodes, true
}

// HasTaxCodes returns a boolean if a field has been set.
func (o *CustomerCreateInputCustomer) HasTaxCodes() bool {
	if o != nil && !IsNil(o.TaxCodes) {
		return true
	}

	return false
}

// SetTaxCodes gets a reference to the given []string and assigns it to the TaxCodes field.
func (o *CustomerCreateInputCustomer) SetTaxCodes(v []string) {
	o.TaxCodes = v
}

// GetTaxIdentificationNumber returns the TaxIdentificationNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerCreateInputCustomer) GetTaxIdentificationNumber() string {
	if o == nil || IsNil(o.TaxIdentificationNumber.Get()) {
		var ret string
		return ret
	}
	return *o.TaxIdentificationNumber.Get()
}

// GetTaxIdentificationNumberOk returns a tuple with the TaxIdentificationNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerCreateInputCustomer) GetTaxIdentificationNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TaxIdentificationNumber.Get(), o.TaxIdentificationNumber.IsSet()
}

// HasTaxIdentificationNumber returns a boolean if a field has been set.
func (o *CustomerCreateInputCustomer) HasTaxIdentificationNumber() bool {
	if o != nil && o.TaxIdentificationNumber.IsSet() {
		return true
	}

	return false
}

// SetTaxIdentificationNumber gets a reference to the given NullableString and assigns it to the TaxIdentificationNumber field.
func (o *CustomerCreateInputCustomer) SetTaxIdentificationNumber(v string) {
	o.TaxIdentificationNumber.Set(&v)
}
// SetTaxIdentificationNumberNil sets the value for TaxIdentificationNumber to be an explicit nil
func (o *CustomerCreateInputCustomer) SetTaxIdentificationNumberNil() {
	o.TaxIdentificationNumber.Set(nil)
}

// UnsetTaxIdentificationNumber ensures that no value is present for TaxIdentificationNumber, not even an explicit nil
func (o *CustomerCreateInputCustomer) UnsetTaxIdentificationNumber() {
	o.TaxIdentificationNumber.Unset()
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *CustomerCreateInputCustomer) GetTimezone() Timezone {
	if o == nil || IsNil(o.Timezone) {
		var ret Timezone
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerCreateInputCustomer) GetTimezoneOk() (*Timezone, bool) {
	if o == nil || IsNil(o.Timezone) {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *CustomerCreateInputCustomer) HasTimezone() bool {
	if o != nil && !IsNil(o.Timezone) {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given Timezone and assigns it to the Timezone field.
func (o *CustomerCreateInputCustomer) SetTimezone(v Timezone) {
	o.Timezone = &v
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerCreateInputCustomer) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerCreateInputCustomer) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *CustomerCreateInputCustomer) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *CustomerCreateInputCustomer) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *CustomerCreateInputCustomer) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *CustomerCreateInputCustomer) UnsetUrl() {
	o.Url.Unset()
}

// GetZipcode returns the Zipcode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerCreateInputCustomer) GetZipcode() string {
	if o == nil || IsNil(o.Zipcode.Get()) {
		var ret string
		return ret
	}
	return *o.Zipcode.Get()
}

// GetZipcodeOk returns a tuple with the Zipcode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerCreateInputCustomer) GetZipcodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Zipcode.Get(), o.Zipcode.IsSet()
}

// HasZipcode returns a boolean if a field has been set.
func (o *CustomerCreateInputCustomer) HasZipcode() bool {
	if o != nil && o.Zipcode.IsSet() {
		return true
	}

	return false
}

// SetZipcode gets a reference to the given NullableString and assigns it to the Zipcode field.
func (o *CustomerCreateInputCustomer) SetZipcode(v string) {
	o.Zipcode.Set(&v)
}
// SetZipcodeNil sets the value for Zipcode to be an explicit nil
func (o *CustomerCreateInputCustomer) SetZipcodeNil() {
	o.Zipcode.Set(nil)
}

// UnsetZipcode ensures that no value is present for Zipcode, not even an explicit nil
func (o *CustomerCreateInputCustomer) UnsetZipcode() {
	o.Zipcode.Unset()
}

// GetNetPaymentTerm returns the NetPaymentTerm field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerCreateInputCustomer) GetNetPaymentTerm() int32 {
	if o == nil || IsNil(o.NetPaymentTerm.Get()) {
		var ret int32
		return ret
	}
	return *o.NetPaymentTerm.Get()
}

// GetNetPaymentTermOk returns a tuple with the NetPaymentTerm field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerCreateInputCustomer) GetNetPaymentTermOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.NetPaymentTerm.Get(), o.NetPaymentTerm.IsSet()
}

// HasNetPaymentTerm returns a boolean if a field has been set.
func (o *CustomerCreateInputCustomer) HasNetPaymentTerm() bool {
	if o != nil && o.NetPaymentTerm.IsSet() {
		return true
	}

	return false
}

// SetNetPaymentTerm gets a reference to the given NullableInt32 and assigns it to the NetPaymentTerm field.
func (o *CustomerCreateInputCustomer) SetNetPaymentTerm(v int32) {
	o.NetPaymentTerm.Set(&v)
}
// SetNetPaymentTermNil sets the value for NetPaymentTerm to be an explicit nil
func (o *CustomerCreateInputCustomer) SetNetPaymentTermNil() {
	o.NetPaymentTerm.Set(nil)
}

// UnsetNetPaymentTerm ensures that no value is present for NetPaymentTerm, not even an explicit nil
func (o *CustomerCreateInputCustomer) UnsetNetPaymentTerm() {
	o.NetPaymentTerm.Unset()
}

// GetBillingConfiguration returns the BillingConfiguration field value if set, zero value otherwise.
func (o *CustomerCreateInputCustomer) GetBillingConfiguration() CustomerBillingConfiguration {
	if o == nil || IsNil(o.BillingConfiguration) {
		var ret CustomerBillingConfiguration
		return ret
	}
	return *o.BillingConfiguration
}

// GetBillingConfigurationOk returns a tuple with the BillingConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerCreateInputCustomer) GetBillingConfigurationOk() (*CustomerBillingConfiguration, bool) {
	if o == nil || IsNil(o.BillingConfiguration) {
		return nil, false
	}
	return o.BillingConfiguration, true
}

// HasBillingConfiguration returns a boolean if a field has been set.
func (o *CustomerCreateInputCustomer) HasBillingConfiguration() bool {
	if o != nil && !IsNil(o.BillingConfiguration) {
		return true
	}

	return false
}

// SetBillingConfiguration gets a reference to the given CustomerBillingConfiguration and assigns it to the BillingConfiguration field.
func (o *CustomerCreateInputCustomer) SetBillingConfiguration(v CustomerBillingConfiguration) {
	o.BillingConfiguration = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *CustomerCreateInputCustomer) GetMetadata() []CustomerCreateInputCustomerMetadataInner {
	if o == nil || IsNil(o.Metadata) {
		var ret []CustomerCreateInputCustomerMetadataInner
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerCreateInputCustomer) GetMetadataOk() ([]CustomerCreateInputCustomerMetadataInner, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CustomerCreateInputCustomer) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given []CustomerCreateInputCustomerMetadataInner and assigns it to the Metadata field.
func (o *CustomerCreateInputCustomer) SetMetadata(v []CustomerCreateInputCustomerMetadataInner) {
	o.Metadata = v
}

func (o CustomerCreateInputCustomer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerCreateInputCustomer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["external_id"] = o.ExternalId
	if o.AddressLine1.IsSet() {
		toSerialize["address_line1"] = o.AddressLine1.Get()
	}
	if o.AddressLine2.IsSet() {
		toSerialize["address_line2"] = o.AddressLine2.Get()
	}
	if o.City.IsSet() {
		toSerialize["city"] = o.City.Get()
	}
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	if o.LegalName.IsSet() {
		toSerialize["legal_name"] = o.LegalName.Get()
	}
	if o.LegalNumber.IsSet() {
		toSerialize["legal_number"] = o.LegalNumber.Get()
	}
	if o.LogoUrl.IsSet() {
		toSerialize["logo_url"] = o.LogoUrl.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Phone.IsSet() {
		toSerialize["phone"] = o.Phone.Get()
	}
	if o.State.IsSet() {
		toSerialize["state"] = o.State.Get()
	}
	if !IsNil(o.TaxCodes) {
		toSerialize["tax_codes"] = o.TaxCodes
	}
	if o.TaxIdentificationNumber.IsSet() {
		toSerialize["tax_identification_number"] = o.TaxIdentificationNumber.Get()
	}
	if !IsNil(o.Timezone) {
		toSerialize["timezone"] = o.Timezone
	}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	if o.Zipcode.IsSet() {
		toSerialize["zipcode"] = o.Zipcode.Get()
	}
	if o.NetPaymentTerm.IsSet() {
		toSerialize["net_payment_term"] = o.NetPaymentTerm.Get()
	}
	if !IsNil(o.BillingConfiguration) {
		toSerialize["billing_configuration"] = o.BillingConfiguration
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

func (o *CustomerCreateInputCustomer) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"external_id",
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

	varCustomerCreateInputCustomer := _CustomerCreateInputCustomer{}

	err = json.Unmarshal(bytes, &varCustomerCreateInputCustomer)

	if err != nil {
		return err
	}

	*o = CustomerCreateInputCustomer(varCustomerCreateInputCustomer)

	return err
}

type NullableCustomerCreateInputCustomer struct {
	value *CustomerCreateInputCustomer
	isSet bool
}

func (v NullableCustomerCreateInputCustomer) Get() *CustomerCreateInputCustomer {
	return v.value
}

func (v *NullableCustomerCreateInputCustomer) Set(val *CustomerCreateInputCustomer) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerCreateInputCustomer) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerCreateInputCustomer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerCreateInputCustomer(val *CustomerCreateInputCustomer) *NullableCustomerCreateInputCustomer {
	return &NullableCustomerCreateInputCustomer{value: val, isSet: true}
}

func (v NullableCustomerCreateInputCustomer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerCreateInputCustomer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


