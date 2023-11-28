# CreditNoteObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LagoId** | **string** | The credit note unique identifier, created by Lago. | 
**SequentialId** | **int32** | The sequential identifier of the credit note, specifically scoped on the associated invoice. It provides a unique numerical identifier for the credit note within the context of the invoice. | 
**Number** | **string** | The credit note unique number. | 
**LagoInvoiceId** | **string** | Unique identifier assigned to the invoice that the credit note belongs to | 
**InvoiceNumber** | **string** | The invoice unique number, related to the credit note. | 
**IssuingDate** | **string** | The date of creation of the credit note. It follows the ISO 8601 date format and provides the specific date when the credit note was created. | 
**CreditStatus** | Pointer to **NullableString** | The status of the credit portion of the credit note. It indicates the current state or condition of the credit amount associated with the credit note. The possible values for this field are:  - &#x60;available&#x60;: this status indicates that an amount remains available for future usage. The credit can be applied towards future transactions or invoices. - &#x60;consumed&#x60;: this status indicates that the credit amount has been fully consumed. The remaining amount is 0, indicating that the credit has been utilized in its entirety. - &#x60;voided&#x60;: this status indicates that the remaining amount of the credit cannot be used any further. The credit has been voided and is no longer available for application or redemption. | [optional] 
**RefundStatus** | Pointer to **NullableString** | The status of the refund portion of the credit note. It indicates the current state or condition of the refund associated with the credit note. The possible values for this field are:  - &#x60;pending&#x60;: this status indicates that the refund is pending execution. The refund request has been initiated but has not been processed or completed yet. - &#x60;succeeded&#x60;: this status indicates that the refund has been successfully executed. The refund amount has been processed and returned to the customer or the designated recipient. - &#x60;failed&#x60;: this status indicates that the refund failed to execute. The refund request encountered an error or unsuccessful processing, and the refund amount could not be returned. | [optional] 
**Reason** | **string** | The reason of the credit note creation. Possible values are &#x60;duplicated_charge&#x60;, &#x60;product_unsatisfactory&#x60;, &#x60;order_change&#x60;, &#x60;order_cancellation&#x60;, &#x60;fraudulent_charge&#x60; or &#x60;other&#x60;. | 
**Description** | Pointer to **NullableString** | The description of the credit note. | [optional] 
**Currency** | [**Currency**](Currency.md) |  | 
**TotalAmountCents** | **int32** | The total amount of the credit note, expressed in cents. | 
**TaxesAmountCents** | **int32** | The tax amount of the credit note, expressed in cents. | 
**TaxesRate** | **float32** | The tax rate associated with this specific credit note. | 
**SubTotalExcludingTaxesAmountCents** | **int32** | The subtotal of the credit note excluding any applicable taxes, expressed in cents. | 
**BalanceAmountCents** | **int32** | The remaining credit note amount, expressed in cents. | 
**CreditAmountCents** | **int32** | The credited amount of the credit note, expressed in cents. | 
**RefundAmountCents** | **int32** | The refunded amount of the credit note, expressed in cents. | 
**CouponsAdjustmentAmountCents** | **int32** | The pro-rated amount of the coupons applied to the source invoice. | 
**CreatedAt** | **time.Time** | The date when the credit note was created. It is expressed in Coordinated Universal Time (UTC). | 
**UpdatedAt** | **time.Time** | The date when the credit note was last updated. It is expressed in Coordinated Universal Time (UTC). | 
**FileUrl** | Pointer to **NullableString** | The PDF file of the credit note. | [optional] 
**Items** | Pointer to [**[]CreditNoteItemObject**](CreditNoteItemObject.md) | Array of credit note’s items. | [optional] 
**AppliedTaxes** | Pointer to [**[]CreditNoteAppliedTaxObject**](CreditNoteAppliedTaxObject.md) |  | [optional] 

## Methods

### NewCreditNoteObject

`func NewCreditNoteObject(lagoId string, sequentialId int32, number string, lagoInvoiceId string, invoiceNumber string, issuingDate string, reason string, currency Currency, totalAmountCents int32, taxesAmountCents int32, taxesRate float32, subTotalExcludingTaxesAmountCents int32, balanceAmountCents int32, creditAmountCents int32, refundAmountCents int32, couponsAdjustmentAmountCents int32, createdAt time.Time, updatedAt time.Time, ) *CreditNoteObject`

