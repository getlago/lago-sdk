# CreditNoteObject


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**lago_id** | **str** | The credit note unique identifier, created by Lago. | 
**sequential_id** | **int** | The sequential identifier of the credit note, specifically scoped on the associated invoice. It provides a unique numerical identifier for the credit note within the context of the invoice. | 
**number** | **str** | The credit note unique number. | 
**lago_invoice_id** | **str** | Unique identifier assigned to the invoice that the credit note belongs to | 
**invoice_number** | **str** | The invoice unique number, related to the credit note. | 
**issuing_date** | **date** | The date of creation of the credit note. It follows the ISO 8601 date format and provides the specific date when the credit note was created. | 
**credit_status** | **str** | The status of the credit portion of the credit note. It indicates the current state or condition of the credit amount associated with the credit note. The possible values for this field are:  - &#x60;available&#x60;: this status indicates that an amount remains available for future usage. The credit can be applied towards future transactions or invoices. - &#x60;consumed&#x60;: this status indicates that the credit amount has been fully consumed. The remaining amount is 0, indicating that the credit has been utilized in its entirety. - &#x60;voided&#x60;: this status indicates that the remaining amount of the credit cannot be used any further. The credit has been voided and is no longer available for application or redemption. | [optional] 
**refund_status** | **str** | The status of the refund portion of the credit note. It indicates the current state or condition of the refund associated with the credit note. The possible values for this field are:  - &#x60;pending&#x60;: this status indicates that the refund is pending execution. The refund request has been initiated but has not been processed or completed yet. - &#x60;succeeded&#x60;: this status indicates that the refund has been successfully executed. The refund amount has been processed and returned to the customer or the designated recipient. - &#x60;failed&#x60;: this status indicates that the refund failed to execute. The refund request encountered an error or unsuccessful processing, and the refund amount could not be returned. | [optional] 
**reason** | **str** | The reason of the credit note creation. Possible values are &#x60;duplicated_charge&#x60;, &#x60;product_unsatisfactory&#x60;, &#x60;order_change&#x60;, &#x60;order_cancellation&#x60;, &#x60;fraudulent_charge&#x60; or &#x60;other&#x60;. | 
**description** | **str** | The description of the credit note. | [optional] 
**currency** | [**Currency**](Currency.md) |  | 
**total_amount_cents** | **int** | The total amount of the credit note, expressed in cents. | 
**taxes_amount_cents** | **int** | The tax amount of the credit note, expressed in cents. | 
**taxes_rate** | **float** | The tax rate associated with this specific credit note. | 
**sub_total_excluding_taxes_amount_cents** | **int** | The subtotal of the credit note excluding any applicable taxes, expressed in cents. | 
**balance_amount_cents** | **int** | The remaining credit note amount, expressed in cents. | 
**credit_amount_cents** | **int** | The credited amount of the credit note, expressed in cents. | 
**refund_amount_cents** | **int** | The refunded amount of the credit note, expressed in cents. | 
**coupons_adjustment_amount_cents** | **int** | The pro-rated amount of the coupons applied to the source invoice. | 
**created_at** | **datetime** | The date when the credit note was created. It is expressed in Coordinated Universal Time (UTC). | 
**updated_at** | **datetime** | The date when the credit note was last updated. It is expressed in Coordinated Universal Time (UTC). | 
**file_url** | **str** | The PDF file of the credit note. | [optional] 
**items** | [**List[CreditNoteItemObject]**](CreditNoteItemObject.md) | Array of credit note’s items. | [optional] 
**applied_taxes** | [**List[CreditNoteAppliedTaxObject]**](CreditNoteAppliedTaxObject.md) |  | [optional] 

## Example

```python
from lago_api.models.credit_note_object import CreditNoteObject

# TODO update the JSON string below
json = "{}"
# create an instance of CreditNoteObject from a JSON string
credit_note_object_instance = CreditNoteObject.from_json(json)
# print the JSON string representation of the object
print CreditNoteObject.to_json()

# convert the object into a dict
credit_note_object_dict = credit_note_object_instance.to_dict()
# create an instance of CreditNoteObject from a dict
credit_note_object_form_dict = credit_note_object.from_dict(credit_note_object_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


