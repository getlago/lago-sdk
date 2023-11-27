# PlanOverridesObjectChargesInner


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **str** | Unique identifier of the charge created by Lago. | [optional] 
**billable_metric_id** | **str** | Unique identifier of the billable metric created by Lago. | [optional] 
**invoice_display_name** | **str** | Specifies the name that will be displayed on an invoice. If no value is set for this field, the name of the actual charge will be used as the default display name. | [optional] 
**min_amount_cents** | **int** | The minimum spending amount required for the charge, measured in cents and excluding any applicable taxes. It indicates the minimum amount that needs to be charged for each billing period. | [optional] 
**properties** | [**ChargeObjectProperties**](ChargeObjectProperties.md) |  | [optional] 
**group_properties** | [**List[PlanCreateInputPlanChargesInnerGroupPropertiesInner]**](PlanCreateInputPlanChargesInnerGroupPropertiesInner.md) | All charge information, sorted by groups. | [optional] 
**tax_codes** | **List[str]** | List of unique code used to identify the taxes. | [optional] 

## Example

```python
from lago_api.models.plan_overrides_object_charges_inner import PlanOverridesObjectChargesInner

# TODO update the JSON string below
json = "{}"
# create an instance of PlanOverridesObjectChargesInner from a JSON string
plan_overrides_object_charges_inner_instance = PlanOverridesObjectChargesInner.from_json(json)
# print the JSON string representation of the object
print PlanOverridesObjectChargesInner.to_json()

# convert the object into a dict
plan_overrides_object_charges_inner_dict = plan_overrides_object_charges_inner_instance.to_dict()
# create an instance of PlanOverridesObjectChargesInner from a dict
plan_overrides_object_charges_inner_form_dict = plan_overrides_object_charges_inner.from_dict(plan_overrides_object_charges_inner_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


