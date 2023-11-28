# AddOnsPaginated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddOns** | [**[]AddOnObject**](AddOnObject.md) |  | 
**Meta** | [**PaginationMeta**](PaginationMeta.md) |  | 

## Methods

### NewAddOnsPaginated

`func NewAddOnsPaginated(addOns []AddOnObject, meta PaginationMeta, ) *AddOnsPaginated`

NewAddOnsPaginated instantiates a new AddOnsPaginated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddOnsPaginatedWithDefaults

`func NewAddOnsPaginatedWithDefaults() *AddOnsPaginated`

NewAddOnsPaginatedWithDefaults instantiates a new AddOnsPaginated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddOns

`func (o *AddOnsPaginated) GetAddOns() []AddOnObject`

GetAddOns returns the AddOns field if non-nil, zero value otherwise.

### GetAddOnsOk

`func (o *AddOnsPaginated) GetAddOnsOk() (*[]AddOnObject, bool)`

GetAddOnsOk returns a tuple with the AddOns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddOns

`func (o *AddOnsPaginated) SetAddOns(v []AddOnObject)`

SetAddOns sets AddOns field to given value.


### GetMeta

`func (o *AddOnsPaginated) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AddOnsPaginated) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AddOnsPaginated) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


