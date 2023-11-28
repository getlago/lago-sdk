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

// checks if the SubscriptionObjectExtended type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionObjectExtended{}

// SubscriptionObjectExtended struct for SubscriptionObjectExtended
type SubscriptionObjectExtended struct {
	// Unique identifier assigned to the subscription within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the subscription’s record within the Lago system
	LagoId string `json:"lago_id"`
	// The subscription external unique identifier (provided by your own application).
	ExternalId string `json:"external_id"`
	// Unique identifier assigned to the customer within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the customer's record within the Lago system
	LagoCustomerId string `json:"lago_customer_id"`
	// The customer external unique identifier (provided by your own application).
	ExternalCustomerId string `json:"external_customer_id"`
	// The billing time for the subscription, which can be set as either `anniversary` or `calendar`. If not explicitly provided, it will default to `calendar`. The billing time determines the timing of recurring billing cycles for the subscription. By specifying `anniversary`, the billing cycle will be based on the specific date the subscription started (billed fully), while `calendar` sets the billing cycle at the first day of the week/month/year (billed with proration).
	BillingTime string `json:"billing_time"`
	// The display name of the subscription on an invoice. This field allows for customization of the subscription's name for billing purposes, especially useful when a single customer has multiple subscriptions using the same plan.
	Name NullableString `json:"name,omitempty"`
	// The unique code representing the plan to be attached to the customer. This code must correspond to the `code` property of one of the active plans.
	PlanCode string `json:"plan_code"`
	// The status of the subscription, which can have the following values: - `pending`: a previous subscription has been downgraded, and the current one is awaiting automatic activation at the end of the billing period. - `active`: the subscription is currently active and applied to the customer. - `terminated`: the subscription is no longer active. - `canceled`: the subscription has been stopped before its activation. This can occur when two consecutive downgrades have been applied to a customer or when a subscription with a pending status is terminated.
	Status string `json:"status"`
	// The creation date of the subscription, represented in ISO 8601 datetime format and expressed in Coordinated Universal Time (UTC). This date provides a timestamp indicating when the subscription was initially created.
	CreatedAt time.Time `json:"created_at"`
	// The cancellation date of the subscription. This field is not null when the subscription is `canceled`. This date should be provided in ISO 8601 datetime format and expressed in Coordinated Universal Time (UTC).
	CanceledAt NullableTime `json:"canceled_at,omitempty"`
	// The effective start date of the subscription. This field can be null if the subscription is `pending` or `canceled`. This date should be provided in ISO 8601 datetime format and expressed in Coordinated Universal Time (UTC).
	StartedAt NullableTime `json:"started_at,omitempty"`
	// The effective end date of the subscription. If this field is set to null, the subscription will automatically renew. This date should be provided in ISO 8601 datetime format, and use Coordinated Universal Time (UTC).
	EndingAt *time.Time `json:"ending_at,omitempty"`
	// The anniversary date and time of the initial subscription. This date serves as the basis for billing subscriptions with `anniversary` billing time. The `anniversary_date` should be provided in ISO 8601 datetime format and expressed in Coordinated Universal Time (UTC).
	SubscriptionAt time.Time `json:"subscription_at"`
	// The termination date of the subscription. This field is not null when the subscription is `terminated`. This date should be provided in ISO 8601 datetime format and expressed in Coordinated Universal Time (UTC)
	TerminatedAt NullableTime `json:"terminated_at,omitempty"`
	// The code identifying the previous plan associated with this subscription.
	PreviousPlanCode NullableString `json:"previous_plan_code,omitempty"`
	// The code identifying the next plan in the case of a downgrade.
	NextPlanCode NullableString `json:"next_plan_code,omitempty"`
	// The date when the plan will be downgraded, represented in ISO 8601 date format.
	DowngradePlanDate NullableTime `json:"downgrade_plan_date,omitempty"`
	Plan *PlanObject `json:"plan,omitempty"`
}

type _SubscriptionObjectExtended SubscriptionObjectExtended

