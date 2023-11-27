# CreditNoteCreateInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**credit_note** | [**CreditNoteCreateInputCreditNote**](CreditNoteCreateInputCreditNote.md) |  | 

## Example

```python
from lago_api.models.credit_note_create_input import CreditNoteCreateInput

# TODO update the JSON string below
json = "{}"
# create an instance of CreditNoteCreateInput from a JSON string
credit_note_create_input_instance = CreditNoteCreateInput.from_json(json)
# print the JSON string representation of the object
print CreditNoteCreateInput.to_json()

# convert the object into a dict
credit_note_create_input_dict = credit_note_create_input_instance.to_dict()
# create an instance of CreditNoteCreateInput from a dict
credit_note_create_input_form_dict = credit_note_create_input.from_dict(credit_note_create_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


