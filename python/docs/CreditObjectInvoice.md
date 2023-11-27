# CreditObjectInvoice


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**lago_id** | **str** |  | 
**payment_status** | **str** |  | 

## Example

```python
from lago_api.models.credit_object_invoice import CreditObjectInvoice

# TODO update the JSON string below
json = "{}"
# create an instance of CreditObjectInvoice from a JSON string
credit_object_invoice_instance = CreditObjectInvoice.from_json(json)
# print the JSON string representation of the object
print CreditObjectInvoice.to_json()

# convert the object into a dict
credit_object_invoice_dict = credit_object_invoice_instance.to_dict()
# create an instance of CreditObjectInvoice from a dict
credit_object_invoice_form_dict = credit_object_invoice.from_dict(credit_object_invoice_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


