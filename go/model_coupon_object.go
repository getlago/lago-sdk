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

// checks if the CouponObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CouponObject{}

// CouponObject struct for CouponObject
type CouponObject struct {
	// Unique identifier of the coupon, created by Lago.
	LagoId string `json:"lago_id"`
	// The name of the coupon.
	Name string `json:"name"`
	// Unique code used to identify the coupon.
	Code string `json:"code"`
	// Description of the coupon.
	Description NullableString `json:"description,omitempty"`
	// The type of the coupon. It can have two possible values: `fixed_amount` or `percentage`.  - If set to `fixed_amount`, the coupon represents a fixed amount discount. - If set to `percentage`, the coupon represents a percentage-based discount.
	CouponType string `json:"coupon_type"`
	// The amount of the coupon in cents. This field is required only for coupon with `fixed_amount` type.
	AmountCents NullableInt32 `json:"amount_cents,omitempty"`
	AmountCurrency *Currency `json:"amount_currency,omitempty"`
	// Indicates whether the coupon can be reused or not. If set to `true`, the coupon is reusable, meaning it can be applied multiple times to the same customer. If set to `false`, the coupon can only be used once and is not reusable. If not specified, this field is set to `true` by default.
	Reusable bool `json:"reusable"`
	// The coupon is limited to specific plans. The possible values can be `true` or `false`.
	LimitedPlans bool `json:"limited_plans"`
	// An array of plan codes to which the coupon is applicable. By specifying the plan codes in this field, you can restrict the coupon's usage to specific plans only.
	PlanCodes []string `json:"plan_codes,omitempty"`
	// The coupon is limited to specific billable metrics. The possible values can be `true` or `false`.
	LimitedBillableMetrics bool `json:"limited_billable_metrics"`
	// An array of billable metric codes to which the coupon is applicable. By specifying the billable metric codes in this field, you can restrict the coupon's usage to specific metrics only.
	BillableMetricCodes []string `json:"billable_metric_codes,omitempty"`
	// The percentage rate of the coupon. This field is required only for coupons with a `percentage` coupon type.
	PercentageRate NullableString `json:"percentage_rate,omitempty"`
	// The type of frequency for the coupon. It can have three possible values: `once`, `recurring`, or `forever`.  - If set to `once`, the coupon is applicable only for a single use. - If set to `recurring`, the coupon can be used multiple times for recurring billing periods. - If set to `forever`, the coupon has unlimited usage and can be applied indefinitely.
	Frequency string `json:"frequency"`
	// Specifies the number of billing periods to which the coupon applies. This field is required only for coupons with a `recurring` frequency type
	FrequencyDuration NullableInt32 `json:"frequency_duration,omitempty"`
	// Specifies the type of expiration for the coupon. It can have two possible values: `time_limit` or `no_expiration`.  - If set to `time_limit`, the coupon has an expiration based on a specified time limit. - If set to `no_expiration`, the coupon does not have an expiration date and remains valid indefinitely.
	Expiration string `json:"expiration"`
	// The expiration date and time of the coupon. This field is required only for coupons with `expiration` set to `time_limit`. The expiration date and time should be specified in UTC format according to the ISO 8601 datetime standard. It indicates the exact moment when the coupon will expire and is no longer valid.
	ExpirationAt NullableTime `json:"expiration_at,omitempty"`
	// The date and time when the coupon was created. It is expressed in UTC format according to the ISO 8601 datetime standard. This field provides the timestamp for the exact moment when the coupon was initially created.
	CreatedAt time.Time `json:"created_at"`
	// This field indicates if the coupon has been terminated and is no longer usable. If it's not null, it won't be removed for existing customers using it, but it prevents the coupon from being applied in the future.
	TerminatedAt *time.Time `json:"terminated_at,omitempty"`
}

type _CouponObject CouponObject

// NewCouponObject instantiates a new CouponObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCouponObject(lagoId string, name string, code string, couponType string, reusable bool, limitedPlans bool, limitedBillableMetrics bool, frequency string, expiration string, createdAt time.Time) *CouponObject {
	this := CouponObject{}
	this.LagoId = lagoId
	this.Name = name
	this.Code = code
	this.CouponType = couponType
	this.Reusable = reusable
	this.LimitedPlans = limitedPlans
	this.LimitedBillableMetrics = limitedBillableMetrics
	this.Frequency = frequency
	this.Expiration = expiration
	this.CreatedAt = createdAt
	return &this
}

// NewCouponObjectWithDefaults instantiates a new CouponObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCouponObjectWithDefaults() *CouponObject {
	this := CouponObject{}
	return &this
}

// GetLagoId returns the LagoId field value
func (o *CouponObject) GetLagoId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LagoId
}

// GetLagoIdOk returns a tuple with the LagoId field value
// and a boolean to check if the value has been set.
func (o *CouponObject) GetLagoIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LagoId, true
}

