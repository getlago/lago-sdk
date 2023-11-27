# CreditNoteEstimatedEstimatedCreditNoteItemsInner


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**amount_cents** | **int** | The credit note’s item amount, expressed in cents. | 
**lago_fee_id** | **str** | Unique identifier assigned to the fee within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the fee’s record within the Lago system. | 

## Example

```python
from lago_api.models.credit_note_estimated_estimated_credit_note_items_inner import CreditNoteEstimatedEstimatedCreditNoteItemsInner

# TODO update the JSON string below
json = "{}"
# create an instance of CreditNoteEstimatedEstimatedCreditNoteItemsInner from a JSON string
credit_note_estimated_estimated_credit_note_items_inner_instance = CreditNoteEstimatedEstimatedCreditNoteItemsInner.from_json(json)
# print the JSON string representation of the object
print CreditNoteEstimatedEstimatedCreditNoteItemsInner.to_json()

# convert the object into a dict
credit_note_estimated_estimated_credit_note_items_inner_dict = credit_note_estimated_estimated_credit_note_items_inner_instance.to_dict()
# create an instance of CreditNoteEstimatedEstimatedCreditNoteItemsInner from a dict
credit_note_estimated_estimated_credit_note_items_inner_form_dict = credit_note_estimated_estimated_credit_note_items_inner.from_dict(credit_note_estimated_estimated_credit_note_items_inner_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


