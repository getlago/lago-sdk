# CreditNoteCreateInputCreditNote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvoiceId** | **string** | The invoice unique identifier, created by Lago. | 
**Reason** | Pointer to **NullableString** | The reason of the credit note creation. Possible values are &#x60;duplicated_charge&#x60;, &#x60;product_unsatisfactory&#x60;, &#x60;order_change&#x60;, &#x60;order_cancellation&#x60;, &#x60;fraudulent_charge&#x60; or &#x60;other&#x60;. | [optional] 
**Description** | Pointer to **string** | The description of the credit note. | [optional] 
**CreditAmountCents** | Pointer to **NullableInt32** | The total amount to be credited on the customer balance. | [optional] 
**RefundAmountCents** | Pointer to **NullableInt32** | The total amount to be refunded to the customer. | [optional] 
**Items** | [**[]CreditNoteEstimateInputCreditNoteItemsInner**](CreditNoteEstimateInputCreditNoteItemsInner.md) | The list of credit note’s items. | 

## Methods

### NewCreditNoteCreateInputCreditNote

`func NewCreditNoteCreateInputCreditNote(invoiceId string, items []CreditNoteEstimateInputCreditNoteItemsInner, ) *CreditNoteCreateInputCreditNote`

NewCreditNoteCreateInputCreditNote instantiates a new CreditNoteCreateInputCreditNote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditNoteCreateInputCreditNoteWithDefaults

`func NewCreditNoteCreateInputCreditNoteWithDefaults() *CreditNoteCreateInputCreditNote`

NewCreditNoteCreateInputCreditNoteWithDefaults instantiates a new CreditNoteCreateInputCreditNote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoiceId

`func (o *CreditNoteCreateInputCreditNote) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *CreditNoteCreateInputCreditNote) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *CreditNoteCreateInputCreditNote) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.


### GetReason

`func (o *CreditNoteCreateInputCreditNote) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CreditNoteCreateInputCreditNote) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CreditNoteCreateInputCreditNote) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CreditNoteCreateInputCreditNote) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *CreditNoteCreateInputCreditNote) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *CreditNoteCreateInputCreditNote) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetDescription

`func (o *CreditNoteCreateInputCreditNote) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreditNoteCreateInputCreditNote) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreditNoteCreateInputCreditNote) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreditNoteCreateInputCreditNote) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreditAmountCents

`func (o *CreditNoteCreateInputCreditNote) GetCreditAmountCents() int32`

GetCreditAmountCents returns the CreditAmountCents field if non-nil, zero value otherwise.

### GetCreditAmountCentsOk

`func (o *CreditNoteCreateInputCreditNote) GetCreditAmountCentsOk() (*int32, bool)`

GetCreditAmountCentsOk returns a tuple with the CreditAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditAmountCents

`func (o *CreditNoteCreateInputCreditNote) SetCreditAmountCents(v int32)`

SetCreditAmountCents sets CreditAmountCents field to given value.

### HasCreditAmountCents

`func (o *CreditNoteCreateInputCreditNote) HasCreditAmountCents() bool`

HasCreditAmountCents returns a boolean if a field has been set.

### SetCreditAmountCentsNil

`func (o *CreditNoteCreateInputCreditNote) SetCreditAmountCentsNil(b bool)`

 SetCreditAmountCentsNil sets the value for CreditAmountCents to be an explicit nil

### UnsetCreditAmountCents
`func (o *CreditNoteCreateInputCreditNote) UnsetCreditAmountCents()`

UnsetCreditAmountCents ensures that no value is present for CreditAmountCents, not even an explicit nil
### GetRefundAmountCents

`func (o *CreditNoteCreateInputCreditNote) GetRefundAmountCents() int32`

GetRefundAmountCents returns the RefundAmountCents field if non-nil, zero value otherwise.

### GetRefundAmountCentsOk

`func (o *CreditNoteCreateInputCreditNote) GetRefundAmountCentsOk() (*int32, bool)`

GetRefundAmountCentsOk returns a tuple with the RefundAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundAmountCents

`func (o *CreditNoteCreateInputCreditNote) SetRefundAmountCents(v int32)`

SetRefundAmountCents sets RefundAmountCents field to given value.

### HasRefundAmountCents

`func (o *CreditNoteCreateInputCreditNote) HasRefundAmountCents() bool`

HasRefundAmountCents returns a boolean if a field has been set.

### SetRefundAmountCentsNil

`func (o *CreditNoteCreateInputCreditNote) SetRefundAmountCentsNil(b bool)`

 SetRefundAmountCentsNil sets the value for RefundAmountCents to be an explicit nil

### UnsetRefundAmountCents
`func (o *CreditNoteCreateInputCreditNote) UnsetRefundAmountCents()`

UnsetRefundAmountCents ensures that no value is present for RefundAmountCents, not even an explicit nil
### GetItems

`func (o *CreditNoteCreateInputCreditNote) GetItems() []CreditNoteEstimateInputCreditNoteItemsInner`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CreditNoteCreateInputCreditNote) GetItemsOk() (*[]CreditNoteEstimateInputCreditNoteItemsInner, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CreditNoteCreateInputCreditNote) SetItems(v []CreditNoteEstimateInputCreditNoteItemsInner)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


