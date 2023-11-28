# BillableMetricObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LagoId** | **string** | Unique identifier of the billable metric created by Lago. | 
**Name** | **string** | Name of the billable metric. | 
**Code** | **string** | Unique code used to identify the billable metric associated with the API request. This code associates each event with the correct metric. | 
**Description** | Pointer to **NullableString** | Internal description of the billable metric. | [optional] 
**Recurring** | **bool** | Defines if the billable metric is persisted billing period over billing period.  - If set to &#x60;true&#x60;: the accumulated number of units calculated from the previous billing period is persisted to the next billing period. - If set to &#x60;false&#x60;: the accumulated number of units is reset to 0 at the end of the billing period. - If not defined in the request, default value is &#x60;false&#x60;. | 
**CreatedAt** | **time.Time** | Creation date of the billable metric. | 
**FieldName** | Pointer to **NullableString** | Property of the billable metric used for aggregating usage data. This field is not required for &#x60;count_agg&#x60;. | [optional] 
**AggregationType** | **string** | Aggregation method used to compute usage for this billable metric. | 
**WeightedInterval** | Pointer to **NullableString** | Parameter exclusively utilized in conjunction with the &#x60;weighted_sum&#x60; aggregation type. It serves to adjust the aggregation result by assigning weights and proration to the result based on time intervals. When this field is not provided, the default time interval is assumed to be in &#x60;seconds&#x60;. | [optional] 
**Group** | Pointer to [**BillableMetricGroup**](BillableMetricGroup.md) |  | [optional] 
**ActiveSubscriptionsCount** | **int32** | Number of active subscriptions using this billable metric. | 
**DraftInvoicesCount** | **int32** | Number of draft invoices for which this billable metric is listed as an invoice item. | 
**PlansCount** | **int32** | Number of plans using this billable metric. | 

## Methods

### NewBillableMetricObject

`func NewBillableMetricObject(lagoId string, name string, code string, recurring bool, createdAt time.Time, aggregationType string, activeSubscriptionsCount int32, draftInvoicesCount int32, plansCount int32, ) *BillableMetricObject`

NewBillableMetricObject instantiates a new BillableMetricObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillableMetricObjectWithDefaults

`func NewBillableMetricObjectWithDefaults() *BillableMetricObject`

NewBillableMetricObjectWithDefaults instantiates a new BillableMetricObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLagoId

`func (o *BillableMetricObject) GetLagoId() string`

GetLagoId returns the LagoId field if non-nil, zero value otherwise.

### GetLagoIdOk

`func (o *BillableMetricObject) GetLagoIdOk() (*string, bool)`

GetLagoIdOk returns a tuple with the LagoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagoId

`func (o *BillableMetricObject) SetLagoId(v string)`

SetLagoId sets LagoId field to given value.


### GetName

`func (o *BillableMetricObject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BillableMetricObject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BillableMetricObject) SetName(v string)`

SetName sets Name field to given value.


### GetCode

`func (o *BillableMetricObject) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *BillableMetricObject) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *BillableMetricObject) SetCode(v string)`

SetCode sets Code field to given value.


### GetDescription

`func (o *BillableMetricObject) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BillableMetricObject) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BillableMetricObject) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BillableMetricObject) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BillableMetricObject) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BillableMetricObject) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetRecurring

`func (o *BillableMetricObject) GetRecurring() bool`

GetRecurring returns the Recurring field if non-nil, zero value otherwise.

### GetRecurringOk

`func (o *BillableMetricObject) GetRecurringOk() (*bool, bool)`

GetRecurringOk returns a tuple with the Recurring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurring

`func (o *BillableMetricObject) SetRecurring(v bool)`

SetRecurring sets Recurring field to given value.


### GetCreatedAt

`func (o *BillableMetricObject) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BillableMetricObject) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BillableMetricObject) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetFieldName

