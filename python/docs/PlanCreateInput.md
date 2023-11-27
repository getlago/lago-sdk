# PlanCreateInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**plan** | [**PlanCreateInputPlan**](PlanCreateInputPlan.md) |  | 

## Example

```python
from lago_api.models.plan_create_input import PlanCreateInput

# TODO update the JSON string below
json = "{}"
# create an instance of PlanCreateInput from a JSON string
plan_create_input_instance = PlanCreateInput.from_json(json)
# print the JSON string representation of the object
print PlanCreateInput.to_json()

# convert the object into a dict
plan_create_input_dict = plan_create_input_instance.to_dict()
# create an instance of PlanCreateInput from a dict
plan_create_input_form_dict = plan_create_input.from_dict(plan_create_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


