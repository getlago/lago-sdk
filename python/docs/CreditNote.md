# CreditNote


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**credit_note** | [**CreditNoteObject**](CreditNoteObject.md) |  | 

## Example

```python
from lago_api.models.credit_note import CreditNote

# TODO update the JSON string below
json = "{}"
# create an instance of CreditNote from a JSON string
credit_note_instance = CreditNote.from_json(json)
# print the JSON string representation of the object
print CreditNote.to_json()

# convert the object into a dict
credit_note_dict = credit_note_instance.to_dict()
# create an instance of CreditNote from a dict
credit_note_form_dict = credit_note.from_dict(credit_note_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


