# GrossRevenueObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Month** | **string** | Identifies the month to analyze revenue. | 
**AmountCents** | Pointer to **int32** | The total amount of revenue for a period, expressed in cents. | [optional] 
**Currency** | Pointer to [**Currency**](Currency.md) |  | [optional] 

## Methods

### NewGrossRevenueObject

`func NewGrossRevenueObject(month string, ) *GrossRevenueObject`

NewGrossRevenueObject instantiates a new GrossRevenueObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGrossRevenueObjectWithDefaults

`func NewGrossRevenueObjectWithDefaults() *GrossRevenueObject`

NewGrossRevenueObjectWithDefaults instantiates a new GrossRevenueObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonth

`func (o *GrossRevenueObject) GetMonth() string`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *GrossRevenueObject) GetMonthOk() (*string, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *GrossRevenueObject) SetMonth(v string)`

SetMonth sets Month field to given value.


### GetAmountCents

`func (o *GrossRevenueObject) GetAmountCents() int32`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *GrossRevenueObject) GetAmountCentsOk() (*int32, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *GrossRevenueObject) SetAmountCents(v int32)`

SetAmountCents sets AmountCents field to given value.

### HasAmountCents

`func (o *GrossRevenueObject) HasAmountCents() bool`

HasAmountCents returns a boolean if a field has been set.

### GetCurrency

`func (o *GrossRevenueObject) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GrossRevenueObject) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GrossRevenueObject) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *GrossRevenueObject) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