// NewSubscriptionObjectExtended instantiates a new SubscriptionObjectExtended object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionObjectExtended(lagoId string, externalId string, lagoCustomerId string, externalCustomerId string, billingTime string, planCode string, status string, createdAt time.Time, subscriptionAt time.Time) *SubscriptionObjectExtended {
	this := SubscriptionObjectExtended{}
	this.LagoId = lagoId
	this.ExternalId = externalId
	this.LagoCustomerId = lagoCustomerId
	this.ExternalCustomerId = externalCustomerId
	this.BillingTime = billingTime
	this.PlanCode = planCode
	this.Status = status
	this.CreatedAt = createdAt
	this.SubscriptionAt = subscriptionAt
	return &this
}

// NewSubscriptionObjectExtendedWithDefaults instantiates a new SubscriptionObjectExtended object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionObjectExtendedWithDefaults() *SubscriptionObjectExtended {
	this := SubscriptionObjectExtended{}
	return &this
}

// GetLagoId returns the LagoId field value
func (o *SubscriptionObjectExtended) GetLagoId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LagoId
}

// GetLagoIdOk returns a tuple with the LagoId field value
// and a boolean to check if the value has been set.
func (o *SubscriptionObjectExtended) GetLagoIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LagoId, true
}

// SetLagoId sets field value
func (o *SubscriptionObjectExtended) SetLagoId(v string) {
	o.LagoId = v
}

// GetExternalId returns the ExternalId field value
func (o *SubscriptionObjectExtended) GetExternalId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value
// and a boolean to check if the value has been set.
func (o *SubscriptionObjectExtended) GetExternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalId, true
}

// SetExternalId sets field value
func (o *SubscriptionObjectExtended) SetExternalId(v string) {
	o.ExternalId = v
}

// GetLagoCustomerId returns the LagoCustomerId field value
func (o *SubscriptionObjectExtended) GetLagoCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LagoCustomerId
}

// GetLagoCustomerIdOk returns a tuple with the LagoCustomerId field value
// and a boolean to check if the value has been set.
func (o *SubscriptionObjectExtended) GetLagoCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LagoCustomerId, true
}

// SetLagoCustomerId sets field value
func (o *SubscriptionObjectExtended) SetLagoCustomerId(v string) {
	o.LagoCustomerId = v
}

// GetExternalCustomerId returns the ExternalCustomerId field value
func (o *SubscriptionObjectExtended) GetExternalCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalCustomerId
}

// GetExternalCustomerIdOk returns a tuple with the ExternalCustomerId field value
// and a boolean to check if the value has been set.
func (o *SubscriptionObjectExtended) GetExternalCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalCustomerId, true
}

// SetExternalCustomerId sets field value
func (o *SubscriptionObjectExtended) SetExternalCustomerId(v string) {
	o.ExternalCustomerId = v
}

// GetBillingTime returns the BillingTime field value
func (o *SubscriptionObjectExtended) GetBillingTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BillingTime
}

// GetBillingTimeOk returns a tuple with the BillingTime field value
// and a boolean to check if the value has been set.
func (o *SubscriptionObjectExtended) GetBillingTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BillingTime, true
}

// SetBillingTime sets field value
func (o *SubscriptionObjectExtended) SetBillingTime(v string) {
	o.BillingTime = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubscriptionObjectExtended) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionObjectExtended) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *SubscriptionObjectExtended) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *SubscriptionObjectExtended) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *SubscriptionObjectExtended) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *SubscriptionObjectExtended) UnsetName() {
	o.Name.Unset()
}

// GetPlanCode returns the PlanCode field value
func (o *SubscriptionObjectExtended) GetPlanCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlanCode
}

// GetPlanCodeOk returns a tuple with the PlanCode field value
// and a boolean to check if the value has been set.
func (o *SubscriptionObjectExtended) GetPlanCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlanCode, true
}

// SetPlanCode sets field value
func (o *SubscriptionObjectExtended) SetPlanCode(v string) {
	o.PlanCode = v
}

// GetStatus returns the Status field value
func (o *SubscriptionObjectExtended) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *SubscriptionObjectExtended) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *SubscriptionObjectExtended) SetStatus(v string) {
	o.Status = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *SubscriptionObjectExtended) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *SubscriptionObjectExtended) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *SubscriptionObjectExtended) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCanceledAt returns the CanceledAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubscriptionObjectExtended) GetCanceledAt() time.Time {
	if o == nil || IsNil(o.CanceledAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.CanceledAt.Get()
}

// GetCanceledAtOk returns a tuple with the CanceledAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionObjectExtended) GetCanceledAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CanceledAt.Get(), o.CanceledAt.IsSet()
}

