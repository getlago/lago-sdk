# PlanUpdateInputPlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the plan. | [optional] 
**InvoiceDisplayName** | Pointer to **string** | Specifies the name that will be displayed on an invoice. If no value is set for this field, the name of the plan will be used as the default display name. | [optional] 
**Code** | Pointer to **string** | The code of the plan. It serves as a unique identifier associated with a particular plan. The code is typically used for internal or system-level identification purposes, like assigning a subscription, for instance. | [optional] 
**Interval** | Pointer to **string** | The interval used for recurring billing. It represents the frequency at which subscription billing occurs. The interval can be one of the following values: &#x60;yearly&#x60;, &#x60;quarterly&#x60;, &#x60;monthly&#x60;, or &#x60;weekly&#x60;. | [optional] 
**Description** | Pointer to **string** | The description on the plan. | [optional] 
**AmountCents** | Pointer to **int32** | The base cost of the plan, excluding any applicable taxes, that is billed on a recurring basis. This value is defined at 0 if your plan is a pay-as-you-go plan. | [optional] 
**AmountCurrency** | Pointer to [**Currency**](Currency.md) |  | [optional] 
**TrialPeriod** | Pointer to **float32** | The duration in days during which the base cost of the plan is offered for free. | [optional] 
**PayInAdvance** | Pointer to **bool** | This field determines the billing timing for the plan. When set to &#x60;true&#x60;, the base cost of the plan is due at the beginning of each billing period. Conversely, when set to &#x60;false&#x60;, the base cost of the plan is due at the end of each billing period. | [optional] 
**BillChargesMonthly** | Pointer to **NullableBool** | This field, when set to &#x60;true&#x60;, enables to invoice usage-based charges on monthly basis, even if the cadence of the plan is yearly. This allows customers to pay charges overage on a monthly basis. This can be set to true only if the plan’s interval is &#x60;yearly&#x60;. | [optional] 
**TaxCodes** | Pointer to **[]string** | List of unique code used to identify the taxes. | [optional] 
**Charges** | Pointer to [**[]PlanUpdateInputPlanChargesInner**](PlanUpdateInputPlanChargesInner.md) | Additional usage-based charges for this plan. | [optional] 

## Methods

### NewPlanUpdateInputPlan

`func NewPlanUpdateInputPlan() *PlanUpdateInputPlan`

NewPlanUpdateInputPlan instantiates a new PlanUpdateInputPlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanUpdateInputPlanWithDefaults

`func NewPlanUpdateInputPlanWithDefaults() *PlanUpdateInputPlan`

NewPlanUpdateInputPlanWithDefaults instantiates a new PlanUpdateInputPlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PlanUpdateInputPlan) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlanUpdateInputPlan) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlanUpdateInputPlan) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PlanUpdateInputPlan) HasName() bool`

HasName returns a boolean if a field has been set.

### GetInvoiceDisplayName

`func (o *PlanUpdateInputPlan) GetInvoiceDisplayName() string`

GetInvoiceDisplayName returns the InvoiceDisplayName field if non-nil, zero value otherwise.

### GetInvoiceDisplayNameOk

`func (o *PlanUpdateInputPlan) GetInvoiceDisplayNameOk() (*string, bool)`

GetInvoiceDisplayNameOk returns a tuple with the InvoiceDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceDisplayName

`func (o *PlanUpdateInputPlan) SetInvoiceDisplayName(v string)`

SetInvoiceDisplayName sets InvoiceDisplayName field to given value.

### HasInvoiceDisplayName

`func (o *PlanUpdateInputPlan) HasInvoiceDisplayName() bool`

HasInvoiceDisplayName returns a boolean if a field has been set.

### GetCode

`func (o *PlanUpdateInputPlan) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PlanUpdateInputPlan) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PlanUpdateInputPlan) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *PlanUpdateInputPlan) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetInterval

`func (o *PlanUpdateInputPlan) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *PlanUpdateInputPlan) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *PlanUpdateInputPlan) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *PlanUpdateInputPlan) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetDescription

`func (o *PlanUpdateInputPlan) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PlanUpdateInputPlan) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PlanUpdateInputPlan) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PlanUpdateInputPlan) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAmountCents

