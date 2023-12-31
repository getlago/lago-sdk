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

// checks if the AppliedCouponInputAppliedCoupon type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppliedCouponInputAppliedCoupon{}

// AppliedCouponInputAppliedCoupon struct for AppliedCouponInputAppliedCoupon
type AppliedCouponInputAppliedCoupon struct {
	// The customer external unique identifier (provided by your own application)
	ExternalCustomerId string `json:"external_customer_id"`
	// Unique code used to identify the coupon.
	CouponCode string `json:"coupon_code"`
	// The type of frequency for the coupon. It can have three possible values: `once`, `recurring` or `forever`.  - If set to `once`, the coupon is applicable only for a single use. - If set to `recurring`, the coupon can be used multiple times for recurring billing periods. - If set to `forever`, the coupon has unlimited usage and can be applied indefinitely.
	Frequency NullableString `json:"frequency,omitempty"`
	// Specifies the number of billing periods to which the coupon applies. This field is required only for coupons with a `recurring` frequency type
	FrequencyDuration NullableInt32 `json:"frequency_duration,omitempty"`
	// The amount of the coupon in cents. This field is required only for coupon with `fixed_amount` type.
	AmountCents NullableInt32 `json:"amount_cents,omitempty"`
	AmountCurrency *Currency `json:"amount_currency,omitempty"`
	// The percentage rate of the coupon. This field is required only for coupons with a `percentage` coupon type.
	PercentageRate NullableString `json:"percentage_rate,omitempty"`
}

type _AppliedCouponInputAppliedCoupon AppliedCouponInputAppliedCoupon

// NewAppliedCouponInputAppliedCoupon instantiates a new AppliedCouponInputAppliedCoupon object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppliedCouponInputAppliedCoupon(externalCustomerId string, couponCode string) *AppliedCouponInputAppliedCoupon {
	this := AppliedCouponInputAppliedCoupon{}
	this.ExternalCustomerId = externalCustomerId
	this.CouponCode = couponCode
	return &this
}

// NewAppliedCouponInputAppliedCouponWithDefaults instantiates a new AppliedCouponInputAppliedCoupon object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppliedCouponInputAppliedCouponWithDefaults() *AppliedCouponInputAppliedCoupon {
	this := AppliedCouponInputAppliedCoupon{}
	return &this
}

// GetExternalCustomerId returns the ExternalCustomerId field value
func (o *AppliedCouponInputAppliedCoupon) GetExternalCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalCustomerId
}

// GetExternalCustomerIdOk returns a tuple with the ExternalCustomerId field value
// and a boolean to check if the value has been set.
func (o *AppliedCouponInputAppliedCoupon) GetExternalCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalCustomerId, true
}

// SetExternalCustomerId sets field value
func (o *AppliedCouponInputAppliedCoupon) SetExternalCustomerId(v string) {
	o.ExternalCustomerId = v
}

// GetCouponCode returns the CouponCode field value
func (o *AppliedCouponInputAppliedCoupon) GetCouponCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CouponCode
}

// GetCouponCodeOk returns a tuple with the CouponCode field value
// and a boolean to check if the value has been set.
func (o *AppliedCouponInputAppliedCoupon) GetCouponCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CouponCode, true
}

// SetCouponCode sets field value
func (o *AppliedCouponInputAppliedCoupon) SetCouponCode(v string) {
	o.CouponCode = v
}

// GetFrequency returns the Frequency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppliedCouponInputAppliedCoupon) GetFrequency() string {
	if o == nil || IsNil(o.Frequency.Get()) {
		var ret string
		return ret
	}
	return *o.Frequency.Get()
}

// GetFrequencyOk returns a tuple with the Frequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppliedCouponInputAppliedCoupon) GetFrequencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Frequency.Get(), o.Frequency.IsSet()
}

// HasFrequency returns a boolean if a field has been set.
func (o *AppliedCouponInputAppliedCoupon) HasFrequency() bool {
	if o != nil && o.Frequency.IsSet() {
		return true
	}

	return false
}

