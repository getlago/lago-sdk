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

// checks if the ChargeObjectProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChargeObjectProperties{}

// ChargeObjectProperties struct for ChargeObjectProperties
type ChargeObjectProperties struct {
	// Graduated ranges, sorted from bottom to top tiers, used for a `graduated` charge model.
	GraduatedRanges []ChargePropertiesGraduatedRangesInner `json:"graduated_ranges,omitempty"`
	// Graduated percentage ranges, sorted from bottom to top tiers, used for a `graduated_percentage` charge model.
	GraduatedPercentageRanges []ChargePropertiesGraduatedPercentageRangesInner `json:"graduated_percentage_ranges,omitempty"`
	// - The unit price, excluding tax, for a `standard` charge model. It is expressed as a decimal value. - The amount, excluding tax, for a complete set of units in a `package` charge model. It is expressed as a decimal value.
	Amount *string `json:"amount,omitempty"`
	// The quantity of units that are provided free of charge for each billing period in a `package` charge model. This field specifies the number of units that customers can use without incurring any additional cost during each billing cycle.
	FreeUnits *int32 `json:"free_units,omitempty"`
	// The quantity of units included in each pack or set for a `package` charge model. It indicates the number of units that are bundled together as a single package or set within the pricing structure.
	PackageSize *int32 `json:"package_size,omitempty"`
	// The percentage rate that is applied to the amount of each transaction for a `percentage` charge model. It is expressed as a decimal value.
	Rate *string `json:"rate,omitempty"`
	// The fixed fee that is applied to each transaction for a `percentage` charge model. It is expressed as a decimal value.
	FixedAmount *string `json:"fixed_amount,omitempty"`
	// The count of transactions that are not impacted by the `percentage` rate and fixed fee in a percentage charge model. This field indicates the number of transactions that are exempt from the calculation of charges based on the specified percentage rate and fixed fee.
	FreeUnitsPerEvents NullableInt32 `json:"free_units_per_events,omitempty"`
	// The transaction amount that is not impacted by the `percentage` rate and fixed fee in a percentage charge model. This field indicates the portion of the transaction amount that is exempt from the calculation of charges based on the specified percentage rate and fixed fee.
	FreeUnitsPerTotalAggregation NullableString `json:"free_units_per_total_aggregation,omitempty"`
	// Specifies the maximum allowable spending for a single transaction. Working as a transaction cap.
	PerTransactionMaxAmount NullableString `json:"per_transaction_max_amount,omitempty"`
	// Specifies the minimum allowable spending for a single transaction. Working as a transaction floor.
	PerTransactionMinAmount NullableString `json:"per_transaction_min_amount,omitempty"`
	// Volume ranges, sorted from bottom to top tiers, used for a `volume` charge model.
	VolumeRanges []ChargePropertiesVolumeRangesInner `json:"volume_ranges,omitempty"`
}

// NewChargeObjectProperties instantiates a new ChargeObjectProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChargeObjectProperties() *ChargeObjectProperties {
	this := ChargeObjectProperties{}
	return &this
}

// NewChargeObjectPropertiesWithDefaults instantiates a new ChargeObjectProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChargeObjectPropertiesWithDefaults() *ChargeObjectProperties {
	this := ChargeObjectProperties{}
	return &this
}

// GetGraduatedRanges returns the GraduatedRanges field value if set, zero value otherwise.
func (o *ChargeObjectProperties) GetGraduatedRanges() []ChargePropertiesGraduatedRangesInner {
	if o == nil || IsNil(o.GraduatedRanges) {
		var ret []ChargePropertiesGraduatedRangesInner
		return ret
	}
	return o.GraduatedRanges
}

// GetGraduatedRangesOk returns a tuple with the GraduatedRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeObjectProperties) GetGraduatedRangesOk() ([]ChargePropertiesGraduatedRangesInner, bool) {
	if o == nil || IsNil(o.GraduatedRanges) {
		return nil, false
	}
	return o.GraduatedRanges, true
}

