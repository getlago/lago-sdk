# CustomerPastUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UsagePeriods** | [**[]CustomerUsage**](CustomerUsage.md) |  | 
**Meta** | [**PaginationMeta**](PaginationMeta.md) |  | 

## Methods

### NewCustomerPastUsage

`func NewCustomerPastUsage(usagePeriods []CustomerUsage, meta PaginationMeta, ) *CustomerPastUsage`

NewCustomerPastUsage instantiates a new CustomerPastUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerPastUsageWithDefaults

`func NewCustomerPastUsageWithDefaults() *CustomerPastUsage`

NewCustomerPastUsageWithDefaults instantiates a new CustomerPastUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsagePeriods

`func (o *CustomerPastUsage) GetUsagePeriods() []CustomerUsage`

GetUsagePeriods returns the UsagePeriods field if non-nil, zero value otherwise.

### GetUsagePeriodsOk

`func (o *CustomerPastUsage) GetUsagePeriodsOk() (*[]CustomerUsage, bool)`

GetUsagePeriodsOk returns a tuple with the UsagePeriods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsagePeriods

`func (o *CustomerPastUsage) SetUsagePeriods(v []CustomerUsage)`

SetUsagePeriods sets UsagePeriods field to given value.


### GetMeta

`func (o *CustomerPastUsage) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CustomerPastUsage) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CustomerPastUsage) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


