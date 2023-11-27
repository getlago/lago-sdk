# InvoiceMetadataObject


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**lago_id** | **str** | Unique identifier assigned to the invoice metadata within the Lago application. | [optional] 
**key** | **str** | Represents the key of the metadata’s key-value pair. | [optional] 
**value** | **str** | Represents the value of the metadata’s key-value pair. | [optional] 
**created_at** | **datetime** | The date and time when the metadata object was created. It follows the ISO 8601 datetime format and is expressed in Coordinated Universal Time (UTC). | [optional] 

## Example

```python
from lago_api.models.invoice_metadata_object import InvoiceMetadataObject

# TODO update the JSON string below
json = "{}"
# create an instance of InvoiceMetadataObject from a JSON string
invoice_metadata_object_instance = InvoiceMetadataObject.from_json(json)
# print the JSON string representation of the object
print InvoiceMetadataObject.to_json()

# convert the object into a dict
invoice_metadata_object_dict = invoice_metadata_object_instance.to_dict()
# create an instance of InvoiceMetadataObject from a dict
invoice_metadata_object_form_dict = invoice_metadata_object.from_dict(invoice_metadata_object_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


