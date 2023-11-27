# PlanUpdateInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**plan** | [**PlanUpdateInputPlan**](PlanUpdateInputPlan.md) |  | 

## Example

```python
from lago_api.models.plan_update_input import PlanUpdateInput

# TODO update the JSON string below
json = "{}"
# create an instance of PlanUpdateInput from a JSON string
plan_update_input_instance = PlanUpdateInput.from_json(json)
# print the JSON string representation of the object
print PlanUpdateInput.to_json()

# convert the object into a dict
plan_update_input_dict = plan_update_input_instance.to_dict()
# create an instance of PlanUpdateInput from a dict
plan_update_input_form_dict = plan_update_input.from_dict(plan_update_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


