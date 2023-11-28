# InvoiceObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LagoId** | **string** | Unique identifier assigned to the fee within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the fee’s record within the Lago system. | 
**SequentialId** | **int32** | This ID helps in uniquely identifying and organizing the invoices associated with a specific customer. It provides a sequential numbering system specific to the customer, allowing for easy tracking and management of invoices within the customer&#39;s context. | 
**Number** | **string** | The unique number assigned to the invoice. This number serves as a distinct identifier for the invoice and helps in differentiating it from other invoices in the system. | 
**IssuingDate** | **string** | The date when the invoice was issued. It is provided in the ISO 8601 date format. | 
**PaymentDueDate** | Pointer to **string** | The payment due date for the invoice, specified in the ISO 8601 date format. | [optional] 
**NetPaymentTerm** | Pointer to **int32** | The net payment term, expressed in days, specifies the duration within which a customer is expected to remit payment after the invoice is finalized. | [optional] 
**InvoiceType** | **string** | The type of invoice issued. Possible values are &#x60;subscription&#x60;, &#x60;one-off&#x60; or &#x60;credit&#x60;. | 
**Status** | **string** | The status of the invoice. It indicates the current state of the invoice and can have two possible values: - &#x60;draft&#x60;: the invoice is in the draft state, waiting for the end of the grace period to be finalized. During this period, events can still be ingested and added to the invoice. - &#x60;finalized&#x60;: the invoice has been issued and finalized. In this state, events cannot be ingested or added to the invoice anymore. | 
**PaymentStatus** | **string** | The status of the payment associated with the invoice. It can have one of the following values: - &#x60;pending&#x60;: the payment is pending, waiting for payment processing in Stripe or when the invoice is emitted but users have not updated the payment status through the endpoint. - &#x60;succeeded&#x60;: the payment of the invoice has been successfully processed. - &#x60;failed&#x60;: the payment of the invoice has failed or encountered an error during processing. | 
**Currency** | [**Currency**](Currency.md) |  | 
**FeesAmountCents** | **int32** | The total sum of fees amount in cents. It calculates the cumulative amount of all the fees associated with the invoice, providing a consolidated value. | 
**CouponsAmountCents** | **int32** | The total sum of all coupons discounted on the invoice. It calculates the cumulative discount amount applied by coupons, expressed in cents. | 
**CreditNotesAmountCents** | **int32** | The total sum of all credit notes discounted on the invoice. It calculates the cumulative discount amount applied by credit notes, expressed in cents. | 
**SubTotalExcludingTaxesAmountCents** | **int32** | Subtotal amount, excluding taxes, expressed in cents. This field depends on the version number. Here are the definitions based on the version: - Version 1: is equal to the sum of &#x60;fees_amount_cents&#x60;, minus &#x60;coupons_amount_cents&#x60;, and minus &#x60;prepaid_credit_amount_cents&#x60;. - Version 2: is equal to the &#x60;fees_amount_cents&#x60;. - Version 3: is equal to the &#x60;fees_amount_cents&#x60;, minus &#x60;coupons_amount_cents&#x60; | 
**TaxesAmountCents** | **int32** | The sum of tax amount associated with the invoice, expressed in cents. | 
**SubTotalIncludingTaxesAmountCents** | **int32** | Subtotal amount, including taxes, expressed in cents. This field depends on the version number. Here are the definitions based on the version: - Version 1: is equal to the &#x60;total_amount_cents&#x60;. - Version 2: is equal to the sum of &#x60;fees_amount_cents&#x60; and &#x60;taxes_amount_cents&#x60;. - Version 3: is equal to the sum &#x60;sub_total_excluding_taxes_amount_cents&#x60; and &#x60;taxes_amount_cents&#x60; | 
**PrepaidCreditAmountCents** | **int32** | The total sum of all prepaid credits discounted on the invoice. It calculates the cumulative discount amount applied by prepaid credits, expressed in cents. | 
**TotalAmountCents** | **int32** | The sum of the amount and taxes amount on the invoice, expressed in cents. It calculates the total financial value of the invoice, including both the original amount and any applicable taxes. | 
**VersionNumber** | **int32** |  | 
**FileUrl** | Pointer to **string** | Contains the URL that provides direct access to the invoice PDF file. You can use this URL to download or view the PDF document of the invoice | [optional] 
**Customer** | Pointer to [**InvoiceObjectCustomer**](InvoiceObjectCustomer.md) |  | [optional] 
**Metadata** | Pointer to [**[]InvoiceMetadataObject**](InvoiceMetadataObject.md) |  | [optional] 
**AppliedTaxes** | Pointer to [**[]InvoiceAppliedTaxObject**](InvoiceAppliedTaxObject.md) |  | [optional] 

