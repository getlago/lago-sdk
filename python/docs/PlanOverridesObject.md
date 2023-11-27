# PlanOverridesObject

Based plan overrides.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**amount_cents** | **int** | The base cost of the plan, excluding any applicable taxes, that is billed on a recurring basis. This value is defined at 0 if your plan is a pay-as-you-go plan. | [optional] 
**amount_currency** | [**Currency**](Currency.md) |  | [optional] 
**description** | **str** | The description on the plan. | [optional] 
**invoice_display_name** | **str** | Specifies the name that will be displayed on an invoice. If no value is set for this field, the name of the plan will be used as the default display name. | [optional] 
**name** | **str** | The name of the plan. | [optional] 
**tax_codes** | **List[str]** | List of unique code used to identify the taxes. | [optional] 
**trial_period** | **float** | The duration in days during which the base cost of the plan is offered for free. | [optional] 
**charges** | [**List[PlanOverridesObjectChargesInner]**](PlanOverridesObjectChargesInner.md) | Additional usage-based charges for this plan. | [optional] 

## Example

```python
from lago_api.models.plan_overrides_object import PlanOverridesObject

# TODO update the JSON string below
json = "{}"
# create an instance of PlanOverridesObject from a JSON string
plan_overrides_object_instance = PlanOverridesObject.from_json(json)
# print the JSON string representation of the object
print PlanOverridesObject.to_json()

# convert the object into a dict
plan_overrides_object_dict = plan_overrides_object_instance.to_dict()
# create an instance of PlanOverridesObject from a dict
plan_overrides_object_form_dict = plan_overrides_object.from_dict(plan_overrides_object_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


