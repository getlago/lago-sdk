# CreditNoteAppliedTaxObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LagoCreditNoteId** | Pointer to **string** | Unique identifier of the credit note, created by Lago. | [optional] 
**BaseAmountCents** | Pointer to **int32** |  | [optional] 
**LagoId** | Pointer to **string** | Unique identifier of the applied tax, created by Lago. | [optional] 
**LagoTaxId** | Pointer to **string** | Unique identifier of the tax, created by Lago. | [optional] 
**TaxName** | Pointer to **string** | Name of the tax. | [optional] 
**TaxCode** | Pointer to **string** | Unique code used to identify the tax associated with the API request. | [optional] 
**TaxRate** | Pointer to **float32** | The percentage rate of the tax | [optional] 
**TaxDescription** | Pointer to **string** | Internal description of the taxe | [optional] 
**AmountCents** | Pointer to **int32** | Amount of the tax | [optional] 
**AmountCurrency** | Pointer to [**Currency**](Currency.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date and time when the applied tax was created. It is expressed in UTC format according to the ISO 8601 datetime standard. This field provides the timestamp for the exact moment when the applied tax was initially created. | [optional] 

## Methods

### NewCreditNoteAppliedTaxObject

`func NewCreditNoteAppliedTaxObject() *CreditNoteAppliedTaxObject`

NewCreditNoteAppliedTaxObject instantiates a new CreditNoteAppliedTaxObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditNoteAppliedTaxObjectWithDefaults

`func NewCreditNoteAppliedTaxObjectWithDefaults() *CreditNoteAppliedTaxObject`

NewCreditNoteAppliedTaxObjectWithDefaults instantiates a new CreditNoteAppliedTaxObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLagoCreditNoteId

`func (o *CreditNoteAppliedTaxObject) GetLagoCreditNoteId() string`

GetLagoCreditNoteId returns the LagoCreditNoteId field if non-nil, zero value otherwise.

### GetLagoCreditNoteIdOk

`func (o *CreditNoteAppliedTaxObject) GetLagoCreditNoteIdOk() (*string, bool)`

GetLagoCreditNoteIdOk returns a tuple with the LagoCreditNoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagoCreditNoteId

`func (o *CreditNoteAppliedTaxObject) SetLagoCreditNoteId(v string)`

SetLagoCreditNoteId sets LagoCreditNoteId field to given value.

### HasLagoCreditNoteId

`func (o *CreditNoteAppliedTaxObject) HasLagoCreditNoteId() bool`

HasLagoCreditNoteId returns a boolean if a field has been set.

### GetBaseAmountCents

`func (o *CreditNoteAppliedTaxObject) GetBaseAmountCents() int32`

GetBaseAmountCents returns the BaseAmountCents field if non-nil, zero value otherwise.

### GetBaseAmountCentsOk

`func (o *CreditNoteAppliedTaxObject) GetBaseAmountCentsOk() (*int32, bool)`

GetBaseAmountCentsOk returns a tuple with the BaseAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAmountCents

`func (o *CreditNoteAppliedTaxObject) SetBaseAmountCents(v int32)`

SetBaseAmountCents sets BaseAmountCents field to given value.

### HasBaseAmountCents

`func (o *CreditNoteAppliedTaxObject) HasBaseAmountCents() bool`

HasBaseAmountCents returns a boolean if a field has been set.

### GetLagoId

`func (o *CreditNoteAppliedTaxObject) GetLagoId() string`

GetLagoId returns the LagoId field if non-nil, zero value otherwise.

### GetLagoIdOk

`func (o *CreditNoteAppliedTaxObject) GetLagoIdOk() (*string, bool)`

GetLagoIdOk returns a tuple with the LagoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagoId

`func (o *CreditNoteAppliedTaxObject) SetLagoId(v string)`

SetLagoId sets LagoId field to given value.

### HasLagoId

`func (o *CreditNoteAppliedTaxObject) HasLagoId() bool`

HasLagoId returns a boolean if a field has been set.

### GetLagoTaxId

`func (o *CreditNoteAppliedTaxObject) GetLagoTaxId() string`

GetLagoTaxId returns the LagoTaxId field if non-nil, zero value otherwise.

### GetLagoTaxIdOk

`func (o *CreditNoteAppliedTaxObject) GetLagoTaxIdOk() (*string, bool)`

GetLagoTaxIdOk returns a tuple with the LagoTaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagoTaxId

`func (o *CreditNoteAppliedTaxObject) SetLagoTaxId(v string)`

SetLagoTaxId sets LagoTaxId field to given value.

### HasLagoTaxId

`func (o *CreditNoteAppliedTaxObject) HasLagoTaxId() bool`

HasLagoTaxId returns a boolean if a field has been set.

### GetTaxName

`func (o *CreditNoteAppliedTaxObject) GetTaxName() string`

GetTaxName returns the TaxName field if non-nil, zero value otherwise.

### GetTaxNameOk

`func (o *CreditNoteAppliedTaxObject) GetTaxNameOk() (*string, bool)`

GetTaxNameOk returns a tuple with the TaxName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxName

`func (o *CreditNoteAppliedTaxObject) SetTaxName(v string)`

SetTaxName sets TaxName field to given value.

### HasTaxName

`func (o *CreditNoteAppliedTaxObject) HasTaxName() bool`

HasTaxName returns a boolean if a field has been set.

### GetTaxCode

`func (o *CreditNoteAppliedTaxObject) GetTaxCode() string`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *CreditNoteAppliedTaxObject) GetTaxCodeOk() (*string, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *CreditNoteAppliedTaxObject) SetTaxCode(v string)`

SetTaxCode sets TaxCode field to given value.

### HasTaxCode

`func (o *CreditNoteAppliedTaxObject) HasTaxCode() bool`

HasTaxCode returns a boolean if a field has been set.

### GetTaxRate

`func (o *CreditNoteAppliedTaxObject) GetTaxRate() float32`

GetTaxRate returns the TaxRate field if non-nil, zero value otherwise.

### GetTaxRateOk

`func (o *CreditNoteAppliedTaxObject) GetTaxRateOk() (*float32, bool)`

GetTaxRateOk returns a tuple with the TaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate

`func (o *CreditNoteAppliedTaxObject) SetTaxRate(v float32)`

SetTaxRate sets TaxRate field to given value.

### HasTaxRate

`func (o *CreditNoteAppliedTaxObject) HasTaxRate() bool`

HasTaxRate returns a boolean if a field has been set.

### GetTaxDescription

`func (o *CreditNoteAppliedTaxObject) GetTaxDescription() string`

GetTaxDescription returns the TaxDescription field if non-nil, zero value otherwise.

### GetTaxDescriptionOk

`func (o *CreditNoteAppliedTaxObject) GetTaxDescriptionOk() (*string, bool)`

GetTaxDescriptionOk returns a tuple with the TaxDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxDescription

`func (o *CreditNoteAppliedTaxObject) SetTaxDescription(v string)`

SetTaxDescription sets TaxDescription field to given value.

### HasTaxDescription

`func (o *CreditNoteAppliedTaxObject) HasTaxDescription() bool`

HasTaxDescription returns a boolean if a field has been set.

### GetAmountCents

`func (o *CreditNoteAppliedTaxObject) GetAmountCents() int32`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *CreditNoteAppliedTaxObject) GetAmountCentsOk() (*int32, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *CreditNoteAppliedTaxObject) SetAmountCents(v int32)`

SetAmountCents sets AmountCents field to given value.

### HasAmountCents

`func (o *CreditNoteAppliedTaxObject) HasAmountCents() bool`

HasAmountCents returns a boolean if a field has been set.

### GetAmountCurrency

`func (o *CreditNoteAppliedTaxObject) GetAmountCurrency() Currency`

GetAmountCurrency returns the AmountCurrency field if non-nil, zero value otherwise.

### GetAmountCurrencyOk

`func (o *CreditNoteAppliedTaxObject) GetAmountCurrencyOk() (*Currency, bool)`

GetAmountCurrencyOk returns a tuple with the AmountCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCurrency

`func (o *CreditNoteAppliedTaxObject) SetAmountCurrency(v Currency)`

SetAmountCurrency sets AmountCurrency field to given value.

### HasAmountCurrency

`func (o *CreditNoteAppliedTaxObject) HasAmountCurrency() bool`

HasAmountCurrency returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CreditNoteAppliedTaxObject) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreditNoteAppliedTaxObject) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreditNoteAppliedTaxObject) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CreditNoteAppliedTaxObject) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


