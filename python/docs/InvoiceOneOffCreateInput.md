# InvoiceOneOffCreateInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**invoice** | [**InvoiceOneOffCreateInputInvoice**](InvoiceOneOffCreateInputInvoice.md) |  | 

## Example

```python
from lago_api.models.invoice_one_off_create_input import InvoiceOneOffCreateInput

# TODO update the JSON string below
json = "{}"
# create an instance of InvoiceOneOffCreateInput from a JSON string
invoice_one_off_create_input_instance = InvoiceOneOffCreateInput.from_json(json)
# print the JSON string representation of the object
print InvoiceOneOffCreateInput.to_json()

# convert the object into a dict
invoice_one_off_create_input_dict = invoice_one_off_create_input_instance.to_dict()
# create an instance of InvoiceOneOffCreateInput from a dict
invoice_one_off_create_input_form_dict = invoice_one_off_create_input.from_dict(invoice_one_off_create_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


