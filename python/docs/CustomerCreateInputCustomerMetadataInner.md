# CustomerCreateInputCustomerMetadataInner


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **str** | Identifier for the metadata object, only required when updating a key-value pair | [optional] 
**key** | **str** | The metadata object key | 
**value** | **str** | The metadata object value | 
**display_in_invoice** | **bool** | Determines whether the item or information should be displayed in the invoice. If set to true, the item or information will be included and visible in the generated invoice. If set to false, the item or information will be excluded and not displayed in the invoice. | 

## Example

```python
from lago_api.models.customer_create_input_customer_metadata_inner import CustomerCreateInputCustomerMetadataInner

# TODO update the JSON string below
json = "{}"
# create an instance of CustomerCreateInputCustomerMetadataInner from a JSON string
customer_create_input_customer_metadata_inner_instance = CustomerCreateInputCustomerMetadataInner.from_json(json)
# print the JSON string representation of the object
print CustomerCreateInputCustomerMetadataInner.to_json()

# convert the object into a dict
customer_create_input_customer_metadata_inner_dict = customer_create_input_customer_metadata_inner_instance.to_dict()
# create an instance of CustomerCreateInputCustomerMetadataInner from a dict
customer_create_input_customer_metadata_inner_form_dict = customer_create_input_customer_metadata_inner.from_dict(customer_create_input_customer_metadata_inner_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


