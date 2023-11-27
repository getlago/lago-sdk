# InvoiceUpdateInputInvoiceMetadataInner


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **str** | The metadata object identifier. Only required when updating existing metadata. | [optional] 
**key** | **str** | The key in the key-value pair of the metadata. | [optional] 
**value** | **str** | The value in the key-value pair of the metadata. | [optional] 

## Example

```python
from lago_api.models.invoice_update_input_invoice_metadata_inner import InvoiceUpdateInputInvoiceMetadataInner

# TODO update the JSON string below
json = "{}"
# create an instance of InvoiceUpdateInputInvoiceMetadataInner from a JSON string
invoice_update_input_invoice_metadata_inner_instance = InvoiceUpdateInputInvoiceMetadataInner.from_json(json)
# print the JSON string representation of the object
print InvoiceUpdateInputInvoiceMetadataInner.to_json()

# convert the object into a dict
invoice_update_input_invoice_metadata_inner_dict = invoice_update_input_invoice_metadata_inner_instance.to_dict()
# create an instance of InvoiceUpdateInputInvoiceMetadataInner from a dict
invoice_update_input_invoice_metadata_inner_form_dict = invoice_update_input_invoice_metadata_inner.from_dict(invoice_update_input_invoice_metadata_inner_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