// HasCanceledAt returns a boolean if a field has been set.
func (o *SubscriptionObjectExtended) HasCanceledAt() bool {
	if o != nil && o.CanceledAt.IsSet() {
		return true
	}

	return false
}

// SetCanceledAt gets a reference to the given NullableTime and assigns it to the CanceledAt field.
func (o *SubscriptionObjectExtended) SetCanceledAt(v time.Time) {
	o.CanceledAt.Set(&v)
}
// SetCanceledAtNil sets the value for CanceledAt to be an explicit nil
func (o *SubscriptionObjectExtended) SetCanceledAtNil() {
	o.CanceledAt.Set(nil)
}

// UnsetCanceledAt ensures that no value is present for CanceledAt, not even an explicit nil
func (o *SubscriptionObjectExtended) UnsetCanceledAt() {
	o.CanceledAt.Unset()
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubscriptionObjectExtended) GetStartedAt() time.Time {
	if o == nil || IsNil(o.StartedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.StartedAt.Get()
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionObjectExtended) GetStartedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartedAt.Get(), o.StartedAt.IsSet()
}

// HasStartedAt returns a boolean if a field has been set.
func (o *SubscriptionObjectExtended) HasStartedAt() bool {
	if o != nil && o.StartedAt.IsSet() {
		return true
	}

	return false
}

// SetStartedAt gets a reference to the given NullableTime and assigns it to the StartedAt field.
func (o *SubscriptionObjectExtended) SetStartedAt(v time.Time) {
	o.StartedAt.Set(&v)
}
// SetStartedAtNil sets the value for StartedAt to be an explicit nil
func (o *SubscriptionObjectExtended) SetStartedAtNil() {
	o.StartedAt.Set(nil)
}

// UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
func (o *SubscriptionObjectExtended) UnsetStartedAt() {
	o.StartedAt.Unset()
}

// GetEndingAt returns the EndingAt field value if set, zero value otherwise.
func (o *SubscriptionObjectExtended) GetEndingAt() time.Time {
	if o == nil || IsNil(o.EndingAt) {
		var ret time.Time
		return ret
	}
	return *o.EndingAt
}

// GetEndingAtOk returns a tuple with the EndingAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionObjectExtended) GetEndingAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndingAt) {
		return nil, false
	}
	return o.EndingAt, true
}

// HasEndingAt returns a boolean if a field has been set.
func (o *SubscriptionObjectExtended) HasEndingAt() bool {
	if o != nil && !IsNil(o.EndingAt) {
		return true
	}

	return false
}

// SetEndingAt gets a reference to the given time.Time and assigns it to the EndingAt field.
func (o *SubscriptionObjectExtended) SetEndingAt(v time.Time) {
	o.EndingAt = &v
}

// GetSubscriptionAt returns the SubscriptionAt field value
func (o *SubscriptionObjectExtended) GetSubscriptionAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.SubscriptionAt
}

// GetSubscriptionAtOk returns a tuple with the SubscriptionAt field value
// and a boolean to check if the value has been set.
func (o *SubscriptionObjectExtended) GetSubscriptionAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscriptionAt, true
}

// SetSubscriptionAt sets field value
func (o *SubscriptionObjectExtended) SetSubscriptionAt(v time.Time) {
	o.SubscriptionAt = v
}

// GetTerminatedAt returns the TerminatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubscriptionObjectExtended) GetTerminatedAt() time.Time {
	if o == nil || IsNil(o.TerminatedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.TerminatedAt.Get()
}

// GetTerminatedAtOk returns a tuple with the TerminatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionObjectExtended) GetTerminatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.TerminatedAt.Get(), o.TerminatedAt.IsSet()
}

// HasTerminatedAt returns a boolean if a field has been set.
func (o *SubscriptionObjectExtended) HasTerminatedAt() bool {
	if o != nil && o.TerminatedAt.IsSet() {
		return true
	}

	return false
}

