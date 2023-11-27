# InvoiceOneOffCreateInputInvoiceFeesInner


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**add_on_code** | **str** | The code of the add-on used as invoice item. | 
**invoice_display_name** | **str** | Specifies the name that will be displayed on an invoice. If no value is set for this field, the name of the actual charge will be used as the default display name. | [optional] 
**unit_amount_cents** | **int** | The amount of the fee per unit, expressed in cents. By default, the amount of the add-on is used. | [optional] 
**units** | **str** | The quantity of units associated with the fee. By default, only 1 unit is added to the invoice. | [optional] 
**description** | **str** | This is a description | [optional] 
**tax_codes** | **List[str]** | List of unique code used to identify the taxes. | [optional] 

## Example

```python
from lago_api.models.invoice_one_off_create_input_invoice_fees_inner import InvoiceOneOffCreateInputInvoiceFeesInner

# TODO update the JSON string below
json = "{}"
# create an instance of InvoiceOneOffCreateInputInvoiceFeesInner from a JSON string
invoice_one_off_create_input_invoice_fees_inner_instance = InvoiceOneOffCreateInputInvoiceFeesInner.from_json(json)
# print the JSON string representation of the object
print InvoiceOneOffCreateInputInvoiceFeesInner.to_json()

# convert the object into a dict
invoice_one_off_create_input_invoice_fees_inner_dict = invoice_one_off_create_input_invoice_fees_inner_instance.to_dict()
# create an instance of InvoiceOneOffCreateInputInvoiceFeesInner from a dict
invoice_one_off_create_input_invoice_fees_inner_form_dict = invoice_one_off_create_input_invoice_fees_inner.from_dict(invoice_one_off_create_input_invoice_fees_inner_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


