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

// checks if the CustomerChargeUsageObjectCharge type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerChargeUsageObjectCharge{}

// CustomerChargeUsageObjectCharge Object listing all the properties for a specific charge item.
type CustomerChargeUsageObjectCharge struct {
	// Unique identifier assigned to the charge within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the charge’s record within the Lago system.
	LagoId string `json:"lago_id"`
	// The pricing model applied to this charge. Possible values are standard, `graduated`, `percentage`, `package` or `volume`.
	ChargeModel string `json:"charge_model"`
	// Specifies the name that will be displayed on an invoice. If no value is set for this field, the name of the actual charge will be used as the default display name.
	InvoiceDisplayName *string `json:"invoice_display_name,omitempty"`
}

type _CustomerChargeUsageObjectCharge CustomerChargeUsageObjectCharge

// NewCustomerChargeUsageObjectCharge instantiates a new CustomerChargeUsageObjectCharge object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerChargeUsageObjectCharge(lagoId string, chargeModel string) *CustomerChargeUsageObjectCharge {
	this := CustomerChargeUsageObjectCharge{}
	this.LagoId = lagoId
	this.ChargeModel = chargeModel
	return &this
}

// NewCustomerChargeUsageObjectChargeWithDefaults instantiates a new CustomerChargeUsageObjectCharge object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerChargeUsageObjectChargeWithDefaults() *CustomerChargeUsageObjectCharge {
	this := CustomerChargeUsageObjectCharge{}
	return &this
}

// GetLagoId returns the LagoId field value
func (o *CustomerChargeUsageObjectCharge) GetLagoId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LagoId
}

// GetLagoIdOk returns a tuple with the LagoId field value
// and a boolean to check if the value has been set.
func (o *CustomerChargeUsageObjectCharge) GetLagoIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LagoId, true
}

// SetLagoId sets field value
func (o *CustomerChargeUsageObjectCharge) SetLagoId(v string) {
	o.LagoId = v
}

// GetChargeModel returns the ChargeModel field value
func (o *CustomerChargeUsageObjectCharge) GetChargeModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChargeModel
}

// GetChargeModelOk returns a tuple with the ChargeModel field value
// and a boolean to check if the value has been set.
func (o *CustomerChargeUsageObjectCharge) GetChargeModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChargeModel, true
}

// SetChargeModel sets field value
func (o *CustomerChargeUsageObjectCharge) SetChargeModel(v string) {
	o.ChargeModel = v
}

// GetInvoiceDisplayName returns the InvoiceDisplayName field value if set, zero value otherwise.
func (o *CustomerChargeUsageObjectCharge) GetInvoiceDisplayName() string {
	if o == nil || IsNil(o.InvoiceDisplayName) {
		var ret string
		return ret
	}
	return *o.InvoiceDisplayName
}

// GetInvoiceDisplayNameOk returns a tuple with the InvoiceDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerChargeUsageObjectCharge) GetInvoiceDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceDisplayName) {
		return nil, false
	}
	return o.InvoiceDisplayName, true
}

// HasInvoiceDisplayName returns a boolean if a field has been set.
func (o *CustomerChargeUsageObjectCharge) HasInvoiceDisplayName() bool {
	if o != nil && !IsNil(o.InvoiceDisplayName) {
		return true
	}

	return false
}

// SetInvoiceDisplayName gets a reference to the given string and assigns it to the InvoiceDisplayName field.
func (o *CustomerChargeUsageObjectCharge) SetInvoiceDisplayName(v string) {
	o.InvoiceDisplayName = &v
}

func (o CustomerChargeUsageObjectCharge) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerChargeUsageObjectCharge) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["lago_id"] = o.LagoId
	toSerialize["charge_model"] = o.ChargeModel
	if !IsNil(o.InvoiceDisplayName) {
		toSerialize["invoice_display_name"] = o.InvoiceDisplayName
	}
	return toSerialize, nil
}

func (o *CustomerChargeUsageObjectCharge) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"lago_id",
		"charge_model",
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

	varCustomerChargeUsageObjectCharge := _CustomerChargeUsageObjectCharge{}

	err = json.Unmarshal(bytes, &varCustomerChargeUsageObjectCharge)

	if err != nil {
		return err
	}

	*o = CustomerChargeUsageObjectCharge(varCustomerChargeUsageObjectCharge)

	return err
}

type NullableCustomerChargeUsageObjectCharge struct {
	value *CustomerChargeUsageObjectCharge
	isSet bool
}

func (v NullableCustomerChargeUsageObjectCharge) Get() *CustomerChargeUsageObjectCharge {
	return v.value
}

func (v *NullableCustomerChargeUsageObjectCharge) Set(val *CustomerChargeUsageObjectCharge) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerChargeUsageObjectCharge) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerChargeUsageObjectCharge) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerChargeUsageObjectCharge(val *CustomerChargeUsageObjectCharge) *NullableCustomerChargeUsageObjectCharge {
	return &NullableCustomerChargeUsageObjectCharge{value: val, isSet: true}
}

func (v NullableCustomerChargeUsageObjectCharge) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerChargeUsageObjectCharge) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


