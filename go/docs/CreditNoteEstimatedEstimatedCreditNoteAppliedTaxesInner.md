# CreditNoteEstimatedEstimatedCreditNoteAppliedTaxesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LagoTaxId** | Pointer to **string** | Unique identifier of the tax, created by Lago. | [optional] 
**TaxName** | Pointer to **string** | Name of the tax. | [optional] 
**TaxCode** | Pointer to **string** | Unique code used to identify the tax associated with the API request. | [optional] 
**TaxRate** | Pointer to **float32** | The percentage rate of the tax | [optional] 
**TaxDescription** | Pointer to **string** | Internal description of the taxe | [optional] 
**BaseAmountCents** | Pointer to **int32** |  | [optional] 
**AmountCents** | Pointer to **int32** | Amount of the tax | [optional] 
**AmountCurrency** | Pointer to [**Currency**](Currency.md) |  | [optional] 

## Methods

### NewCreditNoteEstimatedEstimatedCreditNoteAppliedTaxesInner

`func NewCreditNoteEstimatedEstimatedCreditNoteAppliedTaxesInner() *CreditNoteEstimatedEstimatedCreditNoteAppliedTaxesInner`

NewCreditNoteEstimatedEstimatedCreditNoteAppliedTaxesInner instantiates a new CreditNoteEstimatedEstimatedCreditNoteAppliedTaxesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditNoteEstimatedEstimatedCreditNoteAppliedTaxesInnerWithDefaults

`func NewCreditNoteEstimatedEstimatedCreditNoteAppliedTaxesInnerWithDefaults() *CreditNoteEstimatedEstimatedCreditNoteAppliedTaxesInner`

NewCreditNoteEstimatedEstimatedCreditNoteAppliedTaxesInnerWithDefaults instantiates a new CreditNoteEstimatedEstimatedCreditNoteAppliedTaxesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLagoTaxId

`func (o *CreditNoteEstimatedEstimatedCreditNoteAppliedTaxesInner) GetLagoTaxId() string`

GetLagoTaxId returns the LagoTaxId field if non-nil, zero value otherwise.

### GetLagoTaxIdOk

`func (o *CreditNoteEstimatedEstimatedCreditNoteAppliedTaxesInner) GetLagoTaxIdOk() (*string, bool)`

GetLagoTaxIdOk returns a tuple with the LagoTaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagoTaxId

`func (o *CreditNoteEstimatedEstimatedCreditNoteAppliedTaxesInner) SetLagoTaxId(v string)`

SetLagoTaxId sets LagoTaxId field to given value.

### HasLagoTaxId

`func (o *CreditNoteEstimatedEstimatedCreditNoteAppliedTaxesInner) HasLagoTaxId() bool`

HasLagoTaxId returns a boolean if a field has been set.

### GetTaxName

`func (o *CreditNoteEstimatedEstimatedCreditNoteAppliedTaxesInner) GetTaxName() string`

GetTaxName returns the TaxName field if non-nil, zero value otherwise.

### GetTaxNameOk

`func (o *CreditNoteEstimatedEstimatedCreditNoteAppliedTaxesInner) GetTaxNameOk() (*string, bool)`

GetTaxNameOk returns a tuple with the TaxName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxName

`func (o *CreditNoteEstimatedEstimatedCreditNoteAppliedTaxesInner) SetTaxName(v string)`

SetTaxName sets TaxName field to given value.

### HasTaxName

`func (o *CreditNoteEstimatedEstimatedCreditNoteAppliedTaxesInner) HasTaxName() bool`

HasTaxName returns a boolean if a field has been set.

### GetTaxCode

