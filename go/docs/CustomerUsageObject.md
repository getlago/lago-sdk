# CustomerUsageObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromDatetime** | **time.Time** | The lower bound of the billing period, expressed in the ISO 8601 datetime format in Coordinated Universal Time (UTC). | 
**ToDatetime** | **time.Time** | The upper bound of the billing period, expressed in the ISO 8601 datetime format in Coordinated Universal Time (UTC). | 
**IssuingDate** | **time.Time** | The date of creation of the invoice. | 
**LagoInvoiceId** | Pointer to **NullableString** | A unique identifier associated with the invoice related to this particular usage record. | [optional] 
**Currency** | Pointer to [**Currency**](Currency.md) |  | [optional] 
**AmountCents** | **int32** | The amount in cents, tax excluded. | 
**TaxesAmountCents** | **int32** | The tax amount in cents. | 
**TotalAmountCents** | **int32** | The total amount in cents, tax included. | 
**ChargesUsage** | [**[]CustomerChargeUsageObject**](CustomerChargeUsageObject.md) | Array of charges that comprise the current usage. It contains detailed information about individual charge items associated with the usage. | 

## Methods

### NewCustomerUsageObject

`func NewCustomerUsageObject(fromDatetime time.Time, toDatetime time.Time, issuingDate time.Time, amountCents int32, taxesAmountCents int32, totalAmountCents int32, chargesUsage []CustomerChargeUsageObject, ) *CustomerUsageObject`

NewCustomerUsageObject instantiates a new CustomerUsageObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerUsageObjectWithDefaults

`func NewCustomerUsageObjectWithDefaults() *CustomerUsageObject`

NewCustomerUsageObjectWithDefaults instantiates a new CustomerUsageObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromDatetime

`func (o *CustomerUsageObject) GetFromDatetime() time.Time`

GetFromDatetime returns the FromDatetime field if non-nil, zero value otherwise.

### GetFromDatetimeOk

`func (o *CustomerUsageObject) GetFromDatetimeOk() (*time.Time, bool)`

GetFromDatetimeOk returns a tuple with the FromDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDatetime

`func (o *CustomerUsageObject) SetFromDatetime(v time.Time)`

SetFromDatetime sets FromDatetime field to given value.


### GetToDatetime

`func (o *CustomerUsageObject) GetToDatetime() time.Time`

GetToDatetime returns the ToDatetime field if non-nil, zero value otherwise.

### GetToDatetimeOk

`func (o *CustomerUsageObject) GetToDatetimeOk() (*time.Time, bool)`

GetToDatetimeOk returns a tuple with the ToDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToDatetime

`func (o *CustomerUsageObject) SetToDatetime(v time.Time)`

SetToDatetime sets ToDatetime field to given value.


### GetIssuingDate

`func (o *CustomerUsageObject) GetIssuingDate() time.Time`

GetIssuingDate returns the IssuingDate field if non-nil, zero value otherwise.

### GetIssuingDateOk

`func (o *CustomerUsageObject) GetIssuingDateOk() (*time.Time, bool)`

GetIssuingDateOk returns a tuple with the IssuingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingDate

`func (o *CustomerUsageObject) SetIssuingDate(v time.Time)`

SetIssuingDate sets IssuingDate field to given value.


### GetLagoInvoiceId

`func (o *CustomerUsageObject) GetLagoInvoiceId() string`

GetLagoInvoiceId returns the LagoInvoiceId field if non-nil, zero value otherwise.

### GetLagoInvoiceIdOk

`func (o *CustomerUsageObject) GetLagoInvoiceIdOk() (*string, bool)`

GetLagoInvoiceIdOk returns a tuple with the LagoInvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagoInvoiceId

`func (o *CustomerUsageObject) SetLagoInvoiceId(v string)`

SetLagoInvoiceId sets LagoInvoiceId field to given value.

### HasLagoInvoiceId

`func (o *CustomerUsageObject) HasLagoInvoiceId() bool`

HasLagoInvoiceId returns a boolean if a field has been set.

### SetLagoInvoiceIdNil

`func (o *CustomerUsageObject) SetLagoInvoiceIdNil(b bool)`

 SetLagoInvoiceIdNil sets the value for LagoInvoiceId to be an explicit nil

### UnsetLagoInvoiceId
`func (o *CustomerUsageObject) UnsetLagoInvoiceId()`

UnsetLagoInvoiceId ensures that no value is present for LagoInvoiceId, not even an explicit nil
### GetCurrency

`func (o *CustomerUsageObject) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CustomerUsageObject) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CustomerUsageObject) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CustomerUsageObject) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetAmountCents

`func (o *CustomerUsageObject) GetAmountCents() int32`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *CustomerUsageObject) GetAmountCentsOk() (*int32, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *CustomerUsageObject) SetAmountCents(v int32)`

SetAmountCents sets AmountCents field to given value.


### GetTaxesAmountCents

`func (o *CustomerUsageObject) GetTaxesAmountCents() int32`

GetTaxesAmountCents returns the TaxesAmountCents field if non-nil, zero value otherwise.

### GetTaxesAmountCentsOk

`func (o *CustomerUsageObject) GetTaxesAmountCentsOk() (*int32, bool)`

GetTaxesAmountCentsOk returns a tuple with the TaxesAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxesAmountCents

`func (o *CustomerUsageObject) SetTaxesAmountCents(v int32)`

SetTaxesAmountCents sets TaxesAmountCents field to given value.


### GetTotalAmountCents

`func (o *CustomerUsageObject) GetTotalAmountCents() int32`

GetTotalAmountCents returns the TotalAmountCents field if non-nil, zero value otherwise.

### GetTotalAmountCentsOk

`func (o *CustomerUsageObject) GetTotalAmountCentsOk() (*int32, bool)`

GetTotalAmountCentsOk returns a tuple with the TotalAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountCents

`func (o *CustomerUsageObject) SetTotalAmountCents(v int32)`

SetTotalAmountCents sets TotalAmountCents field to given value.


### GetChargesUsage

`func (o *CustomerUsageObject) GetChargesUsage() []CustomerChargeUsageObject`

GetChargesUsage returns the ChargesUsage field if non-nil, zero value otherwise.

### GetChargesUsageOk

`func (o *CustomerUsageObject) GetChargesUsageOk() (*[]CustomerChargeUsageObject, bool)`

GetChargesUsageOk returns a tuple with the ChargesUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargesUsage

`func (o *CustomerUsageObject) SetChargesUsage(v []CustomerChargeUsageObject)`

SetChargesUsage sets ChargesUsage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


