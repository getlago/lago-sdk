# PlanOverridesObjectChargesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier of the charge created by Lago. | [optional] 
**BillableMetricId** | Pointer to **string** | Unique identifier of the billable metric created by Lago. | [optional] 
**InvoiceDisplayName** | Pointer to **string** | Specifies the name that will be displayed on an invoice. If no value is set for this field, the name of the actual charge will be used as the default display name. | [optional] 
**MinAmountCents** | Pointer to **int32** | The minimum spending amount required for the charge, measured in cents and excluding any applicable taxes. It indicates the minimum amount that needs to be charged for each billing period. | [optional] 
**Properties** | Pointer to [**ChargeObjectProperties**](ChargeObjectProperties.md) |  | [optional] 
**GroupProperties** | Pointer to [**[]PlanCreateInputPlanChargesInnerGroupPropertiesInner**](PlanCreateInputPlanChargesInnerGroupPropertiesInner.md) | All charge information, sorted by groups. | [optional] 
**TaxCodes** | Pointer to **[]string** | List of unique code used to identify the taxes. | [optional] 

## Methods

### NewPlanOverridesObjectChargesInner

`func NewPlanOverridesObjectChargesInner() *PlanOverridesObjectChargesInner`

NewPlanOverridesObjectChargesInner instantiates a new PlanOverridesObjectChargesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanOverridesObjectChargesInnerWithDefaults

`func NewPlanOverridesObjectChargesInnerWithDefaults() *PlanOverridesObjectChargesInner`

NewPlanOverridesObjectChargesInnerWithDefaults instantiates a new PlanOverridesObjectChargesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PlanOverridesObjectChargesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlanOverridesObjectChargesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlanOverridesObjectChargesInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PlanOverridesObjectChargesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBillableMetricId

`func (o *PlanOverridesObjectChargesInner) GetBillableMetricId() string`

GetBillableMetricId returns the BillableMetricId field if non-nil, zero value otherwise.

### GetBillableMetricIdOk

`func (o *PlanOverridesObjectChargesInner) GetBillableMetricIdOk() (*string, bool)`

GetBillableMetricIdOk returns a tuple with the BillableMetricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillableMetricId

`func (o *PlanOverridesObjectChargesInner) SetBillableMetricId(v string)`

SetBillableMetricId sets BillableMetricId field to given value.

### HasBillableMetricId

`func (o *PlanOverridesObjectChargesInner) HasBillableMetricId() bool`

HasBillableMetricId returns a boolean if a field has been set.

### GetInvoiceDisplayName

`func (o *PlanOverridesObjectChargesInner) GetInvoiceDisplayName() string`

GetInvoiceDisplayName returns the InvoiceDisplayName field if non-nil, zero value otherwise.

### GetInvoiceDisplayNameOk

`func (o *PlanOverridesObjectChargesInner) GetInvoiceDisplayNameOk() (*string, bool)`

GetInvoiceDisplayNameOk returns a tuple with the InvoiceDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceDisplayName

`func (o *PlanOverridesObjectChargesInner) SetInvoiceDisplayName(v string)`

SetInvoiceDisplayName sets InvoiceDisplayName field to given value.

### HasInvoiceDisplayName

`func (o *PlanOverridesObjectChargesInner) HasInvoiceDisplayName() bool`

HasInvoiceDisplayName returns a boolean if a field has been set.

### GetMinAmountCents

`func (o *PlanOverridesObjectChargesInner) GetMinAmountCents() int32`

GetMinAmountCents returns the MinAmountCents field if non-nil, zero value otherwise.

### GetMinAmountCentsOk

`func (o *PlanOverridesObjectChargesInner) GetMinAmountCentsOk() (*int32, bool)`

GetMinAmountCentsOk returns a tuple with the MinAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAmountCents

`func (o *PlanOverridesObjectChargesInner) SetMinAmountCents(v int32)`

SetMinAmountCents sets MinAmountCents field to given value.

### HasMinAmountCents

`func (o *PlanOverridesObjectChargesInner) HasMinAmountCents() bool`

HasMinAmountCents returns a boolean if a field has been set.

### GetProperties

`func (o *PlanOverridesObjectChargesInner) GetProperties() ChargeObjectProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *PlanOverridesObjectChargesInner) GetPropertiesOk() (*ChargeObjectProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *PlanOverridesObjectChargesInner) SetProperties(v ChargeObjectProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *PlanOverridesObjectChargesInner) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetGroupProperties

`func (o *PlanOverridesObjectChargesInner) GetGroupProperties() []PlanCreateInputPlanChargesInnerGroupPropertiesInner`

GetGroupProperties returns the GroupProperties field if non-nil, zero value otherwise.

### GetGroupPropertiesOk

`func (o *PlanOverridesObjectChargesInner) GetGroupPropertiesOk() (*[]PlanCreateInputPlanChargesInnerGroupPropertiesInner, bool)`

GetGroupPropertiesOk returns a tuple with the GroupProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupProperties

`func (o *PlanOverridesObjectChargesInner) SetGroupProperties(v []PlanCreateInputPlanChargesInnerGroupPropertiesInner)`

SetGroupProperties sets GroupProperties field to given value.

### HasGroupProperties

`func (o *PlanOverridesObjectChargesInner) HasGroupProperties() bool`

HasGroupProperties returns a boolean if a field has been set.

### GetTaxCodes

`func (o *PlanOverridesObjectChargesInner) GetTaxCodes() []string`

GetTaxCodes returns the TaxCodes field if non-nil, zero value otherwise.

### GetTaxCodesOk

`func (o *PlanOverridesObjectChargesInner) GetTaxCodesOk() (*[]string, bool)`

GetTaxCodesOk returns a tuple with the TaxCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCodes

`func (o *PlanOverridesObjectChargesInner) SetTaxCodes(v []string)`

SetTaxCodes sets TaxCodes field to given value.

### HasTaxCodes

`func (o *PlanOverridesObjectChargesInner) HasTaxCodes() bool`

HasTaxCodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


