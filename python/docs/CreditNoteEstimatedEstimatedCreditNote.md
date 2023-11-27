# CreditNoteEstimatedEstimatedCreditNote


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**lago_invoice_id** | **str** | Unique identifier assigned to the invoice that the credit note belongs to | 
**invoice_number** | **str** | The invoice unique number, related to the credit note. | 
**currency** | [**Currency**](Currency.md) |  | 
**taxes_amount_cents** | **int** | The tax amount of the credit note, expressed in cents. | 
**taxes_rate** | **float** | The tax rate associated with this specific credit note. | 
**sub_total_excluding_taxes_amount_cents** | **int** | The subtotal of the credit note excluding any applicable taxes, expressed in cents. | 
**max_creditable_amount_cents** | **int** | The credited amount of the credit note, expressed in cents. | 
**max_refundable_amount_cents** | **int** | The refunded amount of the credit note, expressed in cents. | 
**coupons_adjustment_amount_cents** | **int** | The pro-rated amount of the coupons applied to the source invoice. | 
**items** | [**List[CreditNoteEstimatedEstimatedCreditNoteItemsInner]**](CreditNoteEstimatedEstimatedCreditNoteItemsInner.md) | Array of credit note’s items. | 
**applied_taxes** | [**List[CreditNoteEstimatedEstimatedCreditNoteAppliedTaxesInner]**](CreditNoteEstimatedEstimatedCreditNoteAppliedTaxesInner.md) |  | [optional] 

## Example

```python
from lago_api.models.credit_note_estimated_estimated_credit_note import CreditNoteEstimatedEstimatedCreditNote

# TODO update the JSON string below
json = "{}"
# create an instance of CreditNoteEstimatedEstimatedCreditNote from a JSON string
credit_note_estimated_estimated_credit_note_instance = CreditNoteEstimatedEstimatedCreditNote.from_json(json)
# print the JSON string representation of the object
print CreditNoteEstimatedEstimatedCreditNote.to_json()

# convert the object into a dict
credit_note_estimated_estimated_credit_note_dict = credit_note_estimated_estimated_credit_note_instance.to_dict()
# create an instance of CreditNoteEstimatedEstimatedCreditNote from a dict
credit_note_estimated_estimated_credit_note_form_dict = credit_note_estimated_estimated_credit_note.from_dict(credit_note_estimated_estimated_credit_note_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


