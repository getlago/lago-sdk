# PlanCreateInputPlanChargesInnerGroupPropertiesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | **string** | Unique identifier of a billable metric group, created by Lago. | 
**Values** | [**GroupPropertiesObjectValues**](GroupPropertiesObjectValues.md) |  | 

## Methods

### NewPlanCreateInputPlanChargesInnerGroupPropertiesInner

`func NewPlanCreateInputPlanChargesInnerGroupPropertiesInner(groupId string, values GroupPropertiesObjectValues, ) *PlanCreateInputPlanChargesInnerGroupPropertiesInner`

NewPlanCreateInputPlanChargesInnerGroupPropertiesInner instantiates a new PlanCreateInputPlanChargesInnerGroupPropertiesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanCreateInputPlanChargesInnerGroupPropertiesInnerWithDefaults

`func NewPlanCreateInputPlanChargesInnerGroupPropertiesInnerWithDefaults() *PlanCreateInputPlanChargesInnerGroupPropertiesInner`

NewPlanCreateInputPlanChargesInnerGroupPropertiesInnerWithDefaults instantiates a new PlanCreateInputPlanChargesInnerGroupPropertiesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *PlanCreateInputPlanChargesInnerGroupPropertiesInner) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *PlanCreateInputPlanChargesInnerGroupPropertiesInner) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *PlanCreateInputPlanChargesInnerGroupPropertiesInner) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetValues

`func (o *PlanCreateInputPlanChargesInnerGroupPropertiesInner) GetValues() GroupPropertiesObjectValues`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *PlanCreateInputPlanChargesInnerGroupPropertiesInner) GetValuesOk() (*GroupPropertiesObjectValues, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *PlanCreateInputPlanChargesInnerGroupPropertiesInner) SetValues(v GroupPropertiesObjectValues)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


