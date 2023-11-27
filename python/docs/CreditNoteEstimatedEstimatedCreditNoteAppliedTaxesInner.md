# CreditNoteEstimatedEstimatedCreditNoteAppliedTaxesInner


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**lago_tax_id** | **str** | Unique identifier of the tax, created by Lago. | [optional] 
**tax_name** | **str** | Name of the tax. | [optional] 
**tax_code** | **str** | Unique code used to identify the tax associated with the API request. | [optional] 
**tax_rate** | **float** | The percentage rate of the tax | [optional] 
**tax_description** | **str** | Internal description of the taxe | [optional] 
**base_amount_cents** | **int** |  | [optional] 
**amount_cents** | **int** | Amount of the tax | [optional] 
**amount_currency** | [**Currency**](Currency.md) |  | [optional] 

## Example

```python
from lago_api.models.credit_note_estimated_estimated_credit_note_applied_taxes_inner import CreditNoteEstimatedEstimatedCreditNoteAppliedTaxesInner

# TODO update the JSON string below
json = "{}"
# create an instance of CreditNoteEstimatedEstimatedCreditNoteAppliedTaxesInner from a JSON string
credit_note_estimated_estimated_credit_note_applied_taxes_inner_instance = CreditNoteEstimatedEstimatedCreditNoteAppliedTaxesInner.from_json(json)
# print the JSON string representation of the object
print CreditNoteEstimatedEstimatedCreditNoteAppliedTaxesInner.to_json()

# convert the object into a dict
credit_note_estimated_estimated_credit_note_applied_taxes_inner_dict = credit_note_estimated_estimated_credit_note_applied_taxes_inner_instance.to_dict()
# create an instance of CreditNoteEstimatedEstimatedCreditNoteAppliedTaxesInner from a dict
credit_note_estimated_estimated_credit_note_applied_taxes_inner_form_dict = credit_note_estimated_estimated_credit_note_applied_taxes_inner.from_dict(credit_note_estimated_estimated_credit_note_applied_taxes_inner_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


