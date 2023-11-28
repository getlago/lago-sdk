# MrrObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Month** | **string** | Identifies the month to analyze MRR. | 
**AmountCents** | Pointer to **int32** | The total amount of MRR, expressed in cents. | [optional] 
**Currency** | Pointer to [**Currency**](Currency.md) |  | [optional] 

## Methods

### NewMrrObject

`func NewMrrObject(month string, ) *MrrObject`

NewMrrObject instantiates a new MrrObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMrrObjectWithDefaults

`func NewMrrObjectWithDefaults() *MrrObject`

NewMrrObjectWithDefaults instantiates a new MrrObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonth

`func (o *MrrObject) GetMonth() string`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *MrrObject) GetMonthOk() (*string, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *MrrObject) SetMonth(v string)`

SetMonth sets Month field to given value.


### GetAmountCents

`func (o *MrrObject) GetAmountCents() int32`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *MrrObject) GetAmountCentsOk() (*int32, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *MrrObject) SetAmountCents(v int32)`

SetAmountCents sets AmountCents field to given value.

### HasAmountCents

`func (o *MrrObject) HasAmountCents() bool`

HasAmountCents returns a boolean if a field has been set.

### GetCurrency

`func (o *MrrObject) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *MrrObject) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *MrrObject) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *MrrObject) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


