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

// checks if the CustomerChargeUsageObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerChargeUsageObject{}

// CustomerChargeUsageObject struct for CustomerChargeUsageObject
type CustomerChargeUsageObject struct {
	// The number of units consumed by the customer for a specific charge item.
	Units string `json:"units"`
	// The quantity of usage events that have been recorded for a particular charge during the specified time period. These events may also be referred to as the number of transactions in some contexts.
	EventsCount int32 `json:"events_count"`
	// The amount in cents, tax excluded, consumed by the customer for a specific charge item.
	AmountCents int32 `json:"amount_cents"`
	AmountCurrency Currency `json:"amount_currency"`
	Charge CustomerChargeUsageObjectCharge `json:"charge"`
	BillableMetric CustomerChargeUsageObjectBillableMetric `json:"billable_metric"`
	// Array of group object, representing multiple dimensions for a charge item.
	Groups []CustomerChargeUsageObjectGroupsInner `json:"groups"`
}

type _CustomerChargeUsageObject CustomerChargeUsageObject

// NewCustomerChargeUsageObject instantiates a new CustomerChargeUsageObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerChargeUsageObject(units string, eventsCount int32, amountCents int32, amountCurrency Currency, charge CustomerChargeUsageObjectCharge, billableMetric CustomerChargeUsageObjectBillableMetric, groups []CustomerChargeUsageObjectGroupsInner) *CustomerChargeUsageObject {
	this := CustomerChargeUsageObject{}
	this.Units = units
	this.EventsCount = eventsCount
	this.AmountCents = amountCents
	this.AmountCurrency = amountCurrency
	this.Charge = charge
	this.BillableMetric = billableMetric
	this.Groups = groups
	return &this
}

// NewCustomerChargeUsageObjectWithDefaults instantiates a new CustomerChargeUsageObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerChargeUsageObjectWithDefaults() *CustomerChargeUsageObject {
	this := CustomerChargeUsageObject{}
	return &this
}

// GetUnits returns the Units field value
func (o *CustomerChargeUsageObject) GetUnits() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Units
}

// GetUnitsOk returns a tuple with the Units field value
// and a boolean to check if the value has been set.
func (o *CustomerChargeUsageObject) GetUnitsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Units, true
}

// SetUnits sets field value
func (o *CustomerChargeUsageObject) SetUnits(v string) {
	o.Units = v
}

// GetEventsCount returns the EventsCount field value
func (o *CustomerChargeUsageObject) GetEventsCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EventsCount
}

// GetEventsCountOk returns a tuple with the EventsCount field value
// and a boolean to check if the value has been set.
func (o *CustomerChargeUsageObject) GetEventsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventsCount, true
}

// SetEventsCount sets field value
func (o *CustomerChargeUsageObject) SetEventsCount(v int32) {
	o.EventsCount = v
}

// GetAmountCents returns the AmountCents field value
func (o *CustomerChargeUsageObject) GetAmountCents() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AmountCents
}

// GetAmountCentsOk returns a tuple with the AmountCents field value
// and a boolean to check if the value has been set.
func (o *CustomerChargeUsageObject) GetAmountCentsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountCents, true
}

// SetAmountCents sets field value
func (o *CustomerChargeUsageObject) SetAmountCents(v int32) {
	o.AmountCents = v
}

// GetAmountCurrency returns the AmountCurrency field value
func (o *CustomerChargeUsageObject) GetAmountCurrency() Currency {
	if o == nil {
		var ret Currency
		return ret
	}

	return o.AmountCurrency
}

// GetAmountCurrencyOk returns a tuple with the AmountCurrency field value
// and a boolean to check if the value has been set.
func (o *CustomerChargeUsageObject) GetAmountCurrencyOk() (*Currency, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountCurrency, true
}

// SetAmountCurrency sets field value
func (o *CustomerChargeUsageObject) SetAmountCurrency(v Currency) {
	o.AmountCurrency = v
}

// GetCharge returns the Charge field value
func (o *CustomerChargeUsageObject) GetCharge() CustomerChargeUsageObjectCharge {
	if o == nil {
		var ret CustomerChargeUsageObjectCharge
		return ret
	}

	return o.Charge
}

// GetChargeOk returns a tuple with the Charge field value
// and a boolean to check if the value has been set.
func (o *CustomerChargeUsageObject) GetChargeOk() (*CustomerChargeUsageObjectCharge, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Charge, true
}

// SetCharge sets field value
func (o *CustomerChargeUsageObject) SetCharge(v CustomerChargeUsageObjectCharge) {
	o.Charge = v
}

// GetBillableMetric returns the BillableMetric field value
func (o *CustomerChargeUsageObject) GetBillableMetric() CustomerChargeUsageObjectBillableMetric {
	if o == nil {
		var ret CustomerChargeUsageObjectBillableMetric
		return ret
	}

	return o.BillableMetric
}

// GetBillableMetricOk returns a tuple with the BillableMetric field value
// and a boolean to check if the value has been set.
func (o *CustomerChargeUsageObject) GetBillableMetricOk() (*CustomerChargeUsageObjectBillableMetric, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BillableMetric, true
}

// SetBillableMetric sets field value
func (o *CustomerChargeUsageObject) SetBillableMetric(v CustomerChargeUsageObjectBillableMetric) {
	o.BillableMetric = v
}

// GetGroups returns the Groups field value
func (o *CustomerChargeUsageObject) GetGroups() []CustomerChargeUsageObjectGroupsInner {
	if o == nil {
		var ret []CustomerChargeUsageObjectGroupsInner
		return ret
	}

	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value
// and a boolean to check if the value has been set.
func (o *CustomerChargeUsageObject) GetGroupsOk() ([]CustomerChargeUsageObjectGroupsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Groups, true
}

// SetGroups sets field value
func (o *CustomerChargeUsageObject) SetGroups(v []CustomerChargeUsageObjectGroupsInner) {
	o.Groups = v
}

func (o CustomerChargeUsageObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerChargeUsageObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["units"] = o.Units
	toSerialize["events_count"] = o.EventsCount
	toSerialize["amount_cents"] = o.AmountCents
	toSerialize["amount_currency"] = o.AmountCurrency
	toSerialize["charge"] = o.Charge
	toSerialize["billable_metric"] = o.BillableMetric
	toSerialize["groups"] = o.Groups
	return toSerialize, nil
}

func (o *CustomerChargeUsageObject) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"units",
		"events_count",
		"amount_cents",
		"amount_currency",
		"charge",
		"billable_metric",
		"groups",
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

	varCustomerChargeUsageObject := _CustomerChargeUsageObject{}

	err = json.Unmarshal(bytes, &varCustomerChargeUsageObject)

	if err != nil {
		return err
	}

	*o = CustomerChargeUsageObject(varCustomerChargeUsageObject)

	return err
}

type NullableCustomerChargeUsageObject struct {
	value *CustomerChargeUsageObject
	isSet bool
}

func (v NullableCustomerChargeUsageObject) Get() *CustomerChargeUsageObject {
	return v.value
}

func (v *NullableCustomerChargeUsageObject) Set(val *CustomerChargeUsageObject) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerChargeUsageObject) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerChargeUsageObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerChargeUsageObject(val *CustomerChargeUsageObject) *NullableCustomerChargeUsageObject {
	return &NullableCustomerChargeUsageObject{value: val, isSet: true}
}

func (v NullableCustomerChargeUsageObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerChargeUsageObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