NewCreditNoteObject instantiates a new CreditNoteObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditNoteObjectWithDefaults

`func NewCreditNoteObjectWithDefaults() *CreditNoteObject`

NewCreditNoteObjectWithDefaults instantiates a new CreditNoteObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLagoId

`func (o *CreditNoteObject) GetLagoId() string`

GetLagoId returns the LagoId field if non-nil, zero value otherwise.

### GetLagoIdOk

`func (o *CreditNoteObject) GetLagoIdOk() (*string, bool)`

GetLagoIdOk returns a tuple with the LagoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagoId

`func (o *CreditNoteObject) SetLagoId(v string)`

SetLagoId sets LagoId field to given value.


### GetSequentialId

`func (o *CreditNoteObject) GetSequentialId() int32`

GetSequentialId returns the SequentialId field if non-nil, zero value otherwise.

### GetSequentialIdOk

`func (o *CreditNoteObject) GetSequentialIdOk() (*int32, bool)`

GetSequentialIdOk returns a tuple with the SequentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequentialId

`func (o *CreditNoteObject) SetSequentialId(v int32)`

SetSequentialId sets SequentialId field to given value.


### GetNumber

`func (o *CreditNoteObject) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *CreditNoteObject) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *CreditNoteObject) SetNumber(v string)`

SetNumber sets Number field to given value.


### GetLagoInvoiceId

`func (o *CreditNoteObject) GetLagoInvoiceId() string`

GetLagoInvoiceId returns the LagoInvoiceId field if non-nil, zero value otherwise.

### GetLagoInvoiceIdOk

`func (o *CreditNoteObject) GetLagoInvoiceIdOk() (*string, bool)`

GetLagoInvoiceIdOk returns a tuple with the LagoInvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagoInvoiceId

`func (o *CreditNoteObject) SetLagoInvoiceId(v string)`

SetLagoInvoiceId sets LagoInvoiceId field to given value.


### GetInvoiceNumber

`func (o *CreditNoteObject) GetInvoiceNumber() string`

GetInvoiceNumber returns the InvoiceNumber field if non-nil, zero value otherwise.

### GetInvoiceNumberOk

`func (o *CreditNoteObject) GetInvoiceNumberOk() (*string, bool)`

GetInvoiceNumberOk returns a tuple with the InvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNumber

`func (o *CreditNoteObject) SetInvoiceNumber(v string)`

SetInvoiceNumber sets InvoiceNumber field to given value.


### GetIssuingDate

`func (o *CreditNoteObject) GetIssuingDate() string`

GetIssuingDate returns the IssuingDate field if non-nil, zero value otherwise.

### GetIssuingDateOk

`func (o *CreditNoteObject) GetIssuingDateOk() (*string, bool)`

GetIssuingDateOk returns a tuple with the IssuingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingDate

`func (o *CreditNoteObject) SetIssuingDate(v string)`

SetIssuingDate sets IssuingDate field to given value.


### GetCreditStatus

`func (o *CreditNoteObject) GetCreditStatus() string`

GetCreditStatus returns the CreditStatus field if non-nil, zero value otherwise.

### GetCreditStatusOk

`func (o *CreditNoteObject) GetCreditStatusOk() (*string, bool)`

GetCreditStatusOk returns a tuple with the CreditStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditStatus

`func (o *CreditNoteObject) SetCreditStatus(v string)`

SetCreditStatus sets CreditStatus field to given value.

### HasCreditStatus

`func (o *CreditNoteObject) HasCreditStatus() bool`

HasCreditStatus returns a boolean if a field has been set.

### SetCreditStatusNil

`func (o *CreditNoteObject) SetCreditStatusNil(b bool)`

 SetCreditStatusNil sets the value for CreditStatus to be an explicit nil

### UnsetCreditStatus
`func (o *CreditNoteObject) UnsetCreditStatus()`

UnsetCreditStatus ensures that no value is present for CreditStatus, not even an explicit nil
### GetRefundStatus

`func (o *CreditNoteObject) GetRefundStatus() string`

GetRefundStatus returns the RefundStatus field if non-nil, zero value otherwise.

