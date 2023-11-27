# CreditNoteUpdateInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**credit_note** | [**CreditNoteUpdateInputCreditNote**](CreditNoteUpdateInputCreditNote.md) |  | 

## Example

```python
from lago_api.models.credit_note_update_input import CreditNoteUpdateInput

# TODO update the JSON string below
json = "{}"
# create an instance of CreditNoteUpdateInput from a JSON string
credit_note_update_input_instance = CreditNoteUpdateInput.from_json(json)
# print the JSON string representation of the object
print CreditNoteUpdateInput.to_json()

# convert the object into a dict
credit_note_update_input_dict = credit_note_update_input_instance.to_dict()
# create an instance of CreditNoteUpdateInput from a dict
credit_note_update_input_form_dict = credit_note_update_input.from_dict(credit_note_update_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


