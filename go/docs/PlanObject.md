# PlanObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LagoId** | **string** | Unique identifier of the plan created by Lago. | 
**Name** | **string** | The name of the plan. | 
**InvoiceDisplayName** | Pointer to **string** | Specifies the name that will be displayed on an invoice. If no value is set for this field, the name of the plan will be used as the default display name. | [optional] 
**CreatedAt** | **time.Time** | The date and time when the plan was created. It is expressed in UTC format according to the ISO 8601 datetime standard. This field provides the timestamp for the exact moment when the plan was initially created. | 
**Code** | **string** | The code of the plan. It serves as a unique identifier associated with a particular plan. The code is typically used for internal or system-level identification purposes, like assigning a subscription, for instance. | 
**Interval** | **string** | The interval used for recurring billing. It represents the frequency at which subscription billing occurs. The interval can be one of the following values: &#x60;yearly&#x60;, &#x60;quarterly&#x60;, &#x60;monthly&#x60; or &#x60;weekly&#x60;. | 
**Description** | Pointer to **string** | The description on the plan. | [optional] 
**AmountCents** | **int32** | The base cost of the plan, excluding any applicable taxes, that is billed on a recurring basis. This value is defined at 0 if your plan is a pay-as-you-go plan. | 
**AmountCurrency** | [**Currency**](Currency.md) |  | 
**TrialPeriod** | Pointer to **float32** | The duration in days during which the base cost of the plan is offered for free. | [optional] 
**PayInAdvance** | Pointer to **bool** | This field determines the billing timing for the plan. When set to &#x60;true&#x60;, the base cost of the plan is due at the beginning of each billing period. Conversely, when set to &#x60;false&#x60;, the base cost of the plan is due at the end of each billing period. | [optional] 
**BillChargesMonthly** | Pointer to **NullableBool** | This field, when set to &#x60;true&#x60;, enables to invoice usage-based charges on monthly basis, even if the cadence of the plan is yearly. This allows customers to pay charges overage on a monthly basis. This can be set to true only if the plan’s interval is &#x60;yearly&#x60;. | [optional] 
**ActiveSubscriptionsCount** | **int32** | The count of active subscriptions that are currently associated with the plan. This field provides valuable information regarding the impact of deleting the plan. By checking the value of this field, you can determine the number of subscriptions that will be affected if the plan is deleted. | 
**DraftInvoicesCount** | **int32** | The number of draft invoices that include a subscription attached to the plan. This field provides valuable information about the impact of deleting the plan. By checking the value of this field, you can determine the number of draft invoices that will be affected if the plan is deleted. | 
**Charges** | Pointer to [**[]ChargeObject**](ChargeObject.md) | Additional usage-based charges for this plan. | [optional] 
**Taxes** | Pointer to [**[]TaxObject**](TaxObject.md) | All taxes applied to the plan. | [optional] 

## Methods

### NewPlanObject

`func NewPlanObject(lagoId string, name string, createdAt time.Time, code string, interval string, amountCents int32, amountCurrency Currency, activeSubscriptionsCount int32, draftInvoicesCount int32, ) *PlanObject`

NewPlanObject instantiates a new PlanObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanObjectWithDefaults

`func NewPlanObjectWithDefaults() *PlanObject`

NewPlanObjectWithDefaults instantiates a new PlanObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLagoId

`func (o *PlanObject) GetLagoId() string`

GetLagoId returns the LagoId field if non-nil, zero value otherwise.

### GetLagoIdOk

`func (o *PlanObject) GetLagoIdOk() (*string, bool)`

GetLagoIdOk returns a tuple with the LagoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagoId

`func (o *PlanObject) SetLagoId(v string)`

SetLagoId sets LagoId field to given value.


### GetName

`func (o *PlanObject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlanObject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlanObject) SetName(v string)`

SetName sets Name field to given value.


### GetInvoiceDisplayName

`func (o *PlanObject) GetInvoiceDisplayName() string`

GetInvoiceDisplayName returns the InvoiceDisplayName field if non-nil, zero value otherwise.

### GetInvoiceDisplayNameOk

`func (o *PlanObject) GetInvoiceDisplayNameOk() (*string, bool)`

GetInvoiceDisplayNameOk returns a tuple with the InvoiceDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceDisplayName

`func (o *PlanObject) SetInvoiceDisplayName(v string)`

SetInvoiceDisplayName sets InvoiceDisplayName field to given value.

### HasInvoiceDisplayName

`func (o *PlanObject) HasInvoiceDisplayName() bool`

HasInvoiceDisplayName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PlanObject) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PlanObject) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PlanObject) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCode

`func (o *PlanObject) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PlanObject) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PlanObject) SetCode(v string)`

SetCode sets Code field to given value.


### GetInterval

`func (o *PlanObject) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *PlanObject) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *PlanObject) SetInterval(v string)`

SetInterval sets Interval field to given value.


### GetDescription

`func (o *PlanObject) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PlanObject) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PlanObject) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PlanObject) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAmountCents

`func (o *PlanObject) GetAmountCents() int32`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *PlanObject) GetAmountCentsOk() (*int32, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *PlanObject) SetAmountCents(v int32)`

SetAmountCents sets AmountCents field to given value.


### GetAmountCurrency

