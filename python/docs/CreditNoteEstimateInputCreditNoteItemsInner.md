# CreditNoteEstimateInputCreditNoteItemsInner


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**fee_id** | **str** | The fee unique identifier, created by Lago. | 
**amount_cents** | **int** | The amount of the credit note item, expressed in cents. | 

## Example

```python
from lago_api.models.credit_note_estimate_input_credit_note_items_inner import CreditNoteEstimateInputCreditNoteItemsInner

# TODO update the JSON string below
json = "{}"
# create an instance of CreditNoteEstimateInputCreditNoteItemsInner from a JSON string
credit_note_estimate_input_credit_note_items_inner_instance = CreditNoteEstimateInputCreditNoteItemsInner.from_json(json)
# print the JSON string representation of the object
print CreditNoteEstimateInputCreditNoteItemsInner.to_json()

# convert the object into a dict
credit_note_estimate_input_credit_note_items_inner_dict = credit_note_estimate_input_credit_note_items_inner_instance.to_dict()
# create an instance of CreditNoteEstimateInputCreditNoteItemsInner from a dict
credit_note_estimate_input_credit_note_items_inner_form_dict = credit_note_estimate_input_credit_note_items_inner.from_dict(credit_note_estimate_input_credit_note_items_inner_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


