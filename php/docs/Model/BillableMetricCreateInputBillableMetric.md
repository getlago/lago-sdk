# # BillableMetricCreateInputBillableMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **string** | Name of the billable metric. |
**code** | **string** | Unique code used to identify the billable metric associated with the API request. This code associates each event with the correct metric. |
**description** | **string** | Internal description of the billable metric. | [optional]
**recurring** | **bool** | Defines if the billable metric is persisted billing period over billing period.  - If set to &#x60;true&#x60;: the accumulated number of units calculated from the previous billing period is persisted to the next billing period. - If set to &#x60;false&#x60;: the accumulated number of units is reset to 0 at the end of the billing period. - If not defined in the request, default value is &#x60;false&#x60;. | [optional]
**field_name** | **string** | Property of the billable metric used for aggregating usage data. This field is not required for &#x60;count_agg&#x60;. | [optional]
**aggregation_type** | **string** | Aggregation method used to compute usage for this billable metric. |
**weighted_interval** | **string** | Parameter exclusively utilized in conjunction with the &#x60;weighted_sum&#x60; aggregation type. It serves to adjust the aggregation result by assigning weights and proration to the result based on time intervals. When this field is not provided, the default time interval is assumed to be in &#x60;seconds&#x60;. | [optional]
**group** | [**\OpenAPI\Client\Model\BillableMetricGroup**](BillableMetricGroup.md) |  | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
