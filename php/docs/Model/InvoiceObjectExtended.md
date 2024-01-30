# # InvoiceObjectExtended

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**lago_id** | **string** | Unique identifier assigned to the fee within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the fee’s record within the Lago system. |
**sequential_id** | **int** | This ID helps in uniquely identifying and organizing the invoices associated with a specific customer. It provides a sequential numbering system specific to the customer, allowing for easy tracking and management of invoices within the customer&#39;s context. |
**number** | **string** | The unique number assigned to the invoice. This number serves as a distinct identifier for the invoice and helps in differentiating it from other invoices in the system. |
**issuing_date** | **\DateTime** | The date when the invoice was issued. It is provided in the ISO 8601 date format. |
**payment_due_date** | **\DateTime** | The payment due date for the invoice, specified in the ISO 8601 date format. | [optional]
**net_payment_term** | **int** | The net payment term, expressed in days, specifies the duration within which a customer is expected to remit payment after the invoice is finalized. | [optional]
**invoice_type** | **string** | The type of invoice issued. Possible values are &#x60;subscription&#x60;, &#x60;one-off&#x60; or &#x60;credit&#x60;. |
**status** | **string** | The status of the invoice. It indicates the current state of the invoice and can have two possible values: - &#x60;draft&#x60;: the invoice is in the draft state, waiting for the end of the grace period to be finalized. During this period, events can still be ingested and added to the invoice. - &#x60;finalized&#x60;: the invoice has been issued and finalized. In this state, events cannot be ingested or added to the invoice anymore. - &#x60;voided&#x60;: the invoice has been issued and subsequently voided. In this state, events cannot be ingested or added to the invoice anymore. |
**payment_status** | **string** | The status of the payment associated with the invoice. It can have one of the following values: - &#x60;pending&#x60;: the payment is pending, waiting for payment processing in Stripe or when the invoice is emitted but users have not updated the payment status through the endpoint. - &#x60;succeeded&#x60;: the payment of the invoice has been successfully processed. - &#x60;failed&#x60;: the payment of the invoice has failed or encountered an error during processing. |
**currency** | [**Currency**](Currency.md) |  |
**fees_amount_cents** | **int** | The total sum of fees amount in cents. It calculates the cumulative amount of all the fees associated with the invoice, providing a consolidated value. |
**coupons_amount_cents** | **int** | The total sum of all coupons discounted on the invoice. It calculates the cumulative discount amount applied by coupons, expressed in cents. |
**credit_notes_amount_cents** | **int** | The total sum of all credit notes discounted on the invoice. It calculates the cumulative discount amount applied by credit notes, expressed in cents. |
**sub_total_excluding_taxes_amount_cents** | **int** | Subtotal amount, excluding taxes, expressed in cents. This field depends on the version number. Here are the definitions based on the version: - Version 1: is equal to the sum of &#x60;fees_amount_cents&#x60;, minus &#x60;coupons_amount_cents&#x60;, and minus &#x60;prepaid_credit_amount_cents&#x60;. - Version 2: is equal to the &#x60;fees_amount_cents&#x60;. - Version 3: is equal to the &#x60;fees_amount_cents&#x60;, minus &#x60;coupons_amount_cents&#x60; |
**taxes_amount_cents** | **int** | The sum of tax amount associated with the invoice, expressed in cents. |
**sub_total_including_taxes_amount_cents** | **int** | Subtotal amount, including taxes, expressed in cents. This field depends on the version number. Here are the definitions based on the version: - Version 1: is equal to the &#x60;total_amount_cents&#x60;. - Version 2: is equal to the sum of &#x60;fees_amount_cents&#x60; and &#x60;taxes_amount_cents&#x60;. - Version 3: is equal to the sum &#x60;sub_total_excluding_taxes_amount_cents&#x60; and &#x60;taxes_amount_cents&#x60; |
**prepaid_credit_amount_cents** | **int** | The total sum of all prepaid credits discounted on the invoice. It calculates the cumulative discount amount applied by prepaid credits, expressed in cents. |
**total_amount_cents** | **int** | The sum of the amount and taxes amount on the invoice, expressed in cents. It calculates the total financial value of the invoice, including both the original amount and any applicable taxes. |
**version_number** | **int** |  |
**file_url** | **string** | Contains the URL that provides direct access to the invoice PDF file. You can use this URL to download or view the PDF document of the invoice | [optional]
**customer** | [**\OpenAPI\Client\Model\InvoiceObjectCustomer**](InvoiceObjectCustomer.md) |  | [optional]
**metadata** | [**\OpenAPI\Client\Model\InvoiceMetadataObject[]**](InvoiceMetadataObject.md) |  | [optional]
**applied_taxes** | [**\OpenAPI\Client\Model\InvoiceAppliedTaxObject[]**](InvoiceAppliedTaxObject.md) |  | [optional]
**credits** | [**\OpenAPI\Client\Model\CreditObject[]**](CreditObject.md) |  | [optional]
**fees** | [**\OpenAPI\Client\Model\FeeObject[]**](FeeObject.md) |  | [optional]
**subscriptions** | [**\OpenAPI\Client\Model\SubscriptionObject[]**](SubscriptionObject.md) |  | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