// SetLagoId sets field value
func (o *CouponObject) SetLagoId(v string) {
	o.LagoId = v
}

// GetName returns the Name field value
func (o *CouponObject) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CouponObject) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CouponObject) SetName(v string) {
	o.Name = v
}

// GetCode returns the Code field value
func (o *CouponObject) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *CouponObject) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *CouponObject) SetCode(v string) {
	o.Code = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CouponObject) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CouponObject) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *CouponObject) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *CouponObject) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *CouponObject) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *CouponObject) UnsetDescription() {
	o.Description.Unset()
}

// GetCouponType returns the CouponType field value
func (o *CouponObject) GetCouponType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CouponType
}

// GetCouponTypeOk returns a tuple with the CouponType field value
// and a boolean to check if the value has been set.
func (o *CouponObject) GetCouponTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CouponType, true
}

// SetCouponType sets field value
func (o *CouponObject) SetCouponType(v string) {
	o.CouponType = v
}

// GetAmountCents returns the AmountCents field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CouponObject) GetAmountCents() int32 {
	if o == nil || IsNil(o.AmountCents.Get()) {
		var ret int32
		return ret
	}
	return *o.AmountCents.Get()
}

// GetAmountCentsOk returns a tuple with the AmountCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CouponObject) GetAmountCentsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AmountCents.Get(), o.AmountCents.IsSet()
}

// HasAmountCents returns a boolean if a field has been set.
func (o *CouponObject) HasAmountCents() bool {
	if o != nil && o.AmountCents.IsSet() {
		return true
	}

	return false
}

// SetAmountCents gets a reference to the given NullableInt32 and assigns it to the AmountCents field.
func (o *CouponObject) SetAmountCents(v int32) {
	o.AmountCents.Set(&v)
}
// SetAmountCentsNil sets the value for AmountCents to be an explicit nil
func (o *CouponObject) SetAmountCentsNil() {
	o.AmountCents.Set(nil)
}

// UnsetAmountCents ensures that no value is present for AmountCents, not even an explicit nil
func (o *CouponObject) UnsetAmountCents() {
	o.AmountCents.Unset()
}

// GetAmountCurrency returns the AmountCurrency field value if set, zero value otherwise.
func (o *CouponObject) GetAmountCurrency() Currency {
	if o == nil || IsNil(o.AmountCurrency) {
		var ret Currency
		return ret
	}
	return *o.AmountCurrency
}

// GetAmountCurrencyOk returns a tuple with the AmountCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CouponObject) GetAmountCurrencyOk() (*Currency, bool) {
	if o == nil || IsNil(o.AmountCurrency) {
		return nil, false
	}
	return o.AmountCurrency, true
}

// HasAmountCurrency returns a boolean if a field has been set.
func (o *CouponObject) HasAmountCurrency() bool {
	if o != nil && !IsNil(o.AmountCurrency) {
		return true
	}

	return false
}

// SetAmountCurrency gets a reference to the given Currency and assigns it to the AmountCurrency field.
func (o *CouponObject) SetAmountCurrency(v Currency) {
	o.AmountCurrency = &v
}

// GetReusable returns the Reusable field value
func (o *CouponObject) GetReusable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Reusable
}

// GetReusableOk returns a tuple with the Reusable field value
// and a boolean to check if the value has been set.
func (o *CouponObject) GetReusableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reusable, true
}

// SetReusable sets field value
func (o *CouponObject) SetReusable(v bool) {
	o.Reusable = v
}

// GetLimitedPlans returns the LimitedPlans field value
func (o *CouponObject) GetLimitedPlans() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.LimitedPlans
}

// GetLimitedPlansOk returns a tuple with the LimitedPlans field value
// and a boolean to check if the value has been set.
func (o *CouponObject) GetLimitedPlansOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LimitedPlans, true
}

// SetLimitedPlans sets field value
func (o *CouponObject) SetLimitedPlans(v bool) {
	o.LimitedPlans = v
}

// GetPlanCodes returns the PlanCodes field value if set, zero value otherwise.
func (o *CouponObject) GetPlanCodes() []string {
	if o == nil || IsNil(o.PlanCodes) {
		var ret []string
		return ret
	}
	return o.PlanCodes
}

// GetPlanCodesOk returns a tuple with the PlanCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CouponObject) GetPlanCodesOk() ([]string, bool) {
	if o == nil || IsNil(o.PlanCodes) {
		return nil, false
	}
	return o.PlanCodes, true
}

// HasPlanCodes returns a boolean if a field has been set.
func (o *CouponObject) HasPlanCodes() bool {
	if o != nil && !IsNil(o.PlanCodes) {
		return true
	}

	return false
}

