# CreditNoteAppliedTaxObject


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**lago_credit_note_id** | **str** | Unique identifier of the credit note, created by Lago. | [optional] 
**base_amount_cents** | **int** |  | [optional] 
**lago_id** | **str** | Unique identifier of the applied tax, created by Lago. | [optional] 
**lago_tax_id** | **str** | Unique identifier of the tax, created by Lago. | [optional] 
**tax_name** | **str** | Name of the tax. | [optional] 
**tax_code** | **str** | Unique code used to identify the tax associated with the API request. | [optional] 
**tax_rate** | **float** | The percentage rate of the tax | [optional] 
**tax_description** | **str** | Internal description of the taxe | [optional] 
**amount_cents** | **int** | Amount of the tax | [optional] 
**amount_currency** | [**Currency**](Currency.md) |  | [optional] 
**created_at** | **datetime** | The date and time when the applied tax was created. It is expressed in UTC format according to the ISO 8601 datetime standard. This field provides the timestamp for the exact moment when the applied tax was initially created. | [optional] 

## Example

```python
from lago_api.models.credit_note_applied_tax_object import CreditNoteAppliedTaxObject

# TODO update the JSON string below
json = "{}"
# create an instance of CreditNoteAppliedTaxObject from a JSON string
credit_note_applied_tax_object_instance = CreditNoteAppliedTaxObject.from_json(json)
# print the JSON string representation of the object
print CreditNoteAppliedTaxObject.to_json()

# convert the object into a dict
credit_note_applied_tax_object_dict = credit_note_applied_tax_object_instance.to_dict()
# create an instance of CreditNoteAppliedTaxObject from a dict
credit_note_applied_tax_object_form_dict = credit_note_applied_tax_object.from_dict(credit_note_applied_tax_object_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


