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

// checks if the PlanObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlanObject{}

// PlanObject struct for PlanObject
type PlanObject struct {
	// Unique identifier of the plan created by Lago.
	LagoId string `json:"lago_id"`
	// The name of the plan.
	Name string `json:"name"`
	// Specifies the name that will be displayed on an invoice. If no value is set for this field, the name of the plan will be used as the default display name.
	InvoiceDisplayName *string `json:"invoice_display_name,omitempty"`
	// The date and time when the plan was created. It is expressed in UTC format according to the ISO 8601 datetime standard. This field provides the timestamp for the exact moment when the plan was initially created.
	CreatedAt time.Time `json:"created_at"`
	// The code of the plan. It serves as a unique identifier associated with a particular plan. The code is typically used for internal or system-level identification purposes, like assigning a subscription, for instance.
	Code string `json:"code"`
	// The interval used for recurring billing. It represents the frequency at which subscription billing occurs. The interval can be one of the following values: `yearly`, `quarterly`, `monthly` or `weekly`.
	Interval string `json:"interval"`
	// The description on the plan.
	Description *string `json:"description,omitempty"`
	// The base cost of the plan, excluding any applicable taxes, that is billed on a recurring basis. This value is defined at 0 if your plan is a pay-as-you-go plan.
	AmountCents int32 `json:"amount_cents"`
	AmountCurrency Currency `json:"amount_currency"`
	// The duration in days during which the base cost of the plan is offered for free.
	TrialPeriod *float32 `json:"trial_period,omitempty"`
	// This field determines the billing timing for the plan. When set to `true`, the base cost of the plan is due at the beginning of each billing period. Conversely, when set to `false`, the base cost of the plan is due at the end of each billing period.
	PayInAdvance *bool `json:"pay_in_advance,omitempty"`
	// This field, when set to `true`, enables to invoice usage-based charges on monthly basis, even if the cadence of the plan is yearly. This allows customers to pay charges overage on a monthly basis. This can be set to true only if the plan’s interval is `yearly`.
	BillChargesMonthly NullableBool `json:"bill_charges_monthly,omitempty"`
	// The count of active subscriptions that are currently associated with the plan. This field provides valuable information regarding the impact of deleting the plan. By checking the value of this field, you can determine the number of subscriptions that will be affected if the plan is deleted.
	ActiveSubscriptionsCount int32 `json:"active_subscriptions_count"`
	// The number of draft invoices that include a subscription attached to the plan. This field provides valuable information about the impact of deleting the plan. By checking the value of this field, you can determine the number of draft invoices that will be affected if the plan is deleted.
	DraftInvoicesCount int32 `json:"draft_invoices_count"`
	// Additional usage-based charges for this plan.
	Charges []ChargeObject `json:"charges,omitempty"`
	// All taxes applied to the plan.
	Taxes []TaxObject `json:"taxes,omitempty"`
}

type _PlanObject PlanObject

// NewPlanObject instantiates a new PlanObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlanObject(lagoId string, name string, createdAt time.Time, code string, interval string, amountCents int32, amountCurrency Currency, activeSubscriptionsCount int32, draftInvoicesCount int32) *PlanObject {
	this := PlanObject{}
	this.LagoId = lagoId
	this.Name = name
	this.CreatedAt = createdAt
	this.Code = code
	this.Interval = interval
	this.AmountCents = amountCents
	this.AmountCurrency = amountCurrency
	this.ActiveSubscriptionsCount = activeSubscriptionsCount
	this.DraftInvoicesCount = draftInvoicesCount
	return &this
}

// NewPlanObjectWithDefaults instantiates a new PlanObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlanObjectWithDefaults() *PlanObject {
	this := PlanObject{}
	return &this
}

// GetLagoId returns the LagoId field value
func (o *PlanObject) GetLagoId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LagoId
}

// GetLagoIdOk returns a tuple with the LagoId field value
// and a boolean to check if the value has been set.
func (o *PlanObject) GetLagoIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LagoId, true
}

// SetLagoId sets field value
func (o *PlanObject) SetLagoId(v string) {
	o.LagoId = v
}

// GetName returns the Name field value
func (o *PlanObject) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PlanObject) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PlanObject) SetName(v string) {
	o.Name = v
}

// GetInvoiceDisplayName returns the InvoiceDisplayName field value if set, zero value otherwise.
func (o *PlanObject) GetInvoiceDisplayName() string {
	if o == nil || IsNil(o.InvoiceDisplayName) {
		var ret string
		return ret
	}
	return *o.InvoiceDisplayName
}

// GetInvoiceDisplayNameOk returns a tuple with the InvoiceDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlanObject) GetInvoiceDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceDisplayName) {
		return nil, false
	}
	return o.InvoiceDisplayName, true
}

// HasInvoiceDisplayName returns a boolean if a field has been set.
func (o *PlanObject) HasInvoiceDisplayName() bool {
	if o != nil && !IsNil(o.InvoiceDisplayName) {
		return true
	}

	return false
}

