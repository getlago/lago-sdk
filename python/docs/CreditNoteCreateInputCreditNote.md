# CreditNoteCreateInputCreditNote


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**invoice_id** | **str** | The invoice unique identifier, created by Lago. | 
**reason** | **str** | The reason of the credit note creation. Possible values are &#x60;duplicated_charge&#x60;, &#x60;product_unsatisfactory&#x60;, &#x60;order_change&#x60;, &#x60;order_cancellation&#x60;, &#x60;fraudulent_charge&#x60; or &#x60;other&#x60;. | [optional] 
**description** | **str** | The description of the credit note. | [optional] 
**credit_amount_cents** | **int** | The total amount to be credited on the customer balance. | [optional] 
**refund_amount_cents** | **int** | The total amount to be refunded to the customer. | [optional] 
**items** | [**List[CreditNoteEstimateInputCreditNoteItemsInner]**](CreditNoteEstimateInputCreditNoteItemsInner.md) | The list of credit note’s items. | 

## Example

```python
from lago_api.models.credit_note_create_input_credit_note import CreditNoteCreateInputCreditNote

# TODO update the JSON string below
json = "{}"
# create an instance of CreditNoteCreateInputCreditNote from a JSON string
credit_note_create_input_credit_note_instance = CreditNoteCreateInputCreditNote.from_json(json)
# print the JSON string representation of the object
print CreditNoteCreateInputCreditNote.to_json()

# convert the object into a dict
credit_note_create_input_credit_note_dict = credit_note_create_input_credit_note_instance.to_dict()
# create an instance of CreditNoteCreateInputCreditNote from a dict
credit_note_create_input_credit_note_form_dict = credit_note_create_input_credit_note.from_dict(credit_note_create_input_credit_note_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