`func (o *PlanUpdateInputPlan) GetAmountCents() int32`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *PlanUpdateInputPlan) GetAmountCentsOk() (*int32, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *PlanUpdateInputPlan) SetAmountCents(v int32)`

SetAmountCents sets AmountCents field to given value.

### HasAmountCents

`func (o *PlanUpdateInputPlan) HasAmountCents() bool`

HasAmountCents returns a boolean if a field has been set.

### GetAmountCurrency

`func (o *PlanUpdateInputPlan) GetAmountCurrency() Currency`

GetAmountCurrency returns the AmountCurrency field if non-nil, zero value otherwise.

### GetAmountCurrencyOk

`func (o *PlanUpdateInputPlan) GetAmountCurrencyOk() (*Currency, bool)`

GetAmountCurrencyOk returns a tuple with the AmountCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCurrency

`func (o *PlanUpdateInputPlan) SetAmountCurrency(v Currency)`

SetAmountCurrency sets AmountCurrency field to given value.

### HasAmountCurrency

`func (o *PlanUpdateInputPlan) HasAmountCurrency() bool`

HasAmountCurrency returns a boolean if a field has been set.

### GetTrialPeriod

`func (o *PlanUpdateInputPlan) GetTrialPeriod() float32`

GetTrialPeriod returns the TrialPeriod field if non-nil, zero value otherwise.

### GetTrialPeriodOk

`func (o *PlanUpdateInputPlan) GetTrialPeriodOk() (*float32, bool)`

GetTrialPeriodOk returns a tuple with the TrialPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialPeriod

`func (o *PlanUpdateInputPlan) SetTrialPeriod(v float32)`

SetTrialPeriod sets TrialPeriod field to given value.

### HasTrialPeriod

`func (o *PlanUpdateInputPlan) HasTrialPeriod() bool`

HasTrialPeriod returns a boolean if a field has been set.

### GetPayInAdvance

`func (o *PlanUpdateInputPlan) GetPayInAdvance() bool`

GetPayInAdvance returns the PayInAdvance field if non-nil, zero value otherwise.

### GetPayInAdvanceOk

`func (o *PlanUpdateInputPlan) GetPayInAdvanceOk() (*bool, bool)`

GetPayInAdvanceOk returns a tuple with the PayInAdvance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayInAdvance

`func (o *PlanUpdateInputPlan) SetPayInAdvance(v bool)`

SetPayInAdvance sets PayInAdvance field to given value.

### HasPayInAdvance

`func (o *PlanUpdateInputPlan) HasPayInAdvance() bool`

HasPayInAdvance returns a boolean if a field has been set.

### GetBillChargesMonthly

`func (o *PlanUpdateInputPlan) GetBillChargesMonthly() bool`

GetBillChargesMonthly returns the BillChargesMonthly field if non-nil, zero value otherwise.

### GetBillChargesMonthlyOk

`func (o *PlanUpdateInputPlan) GetBillChargesMonthlyOk() (*bool, bool)`

GetBillChargesMonthlyOk returns a tuple with the BillChargesMonthly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillChargesMonthly

`func (o *PlanUpdateInputPlan) SetBillChargesMonthly(v bool)`

SetBillChargesMonthly sets BillChargesMonthly field to given value.

### HasBillChargesMonthly

`func (o *PlanUpdateInputPlan) HasBillChargesMonthly() bool`

HasBillChargesMonthly returns a boolean if a field has been set.

### SetBillChargesMonthlyNil

`func (o *PlanUpdateInputPlan) SetBillChargesMonthlyNil(b bool)`

 SetBillChargesMonthlyNil sets the value for BillChargesMonthly to be an explicit nil

### UnsetBillChargesMonthly
`func (o *PlanUpdateInputPlan) UnsetBillChargesMonthly()`

UnsetBillChargesMonthly ensures that no value is present for BillChargesMonthly, not even an explicit nil
### GetTaxCodes

`func (o *PlanUpdateInputPlan) GetTaxCodes() []string`

GetTaxCodes returns the TaxCodes field if non-nil, zero value otherwise.

### GetTaxCodesOk

`func (o *PlanUpdateInputPlan) GetTaxCodesOk() (*[]string, bool)`

GetTaxCodesOk returns a tuple with the TaxCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCodes

`func (o *PlanUpdateInputPlan) SetTaxCodes(v []string)`

SetTaxCodes sets TaxCodes field to given value.

### HasTaxCodes

`func (o *PlanUpdateInputPlan) HasTaxCodes() bool`

HasTaxCodes returns a boolean if a field has been set.

### GetCharges

`func (o *PlanUpdateInputPlan) GetCharges() []PlanUpdateInputPlanChargesInner`

GetCharges returns the Charges field if non-nil, zero value otherwise.

### GetChargesOk

`func (o *PlanUpdateInputPlan) GetChargesOk() (*[]PlanUpdateInputPlanChargesInner, bool)`

GetChargesOk returns a tuple with the Charges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharges

`func (o *PlanUpdateInputPlan) SetCharges(v []PlanUpdateInputPlanChargesInner)`

SetCharges sets Charges field to given value.

### HasCharges

`func (o *PlanUpdateInputPlan) HasCharges() bool`

HasCharges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


