# PlanCreateInputPlanChargesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillableMetricId** | Pointer to **string** | Unique identifier of the billable metric created by Lago. | [optional] 
**ChargeModel** | Pointer to **string** | Specifies the pricing model used for the calculation of the final fee. It can be &#x60;standard&#x60;, &#x60;graduated&#x60;, &#x60;graduated_percentage&#x60; &#x60;package&#x60;, &#x60;percentage&#x60; or &#x60;volume&#x60;. | [optional] 
**PayInAdvance** | Pointer to **bool** | This field determines the billing timing for this specific usage-based charge. When set to &#x60;true&#x60;, the charge is due and invoiced immediately. Conversely, when set to false, the charge is due and invoiced at the end of each billing period. | [optional] 
**Invoiceable** | Pointer to **bool** | This field specifies whether the charge should be included in a proper invoice. If set to false, no invoice will be issued for this charge. You can only set it to &#x60;false&#x60; when &#x60;pay_in_advance&#x60; is &#x60;true&#x60;. | [optional] 
**InvoiceDisplayName** | Pointer to **string** | Specifies the name that will be displayed on an invoice. If no value is set for this field, the name of the actual charge will be used as the default display name. | [optional] 
**Prorated** | Pointer to **bool** | Specifies whether a charge is prorated based on the remaining number of days in the billing period or billed fully.  - If set to &#x60;true&#x60;, the charge is prorated based on the remaining days in the current billing period. - If set to &#x60;false&#x60;, the charge is billed in full. - If not defined in the request, default value is &#x60;false&#x60;. | [optional] 
**MinAmountCents** | Pointer to **int32** | The minimum spending amount required for the charge, measured in cents and excluding any applicable taxes. It indicates the minimum amount that needs to be charged for each billing period. | [optional] 
**Properties** | Pointer to [**ChargeObjectProperties**](ChargeObjectProperties.md) |  | [optional] 
**GroupProperties** | Pointer to [**[]PlanCreateInputPlanChargesInnerGroupPropertiesInner**](PlanCreateInputPlanChargesInnerGroupPropertiesInner.md) | All charge information, sorted by groups. | [optional] 
**TaxCodes** | Pointer to **[]string** | List of unique code used to identify the taxes. | [optional] 

## Methods

### NewPlanCreateInputPlanChargesInner

`func NewPlanCreateInputPlanChargesInner() *PlanCreateInputPlanChargesInner`

NewPlanCreateInputPlanChargesInner instantiates a new PlanCreateInputPlanChargesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanCreateInputPlanChargesInnerWithDefaults

`func NewPlanCreateInputPlanChargesInnerWithDefaults() *PlanCreateInputPlanChargesInner`

NewPlanCreateInputPlanChargesInnerWithDefaults instantiates a new PlanCreateInputPlanChargesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillableMetricId

`func (o *PlanCreateInputPlanChargesInner) GetBillableMetricId() string`

GetBillableMetricId returns the BillableMetricId field if non-nil, zero value otherwise.

### GetBillableMetricIdOk

`func (o *PlanCreateInputPlanChargesInner) GetBillableMetricIdOk() (*string, bool)`

GetBillableMetricIdOk returns a tuple with the BillableMetricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillableMetricId

`func (o *PlanCreateInputPlanChargesInner) SetBillableMetricId(v string)`

SetBillableMetricId sets BillableMetricId field to given value.

### HasBillableMetricId

`func (o *PlanCreateInputPlanChargesInner) HasBillableMetricId() bool`

HasBillableMetricId returns a boolean if a field has been set.

### GetChargeModel

`func (o *PlanCreateInputPlanChargesInner) GetChargeModel() string`

GetChargeModel returns the ChargeModel field if non-nil, zero value otherwise.

### GetChargeModelOk

`func (o *PlanCreateInputPlanChargesInner) GetChargeModelOk() (*string, bool)`

GetChargeModelOk returns a tuple with the ChargeModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeModel

`func (o *PlanCreateInputPlanChargesInner) SetChargeModel(v string)`

SetChargeModel sets ChargeModel field to given value.

### HasChargeModel

`func (o *PlanCreateInputPlanChargesInner) HasChargeModel() bool`

HasChargeModel returns a boolean if a field has been set.

### GetPayInAdvance

`func (o *PlanCreateInputPlanChargesInner) GetPayInAdvance() bool`

GetPayInAdvance returns the PayInAdvance field if non-nil, zero value otherwise.

### GetPayInAdvanceOk

`func (o *PlanCreateInputPlanChargesInner) GetPayInAdvanceOk() (*bool, bool)`

GetPayInAdvanceOk returns a tuple with the PayInAdvance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayInAdvance

`func (o *PlanCreateInputPlanChargesInner) SetPayInAdvance(v bool)`

SetPayInAdvance sets PayInAdvance field to given value.

### HasPayInAdvance

`func (o *PlanCreateInputPlanChargesInner) HasPayInAdvance() bool`

