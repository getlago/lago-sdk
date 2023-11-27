# PlanCreateInputPlanChargesInnerGroupPropertiesInner


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**group_id** | **str** | Unique identifier of a billable metric group, created by Lago. | 
**values** | [**GroupPropertiesObjectValues**](GroupPropertiesObjectValues.md) |  | 

## Example

```python
from lago_api.models.plan_create_input_plan_charges_inner_group_properties_inner import PlanCreateInputPlanChargesInnerGroupPropertiesInner

# TODO update the JSON string below
json = "{}"
# create an instance of PlanCreateInputPlanChargesInnerGroupPropertiesInner from a JSON string
plan_create_input_plan_charges_inner_group_properties_inner_instance = PlanCreateInputPlanChargesInnerGroupPropertiesInner.from_json(json)
# print the JSON string representation of the object
print PlanCreateInputPlanChargesInnerGroupPropertiesInner.to_json()

# convert the object into a dict
plan_create_input_plan_charges_inner_group_properties_inner_dict = plan_create_input_plan_charges_inner_group_properties_inner_instance.to_dict()
# create an instance of PlanCreateInputPlanChargesInnerGroupPropertiesInner from a dict
plan_create_input_plan_charges_inner_group_properties_inner_form_dict = plan_create_input_plan_charges_inner_group_properties_inner.from_dict(plan_create_input_plan_charges_inner_group_properties_inner_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


