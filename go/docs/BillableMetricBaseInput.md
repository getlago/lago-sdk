# BillableMetricBaseInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the billable metric. | [optional] 
**Code** | Pointer to **string** | Unique code used to identify the billable metric associated with the API request. This code associates each event with the correct metric. | [optional] 
**Description** | Pointer to **NullableString** | Internal description of the billable metric. | [optional] 
**Recurring** | Pointer to **bool** | Defines if the billable metric is persisted billing period over billing period.  - If set to &#x60;true&#x60;: the accumulated number of units calculated from the previous billing period is persisted to the next billing period. - If set to &#x60;false&#x60;: the accumulated number of units is reset to 0 at the end of the billing period. - If not defined in the request, default value is &#x60;false&#x60;. | [optional] 
**FieldName** | Pointer to **NullableString** | Property of the billable metric used for aggregating usage data. This field is not required for &#x60;count_agg&#x60;. | [optional] 
**AggregationType** | Pointer to **string** | Aggregation method used to compute usage for this billable metric. | [optional] 
**WeightedInterval** | Pointer to **NullableString** | Parameter exclusively utilized in conjunction with the &#x60;weighted_sum&#x60; aggregation type. It serves to adjust the aggregation result by assigning weights and proration to the result based on time intervals. When this field is not provided, the default time interval is assumed to be in &#x60;seconds&#x60;. | [optional] 
**Group** | Pointer to [**BillableMetricGroup**](BillableMetricGroup.md) |  | [optional] 

## Methods

### NewBillableMetricBaseInput

`func NewBillableMetricBaseInput() *BillableMetricBaseInput`

NewBillableMetricBaseInput instantiates a new BillableMetricBaseInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillableMetricBaseInputWithDefaults

`func NewBillableMetricBaseInputWithDefaults() *BillableMetricBaseInput`

NewBillableMetricBaseInputWithDefaults instantiates a new BillableMetricBaseInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BillableMetricBaseInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BillableMetricBaseInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BillableMetricBaseInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BillableMetricBaseInput) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *BillableMetricBaseInput) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *BillableMetricBaseInput) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *BillableMetricBaseInput) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *BillableMetricBaseInput) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDescription

`func (o *BillableMetricBaseInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BillableMetricBaseInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BillableMetricBaseInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BillableMetricBaseInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BillableMetricBaseInput) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BillableMetricBaseInput) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetRecurring

`func (o *BillableMetricBaseInput) GetRecurring() bool`

GetRecurring returns the Recurring field if non-nil, zero value otherwise.

### GetRecurringOk

`func (o *BillableMetricBaseInput) GetRecurringOk() (*bool, bool)`

GetRecurringOk returns a tuple with the Recurring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurring

`func (o *BillableMetricBaseInput) SetRecurring(v bool)`

SetRecurring sets Recurring field to given value.

### HasRecurring

`func (o *BillableMetricBaseInput) HasRecurring() bool`

HasRecurring returns a boolean if a field has been set.

### GetFieldName

`func (o *BillableMetricBaseInput) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *BillableMetricBaseInput) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *BillableMetricBaseInput) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.

### HasFieldName

`func (o *BillableMetricBaseInput) HasFieldName() bool`

HasFieldName returns a boolean if a field has been set.

### SetFieldNameNil

`func (o *BillableMetricBaseInput) SetFieldNameNil(b bool)`

 SetFieldNameNil sets the value for FieldName to be an explicit nil

### UnsetFieldName
`func (o *BillableMetricBaseInput) UnsetFieldName()`

UnsetFieldName ensures that no value is present for FieldName, not even an explicit nil
### GetAggregationType

`func (o *BillableMetricBaseInput) GetAggregationType() string`

GetAggregationType returns the AggregationType field if non-nil, zero value otherwise.

### GetAggregationTypeOk

`func (o *BillableMetricBaseInput) GetAggregationTypeOk() (*string, bool)`

GetAggregationTypeOk returns a tuple with the AggregationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationType

`func (o *BillableMetricBaseInput) SetAggregationType(v string)`

SetAggregationType sets AggregationType field to given value.

### HasAggregationType

`func (o *BillableMetricBaseInput) HasAggregationType() bool`

HasAggregationType returns a boolean if a field has been set.

### GetWeightedInterval

`func (o *BillableMetricBaseInput) GetWeightedInterval() string`

GetWeightedInterval returns the WeightedInterval field if non-nil, zero value otherwise.

### GetWeightedIntervalOk

`func (o *BillableMetricBaseInput) GetWeightedIntervalOk() (*string, bool)`

GetWeightedIntervalOk returns a tuple with the WeightedInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightedInterval

`func (o *BillableMetricBaseInput) SetWeightedInterval(v string)`

SetWeightedInterval sets WeightedInterval field to given value.

### HasWeightedInterval

`func (o *BillableMetricBaseInput) HasWeightedInterval() bool`

HasWeightedInterval returns a boolean if a field has been set.

### SetWeightedIntervalNil

`func (o *BillableMetricBaseInput) SetWeightedIntervalNil(b bool)`

 SetWeightedIntervalNil sets the value for WeightedInterval to be an explicit nil

### UnsetWeightedInterval
`func (o *BillableMetricBaseInput) UnsetWeightedInterval()`

UnsetWeightedInterval ensures that no value is present for WeightedInterval, not even an explicit nil
### GetGroup

`func (o *BillableMetricBaseInput) GetGroup() BillableMetricGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *BillableMetricBaseInput) GetGroupOk() (*BillableMetricGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *BillableMetricBaseInput) SetGroup(v BillableMetricGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *BillableMetricBaseInput) HasGroup() bool`

HasGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


