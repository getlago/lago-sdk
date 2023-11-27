# CustomerChargeUsageObjectBillableMetric

The related billable metric object.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**lago_id** | **str** | Unique identifier assigned to the billable metric within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the billable metric’s record within the Lago system. | 
**name** | **str** | The name of the billable metric used for this charge. | 
**code** | **str** | The code of the billable metric used for this charge. | 
**aggregation_type** | **str** | The aggregation type of the billable metric used for this charge. Possible values are &#x60;count_agg&#x60;, &#x60;sum_agg&#x60;, &#x60;max_agg&#x60; or &#x60;unique_count_agg&#x60;. | 

## Example

```python
from lago_api.models.customer_charge_usage_object_billable_metric import CustomerChargeUsageObjectBillableMetric

# TODO update the JSON string below
json = "{}"
# create an instance of CustomerChargeUsageObjectBillableMetric from a JSON string
customer_charge_usage_object_billable_metric_instance = CustomerChargeUsageObjectBillableMetric.from_json(json)
# print the JSON string representation of the object
print CustomerChargeUsageObjectBillableMetric.to_json()

# convert the object into a dict
customer_charge_usage_object_billable_metric_dict = customer_charge_usage_object_billable_metric_instance.to_dict()
# create an instance of CustomerChargeUsageObjectBillableMetric from a dict
customer_charge_usage_object_billable_metric_form_dict = customer_charge_usage_object_billable_metric.from_dict(customer_charge_usage_object_billable_metric_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


