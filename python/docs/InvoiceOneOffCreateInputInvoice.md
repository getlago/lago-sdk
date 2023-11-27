# InvoiceOneOffCreateInputInvoice


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**external_customer_id** | **str** | Unique identifier assigned to the customer in your application. | 
**currency** | [**Currency**](Currency.md) |  | [optional] 
**fees** | [**List[InvoiceOneOffCreateInputInvoiceFeesInner]**](InvoiceOneOffCreateInputInvoiceFeesInner.md) |  | 

## Example

```python
from lago_api.models.invoice_one_off_create_input_invoice import InvoiceOneOffCreateInputInvoice

# TODO update the JSON string below
json = "{}"
# create an instance of InvoiceOneOffCreateInputInvoice from a JSON string
invoice_one_off_create_input_invoice_instance = InvoiceOneOffCreateInputInvoice.from_json(json)
# print the JSON string representation of the object
print InvoiceOneOffCreateInputInvoice.to_json()

# convert the object into a dict
invoice_one_off_create_input_invoice_dict = invoice_one_off_create_input_invoice_instance.to_dict()
# create an instance of InvoiceOneOffCreateInputInvoice from a dict
invoice_one_off_create_input_invoice_form_dict = invoice_one_off_create_input_invoice.from_dict(invoice_one_off_create_input_invoice_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


