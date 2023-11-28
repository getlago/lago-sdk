# InvoicedUsageObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Month** | **string** | Identifies the month to analyze revenue. | 
**Code** | Pointer to **string** | The code of the usage-based billable metrics. | [optional] 
**AmountCents** | **int32** | The total amount of revenue for a period, expressed in cents. | 
**Currency** | [**Currency**](Currency.md) |  | 

## Methods

### NewInvoicedUsageObject

`func NewInvoicedUsageObject(month string, amountCents int32, currency Currency, ) *InvoicedUsageObject`

NewInvoicedUsageObject instantiates a new InvoicedUsageObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoicedUsageObjectWithDefaults

`func NewInvoicedUsageObjectWithDefaults() *InvoicedUsageObject`

NewInvoicedUsageObjectWithDefaults instantiates a new InvoicedUsageObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonth

`func (o *InvoicedUsageObject) GetMonth() string`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *InvoicedUsageObject) GetMonthOk() (*string, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *InvoicedUsageObject) SetMonth(v string)`

SetMonth sets Month field to given value.


### GetCode

`func (o *InvoicedUsageObject) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *InvoicedUsageObject) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *InvoicedUsageObject) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *InvoicedUsageObject) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetAmountCents

`func (o *InvoicedUsageObject) GetAmountCents() int32`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *InvoicedUsageObject) GetAmountCentsOk() (*int32, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *InvoicedUsageObject) SetAmountCents(v int32)`

SetAmountCents sets AmountCents field to given value.


### GetCurrency

`func (o *InvoicedUsageObject) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *InvoicedUsageObject) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *InvoicedUsageObject) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


