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

// checks if the CustomerChargeUsageObjectGroupsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerChargeUsageObjectGroupsInner{}

// CustomerChargeUsageObjectGroupsInner struct for CustomerChargeUsageObjectGroupsInner
type CustomerChargeUsageObjectGroupsInner struct {
	// Unique identifier assigned to the group within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the group record within the Lago system.
	LagoId *string `json:"lago_id,omitempty"`
	// The group key, only returned for groups with two dimensions.
	Key NullableString `json:"key,omitempty"`
	// The group value.
	Value *string `json:"value,omitempty"`
	// The number of units consumed for a specific group related to a charge item.
	Units *string `json:"units,omitempty"`
	// The quantity of usage events that have been recorded for a particular charge during the specified time period. These events may also be referred to as the number of transactions in some contexts.
	EventsCount *int32 `json:"events_count,omitempty"`
	// The amount in cents, tax excluded, consumed for a specific group related to a charge item.
	AmountCents *int32 `json:"amount_cents,omitempty"`
}

// NewCustomerChargeUsageObjectGroupsInner instantiates a new CustomerChargeUsageObjectGroupsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerChargeUsageObjectGroupsInner() *CustomerChargeUsageObjectGroupsInner {
	this := CustomerChargeUsageObjectGroupsInner{}
	return &this
}

// NewCustomerChargeUsageObjectGroupsInnerWithDefaults instantiates a new CustomerChargeUsageObjectGroupsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerChargeUsageObjectGroupsInnerWithDefaults() *CustomerChargeUsageObjectGroupsInner {
	this := CustomerChargeUsageObjectGroupsInner{}
	return &this
}

// GetLagoId returns the LagoId field value if set, zero value otherwise.
func (o *CustomerChargeUsageObjectGroupsInner) GetLagoId() string {
	if o == nil || IsNil(o.LagoId) {
		var ret string
		return ret
	}
	return *o.LagoId
}

// GetLagoIdOk returns a tuple with the LagoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerChargeUsageObjectGroupsInner) GetLagoIdOk() (*string, bool) {
	if o == nil || IsNil(o.LagoId) {
		return nil, false
	}
	return o.LagoId, true
}

// HasLagoId returns a boolean if a field has been set.
func (o *CustomerChargeUsageObjectGroupsInner) HasLagoId() bool {
	if o != nil && !IsNil(o.LagoId) {
		return true
	}

	return false
}

// SetLagoId gets a reference to the given string and assigns it to the LagoId field.
func (o *CustomerChargeUsageObjectGroupsInner) SetLagoId(v string) {
	o.LagoId = &v
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerChargeUsageObjectGroupsInner) GetKey() string {
	if o == nil || IsNil(o.Key.Get()) {
		var ret string
		return ret
	}
	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerChargeUsageObjectGroupsInner) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// HasKey returns a boolean if a field has been set.
func (o *CustomerChargeUsageObjectGroupsInner) HasKey() bool {
	if o != nil && o.Key.IsSet() {
		return true
	}

	return false
}

// SetKey gets a reference to the given NullableString and assigns it to the Key field.
func (o *CustomerChargeUsageObjectGroupsInner) SetKey(v string) {
	o.Key.Set(&v)
}
// SetKeyNil sets the value for Key to be an explicit nil
func (o *CustomerChargeUsageObjectGroupsInner) SetKeyNil() {
	o.Key.Set(nil)
}

// UnsetKey ensures that no value is present for Key, not even an explicit nil
func (o *CustomerChargeUsageObjectGroupsInner) UnsetKey() {
	o.Key.Unset()
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CustomerChargeUsageObjectGroupsInner) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerChargeUsageObjectGroupsInner) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CustomerChargeUsageObjectGroupsInner) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *CustomerChargeUsageObjectGroupsInner) SetValue(v string) {
	o.Value = &v
}

// GetUnits returns the Units field value if set, zero value otherwise.
func (o *CustomerChargeUsageObjectGroupsInner) GetUnits() string {
	if o == nil || IsNil(o.Units) {
		var ret string
		return ret
	}
	return *o.Units
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerChargeUsageObjectGroupsInner) GetUnitsOk() (*string, bool) {
	if o == nil || IsNil(o.Units) {
		return nil, false
	}
	return o.Units, true
}

