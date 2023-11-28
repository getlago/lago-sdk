# TaxesPaginated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Taxes** | [**[]TaxObject**](TaxObject.md) |  | 
**Meta** | [**PaginationMeta**](PaginationMeta.md) |  | 

## Methods

### NewTaxesPaginated

`func NewTaxesPaginated(taxes []TaxObject, meta PaginationMeta, ) *TaxesPaginated`

NewTaxesPaginated instantiates a new TaxesPaginated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxesPaginatedWithDefaults

`func NewTaxesPaginatedWithDefaults() *TaxesPaginated`

NewTaxesPaginatedWithDefaults instantiates a new TaxesPaginated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaxes

`func (o *TaxesPaginated) GetTaxes() []TaxObject`

GetTaxes returns the Taxes field if non-nil, zero value otherwise.

### GetTaxesOk

`func (o *TaxesPaginated) GetTaxesOk() (*[]TaxObject, bool)`

GetTaxesOk returns a tuple with the Taxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxes

`func (o *TaxesPaginated) SetTaxes(v []TaxObject)`

SetTaxes sets Taxes field to given value.


### GetMeta

`func (o *TaxesPaginated) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *TaxesPaginated) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *TaxesPaginated) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


