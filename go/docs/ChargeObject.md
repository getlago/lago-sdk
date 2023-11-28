# ChargeObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LagoId** | **string** | Unique identifier of charge, created by Lago. | 
**LagoBillableMetricId** | **string** | Unique identifier of the billable metric created by Lago. | 
**BillableMetricCode** | **string** | Unique code identifying a billable metric. | 
**InvoiceDisplayName** | Pointer to **string** | Specifies the name that will be displayed on an invoice. If no value is set for this field, the name of the actual charge will be used as the default display name. | [optional] 
**CreatedAt** | **time.Time** | The date and time when the charge was created. It is expressed in UTC format according to the ISO 8601 datetime standard. | 
**ChargeModel** | **string** | Specifies the pricing model used for the calculation of the final fee. It can be &#x60;standard&#x60;, &#x60;graduated&#x60;, &#x60;graduated_percentage&#x60;, &#x60;package&#x60;, &#x60;percentage&#x60; or &#x60;volume&#x60;. | 
**PayInAdvance** | Pointer to **bool** | This field determines the billing timing for this specific usage-based charge. When set to &#x60;true&#x60;, the charge is due and invoiced immediately. Conversely, when set to &#x60;false&#x60;, the charge is due and invoiced at the end of each billing period. | [optional] 
**Invoiceable** | Pointer to **bool** | This field specifies whether the charge should be included in a proper invoice. If set to &#x60;false&#x60;, no invoice will be issued for this charge. You can only set it to &#x60;false&#x60; when &#x60;pay_in_advance&#x60; is &#x60;true&#x60;. | [optional] 
**Prorated** | Pointer to **bool** | Specifies whether a charge is prorated based on the remaining number of days in the billing period or billed fully.  - If set to &#x60;true&#x60;, the charge is prorated based on the remaining days in the current billing period. - If set to &#x60;false&#x60;, the charge is billed in full. - If not defined in the request, default value is &#x60;false&#x60;. | [optional] 
**MinAmountCents** | Pointer to **int32** | The minimum spending amount required for the charge, measured in cents and excluding any applicable taxes. It indicates the minimum amount that needs to be charged for each billing period. | [optional] 
**Properties** | Pointer to [**ChargeObjectProperties**](ChargeObjectProperties.md) |  | [optional] 
**GroupProperties** | Pointer to [**[]GroupPropertiesObject**](GroupPropertiesObject.md) | All charge information, sorted by groups. | [optional] 
**Taxes** | Pointer to [**[]TaxObject**](TaxObject.md) | All taxes applied to the charge. | [optional] 

## Methods

### NewChargeObject

`func NewChargeObject(lagoId string, lagoBillableMetricId string, billableMetricCode string, createdAt time.Time, chargeModel string, ) *ChargeObject`

NewChargeObject instantiates a new ChargeObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChargeObjectWithDefaults

`func NewChargeObjectWithDefaults() *ChargeObject`

NewChargeObjectWithDefaults instantiates a new ChargeObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLagoId

`func (o *ChargeObject) GetLagoId() string`

GetLagoId returns the LagoId field if non-nil, zero value otherwise.

### GetLagoIdOk

`func (o *ChargeObject) GetLagoIdOk() (*string, bool)`

GetLagoIdOk returns a tuple with the LagoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagoId

`func (o *ChargeObject) SetLagoId(v string)`

SetLagoId sets LagoId field to given value.


### GetLagoBillableMetricId

`func (o *ChargeObject) GetLagoBillableMetricId() string`

GetLagoBillableMetricId returns the LagoBillableMetricId field if non-nil, zero value otherwise.

### GetLagoBillableMetricIdOk

`func (o *ChargeObject) GetLagoBillableMetricIdOk() (*string, bool)`

GetLagoBillableMetricIdOk returns a tuple with the LagoBillableMetricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagoBillableMetricId

`func (o *ChargeObject) SetLagoBillableMetricId(v string)`

SetLagoBillableMetricId sets LagoBillableMetricId field to given value.


### GetBillableMetricCode

`func (o *ChargeObject) GetBillableMetricCode() string`

GetBillableMetricCode returns the BillableMetricCode field if non-nil, zero value otherwise.

### GetBillableMetricCodeOk

`func (o *ChargeObject) GetBillableMetricCodeOk() (*string, bool)`

GetBillableMetricCodeOk returns a tuple with the BillableMetricCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillableMetricCode

`func (o *ChargeObject) SetBillableMetricCode(v string)`

SetBillableMetricCode sets BillableMetricCode field to given value.


### GetInvoiceDisplayName

`func (o *ChargeObject) GetInvoiceDisplayName() string`

GetInvoiceDisplayName returns the InvoiceDisplayName field if non-nil, zero value otherwise.

### GetInvoiceDisplayNameOk

`func (o *ChargeObject) GetInvoiceDisplayNameOk() (*string, bool)`

