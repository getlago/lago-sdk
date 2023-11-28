# EventObjectProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OperationType** | Pointer to **string** | The &#x60;operation_type&#x60; field is only necessary when adding or removing a specific unit when the targeted billable metric adopts a &#x60;unique_count_agg&#x60; aggregation method. In other cases, the &#x60;operation_type&#x60; field is not required. The valid values for the &#x60;operation_type&#x60; field are &#x60;add&#x60; or &#x60;remove&#x60;, which indicate whether the unit is being added or removed from the unique count aggregation, respectively. | [optional] 

## Methods

### NewEventObjectProperties

`func NewEventObjectProperties() *EventObjectProperties`

NewEventObjectProperties instantiates a new EventObjectProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventObjectPropertiesWithDefaults

`func NewEventObjectPropertiesWithDefaults() *EventObjectProperties`

NewEventObjectPropertiesWithDefaults instantiates a new EventObjectProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperationType

`func (o *EventObjectProperties) GetOperationType() string`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *EventObjectProperties) GetOperationTypeOk() (*string, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *EventObjectProperties) SetOperationType(v string)`

SetOperationType sets OperationType field to given value.

### HasOperationType

`func (o *EventObjectProperties) HasOperationType() bool`

HasOperationType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