## Methods

### NewInvoiceObject

`func NewInvoiceObject(lagoId string, sequentialId int32, number string, issuingDate string, invoiceType string, status string, paymentStatus string, currency Currency, feesAmountCents int32, couponsAmountCents int32, creditNotesAmountCents int32, subTotalExcludingTaxesAmountCents int32, taxesAmountCents int32, subTotalIncludingTaxesAmountCents int32, prepaidCreditAmountCents int32, totalAmountCents int32, versionNumber int32, ) *InvoiceObject`

NewInvoiceObject instantiates a new InvoiceObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceObjectWithDefaults

`func NewInvoiceObjectWithDefaults() *InvoiceObject`

NewInvoiceObjectWithDefaults instantiates a new InvoiceObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLagoId

`func (o *InvoiceObject) GetLagoId() string`

GetLagoId returns the LagoId field if non-nil, zero value otherwise.

### GetLagoIdOk

`func (o *InvoiceObject) GetLagoIdOk() (*string, bool)`

GetLagoIdOk returns a tuple with the LagoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagoId

`func (o *InvoiceObject) SetLagoId(v string)`

SetLagoId sets LagoId field to given value.


### GetSequentialId

`func (o *InvoiceObject) GetSequentialId() int32`

GetSequentialId returns the SequentialId field if non-nil, zero value otherwise.

### GetSequentialIdOk

`func (o *InvoiceObject) GetSequentialIdOk() (*int32, bool)`

GetSequentialIdOk returns a tuple with the SequentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequentialId

`func (o *InvoiceObject) SetSequentialId(v int32)`

SetSequentialId sets SequentialId field to given value.


### GetNumber

`func (o *InvoiceObject) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *InvoiceObject) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *InvoiceObject) SetNumber(v string)`

SetNumber sets Number field to given value.


### GetIssuingDate

`func (o *InvoiceObject) GetIssuingDate() string`

GetIssuingDate returns the IssuingDate field if non-nil, zero value otherwise.

### GetIssuingDateOk

`func (o *InvoiceObject) GetIssuingDateOk() (*string, bool)`

GetIssuingDateOk returns a tuple with the IssuingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingDate

`func (o *InvoiceObject) SetIssuingDate(v string)`

SetIssuingDate sets IssuingDate field to given value.


### GetPaymentDueDate

`func (o *InvoiceObject) GetPaymentDueDate() string`

GetPaymentDueDate returns the PaymentDueDate field if non-nil, zero value otherwise.

### GetPaymentDueDateOk

`func (o *InvoiceObject) GetPaymentDueDateOk() (*string, bool)`

GetPaymentDueDateOk returns a tuple with the PaymentDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDueDate

`func (o *InvoiceObject) SetPaymentDueDate(v string)`

SetPaymentDueDate sets PaymentDueDate field to given value.

### HasPaymentDueDate

`func (o *InvoiceObject) HasPaymentDueDate() bool`

HasPaymentDueDate returns a boolean if a field has been set.

### GetNetPaymentTerm

`func (o *InvoiceObject) GetNetPaymentTerm() int32`

GetNetPaymentTerm returns the NetPaymentTerm field if non-nil, zero value otherwise.

### GetNetPaymentTermOk

`func (o *InvoiceObject) GetNetPaymentTermOk() (*int32, bool)`

GetNetPaymentTermOk returns a tuple with the NetPaymentTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetPaymentTerm

`func (o *InvoiceObject) SetNetPaymentTerm(v int32)`

SetNetPaymentTerm sets NetPaymentTerm field to given value.

### HasNetPaymentTerm

`func (o *InvoiceObject) HasNetPaymentTerm() bool`

HasNetPaymentTerm returns a boolean if a field has been set.

### GetInvoiceType

`func (o *InvoiceObject) GetInvoiceType() string`

GetInvoiceType returns the InvoiceType field if non-nil, zero value otherwise.

### GetInvoiceTypeOk

