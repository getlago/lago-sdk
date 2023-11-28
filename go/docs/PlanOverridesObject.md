# PlanOverridesObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountCents** | Pointer to **int32** | The base cost of the plan, excluding any applicable taxes, that is billed on a recurring basis. This value is defined at 0 if your plan is a pay-as-you-go plan. | [optional] 
**AmountCurrency** | Pointer to [**Currency**](Currency.md) |  | [optional] 
**Description** | Pointer to **string** | The description on the plan. | [optional] 
**InvoiceDisplayName** | Pointer to **string** | Specifies the name that will be displayed on an invoice. If no value is set for this field, the name of the plan will be used as the default display name. | [optional] 
**Name** | Pointer to **string** | The name of the plan. | [optional] 
**TaxCodes** | Pointer to **[]string** | List of unique code used to identify the taxes. | [optional] 
**TrialPeriod** | Pointer to **float32** | The duration in days during which the base cost of the plan is offered for free. | [optional] 
**Charges** | Pointer to [**[]PlanOverridesObjectChargesInner**](PlanOverridesObjectChargesInner.md) | Additional usage-based charges for this plan. | [optional] 

## Methods

### NewPlanOverridesObject

`func NewPlanOverridesObject() *PlanOverridesObject`

NewPlanOverridesObject instantiates a new PlanOverridesObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanOverridesObjectWithDefaults

`func NewPlanOverridesObjectWithDefaults() *PlanOverridesObject`

NewPlanOverridesObjectWithDefaults instantiates a new PlanOverridesObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountCents

`func (o *PlanOverridesObject) GetAmountCents() int32`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *PlanOverridesObject) GetAmountCentsOk() (*int32, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *PlanOverridesObject) SetAmountCents(v int32)`

SetAmountCents sets AmountCents field to given value.

### HasAmountCents

`func (o *PlanOverridesObject) HasAmountCents() bool`

HasAmountCents returns a boolean if a field has been set.

### GetAmountCurrency

`func (o *PlanOverridesObject) GetAmountCurrency() Currency`

GetAmountCurrency returns the AmountCurrency field if non-nil, zero value otherwise.

### GetAmountCurrencyOk

`func (o *PlanOverridesObject) GetAmountCurrencyOk() (*Currency, bool)`

GetAmountCurrencyOk returns a tuple with the AmountCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCurrency

`func (o *PlanOverridesObject) SetAmountCurrency(v Currency)`

SetAmountCurrency sets AmountCurrency field to given value.

### HasAmountCurrency

`func (o *PlanOverridesObject) HasAmountCurrency() bool`

HasAmountCurrency returns a boolean if a field has been set.

### GetDescription

`func (o *PlanOverridesObject) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PlanOverridesObject) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PlanOverridesObject) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PlanOverridesObject) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInvoiceDisplayName

`func (o *PlanOverridesObject) GetInvoiceDisplayName() string`

GetInvoiceDisplayName returns the InvoiceDisplayName field if non-nil, zero value otherwise.

### GetInvoiceDisplayNameOk

`func (o *PlanOverridesObject) GetInvoiceDisplayNameOk() (*string, bool)`

GetInvoiceDisplayNameOk returns a tuple with the InvoiceDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceDisplayName

`func (o *PlanOverridesObject) SetInvoiceDisplayName(v string)`

SetInvoiceDisplayName sets InvoiceDisplayName field to given value.

### HasInvoiceDisplayName

`func (o *PlanOverridesObject) HasInvoiceDisplayName() bool`

HasInvoiceDisplayName returns a boolean if a field has been set.

### GetName

`func (o *PlanOverridesObject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlanOverridesObject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlanOverridesObject) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PlanOverridesObject) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTaxCodes

`func (o *PlanOverridesObject) GetTaxCodes() []string`

GetTaxCodes returns the TaxCodes field if non-nil, zero value otherwise.

### GetTaxCodesOk

`func (o *PlanOverridesObject) GetTaxCodesOk() (*[]string, bool)`

GetTaxCodesOk returns a tuple with the TaxCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCodes

`func (o *PlanOverridesObject) SetTaxCodes(v []string)`

SetTaxCodes sets TaxCodes field to given value.

### HasTaxCodes

`func (o *PlanOverridesObject) HasTaxCodes() bool`

HasTaxCodes returns a boolean if a field has been set.

### GetTrialPeriod

`func (o *PlanOverridesObject) GetTrialPeriod() float32`

GetTrialPeriod returns the TrialPeriod field if non-nil, zero value otherwise.

### GetTrialPeriodOk

`func (o *PlanOverridesObject) GetTrialPeriodOk() (*float32, bool)`

GetTrialPeriodOk returns a tuple with the TrialPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialPeriod

`func (o *PlanOverridesObject) SetTrialPeriod(v float32)`

SetTrialPeriod sets TrialPeriod field to given value.

### HasTrialPeriod

`func (o *PlanOverridesObject) HasTrialPeriod() bool`

HasTrialPeriod returns a boolean if a field has been set.

### GetCharges

`func (o *PlanOverridesObject) GetCharges() []PlanOverridesObjectChargesInner`

GetCharges returns the Charges field if non-nil, zero value otherwise.

### GetChargesOk

`func (o *PlanOverridesObject) GetChargesOk() (*[]PlanOverridesObjectChargesInner, bool)`

GetChargesOk returns a tuple with the Charges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharges

`func (o *PlanOverridesObject) SetCharges(v []PlanOverridesObjectChargesInner)`

SetCharges sets Charges field to given value.

### HasCharges

`func (o *PlanOverridesObject) HasCharges() bool`

HasCharges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


