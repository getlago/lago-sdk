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

// checks if the AppliedCouponObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppliedCouponObject{}

// AppliedCouponObject struct for AppliedCouponObject
type AppliedCouponObject struct {
	// Unique identifier of the applied coupon, created by Lago.
	LagoId string `json:"lago_id"`
	// Unique identifier of the coupon, created by Lago.
	LagoCouponId string `json:"lago_coupon_id"`
	// Unique code used to identify the coupon.
	CouponCode string `json:"coupon_code"`
	// The name of the coupon.
	CouponName string `json:"coupon_name"`
	// Unique identifier of the customer, created by Lago.
	LagoCustomerId string `json:"lago_customer_id"`
	// The customer external unique identifier (provided by your own application)
	ExternalCustomerId string `json:"external_customer_id"`
	// The status of the coupon. Can be either `active` or `terminated`.
	Status string `json:"status"`
	// The amount of the coupon in cents. This field is required only for coupon with `fixed_amount` type.
	AmountCents NullableInt32 `json:"amount_cents,omitempty"`
	// The remaining amount in cents for a `fixed_amount` coupon with a frequency set to `once`. This field indicates the remaining balance or value that can still be utilized from the coupon.
	AmountCentsRemaining NullableInt32 `json:"amount_cents_remaining,omitempty"`
	AmountCurrency *Currency `json:"amount_currency,omitempty"`
	// The percentage rate of the coupon. This field is required only for coupons with a `percentage` coupon type.
	PercentageRate NullableString `json:"percentage_rate,omitempty"`
	// The type of frequency for the coupon. It can have three possible values: `once`, `recurring` or `forever`.  - If set to `once`, the coupon is applicable only for a single use. - If set to `recurring`, the coupon can be used multiple times for recurring billing periods. - If set to `forever`, the coupon has unlimited usage and can be applied indefinitely.
	Frequency string `json:"frequency"`
	// Specifies the number of billing periods to which the coupon applies. This field is required only for coupons with a `recurring` frequency type
	FrequencyDuration NullableInt32 `json:"frequency_duration,omitempty"`
	// The remaining number of billing periods to which the coupon is applicable. This field determines the remaining usage or availability of the coupon based on the remaining billing periods.
	FrequencyDurationRemaining NullableInt32 `json:"frequency_duration_remaining,omitempty"`
	// The date and time after which the coupon will stop applying to customer's invoices. Once the expiration date is reached, the coupon will no longer be applicable, and any further invoices generated for the customer will not include the coupon discount.
	ExpirationAt NullableTime `json:"expiration_at,omitempty"`
	// The date and time when the coupon was assigned to a customer. It is expressed in UTC format according to the ISO 8601 datetime standard.
	CreatedAt time.Time `json:"created_at"`
	// This field indicates the specific moment when the coupon amount is fully utilized or when the coupon is removed from the customer's coupon list. It is expressed in UTC format according to the ISO 8601 datetime standard.
	TerminatedAt NullableTime `json:"terminated_at,omitempty"`
}

type _AppliedCouponObject AppliedCouponObject

// NewAppliedCouponObject instantiates a new AppliedCouponObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppliedCouponObject(lagoId string, lagoCouponId string, couponCode string, couponName string, lagoCustomerId string, externalCustomerId string, status string, frequency string, createdAt time.Time) *AppliedCouponObject {
	this := AppliedCouponObject{}
	this.LagoId = lagoId
	this.LagoCouponId = lagoCouponId
	this.CouponCode = couponCode
	this.CouponName = couponName
	this.LagoCustomerId = lagoCustomerId
	this.ExternalCustomerId = externalCustomerId
	this.Status = status
	this.Frequency = frequency
	this.CreatedAt = createdAt
	return &this
}

// NewAppliedCouponObjectWithDefaults instantiates a new AppliedCouponObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppliedCouponObjectWithDefaults() *AppliedCouponObject {
	this := AppliedCouponObject{}
	return &this
}

// GetLagoId returns the LagoId field value
func (o *AppliedCouponObject) GetLagoId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LagoId
}

// GetLagoIdOk returns a tuple with the LagoId field value
// and a boolean to check if the value has been set.
func (o *AppliedCouponObject) GetLagoIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LagoId, true
}