### GetRefundStatusOk

`func (o *CreditNoteObject) GetRefundStatusOk() (*string, bool)`

GetRefundStatusOk returns a tuple with the RefundStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundStatus

`func (o *CreditNoteObject) SetRefundStatus(v string)`

SetRefundStatus sets RefundStatus field to given value.

### HasRefundStatus

`func (o *CreditNoteObject) HasRefundStatus() bool`

HasRefundStatus returns a boolean if a field has been set.

### SetRefundStatusNil

`func (o *CreditNoteObject) SetRefundStatusNil(b bool)`

 SetRefundStatusNil sets the value for RefundStatus to be an explicit nil

### UnsetRefundStatus
`func (o *CreditNoteObject) UnsetRefundStatus()`

UnsetRefundStatus ensures that no value is present for RefundStatus, not even an explicit nil
### GetReason

`func (o *CreditNoteObject) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CreditNoteObject) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CreditNoteObject) SetReason(v string)`

SetReason sets Reason field to given value.


### GetDescription

`func (o *CreditNoteObject) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreditNoteObject) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreditNoteObject) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreditNoteObject) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreditNoteObject) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreditNoteObject) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCurrency

`func (o *CreditNoteObject) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreditNoteObject) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreditNoteObject) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.


### GetTotalAmountCents

`func (o *CreditNoteObject) GetTotalAmountCents() int32`

GetTotalAmountCents returns the TotalAmountCents field if non-nil, zero value otherwise.

### GetTotalAmountCentsOk

`func (o *CreditNoteObject) GetTotalAmountCentsOk() (*int32, bool)`

GetTotalAmountCentsOk returns a tuple with the TotalAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountCents

`func (o *CreditNoteObject) SetTotalAmountCents(v int32)`

SetTotalAmountCents sets TotalAmountCents field to given value.


### GetTaxesAmountCents

`func (o *CreditNoteObject) GetTaxesAmountCents() int32`

GetTaxesAmountCents returns the TaxesAmountCents field if non-nil, zero value otherwise.

### GetTaxesAmountCentsOk

`func (o *CreditNoteObject) GetTaxesAmountCentsOk() (*int32, bool)`

GetTaxesAmountCentsOk returns a tuple with the TaxesAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxesAmountCents

`func (o *CreditNoteObject) SetTaxesAmountCents(v int32)`

SetTaxesAmountCents sets TaxesAmountCents field to given value.


### GetTaxesRate

`func (o *CreditNoteObject) GetTaxesRate() float32`

GetTaxesRate returns the TaxesRate field if non-nil, zero value otherwise.

### GetTaxesRateOk

`func (o *CreditNoteObject) GetTaxesRateOk() (*float32, bool)`

GetTaxesRateOk returns a tuple with the TaxesRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxesRate

`func (o *CreditNoteObject) SetTaxesRate(v float32)`

SetTaxesRate sets TaxesRate field to given value.


### GetSubTotalExcludingTaxesAmountCents

`func (o *CreditNoteObject) GetSubTotalExcludingTaxesAmountCents() int32`

GetSubTotalExcludingTaxesAmountCents returns the SubTotalExcludingTaxesAmountCents field if non-nil, zero value otherwise.

### GetSubTotalExcludingTaxesAmountCentsOk

`func (o *CreditNoteObject) GetSubTotalExcludingTaxesAmountCentsOk() (*int32, bool)`

GetSubTotalExcludingTaxesAmountCentsOk returns a tuple with the SubTotalExcludingTaxesAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTotalExcludingTaxesAmountCents

`func (o *CreditNoteObject) SetSubTotalExcludingTaxesAmountCents(v int32)`

SetSubTotalExcludingTaxesAmountCents sets SubTotalExcludingTaxesAmountCents field to given value.


### GetBalanceAmountCents

`func (o *CreditNoteObject) GetBalanceAmountCents() int32`

GetBalanceAmountCents returns the BalanceAmountCents field if non-nil, zero value otherwise.

### GetBalanceAmountCentsOk

`func (o *CreditNoteObject) GetBalanceAmountCentsOk() (*int32, bool)`

GetBalanceAmountCentsOk returns a tuple with the BalanceAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceAmountCents

`func (o *CreditNoteObject) SetBalanceAmountCents(v int32)`

