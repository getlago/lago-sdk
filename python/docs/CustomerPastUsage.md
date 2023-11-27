# CustomerPastUsage


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**usage_periods** | [**List[CustomerUsage]**](CustomerUsage.md) |  | 
**meta** | [**PaginationMeta**](PaginationMeta.md) |  | 

## Example

```python
from lago_api.models.customer_past_usage import CustomerPastUsage

# TODO update the JSON string below
json = "{}"
# create an instance of CustomerPastUsage from a JSON string
customer_past_usage_instance = CustomerPastUsage.from_json(json)
# print the JSON string representation of the object
print CustomerPastUsage.to_json()

# convert the object into a dict
customer_past_usage_dict = customer_past_usage_instance.to_dict()
# create an instance of CustomerPastUsage from a dict
customer_past_usage_form_dict = customer_past_usage.from_dict(customer_past_usage_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