// SetTerminatedAt gets a reference to the given NullableTime and assigns it to the TerminatedAt field.
func (o *SubscriptionObjectExtended) SetTerminatedAt(v time.Time) {
	o.TerminatedAt.Set(&v)
}
// SetTerminatedAtNil sets the value for TerminatedAt to be an explicit nil
func (o *SubscriptionObjectExtended) SetTerminatedAtNil() {
	o.TerminatedAt.Set(nil)
}

// UnsetTerminatedAt ensures that no value is present for TerminatedAt, not even an explicit nil
func (o *SubscriptionObjectExtended) UnsetTerminatedAt() {
	o.TerminatedAt.Unset()
}

// GetPreviousPlanCode returns the PreviousPlanCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubscriptionObjectExtended) GetPreviousPlanCode() string {
	if o == nil || IsNil(o.PreviousPlanCode.Get()) {
		var ret string
		return ret
	}
	return *o.PreviousPlanCode.Get()
}

// GetPreviousPlanCodeOk returns a tuple with the PreviousPlanCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionObjectExtended) GetPreviousPlanCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PreviousPlanCode.Get(), o.PreviousPlanCode.IsSet()
}

// HasPreviousPlanCode returns a boolean if a field has been set.
func (o *SubscriptionObjectExtended) HasPreviousPlanCode() bool {
	if o != nil && o.PreviousPlanCode.IsSet() {
		return true
	}

	return false
}

// SetPreviousPlanCode gets a reference to the given NullableString and assigns it to the PreviousPlanCode field.
func (o *SubscriptionObjectExtended) SetPreviousPlanCode(v string) {
	o.PreviousPlanCode.Set(&v)
}
// SetPreviousPlanCodeNil sets the value for PreviousPlanCode to be an explicit nil
func (o *SubscriptionObjectExtended) SetPreviousPlanCodeNil() {
	o.PreviousPlanCode.Set(nil)
}

// UnsetPreviousPlanCode ensures that no value is present for PreviousPlanCode, not even an explicit nil
func (o *SubscriptionObjectExtended) UnsetPreviousPlanCode() {
	o.PreviousPlanCode.Unset()
}

// GetNextPlanCode returns the NextPlanCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubscriptionObjectExtended) GetNextPlanCode() string {
	if o == nil || IsNil(o.NextPlanCode.Get()) {
		var ret string
		return ret
	}
	return *o.NextPlanCode.Get()
}

// GetNextPlanCodeOk returns a tuple with the NextPlanCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionObjectExtended) GetNextPlanCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextPlanCode.Get(), o.NextPlanCode.IsSet()
}

// HasNextPlanCode returns a boolean if a field has been set.
func (o *SubscriptionObjectExtended) HasNextPlanCode() bool {
	if o != nil && o.NextPlanCode.IsSet() {
		return true
	}

	return false
}

// SetNextPlanCode gets a reference to the given NullableString and assigns it to the NextPlanCode field.
func (o *SubscriptionObjectExtended) SetNextPlanCode(v string) {
	o.NextPlanCode.Set(&v)
}
// SetNextPlanCodeNil sets the value for NextPlanCode to be an explicit nil
func (o *SubscriptionObjectExtended) SetNextPlanCodeNil() {
	o.NextPlanCode.Set(nil)
}

// UnsetNextPlanCode ensures that no value is present for NextPlanCode, not even an explicit nil
func (o *SubscriptionObjectExtended) UnsetNextPlanCode() {
	o.NextPlanCode.Unset()
}

// GetDowngradePlanDate returns the DowngradePlanDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubscriptionObjectExtended) GetDowngradePlanDate() time.Time {
	if o == nil || IsNil(o.DowngradePlanDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.DowngradePlanDate.Get()
}

// GetDowngradePlanDateOk returns a tuple with the DowngradePlanDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionObjectExtended) GetDowngradePlanDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DowngradePlanDate.Get(), o.DowngradePlanDate.IsSet()
}

// HasDowngradePlanDate returns a boolean if a field has been set.
func (o *SubscriptionObjectExtended) HasDowngradePlanDate() bool {
	if o != nil && o.DowngradePlanDate.IsSet() {
		return true
	}

	return false
}

