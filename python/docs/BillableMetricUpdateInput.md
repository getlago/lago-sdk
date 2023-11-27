# BillableMetricUpdateInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**billable_metric** | [**BillableMetricBaseInput**](BillableMetricBaseInput.md) |  | 

## Example

```python
from lago_api.models.billable_metric_update_input import BillableMetricUpdateInput

# TODO update the JSON string below
json = "{}"
# create an instance of BillableMetricUpdateInput from a JSON string
billable_metric_update_input_instance = BillableMetricUpdateInput.from_json(json)
# print the JSON string representation of the object
print BillableMetricUpdateInput.to_json()

# convert the object into a dict
billable_metric_update_input_dict = billable_metric_update_input_instance.to_dict()
# create an instance of BillableMetricUpdateInput from a dict
billable_metric_update_input_form_dict = billable_metric_update_input.from_dict(billable_metric_update_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


