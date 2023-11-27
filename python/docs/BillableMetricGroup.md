# BillableMetricGroup

Group with one or two dimensions, used to apply differentiated pricing based on additional properties of the billable metric.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**key** | **str** | Name of the event property used to group values. | 
**values** | [**List[BillableMetricGroupValuesInner]**](BillableMetricGroupValuesInner.md) | Array of strings or objects representing all possible values. | 

## Example

```python
from lago_api.models.billable_metric_group import BillableMetricGroup

# TODO update the JSON string below
json = "{}"
# create an instance of BillableMetricGroup from a JSON string
billable_metric_group_instance = BillableMetricGroup.from_json(json)
# print the JSON string representation of the object
print BillableMetricGroup.to_json()

# convert the object into a dict
billable_metric_group_dict = billable_metric_group_instance.to_dict()
# create an instance of BillableMetricGroup from a dict
billable_metric_group_form_dict = billable_metric_group.from_dict(billable_metric_group_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