// HasGraduatedRanges returns a boolean if a field has been set.
func (o *ChargeObjectProperties) HasGraduatedRanges() bool {
	if o != nil && !IsNil(o.GraduatedRanges) {
		return true
	}

	return false
}

// SetGraduatedRanges gets a reference to the given []ChargePropertiesGraduatedRangesInner and assigns it to the GraduatedRanges field.
func (o *ChargeObjectProperties) SetGraduatedRanges(v []ChargePropertiesGraduatedRangesInner) {
	o.GraduatedRanges = v
}

// GetGraduatedPercentageRanges returns the GraduatedPercentageRanges field value if set, zero value otherwise.
func (o *ChargeObjectProperties) GetGraduatedPercentageRanges() []ChargePropertiesGraduatedPercentageRangesInner {
	if o == nil || IsNil(o.GraduatedPercentageRanges) {
		var ret []ChargePropertiesGraduatedPercentageRangesInner
		return ret
	}
	return o.GraduatedPercentageRanges
}

// GetGraduatedPercentageRangesOk returns a tuple with the GraduatedPercentageRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeObjectProperties) GetGraduatedPercentageRangesOk() ([]ChargePropertiesGraduatedPercentageRangesInner, bool) {
	if o == nil || IsNil(o.GraduatedPercentageRanges) {
		return nil, false
	}
	return o.GraduatedPercentageRanges, true
}

// HasGraduatedPercentageRanges returns a boolean if a field has been set.
func (o *ChargeObjectProperties) HasGraduatedPercentageRanges() bool {
	if o != nil && !IsNil(o.GraduatedPercentageRanges) {
		return true
	}

	return false
}

// SetGraduatedPercentageRanges gets a reference to the given []ChargePropertiesGraduatedPercentageRangesInner and assigns it to the GraduatedPercentageRanges field.
func (o *ChargeObjectProperties) SetGraduatedPercentageRanges(v []ChargePropertiesGraduatedPercentageRangesInner) {
	o.GraduatedPercentageRanges = v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *ChargeObjectProperties) GetAmount() string {
	if o == nil || IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeObjectProperties) GetAmountOk() (*string, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *ChargeObjectProperties) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *ChargeObjectProperties) SetAmount(v string) {
	o.Amount = &v
}

// GetFreeUnits returns the FreeUnits field value if set, zero value otherwise.
func (o *ChargeObjectProperties) GetFreeUnits() int32 {
	if o == nil || IsNil(o.FreeUnits) {
		var ret int32
		return ret
	}
	return *o.FreeUnits
}

// GetFreeUnitsOk returns a tuple with the FreeUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeObjectProperties) GetFreeUnitsOk() (*int32, bool) {
	if o == nil || IsNil(o.FreeUnits) {
		return nil, false
	}
	return o.FreeUnits, true
}

// HasFreeUnits returns a boolean if a field has been set.
func (o *ChargeObjectProperties) HasFreeUnits() bool {
	if o != nil && !IsNil(o.FreeUnits) {
		return true
	}

	return false
}

// SetFreeUnits gets a reference to the given int32 and assigns it to the FreeUnits field.
func (o *ChargeObjectProperties) SetFreeUnits(v int32) {
	o.FreeUnits = &v
}

// GetPackageSize returns the PackageSize field value if set, zero value otherwise.
func (o *ChargeObjectProperties) GetPackageSize() int32 {
	if o == nil || IsNil(o.PackageSize) {
		var ret int32
		return ret
	}
	return *o.PackageSize
}

// GetPackageSizeOk returns a tuple with the PackageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeObjectProperties) GetPackageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PackageSize) {
		return nil, false
	}
	return o.PackageSize, true
}

// HasPackageSize returns a boolean if a field has been set.
func (o *ChargeObjectProperties) HasPackageSize() bool {
	if o != nil && !IsNil(o.PackageSize) {
		return true
	}

	return false
}

// SetPackageSize gets a reference to the given int32 and assigns it to the PackageSize field.
func (o *ChargeObjectProperties) SetPackageSize(v int32) {
	o.PackageSize = &v
}