// HasUnits returns a boolean if a field has been set.
func (o *CustomerChargeUsageObjectGroupsInner) HasUnits() bool {
	if o != nil && !IsNil(o.Units) {
		return true
	}

	return false
}

// SetUnits gets a reference to the given string and assigns it to the Units field.
func (o *CustomerChargeUsageObjectGroupsInner) SetUnits(v string) {
	o.Units = &v
}

// GetEventsCount returns the EventsCount field value if set, zero value otherwise.
func (o *CustomerChargeUsageObjectGroupsInner) GetEventsCount() int32 {
	if o == nil || IsNil(o.EventsCount) {
		var ret int32
		return ret
	}
	return *o.EventsCount
}

// GetEventsCountOk returns a tuple with the EventsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerChargeUsageObjectGroupsInner) GetEventsCountOk() (*int32, bool) {
	if o == nil || IsNil(o.EventsCount) {
		return nil, false
	}
	return o.EventsCount, true
}

// HasEventsCount returns a boolean if a field has been set.
func (o *CustomerChargeUsageObjectGroupsInner) HasEventsCount() bool {
	if o != nil && !IsNil(o.EventsCount) {
		return true
	}

	return false
}

// SetEventsCount gets a reference to the given int32 and assigns it to the EventsCount field.
func (o *CustomerChargeUsageObjectGroupsInner) SetEventsCount(v int32) {
	o.EventsCount = &v
}

// GetAmountCents returns the AmountCents field value if set, zero value otherwise.
func (o *CustomerChargeUsageObjectGroupsInner) GetAmountCents() int32 {
	if o == nil || IsNil(o.AmountCents) {
		var ret int32
		return ret
	}
	return *o.AmountCents
}

// GetAmountCentsOk returns a tuple with the AmountCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerChargeUsageObjectGroupsInner) GetAmountCentsOk() (*int32, bool) {
	if o == nil || IsNil(o.AmountCents) {
		return nil, false
	}
	return o.AmountCents, true
}

// HasAmountCents returns a boolean if a field has been set.
func (o *CustomerChargeUsageObjectGroupsInner) HasAmountCents() bool {
	if o != nil && !IsNil(o.AmountCents) {
		return true
	}

	return false
}

// SetAmountCents gets a reference to the given int32 and assigns it to the AmountCents field.
func (o *CustomerChargeUsageObjectGroupsInner) SetAmountCents(v int32) {
	o.AmountCents = &v
}

func (o CustomerChargeUsageObjectGroupsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerChargeUsageObjectGroupsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LagoId) {
		toSerialize["lago_id"] = o.LagoId
	}
	if o.Key.IsSet() {
		toSerialize["key"] = o.Key.Get()
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Units) {
		toSerialize["units"] = o.Units
	}
	if !IsNil(o.EventsCount) {
		toSerialize["events_count"] = o.EventsCount
	}
	if !IsNil(o.AmountCents) {
		toSerialize["amount_cents"] = o.AmountCents
	}
	return toSerialize, nil
}

type NullableCustomerChargeUsageObjectGroupsInner struct {
	value *CustomerChargeUsageObjectGroupsInner
	isSet bool
}

func (v NullableCustomerChargeUsageObjectGroupsInner) Get() *CustomerChargeUsageObjectGroupsInner {
	return v.value
}

func (v *NullableCustomerChargeUsageObjectGroupsInner) Set(val *CustomerChargeUsageObjectGroupsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerChargeUsageObjectGroupsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerChargeUsageObjectGroupsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerChargeUsageObjectGroupsInner(val *CustomerChargeUsageObjectGroupsInner) *NullableCustomerChargeUsageObjectGroupsInner {
	return &NullableCustomerChargeUsageObjectGroupsInner{value: val, isSet: true}
}

func (v NullableCustomerChargeUsageObjectGroupsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerChargeUsageObjectGroupsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


