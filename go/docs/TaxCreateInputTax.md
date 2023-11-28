# TaxCreateInputTax

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the tax. | 
**Code** | **string** | Unique code used to identify the tax associated with the API request. | 
**Rate** | **string** | The percentage rate of the tax | 
**Description** | Pointer to **NullableString** | Internal description of the taxe | [optional] 
**AppliedToOrganization** | Pointer to **bool** | Set to &#x60;true&#x60; if the tax is used as one of the organization&#39;s default | [optional] 

## Methods

### NewTaxCreateInputTax

`func NewTaxCreateInputTax(name string, code string, rate string, ) *TaxCreateInputTax`

NewTaxCreateInputTax instantiates a new TaxCreateInputTax object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxCreateInputTaxWithDefaults

`func NewTaxCreateInputTaxWithDefaults() *TaxCreateInputTax`

NewTaxCreateInputTaxWithDefaults instantiates a new TaxCreateInputTax object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TaxCreateInputTax) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TaxCreateInputTax) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TaxCreateInputTax) SetName(v string)`

SetName sets Name field to given value.


### GetCode

`func (o *TaxCreateInputTax) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *TaxCreateInputTax) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *TaxCreateInputTax) SetCode(v string)`

SetCode sets Code field to given value.


### GetRate

`func (o *TaxCreateInputTax) GetRate() string`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *TaxCreateInputTax) GetRateOk() (*string, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *TaxCreateInputTax) SetRate(v string)`

SetRate sets Rate field to given value.


### GetDescription

`func (o *TaxCreateInputTax) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TaxCreateInputTax) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TaxCreateInputTax) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TaxCreateInputTax) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TaxCreateInputTax) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TaxCreateInputTax) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetAppliedToOrganization

`func (o *TaxCreateInputTax) GetAppliedToOrganization() bool`

GetAppliedToOrganization returns the AppliedToOrganization field if non-nil, zero value otherwise.

### GetAppliedToOrganizationOk

`func (o *TaxCreateInputTax) GetAppliedToOrganizationOk() (*bool, bool)`

GetAppliedToOrganizationOk returns a tuple with the AppliedToOrganization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedToOrganization

`func (o *TaxCreateInputTax) SetAppliedToOrganization(v bool)`

SetAppliedToOrganization sets AppliedToOrganization field to given value.

### HasAppliedToOrganization

`func (o *TaxCreateInputTax) HasAppliedToOrganization() bool`

HasAppliedToOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