// GetRate returns the Rate field value if set, zero value otherwise.
func (o *ChargeObjectProperties) GetRate() string {
	if o == nil || IsNil(o.Rate) {
		var ret string
		return ret
	}
	return *o.Rate
}

// GetRateOk returns a tuple with the Rate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeObjectProperties) GetRateOk() (*string, bool) {
	if o == nil || IsNil(o.Rate) {
		return nil, false
	}
	return o.Rate, true
}

// HasRate returns a boolean if a field has been set.
func (o *ChargeObjectProperties) HasRate() bool {
	if o != nil && !IsNil(o.Rate) {
		return true
	}

	return false
}

// SetRate gets a reference to the given string and assigns it to the Rate field.
func (o *ChargeObjectProperties) SetRate(v string) {
	o.Rate = &v
}

// GetFixedAmount returns the FixedAmount field value if set, zero value otherwise.
func (o *ChargeObjectProperties) GetFixedAmount() string {
	if o == nil || IsNil(o.FixedAmount) {
		var ret string
		return ret
	}
	return *o.FixedAmount
}

// GetFixedAmountOk returns a tuple with the FixedAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeObjectProperties) GetFixedAmountOk() (*string, bool) {
	if o == nil || IsNil(o.FixedAmount) {
		return nil, false
	}
	return o.FixedAmount, true
}

// HasFixedAmount returns a boolean if a field has been set.
func (o *ChargeObjectProperties) HasFixedAmount() bool {
	if o != nil && !IsNil(o.FixedAmount) {
		return true
	}

	return false
}

// SetFixedAmount gets a reference to the given string and assigns it to the FixedAmount field.
func (o *ChargeObjectProperties) SetFixedAmount(v string) {
	o.FixedAmount = &v
}

// GetFreeUnitsPerEvents returns the FreeUnitsPerEvents field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChargeObjectProperties) GetFreeUnitsPerEvents() int32 {
	if o == nil || IsNil(o.FreeUnitsPerEvents.Get()) {
		var ret int32
		return ret
	}
	return *o.FreeUnitsPerEvents.Get()
}

// GetFreeUnitsPerEventsOk returns a tuple with the FreeUnitsPerEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChargeObjectProperties) GetFreeUnitsPerEventsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.FreeUnitsPerEvents.Get(), o.FreeUnitsPerEvents.IsSet()
}

// HasFreeUnitsPerEvents returns a boolean if a field has been set.
func (o *ChargeObjectProperties) HasFreeUnitsPerEvents() bool {
	if o != nil && o.FreeUnitsPerEvents.IsSet() {
		return true
	}

	return false
}

// SetFreeUnitsPerEvents gets a reference to the given NullableInt32 and assigns it to the FreeUnitsPerEvents field.
func (o *ChargeObjectProperties) SetFreeUnitsPerEvents(v int32) {
	o.FreeUnitsPerEvents.Set(&v)
}
// SetFreeUnitsPerEventsNil sets the value for FreeUnitsPerEvents to be an explicit nil
func (o *ChargeObjectProperties) SetFreeUnitsPerEventsNil() {
	o.FreeUnitsPerEvents.Set(nil)
}

// UnsetFreeUnitsPerEvents ensures that no value is present for FreeUnitsPerEvents, not even an explicit nil
func (o *ChargeObjectProperties) UnsetFreeUnitsPerEvents() {
	o.FreeUnitsPerEvents.Unset()
}

// GetFreeUnitsPerTotalAggregation returns the FreeUnitsPerTotalAggregation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChargeObjectProperties) GetFreeUnitsPerTotalAggregation() string {
	if o == nil || IsNil(o.FreeUnitsPerTotalAggregation.Get()) {
		var ret string
		return ret
	}
	return *o.FreeUnitsPerTotalAggregation.Get()
}

