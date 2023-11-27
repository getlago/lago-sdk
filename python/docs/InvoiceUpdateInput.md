# InvoiceUpdateInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**invoice** | [**InvoiceUpdateInputInvoice**](InvoiceUpdateInputInvoice.md) |  | 

## Example

```python
from lago_api.models.invoice_update_input import InvoiceUpdateInput

# TODO update the JSON string below
json = "{}"
# create an instance of InvoiceUpdateInput from a JSON string
invoice_update_input_instance = InvoiceUpdateInput.from_json(json)
# print the JSON string representation of the object
print InvoiceUpdateInput.to_json()

# convert the object into a dict
invoice_update_input_dict = invoice_update_input_instance.to_dict()
# create an instance of InvoiceUpdateInput from a dict
invoice_update_input_form_dict = invoice_update_input.from_dict(invoice_update_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


