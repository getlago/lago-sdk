# CustomerChargeUsageObjectGroupsInner


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**lago_id** | **str** | Unique identifier assigned to the group within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the group record within the Lago system. | [optional] 
**key** | **str** | The group key, only returned for groups with two dimensions. | [optional] 
**value** | **str** | The group value. | [optional] 
**units** | **str** | The number of units consumed for a specific group related to a charge item. | [optional] 
**events_count** | **int** | The quantity of usage events that have been recorded for a particular charge during the specified time period. These events may also be referred to as the number of transactions in some contexts. | [optional] 
**amount_cents** | **int** | The amount in cents, tax excluded, consumed for a specific group related to a charge item. | [optional] 

## Example

```python
from lago_api.models.customer_charge_usage_object_groups_inner import CustomerChargeUsageObjectGroupsInner

# TODO update the JSON string below
json = "{}"
# create an instance of CustomerChargeUsageObjectGroupsInner from a JSON string
customer_charge_usage_object_groups_inner_instance = CustomerChargeUsageObjectGroupsInner.from_json(json)
# print the JSON string representation of the object
print CustomerChargeUsageObjectGroupsInner.to_json()

# convert the object into a dict
customer_charge_usage_object_groups_inner_dict = customer_charge_usage_object_groups_inner_instance.to_dict()
# create an instance of CustomerChargeUsageObjectGroupsInner from a dict
customer_charge_usage_object_groups_inner_form_dict = customer_charge_usage_object_groups_inner.from_dict(customer_charge_usage_object_groups_inner_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


