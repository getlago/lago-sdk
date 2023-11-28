# CreditNoteEstimateInputCreditNote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvoiceId** | **string** | The invoice unique identifier, created by Lago. | 
**Items** | [**[]CreditNoteEstimateInputCreditNoteItemsInner**](CreditNoteEstimateInputCreditNoteItemsInner.md) | The list of credit note’s items. | 

## Methods

### NewCreditNoteEstimateInputCreditNote

`func NewCreditNoteEstimateInputCreditNote(invoiceId string, items []CreditNoteEstimateInputCreditNoteItemsInner, ) *CreditNoteEstimateInputCreditNote`

NewCreditNoteEstimateInputCreditNote instantiates a new CreditNoteEstimateInputCreditNote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditNoteEstimateInputCreditNoteWithDefaults

`func NewCreditNoteEstimateInputCreditNoteWithDefaults() *CreditNoteEstimateInputCreditNote`

NewCreditNoteEstimateInputCreditNoteWithDefaults instantiates a new CreditNoteEstimateInputCreditNote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoiceId

`func (o *CreditNoteEstimateInputCreditNote) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *CreditNoteEstimateInputCreditNote) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *CreditNoteEstimateInputCreditNote) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.


### GetItems

`func (o *CreditNoteEstimateInputCreditNote) GetItems() []CreditNoteEstimateInputCreditNoteItemsInner`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CreditNoteEstimateInputCreditNote) GetItemsOk() (*[]CreditNoteEstimateInputCreditNoteItemsInner, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CreditNoteEstimateInputCreditNote) SetItems(v []CreditNoteEstimateInputCreditNoteItemsInner)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


