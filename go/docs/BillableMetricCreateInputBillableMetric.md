# BillableMetricCreateInputBillableMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the billable metric. | 
**Code** | **string** | Unique code used to identify the billable metric associated with the API request. This code associates each event with the correct metric. | 
**Description** | Pointer to **NullableString** | Internal description of the billable metric. | [optional] 
**Recurring** | Pointer to **bool** | Defines if the billable metric is persisted billing period over billing period.  - If set to &#x60;true&#x60;: the accumulated number of units calculated from the previous billing period is persisted to the next billing period. - If set to &#x60;false&#x60;: the accumulated number of units is reset to 0 at the end of the billing period. - If not defined in the request, default value is &#x60;false&#x60;. | [optional] 
**FieldName** | Pointer to **NullableString** | Property of the billable metric used for aggregating usage data. This field is not required for &#x60;count_agg&#x60;. | [optional] 
**AggregationType** | **string** | Aggregation method used to compute usage for this billable metric. | 
**WeightedInterval** | Pointer to **NullableString** | Parameter exclusively utilized in conjunction with the &#x60;weighted_sum&#x60; aggregation type. It serves to adjust the aggregation result by assigning weights and proration to the result based on time intervals. When this field is not provided, the default time interval is assumed to be in &#x60;seconds&#x60;. | [optional] 
**Group** | Pointer to [**BillableMetricGroup**](BillableMetricGroup.md) |  | [optional] 

## Methods

### NewBillableMetricCreateInputBillableMetric

`func NewBillableMetricCreateInputBillableMetric(name string, code string, aggregationType string, ) *BillableMetricCreateInputBillableMetric`

NewBillableMetricCreateInputBillableMetric instantiates a new BillableMetricCreateInputBillableMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillableMetricCreateInputBillableMetricWithDefaults

`func NewBillableMetricCreateInputBillableMetricWithDefaults() *BillableMetricCreateInputBillableMetric`

NewBillableMetricCreateInputBillableMetricWithDefaults instantiates a new BillableMetricCreateInputBillableMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BillableMetricCreateInputBillableMetric) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BillableMetricCreateInputBillableMetric) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BillableMetricCreateInputBillableMetric) SetName(v string)`

SetName sets Name field to given value.


### GetCode

`func (o *BillableMetricCreateInputBillableMetric) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *BillableMetricCreateInputBillableMetric) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *BillableMetricCreateInputBillableMetric) SetCode(v string)`

SetCode sets Code field to given value.


### GetDescription

`func (o *BillableMetricCreateInputBillableMetric) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BillableMetricCreateInputBillableMetric) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BillableMetricCreateInputBillableMetric) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BillableMetricCreateInputBillableMetric) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BillableMetricCreateInputBillableMetric) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BillableMetricCreateInputBillableMetric) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetRecurring

`func (o *BillableMetricCreateInputBillableMetric) GetRecurring() bool`

GetRecurring returns the Recurring field if non-nil, zero value otherwise.

### GetRecurringOk

`func (o *BillableMetricCreateInputBillableMetric) GetRecurringOk() (*bool, bool)`

GetRecurringOk returns a tuple with the Recurring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurring

`func (o *BillableMetricCreateInputBillableMetric) SetRecurring(v bool)`

SetRecurring sets Recurring field to given value.

### HasRecurring

`func (o *BillableMetricCreateInputBillableMetric) HasRecurring() bool`

HasRecurring returns a boolean if a field has been set.

### GetFieldName

`func (o *BillableMetricCreateInputBillableMetric) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *BillableMetricCreateInputBillableMetric) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *BillableMetricCreateInputBillableMetric) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.

### HasFieldName

`func (o *BillableMetricCreateInputBillableMetric) HasFieldName() bool`

HasFieldName returns a boolean if a field has been set.

### SetFieldNameNil

`func (o *BillableMetricCreateInputBillableMetric) SetFieldNameNil(b bool)`

 SetFieldNameNil sets the value for FieldName to be an explicit nil

### UnsetFieldName
`func (o *BillableMetricCreateInputBillableMetric) UnsetFieldName()`

UnsetFieldName ensures that no value is present for FieldName, not even an explicit nil
### GetAggregationType

`func (o *BillableMetricCreateInputBillableMetric) GetAggregationType() string`

GetAggregationType returns the AggregationType field if non-nil, zero value otherwise.

### GetAggregationTypeOk

`func (o *BillableMetricCreateInputBillableMetric) GetAggregationTypeOk() (*string, bool)`

GetAggregationTypeOk returns a tuple with the AggregationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationType

`func (o *BillableMetricCreateInputBillableMetric) SetAggregationType(v string)`

SetAggregationType sets AggregationType field to given value.


### GetWeightedInterval

`func (o *BillableMetricCreateInputBillableMetric) GetWeightedInterval() string`

GetWeightedInterval returns the WeightedInterval field if non-nil, zero value otherwise.

### GetWeightedIntervalOk

`func (o *BillableMetricCreateInputBillableMetric) GetWeightedIntervalOk() (*string, bool)`

GetWeightedIntervalOk returns a tuple with the WeightedInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightedInterval

`func (o *BillableMetricCreateInputBillableMetric) SetWeightedInterval(v string)`

SetWeightedInterval sets WeightedInterval field to given value.

### HasWeightedInterval

`func (o *BillableMetricCreateInputBillableMetric) HasWeightedInterval() bool`

HasWeightedInterval returns a boolean if a field has been set.

### SetWeightedIntervalNil

`func (o *BillableMetricCreateInputBillableMetric) SetWeightedIntervalNil(b bool)`

 SetWeightedIntervalNil sets the value for WeightedInterval to be an explicit nil

### UnsetWeightedInterval
`func (o *BillableMetricCreateInputBillableMetric) UnsetWeightedInterval()`

UnsetWeightedInterval ensures that no value is present for WeightedInterval, not even an explicit nil
### GetGroup

`func (o *BillableMetricCreateInputBillableMetric) GetGroup() BillableMetricGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *BillableMetricCreateInputBillableMetric) GetGroupOk() (*BillableMetricGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *BillableMetricCreateInputBillableMetric) SetGroup(v BillableMetricGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *BillableMetricCreateInputBillableMetric) HasGroup() bool`

HasGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


