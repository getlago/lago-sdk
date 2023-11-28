# PlanCreateInputPlan

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
**Charges** | Pointer to [**[]PlanCreateInputPlanChargesInner**](PlanCreateInputPlanChargesInner.md) | Additional usage-based charges for this plan. | [optional] 

## Methods

### NewPlanCreateInputPlan

`func NewPlanCreateInputPlan() *PlanCreateInputPlan`

NewPlanCreateInputPlan instantiates a new PlanCreateInputPlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanCreateInputPlanWithDefaults

`func NewPlanCreateInputPlanWithDefaults() *PlanCreateInputPlan`

NewPlanCreateInputPlanWithDefaults instantiates a new PlanCreateInputPlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PlanCreateInputPlan) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlanCreateInputPlan) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlanCreateInputPlan) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PlanCreateInputPlan) HasName() bool`

HasName returns a boolean if a field has been set.

### GetInvoiceDisplayName

`func (o *PlanCreateInputPlan) GetInvoiceDisplayName() string`

GetInvoiceDisplayName returns the InvoiceDisplayName field if non-nil, zero value otherwise.

### GetInvoiceDisplayNameOk

`func (o *PlanCreateInputPlan) GetInvoiceDisplayNameOk() (*string, bool)`

GetInvoiceDisplayNameOk returns a tuple with the InvoiceDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceDisplayName

`func (o *PlanCreateInputPlan) SetInvoiceDisplayName(v string)`

SetInvoiceDisplayName sets InvoiceDisplayName field to given value.

### HasInvoiceDisplayName

`func (o *PlanCreateInputPlan) HasInvoiceDisplayName() bool`

HasInvoiceDisplayName returns a boolean if a field has been set.

### GetCode

`func (o *PlanCreateInputPlan) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PlanCreateInputPlan) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PlanCreateInputPlan) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *PlanCreateInputPlan) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetInterval

`func (o *PlanCreateInputPlan) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *PlanCreateInputPlan) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *PlanCreateInputPlan) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *PlanCreateInputPlan) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetDescription

`func (o *PlanCreateInputPlan) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PlanCreateInputPlan) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PlanCreateInputPlan) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PlanCreateInputPlan) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAmountCents

`func (o *PlanCreateInputPlan) GetAmountCents() int32`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *PlanCreateInputPlan) GetAmountCentsOk() (*int32, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *PlanCreateInputPlan) SetAmountCents(v int32)`

SetAmountCents sets AmountCents field to given value.

### HasAmountCents

`func (o *PlanCreateInputPlan) HasAmountCents() bool`

HasAmountCents returns a boolean if a field has been set.

### GetAmountCurrency

`func (o *PlanCreateInputPlan) GetAmountCurrency() Currency`

GetAmountCurrency returns the AmountCurrency field if non-nil, zero value otherwise.

### GetAmountCurrencyOk

`func (o *PlanCreateInputPlan) GetAmountCurrencyOk() (*Currency, bool)`

GetAmountCurrencyOk returns a tuple with the AmountCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCurrency

`func (o *PlanCreateInputPlan) SetAmountCurrency(v Currency)`

SetAmountCurrency sets AmountCurrency field to given value.

### HasAmountCurrency

`func (o *PlanCreateInputPlan) HasAmountCurrency() bool`

HasAmountCurrency returns a boolean if a field has been set.

### GetTrialPeriod

`func (o *PlanCreateInputPlan) GetTrialPeriod() float32`

GetTrialPeriod returns the TrialPeriod field if non-nil, zero value otherwise.

### GetTrialPeriodOk

`func (o *PlanCreateInputPlan) GetTrialPeriodOk() (*float32, bool)`

GetTrialPeriodOk returns a tuple with the TrialPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialPeriod

`func (o *PlanCreateInputPlan) SetTrialPeriod(v float32)`

SetTrialPeriod sets TrialPeriod field to given value.

### HasTrialPeriod

`func (o *PlanCreateInputPlan) HasTrialPeriod() bool`

HasTrialPeriod returns a boolean if a field has been set.

### GetPayInAdvance

`func (o *PlanCreateInputPlan) GetPayInAdvance() bool`

GetPayInAdvance returns the PayInAdvance field if non-nil, zero value otherwise.

### GetPayInAdvanceOk

`func (o *PlanCreateInputPlan) GetPayInAdvanceOk() (*bool, bool)`

GetPayInAdvanceOk returns a tuple with the PayInAdvance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayInAdvance

`func (o *PlanCreateInputPlan) SetPayInAdvance(v bool)`

SetPayInAdvance sets PayInAdvance field to given value.

### HasPayInAdvance

`func (o *PlanCreateInputPlan) HasPayInAdvance() bool`

HasPayInAdvance returns a boolean if a field has been set.

### GetBillChargesMonthly

`func (o *PlanCreateInputPlan) GetBillChargesMonthly() bool`

GetBillChargesMonthly returns the BillChargesMonthly field if non-nil, zero value otherwise.

### GetBillChargesMonthlyOk

`func (o *PlanCreateInputPlan) GetBillChargesMonthlyOk() (*bool, bool)`

GetBillChargesMonthlyOk returns a tuple with the BillChargesMonthly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillChargesMonthly

`func (o *PlanCreateInputPlan) SetBillChargesMonthly(v bool)`

SetBillChargesMonthly sets BillChargesMonthly field to given value.

### HasBillChargesMonthly

`func (o *PlanCreateInputPlan) HasBillChargesMonthly() bool`

HasBillChargesMonthly returns a boolean if a field has been set.

### SetBillChargesMonthlyNil

`func (o *PlanCreateInputPlan) SetBillChargesMonthlyNil(b bool)`

 SetBillChargesMonthlyNil sets the value for BillChargesMonthly to be an explicit nil

### UnsetBillChargesMonthly
`func (o *PlanCreateInputPlan) UnsetBillChargesMonthly()`

UnsetBillChargesMonthly ensures that no value is present for BillChargesMonthly, not even an explicit nil
### GetTaxCodes

`func (o *PlanCreateInputPlan) GetTaxCodes() []string`

GetTaxCodes returns the TaxCodes field if non-nil, zero value otherwise.

### GetTaxCodesOk

`func (o *PlanCreateInputPlan) GetTaxCodesOk() (*[]string, bool)`

GetTaxCodesOk returns a tuple with the TaxCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCodes

`func (o *PlanCreateInputPlan) SetTaxCodes(v []string)`

SetTaxCodes sets TaxCodes field to given value.

### HasTaxCodes

`func (o *PlanCreateInputPlan) HasTaxCodes() bool`

HasTaxCodes returns a boolean if a field has been set.

### GetCharges

`func (o *PlanCreateInputPlan) GetCharges() []PlanCreateInputPlanChargesInner`

GetCharges returns the Charges field if non-nil, zero value otherwise.

### GetChargesOk

`func (o *PlanCreateInputPlan) GetChargesOk() (*[]PlanCreateInputPlanChargesInner, bool)`

GetChargesOk returns a tuple with the Charges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharges

`func (o *PlanCreateInputPlan) SetCharges(v []PlanCreateInputPlanChargesInner)`

SetCharges sets Charges field to given value.

### HasCharges

`func (o *PlanCreateInputPlan) HasCharges() bool`

HasCharges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


