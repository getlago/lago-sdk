# CreditNoteItemObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LagoId** | **string** | The credit note’s item unique identifier, created by Lago. | 
**AmountCents** | **int32** | The credit note’s item amount, expressed in cents. | 
**AmountCurrency** | [**Currency**](Currency.md) |  | 
**Fee** | [**CreditNoteItemObjectFee**](CreditNoteItemObjectFee.md) |  | 

## Methods

### NewCreditNoteItemObject

`func NewCreditNoteItemObject(lagoId string, amountCents int32, amountCurrency Currency, fee CreditNoteItemObjectFee, ) *CreditNoteItemObject`

NewCreditNoteItemObject instantiates a new CreditNoteItemObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditNoteItemObjectWithDefaults

`func NewCreditNoteItemObjectWithDefaults() *CreditNoteItemObject`

NewCreditNoteItemObjectWithDefaults instantiates a new CreditNoteItemObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLagoId

`func (o *CreditNoteItemObject) GetLagoId() string`

GetLagoId returns the LagoId field if non-nil, zero value otherwise.

### GetLagoIdOk

`func (o *CreditNoteItemObject) GetLagoIdOk() (*string, bool)`

GetLagoIdOk returns a tuple with the LagoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagoId

`func (o *CreditNoteItemObject) SetLagoId(v string)`

SetLagoId sets LagoId field to given value.


### GetAmountCents

`func (o *CreditNoteItemObject) GetAmountCents() int32`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *CreditNoteItemObject) GetAmountCentsOk() (*int32, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *CreditNoteItemObject) SetAmountCents(v int32)`

SetAmountCents sets AmountCents field to given value.


### GetAmountCurrency

`func (o *CreditNoteItemObject) GetAmountCurrency() Currency`

GetAmountCurrency returns the AmountCurrency field if non-nil, zero value otherwise.

### GetAmountCurrencyOk

`func (o *CreditNoteItemObject) GetAmountCurrencyOk() (*Currency, bool)`

GetAmountCurrencyOk returns a tuple with the AmountCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCurrency

`func (o *CreditNoteItemObject) SetAmountCurrency(v Currency)`

SetAmountCurrency sets AmountCurrency field to given value.


### GetFee

`func (o *CreditNoteItemObject) GetFee() CreditNoteItemObjectFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *CreditNoteItemObject) GetFeeOk() (*CreditNoteItemObjectFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *CreditNoteItemObject) SetFee(v CreditNoteItemObjectFee)`

SetFee sets Fee field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


