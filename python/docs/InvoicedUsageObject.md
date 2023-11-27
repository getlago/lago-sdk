# InvoicedUsageObject


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**month** | **str** | Identifies the month to analyze revenue. | 
**code** | **str** | The code of the usage-based billable metrics. | [optional] 
**amount_cents** | **int** | The total amount of revenue for a period, expressed in cents. | 
**currency** | [**Currency**](Currency.md) |  | 

## Example

```python
from lago_api.models.invoiced_usage_object import InvoicedUsageObject

# TODO update the JSON string below
json = "{}"
# create an instance of InvoicedUsageObject from a JSON string
invoiced_usage_object_instance = InvoicedUsageObject.from_json(json)
# print the JSON string representation of the object
print InvoicedUsageObject.to_json()

# convert the object into a dict
invoiced_usage_object_dict = invoiced_usage_object_instance.to_dict()
# create an instance of InvoicedUsageObject from a dict
invoiced_usage_object_form_dict = invoiced_usage_object.from_dict(invoiced_usage_object_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


