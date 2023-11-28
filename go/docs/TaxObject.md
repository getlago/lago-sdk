# TaxObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LagoId** | **string** | Unique identifier of the tax, created by Lago. | 
**Name** | **string** | Name of the tax. | 
**Code** | **string** | Unique code used to identify the tax associated with the API request. | 
**Description** | Pointer to **string** | Internal description of the taxe | [optional] 
**Rate** | **float32** | The percentage rate of the tax | 
**AppliedToOrganization** | **bool** | Set to &#x60;true&#x60; if the tax is used as one of the organization&#39;s default | 
**AddOnsCount** | Pointer to **int32** | Number of add-ons this tax is applied to. | [optional] 
**ChargesCount** | Pointer to **int32** | Number of charges this tax is applied to. | [optional] 
**CustomersCount** | **int32** | Number of customers this tax is applied to (directly or via the organization&#39;s default). | 
**PlansCount** | Pointer to **int32** | Number of plans this tax is applied to. | [optional] 
**CreatedAt** | **time.Time** | Creation date of the tax. | 

## Methods

### NewTaxObject

`func NewTaxObject(lagoId string, name string, code string, rate float32, appliedToOrganization bool, customersCount int32, createdAt time.Time, ) *TaxObject`

NewTaxObject instantiates a new TaxObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxObjectWithDefaults

`func NewTaxObjectWithDefaults() *TaxObject`

NewTaxObjectWithDefaults instantiates a new TaxObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLagoId

`func (o *TaxObject) GetLagoId() string`

GetLagoId returns the LagoId field if non-nil, zero value otherwise.

### GetLagoIdOk

`func (o *TaxObject) GetLagoIdOk() (*string, bool)`

GetLagoIdOk returns a tuple with the LagoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagoId

`func (o *TaxObject) SetLagoId(v string)`

SetLagoId sets LagoId field to given value.


### GetName

`func (o *TaxObject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TaxObject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TaxObject) SetName(v string)`

SetName sets Name field to given value.


### GetCode

`func (o *TaxObject) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *TaxObject) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *TaxObject) SetCode(v string)`

SetCode sets Code field to given value.


### GetDescription

`func (o *TaxObject) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TaxObject) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TaxObject) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TaxObject) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRate

`func (o *TaxObject) GetRate() float32`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *TaxObject) GetRateOk() (*float32, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *TaxObject) SetRate(v float32)`

SetRate sets Rate field to given value.


### GetAppliedToOrganization

`func (o *TaxObject) GetAppliedToOrganization() bool`

GetAppliedToOrganization returns the AppliedToOrganization field if non-nil, zero value otherwise.

### GetAppliedToOrganizationOk

`func (o *TaxObject) GetAppliedToOrganizationOk() (*bool, bool)`

GetAppliedToOrganizationOk returns a tuple with the AppliedToOrganization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedToOrganization

`func (o *TaxObject) SetAppliedToOrganization(v bool)`

SetAppliedToOrganization sets AppliedToOrganization field to given value.


### GetAddOnsCount

`func (o *TaxObject) GetAddOnsCount() int32`

GetAddOnsCount returns the AddOnsCount field if non-nil, zero value otherwise.

### GetAddOnsCountOk

`func (o *TaxObject) GetAddOnsCountOk() (*int32, bool)`

GetAddOnsCountOk returns a tuple with the AddOnsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddOnsCount

`func (o *TaxObject) SetAddOnsCount(v int32)`

SetAddOnsCount sets AddOnsCount field to given value.

### HasAddOnsCount

`func (o *TaxObject) HasAddOnsCount() bool`

HasAddOnsCount returns a boolean if a field has been set.

### GetChargesCount

`func (o *TaxObject) GetChargesCount() int32`

GetChargesCount returns the ChargesCount field if non-nil, zero value otherwise.

### GetChargesCountOk

`func (o *TaxObject) GetChargesCountOk() (*int32, bool)`

GetChargesCountOk returns a tuple with the ChargesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargesCount

`func (o *TaxObject) SetChargesCount(v int32)`

SetChargesCount sets ChargesCount field to given value.

### HasChargesCount

`func (o *TaxObject) HasChargesCount() bool`

HasChargesCount returns a boolean if a field has been set.

### GetCustomersCount

`func (o *TaxObject) GetCustomersCount() int32`

GetCustomersCount returns the CustomersCount field if non-nil, zero value otherwise.

### GetCustomersCountOk

`func (o *TaxObject) GetCustomersCountOk() (*int32, bool)`

GetCustomersCountOk returns a tuple with the CustomersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomersCount

`func (o *TaxObject) SetCustomersCount(v int32)`

SetCustomersCount sets CustomersCount field to given value.


### GetPlansCount

`func (o *TaxObject) GetPlansCount() int32`

GetPlansCount returns the PlansCount field if non-nil, zero value otherwise.

### GetPlansCountOk

`func (o *TaxObject) GetPlansCountOk() (*int32, bool)`

GetPlansCountOk returns a tuple with the PlansCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlansCount

`func (o *TaxObject) SetPlansCount(v int32)`

SetPlansCount sets PlansCount field to given value.

### HasPlansCount

`func (o *TaxObject) HasPlansCount() bool`

HasPlansCount returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TaxObject) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TaxObject) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TaxObject) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


