# BillableMetricsPaginated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillableMetrics** | [**[]BillableMetricObject**](BillableMetricObject.md) |  | 
**Meta** | [**PaginationMeta**](PaginationMeta.md) |  | 

## Methods

### NewBillableMetricsPaginated

`func NewBillableMetricsPaginated(billableMetrics []BillableMetricObject, meta PaginationMeta, ) *BillableMetricsPaginated`

NewBillableMetricsPaginated instantiates a new BillableMetricsPaginated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillableMetricsPaginatedWithDefaults

`func NewBillableMetricsPaginatedWithDefaults() *BillableMetricsPaginated`

NewBillableMetricsPaginatedWithDefaults instantiates a new BillableMetricsPaginated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillableMetrics

`func (o *BillableMetricsPaginated) GetBillableMetrics() []BillableMetricObject`

GetBillableMetrics returns the BillableMetrics field if non-nil, zero value otherwise.

### GetBillableMetricsOk

`func (o *BillableMetricsPaginated) GetBillableMetricsOk() (*[]BillableMetricObject, bool)`

GetBillableMetricsOk returns a tuple with the BillableMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillableMetrics

`func (o *BillableMetricsPaginated) SetBillableMetrics(v []BillableMetricObject)`

SetBillableMetrics sets BillableMetrics field to given value.


### GetMeta

`func (o *BillableMetricsPaginated) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *BillableMetricsPaginated) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *BillableMetricsPaginated) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