// GetFreeUnitsPerTotalAggregationOk returns a tuple with the FreeUnitsPerTotalAggregation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChargeObjectProperties) GetFreeUnitsPerTotalAggregationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FreeUnitsPerTotalAggregation.Get(), o.FreeUnitsPerTotalAggregation.IsSet()
}

// HasFreeUnitsPerTotalAggregation returns a boolean if a field has been set.
func (o *ChargeObjectProperties) HasFreeUnitsPerTotalAggregation() bool {
	if o != nil && o.FreeUnitsPerTotalAggregation.IsSet() {
		return true
	}

	return false
}

// SetFreeUnitsPerTotalAggregation gets a reference to the given NullableString and assigns it to the FreeUnitsPerTotalAggregation field.
func (o *ChargeObjectProperties) SetFreeUnitsPerTotalAggregation(v string) {
	o.FreeUnitsPerTotalAggregation.Set(&v)
}
// SetFreeUnitsPerTotalAggregationNil sets the value for FreeUnitsPerTotalAggregation to be an explicit nil
func (o *ChargeObjectProperties) SetFreeUnitsPerTotalAggregationNil() {
	o.FreeUnitsPerTotalAggregation.Set(nil)
}

// UnsetFreeUnitsPerTotalAggregation ensures that no value is present for FreeUnitsPerTotalAggregation, not even an explicit nil
func (o *ChargeObjectProperties) UnsetFreeUnitsPerTotalAggregation() {
	o.FreeUnitsPerTotalAggregation.Unset()
}

// GetPerTransactionMaxAmount returns the PerTransactionMaxAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChargeObjectProperties) GetPerTransactionMaxAmount() string {
	if o == nil || IsNil(o.PerTransactionMaxAmount.Get()) {
		var ret string
		return ret
	}
	return *o.PerTransactionMaxAmount.Get()
}

// GetPerTransactionMaxAmountOk returns a tuple with the PerTransactionMaxAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChargeObjectProperties) GetPerTransactionMaxAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PerTransactionMaxAmount.Get(), o.PerTransactionMaxAmount.IsSet()
}

// HasPerTransactionMaxAmount returns a boolean if a field has been set.
func (o *ChargeObjectProperties) HasPerTransactionMaxAmount() bool {
	if o != nil && o.PerTransactionMaxAmount.IsSet() {
		return true
	}

	return false
}

// SetPerTransactionMaxAmount gets a reference to the given NullableString and assigns it to the PerTransactionMaxAmount field.
func (o *ChargeObjectProperties) SetPerTransactionMaxAmount(v string) {
	o.PerTransactionMaxAmount.Set(&v)
}
// SetPerTransactionMaxAmountNil sets the value for PerTransactionMaxAmount to be an explicit nil
func (o *ChargeObjectProperties) SetPerTransactionMaxAmountNil() {
	o.PerTransactionMaxAmount.Set(nil)
}

// UnsetPerTransactionMaxAmount ensures that no value is present for PerTransactionMaxAmount, not even an explicit nil
func (o *ChargeObjectProperties) UnsetPerTransactionMaxAmount() {
	o.PerTransactionMaxAmount.Unset()
}

// GetPerTransactionMinAmount returns the PerTransactionMinAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChargeObjectProperties) GetPerTransactionMinAmount() string {
	if o == nil || IsNil(o.PerTransactionMinAmount.Get()) {
		var ret string
		return ret
	}
	return *o.PerTransactionMinAmount.Get()
}

// GetPerTransactionMinAmountOk returns a tuple with the PerTransactionMinAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChargeObjectProperties) GetPerTransactionMinAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PerTransactionMinAmount.Get(), o.PerTransactionMinAmount.IsSet()
}

// HasPerTransactionMinAmount returns a boolean if a field has been set.
func (o *ChargeObjectProperties) HasPerTransactionMinAmount() bool {
	if o != nil && o.PerTransactionMinAmount.IsSet() {
		return true
	}

	return false
}