// SetLagoId sets field value
func (o *AppliedCouponObject) SetLagoId(v string) {
	o.LagoId = v
}

// GetLagoCouponId returns the LagoCouponId field value
func (o *AppliedCouponObject) GetLagoCouponId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LagoCouponId
}

// GetLagoCouponIdOk returns a tuple with the LagoCouponId field value
// and a boolean to check if the value has been set.
func (o *AppliedCouponObject) GetLagoCouponIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LagoCouponId, true
}

// SetLagoCouponId sets field value
func (o *AppliedCouponObject) SetLagoCouponId(v string) {
	o.LagoCouponId = v
}

// GetCouponCode returns the CouponCode field value
func (o *AppliedCouponObject) GetCouponCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CouponCode
}

// GetCouponCodeOk returns a tuple with the CouponCode field value
// and a boolean to check if the value has been set.
func (o *AppliedCouponObject) GetCouponCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CouponCode, true
}

// SetCouponCode sets field value
func (o *AppliedCouponObject) SetCouponCode(v string) {
	o.CouponCode = v
}

// GetCouponName returns the CouponName field value
func (o *AppliedCouponObject) GetCouponName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CouponName
}

// GetCouponNameOk returns a tuple with the CouponName field value
// and a boolean to check if the value has been set.
func (o *AppliedCouponObject) GetCouponNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CouponName, true
}

// SetCouponName sets field value
func (o *AppliedCouponObject) SetCouponName(v string) {
	o.CouponName = v
}

// GetLagoCustomerId returns the LagoCustomerId field value
func (o *AppliedCouponObject) GetLagoCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LagoCustomerId
}

// GetLagoCustomerIdOk returns a tuple with the LagoCustomerId field value
// and a boolean to check if the value has been set.
func (o *AppliedCouponObject) GetLagoCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LagoCustomerId, true
}

// SetLagoCustomerId sets field value
func (o *AppliedCouponObject) SetLagoCustomerId(v string) {
	o.LagoCustomerId = v
}

// GetExternalCustomerId returns the ExternalCustomerId field value
func (o *AppliedCouponObject) GetExternalCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalCustomerId
}

// GetExternalCustomerIdOk returns a tuple with the ExternalCustomerId field value
// and a boolean to check if the value has been set.
func (o *AppliedCouponObject) GetExternalCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalCustomerId, true
}

// SetExternalCustomerId sets field value
func (o *AppliedCouponObject) SetExternalCustomerId(v string) {
	o.ExternalCustomerId = v
}

// GetStatus returns the Status field value
func (o *AppliedCouponObject) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *AppliedCouponObject) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *AppliedCouponObject) SetStatus(v string) {
	o.Status = v
}

// GetAmountCents returns the AmountCents field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppliedCouponObject) GetAmountCents() int32 {
	if o == nil || IsNil(o.AmountCents.Get()) {
		var ret int32
		return ret
	}
	return *o.AmountCents.Get()
}

// GetAmountCentsOk returns a tuple with the AmountCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppliedCouponObject) GetAmountCentsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AmountCents.Get(), o.AmountCents.IsSet()
}

// HasAmountCents returns a boolean if a field has been set.
func (o *AppliedCouponObject) HasAmountCents() bool {
	if o != nil && o.AmountCents.IsSet() {
		return true
	}

	return false
}

// SetAmountCents gets a reference to the given NullableInt32 and assigns it to the AmountCents field.
func (o *AppliedCouponObject) SetAmountCents(v int32) {
	o.AmountCents.Set(&v)
}
// SetAmountCentsNil sets the value for AmountCents to be an explicit nil
func (o *AppliedCouponObject) SetAmountCentsNil() {
	o.AmountCents.Set(nil)
}

// UnsetAmountCents ensures that no value is present for AmountCents, not even an explicit nil
func (o *AppliedCouponObject) UnsetAmountCents() {
	o.AmountCents.Unset()
}

// GetAmountCentsRemaining returns the AmountCentsRemaining field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppliedCouponObject) GetAmountCentsRemaining() int32 {
	if o == nil || IsNil(o.AmountCentsRemaining.Get()) {
		var ret int32
		return ret
	}
	return *o.AmountCentsRemaining.Get()
}