`func (o *InvoiceObject) GetInvoiceTypeOk() (*string, bool)`

GetInvoiceTypeOk returns a tuple with the InvoiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceType

`func (o *InvoiceObject) SetInvoiceType(v string)`

SetInvoiceType sets InvoiceType field to given value.


### GetStatus

`func (o *InvoiceObject) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InvoiceObject) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InvoiceObject) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetPaymentStatus

`func (o *InvoiceObject) GetPaymentStatus() string`

GetPaymentStatus returns the PaymentStatus field if non-nil, zero value otherwise.

### GetPaymentStatusOk

`func (o *InvoiceObject) GetPaymentStatusOk() (*string, bool)`

GetPaymentStatusOk returns a tuple with the PaymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentStatus

`func (o *InvoiceObject) SetPaymentStatus(v string)`

SetPaymentStatus sets PaymentStatus field to given value.


### GetCurrency

`func (o *InvoiceObject) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *InvoiceObject) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *InvoiceObject) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.


### GetFeesAmountCents

`func (o *InvoiceObject) GetFeesAmountCents() int32`

GetFeesAmountCents returns the FeesAmountCents field if non-nil, zero value otherwise.

### GetFeesAmountCentsOk

`func (o *InvoiceObject) GetFeesAmountCentsOk() (*int32, bool)`

GetFeesAmountCentsOk returns a tuple with the FeesAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeesAmountCents

`func (o *InvoiceObject) SetFeesAmountCents(v int32)`

SetFeesAmountCents sets FeesAmountCents field to given value.


### GetCouponsAmountCents

`func (o *InvoiceObject) GetCouponsAmountCents() int32`

GetCouponsAmountCents returns the CouponsAmountCents field if non-nil, zero value otherwise.

### GetCouponsAmountCentsOk

`func (o *InvoiceObject) GetCouponsAmountCentsOk() (*int32, bool)`

GetCouponsAmountCentsOk returns a tuple with the CouponsAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponsAmountCents

`func (o *InvoiceObject) SetCouponsAmountCents(v int32)`

SetCouponsAmountCents sets CouponsAmountCents field to given value.


### GetCreditNotesAmountCents

`func (o *InvoiceObject) GetCreditNotesAmountCents() int32`

GetCreditNotesAmountCents returns the CreditNotesAmountCents field if non-nil, zero value otherwise.

### GetCreditNotesAmountCentsOk

`func (o *InvoiceObject) GetCreditNotesAmountCentsOk() (*int32, bool)`

GetCreditNotesAmountCentsOk returns a tuple with the CreditNotesAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditNotesAmountCents

`func (o *InvoiceObject) SetCreditNotesAmountCents(v int32)`

SetCreditNotesAmountCents sets CreditNotesAmountCents field to given value.


### GetSubTotalExcludingTaxesAmountCents

`func (o *InvoiceObject) GetSubTotalExcludingTaxesAmountCents() int32`

GetSubTotalExcludingTaxesAmountCents returns the SubTotalExcludingTaxesAmountCents field if non-nil, zero value otherwise.

### GetSubTotalExcludingTaxesAmountCentsOk

`func (o *InvoiceObject) GetSubTotalExcludingTaxesAmountCentsOk() (*int32, bool)`

GetSubTotalExcludingTaxesAmountCentsOk returns a tuple with the SubTotalExcludingTaxesAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTotalExcludingTaxesAmountCents

`func (o *InvoiceObject) SetSubTotalExcludingTaxesAmountCents(v int32)`

SetSubTotalExcludingTaxesAmountCents sets SubTotalExcludingTaxesAmountCents field to given value.


### GetTaxesAmountCents

`func (o *InvoiceObject) GetTaxesAmountCents() int32`

GetTaxesAmountCents returns the TaxesAmountCents field if non-nil, zero value otherwise.

### GetTaxesAmountCentsOk

`func (o *InvoiceObject) GetTaxesAmountCentsOk() (*int32, bool)`

GetTaxesAmountCentsOk returns a tuple with the TaxesAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxesAmountCents

`func (o *InvoiceObject) SetTaxesAmountCents(v int32)`

SetTaxesAmountCents sets TaxesAmountCents field to given value.


### GetSubTotalIncludingTaxesAmountCents

`func (o *InvoiceObject) GetSubTotalIncludingTaxesAmountCents() int32`

