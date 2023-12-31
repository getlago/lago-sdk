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

// checks if the BillableMetricBaseInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillableMetricBaseInput{}

// BillableMetricBaseInput struct for BillableMetricBaseInput
type BillableMetricBaseInput struct {
	// Name of the billable metric.
	Name *string `json:"name,omitempty"`
	// Unique code used to identify the billable metric associated with the API request. This code associates each event with the correct metric.
	Code *string `json:"code,omitempty"`
	// Internal description of the billable metric.
	Description NullableString `json:"description,omitempty"`
	// Defines if the billable metric is persisted billing period over billing period.  - If set to `true`: the accumulated number of units calculated from the previous billing period is persisted to the next billing period. - If set to `false`: the accumulated number of units is reset to 0 at the end of the billing period. - If not defined in the request, default value is `false`.
	Recurring *bool `json:"recurring,omitempty"`
	// Property of the billable metric used for aggregating usage data. This field is not required for `count_agg`.
	FieldName NullableString `json:"field_name,omitempty"`
	// Aggregation method used to compute usage for this billable metric.
	AggregationType *string `json:"aggregation_type,omitempty"`
	// Parameter exclusively utilized in conjunction with the `weighted_sum` aggregation type. It serves to adjust the aggregation result by assigning weights and proration to the result based on time intervals. When this field is not provided, the default time interval is assumed to be in `seconds`.
	WeightedInterval NullableString `json:"weighted_interval,omitempty"`
	Group *BillableMetricGroup `json:"group,omitempty"`
}

// NewBillableMetricBaseInput instantiates a new BillableMetricBaseInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillableMetricBaseInput() *BillableMetricBaseInput {
	this := BillableMetricBaseInput{}
	return &this
}

// NewBillableMetricBaseInputWithDefaults instantiates a new BillableMetricBaseInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillableMetricBaseInputWithDefaults() *BillableMetricBaseInput {
	this := BillableMetricBaseInput{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BillableMetricBaseInput) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillableMetricBaseInput) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BillableMetricBaseInput) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BillableMetricBaseInput) SetName(v string) {
	o.Name = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *BillableMetricBaseInput) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillableMetricBaseInput) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *BillableMetricBaseInput) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *BillableMetricBaseInput) SetCode(v string) {
	o.Code = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillableMetricBaseInput) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillableMetricBaseInput) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *BillableMetricBaseInput) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *BillableMetricBaseInput) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *BillableMetricBaseInput) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *BillableMetricBaseInput) UnsetDescription() {
	o.Description.Unset()
}

// GetRecurring returns the Recurring field value if set, zero value otherwise.
func (o *BillableMetricBaseInput) GetRecurring() bool {
	if o == nil || IsNil(o.Recurring) {
		var ret bool
		return ret
	}
	return *o.Recurring
}

// GetRecurringOk returns a tuple with the Recurring field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillableMetricBaseInput) GetRecurringOk() (*bool, bool) {
	if o == nil || IsNil(o.Recurring) {
		return nil, false
	}
	return o.Recurring, true
}

// HasRecurring returns a boolean if a field has been set.
func (o *BillableMetricBaseInput) HasRecurring() bool {
	if o != nil && !IsNil(o.Recurring) {
		return true
	}

	return false
}

// SetRecurring gets a reference to the given bool and assigns it to the Recurring field.
func (o *BillableMetricBaseInput) SetRecurring(v bool) {
	o.Recurring = &v
}

// GetFieldName returns the FieldName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillableMetricBaseInput) GetFieldName() string {
	if o == nil || IsNil(o.FieldName.Get()) {
		var ret string
		return ret
	}
	return *o.FieldName.Get()
}

// GetFieldNameOk returns a tuple with the FieldName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillableMetricBaseInput) GetFieldNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FieldName.Get(), o.FieldName.IsSet()
}

// HasFieldName returns a boolean if a field has been set.
func (o *BillableMetricBaseInput) HasFieldName() bool {
	if o != nil && o.FieldName.IsSet() {
		return true
	}

	return false
}

// SetFieldName gets a reference to the given NullableString and assigns it to the FieldName field.
func (o *BillableMetricBaseInput) SetFieldName(v string) {
	o.FieldName.Set(&v)
}
// SetFieldNameNil sets the value for FieldName to be an explicit nil
func (o *BillableMetricBaseInput) SetFieldNameNil() {
	o.FieldName.Set(nil)
}