// SetPlanCodes gets a reference to the given []string and assigns it to the PlanCodes field.
func (o *CouponObject) SetPlanCodes(v []string) {
	o.PlanCodes = v
}

// GetLimitedBillableMetrics returns the LimitedBillableMetrics field value
func (o *CouponObject) GetLimitedBillableMetrics() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.LimitedBillableMetrics
}

// GetLimitedBillableMetricsOk returns a tuple with the LimitedBillableMetrics field value
// and a boolean to check if the value has been set.
func (o *CouponObject) GetLimitedBillableMetricsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LimitedBillableMetrics, true
}

// SetLimitedBillableMetrics sets field value
func (o *CouponObject) SetLimitedBillableMetrics(v bool) {
	o.LimitedBillableMetrics = v
}

// GetBillableMetricCodes returns the BillableMetricCodes field value if set, zero value otherwise.
func (o *CouponObject) GetBillableMetricCodes() []string {
	if o == nil || IsNil(o.BillableMetricCodes) {
		var ret []string
		return ret
	}
	return o.BillableMetricCodes
}

// GetBillableMetricCodesOk returns a tuple with the BillableMetricCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CouponObject) GetBillableMetricCodesOk() ([]string, bool) {
	if o == nil || IsNil(o.BillableMetricCodes) {
		return nil, false
	}
	return o.BillableMetricCodes, true
}

// HasBillableMetricCodes returns a boolean if a field has been set.
func (o *CouponObject) HasBillableMetricCodes() bool {
	if o != nil && !IsNil(o.BillableMetricCodes) {
		return true
	}

	return false
}

// SetBillableMetricCodes gets a reference to the given []string and assigns it to the BillableMetricCodes field.
func (o *CouponObject) SetBillableMetricCodes(v []string) {
	o.BillableMetricCodes = v
}

// GetPercentageRate returns the PercentageRate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CouponObject) GetPercentageRate() string {
	if o == nil || IsNil(o.PercentageRate.Get()) {
		var ret string
		return ret
	}
	return *o.PercentageRate.Get()
}

// GetPercentageRateOk returns a tuple with the PercentageRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CouponObject) GetPercentageRateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PercentageRate.Get(), o.PercentageRate.IsSet()
}

// HasPercentageRate returns a boolean if a field has been set.
func (o *CouponObject) HasPercentageRate() bool {
	if o != nil && o.PercentageRate.IsSet() {
		return true
	}

	return false
}

// SetPercentageRate gets a reference to the given NullableString and assigns it to the PercentageRate field.
func (o *CouponObject) SetPercentageRate(v string) {
	o.PercentageRate.Set(&v)
}
// SetPercentageRateNil sets the value for PercentageRate to be an explicit nil
func (o *CouponObject) SetPercentageRateNil() {
	o.PercentageRate.Set(nil)
}

// UnsetPercentageRate ensures that no value is present for PercentageRate, not even an explicit nil
func (o *CouponObject) UnsetPercentageRate() {
	o.PercentageRate.Unset()
}

// GetFrequency returns the Frequency field value
func (o *CouponObject) GetFrequency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field value
// and a boolean to check if the value has been set.
func (o *CouponObject) GetFrequencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Frequency, true
}

// SetFrequency sets field value
func (o *CouponObject) SetFrequency(v string) {
	o.Frequency = v
}

// GetFrequencyDuration returns the FrequencyDuration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CouponObject) GetFrequencyDuration() int32 {
	if o == nil || IsNil(o.FrequencyDuration.Get()) {
		var ret int32
		return ret
	}
	return *o.FrequencyDuration.Get()
}

// GetFrequencyDurationOk returns a tuple with the FrequencyDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CouponObject) GetFrequencyDurationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.FrequencyDuration.Get(), o.FrequencyDuration.IsSet()
}

// HasFrequencyDuration returns a boolean if a field has been set.
func (o *CouponObject) HasFrequencyDuration() bool {
	if o != nil && o.FrequencyDuration.IsSet() {
		return true
	}

	return false
}

// SetFrequencyDuration gets a reference to the given NullableInt32 and assigns it to the FrequencyDuration field.
func (o *CouponObject) SetFrequencyDuration(v int32) {
	o.FrequencyDuration.Set(&v)
}
// SetFrequencyDurationNil sets the value for FrequencyDuration to be an explicit nil
func (o *CouponObject) SetFrequencyDurationNil() {
	o.FrequencyDuration.Set(nil)
}

// UnsetFrequencyDuration ensures that no value is present for FrequencyDuration, not even an explicit nil
func (o *CouponObject) UnsetFrequencyDuration() {
	o.FrequencyDuration.Unset()
}

// GetExpiration returns the Expiration field value
func (o *CouponObject) GetExpiration() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value
// and a boolean to check if the value has been set.
func (o *CouponObject) GetExpirationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expiration, true
}

