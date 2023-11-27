# InvoiceCollectionObject


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**month** | **str** | Identifies the month to analyze revenue. | 
**payment_status** | **str** | The payment status of the invoices. | [optional] 
**invoices_count** | **int** | Contains invoices count. | 
**amount_cents** | **int** | The total amount of revenue for a period, expressed in cents. | 
**currency** | [**Currency**](Currency.md) |  | [optional] 

## Example

```python
from lago_api.models.invoice_collection_object import InvoiceCollectionObject

# TODO update the JSON string below
json = "{}"
# create an instance of InvoiceCollectionObject from a JSON string
invoice_collection_object_instance = InvoiceCollectionObject.from_json(json)
# print the JSON string representation of the object
print InvoiceCollectionObject.to_json()

# convert the object into a dict
invoice_collection_object_dict = invoice_collection_object_instance.to_dict()
# create an instance of InvoiceCollectionObject from a dict
invoice_collection_object_form_dict = invoice_collection_object.from_dict(invoice_collection_object_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


