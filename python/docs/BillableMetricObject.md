# BillableMetricObject


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**lago_id** | **str** | Unique identifier of the billable metric created by Lago. | 
**name** | **str** | Name of the billable metric. | 
**code** | **str** | Unique code used to identify the billable metric associated with the API request. This code associates each event with the correct metric. | 
**description** | **str** | Internal description of the billable metric. | [optional] 
**recurring** | **bool** | Defines if the billable metric is persisted billing period over billing period.  - If set to &#x60;true&#x60;: the accumulated number of units calculated from the previous billing period is persisted to the next billing period. - If set to &#x60;false&#x60;: the accumulated number of units is reset to 0 at the end of the billing period. - If not defined in the request, default value is &#x60;false&#x60;. | 
**created_at** | **datetime** | Creation date of the billable metric. | 
**field_name** | **str** | Property of the billable metric used for aggregating usage data. This field is not required for &#x60;count_agg&#x60;. | [optional] 
**aggregation_type** | **str** | Aggregation method used to compute usage for this billable metric. | 
**weighted_interval** | **str** | Parameter exclusively utilized in conjunction with the &#x60;weighted_sum&#x60; aggregation type. It serves to adjust the aggregation result by assigning weights and proration to the result based on time intervals. When this field is not provided, the default time interval is assumed to be in &#x60;seconds&#x60;. | [optional] 
**group** | [**BillableMetricGroup**](BillableMetricGroup.md) |  | [optional] 
**active_subscriptions_count** | **int** | Number of active subscriptions using this billable metric. | 
**draft_invoices_count** | **int** | Number of draft invoices for which this billable metric is listed as an invoice item. | 
**plans_count** | **int** | Number of plans using this billable metric. | 

## Example

```python
from lago_api.models.billable_metric_object import BillableMetricObject

# TODO update the JSON string below
json = "{}"
# create an instance of BillableMetricObject from a JSON string
billable_metric_object_instance = BillableMetricObject.from_json(json)
# print the JSON string representation of the object
print BillableMetricObject.to_json()

# convert the object into a dict
billable_metric_object_dict = billable_metric_object_instance.to_dict()
# create an instance of BillableMetricObject from a dict
billable_metric_object_form_dict = billable_metric_object.from_dict(billable_metric_object_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


