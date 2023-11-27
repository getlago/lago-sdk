# CreditNoteEstimated


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**estimated_credit_note** | [**CreditNoteEstimatedEstimatedCreditNote**](CreditNoteEstimatedEstimatedCreditNote.md) |  | 

## Example

```python
from lago_api.models.credit_note_estimated import CreditNoteEstimated

# TODO update the JSON string below
json = "{}"
# create an instance of CreditNoteEstimated from a JSON string
credit_note_estimated_instance = CreditNoteEstimated.from_json(json)
# print the JSON string representation of the object
print CreditNoteEstimated.to_json()

# convert the object into a dict
credit_note_estimated_dict = credit_note_estimated_instance.to_dict()
# create an instance of CreditNoteEstimated from a dict
credit_note_estimated_form_dict = credit_note_estimated.from_dict(credit_note_estimated_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