GetSubTotalIncludingTaxesAmountCents returns the SubTotalIncludingTaxesAmountCents field if non-nil, zero value otherwise.

### GetSubTotalIncludingTaxesAmountCentsOk

`func (o *InvoiceObject) GetSubTotalIncludingTaxesAmountCentsOk() (*int32, bool)`

GetSubTotalIncludingTaxesAmountCentsOk returns a tuple with the SubTotalIncludingTaxesAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTotalIncludingTaxesAmountCents

`func (o *InvoiceObject) SetSubTotalIncludingTaxesAmountCents(v int32)`

SetSubTotalIncludingTaxesAmountCents sets SubTotalIncludingTaxesAmountCents field to given value.


### GetPrepaidCreditAmountCents

`func (o *InvoiceObject) GetPrepaidCreditAmountCents() int32`

GetPrepaidCreditAmountCents returns the PrepaidCreditAmountCents field if non-nil, zero value otherwise.

### GetPrepaidCreditAmountCentsOk

`func (o *InvoiceObject) GetPrepaidCreditAmountCentsOk() (*int32, bool)`

GetPrepaidCreditAmountCentsOk returns a tuple with the PrepaidCreditAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrepaidCreditAmountCents

`func (o *InvoiceObject) SetPrepaidCreditAmountCents(v int32)`

SetPrepaidCreditAmountCents sets PrepaidCreditAmountCents field to given value.


### GetTotalAmountCents

`func (o *InvoiceObject) GetTotalAmountCents() int32`

GetTotalAmountCents returns the TotalAmountCents field if non-nil, zero value otherwise.

### GetTotalAmountCentsOk

`func (o *InvoiceObject) GetTotalAmountCentsOk() (*int32, bool)`

GetTotalAmountCentsOk returns a tuple with the TotalAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountCents

`func (o *InvoiceObject) SetTotalAmountCents(v int32)`

SetTotalAmountCents sets TotalAmountCents field to given value.


### GetVersionNumber

`func (o *InvoiceObject) GetVersionNumber() int32`

GetVersionNumber returns the VersionNumber field if non-nil, zero value otherwise.

### GetVersionNumberOk

`func (o *InvoiceObject) GetVersionNumberOk() (*int32, bool)`

GetVersionNumberOk returns a tuple with the VersionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionNumber

`func (o *InvoiceObject) SetVersionNumber(v int32)`

SetVersionNumber sets VersionNumber field to given value.


### GetFileUrl

`func (o *InvoiceObject) GetFileUrl() string`

GetFileUrl returns the FileUrl field if non-nil, zero value otherwise.

### GetFileUrlOk

`func (o *InvoiceObject) GetFileUrlOk() (*string, bool)`

GetFileUrlOk returns a tuple with the FileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUrl

`func (o *InvoiceObject) SetFileUrl(v string)`

SetFileUrl sets FileUrl field to given value.

### HasFileUrl

`func (o *InvoiceObject) HasFileUrl() bool`

HasFileUrl returns a boolean if a field has been set.

### GetCustomer

`func (o *InvoiceObject) GetCustomer() InvoiceObjectCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *InvoiceObject) GetCustomerOk() (*InvoiceObjectCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *InvoiceObject) SetCustomer(v InvoiceObjectCustomer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *InvoiceObject) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetMetadata

`func (o *InvoiceObject) GetMetadata() []InvoiceMetadataObject`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *InvoiceObject) GetMetadataOk() (*[]InvoiceMetadataObject, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *InvoiceObject) SetMetadata(v []InvoiceMetadataObject)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *InvoiceObject) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetAppliedTaxes

`func (o *InvoiceObject) GetAppliedTaxes() []InvoiceAppliedTaxObject`

GetAppliedTaxes returns the AppliedTaxes field if non-nil, zero value otherwise.

### GetAppliedTaxesOk

`func (o *InvoiceObject) GetAppliedTaxesOk() (*[]InvoiceAppliedTaxObject, bool)`

GetAppliedTaxesOk returns a tuple with the AppliedTaxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedTaxes

`func (o *InvoiceObject) SetAppliedTaxes(v []InvoiceAppliedTaxObject)`

SetAppliedTaxes sets AppliedTaxes field to given value.

### HasAppliedTaxes

`func (o *InvoiceObject) HasAppliedTaxes() bool`

HasAppliedTaxes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


