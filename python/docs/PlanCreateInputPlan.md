# PlanCreateInputPlan


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | The name of the plan. | [optional] 
**invoice_display_name** | **str** | Specifies the name that will be displayed on an invoice. If no value is set for this field, the name of the plan will be used as the default display name. | [optional] 
**code** | **str** | The code of the plan. It serves as a unique identifier associated with a particular plan. The code is typically used for internal or system-level identification purposes, like assigning a subscription, for instance. | [optional] 
**interval** | **str** | The interval used for recurring billing. It represents the frequency at which subscription billing occurs. The interval can be one of the following values: &#x60;yearly&#x60;, &#x60;quarterly&#x60;, &#x60;monthly&#x60;, or &#x60;weekly&#x60;. | [optional] 
**description** | **str** | The description on the plan. | [optional] 
**amount_cents** | **int** | The base cost of the plan, excluding any applicable taxes, that is billed on a recurring basis. This value is defined at 0 if your plan is a pay-as-you-go plan. | [optional] 
**amount_currency** | [**Currency**](Currency.md) |  | [optional] 
**trial_period** | **float** | The duration in days during which the base cost of the plan is offered for free. | [optional] 
**pay_in_advance** | **bool** | This field determines the billing timing for the plan. When set to &#x60;true&#x60;, the base cost of the plan is due at the beginning of each billing period. Conversely, when set to &#x60;false&#x60;, the base cost of the plan is due at the end of each billing period. | [optional] 
**bill_charges_monthly** | **bool** | This field, when set to &#x60;true&#x60;, enables to invoice usage-based charges on monthly basis, even if the cadence of the plan is yearly. This allows customers to pay charges overage on a monthly basis. This can be set to true only if the plan’s interval is &#x60;yearly&#x60;. | [optional] 
**tax_codes** | **List[str]** | List of unique code used to identify the taxes. | [optional] 
**charges** | [**List[PlanCreateInputPlanChargesInner]**](PlanCreateInputPlanChargesInner.md) | Additional usage-based charges for this plan. | [optional] 

## Example

```python
from lago_api.models.plan_create_input_plan import PlanCreateInputPlan

# TODO update the JSON string below
json = "{}"
# create an instance of PlanCreateInputPlan from a JSON string
plan_create_input_plan_instance = PlanCreateInputPlan.from_json(json)
# print the JSON string representation of the object
print PlanCreateInputPlan.to_json()

# convert the object into a dict
plan_create_input_plan_dict = plan_create_input_plan_instance.to_dict()
# create an instance of PlanCreateInputPlan from a dict
plan_create_input_plan_form_dict = plan_create_input_plan.from_dict(plan_create_input_plan_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