// SetFrequency gets a reference to the given NullableString and assigns it to the Frequency field.
func (o *AppliedCouponInputAppliedCoupon) SetFrequency(v string) {
	o.Frequency.Set(&v)
}
// SetFrequencyNil sets the value for Frequency to be an explicit nil
func (o *AppliedCouponInputAppliedCoupon) SetFrequencyNil() {
	o.Frequency.Set(nil)
}

// UnsetFrequency ensures that no value is present for Frequency, not even an explicit nil
func (o *AppliedCouponInputAppliedCoupon) UnsetFrequency() {
	o.Frequency.Unset()
}

// GetFrequencyDuration returns the FrequencyDuration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppliedCouponInputAppliedCoupon) GetFrequencyDuration() int32 {
	if o == nil || IsNil(o.FrequencyDuration.Get()) {
		var ret int32
		return ret
	}
	return *o.FrequencyDuration.Get()
}

// GetFrequencyDurationOk returns a tuple with the FrequencyDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppliedCouponInputAppliedCoupon) GetFrequencyDurationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.FrequencyDuration.Get(), o.FrequencyDuration.IsSet()
}

// HasFrequencyDuration returns a boolean if a field has been set.
func (o *AppliedCouponInputAppliedCoupon) HasFrequencyDuration() bool {
	if o != nil && o.FrequencyDuration.IsSet() {
		return true
	}

	return false
}

// SetFrequencyDuration gets a reference to the given NullableInt32 and assigns it to the FrequencyDuration field.
func (o *AppliedCouponInputAppliedCoupon) SetFrequencyDuration(v int32) {
	o.FrequencyDuration.Set(&v)
}
// SetFrequencyDurationNil sets the value for FrequencyDuration to be an explicit nil
func (o *AppliedCouponInputAppliedCoupon) SetFrequencyDurationNil() {
	o.FrequencyDuration.Set(nil)
}

// UnsetFrequencyDuration ensures that no value is present for FrequencyDuration, not even an explicit nil
func (o *AppliedCouponInputAppliedCoupon) UnsetFrequencyDuration() {
	o.FrequencyDuration.Unset()
}

// GetAmountCents returns the AmountCents field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppliedCouponInputAppliedCoupon) GetAmountCents() int32 {
	if o == nil || IsNil(o.AmountCents.Get()) {
		var ret int32
		return ret
	}
	return *o.AmountCents.Get()
}

// GetAmountCentsOk returns a tuple with the AmountCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppliedCouponInputAppliedCoupon) GetAmountCentsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AmountCents.Get(), o.AmountCents.IsSet()
}

// HasAmountCents returns a boolean if a field has been set.
func (o *AppliedCouponInputAppliedCoupon) HasAmountCents() bool {
	if o != nil && o.AmountCents.IsSet() {
		return true
	}

	return false
}

// SetAmountCents gets a reference to the given NullableInt32 and assigns it to the AmountCents field.
func (o *AppliedCouponInputAppliedCoupon) SetAmountCents(v int32) {
	o.AmountCents.Set(&v)
}
// SetAmountCentsNil sets the value for AmountCents to be an explicit nil
func (o *AppliedCouponInputAppliedCoupon) SetAmountCentsNil() {
	o.AmountCents.Set(nil)
}

// UnsetAmountCents ensures that no value is present for AmountCents, not even an explicit nil
func (o *AppliedCouponInputAppliedCoupon) UnsetAmountCents() {
	o.AmountCents.Unset()
}

// GetAmountCurrency returns the AmountCurrency field value if set, zero value otherwise.
func (o *AppliedCouponInputAppliedCoupon) GetAmountCurrency() Currency {
	if o == nil || IsNil(o.AmountCurrency) {
		var ret Currency
		return ret
	}
	return *o.AmountCurrency
}

// GetAmountCurrencyOk returns a tuple with the AmountCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppliedCouponInputAppliedCoupon) GetAmountCurrencyOk() (*Currency, bool) {
	if o == nil || IsNil(o.AmountCurrency) {
		return nil, false
	}
	return o.AmountCurrency, true
}

