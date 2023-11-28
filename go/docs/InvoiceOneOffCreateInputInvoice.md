# InvoiceOneOffCreateInputInvoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalCustomerId** | **string** | Unique identifier assigned to the customer in your application. | 
**Currency** | Pointer to [**Currency**](Currency.md) |  | [optional] 
**Fees** | [**[]InvoiceOneOffCreateInputInvoiceFeesInner**](InvoiceOneOffCreateInputInvoiceFeesInner.md) |  | 

## Methods

### NewInvoiceOneOffCreateInputInvoice

`func NewInvoiceOneOffCreateInputInvoice(externalCustomerId string, fees []InvoiceOneOffCreateInputInvoiceFeesInner, ) *InvoiceOneOffCreateInputInvoice`

NewInvoiceOneOffCreateInputInvoice instantiates a new InvoiceOneOffCreateInputInvoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceOneOffCreateInputInvoiceWithDefaults

`func NewInvoiceOneOffCreateInputInvoiceWithDefaults() *InvoiceOneOffCreateInputInvoice`

NewInvoiceOneOffCreateInputInvoiceWithDefaults instantiates a new InvoiceOneOffCreateInputInvoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalCustomerId

`func (o *InvoiceOneOffCreateInputInvoice) GetExternalCustomerId() string`

GetExternalCustomerId returns the ExternalCustomerId field if non-nil, zero value otherwise.

### GetExternalCustomerIdOk

`func (o *InvoiceOneOffCreateInputInvoice) GetExternalCustomerIdOk() (*string, bool)`

GetExternalCustomerIdOk returns a tuple with the ExternalCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalCustomerId

`func (o *InvoiceOneOffCreateInputInvoice) SetExternalCustomerId(v string)`

SetExternalCustomerId sets ExternalCustomerId field to given value.


### GetCurrency

`func (o *InvoiceOneOffCreateInputInvoice) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *InvoiceOneOffCreateInputInvoice) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *InvoiceOneOffCreateInputInvoice) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *InvoiceOneOffCreateInputInvoice) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetFees

`func (o *InvoiceOneOffCreateInputInvoice) GetFees() []InvoiceOneOffCreateInputInvoiceFeesInner`

GetFees returns the Fees field if non-nil, zero value otherwise.

### GetFeesOk

`func (o *InvoiceOneOffCreateInputInvoice) GetFeesOk() (*[]InvoiceOneOffCreateInputInvoiceFeesInner, bool)`

GetFeesOk returns a tuple with the Fees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFees

`func (o *InvoiceOneOffCreateInputInvoice) SetFees(v []InvoiceOneOffCreateInputInvoiceFeesInner)`

SetFees sets Fees field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


