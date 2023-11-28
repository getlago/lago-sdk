# CreditObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LagoId** | **string** | Unique identifier assigned to the credit within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the credit’s item record within the Lago system. | 
**AmountCents** | **int32** | The amount of credit associated with the invoice, expressed in cents. | 
**AmountCurrency** | [**Currency**](Currency.md) |  | 
**BeforeTaxes** | **bool** | Indicates whether the credit is applied on the amount before taxes (coupons) or after taxes (credit notes). This flag helps determine the order in which credits are applied to the invoice calculation | 
**Item** | [**CreditObjectItem**](CreditObjectItem.md) |  | 
**Invoice** | [**CreditObjectInvoice**](CreditObjectInvoice.md) |  | 

## Methods

### NewCreditObject

`func NewCreditObject(lagoId string, amountCents int32, amountCurrency Currency, beforeTaxes bool, item CreditObjectItem, invoice CreditObjectInvoice, ) *CreditObject`

NewCreditObject instantiates a new CreditObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditObjectWithDefaults

`func NewCreditObjectWithDefaults() *CreditObject`

NewCreditObjectWithDefaults instantiates a new CreditObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLagoId

`func (o *CreditObject) GetLagoId() string`

GetLagoId returns the LagoId field if non-nil, zero value otherwise.

### GetLagoIdOk

`func (o *CreditObject) GetLagoIdOk() (*string, bool)`

GetLagoIdOk returns a tuple with the LagoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagoId

`func (o *CreditObject) SetLagoId(v string)`

SetLagoId sets LagoId field to given value.


### GetAmountCents

`func (o *CreditObject) GetAmountCents() int32`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *CreditObject) GetAmountCentsOk() (*int32, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *CreditObject) SetAmountCents(v int32)`

SetAmountCents sets AmountCents field to given value.


### GetAmountCurrency

`func (o *CreditObject) GetAmountCurrency() Currency`

GetAmountCurrency returns the AmountCurrency field if non-nil, zero value otherwise.

### GetAmountCurrencyOk

`func (o *CreditObject) GetAmountCurrencyOk() (*Currency, bool)`

GetAmountCurrencyOk returns a tuple with the AmountCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCurrency

`func (o *CreditObject) SetAmountCurrency(v Currency)`

SetAmountCurrency sets AmountCurrency field to given value.


### GetBeforeTaxes

`func (o *CreditObject) GetBeforeTaxes() bool`

GetBeforeTaxes returns the BeforeTaxes field if non-nil, zero value otherwise.

### GetBeforeTaxesOk

`func (o *CreditObject) GetBeforeTaxesOk() (*bool, bool)`

GetBeforeTaxesOk returns a tuple with the BeforeTaxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeforeTaxes

`func (o *CreditObject) SetBeforeTaxes(v bool)`

SetBeforeTaxes sets BeforeTaxes field to given value.


### GetItem

`func (o *CreditObject) GetItem() CreditObjectItem`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *CreditObject) GetItemOk() (*CreditObjectItem, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *CreditObject) SetItem(v CreditObjectItem)`

SetItem sets Item field to given value.


### GetInvoice

`func (o *CreditObject) GetInvoice() CreditObjectInvoice`

GetInvoice returns the Invoice field if non-nil, zero value otherwise.

### GetInvoiceOk

`func (o *CreditObject) GetInvoiceOk() (*CreditObjectInvoice, bool)`

GetInvoiceOk returns a tuple with the Invoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoice

`func (o *CreditObject) SetInvoice(v CreditObjectInvoice)`

SetInvoice sets Invoice field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


