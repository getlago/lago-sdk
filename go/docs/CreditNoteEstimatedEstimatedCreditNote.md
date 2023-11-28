# CreditNoteEstimatedEstimatedCreditNote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LagoInvoiceId** | **string** | Unique identifier assigned to the invoice that the credit note belongs to | 
**InvoiceNumber** | **string** | The invoice unique number, related to the credit note. | 
**Currency** | [**Currency**](Currency.md) |  | 
**TaxesAmountCents** | **int32** | The tax amount of the credit note, expressed in cents. | 
**TaxesRate** | **float32** | The tax rate associated with this specific credit note. | 
**SubTotalExcludingTaxesAmountCents** | **int32** | The subtotal of the credit note excluding any applicable taxes, expressed in cents. | 
**MaxCreditableAmountCents** | **int32** | The credited amount of the credit note, expressed in cents. | 
**MaxRefundableAmountCents** | **int32** | The refunded amount of the credit note, expressed in cents. | 
**CouponsAdjustmentAmountCents** | **int32** | The pro-rated amount of the coupons applied to the source invoice. | 
**Items** | [**[]CreditNoteEstimatedEstimatedCreditNoteItemsInner**](CreditNoteEstimatedEstimatedCreditNoteItemsInner.md) | Array of credit note’s items. | 
**AppliedTaxes** | Pointer to [**[]CreditNoteEstimatedEstimatedCreditNoteAppliedTaxesInner**](CreditNoteEstimatedEstimatedCreditNoteAppliedTaxesInner.md) |  | [optional] 

## Methods

### NewCreditNoteEstimatedEstimatedCreditNote

`func NewCreditNoteEstimatedEstimatedCreditNote(lagoInvoiceId string, invoiceNumber string, currency Currency, taxesAmountCents int32, taxesRate float32, subTotalExcludingTaxesAmountCents int32, maxCreditableAmountCents int32, maxRefundableAmountCents int32, couponsAdjustmentAmountCents int32, items []CreditNoteEstimatedEstimatedCreditNoteItemsInner, ) *CreditNoteEstimatedEstimatedCreditNote`

NewCreditNoteEstimatedEstimatedCreditNote instantiates a new CreditNoteEstimatedEstimatedCreditNote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditNoteEstimatedEstimatedCreditNoteWithDefaults

`func NewCreditNoteEstimatedEstimatedCreditNoteWithDefaults() *CreditNoteEstimatedEstimatedCreditNote`

NewCreditNoteEstimatedEstimatedCreditNoteWithDefaults instantiates a new CreditNoteEstimatedEstimatedCreditNote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLagoInvoiceId

`func (o *CreditNoteEstimatedEstimatedCreditNote) GetLagoInvoiceId() string`

GetLagoInvoiceId returns the LagoInvoiceId field if non-nil, zero value otherwise.

### GetLagoInvoiceIdOk

`func (o *CreditNoteEstimatedEstimatedCreditNote) GetLagoInvoiceIdOk() (*string, bool)`

GetLagoInvoiceIdOk returns a tuple with the LagoInvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagoInvoiceId

`func (o *CreditNoteEstimatedEstimatedCreditNote) SetLagoInvoiceId(v string)`

SetLagoInvoiceId sets LagoInvoiceId field to given value.


### GetInvoiceNumber

`func (o *CreditNoteEstimatedEstimatedCreditNote) GetInvoiceNumber() string`

GetInvoiceNumber returns the InvoiceNumber field if non-nil, zero value otherwise.

### GetInvoiceNumberOk

`func (o *CreditNoteEstimatedEstimatedCreditNote) GetInvoiceNumberOk() (*string, bool)`

GetInvoiceNumberOk returns a tuple with the InvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNumber

`func (o *CreditNoteEstimatedEstimatedCreditNote) SetInvoiceNumber(v string)`

SetInvoiceNumber sets InvoiceNumber field to given value.


### GetCurrency