HasPayInAdvance returns a boolean if a field has been set.

### GetInvoiceable

`func (o *PlanCreateInputPlanChargesInner) GetInvoiceable() bool`

GetInvoiceable returns the Invoiceable field if non-nil, zero value otherwise.

### GetInvoiceableOk

`func (o *PlanCreateInputPlanChargesInner) GetInvoiceableOk() (*bool, bool)`

GetInvoiceableOk returns a tuple with the Invoiceable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceable

`func (o *PlanCreateInputPlanChargesInner) SetInvoiceable(v bool)`

SetInvoiceable sets Invoiceable field to given value.

### HasInvoiceable

`func (o *PlanCreateInputPlanChargesInner) HasInvoiceable() bool`

HasInvoiceable returns a boolean if a field has been set.

### GetInvoiceDisplayName

`func (o *PlanCreateInputPlanChargesInner) GetInvoiceDisplayName() string`

GetInvoiceDisplayName returns the InvoiceDisplayName field if non-nil, zero value otherwise.

### GetInvoiceDisplayNameOk

`func (o *PlanCreateInputPlanChargesInner) GetInvoiceDisplayNameOk() (*string, bool)`

GetInvoiceDisplayNameOk returns a tuple with the InvoiceDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceDisplayName

`func (o *PlanCreateInputPlanChargesInner) SetInvoiceDisplayName(v string)`

SetInvoiceDisplayName sets InvoiceDisplayName field to given value.

### HasInvoiceDisplayName

`func (o *PlanCreateInputPlanChargesInner) HasInvoiceDisplayName() bool`

HasInvoiceDisplayName returns a boolean if a field has been set.

### GetProrated

`func (o *PlanCreateInputPlanChargesInner) GetProrated() bool`

GetProrated returns the Prorated field if non-nil, zero value otherwise.

### GetProratedOk

`func (o *PlanCreateInputPlanChargesInner) GetProratedOk() (*bool, bool)`

GetProratedOk returns a tuple with the Prorated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProrated

`func (o *PlanCreateInputPlanChargesInner) SetProrated(v bool)`

SetProrated sets Prorated field to given value.

### HasProrated

`func (o *PlanCreateInputPlanChargesInner) HasProrated() bool`

HasProrated returns a boolean if a field has been set.

### GetMinAmountCents

`func (o *PlanCreateInputPlanChargesInner) GetMinAmountCents() int32`

GetMinAmountCents returns the MinAmountCents field if non-nil, zero value otherwise.

### GetMinAmountCentsOk

`func (o *PlanCreateInputPlanChargesInner) GetMinAmountCentsOk() (*int32, bool)`

GetMinAmountCentsOk returns a tuple with the MinAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAmountCents

`func (o *PlanCreateInputPlanChargesInner) SetMinAmountCents(v int32)`

SetMinAmountCents sets MinAmountCents field to given value.

### HasMinAmountCents

`func (o *PlanCreateInputPlanChargesInner) HasMinAmountCents() bool`

HasMinAmountCents returns a boolean if a field has been set.

### GetProperties

`func (o *PlanCreateInputPlanChargesInner) GetProperties() ChargeObjectProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *PlanCreateInputPlanChargesInner) GetPropertiesOk() (*ChargeObjectProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *PlanCreateInputPlanChargesInner) SetProperties(v ChargeObjectProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *PlanCreateInputPlanChargesInner) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetGroupProperties

`func (o *PlanCreateInputPlanChargesInner) GetGroupProperties() []PlanCreateInputPlanChargesInnerGroupPropertiesInner`

GetGroupProperties returns the GroupProperties field if non-nil, zero value otherwise.

### GetGroupPropertiesOk

`func (o *PlanCreateInputPlanChargesInner) GetGroupPropertiesOk() (*[]PlanCreateInputPlanChargesInnerGroupPropertiesInner, bool)`

GetGroupPropertiesOk returns a tuple with the GroupProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupProperties

`func (o *PlanCreateInputPlanChargesInner) SetGroupProperties(v []PlanCreateInputPlanChargesInnerGroupPropertiesInner)`

SetGroupProperties sets GroupProperties field to given value.

### HasGroupProperties

`func (o *PlanCreateInputPlanChargesInner) HasGroupProperties() bool`

HasGroupProperties returns a boolean if a field has been set.

### GetTaxCodes

`func (o *PlanCreateInputPlanChargesInner) GetTaxCodes() []string`

GetTaxCodes returns the TaxCodes field if non-nil, zero value otherwise.

### GetTaxCodesOk

`func (o *PlanCreateInputPlanChargesInner) GetTaxCodesOk() (*[]string, bool)`

GetTaxCodesOk returns a tuple with the TaxCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCodes

`func (o *PlanCreateInputPlanChargesInner) SetTaxCodes(v []string)`

SetTaxCodes sets TaxCodes field to given value.

### HasTaxCodes

`func (o *PlanCreateInputPlanChargesInner) HasTaxCodes() bool`

HasTaxCodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