// SetInvoiceDisplayName gets a reference to the given string and assigns it to the InvoiceDisplayName field.
func (o *PlanObject) SetInvoiceDisplayName(v string) {
	o.InvoiceDisplayName = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *PlanObject) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *PlanObject) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *PlanObject) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCode returns the Code field value
func (o *PlanObject) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *PlanObject) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *PlanObject) SetCode(v string) {
	o.Code = v
}

// GetInterval returns the Interval field value
func (o *PlanObject) GetInterval() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value
// and a boolean to check if the value has been set.
func (o *PlanObject) GetIntervalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interval, true
}

// SetInterval sets field value
func (o *PlanObject) SetInterval(v string) {
	o.Interval = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PlanObject) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlanObject) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PlanObject) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PlanObject) SetDescription(v string) {
	o.Description = &v
}

// GetAmountCents returns the AmountCents field value
func (o *PlanObject) GetAmountCents() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AmountCents
}

// GetAmountCentsOk returns a tuple with the AmountCents field value
// and a boolean to check if the value has been set.
func (o *PlanObject) GetAmountCentsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountCents, true
}

// SetAmountCents sets field value
func (o *PlanObject) SetAmountCents(v int32) {
	o.AmountCents = v
}

// GetAmountCurrency returns the AmountCurrency field value
func (o *PlanObject) GetAmountCurrency() Currency {
	if o == nil {
		var ret Currency
		return ret
	}

	return o.AmountCurrency
}

// GetAmountCurrencyOk returns a tuple with the AmountCurrency field value
// and a boolean to check if the value has been set.
func (o *PlanObject) GetAmountCurrencyOk() (*Currency, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountCurrency, true
}

// SetAmountCurrency sets field value
func (o *PlanObject) SetAmountCurrency(v Currency) {
	o.AmountCurrency = v
}

// GetTrialPeriod returns the TrialPeriod field value if set, zero value otherwise.
func (o *PlanObject) GetTrialPeriod() float32 {
	if o == nil || IsNil(o.TrialPeriod) {
		var ret float32
		return ret
	}
	return *o.TrialPeriod
}

// GetTrialPeriodOk returns a tuple with the TrialPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlanObject) GetTrialPeriodOk() (*float32, bool) {
	if o == nil || IsNil(o.TrialPeriod) {
		return nil, false
	}
	return o.TrialPeriod, true
}

// HasTrialPeriod returns a boolean if a field has been set.
func (o *PlanObject) HasTrialPeriod() bool {
	if o != nil && !IsNil(o.TrialPeriod) {
		return true
	}

	return false
}

// SetTrialPeriod gets a reference to the given float32 and assigns it to the TrialPeriod field.
func (o *PlanObject) SetTrialPeriod(v float32) {
	o.TrialPeriod = &v
}

// GetPayInAdvance returns the PayInAdvance field value if set, zero value otherwise.
func (o *PlanObject) GetPayInAdvance() bool {
	if o == nil || IsNil(o.PayInAdvance) {
		var ret bool
		return ret
	}
	return *o.PayInAdvance
}

// GetPayInAdvanceOk returns a tuple with the PayInAdvance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlanObject) GetPayInAdvanceOk() (*bool, bool) {
	if o == nil || IsNil(o.PayInAdvance) {
		return nil, false
	}
	return o.PayInAdvance, true
}

// HasPayInAdvance returns a boolean if a field has been set.
func (o *PlanObject) HasPayInAdvance() bool {
	if o != nil && !IsNil(o.PayInAdvance) {
		return true
	}

	return false
}

// SetPayInAdvance gets a reference to the given bool and assigns it to the PayInAdvance field.
func (o *PlanObject) SetPayInAdvance(v bool) {
	o.PayInAdvance = &v
}

// GetBillChargesMonthly returns the BillChargesMonthly field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlanObject) GetBillChargesMonthly() bool {
	if o == nil || IsNil(o.BillChargesMonthly.Get()) {
		var ret bool
		return ret
	}
	return *o.BillChargesMonthly.Get()
}

// GetBillChargesMonthlyOk returns a tuple with the BillChargesMonthly field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlanObject) GetBillChargesMonthlyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.BillChargesMonthly.Get(), o.BillChargesMonthly.IsSet()
}

// HasBillChargesMonthly returns a boolean if a field has been set.
func (o *PlanObject) HasBillChargesMonthly() bool {
	if o != nil && o.BillChargesMonthly.IsSet() {
		return true
	}

	return false
}

// SetBillChargesMonthly gets a reference to the given NullableBool and assigns it to the BillChargesMonthly field.
func (o *PlanObject) SetBillChargesMonthly(v bool) {
	o.BillChargesMonthly.Set(&v)
}
// SetBillChargesMonthlyNil sets the value for BillChargesMonthly to be an explicit nil
func (o *PlanObject) SetBillChargesMonthlyNil() {
	o.BillChargesMonthly.Set(nil)
}

// UnsetBillChargesMonthly ensures that no value is present for BillChargesMonthly, not even an explicit nil
func (o *PlanObject) UnsetBillChargesMonthly() {
	o.BillChargesMonthly.Unset()
}