// SetExpiration sets field value
func (o *CouponObject) SetExpiration(v string) {
	o.Expiration = v
}

// GetExpirationAt returns the ExpirationAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CouponObject) GetExpirationAt() time.Time {
	if o == nil || IsNil(o.ExpirationAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ExpirationAt.Get()
}

// GetExpirationAtOk returns a tuple with the ExpirationAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CouponObject) GetExpirationAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpirationAt.Get(), o.ExpirationAt.IsSet()
}

// HasExpirationAt returns a boolean if a field has been set.
func (o *CouponObject) HasExpirationAt() bool {
	if o != nil && o.ExpirationAt.IsSet() {
		return true
	}

	return false
}

// SetExpirationAt gets a reference to the given NullableTime and assigns it to the ExpirationAt field.
func (o *CouponObject) SetExpirationAt(v time.Time) {
	o.ExpirationAt.Set(&v)
}
// SetExpirationAtNil sets the value for ExpirationAt to be an explicit nil
func (o *CouponObject) SetExpirationAtNil() {
	o.ExpirationAt.Set(nil)
}

// UnsetExpirationAt ensures that no value is present for ExpirationAt, not even an explicit nil
func (o *CouponObject) UnsetExpirationAt() {
	o.ExpirationAt.Unset()
}

// GetCreatedAt returns the CreatedAt field value
func (o *CouponObject) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *CouponObject) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *CouponObject) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetTerminatedAt returns the TerminatedAt field value if set, zero value otherwise.
func (o *CouponObject) GetTerminatedAt() time.Time {
	if o == nil || IsNil(o.TerminatedAt) {
		var ret time.Time
		return ret
	}
	return *o.TerminatedAt
}

// GetTerminatedAtOk returns a tuple with the TerminatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CouponObject) GetTerminatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.TerminatedAt) {
		return nil, false
	}
	return o.TerminatedAt, true
}

// HasTerminatedAt returns a boolean if a field has been set.
func (o *CouponObject) HasTerminatedAt() bool {
	if o != nil && !IsNil(o.TerminatedAt) {
		return true
	}

	return false
}

// SetTerminatedAt gets a reference to the given time.Time and assigns it to the TerminatedAt field.
func (o *CouponObject) SetTerminatedAt(v time.Time) {
	o.TerminatedAt = &v
}

func (o CouponObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CouponObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["lago_id"] = o.LagoId
	toSerialize["name"] = o.Name
	toSerialize["code"] = o.Code
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	toSerialize["coupon_type"] = o.CouponType
	if o.AmountCents.IsSet() {
		toSerialize["amount_cents"] = o.AmountCents.Get()
	}
	if !IsNil(o.AmountCurrency) {
		toSerialize["amount_currency"] = o.AmountCurrency
	}
	toSerialize["reusable"] = o.Reusable
	toSerialize["limited_plans"] = o.LimitedPlans
	if !IsNil(o.PlanCodes) {
		toSerialize["plan_codes"] = o.PlanCodes
	}
	toSerialize["limited_billable_metrics"] = o.LimitedBillableMetrics
	if !IsNil(o.BillableMetricCodes) {
		toSerialize["billable_metric_codes"] = o.BillableMetricCodes
	}
	if o.PercentageRate.IsSet() {
		toSerialize["percentage_rate"] = o.PercentageRate.Get()
	}
	toSerialize["frequency"] = o.Frequency
	if o.FrequencyDuration.IsSet() {
		toSerialize["frequency_duration"] = o.FrequencyDuration.Get()
	}
	toSerialize["expiration"] = o.Expiration
	if o.ExpirationAt.IsSet() {
		toSerialize["expiration_at"] = o.ExpirationAt.Get()
	}
	toSerialize["created_at"] = o.CreatedAt
	if !IsNil(o.TerminatedAt) {
		toSerialize["terminated_at"] = o.TerminatedAt
	}
	return toSerialize, nil
}

func (o *CouponObject) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"lago_id",
		"name",
		"code",
		"coupon_type",
		"reusable",
		"limited_plans",
		"limited_billable_metrics",
		"frequency",
		"expiration",
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

	varCouponObject := _CouponObject{}

	err = json.Unmarshal(bytes, &varCouponObject)

	if err != nil {
		return err
	}

	*o = CouponObject(varCouponObject)

	return err
}

type NullableCouponObject struct {
	value *CouponObject
	isSet bool
}

func (v NullableCouponObject) Get() *CouponObject {
	return v.value
}

func (v *NullableCouponObject) Set(val *CouponObject) {
	v.value = val
	v.isSet = true
}

func (v NullableCouponObject) IsSet() bool {
	return v.isSet
}

func (v *NullableCouponObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCouponObject(val *CouponObject) *NullableCouponObject {
	return &NullableCouponObject{value: val, isSet: true}
}

func (v NullableCouponObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCouponObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


