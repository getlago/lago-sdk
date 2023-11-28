# CustomerChargeUsageObjectBillableMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LagoId** | **string** | Unique identifier assigned to the billable metric within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the billable metric’s record within the Lago system. | 
**Name** | **string** | The name of the billable metric used for this charge. | 
**Code** | **string** | The code of the billable metric used for this charge. | 
**AggregationType** | **string** | The aggregation type of the billable metric used for this charge. Possible values are &#x60;count_agg&#x60;, &#x60;sum_agg&#x60;, &#x60;max_agg&#x60; or &#x60;unique_count_agg&#x60;. | 

## Methods

### NewCustomerChargeUsageObjectBillableMetric

`func NewCustomerChargeUsageObjectBillableMetric(lagoId string, name string, code string, aggregationType string, ) *CustomerChargeUsageObjectBillableMetric`

NewCustomerChargeUsageObjectBillableMetric instantiates a new CustomerChargeUsageObjectBillableMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerChargeUsageObjectBillableMetricWithDefaults

`func NewCustomerChargeUsageObjectBillableMetricWithDefaults() *CustomerChargeUsageObjectBillableMetric`

NewCustomerChargeUsageObjectBillableMetricWithDefaults instantiates a new CustomerChargeUsageObjectBillableMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLagoId

`func (o *CustomerChargeUsageObjectBillableMetric) GetLagoId() string`

GetLagoId returns the LagoId field if non-nil, zero value otherwise.

### GetLagoIdOk

`func (o *CustomerChargeUsageObjectBillableMetric) GetLagoIdOk() (*string, bool)`

GetLagoIdOk returns a tuple with the LagoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagoId

`func (o *CustomerChargeUsageObjectBillableMetric) SetLagoId(v string)`

SetLagoId sets LagoId field to given value.


### GetName

`func (o *CustomerChargeUsageObjectBillableMetric) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomerChargeUsageObjectBillableMetric) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomerChargeUsageObjectBillableMetric) SetName(v string)`

SetName sets Name field to given value.


### GetCode

`func (o *CustomerChargeUsageObjectBillableMetric) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CustomerChargeUsageObjectBillableMetric) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CustomerChargeUsageObjectBillableMetric) SetCode(v string)`

SetCode sets Code field to given value.


### GetAggregationType

`func (o *CustomerChargeUsageObjectBillableMetric) GetAggregationType() string`

GetAggregationType returns the AggregationType field if non-nil, zero value otherwise.

### GetAggregationTypeOk

`func (o *CustomerChargeUsageObjectBillableMetric) GetAggregationTypeOk() (*string, bool)`

GetAggregationTypeOk returns a tuple with the AggregationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationType

`func (o *CustomerChargeUsageObjectBillableMetric) SetAggregationType(v string)`

SetAggregationType sets AggregationType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


