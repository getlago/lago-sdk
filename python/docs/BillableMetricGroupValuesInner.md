# BillableMetricGroupValuesInner


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**key** | **str** | Name of the event property used to group values. | 
**values** | **List[str]** | Array of strings representing all possible values. | 

## Example

```python
from lago_api.models.billable_metric_group_values_inner import BillableMetricGroupValuesInner

# TODO update the JSON string below
json = "{}"
# create an instance of BillableMetricGroupValuesInner from a JSON string
billable_metric_group_values_inner_instance = BillableMetricGroupValuesInner.from_json(json)
# print the JSON string representation of the object
print BillableMetricGroupValuesInner.to_json()

# convert the object into a dict
billable_metric_group_values_inner_dict = billable_metric_group_values_inner_instance.to_dict()
# create an instance of BillableMetricGroupValuesInner from a dict
billable_metric_group_values_inner_form_dict = billable_metric_group_values_inner.from_dict(billable_metric_group_values_inner_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


