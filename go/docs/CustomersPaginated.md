# CustomersPaginated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Customers** | [**[]CustomerObjectExtended**](CustomerObjectExtended.md) |  | 
**Meta** | [**PaginationMeta**](PaginationMeta.md) |  | 

## Methods

### NewCustomersPaginated

`func NewCustomersPaginated(customers []CustomerObjectExtended, meta PaginationMeta, ) *CustomersPaginated`

NewCustomersPaginated instantiates a new CustomersPaginated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomersPaginatedWithDefaults

`func NewCustomersPaginatedWithDefaults() *CustomersPaginated`

NewCustomersPaginatedWithDefaults instantiates a new CustomersPaginated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomers

`func (o *CustomersPaginated) GetCustomers() []CustomerObjectExtended`

GetCustomers returns the Customers field if non-nil, zero value otherwise.

### GetCustomersOk

`func (o *CustomersPaginated) GetCustomersOk() (*[]CustomerObjectExtended, bool)`

GetCustomersOk returns a tuple with the Customers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomers

`func (o *CustomersPaginated) SetCustomers(v []CustomerObjectExtended)`

SetCustomers sets Customers field to given value.


### GetMeta

`func (o *CustomersPaginated) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CustomersPaginated) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CustomersPaginated) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


