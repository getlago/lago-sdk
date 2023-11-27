# BillableMetricCreateInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**billable_metric** | [**BillableMetricCreateInputBillableMetric**](BillableMetricCreateInputBillableMetric.md) |  | 

## Example

```python
from lago_api.models.billable_metric_create_input import BillableMetricCreateInput

# TODO update the JSON string below
json = "{}"
# create an instance of BillableMetricCreateInput from a JSON string
billable_metric_create_input_instance = BillableMetricCreateInput.from_json(json)
# print the JSON string representation of the object
print BillableMetricCreateInput.to_json()

# convert the object into a dict
billable_metric_create_input_dict = billable_metric_create_input_instance.to_dict()
# create an instance of BillableMetricCreateInput from a dict
billable_metric_create_input_form_dict = billable_metric_create_input.from_dict(billable_metric_create_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


