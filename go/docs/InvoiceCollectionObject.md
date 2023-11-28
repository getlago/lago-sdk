# InvoiceCollectionObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Month** | **string** | Identifies the month to analyze revenue. | 
**PaymentStatus** | Pointer to **string** | The payment status of the invoices. | [optional] 
**InvoicesCount** | **int32** | Contains invoices count. | 
**AmountCents** | **int32** | The total amount of revenue for a period, expressed in cents. | 
**Currency** | Pointer to [**Currency**](Currency.md) |  | [optional] 

## Methods

### NewInvoiceCollectionObject

`func NewInvoiceCollectionObject(month string, invoicesCount int32, amountCents int32, ) *InvoiceCollectionObject`

NewInvoiceCollectionObject instantiates a new InvoiceCollectionObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceCollectionObjectWithDefaults

`func NewInvoiceCollectionObjectWithDefaults() *InvoiceCollectionObject`

NewInvoiceCollectionObjectWithDefaults instantiates a new InvoiceCollectionObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonth

`func (o *InvoiceCollectionObject) GetMonth() string`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *InvoiceCollectionObject) GetMonthOk() (*string, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *InvoiceCollectionObject) SetMonth(v string)`

SetMonth sets Month field to given value.


### GetPaymentStatus

`func (o *InvoiceCollectionObject) GetPaymentStatus() string`

GetPaymentStatus returns the PaymentStatus field if non-nil, zero value otherwise.

### GetPaymentStatusOk

`func (o *InvoiceCollectionObject) GetPaymentStatusOk() (*string, bool)`

GetPaymentStatusOk returns a tuple with the PaymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentStatus

`func (o *InvoiceCollectionObject) SetPaymentStatus(v string)`

SetPaymentStatus sets PaymentStatus field to given value.

### HasPaymentStatus

`func (o *InvoiceCollectionObject) HasPaymentStatus() bool`

HasPaymentStatus returns a boolean if a field has been set.

### GetInvoicesCount

`func (o *InvoiceCollectionObject) GetInvoicesCount() int32`

GetInvoicesCount returns the InvoicesCount field if non-nil, zero value otherwise.

### GetInvoicesCountOk

`func (o *InvoiceCollectionObject) GetInvoicesCountOk() (*int32, bool)`

GetInvoicesCountOk returns a tuple with the InvoicesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicesCount

`func (o *InvoiceCollectionObject) SetInvoicesCount(v int32)`

SetInvoicesCount sets InvoicesCount field to given value.


### GetAmountCents

`func (o *InvoiceCollectionObject) GetAmountCents() int32`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *InvoiceCollectionObject) GetAmountCentsOk() (*int32, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *InvoiceCollectionObject) SetAmountCents(v int32)`

SetAmountCents sets AmountCents field to given value.


### GetCurrency

`func (o *InvoiceCollectionObject) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *InvoiceCollectionObject) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *InvoiceCollectionObject) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *InvoiceCollectionObject) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