// UnsetFieldName ensures that no value is present for FieldName, not even an explicit nil
func (o *BillableMetricBaseInput) UnsetFieldName() {
	o.FieldName.Unset()
}

// GetAggregationType returns the AggregationType field value if set, zero value otherwise.
func (o *BillableMetricBaseInput) GetAggregationType() string {
	if o == nil || IsNil(o.AggregationType) {
		var ret string
		return ret
	}
	return *o.AggregationType
}

// GetAggregationTypeOk returns a tuple with the AggregationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillableMetricBaseInput) GetAggregationTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AggregationType) {
		return nil, false
	}
	return o.AggregationType, true
}

// HasAggregationType returns a boolean if a field has been set.
func (o *BillableMetricBaseInput) HasAggregationType() bool {
	if o != nil && !IsNil(o.AggregationType) {
		return true
	}

	return false
}

// SetAggregationType gets a reference to the given string and assigns it to the AggregationType field.
func (o *BillableMetricBaseInput) SetAggregationType(v string) {
	o.AggregationType = &v
}

// GetWeightedInterval returns the WeightedInterval field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillableMetricBaseInput) GetWeightedInterval() string {
	if o == nil || IsNil(o.WeightedInterval.Get()) {
		var ret string
		return ret
	}
	return *o.WeightedInterval.Get()
}

// GetWeightedIntervalOk returns a tuple with the WeightedInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillableMetricBaseInput) GetWeightedIntervalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WeightedInterval.Get(), o.WeightedInterval.IsSet()
}

// HasWeightedInterval returns a boolean if a field has been set.
func (o *BillableMetricBaseInput) HasWeightedInterval() bool {
	if o != nil && o.WeightedInterval.IsSet() {
		return true
	}

	return false
}

// SetWeightedInterval gets a reference to the given NullableString and assigns it to the WeightedInterval field.
func (o *BillableMetricBaseInput) SetWeightedInterval(v string) {
	o.WeightedInterval.Set(&v)
}
// SetWeightedIntervalNil sets the value for WeightedInterval to be an explicit nil
func (o *BillableMetricBaseInput) SetWeightedIntervalNil() {
	o.WeightedInterval.Set(nil)
}

// UnsetWeightedInterval ensures that no value is present for WeightedInterval, not even an explicit nil
func (o *BillableMetricBaseInput) UnsetWeightedInterval() {
	o.WeightedInterval.Unset()
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *BillableMetricBaseInput) GetGroup() BillableMetricGroup {
	if o == nil || IsNil(o.Group) {
		var ret BillableMetricGroup
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillableMetricBaseInput) GetGroupOk() (*BillableMetricGroup, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *BillableMetricBaseInput) HasGroup() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given BillableMetricGroup and assigns it to the Group field.
func (o *BillableMetricBaseInput) SetGroup(v BillableMetricGroup) {
	o.Group = &v
}

func (o BillableMetricBaseInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillableMetricBaseInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !IsNil(o.Recurring) {
		toSerialize["recurring"] = o.Recurring
	}
	if o.FieldName.IsSet() {
		toSerialize["field_name"] = o.FieldName.Get()
	}
	if !IsNil(o.AggregationType) {
		toSerialize["aggregation_type"] = o.AggregationType
	}
	if o.WeightedInterval.IsSet() {
		toSerialize["weighted_interval"] = o.WeightedInterval.Get()
	}
	if !IsNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	return toSerialize, nil
}

type NullableBillableMetricBaseInput struct {
	value *BillableMetricBaseInput
	isSet bool
}

func (v NullableBillableMetricBaseInput) Get() *BillableMetricBaseInput {
	return v.value
}

func (v *NullableBillableMetricBaseInput) Set(val *BillableMetricBaseInput) {
	v.value = val
	v.isSet = true
}

func (v NullableBillableMetricBaseInput) IsSet() bool {
	return v.isSet
}

func (v *NullableBillableMetricBaseInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillableMetricBaseInput(val *BillableMetricBaseInput) *NullableBillableMetricBaseInput {
	return &NullableBillableMetricBaseInput{value: val, isSet: true}
}

func (v NullableBillableMetricBaseInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillableMetricBaseInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


