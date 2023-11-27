# CustomerMetadata

Set of key-value pairs that you can attach to a customer. This can be useful for storing additional information about the customer in a structured format

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**lago_id** | **str** | A unique identifier for the customer metadata object in the Lago application. Can be used to update a key-value pair | 
**key** | **str** | The metadata object key | 
**value** | **str** | The metadata object value | 
**display_in_invoice** | **bool** | Determines whether the item or information should be displayed in the invoice. If set to true, the item or information will be included and visible in the generated invoice. If set to false, the item or information will be excluded and not displayed in the invoice. | 
**created_at** | **datetime** | The date of the metadata object creation, represented in ISO 8601 datetime format and expressed in Coordinated Universal Time (UTC). The creation_date provides a standardized and internationally recognized timestamp for when the metadata object was created | 

## Example

```python
from lago_api.models.customer_metadata import CustomerMetadata

# TODO update the JSON string below
json = "{}"
# create an instance of CustomerMetadata from a JSON string
customer_metadata_instance = CustomerMetadata.from_json(json)
# print the JSON string representation of the object
print CustomerMetadata.to_json()

# convert the object into a dict
customer_metadata_dict = customer_metadata_instance.to_dict()
# create an instance of CustomerMetadata from a dict
customer_metadata_form_dict = customer_metadata.from_dict(customer_metadata_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


