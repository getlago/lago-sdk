# InvoiceCollections


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**invoice_collections** | [**List[InvoiceCollectionObject]**](InvoiceCollectionObject.md) |  | 

## Example

```python
from lago_api.models.invoice_collections import InvoiceCollections

# TODO update the JSON string below
json = "{}"
# create an instance of InvoiceCollections from a JSON string
invoice_collections_instance = InvoiceCollections.from_json(json)
# print the JSON string representation of the object
print InvoiceCollections.to_json()

# convert the object into a dict
invoice_collections_dict = invoice_collections_instance.to_dict()
# create an instance of InvoiceCollections from a dict
invoice_collections_form_dict = invoice_collections.from_dict(invoice_collections_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


