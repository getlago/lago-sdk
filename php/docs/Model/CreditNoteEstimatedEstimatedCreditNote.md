# # CreditNoteEstimatedEstimatedCreditNote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**lago_invoice_id** | **string** | Unique identifier assigned to the invoice that the credit note belongs to |
**invoice_number** | **string** | The invoice unique number, related to the credit note. |
**currency** | [**Currency**](Currency.md) |  |
**taxes_amount_cents** | **int** | The tax amount of the credit note, expressed in cents. |
**taxes_rate** | **float** | The tax rate associated with this specific credit note. |
**sub_total_excluding_taxes_amount_cents** | **int** | The subtotal of the credit note excluding any applicable taxes, expressed in cents. |
**max_creditable_amount_cents** | **int** | The credited amount of the credit note, expressed in cents. |
**max_refundable_amount_cents** | **int** | The refunded amount of the credit note, expressed in cents. |
**coupons_adjustment_amount_cents** | **int** | The pro-rated amount of the coupons applied to the source invoice. |
**items** | [**\OpenAPI\Client\Model\CreditNoteEstimatedEstimatedCreditNoteItemsInner[]**](CreditNoteEstimatedEstimatedCreditNoteItemsInner.md) | Array of credit note’s items. |
**applied_taxes** | [**\OpenAPI\Client\Model\CreditNoteEstimatedEstimatedCreditNoteAppliedTaxesInner[]**](CreditNoteEstimatedEstimatedCreditNoteAppliedTaxesInner.md) |  | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
