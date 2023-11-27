# InvoicedUsages


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**invoiced_usages** | [**List[InvoicedUsageObject]**](InvoicedUsageObject.md) |  | 

## Example

```python
from lago_api.models.invoiced_usages import InvoicedUsages

# TODO update the JSON string below
json = "{}"
# create an instance of InvoicedUsages from a JSON string
invoiced_usages_instance = InvoicedUsages.from_json(json)
# print the JSON string representation of the object
print InvoicedUsages.to_json()

# convert the object into a dict
invoiced_usages_dict = invoiced_usages_instance.to_dict()
# create an instance of InvoicedUsages from a dict
invoiced_usages_form_dict = invoiced_usages.from_dict(invoiced_usages_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