// HasAmountCurrency returns a boolean if a field has been set.
func (o *AppliedCouponInputAppliedCoupon) HasAmountCurrency() bool {
	if o != nil && !IsNil(o.AmountCurrency) {
		return true
	}

	return false
}

// SetAmountCurrency gets a reference to the given Currency and assigns it to the AmountCurrency field.
func (o *AppliedCouponInputAppliedCoupon) SetAmountCurrency(v Currency) {
	o.AmountCurrency = &v
}

// GetPercentageRate returns the PercentageRate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppliedCouponInputAppliedCoupon) GetPercentageRate() string {
	if o == nil || IsNil(o.PercentageRate.Get()) {
		var ret string
		return ret
	}
	return *o.PercentageRate.Get()
}

// GetPercentageRateOk returns a tuple with the PercentageRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppliedCouponInputAppliedCoupon) GetPercentageRateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PercentageRate.Get(), o.PercentageRate.IsSet()
}

// HasPercentageRate returns a boolean if a field has been set.
func (o *AppliedCouponInputAppliedCoupon) HasPercentageRate() bool {
	if o != nil && o.PercentageRate.IsSet() {
		return true
	}

	return false
}

// SetPercentageRate gets a reference to the given NullableString and assigns it to the PercentageRate field.
func (o *AppliedCouponInputAppliedCoupon) SetPercentageRate(v string) {
	o.PercentageRate.Set(&v)
}
// SetPercentageRateNil sets the value for PercentageRate to be an explicit nil
func (o *AppliedCouponInputAppliedCoupon) SetPercentageRateNil() {
	o.PercentageRate.Set(nil)
}

// UnsetPercentageRate ensures that no value is present for PercentageRate, not even an explicit nil
func (o *AppliedCouponInputAppliedCoupon) UnsetPercentageRate() {
	o.PercentageRate.Unset()
}

func (o AppliedCouponInputAppliedCoupon) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppliedCouponInputAppliedCoupon) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["external_customer_id"] = o.ExternalCustomerId
	toSerialize["coupon_code"] = o.CouponCode
	if o.Frequency.IsSet() {
		toSerialize["frequency"] = o.Frequency.Get()
	}
	if o.FrequencyDuration.IsSet() {
		toSerialize["frequency_duration"] = o.FrequencyDuration.Get()
	}
	if o.AmountCents.IsSet() {
		toSerialize["amount_cents"] = o.AmountCents.Get()
	}
	if !IsNil(o.AmountCurrency) {
		toSerialize["amount_currency"] = o.AmountCurrency
	}
	if o.PercentageRate.IsSet() {
		toSerialize["percentage_rate"] = o.PercentageRate.Get()
	}
	return toSerialize, nil
}

func (o *AppliedCouponInputAppliedCoupon) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"external_customer_id",
		"coupon_code",
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

	varAppliedCouponInputAppliedCoupon := _AppliedCouponInputAppliedCoupon{}

	err = json.Unmarshal(bytes, &varAppliedCouponInputAppliedCoupon)

	if err != nil {
		return err
	}

	*o = AppliedCouponInputAppliedCoupon(varAppliedCouponInputAppliedCoupon)

	return err
}

type NullableAppliedCouponInputAppliedCoupon struct {
	value *AppliedCouponInputAppliedCoupon
	isSet bool
}

func (v NullableAppliedCouponInputAppliedCoupon) Get() *AppliedCouponInputAppliedCoupon {
	return v.value
}

func (v *NullableAppliedCouponInputAppliedCoupon) Set(val *AppliedCouponInputAppliedCoupon) {
	v.value = val
	v.isSet = true
}

func (v NullableAppliedCouponInputAppliedCoupon) IsSet() bool {
	return v.isSet
}

func (v *NullableAppliedCouponInputAppliedCoupon) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppliedCouponInputAppliedCoupon(val *AppliedCouponInputAppliedCoupon) *NullableAppliedCouponInputAppliedCoupon {
	return &NullableAppliedCouponInputAppliedCoupon{value: val, isSet: true}
}

func (v NullableAppliedCouponInputAppliedCoupon) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppliedCouponInputAppliedCoupon) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


