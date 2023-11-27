# CustomerUsage


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**customer_usage** | [**CustomerUsageObject**](CustomerUsageObject.md) |  | 

## Example

```python
from lago_api.models.customer_usage import CustomerUsage

# TODO update the JSON string below
json = "{}"
# create an instance of CustomerUsage from a JSON string
customer_usage_instance = CustomerUsage.from_json(json)
# print the JSON string representation of the object
print CustomerUsage.to_json()

# convert the object into a dict
customer_usage_dict = customer_usage_instance.to_dict()
# create an instance of CustomerUsage from a dict
customer_usage_form_dict = customer_usage.from_dict(customer_usage_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


