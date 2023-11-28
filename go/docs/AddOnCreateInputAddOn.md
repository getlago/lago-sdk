# AddOnCreateInputAddOn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the add-on. | 
**InvoiceDisplayName** | Pointer to **string** | Specifies the name that will be displayed on an invoice. If no value is set for this field, the name of the actual charge will be used as the default display name. | [optional] 
**Code** | **string** | Unique code used to identify the add-on. | 
**AmountCents** | **int32** | The cost of the add-on in cents, excluding any applicable taxes, that is billed to a customer. By creating a one-off invoice, you will be able to override this value. | 
**AmountCurrency** | [**Currency**](Currency.md) |  | 
**Description** | Pointer to **NullableString** | The description of the add-on. | [optional] 
**TaxCodes** | Pointer to **[]string** | List of unique code used to identify the taxes. | [optional] 

## Methods

### NewAddOnCreateInputAddOn

`func NewAddOnCreateInputAddOn(name string, code string, amountCents int32, amountCurrency Currency, ) *AddOnCreateInputAddOn`

NewAddOnCreateInputAddOn instantiates a new AddOnCreateInputAddOn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddOnCreateInputAddOnWithDefaults

`func NewAddOnCreateInputAddOnWithDefaults() *AddOnCreateInputAddOn`

NewAddOnCreateInputAddOnWithDefaults instantiates a new AddOnCreateInputAddOn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddOnCreateInputAddOn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddOnCreateInputAddOn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddOnCreateInputAddOn) SetName(v string)`

SetName sets Name field to given value.


### GetInvoiceDisplayName

`func (o *AddOnCreateInputAddOn) GetInvoiceDisplayName() string`

GetInvoiceDisplayName returns the InvoiceDisplayName field if non-nil, zero value otherwise.

### GetInvoiceDisplayNameOk

`func (o *AddOnCreateInputAddOn) GetInvoiceDisplayNameOk() (*string, bool)`

GetInvoiceDisplayNameOk returns a tuple with the InvoiceDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceDisplayName

`func (o *AddOnCreateInputAddOn) SetInvoiceDisplayName(v string)`

SetInvoiceDisplayName sets InvoiceDisplayName field to given value.

### HasInvoiceDisplayName

`func (o *AddOnCreateInputAddOn) HasInvoiceDisplayName() bool`

HasInvoiceDisplayName returns a boolean if a field has been set.

### GetCode

`func (o *AddOnCreateInputAddOn) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AddOnCreateInputAddOn) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AddOnCreateInputAddOn) SetCode(v string)`

SetCode sets Code field to given value.


### GetAmountCents

`func (o *AddOnCreateInputAddOn) GetAmountCents() int32`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *AddOnCreateInputAddOn) GetAmountCentsOk() (*int32, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *AddOnCreateInputAddOn) SetAmountCents(v int32)`

SetAmountCents sets AmountCents field to given value.


### GetAmountCurrency

`func (o *AddOnCreateInputAddOn) GetAmountCurrency() Currency`

GetAmountCurrency returns the AmountCurrency field if non-nil, zero value otherwise.

### GetAmountCurrencyOk

`func (o *AddOnCreateInputAddOn) GetAmountCurrencyOk() (*Currency, bool)`

GetAmountCurrencyOk returns a tuple with the AmountCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCurrency

`func (o *AddOnCreateInputAddOn) SetAmountCurrency(v Currency)`

SetAmountCurrency sets AmountCurrency field to given value.


### GetDescription

`func (o *AddOnCreateInputAddOn) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddOnCreateInputAddOn) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddOnCreateInputAddOn) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddOnCreateInputAddOn) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AddOnCreateInputAddOn) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AddOnCreateInputAddOn) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetTaxCodes

`func (o *AddOnCreateInputAddOn) GetTaxCodes() []string`

GetTaxCodes returns the TaxCodes field if non-nil, zero value otherwise.

### GetTaxCodesOk

`func (o *AddOnCreateInputAddOn) GetTaxCodesOk() (*[]string, bool)`

GetTaxCodesOk returns a tuple with the TaxCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCodes

`func (o *AddOnCreateInputAddOn) SetTaxCodes(v []string)`

SetTaxCodes sets TaxCodes field to given value.

### HasTaxCodes

`func (o *AddOnCreateInputAddOn) HasTaxCodes() bool`

HasTaxCodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