`func (o *PlanObject) GetAmountCurrency() Currency`

GetAmountCurrency returns the AmountCurrency field if non-nil, zero value otherwise.

### GetAmountCurrencyOk

`func (o *PlanObject) GetAmountCurrencyOk() (*Currency, bool)`

GetAmountCurrencyOk returns a tuple with the AmountCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCurrency

`func (o *PlanObject) SetAmountCurrency(v Currency)`

SetAmountCurrency sets AmountCurrency field to given value.


### GetTrialPeriod

`func (o *PlanObject) GetTrialPeriod() float32`

GetTrialPeriod returns the TrialPeriod field if non-nil, zero value otherwise.

### GetTrialPeriodOk

`func (o *PlanObject) GetTrialPeriodOk() (*float32, bool)`

GetTrialPeriodOk returns a tuple with the TrialPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialPeriod

`func (o *PlanObject) SetTrialPeriod(v float32)`

SetTrialPeriod sets TrialPeriod field to given value.

### HasTrialPeriod

`func (o *PlanObject) HasTrialPeriod() bool`

HasTrialPeriod returns a boolean if a field has been set.

### GetPayInAdvance

`func (o *PlanObject) GetPayInAdvance() bool`

GetPayInAdvance returns the PayInAdvance field if non-nil, zero value otherwise.

### GetPayInAdvanceOk

`func (o *PlanObject) GetPayInAdvanceOk() (*bool, bool)`

GetPayInAdvanceOk returns a tuple with the PayInAdvance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayInAdvance

`func (o *PlanObject) SetPayInAdvance(v bool)`

SetPayInAdvance sets PayInAdvance field to given value.

### HasPayInAdvance

`func (o *PlanObject) HasPayInAdvance() bool`

HasPayInAdvance returns a boolean if a field has been set.

### GetBillChargesMonthly

`func (o *PlanObject) GetBillChargesMonthly() bool`

GetBillChargesMonthly returns the BillChargesMonthly field if non-nil, zero value otherwise.

### GetBillChargesMonthlyOk

`func (o *PlanObject) GetBillChargesMonthlyOk() (*bool, bool)`

GetBillChargesMonthlyOk returns a tuple with the BillChargesMonthly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillChargesMonthly

`func (o *PlanObject) SetBillChargesMonthly(v bool)`

SetBillChargesMonthly sets BillChargesMonthly field to given value.

### HasBillChargesMonthly

`func (o *PlanObject) HasBillChargesMonthly() bool`

HasBillChargesMonthly returns a boolean if a field has been set.

### SetBillChargesMonthlyNil

`func (o *PlanObject) SetBillChargesMonthlyNil(b bool)`

 SetBillChargesMonthlyNil sets the value for BillChargesMonthly to be an explicit nil

### UnsetBillChargesMonthly
`func (o *PlanObject) UnsetBillChargesMonthly()`

UnsetBillChargesMonthly ensures that no value is present for BillChargesMonthly, not even an explicit nil
### GetActiveSubscriptionsCount

`func (o *PlanObject) GetActiveSubscriptionsCount() int32`

GetActiveSubscriptionsCount returns the ActiveSubscriptionsCount field if non-nil, zero value otherwise.

### GetActiveSubscriptionsCountOk

`func (o *PlanObject) GetActiveSubscriptionsCountOk() (*int32, bool)`

GetActiveSubscriptionsCountOk returns a tuple with the ActiveSubscriptionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveSubscriptionsCount

`func (o *PlanObject) SetActiveSubscriptionsCount(v int32)`

SetActiveSubscriptionsCount sets ActiveSubscriptionsCount field to given value.


### GetDraftInvoicesCount

`func (o *PlanObject) GetDraftInvoicesCount() int32`

GetDraftInvoicesCount returns the DraftInvoicesCount field if non-nil, zero value otherwise.

### GetDraftInvoicesCountOk

`func (o *PlanObject) GetDraftInvoicesCountOk() (*int32, bool)`

GetDraftInvoicesCountOk returns a tuple with the DraftInvoicesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraftInvoicesCount

`func (o *PlanObject) SetDraftInvoicesCount(v int32)`

SetDraftInvoicesCount sets DraftInvoicesCount field to given value.


### GetCharges

`func (o *PlanObject) GetCharges() []ChargeObject`

GetCharges returns the Charges field if non-nil, zero value otherwise.

### GetChargesOk

`func (o *PlanObject) GetChargesOk() (*[]ChargeObject, bool)`

GetChargesOk returns a tuple with the Charges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharges

`func (o *PlanObject) SetCharges(v []ChargeObject)`

SetCharges sets Charges field to given value.

### HasCharges

`func (o *PlanObject) HasCharges() bool`

HasCharges returns a boolean if a field has been set.

### GetTaxes

`func (o *PlanObject) GetTaxes() []TaxObject`

GetTaxes returns the Taxes field if non-nil, zero value otherwise.

### GetTaxesOk

`func (o *PlanObject) GetTaxesOk() (*[]TaxObject, bool)`

GetTaxesOk returns a tuple with the Taxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxes

`func (o *PlanObject) SetTaxes(v []TaxObject)`

SetTaxes sets Taxes field to given value.

### HasTaxes

`func (o *PlanObject) HasTaxes() bool`

HasTaxes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


