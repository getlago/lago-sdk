# CreditNoteEstimateInputCreditNote


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**invoice_id** | **str** | The invoice unique identifier, created by Lago. | 
**items** | [**List[CreditNoteEstimateInputCreditNoteItemsInner]**](CreditNoteEstimateInputCreditNoteItemsInner.md) | The list of credit note’s items. | 

## Example

```python
from lago_api.models.credit_note_estimate_input_credit_note import CreditNoteEstimateInputCreditNote

# TODO update the JSON string below
json = "{}"
# create an instance of CreditNoteEstimateInputCreditNote from a JSON string
credit_note_estimate_input_credit_note_instance = CreditNoteEstimateInputCreditNote.from_json(json)
# print the JSON string representation of the object
print CreditNoteEstimateInputCreditNote.to_json()

# convert the object into a dict
credit_note_estimate_input_credit_note_dict = credit_note_estimate_input_credit_note_instance.to_dict()
# create an instance of CreditNoteEstimateInputCreditNote from a dict
credit_note_estimate_input_credit_note_form_dict = credit_note_estimate_input_credit_note.from_dict(credit_note_estimate_input_credit_note_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


