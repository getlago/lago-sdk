# BaseAppliedTax

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewBaseAppliedTax

`func NewBaseAppliedTax() *BaseAppliedTax`

NewBaseAppliedTax instantiates a new BaseAppliedTax object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseAppliedTaxWithDefaults

`func NewBaseAppliedTaxWithDefaults() *BaseAppliedTax`

NewBaseAppliedTaxWithDefaults instantiates a new BaseAppliedTax object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLagoId

`func (o *BaseAppliedTax) GetLagoId() string`

GetLagoId returns the LagoId field if non-nil, zero value otherwise.

### GetLagoIdOk

`func (o *BaseAppliedTax) GetLagoIdOk() (*string, bool)`

GetLagoIdOk returns a tuple with the LagoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagoId

`func (o *BaseAppliedTax) SetLagoId(v string)`

SetLagoId sets LagoId field to given value.

### HasLagoId

`func (o *BaseAppliedTax) HasLagoId() bool`

HasLagoId returns a boolean if a field has been set.

### GetLagoTaxId

`func (o *BaseAppliedTax) GetLagoTaxId() string`

GetLagoTaxId returns the LagoTaxId field if non-nil, zero value otherwise.

### GetLagoTaxIdOk

`func (o *BaseAppliedTax) GetLagoTaxIdOk() (*string, bool)`

GetLagoTaxIdOk returns a tuple with the LagoTaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagoTaxId

`func (o *BaseAppliedTax) SetLagoTaxId(v string)`

SetLagoTaxId sets LagoTaxId field to given value.

### HasLagoTaxId

`func (o *BaseAppliedTax) HasLagoTaxId() bool`

HasLagoTaxId returns a boolean if a field has been set.

### GetTaxName

`func (o *BaseAppliedTax) GetTaxName() string`

GetTaxName returns the TaxName field if non-nil, zero value otherwise.

### GetTaxNameOk

`func (o *BaseAppliedTax) GetTaxNameOk() (*string, bool)`

GetTaxNameOk returns a tuple with the TaxName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxName

`func (o *BaseAppliedTax) SetTaxName(v string)`

SetTaxName sets TaxName field to given value.

### HasTaxName

`func (o *BaseAppliedTax) HasTaxName() bool`

HasTaxName returns a boolean if a field has been set.

### GetTaxCode

`func (o *BaseAppliedTax) GetTaxCode() string`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *BaseAppliedTax) GetTaxCodeOk() (*string, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *BaseAppliedTax) SetTaxCode(v string)`

SetTaxCode sets TaxCode field to given value.

### HasTaxCode

`func (o *BaseAppliedTax) HasTaxCode() bool`

HasTaxCode returns a boolean if a field has been set.

### GetTaxRate

`func (o *BaseAppliedTax) GetTaxRate() float32`

GetTaxRate returns the TaxRate field if non-nil, zero value otherwise.

### GetTaxRateOk

`func (o *BaseAppliedTax) GetTaxRateOk() (*float32, bool)`

GetTaxRateOk returns a tuple with the TaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate

`func (o *BaseAppliedTax) SetTaxRate(v float32)`

SetTaxRate sets TaxRate field to given value.

### HasTaxRate

`func (o *BaseAppliedTax) HasTaxRate() bool`

HasTaxRate returns a boolean if a field has been set.

### GetTaxDescription

`func (o *BaseAppliedTax) GetTaxDescription() string`

GetTaxDescription returns the TaxDescription field if non-nil, zero value otherwise.

### GetTaxDescriptionOk

`func (o *BaseAppliedTax) GetTaxDescriptionOk() (*string, bool)`

GetTaxDescriptionOk returns a tuple with the TaxDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxDescription

`func (o *BaseAppliedTax) SetTaxDescription(v string)`

SetTaxDescription sets TaxDescription field to given value.

### HasTaxDescription

`func (o *BaseAppliedTax) HasTaxDescription() bool`

HasTaxDescription returns a boolean if a field has been set.

### GetAmountCents

`func (o *BaseAppliedTax) GetAmountCents() int32`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *BaseAppliedTax) GetAmountCentsOk() (*int32, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *BaseAppliedTax) SetAmountCents(v int32)`

SetAmountCents sets AmountCents field to given value.

### HasAmountCents

`func (o *BaseAppliedTax) HasAmountCents() bool`

HasAmountCents returns a boolean if a field has been set.

### GetAmountCurrency

`func (o *BaseAppliedTax) GetAmountCurrency() Currency`

GetAmountCurrency returns the AmountCurrency field if non-nil, zero value otherwise.

### GetAmountCurrencyOk

`func (o *BaseAppliedTax) GetAmountCurrencyOk() (*Currency, bool)`

GetAmountCurrencyOk returns a tuple with the AmountCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCurrency

`func (o *BaseAppliedTax) SetAmountCurrency(v Currency)`

SetAmountCurrency sets AmountCurrency field to given value.

### HasAmountCurrency

`func (o *BaseAppliedTax) HasAmountCurrency() bool`

HasAmountCurrency returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BaseAppliedTax) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BaseAppliedTax) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BaseAppliedTax) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BaseAppliedTax) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