// GetAmountCentsRemainingOk returns a tuple with the AmountCentsRemaining field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppliedCouponObject) GetAmountCentsRemainingOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AmountCentsRemaining.Get(), o.AmountCentsRemaining.IsSet()
}

// HasAmountCentsRemaining returns a boolean if a field has been set.
func (o *AppliedCouponObject) HasAmountCentsRemaining() bool {
	if o != nil && o.AmountCentsRemaining.IsSet() {
		return true
	}

	return false
}

// SetAmountCentsRemaining gets a reference to the given NullableInt32 and assigns it to the AmountCentsRemaining field.
func (o *AppliedCouponObject) SetAmountCentsRemaining(v int32) {
	o.AmountCentsRemaining.Set(&v)
}
// SetAmountCentsRemainingNil sets the value for AmountCentsRemaining to be an explicit nil
func (o *AppliedCouponObject) SetAmountCentsRemainingNil() {
	o.AmountCentsRemaining.Set(nil)
}

// UnsetAmountCentsRemaining ensures that no value is present for AmountCentsRemaining, not even an explicit nil
func (o *AppliedCouponObject) UnsetAmountCentsRemaining() {
	o.AmountCentsRemaining.Unset()
}

// GetAmountCurrency returns the AmountCurrency field value if set, zero value otherwise.
func (o *AppliedCouponObject) GetAmountCurrency() Currency {
	if o == nil || IsNil(o.AmountCurrency) {
		var ret Currency
		return ret
	}
	return *o.AmountCurrency
}

// GetAmountCurrencyOk returns a tuple with the AmountCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppliedCouponObject) GetAmountCurrencyOk() (*Currency, bool) {
	if o == nil || IsNil(o.AmountCurrency) {
		return nil, false
	}
	return o.AmountCurrency, true
}

// HasAmountCurrency returns a boolean if a field has been set.
func (o *AppliedCouponObject) HasAmountCurrency() bool {
	if o != nil && !IsNil(o.AmountCurrency) {
		return true
	}

	return false
}

// SetAmountCurrency gets a reference to the given Currency and assigns it to the AmountCurrency field.
func (o *AppliedCouponObject) SetAmountCurrency(v Currency) {
	o.AmountCurrency = &v
}

// GetPercentageRate returns the PercentageRate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppliedCouponObject) GetPercentageRate() string {
	if o == nil || IsNil(o.PercentageRate.Get()) {
		var ret string
		return ret
	}
	return *o.PercentageRate.Get()
}

// GetPercentageRateOk returns a tuple with the PercentageRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppliedCouponObject) GetPercentageRateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PercentageRate.Get(), o.PercentageRate.IsSet()
}

// HasPercentageRate returns a boolean if a field has been set.
func (o *AppliedCouponObject) HasPercentageRate() bool {
	if o != nil && o.PercentageRate.IsSet() {
		return true
	}

	return false
}

// SetPercentageRate gets a reference to the given NullableString and assigns it to the PercentageRate field.
func (o *AppliedCouponObject) SetPercentageRate(v string) {
	o.PercentageRate.Set(&v)
}
// SetPercentageRateNil sets the value for PercentageRate to be an explicit nil
func (o *AppliedCouponObject) SetPercentageRateNil() {
	o.PercentageRate.Set(nil)
}

// UnsetPercentageRate ensures that no value is present for PercentageRate, not even an explicit nil
func (o *AppliedCouponObject) UnsetPercentageRate() {
	o.PercentageRate.Unset()
}

// GetFrequency returns the Frequency field value
func (o *AppliedCouponObject) GetFrequency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field value
// and a boolean to check if the value has been set.
func (o *AppliedCouponObject) GetFrequencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Frequency, true
}

// SetFrequency sets field value
func (o *AppliedCouponObject) SetFrequency(v string) {
	o.Frequency = v
}

// GetFrequencyDuration returns the FrequencyDuration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppliedCouponObject) GetFrequencyDuration() int32 {
	if o == nil || IsNil(o.FrequencyDuration.Get()) {
		var ret int32
		return ret
	}
	return *o.FrequencyDuration.Get()
}