`func (o *BillableMetricObject) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *BillableMetricObject) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *BillableMetricObject) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.

### HasFieldName

`func (o *BillableMetricObject) HasFieldName() bool`

HasFieldName returns a boolean if a field has been set.

### SetFieldNameNil

`func (o *BillableMetricObject) SetFieldNameNil(b bool)`

 SetFieldNameNil sets the value for FieldName to be an explicit nil

### UnsetFieldName
`func (o *BillableMetricObject) UnsetFieldName()`

UnsetFieldName ensures that no value is present for FieldName, not even an explicit nil
### GetAggregationType

`func (o *BillableMetricObject) GetAggregationType() string`

GetAggregationType returns the AggregationType field if non-nil, zero value otherwise.

### GetAggregationTypeOk

`func (o *BillableMetricObject) GetAggregationTypeOk() (*string, bool)`

GetAggregationTypeOk returns a tuple with the AggregationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationType

`func (o *BillableMetricObject) SetAggregationType(v string)`

SetAggregationType sets AggregationType field to given value.


### GetWeightedInterval

`func (o *BillableMetricObject) GetWeightedInterval() string`

GetWeightedInterval returns the WeightedInterval field if non-nil, zero value otherwise.

### GetWeightedIntervalOk

`func (o *BillableMetricObject) GetWeightedIntervalOk() (*string, bool)`

GetWeightedIntervalOk returns a tuple with the WeightedInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightedInterval

`func (o *BillableMetricObject) SetWeightedInterval(v string)`

SetWeightedInterval sets WeightedInterval field to given value.

### HasWeightedInterval

`func (o *BillableMetricObject) HasWeightedInterval() bool`

HasWeightedInterval returns a boolean if a field has been set.

### SetWeightedIntervalNil

`func (o *BillableMetricObject) SetWeightedIntervalNil(b bool)`

 SetWeightedIntervalNil sets the value for WeightedInterval to be an explicit nil

### UnsetWeightedInterval
`func (o *BillableMetricObject) UnsetWeightedInterval()`

UnsetWeightedInterval ensures that no value is present for WeightedInterval, not even an explicit nil
### GetGroup

`func (o *BillableMetricObject) GetGroup() BillableMetricGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *BillableMetricObject) GetGroupOk() (*BillableMetricGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *BillableMetricObject) SetGroup(v BillableMetricGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *BillableMetricObject) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetActiveSubscriptionsCount

`func (o *BillableMetricObject) GetActiveSubscriptionsCount() int32`

GetActiveSubscriptionsCount returns the ActiveSubscriptionsCount field if non-nil, zero value otherwise.

### GetActiveSubscriptionsCountOk

`func (o *BillableMetricObject) GetActiveSubscriptionsCountOk() (*int32, bool)`

GetActiveSubscriptionsCountOk returns a tuple with the ActiveSubscriptionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveSubscriptionsCount

`func (o *BillableMetricObject) SetActiveSubscriptionsCount(v int32)`

SetActiveSubscriptionsCount sets ActiveSubscriptionsCount field to given value.


### GetDraftInvoicesCount

`func (o *BillableMetricObject) GetDraftInvoicesCount() int32`

GetDraftInvoicesCount returns the DraftInvoicesCount field if non-nil, zero value otherwise.

### GetDraftInvoicesCountOk

`func (o *BillableMetricObject) GetDraftInvoicesCountOk() (*int32, bool)`

GetDraftInvoicesCountOk returns a tuple with the DraftInvoicesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraftInvoicesCount

`func (o *BillableMetricObject) SetDraftInvoicesCount(v int32)`

SetDraftInvoicesCount sets DraftInvoicesCount field to given value.


### GetPlansCount

`func (o *BillableMetricObject) GetPlansCount() int32`

GetPlansCount returns the PlansCount field if non-nil, zero value otherwise.

### GetPlansCountOk

`func (o *BillableMetricObject) GetPlansCountOk() (*int32, bool)`

GetPlansCountOk returns a tuple with the PlansCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlansCount

`func (o *BillableMetricObject) SetPlansCount(v int32)`

SetPlansCount sets PlansCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