`func (o *CreditNoteEstimatedEstimatedCreditNote) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreditNoteEstimatedEstimatedCreditNote) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreditNoteEstimatedEstimatedCreditNote) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.


### GetTaxesAmountCents

`func (o *CreditNoteEstimatedEstimatedCreditNote) GetTaxesAmountCents() int32`

GetTaxesAmountCents returns the TaxesAmountCents field if non-nil, zero value otherwise.

### GetTaxesAmountCentsOk

`func (o *CreditNoteEstimatedEstimatedCreditNote) GetTaxesAmountCentsOk() (*int32, bool)`

GetTaxesAmountCentsOk returns a tuple with the TaxesAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxesAmountCents

`func (o *CreditNoteEstimatedEstimatedCreditNote) SetTaxesAmountCents(v int32)`

SetTaxesAmountCents sets TaxesAmountCents field to given value.


### GetTaxesRate

`func (o *CreditNoteEstimatedEstimatedCreditNote) GetTaxesRate() float32`

GetTaxesRate returns the TaxesRate field if non-nil, zero value otherwise.

### GetTaxesRateOk

`func (o *CreditNoteEstimatedEstimatedCreditNote) GetTaxesRateOk() (*float32, bool)`

GetTaxesRateOk returns a tuple with the TaxesRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxesRate

`func (o *CreditNoteEstimatedEstimatedCreditNote) SetTaxesRate(v float32)`

SetTaxesRate sets TaxesRate field to given value.


### GetSubTotalExcludingTaxesAmountCents

`func (o *CreditNoteEstimatedEstimatedCreditNote) GetSubTotalExcludingTaxesAmountCents() int32`

GetSubTotalExcludingTaxesAmountCents returns the SubTotalExcludingTaxesAmountCents field if non-nil, zero value otherwise.

### GetSubTotalExcludingTaxesAmountCentsOk

`func (o *CreditNoteEstimatedEstimatedCreditNote) GetSubTotalExcludingTaxesAmountCentsOk() (*int32, bool)`

GetSubTotalExcludingTaxesAmountCentsOk returns a tuple with the SubTotalExcludingTaxesAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTotalExcludingTaxesAmountCents

`func (o *CreditNoteEstimatedEstimatedCreditNote) SetSubTotalExcludingTaxesAmountCents(v int32)`

SetSubTotalExcludingTaxesAmountCents sets SubTotalExcludingTaxesAmountCents field to given value.


### GetMaxCreditableAmountCents

`func (o *CreditNoteEstimatedEstimatedCreditNote) GetMaxCreditableAmountCents() int32`

GetMaxCreditableAmountCents returns the MaxCreditableAmountCents field if non-nil, zero value otherwise.

### GetMaxCreditableAmountCentsOk

`func (o *CreditNoteEstimatedEstimatedCreditNote) GetMaxCreditableAmountCentsOk() (*int32, bool)`

GetMaxCreditableAmountCentsOk returns a tuple with the MaxCreditableAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCreditableAmountCents

`func (o *CreditNoteEstimatedEstimatedCreditNote) SetMaxCreditableAmountCents(v int32)`

SetMaxCreditableAmountCents sets MaxCreditableAmountCents field to given value.


### GetMaxRefundableAmountCents

`func (o *CreditNoteEstimatedEstimatedCreditNote) GetMaxRefundableAmountCents() int32`

GetMaxRefundableAmountCents returns the MaxRefundableAmountCents field if non-nil, zero value otherwise.

### GetMaxRefundableAmountCentsOk

`func (o *CreditNoteEstimatedEstimatedCreditNote) GetMaxRefundableAmountCentsOk() (*int32, bool)`

GetMaxRefundableAmountCentsOk returns a tuple with the MaxRefundableAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRefundableAmountCents

`func (o *CreditNoteEstimatedEstimatedCreditNote) SetMaxRefundableAmountCents(v int32)`

SetMaxRefundableAmountCents sets MaxRefundableAmountCents field to given value.


### GetCouponsAdjustmentAmountCents

`func (o *CreditNoteEstimatedEstimatedCreditNote) GetCouponsAdjustmentAmountCents() int32`

GetCouponsAdjustmentAmountCents returns the CouponsAdjustmentAmountCents field if non-nil, zero value otherwise.

### GetCouponsAdjustmentAmountCentsOk

`func (o *CreditNoteEstimatedEstimatedCreditNote) GetCouponsAdjustmentAmountCentsOk() (*int32, bool)`

GetCouponsAdjustmentAmountCentsOk returns a tuple with the CouponsAdjustmentAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponsAdjustmentAmountCents

`func (o *CreditNoteEstimatedEstimatedCreditNote) SetCouponsAdjustmentAmountCents(v int32)`

SetCouponsAdjustmentAmountCents sets CouponsAdjustmentAmountCents field to given value.


### GetItems

`func (o *CreditNoteEstimatedEstimatedCreditNote) GetItems() []CreditNoteEstimatedEstimatedCreditNoteItemsInner`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CreditNoteEstimatedEstimatedCreditNote) GetItemsOk() (*[]CreditNoteEstimatedEstimatedCreditNoteItemsInner, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CreditNoteEstimatedEstimatedCreditNote) SetItems(v []CreditNoteEstimatedEstimatedCreditNoteItemsInner)`

SetItems sets Items field to given value.


### GetAppliedTaxes

`func (o *CreditNoteEstimatedEstimatedCreditNote) GetAppliedTaxes() []CreditNoteEstimatedEstimatedCreditNoteAppliedTaxesInner`

GetAppliedTaxes returns the AppliedTaxes field if non-nil, zero value otherwise.

### GetAppliedTaxesOk

`func (o *CreditNoteEstimatedEstimatedCreditNote) GetAppliedTaxesOk() (*[]CreditNoteEstimatedEstimatedCreditNoteAppliedTaxesInner, bool)`

GetAppliedTaxesOk returns a tuple with the AppliedTaxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedTaxes

`func (o *CreditNoteEstimatedEstimatedCreditNote) SetAppliedTaxes(v []CreditNoteEstimatedEstimatedCreditNoteAppliedTaxesInner)`

SetAppliedTaxes sets AppliedTaxes field to given value.

### HasAppliedTaxes

`func (o *CreditNoteEstimatedEstimatedCreditNote) HasAppliedTaxes() bool`

HasAppliedTaxes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


