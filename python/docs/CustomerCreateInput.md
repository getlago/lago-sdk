# CustomerCreateInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**customer** | [**CustomerCreateInputCustomer**](CustomerCreateInputCustomer.md) |  | 

## Example

```python
from lago_api.models.customer_create_input import CustomerCreateInput

# TODO update the JSON string below
json = "{}"
# create an instance of CustomerCreateInput from a JSON string
customer_create_input_instance = CustomerCreateInput.from_json(json)
# print the JSON string representation of the object
print CustomerCreateInput.to_json()

# convert the object into a dict
customer_create_input_dict = customer_create_input_instance.to_dict()
# create an instance of CustomerCreateInput from a dict
customer_create_input_form_dict = customer_create_input.from_dict(customer_create_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