SetBalanceAmountCents sets BalanceAmountCents field to given value.


### GetCreditAmountCents

`func (o *CreditNoteObject) GetCreditAmountCents() int32`

GetCreditAmountCents returns the CreditAmountCents field if non-nil, zero value otherwise.

### GetCreditAmountCentsOk

`func (o *CreditNoteObject) GetCreditAmountCentsOk() (*int32, bool)`

GetCreditAmountCentsOk returns a tuple with the CreditAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditAmountCents

`func (o *CreditNoteObject) SetCreditAmountCents(v int32)`

SetCreditAmountCents sets CreditAmountCents field to given value.


### GetRefundAmountCents

`func (o *CreditNoteObject) GetRefundAmountCents() int32`

GetRefundAmountCents returns the RefundAmountCents field if non-nil, zero value otherwise.

### GetRefundAmountCentsOk

`func (o *CreditNoteObject) GetRefundAmountCentsOk() (*int32, bool)`

GetRefundAmountCentsOk returns a tuple with the RefundAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundAmountCents

`func (o *CreditNoteObject) SetRefundAmountCents(v int32)`

SetRefundAmountCents sets RefundAmountCents field to given value.


### GetCouponsAdjustmentAmountCents

`func (o *CreditNoteObject) GetCouponsAdjustmentAmountCents() int32`

GetCouponsAdjustmentAmountCents returns the CouponsAdjustmentAmountCents field if non-nil, zero value otherwise.

### GetCouponsAdjustmentAmountCentsOk

`func (o *CreditNoteObject) GetCouponsAdjustmentAmountCentsOk() (*int32, bool)`

GetCouponsAdjustmentAmountCentsOk returns a tuple with the CouponsAdjustmentAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponsAdjustmentAmountCents

`func (o *CreditNoteObject) SetCouponsAdjustmentAmountCents(v int32)`

SetCouponsAdjustmentAmountCents sets CouponsAdjustmentAmountCents field to given value.


### GetCreatedAt

`func (o *CreditNoteObject) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreditNoteObject) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreditNoteObject) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *CreditNoteObject) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CreditNoteObject) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CreditNoteObject) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetFileUrl

`func (o *CreditNoteObject) GetFileUrl() string`

GetFileUrl returns the FileUrl field if non-nil, zero value otherwise.

### GetFileUrlOk

`func (o *CreditNoteObject) GetFileUrlOk() (*string, bool)`

GetFileUrlOk returns a tuple with the FileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUrl

`func (o *CreditNoteObject) SetFileUrl(v string)`

SetFileUrl sets FileUrl field to given value.

### HasFileUrl

`func (o *CreditNoteObject) HasFileUrl() bool`

HasFileUrl returns a boolean if a field has been set.

### SetFileUrlNil

`func (o *CreditNoteObject) SetFileUrlNil(b bool)`

 SetFileUrlNil sets the value for FileUrl to be an explicit nil

### UnsetFileUrl
`func (o *CreditNoteObject) UnsetFileUrl()`

UnsetFileUrl ensures that no value is present for FileUrl, not even an explicit nil
### GetItems

`func (o *CreditNoteObject) GetItems() []CreditNoteItemObject`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CreditNoteObject) GetItemsOk() (*[]CreditNoteItemObject, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CreditNoteObject) SetItems(v []CreditNoteItemObject)`

SetItems sets Items field to given value.

### HasItems

`func (o *CreditNoteObject) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetAppliedTaxes

`func (o *CreditNoteObject) GetAppliedTaxes() []CreditNoteAppliedTaxObject`

GetAppliedTaxes returns the AppliedTaxes field if non-nil, zero value otherwise.

### GetAppliedTaxesOk

`func (o *CreditNoteObject) GetAppliedTaxesOk() (*[]CreditNoteAppliedTaxObject, bool)`

GetAppliedTaxesOk returns a tuple with the AppliedTaxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedTaxes

`func (o *CreditNoteObject) SetAppliedTaxes(v []CreditNoteAppliedTaxObject)`

SetAppliedTaxes sets AppliedTaxes field to given value.

### HasAppliedTaxes

`func (o *CreditNoteObject) HasAppliedTaxes() bool`

HasAppliedTaxes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