// GetFrequencyDurationOk returns a tuple with the FrequencyDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppliedCouponObject) GetFrequencyDurationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.FrequencyDuration.Get(), o.FrequencyDuration.IsSet()
}

// HasFrequencyDuration returns a boolean if a field has been set.
func (o *AppliedCouponObject) HasFrequencyDuration() bool {
	if o != nil && o.FrequencyDuration.IsSet() {
		return true
	}

	return false
}

// SetFrequencyDuration gets a reference to the given NullableInt32 and assigns it to the FrequencyDuration field.
func (o *AppliedCouponObject) SetFrequencyDuration(v int32) {
	o.FrequencyDuration.Set(&v)
}
// SetFrequencyDurationNil sets the value for FrequencyDuration to be an explicit nil
func (o *AppliedCouponObject) SetFrequencyDurationNil() {
	o.FrequencyDuration.Set(nil)
}

// UnsetFrequencyDuration ensures that no value is present for FrequencyDuration, not even an explicit nil
func (o *AppliedCouponObject) UnsetFrequencyDuration() {
	o.FrequencyDuration.Unset()
}

// GetFrequencyDurationRemaining returns the FrequencyDurationRemaining field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppliedCouponObject) GetFrequencyDurationRemaining() int32 {
	if o == nil || IsNil(o.FrequencyDurationRemaining.Get()) {
		var ret int32
		return ret
	}
	return *o.FrequencyDurationRemaining.Get()
}

// GetFrequencyDurationRemainingOk returns a tuple with the FrequencyDurationRemaining field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppliedCouponObject) GetFrequencyDurationRemainingOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.FrequencyDurationRemaining.Get(), o.FrequencyDurationRemaining.IsSet()
}

// HasFrequencyDurationRemaining returns a boolean if a field has been set.
func (o *AppliedCouponObject) HasFrequencyDurationRemaining() bool {
	if o != nil && o.FrequencyDurationRemaining.IsSet() {
		return true
	}

	return false
}

// SetFrequencyDurationRemaining gets a reference to the given NullableInt32 and assigns it to the FrequencyDurationRemaining field.
func (o *AppliedCouponObject) SetFrequencyDurationRemaining(v int32) {
	o.FrequencyDurationRemaining.Set(&v)
}
// SetFrequencyDurationRemainingNil sets the value for FrequencyDurationRemaining to be an explicit nil
func (o *AppliedCouponObject) SetFrequencyDurationRemainingNil() {
	o.FrequencyDurationRemaining.Set(nil)
}

// UnsetFrequencyDurationRemaining ensures that no value is present for FrequencyDurationRemaining, not even an explicit nil
func (o *AppliedCouponObject) UnsetFrequencyDurationRemaining() {
	o.FrequencyDurationRemaining.Unset()
}

// GetExpirationAt returns the ExpirationAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppliedCouponObject) GetExpirationAt() time.Time {
	if o == nil || IsNil(o.ExpirationAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ExpirationAt.Get()
}

// GetExpirationAtOk returns a tuple with the ExpirationAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppliedCouponObject) GetExpirationAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpirationAt.Get(), o.ExpirationAt.IsSet()
}

// HasExpirationAt returns a boolean if a field has been set.
func (o *AppliedCouponObject) HasExpirationAt() bool {
	if o != nil && o.ExpirationAt.IsSet() {
		return true
	}

	return false
}

// SetExpirationAt gets a reference to the given NullableTime and assigns it to the ExpirationAt field.
func (o *AppliedCouponObject) SetExpirationAt(v time.Time) {
	o.ExpirationAt.Set(&v)
}
// SetExpirationAtNil sets the value for ExpirationAt to be an explicit nil
func (o *AppliedCouponObject) SetExpirationAtNil() {
	o.ExpirationAt.Set(nil)
}

// UnsetExpirationAt ensures that no value is present for ExpirationAt, not even an explicit nil
func (o *AppliedCouponObject) UnsetExpirationAt() {
	o.ExpirationAt.Unset()
}

