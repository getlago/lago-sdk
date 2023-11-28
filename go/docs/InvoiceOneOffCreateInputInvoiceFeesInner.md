# InvoiceOneOffCreateInputInvoiceFeesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddOnCode** | **string** | The code of the add-on used as invoice item. | 
**InvoiceDisplayName** | Pointer to **string** | Specifies the name that will be displayed on an invoice. If no value is set for this field, the name of the actual charge will be used as the default display name. | [optional] 
**UnitAmountCents** | Pointer to **NullableInt32** | The amount of the fee per unit, expressed in cents. By default, the amount of the add-on is used. | [optional] 
**Units** | Pointer to **NullableString** | The quantity of units associated with the fee. By default, only 1 unit is added to the invoice. | [optional] 
**Description** | Pointer to **NullableString** | This is a description | [optional] 
**TaxCodes** | Pointer to **[]string** | List of unique code used to identify the taxes. | [optional] 

## Methods

### NewInvoiceOneOffCreateInputInvoiceFeesInner

`func NewInvoiceOneOffCreateInputInvoiceFeesInner(addOnCode string, ) *InvoiceOneOffCreateInputInvoiceFeesInner`

NewInvoiceOneOffCreateInputInvoiceFeesInner instantiates a new InvoiceOneOffCreateInputInvoiceFeesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceOneOffCreateInputInvoiceFeesInnerWithDefaults

`func NewInvoiceOneOffCreateInputInvoiceFeesInnerWithDefaults() *InvoiceOneOffCreateInputInvoiceFeesInner`

NewInvoiceOneOffCreateInputInvoiceFeesInnerWithDefaults instantiates a new InvoiceOneOffCreateInputInvoiceFeesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddOnCode

`func (o *InvoiceOneOffCreateInputInvoiceFeesInner) GetAddOnCode() string`

GetAddOnCode returns the AddOnCode field if non-nil, zero value otherwise.

### GetAddOnCodeOk

`func (o *InvoiceOneOffCreateInputInvoiceFeesInner) GetAddOnCodeOk() (*string, bool)`

GetAddOnCodeOk returns a tuple with the AddOnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddOnCode

`func (o *InvoiceOneOffCreateInputInvoiceFeesInner) SetAddOnCode(v string)`

SetAddOnCode sets AddOnCode field to given value.


### GetInvoiceDisplayName

`func (o *InvoiceOneOffCreateInputInvoiceFeesInner) GetInvoiceDisplayName() string`

GetInvoiceDisplayName returns the InvoiceDisplayName field if non-nil, zero value otherwise.

### GetInvoiceDisplayNameOk

`func (o *InvoiceOneOffCreateInputInvoiceFeesInner) GetInvoiceDisplayNameOk() (*string, bool)`

GetInvoiceDisplayNameOk returns a tuple with the InvoiceDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceDisplayName

`func (o *InvoiceOneOffCreateInputInvoiceFeesInner) SetInvoiceDisplayName(v string)`

SetInvoiceDisplayName sets InvoiceDisplayName field to given value.

### HasInvoiceDisplayName

`func (o *InvoiceOneOffCreateInputInvoiceFeesInner) HasInvoiceDisplayName() bool`

HasInvoiceDisplayName returns a boolean if a field has been set.

### GetUnitAmountCents

`func (o *InvoiceOneOffCreateInputInvoiceFeesInner) GetUnitAmountCents() int32`

GetUnitAmountCents returns the UnitAmountCents field if non-nil, zero value otherwise.

### GetUnitAmountCentsOk

`func (o *InvoiceOneOffCreateInputInvoiceFeesInner) GetUnitAmountCentsOk() (*int32, bool)`

GetUnitAmountCentsOk returns a tuple with the UnitAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAmountCents

`func (o *InvoiceOneOffCreateInputInvoiceFeesInner) SetUnitAmountCents(v int32)`

SetUnitAmountCents sets UnitAmountCents field to given value.

### HasUnitAmountCents

`func (o *InvoiceOneOffCreateInputInvoiceFeesInner) HasUnitAmountCents() bool`

HasUnitAmountCents returns a boolean if a field has been set.

### SetUnitAmountCentsNil

`func (o *InvoiceOneOffCreateInputInvoiceFeesInner) SetUnitAmountCentsNil(b bool)`

 SetUnitAmountCentsNil sets the value for UnitAmountCents to be an explicit nil

### UnsetUnitAmountCents
`func (o *InvoiceOneOffCreateInputInvoiceFeesInner) UnsetUnitAmountCents()`

UnsetUnitAmountCents ensures that no value is present for UnitAmountCents, not even an explicit nil
### GetUnits

`func (o *InvoiceOneOffCreateInputInvoiceFeesInner) GetUnits() string`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *InvoiceOneOffCreateInputInvoiceFeesInner) GetUnitsOk() (*string, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *InvoiceOneOffCreateInputInvoiceFeesInner) SetUnits(v string)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *InvoiceOneOffCreateInputInvoiceFeesInner) HasUnits() bool`

HasUnits returns a boolean if a field has been set.

### SetUnitsNil

`func (o *InvoiceOneOffCreateInputInvoiceFeesInner) SetUnitsNil(b bool)`

 SetUnitsNil sets the value for Units to be an explicit nil

### UnsetUnits
`func (o *InvoiceOneOffCreateInputInvoiceFeesInner) UnsetUnits()`

UnsetUnits ensures that no value is present for Units, not even an explicit nil
### GetDescription

`func (o *InvoiceOneOffCreateInputInvoiceFeesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InvoiceOneOffCreateInputInvoiceFeesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InvoiceOneOffCreateInputInvoiceFeesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InvoiceOneOffCreateInputInvoiceFeesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *InvoiceOneOffCreateInputInvoiceFeesInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *InvoiceOneOffCreateInputInvoiceFeesInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetTaxCodes

`func (o *InvoiceOneOffCreateInputInvoiceFeesInner) GetTaxCodes() []string`

GetTaxCodes returns the TaxCodes field if non-nil, zero value otherwise.

### GetTaxCodesOk

`func (o *InvoiceOneOffCreateInputInvoiceFeesInner) GetTaxCodesOk() (*[]string, bool)`

GetTaxCodesOk returns a tuple with the TaxCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCodes

`func (o *InvoiceOneOffCreateInputInvoiceFeesInner) SetTaxCodes(v []string)`

SetTaxCodes sets TaxCodes field to given value.

### HasTaxCodes

`func (o *InvoiceOneOffCreateInputInvoiceFeesInner) HasTaxCodes() bool`

HasTaxCodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


