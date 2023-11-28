# TaxBaseInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the tax. | [optional] 
**Code** | Pointer to **string** | Unique code used to identify the tax associated with the API request. | [optional] 
**Rate** | Pointer to **string** | The percentage rate of the tax | [optional] 
**Description** | Pointer to **NullableString** | Internal description of the taxe | [optional] 
**AppliedToOrganization** | Pointer to **bool** | Set to &#x60;true&#x60; if the tax is used as one of the organization&#39;s default | [optional] 

## Methods

### NewTaxBaseInput

`func NewTaxBaseInput() *TaxBaseInput`

NewTaxBaseInput instantiates a new TaxBaseInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxBaseInputWithDefaults

`func NewTaxBaseInputWithDefaults() *TaxBaseInput`

NewTaxBaseInputWithDefaults instantiates a new TaxBaseInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TaxBaseInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TaxBaseInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TaxBaseInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TaxBaseInput) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *TaxBaseInput) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *TaxBaseInput) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *TaxBaseInput) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *TaxBaseInput) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetRate

`func (o *TaxBaseInput) GetRate() string`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *TaxBaseInput) GetRateOk() (*string, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *TaxBaseInput) SetRate(v string)`

SetRate sets Rate field to given value.

### HasRate

`func (o *TaxBaseInput) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetDescription

`func (o *TaxBaseInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TaxBaseInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TaxBaseInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TaxBaseInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TaxBaseInput) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TaxBaseInput) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetAppliedToOrganization

`func (o *TaxBaseInput) GetAppliedToOrganization() bool`

GetAppliedToOrganization returns the AppliedToOrganization field if non-nil, zero value otherwise.

### GetAppliedToOrganizationOk

`func (o *TaxBaseInput) GetAppliedToOrganizationOk() (*bool, bool)`

GetAppliedToOrganizationOk returns a tuple with the AppliedToOrganization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedToOrganization

`func (o *TaxBaseInput) SetAppliedToOrganization(v bool)`

SetAppliedToOrganization sets AppliedToOrganization field to given value.

### HasAppliedToOrganization

`func (o *TaxBaseInput) HasAppliedToOrganization() bool`

HasAppliedToOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


