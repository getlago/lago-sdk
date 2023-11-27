# BillableMetricBaseInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | Name of the billable metric. | [optional] 
**code** | **str** | Unique code used to identify the billable metric associated with the API request. This code associates each event with the correct metric. | [optional] 
**description** | **str** | Internal description of the billable metric. | [optional] 
**recurring** | **bool** | Defines if the billable metric is persisted billing period over billing period.  - If set to &#x60;true&#x60;: the accumulated number of units calculated from the previous billing period is persisted to the next billing period. - If set to &#x60;false&#x60;: the accumulated number of units is reset to 0 at the end of the billing period. - If not defined in the request, default value is &#x60;false&#x60;. | [optional] 
**field_name** | **str** | Property of the billable metric used for aggregating usage data. This field is not required for &#x60;count_agg&#x60;. | [optional] 
**aggregation_type** | **str** | Aggregation method used to compute usage for this billable metric. | [optional] 
**weighted_interval** | **str** | Parameter exclusively utilized in conjunction with the &#x60;weighted_sum&#x60; aggregation type. It serves to adjust the aggregation result by assigning weights and proration to the result based on time intervals. When this field is not provided, the default time interval is assumed to be in &#x60;seconds&#x60;. | [optional] 
**group** | [**BillableMetricGroup**](BillableMetricGroup.md) |  | [optional] 

## Example

```python
from lago_api.models.billable_metric_base_input import BillableMetricBaseInput

# TODO update the JSON string below
json = "{}"
# create an instance of BillableMetricBaseInput from a JSON string
billable_metric_base_input_instance = BillableMetricBaseInput.from_json(json)
# print the JSON string representation of the object
print BillableMetricBaseInput.to_json()

# convert the object into a dict
billable_metric_base_input_dict = billable_metric_base_input_instance.to_dict()
# create an instance of BillableMetricBaseInput from a dict
billable_metric_base_input_form_dict = billable_metric_base_input.from_dict(billable_metric_base_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