// GetActiveSubscriptionsCount returns the ActiveSubscriptionsCount field value
func (o *PlanObject) GetActiveSubscriptionsCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ActiveSubscriptionsCount
}

// GetActiveSubscriptionsCountOk returns a tuple with the ActiveSubscriptionsCount field value
// and a boolean to check if the value has been set.
func (o *PlanObject) GetActiveSubscriptionsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActiveSubscriptionsCount, true
}

// SetActiveSubscriptionsCount sets field value
func (o *PlanObject) SetActiveSubscriptionsCount(v int32) {
	o.ActiveSubscriptionsCount = v
}

// GetDraftInvoicesCount returns the DraftInvoicesCount field value
func (o *PlanObject) GetDraftInvoicesCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DraftInvoicesCount
}

// GetDraftInvoicesCountOk returns a tuple with the DraftInvoicesCount field value
// and a boolean to check if the value has been set.
func (o *PlanObject) GetDraftInvoicesCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DraftInvoicesCount, true
}

// SetDraftInvoicesCount sets field value
func (o *PlanObject) SetDraftInvoicesCount(v int32) {
	o.DraftInvoicesCount = v
}

// GetCharges returns the Charges field value if set, zero value otherwise.
func (o *PlanObject) GetCharges() []ChargeObject {
	if o == nil || IsNil(o.Charges) {
		var ret []ChargeObject
		return ret
	}
	return o.Charges
}

// GetChargesOk returns a tuple with the Charges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlanObject) GetChargesOk() ([]ChargeObject, bool) {
	if o == nil || IsNil(o.Charges) {
		return nil, false
	}
	return o.Charges, true
}

// HasCharges returns a boolean if a field has been set.
func (o *PlanObject) HasCharges() bool {
	if o != nil && !IsNil(o.Charges) {
		return true
	}

	return false
}

// SetCharges gets a reference to the given []ChargeObject and assigns it to the Charges field.
func (o *PlanObject) SetCharges(v []ChargeObject) {
	o.Charges = v
}

// GetTaxes returns the Taxes field value if set, zero value otherwise.
func (o *PlanObject) GetTaxes() []TaxObject {
	if o == nil || IsNil(o.Taxes) {
		var ret []TaxObject
		return ret
	}
	return o.Taxes
}

// GetTaxesOk returns a tuple with the Taxes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlanObject) GetTaxesOk() ([]TaxObject, bool) {
	if o == nil || IsNil(o.Taxes) {
		return nil, false
	}
	return o.Taxes, true
}

// HasTaxes returns a boolean if a field has been set.
func (o *PlanObject) HasTaxes() bool {
	if o != nil && !IsNil(o.Taxes) {
		return true
	}

	return false
}

// SetTaxes gets a reference to the given []TaxObject and assigns it to the Taxes field.
func (o *PlanObject) SetTaxes(v []TaxObject) {
	o.Taxes = v
}

func (o PlanObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlanObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["lago_id"] = o.LagoId
	toSerialize["name"] = o.Name
	if !IsNil(o.InvoiceDisplayName) {
		toSerialize["invoice_display_name"] = o.InvoiceDisplayName
	}
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["code"] = o.Code
	toSerialize["interval"] = o.Interval
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["amount_cents"] = o.AmountCents
	toSerialize["amount_currency"] = o.AmountCurrency
	if !IsNil(o.TrialPeriod) {
		toSerialize["trial_period"] = o.TrialPeriod
	}
	if !IsNil(o.PayInAdvance) {
		toSerialize["pay_in_advance"] = o.PayInAdvance
	}
	if o.BillChargesMonthly.IsSet() {
		toSerialize["bill_charges_monthly"] = o.BillChargesMonthly.Get()
	}
	toSerialize["active_subscriptions_count"] = o.ActiveSubscriptionsCount
	toSerialize["draft_invoices_count"] = o.DraftInvoicesCount
	if !IsNil(o.Charges) {
		toSerialize["charges"] = o.Charges
	}
	if !IsNil(o.Taxes) {
		toSerialize["taxes"] = o.Taxes
	}
	return toSerialize, nil
}

func (o *PlanObject) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"lago_id",
		"name",
		"created_at",
		"code",
		"interval",
		"amount_cents",
		"amount_currency",
		"active_subscriptions_count",
		"draft_invoices_count",
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

	varPlanObject := _PlanObject{}

	err = json.Unmarshal(bytes, &varPlanObject)

	if err != nil {
		return err
	}

	*o = PlanObject(varPlanObject)

	return err
}

type NullablePlanObject struct {
	value *PlanObject
	isSet bool
}

func (v NullablePlanObject) Get() *PlanObject {
	return v.value
}

func (v *NullablePlanObject) Set(val *PlanObject) {
	v.value = val
	v.isSet = true
}

func (v NullablePlanObject) IsSet() bool {
	return v.isSet
}

func (v *NullablePlanObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlanObject(val *PlanObject) *NullablePlanObject {
	return &NullablePlanObject{value: val, isSet: true}
}

func (v NullablePlanObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlanObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