// SetDowngradePlanDate gets a reference to the given NullableTime and assigns it to the DowngradePlanDate field.
func (o *SubscriptionObjectExtended) SetDowngradePlanDate(v time.Time) {
	o.DowngradePlanDate.Set(&v)
}
// SetDowngradePlanDateNil sets the value for DowngradePlanDate to be an explicit nil
func (o *SubscriptionObjectExtended) SetDowngradePlanDateNil() {
	o.DowngradePlanDate.Set(nil)
}

// UnsetDowngradePlanDate ensures that no value is present for DowngradePlanDate, not even an explicit nil
func (o *SubscriptionObjectExtended) UnsetDowngradePlanDate() {
	o.DowngradePlanDate.Unset()
}

// GetPlan returns the Plan field value if set, zero value otherwise.
func (o *SubscriptionObjectExtended) GetPlan() PlanObject {
	if o == nil || IsNil(o.Plan) {
		var ret PlanObject
		return ret
	}
	return *o.Plan
}

// GetPlanOk returns a tuple with the Plan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionObjectExtended) GetPlanOk() (*PlanObject, bool) {
	if o == nil || IsNil(o.Plan) {
		return nil, false
	}
	return o.Plan, true
}

// HasPlan returns a boolean if a field has been set.
func (o *SubscriptionObjectExtended) HasPlan() bool {
	if o != nil && !IsNil(o.Plan) {
		return true
	}

	return false
}

// SetPlan gets a reference to the given PlanObject and assigns it to the Plan field.
func (o *SubscriptionObjectExtended) SetPlan(v PlanObject) {
	o.Plan = &v
}

func (o SubscriptionObjectExtended) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionObjectExtended) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["lago_id"] = o.LagoId
	toSerialize["external_id"] = o.ExternalId
	toSerialize["lago_customer_id"] = o.LagoCustomerId
	toSerialize["external_customer_id"] = o.ExternalCustomerId
	toSerialize["billing_time"] = o.BillingTime
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	toSerialize["plan_code"] = o.PlanCode
	toSerialize["status"] = o.Status
	toSerialize["created_at"] = o.CreatedAt
	if o.CanceledAt.IsSet() {
		toSerialize["canceled_at"] = o.CanceledAt.Get()
	}
	if o.StartedAt.IsSet() {
		toSerialize["started_at"] = o.StartedAt.Get()
	}
	if !IsNil(o.EndingAt) {
		toSerialize["ending_at"] = o.EndingAt
	}
	toSerialize["subscription_at"] = o.SubscriptionAt
	if o.TerminatedAt.IsSet() {
		toSerialize["terminated_at"] = o.TerminatedAt.Get()
	}
	if o.PreviousPlanCode.IsSet() {
		toSerialize["previous_plan_code"] = o.PreviousPlanCode.Get()
	}
	if o.NextPlanCode.IsSet() {
		toSerialize["next_plan_code"] = o.NextPlanCode.Get()
	}
	if o.DowngradePlanDate.IsSet() {
		toSerialize["downgrade_plan_date"] = o.DowngradePlanDate.Get()
	}
	if !IsNil(o.Plan) {
		toSerialize["plan"] = o.Plan
	}
	return toSerialize, nil
}

func (o *SubscriptionObjectExtended) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"lago_id",
		"external_id",
		"lago_customer_id",
		"external_customer_id",
		"billing_time",
		"plan_code",
		"status",
		"created_at",
		"subscription_at",
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

	varSubscriptionObjectExtended := _SubscriptionObjectExtended{}

	err = json.Unmarshal(bytes, &varSubscriptionObjectExtended)

	if err != nil {
		return err
	}

	*o = SubscriptionObjectExtended(varSubscriptionObjectExtended)

	return err
}

type NullableSubscriptionObjectExtended struct {
	value *SubscriptionObjectExtended
	isSet bool
}

func (v NullableSubscriptionObjectExtended) Get() *SubscriptionObjectExtended {
	return v.value
}

func (v *NullableSubscriptionObjectExtended) Set(val *SubscriptionObjectExtended) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionObjectExtended) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionObjectExtended) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionObjectExtended(val *SubscriptionObjectExtended) *NullableSubscriptionObjectExtended {
	return &NullableSubscriptionObjectExtended{value: val, isSet: true}
}

func (v NullableSubscriptionObjectExtended) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionObjectExtended) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


