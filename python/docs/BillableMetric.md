# BillableMetric


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**billable_metric** | [**BillableMetricObject**](BillableMetricObject.md) |  | 

## Example

```python
from lago_api.models.billable_metric import BillableMetric

# TODO update the JSON string below
json = "{}"
# create an instance of BillableMetric from a JSON string
billable_metric_instance = BillableMetric.from_json(json)
# print the JSON string representation of the object
print BillableMetric.to_json()

# convert the object into a dict
billable_metric_dict = billable_metric_instance.to_dict()
# create an instance of BillableMetric from a dict
billable_metric_form_dict = billable_metric.from_dict(billable_metric_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