GetInvoiceDisplayNameOk returns a tuple with the InvoiceDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceDisplayName

`func (o *ChargeObject) SetInvoiceDisplayName(v string)`

SetInvoiceDisplayName sets InvoiceDisplayName field to given value.

### HasInvoiceDisplayName

`func (o *ChargeObject) HasInvoiceDisplayName() bool`

HasInvoiceDisplayName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ChargeObject) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ChargeObject) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ChargeObject) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetChargeModel

`func (o *ChargeObject) GetChargeModel() string`

GetChargeModel returns the ChargeModel field if non-nil, zero value otherwise.

### GetChargeModelOk

`func (o *ChargeObject) GetChargeModelOk() (*string, bool)`

GetChargeModelOk returns a tuple with the ChargeModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeModel

`func (o *ChargeObject) SetChargeModel(v string)`

SetChargeModel sets ChargeModel field to given value.


### GetPayInAdvance

`func (o *ChargeObject) GetPayInAdvance() bool`

GetPayInAdvance returns the PayInAdvance field if non-nil, zero value otherwise.

### GetPayInAdvanceOk

`func (o *ChargeObject) GetPayInAdvanceOk() (*bool, bool)`

GetPayInAdvanceOk returns a tuple with the PayInAdvance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayInAdvance

`func (o *ChargeObject) SetPayInAdvance(v bool)`

SetPayInAdvance sets PayInAdvance field to given value.

### HasPayInAdvance

`func (o *ChargeObject) HasPayInAdvance() bool`

HasPayInAdvance returns a boolean if a field has been set.

### GetInvoiceable

`func (o *ChargeObject) GetInvoiceable() bool`

GetInvoiceable returns the Invoiceable field if non-nil, zero value otherwise.

### GetInvoiceableOk

`func (o *ChargeObject) GetInvoiceableOk() (*bool, bool)`

GetInvoiceableOk returns a tuple with the Invoiceable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceable

`func (o *ChargeObject) SetInvoiceable(v bool)`

SetInvoiceable sets Invoiceable field to given value.

### HasInvoiceable

`func (o *ChargeObject) HasInvoiceable() bool`

HasInvoiceable returns a boolean if a field has been set.

### GetProrated

`func (o *ChargeObject) GetProrated() bool`

GetProrated returns the Prorated field if non-nil, zero value otherwise.

### GetProratedOk

`func (o *ChargeObject) GetProratedOk() (*bool, bool)`

GetProratedOk returns a tuple with the Prorated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProrated

`func (o *ChargeObject) SetProrated(v bool)`

SetProrated sets Prorated field to given value.

### HasProrated

`func (o *ChargeObject) HasProrated() bool`

HasProrated returns a boolean if a field has been set.

### GetMinAmountCents

`func (o *ChargeObject) GetMinAmountCents() int32`

GetMinAmountCents returns the MinAmountCents field if non-nil, zero value otherwise.

### GetMinAmountCentsOk

`func (o *ChargeObject) GetMinAmountCentsOk() (*int32, bool)`

GetMinAmountCentsOk returns a tuple with the MinAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAmountCents

`func (o *ChargeObject) SetMinAmountCents(v int32)`

SetMinAmountCents sets MinAmountCents field to given value.

### HasMinAmountCents

`func (o *ChargeObject) HasMinAmountCents() bool`

HasMinAmountCents returns a boolean if a field has been set.

### GetProperties

`func (o *ChargeObject) GetProperties() ChargeObjectProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ChargeObject) GetPropertiesOk() (*ChargeObjectProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ChargeObject) SetProperties(v ChargeObjectProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *ChargeObject) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetGroupProperties

`func (o *ChargeObject) GetGroupProperties() []GroupPropertiesObject`

GetGroupProperties returns the GroupProperties field if non-nil, zero value otherwise.

### GetGroupPropertiesOk

`func (o *ChargeObject) GetGroupPropertiesOk() (*[]GroupPropertiesObject, bool)`

GetGroupPropertiesOk returns a tuple with the GroupProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupProperties

`func (o *ChargeObject) SetGroupProperties(v []GroupPropertiesObject)`

SetGroupProperties sets GroupProperties field to given value.

### HasGroupProperties

`func (o *ChargeObject) HasGroupProperties() bool`

HasGroupProperties returns a boolean if a field has been set.

### GetTaxes

`func (o *ChargeObject) GetTaxes() []TaxObject`

GetTaxes returns the Taxes field if non-nil, zero value otherwise.

### GetTaxesOk

`func (o *ChargeObject) GetTaxesOk() (*[]TaxObject, bool)`

GetTaxesOk returns a tuple with the Taxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxes

`func (o *ChargeObject) SetTaxes(v []TaxObject)`

SetTaxes sets Taxes field to given value.

### HasTaxes

`func (o *ChargeObject) HasTaxes() bool`

HasTaxes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


