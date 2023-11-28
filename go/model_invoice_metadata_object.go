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

// checks if the InvoiceMetadataObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvoiceMetadataObject{}

// InvoiceMetadataObject struct for InvoiceMetadataObject
type InvoiceMetadataObject struct {
	// Unique identifier assigned to the invoice metadata within the Lago application.
	LagoId *string `json:"lago_id,omitempty"`
	// Represents the key of the metadata’s key-value pair.
	Key *string `json:"key,omitempty"`
	// Represents the value of the metadata’s key-value pair.
	Value *string `json:"value,omitempty"`
	// The date and time when the metadata object was created. It follows the ISO 8601 datetime format and is expressed in Coordinated Universal Time (UTC).
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

// NewInvoiceMetadataObject instantiates a new InvoiceMetadataObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceMetadataObject() *InvoiceMetadataObject {
	this := InvoiceMetadataObject{}
	return &this
}

// NewInvoiceMetadataObjectWithDefaults instantiates a new InvoiceMetadataObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceMetadataObjectWithDefaults() *InvoiceMetadataObject {
	this := InvoiceMetadataObject{}
	return &this
}

// GetLagoId returns the LagoId field value if set, zero value otherwise.
func (o *InvoiceMetadataObject) GetLagoId() string {
	if o == nil || IsNil(o.LagoId) {
		var ret string
		return ret
	}
	return *o.LagoId
}

// GetLagoIdOk returns a tuple with the LagoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceMetadataObject) GetLagoIdOk() (*string, bool) {
	if o == nil || IsNil(o.LagoId) {
		return nil, false
	}
	return o.LagoId, true
}

// HasLagoId returns a boolean if a field has been set.
func (o *InvoiceMetadataObject) HasLagoId() bool {
	if o != nil && !IsNil(o.LagoId) {
		return true
	}

	return false
}

// SetLagoId gets a reference to the given string and assigns it to the LagoId field.
func (o *InvoiceMetadataObject) SetLagoId(v string) {
	o.LagoId = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *InvoiceMetadataObject) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceMetadataObject) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *InvoiceMetadataObject) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *InvoiceMetadataObject) SetKey(v string) {
	o.Key = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *InvoiceMetadataObject) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceMetadataObject) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *InvoiceMetadataObject) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *InvoiceMetadataObject) SetValue(v string) {
	o.Value = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *InvoiceMetadataObject) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceMetadataObject) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *InvoiceMetadataObject) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *InvoiceMetadataObject) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

func (o InvoiceMetadataObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvoiceMetadataObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LagoId) {
		toSerialize["lago_id"] = o.LagoId
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	return toSerialize, nil
}

type NullableInvoiceMetadataObject struct {
	value *InvoiceMetadataObject
	isSet bool
}

func (v NullableInvoiceMetadataObject) Get() *InvoiceMetadataObject {
	return v.value
}

func (v *NullableInvoiceMetadataObject) Set(val *InvoiceMetadataObject) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceMetadataObject) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceMetadataObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceMetadataObject(val *InvoiceMetadataObject) *NullableInvoiceMetadataObject {
	return &NullableInvoiceMetadataObject{value: val, isSet: true}
}

func (v NullableInvoiceMetadataObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceMetadataObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


