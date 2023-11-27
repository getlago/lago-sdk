# CreditNoteUpdateInputCreditNote


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**refund_status** | **str** | The status of the refund portion of the credit note. It indicates the current state or condition of the refund associated with the credit note. The possible values for this field are:  - &#x60;pending&#x60;: this status indicates that the refund is pending execution. The refund request has been initiated but has not been processed or completed yet. - &#x60;succeeded&#x60;: this status indicates that the refund has been successfully executed. The refund amount has been processed and returned to the customer or the designated recipient. - &#x60;failed&#x60;: this status indicates that the refund failed to execute. The refund request encountered an error or unsuccessful processing, and the refund amount could not be returned. | 

## Example

```python
from lago_api.models.credit_note_update_input_credit_note import CreditNoteUpdateInputCreditNote

# TODO update the JSON string below
json = "{}"
# create an instance of CreditNoteUpdateInputCreditNote from a JSON string
credit_note_update_input_credit_note_instance = CreditNoteUpdateInputCreditNote.from_json(json)
# print the JSON string representation of the object
print CreditNoteUpdateInputCreditNote.to_json()

# convert the object into a dict
credit_note_update_input_credit_note_dict = credit_note_update_input_credit_note_instance.to_dict()
# create an instance of CreditNoteUpdateInputCreditNote from a dict
credit_note_update_input_credit_note_form_dict = credit_note_update_input_credit_note.from_dict(credit_note_update_input_credit_note_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