// GetCreatedAt returns the CreatedAt field value
func (o *AppliedCouponObject) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AppliedCouponObject) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *AppliedCouponObject) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetTerminatedAt returns the TerminatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppliedCouponObject) GetTerminatedAt() time.Time {
	if o == nil || IsNil(o.TerminatedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.TerminatedAt.Get()
}

// GetTerminatedAtOk returns a tuple with the TerminatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppliedCouponObject) GetTerminatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.TerminatedAt.Get(), o.TerminatedAt.IsSet()
}

// HasTerminatedAt returns a boolean if a field has been set.
func (o *AppliedCouponObject) HasTerminatedAt() bool {
	if o != nil && o.TerminatedAt.IsSet() {
		return true
	}

	return false
}

// SetTerminatedAt gets a reference to the given NullableTime and assigns it to the TerminatedAt field.
func (o *AppliedCouponObject) SetTerminatedAt(v time.Time) {
	o.TerminatedAt.Set(&v)
}
// SetTerminatedAtNil sets the value for TerminatedAt to be an explicit nil
func (o *AppliedCouponObject) SetTerminatedAtNil() {
	o.TerminatedAt.Set(nil)
}

// UnsetTerminatedAt ensures that no value is present for TerminatedAt, not even an explicit nil
func (o *AppliedCouponObject) UnsetTerminatedAt() {
	o.TerminatedAt.Unset()
}

func (o AppliedCouponObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppliedCouponObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["lago_id"] = o.LagoId
	toSerialize["lago_coupon_id"] = o.LagoCouponId
	toSerialize["coupon_code"] = o.CouponCode
	toSerialize["coupon_name"] = o.CouponName
	toSerialize["lago_customer_id"] = o.LagoCustomerId
	toSerialize["external_customer_id"] = o.ExternalCustomerId
	toSerialize["status"] = o.Status
	if o.AmountCents.IsSet() {
		toSerialize["amount_cents"] = o.AmountCents.Get()
	}
	if o.AmountCentsRemaining.IsSet() {
		toSerialize["amount_cents_remaining"] = o.AmountCentsRemaining.Get()
	}
	if !IsNil(o.AmountCurrency) {
		toSerialize["amount_currency"] = o.AmountCurrency
	}
	if o.PercentageRate.IsSet() {
		toSerialize["percentage_rate"] = o.PercentageRate.Get()
	}
	toSerialize["frequency"] = o.Frequency
	if o.FrequencyDuration.IsSet() {
		toSerialize["frequency_duration"] = o.FrequencyDuration.Get()
	}
	if o.FrequencyDurationRemaining.IsSet() {
		toSerialize["frequency_duration_remaining"] = o.FrequencyDurationRemaining.Get()
	}
	if o.ExpirationAt.IsSet() {
		toSerialize["expiration_at"] = o.ExpirationAt.Get()
	}
	toSerialize["created_at"] = o.CreatedAt
	if o.TerminatedAt.IsSet() {
		toSerialize["terminated_at"] = o.TerminatedAt.Get()
	}
	return toSerialize, nil
}

func (o *AppliedCouponObject) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"lago_id",
		"lago_coupon_id",
		"coupon_code",
		"coupon_name",
		"lago_customer_id",
		"external_customer_id",
		"status",
		"frequency",
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

	varAppliedCouponObject := _AppliedCouponObject{}

	err = json.Unmarshal(bytes, &varAppliedCouponObject)

	if err != nil {
		return err
	}

	*o = AppliedCouponObject(varAppliedCouponObject)

	return err
}

type NullableAppliedCouponObject struct {
	value *AppliedCouponObject
	isSet bool
}

func (v NullableAppliedCouponObject) Get() *AppliedCouponObject {
	return v.value
}

func (v *NullableAppliedCouponObject) Set(val *AppliedCouponObject) {
	v.value = val
	v.isSet = true
}

func (v NullableAppliedCouponObject) IsSet() bool {
	return v.isSet
}

func (v *NullableAppliedCouponObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppliedCouponObject(val *AppliedCouponObject) *NullableAppliedCouponObject {
	return &NullableAppliedCouponObject{value: val, isSet: true}
}

func (v NullableAppliedCouponObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppliedCouponObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


