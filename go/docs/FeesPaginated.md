# FeesPaginated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fees** | [**[]FeeObject**](FeeObject.md) |  | 
**Meta** | [**PaginationMeta**](PaginationMeta.md) |  | 

## Methods

### NewFeesPaginated

`func NewFeesPaginated(fees []FeeObject, meta PaginationMeta, ) *FeesPaginated`

NewFeesPaginated instantiates a new FeesPaginated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeesPaginatedWithDefaults

`func NewFeesPaginatedWithDefaults() *FeesPaginated`

NewFeesPaginatedWithDefaults instantiates a new FeesPaginated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFees

`func (o *FeesPaginated) GetFees() []FeeObject`

GetFees returns the Fees field if non-nil, zero value otherwise.

### GetFeesOk

`func (o *FeesPaginated) GetFeesOk() (*[]FeeObject, bool)`

GetFeesOk returns a tuple with the Fees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFees

`func (o *FeesPaginated) SetFees(v []FeeObject)`

SetFees sets Fees field to given value.


### GetMeta

`func (o *FeesPaginated) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *FeesPaginated) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *FeesPaginated) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