// SetPerTransactionMinAmount gets a reference to the given NullableString and assigns it to the PerTransactionMinAmount field.
func (o *ChargeObjectProperties) SetPerTransactionMinAmount(v string) {
	o.PerTransactionMinAmount.Set(&v)
}
// SetPerTransactionMinAmountNil sets the value for PerTransactionMinAmount to be an explicit nil
func (o *ChargeObjectProperties) SetPerTransactionMinAmountNil() {
	o.PerTransactionMinAmount.Set(nil)
}

// UnsetPerTransactionMinAmount ensures that no value is present for PerTransactionMinAmount, not even an explicit nil
func (o *ChargeObjectProperties) UnsetPerTransactionMinAmount() {
	o.PerTransactionMinAmount.Unset()
}

// GetVolumeRanges returns the VolumeRanges field value if set, zero value otherwise.
func (o *ChargeObjectProperties) GetVolumeRanges() []ChargePropertiesVolumeRangesInner {
	if o == nil || IsNil(o.VolumeRanges) {
		var ret []ChargePropertiesVolumeRangesInner
		return ret
	}
	return o.VolumeRanges
}

// GetVolumeRangesOk returns a tuple with the VolumeRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeObjectProperties) GetVolumeRangesOk() ([]ChargePropertiesVolumeRangesInner, bool) {
	if o == nil || IsNil(o.VolumeRanges) {
		return nil, false
	}
	return o.VolumeRanges, true
}

// HasVolumeRanges returns a boolean if a field has been set.
func (o *ChargeObjectProperties) HasVolumeRanges() bool {
	if o != nil && !IsNil(o.VolumeRanges) {
		return true
	}

	return false
}

// SetVolumeRanges gets a reference to the given []ChargePropertiesVolumeRangesInner and assigns it to the VolumeRanges field.
func (o *ChargeObjectProperties) SetVolumeRanges(v []ChargePropertiesVolumeRangesInner) {
	o.VolumeRanges = v
}

func (o ChargeObjectProperties) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChargeObjectProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GraduatedRanges) {
		toSerialize["graduated_ranges"] = o.GraduatedRanges
	}
	if !IsNil(o.GraduatedPercentageRanges) {
		toSerialize["graduated_percentage_ranges"] = o.GraduatedPercentageRanges
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.FreeUnits) {
		toSerialize["free_units"] = o.FreeUnits
	}
	if !IsNil(o.PackageSize) {
		toSerialize["package_size"] = o.PackageSize
	}
	if !IsNil(o.Rate) {
		toSerialize["rate"] = o.Rate
	}
	if !IsNil(o.FixedAmount) {
		toSerialize["fixed_amount"] = o.FixedAmount
	}
	if o.FreeUnitsPerEvents.IsSet() {
		toSerialize["free_units_per_events"] = o.FreeUnitsPerEvents.Get()
	}
	if o.FreeUnitsPerTotalAggregation.IsSet() {
		toSerialize["free_units_per_total_aggregation"] = o.FreeUnitsPerTotalAggregation.Get()
	}
	if o.PerTransactionMaxAmount.IsSet() {
		toSerialize["per_transaction_max_amount"] = o.PerTransactionMaxAmount.Get()
	}
	if o.PerTransactionMinAmount.IsSet() {
		toSerialize["per_transaction_min_amount"] = o.PerTransactionMinAmount.Get()
	}
	if !IsNil(o.VolumeRanges) {
		toSerialize["volume_ranges"] = o.VolumeRanges
	}
	return toSerialize, nil
}

type NullableChargeObjectProperties struct {
	value *ChargeObjectProperties
	isSet bool
}

func (v NullableChargeObjectProperties) Get() *ChargeObjectProperties {
	return v.value
}

func (v *NullableChargeObjectProperties) Set(val *ChargeObjectProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableChargeObjectProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableChargeObjectProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChargeObjectProperties(val *ChargeObjectProperties) *NullableChargeObjectProperties {
	return &NullableChargeObjectProperties{value: val, isSet: true}
}

func (v NullableChargeObjectProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChargeObjectProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


