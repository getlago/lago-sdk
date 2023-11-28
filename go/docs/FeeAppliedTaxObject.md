# FeeAppliedTaxObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LagoFeeId** | Pointer to **string** | Unique identifier of the fee, created by Lago. | [optional] 
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

### NewFeeAppliedTaxObject

`func NewFeeAppliedTaxObject() *FeeAppliedTaxObject`

NewFeeAppliedTaxObject instantiates a new FeeAppliedTaxObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeeAppliedTaxObjectWithDefaults

`func NewFeeAppliedTaxObjectWithDefaults() *FeeAppliedTaxObject`

NewFeeAppliedTaxObjectWithDefaults instantiates a new FeeAppliedTaxObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLagoFeeId

`func (o *FeeAppliedTaxObject) GetLagoFeeId() string`

GetLagoFeeId returns the LagoFeeId field if non-nil, zero value otherwise.

### GetLagoFeeIdOk

`func (o *FeeAppliedTaxObject) GetLagoFeeIdOk() (*string, bool)`

GetLagoFeeIdOk returns a tuple with the LagoFeeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagoFeeId

`func (o *FeeAppliedTaxObject) SetLagoFeeId(v string)`

SetLagoFeeId sets LagoFeeId field to given value.

### HasLagoFeeId

`func (o *FeeAppliedTaxObject) HasLagoFeeId() bool`

HasLagoFeeId returns a boolean if a field has been set.

### GetLagoId

`func (o *FeeAppliedTaxObject) GetLagoId() string`

GetLagoId returns the LagoId field if non-nil, zero value otherwise.

### GetLagoIdOk

`func (o *FeeAppliedTaxObject) GetLagoIdOk() (*string, bool)`

GetLagoIdOk returns a tuple with the LagoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagoId

`func (o *FeeAppliedTaxObject) SetLagoId(v string)`

SetLagoId sets LagoId field to given value.

### HasLagoId

`func (o *FeeAppliedTaxObject) HasLagoId() bool`

HasLagoId returns a boolean if a field has been set.

### GetLagoTaxId

`func (o *FeeAppliedTaxObject) GetLagoTaxId() string`

GetLagoTaxId returns the LagoTaxId field if non-nil, zero value otherwise.

### GetLagoTaxIdOk

`func (o *FeeAppliedTaxObject) GetLagoTaxIdOk() (*string, bool)`

GetLagoTaxIdOk returns a tuple with the LagoTaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagoTaxId

`func (o *FeeAppliedTaxObject) SetLagoTaxId(v string)`

SetLagoTaxId sets LagoTaxId field to given value.

### HasLagoTaxId

`func (o *FeeAppliedTaxObject) HasLagoTaxId() bool`

HasLagoTaxId returns a boolean if a field has been set.

### GetTaxName

`func (o *FeeAppliedTaxObject) GetTaxName() string`

GetTaxName returns the TaxName field if non-nil, zero value otherwise.

### GetTaxNameOk

`func (o *FeeAppliedTaxObject) GetTaxNameOk() (*string, bool)`

GetTaxNameOk returns a tuple with the TaxName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxName

`func (o *FeeAppliedTaxObject) SetTaxName(v string)`

SetTaxName sets TaxName field to given value.

### HasTaxName

`func (o *FeeAppliedTaxObject) HasTaxName() bool`

HasTaxName returns a boolean if a field has been set.

### GetTaxCode

`func (o *FeeAppliedTaxObject) GetTaxCode() string`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *FeeAppliedTaxObject) GetTaxCodeOk() (*string, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *FeeAppliedTaxObject) SetTaxCode(v string)`

SetTaxCode sets TaxCode field to given value.

### HasTaxCode

`func (o *FeeAppliedTaxObject) HasTaxCode() bool`

HasTaxCode returns a boolean if a field has been set.

### GetTaxRate

`func (o *FeeAppliedTaxObject) GetTaxRate() float32`

GetTaxRate returns the TaxRate field if non-nil, zero value otherwise.

### GetTaxRateOk

`func (o *FeeAppliedTaxObject) GetTaxRateOk() (*float32, bool)`

GetTaxRateOk returns a tuple with the TaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate

`func (o *FeeAppliedTaxObject) SetTaxRate(v float32)`

SetTaxRate sets TaxRate field to given value.

### HasTaxRate

`func (o *FeeAppliedTaxObject) HasTaxRate() bool`

HasTaxRate returns a boolean if a field has been set.

### GetTaxDescription

`func (o *FeeAppliedTaxObject) GetTaxDescription() string`

GetTaxDescription returns the TaxDescription field if non-nil, zero value otherwise.

### GetTaxDescriptionOk

`func (o *FeeAppliedTaxObject) GetTaxDescriptionOk() (*string, bool)`

GetTaxDescriptionOk returns a tuple with the TaxDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxDescription

`func (o *FeeAppliedTaxObject) SetTaxDescription(v string)`

SetTaxDescription sets TaxDescription field to given value.

### HasTaxDescription

`func (o *FeeAppliedTaxObject) HasTaxDescription() bool`

HasTaxDescription returns a boolean if a field has been set.

### GetAmountCents

`func (o *FeeAppliedTaxObject) GetAmountCents() int32`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *FeeAppliedTaxObject) GetAmountCentsOk() (*int32, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *FeeAppliedTaxObject) SetAmountCents(v int32)`

SetAmountCents sets AmountCents field to given value.

### HasAmountCents

`func (o *FeeAppliedTaxObject) HasAmountCents() bool`

HasAmountCents returns a boolean if a field has been set.

### GetAmountCurrency

`func (o *FeeAppliedTaxObject) GetAmountCurrency() Currency`

GetAmountCurrency returns the AmountCurrency field if non-nil, zero value otherwise.

### GetAmountCurrencyOk

`func (o *FeeAppliedTaxObject) GetAmountCurrencyOk() (*Currency, bool)`

GetAmountCurrencyOk returns a tuple with the AmountCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCurrency

`func (o *FeeAppliedTaxObject) SetAmountCurrency(v Currency)`

SetAmountCurrency sets AmountCurrency field to given value.

### HasAmountCurrency

`func (o *FeeAppliedTaxObject) HasAmountCurrency() bool`

HasAmountCurrency returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FeeAppliedTaxObject) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FeeAppliedTaxObject) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FeeAppliedTaxObject) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FeeAppliedTaxObject) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


