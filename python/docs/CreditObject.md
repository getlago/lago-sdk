# CreditObject


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**lago_id** | **str** | Unique identifier assigned to the credit within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the credit’s item record within the Lago system. | 
**amount_cents** | **int** | The amount of credit associated with the invoice, expressed in cents. | 
**amount_currency** | [**Currency**](Currency.md) |  | 
**before_taxes** | **bool** | Indicates whether the credit is applied on the amount before taxes (coupons) or after taxes (credit notes). This flag helps determine the order in which credits are applied to the invoice calculation | 
**item** | [**CreditObjectItem**](CreditObjectItem.md) |  | 
**invoice** | [**CreditObjectInvoice**](CreditObjectInvoice.md) |  | 

## Example

```python
from lago_api.models.credit_object import CreditObject

# TODO update the JSON string below
json = "{}"
# create an instance of CreditObject from a JSON string
credit_object_instance = CreditObject.from_json(json)
# print the JSON string representation of the object
print CreditObject.to_json()

# convert the object into a dict
credit_object_dict = credit_object_instance.to_dict()
# create an instance of CreditObject from a dict
credit_object_form_dict = credit_object.from_dict(credit_object_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


