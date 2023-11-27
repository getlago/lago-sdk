# InvoiceUpdateInputInvoice


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**payment_status** | **str** | The payment status of the invoice. Possible values are &#x60;pending&#x60;, &#x60;failed&#x60; or &#x60;succeeded&#x60;. | [optional] 
**metadata** | [**List[InvoiceUpdateInputInvoiceMetadataInner]**](InvoiceUpdateInputInvoiceMetadataInner.md) |  | [optional] 

## Example

```python
from lago_api.models.invoice_update_input_invoice import InvoiceUpdateInputInvoice

# TODO update the JSON string below
json = "{}"
# create an instance of InvoiceUpdateInputInvoice from a JSON string
invoice_update_input_invoice_instance = InvoiceUpdateInputInvoice.from_json(json)
# print the JSON string representation of the object
print InvoiceUpdateInputInvoice.to_json()

# convert the object into a dict
invoice_update_input_invoice_dict = invoice_update_input_invoice_instance.to_dict()
# create an instance of InvoiceUpdateInputInvoice from a dict
invoice_update_input_invoice_form_dict = invoice_update_input_invoice.from_dict(invoice_update_input_invoice_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