`func (o *CreditNoteEstimatedEstimatedCreditNoteAppliedTaxesInner) GetTaxCode() string`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *CreditNoteEstimatedEstimatedCreditNoteAppliedTaxesInner) GetTaxCodeOk() (*string, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *CreditNoteEstimatedEstimatedCreditNoteAppliedTaxesInner) SetTaxCode(v string)`

SetTaxCode sets TaxCode field to given value.

### HasTaxCode

`func (o *CreditNoteEstimatedEstimatedCreditNoteAppliedTaxesInner) HasTaxCode() bool`

HasTaxCode returns a boolean if a field has been set.

### GetTaxRate

`func (o *CreditNoteEstimatedEstimatedCreditNoteAppliedTaxesInner) GetTaxRate() float32`

GetTaxRate returns the TaxRate field if non-nil, zero value otherwise.

### GetTaxRateOk

`func (o *CreditNoteEstimatedEstimatedCreditNoteAppliedTaxesInner) GetTaxRateOk() (*float32, bool)`

GetTaxRateOk returns a tuple with the TaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate

`func (o *CreditNoteEstimatedEstimatedCreditNoteAppliedTaxesInner) SetTaxRate(v float32)`

SetTaxRate sets TaxRate field to given value.

### HasTaxRate

`func (o *CreditNoteEstimatedEstimatedCreditNoteAppliedTaxesInner) HasTaxRate() bool`

HasTaxRate returns a boolean if a field has been set.

### GetTaxDescription

`func (o *CreditNoteEstimatedEstimatedCreditNoteAppliedTaxesInner) GetTaxDescription() string`

GetTaxDescription returns the TaxDescription field if non-nil, zero value otherwise.

### GetTaxDescriptionOk

`func (o *CreditNoteEstimatedEstimatedCreditNoteAppliedTaxesInner) GetTaxDescriptionOk() (*string, bool)`

GetTaxDescriptionOk returns a tuple with the TaxDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxDescription

`func (o *CreditNoteEstimatedEstimatedCreditNoteAppliedTaxesInner) SetTaxDescription(v string)`

SetTaxDescription sets TaxDescription field to given value.

### HasTaxDescription

`func (o *CreditNoteEstimatedEstimatedCreditNoteAppliedTaxesInner) HasTaxDescription() bool`

HasTaxDescription returns a boolean if a field has been set.

### GetBaseAmountCents

`func (o *CreditNoteEstimatedEstimatedCreditNoteAppliedTaxesInner) GetBaseAmountCents() int32`

GetBaseAmountCents returns the BaseAmountCents field if non-nil, zero value otherwise.

### GetBaseAmountCentsOk

`func (o *CreditNoteEstimatedEstimatedCreditNoteAppliedTaxesInner) GetBaseAmountCentsOk() (*int32, bool)`

GetBaseAmountCentsOk returns a tuple with the BaseAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAmountCents

`func (o *CreditNoteEstimatedEstimatedCreditNoteAppliedTaxesInner) SetBaseAmountCents(v int32)`

SetBaseAmountCents sets BaseAmountCents field to given value.

### HasBaseAmountCents

`func (o *CreditNoteEstimatedEstimatedCreditNoteAppliedTaxesInner) HasBaseAmountCents() bool`

HasBaseAmountCents returns a boolean if a field has been set.

### GetAmountCents

`func (o *CreditNoteEstimatedEstimatedCreditNoteAppliedTaxesInner) GetAmountCents() int32`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *CreditNoteEstimatedEstimatedCreditNoteAppliedTaxesInner) GetAmountCentsOk() (*int32, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *CreditNoteEstimatedEstimatedCreditNoteAppliedTaxesInner) SetAmountCents(v int32)`

SetAmountCents sets AmountCents field to given value.

### HasAmountCents

`func (o *CreditNoteEstimatedEstimatedCreditNoteAppliedTaxesInner) HasAmountCents() bool`

HasAmountCents returns a boolean if a field has been set.

### GetAmountCurrency

`func (o *CreditNoteEstimatedEstimatedCreditNoteAppliedTaxesInner) GetAmountCurrency() Currency`

GetAmountCurrency returns the AmountCurrency field if non-nil, zero value otherwise.

### GetAmountCurrencyOk

`func (o *CreditNoteEstimatedEstimatedCreditNoteAppliedTaxesInner) GetAmountCurrencyOk() (*Currency, bool)`

GetAmountCurrencyOk returns a tuple with the AmountCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCurrency

`func (o *CreditNoteEstimatedEstimatedCreditNoteAppliedTaxesInner) SetAmountCurrency(v Currency)`

SetAmountCurrency sets AmountCurrency field to given value.

### HasAmountCurrency

`func (o *CreditNoteEstimatedEstimatedCreditNoteAppliedTaxesInner) HasAmountCurrency() bool`

HasAmountCurrency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


