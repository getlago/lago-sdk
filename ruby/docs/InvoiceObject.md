# LagoAPI::InvoiceObject

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **lago_id** | **String** | Unique identifier assigned to the fee within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the fee’s record within the Lago system. |  |
| **sequential_id** | **Integer** | This ID helps in uniquely identifying and organizing the invoices associated with a specific customer. It provides a sequential numbering system specific to the customer, allowing for easy tracking and management of invoices within the customer&#39;s context. |  |
| **number** | **String** | The unique number assigned to the invoice. This number serves as a distinct identifier for the invoice and helps in differentiating it from other invoices in the system. |  |
| **issuing_date** | **Date** | The date when the invoice was issued. It is provided in the ISO 8601 date format. |  |
| **payment_due_date** | **Date** | The payment due date for the invoice, specified in the ISO 8601 date format. | [optional] |
| **net_payment_term** | **Integer** | The net payment term, expressed in days, specifies the duration within which a customer is expected to remit payment after the invoice is finalized. | [optional] |
| **invoice_type** | **String** | The type of invoice issued. Possible values are &#x60;subscription&#x60;, &#x60;one-off&#x60; or &#x60;credit&#x60;. |  |
| **status** | **String** | The status of the invoice. It indicates the current state of the invoice and can have two possible values: - &#x60;draft&#x60;: the invoice is in the draft state, waiting for the end of the grace period to be finalized. During this period, events can still be ingested and added to the invoice. - &#x60;finalized&#x60;: the invoice has been issued and finalized. In this state, events cannot be ingested or added to the invoice anymore. |  |
| **payment_status** | **String** | The status of the payment associated with the invoice. It can have one of the following values: - &#x60;pending&#x60;: the payment is pending, waiting for payment processing in Stripe or when the invoice is emitted but users have not updated the payment status through the endpoint. - &#x60;succeeded&#x60;: the payment of the invoice has been successfully processed. - &#x60;failed&#x60;: the payment of the invoice has failed or encountered an error during processing. |  |
| **currency** | [**Currency**](Currency.md) |  |  |
| **fees_amount_cents** | **Integer** | The total sum of fees amount in cents. It calculates the cumulative amount of all the fees associated with the invoice, providing a consolidated value. |  |
| **coupons_amount_cents** | **Integer** | The total sum of all coupons discounted on the invoice. It calculates the cumulative discount amount applied by coupons, expressed in cents. |  |
| **credit_notes_amount_cents** | **Integer** | The total sum of all credit notes discounted on the invoice. It calculates the cumulative discount amount applied by credit notes, expressed in cents. |  |
| **sub_total_excluding_taxes_amount_cents** | **Integer** | Subtotal amount, excluding taxes, expressed in cents. This field depends on the version number. Here are the definitions based on the version: - Version 1: is equal to the sum of &#x60;fees_amount_cents&#x60;, minus &#x60;coupons_amount_cents&#x60;, and minus &#x60;prepaid_credit_amount_cents&#x60;. - Version 2: is equal to the &#x60;fees_amount_cents&#x60;. - Version 3: is equal to the &#x60;fees_amount_cents&#x60;, minus &#x60;coupons_amount_cents&#x60; |  |
| **taxes_amount_cents** | **Integer** | The sum of tax amount associated with the invoice, expressed in cents. |  |
| **sub_total_including_taxes_amount_cents** | **Integer** | Subtotal amount, including taxes, expressed in cents. This field depends on the version number. Here are the definitions based on the version: - Version 1: is equal to the &#x60;total_amount_cents&#x60;. - Version 2: is equal to the sum of &#x60;fees_amount_cents&#x60; and &#x60;taxes_amount_cents&#x60;. - Version 3: is equal to the sum &#x60;sub_total_excluding_taxes_amount_cents&#x60; and &#x60;taxes_amount_cents&#x60; |  |
| **prepaid_credit_amount_cents** | **Integer** | The total sum of all prepaid credits discounted on the invoice. It calculates the cumulative discount amount applied by prepaid credits, expressed in cents. |  |
| **total_amount_cents** | **Integer** | The sum of the amount and taxes amount on the invoice, expressed in cents. It calculates the total financial value of the invoice, including both the original amount and any applicable taxes. |  |
| **version_number** | **Integer** |  |  |
| **file_url** | **String** | Contains the URL that provides direct access to the invoice PDF file. You can use this URL to download or view the PDF document of the invoice | [optional] |
| **customer** | [**InvoiceObjectCustomer**](InvoiceObjectCustomer.md) |  | [optional] |
| **metadata** | [**Array&lt;InvoiceMetadataObject&gt;**](InvoiceMetadataObject.md) |  | [optional] |
| **applied_taxes** | [**Array&lt;InvoiceAppliedTaxObject&gt;**](InvoiceAppliedTaxObject.md) |  | [optional] |

## Example

```ruby
require 'lago_ruby'

instance = LagoAPI::InvoiceObject.new(
  lago_id: 1a901a90-1a90-1a90-1a90-1a901a901a90,
  sequential_id: 2,
  number: LAG-1234-001-002,
  issuing_date: Sat Apr 30 00:00:00 UTC 2022,
  payment_due_date: Sat Apr 30 00:00:00 UTC 2022,
  net_payment_term: 30,
  invoice_type: subscription,
  status: finalized,
  payment_status: succeeded,
  currency: null,
  fees_amount_cents: 100,
  coupons_amount_cents: 10,
  credit_notes_amount_cents: 10,
  sub_total_excluding_taxes_amount_cents: 100,
  taxes_amount_cents: 20,
  sub_total_including_taxes_amount_cents: 120,
  prepaid_credit_amount_cents: 0,
  total_amount_cents: 100,
  version_number: 3,
  file_url: https://getlago.com/invoice/file,
  customer: null,
  metadata: null,
  applied_taxes: null
)
```

