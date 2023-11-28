/*
Lago API documentation

Lago API allows your application to push customer information and metrics (events) from your application to the billing application.

API version: 0.52.0-beta
Contact: tech@getlago.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lagoapi

import (
	"encoding/json"
	"fmt"
)

// checks if the InvoiceObjectExtended type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvoiceObjectExtended{}

// InvoiceObjectExtended struct for InvoiceObjectExtended
type InvoiceObjectExtended struct {
	// Unique identifier assigned to the fee within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the fee’s record within the Lago system.
	LagoId string `json:"lago_id"`
	// This ID helps in uniquely identifying and organizing the invoices associated with a specific customer. It provides a sequential numbering system specific to the customer, allowing for easy tracking and management of invoices within the customer's context.
	SequentialId int32 `json:"sequential_id"`
	// The unique number assigned to the invoice. This number serves as a distinct identifier for the invoice and helps in differentiating it from other invoices in the system.
	Number string `json:"number"`
	// The date when the invoice was issued. It is provided in the ISO 8601 date format.
	IssuingDate string `json:"issuing_date"`
	// The payment due date for the invoice, specified in the ISO 8601 date format.
	PaymentDueDate *string `json:"payment_due_date,omitempty"`
	// The net payment term, expressed in days, specifies the duration within which a customer is expected to remit payment after the invoice is finalized.
	NetPaymentTerm *int32 `json:"net_payment_term,omitempty"`
	// The type of invoice issued. Possible values are `subscription`, `one-off` or `credit`.
	InvoiceType string `json:"invoice_type"`
	// The status of the invoice. It indicates the current state of the invoice and can have two possible values: - `draft`: the invoice is in the draft state, waiting for the end of the grace period to be finalized. During this period, events can still be ingested and added to the invoice. - `finalized`: the invoice has been issued and finalized. In this state, events cannot be ingested or added to the invoice anymore.
	Status string `json:"status"`
	// The status of the payment associated with the invoice. It can have one of the following values: - `pending`: the payment is pending, waiting for payment processing in Stripe or when the invoice is emitted but users have not updated the payment status through the endpoint. - `succeeded`: the payment of the invoice has been successfully processed. - `failed`: the payment of the invoice has failed or encountered an error during processing.
	PaymentStatus string `json:"payment_status"`
	Currency Currency `json:"currency"`
	// The total sum of fees amount in cents. It calculates the cumulative amount of all the fees associated with the invoice, providing a consolidated value.
	FeesAmountCents int32 `json:"fees_amount_cents"`
	// The total sum of all coupons discounted on the invoice. It calculates the cumulative discount amount applied by coupons, expressed in cents.
	CouponsAmountCents int32 `json:"coupons_amount_cents"`
	// The total sum of all credit notes discounted on the invoice. It calculates the cumulative discount amount applied by credit notes, expressed in cents.
	CreditNotesAmountCents int32 `json:"credit_notes_amount_cents"`
	// Subtotal amount, excluding taxes, expressed in cents. This field depends on the version number. Here are the definitions based on the version: - Version 1: is equal to the sum of `fees_amount_cents`, minus `coupons_amount_cents`, and minus `prepaid_credit_amount_cents`. - Version 2: is equal to the `fees_amount_cents`. - Version 3: is equal to the `fees_amount_cents`, minus `coupons_amount_cents`
	SubTotalExcludingTaxesAmountCents int32 `json:"sub_total_excluding_taxes_amount_cents"`
	// The sum of tax amount associated with the invoice, expressed in cents.
	TaxesAmountCents int32 `json:"taxes_amount_cents"`
	// Subtotal amount, including taxes, expressed in cents. This field depends on the version number. Here are the definitions based on the version: - Version 1: is equal to the `total_amount_cents`. - Version 2: is equal to the sum of `fees_amount_cents` and `taxes_amount_cents`. - Version 3: is equal to the sum `sub_total_excluding_taxes_amount_cents` and `taxes_amount_cents`
	SubTotalIncludingTaxesAmountCents int32 `json:"sub_total_including_taxes_amount_cents"`
	// The total sum of all prepaid credits discounted on the invoice. It calculates the cumulative discount amount applied by prepaid credits, expressed in cents.
	PrepaidCreditAmountCents int32 `json:"prepaid_credit_amount_cents"`
	// The sum of the amount and taxes amount on the invoice, expressed in cents. It calculates the total financial value of the invoice, including both the original amount and any applicable taxes.
	TotalAmountCents int32 `json:"total_amount_cents"`
	VersionNumber int32 `json:"version_number"`
	// Contains the URL that provides direct access to the invoice PDF file. You can use this URL to download or view the PDF document of the invoice
	FileUrl *string `json:"file_url,omitempty"`
	Customer *InvoiceObjectCustomer `json:"customer,omitempty"`
	Metadata []InvoiceMetadataObject `json:"metadata,omitempty"`
	AppliedTaxes []InvoiceAppliedTaxObject `json:"applied_taxes,omitempty"`
	Credits []CreditObject `json:"credits,omitempty"`
	Fees []FeeObject `json:"fees,omitempty"`
	Subscriptions []SubscriptionObject `json:"subscriptions,omitempty"`
}

type _InvoiceObjectExtended InvoiceObjectExtended

// NewInvoiceObjectExtended instantiates a new InvoiceObjectExtended object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceObjectExtended(lagoId string, sequentialId int32, number string, issuingDate string, invoiceType string, status string, paymentStatus string, currency Currency, feesAmountCents int32, couponsAmountCents int32, creditNotesAmountCents int32, subTotalExcludingTaxesAmountCents int32, taxesAmountCents int32, subTotalIncludingTaxesAmountCents int32, prepaidCreditAmountCents int32, totalAmountCents int32, versionNumber int32) *InvoiceObjectExtended {
	this := InvoiceObjectExtended{}
	this.LagoId = lagoId
	this.SequentialId = sequentialId
	this.Number = number
	this.IssuingDate = issuingDate
	this.InvoiceType = invoiceType
	this.Status = status
	this.PaymentStatus = paymentStatus
	this.Currency = currency
	this.FeesAmountCents = feesAmountCents
	this.CouponsAmountCents = couponsAmountCents
	this.CreditNotesAmountCents = creditNotesAmountCents
	this.SubTotalExcludingTaxesAmountCents = subTotalExcludingTaxesAmountCents
	this.TaxesAmountCents = taxesAmountCents
	this.SubTotalIncludingTaxesAmountCents = subTotalIncludingTaxesAmountCents
	this.PrepaidCreditAmountCents = prepaidCreditAmountCents
	this.TotalAmountCents = totalAmountCents
	this.VersionNumber = versionNumber
	return &this
}

// NewInvoiceObjectExtendedWithDefaults instantiates a new InvoiceObjectExtended object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceObjectExtendedWithDefaults() *InvoiceObjectExtended {
	this := InvoiceObjectExtended{}
	return &this
}

// GetLagoId returns the LagoId field value
func (o *InvoiceObjectExtended) GetLagoId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LagoId
}

// GetLagoIdOk returns a tuple with the LagoId field value
// and a boolean to check if the value has been set.
func (o *InvoiceObjectExtended) GetLagoIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LagoId, true
}

// SetLagoId sets field value
func (o *InvoiceObjectExtended) SetLagoId(v string) {
	o.LagoId = v
}

// GetSequentialId returns the SequentialId field value
func (o *InvoiceObjectExtended) GetSequentialId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SequentialId
}

// GetSequentialIdOk returns a tuple with the SequentialId field value
// and a boolean to check if the value has been set.
func (o *InvoiceObjectExtended) GetSequentialIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SequentialId, true
}

// SetSequentialId sets field value
func (o *InvoiceObjectExtended) SetSequentialId(v int32) {
	o.SequentialId = v
}

// GetNumber returns the Number field value
func (o *InvoiceObjectExtended) GetNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Number
}

// GetNumberOk returns a tuple with the Number field value
// and a boolean to check if the value has been set.
func (o *InvoiceObjectExtended) GetNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Number, true
}

// SetNumber sets field value
func (o *InvoiceObjectExtended) SetNumber(v string) {
	o.Number = v
}

// GetIssuingDate returns the IssuingDate field value
func (o *InvoiceObjectExtended) GetIssuingDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IssuingDate
}

// GetIssuingDateOk returns a tuple with the IssuingDate field value
// and a boolean to check if the value has been set.
func (o *InvoiceObjectExtended) GetIssuingDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssuingDate, true
}

// SetIssuingDate sets field value
func (o *InvoiceObjectExtended) SetIssuingDate(v string) {
	o.IssuingDate = v
}

// GetPaymentDueDate returns the PaymentDueDate field value if set, zero value otherwise.
func (o *InvoiceObjectExtended) GetPaymentDueDate() string {
	if o == nil || IsNil(o.PaymentDueDate) {
		var ret string
		return ret
	}
	return *o.PaymentDueDate
}

// GetPaymentDueDateOk returns a tuple with the PaymentDueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceObjectExtended) GetPaymentDueDateOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentDueDate) {
		return nil, false
	}
	return o.PaymentDueDate, true
}

// HasPaymentDueDate returns a boolean if a field has been set.
func (o *InvoiceObjectExtended) HasPaymentDueDate() bool {
	if o != nil && !IsNil(o.PaymentDueDate) {
		return true
	}

	return false
}

// SetPaymentDueDate gets a reference to the given string and assigns it to the PaymentDueDate field.
func (o *InvoiceObjectExtended) SetPaymentDueDate(v string) {
	o.PaymentDueDate = &v
}

// GetNetPaymentTerm returns the NetPaymentTerm field value if set, zero value otherwise.
func (o *InvoiceObjectExtended) GetNetPaymentTerm() int32 {
	if o == nil || IsNil(o.NetPaymentTerm) {
		var ret int32
		return ret
	}
	return *o.NetPaymentTerm
}

// GetNetPaymentTermOk returns a tuple with the NetPaymentTerm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceObjectExtended) GetNetPaymentTermOk() (*int32, bool) {
	if o == nil || IsNil(o.NetPaymentTerm) {
		return nil, false
	}
	return o.NetPaymentTerm, true
}

// HasNetPaymentTerm returns a boolean if a field has been set.
func (o *InvoiceObjectExtended) HasNetPaymentTerm() bool {
	if o != nil && !IsNil(o.NetPaymentTerm) {
		return true
	}

	return false
}

// SetNetPaymentTerm gets a reference to the given int32 and assigns it to the NetPaymentTerm field.
func (o *InvoiceObjectExtended) SetNetPaymentTerm(v int32) {
	o.NetPaymentTerm = &v
}

// GetInvoiceType returns the InvoiceType field value
func (o *InvoiceObjectExtended) GetInvoiceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InvoiceType
}

// GetInvoiceTypeOk returns a tuple with the InvoiceType field value
// and a boolean to check if the value has been set.
func (o *InvoiceObjectExtended) GetInvoiceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvoiceType, true
}

// SetInvoiceType sets field value
func (o *InvoiceObjectExtended) SetInvoiceType(v string) {
	o.InvoiceType = v
}

// GetStatus returns the Status field value
func (o *InvoiceObjectExtended) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *InvoiceObjectExtended) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *InvoiceObjectExtended) SetStatus(v string) {
	o.Status = v
}

// GetPaymentStatus returns the PaymentStatus field value
func (o *InvoiceObjectExtended) GetPaymentStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentStatus
}

// GetPaymentStatusOk returns a tuple with the PaymentStatus field value
// and a boolean to check if the value has been set.
func (o *InvoiceObjectExtended) GetPaymentStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentStatus, true
}

// SetPaymentStatus sets field value
func (o *InvoiceObjectExtended) SetPaymentStatus(v string) {
	o.PaymentStatus = v
}

// GetCurrency returns the Currency field value
func (o *InvoiceObjectExtended) GetCurrency() Currency {
	if o == nil {
		var ret Currency
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *InvoiceObjectExtended) GetCurrencyOk() (*Currency, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *InvoiceObjectExtended) SetCurrency(v Currency) {
	o.Currency = v
}

// GetFeesAmountCents returns the FeesAmountCents field value
func (o *InvoiceObjectExtended) GetFeesAmountCents() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FeesAmountCents
}

// GetFeesAmountCentsOk returns a tuple with the FeesAmountCents field value
// and a boolean to check if the value has been set.
func (o *InvoiceObjectExtended) GetFeesAmountCentsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeesAmountCents, true
}

// SetFeesAmountCents sets field value
func (o *InvoiceObjectExtended) SetFeesAmountCents(v int32) {
	o.FeesAmountCents = v
}

// GetCouponsAmountCents returns the CouponsAmountCents field value
func (o *InvoiceObjectExtended) GetCouponsAmountCents() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CouponsAmountCents
}

// GetCouponsAmountCentsOk returns a tuple with the CouponsAmountCents field value
// and a boolean to check if the value has been set.
func (o *InvoiceObjectExtended) GetCouponsAmountCentsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CouponsAmountCents, true
}

// SetCouponsAmountCents sets field value
func (o *InvoiceObjectExtended) SetCouponsAmountCents(v int32) {
	o.CouponsAmountCents = v
}

// GetCreditNotesAmountCents returns the CreditNotesAmountCents field value
func (o *InvoiceObjectExtended) GetCreditNotesAmountCents() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CreditNotesAmountCents
}

// GetCreditNotesAmountCentsOk returns a tuple with the CreditNotesAmountCents field value
// and a boolean to check if the value has been set.
func (o *InvoiceObjectExtended) GetCreditNotesAmountCentsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreditNotesAmountCents, true
}

// SetCreditNotesAmountCents sets field value
func (o *InvoiceObjectExtended) SetCreditNotesAmountCents(v int32) {
	o.CreditNotesAmountCents = v
}

// GetSubTotalExcludingTaxesAmountCents returns the SubTotalExcludingTaxesAmountCents field value
func (o *InvoiceObjectExtended) GetSubTotalExcludingTaxesAmountCents() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SubTotalExcludingTaxesAmountCents
}

// GetSubTotalExcludingTaxesAmountCentsOk returns a tuple with the SubTotalExcludingTaxesAmountCents field value
// and a boolean to check if the value has been set.
func (o *InvoiceObjectExtended) GetSubTotalExcludingTaxesAmountCentsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubTotalExcludingTaxesAmountCents, true
}

// SetSubTotalExcludingTaxesAmountCents sets field value
func (o *InvoiceObjectExtended) SetSubTotalExcludingTaxesAmountCents(v int32) {
	o.SubTotalExcludingTaxesAmountCents = v
}

// GetTaxesAmountCents returns the TaxesAmountCents field value
func (o *InvoiceObjectExtended) GetTaxesAmountCents() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TaxesAmountCents
}

// GetTaxesAmountCentsOk returns a tuple with the TaxesAmountCents field value
// and a boolean to check if the value has been set.
func (o *InvoiceObjectExtended) GetTaxesAmountCentsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaxesAmountCents, true
}

// SetTaxesAmountCents sets field value
func (o *InvoiceObjectExtended) SetTaxesAmountCents(v int32) {
	o.TaxesAmountCents = v
}

// GetSubTotalIncludingTaxesAmountCents returns the SubTotalIncludingTaxesAmountCents field value
func (o *InvoiceObjectExtended) GetSubTotalIncludingTaxesAmountCents() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SubTotalIncludingTaxesAmountCents
}

// GetSubTotalIncludingTaxesAmountCentsOk returns a tuple with the SubTotalIncludingTaxesAmountCents field value
// and a boolean to check if the value has been set.
func (o *InvoiceObjectExtended) GetSubTotalIncludingTaxesAmountCentsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubTotalIncludingTaxesAmountCents, true
}

// SetSubTotalIncludingTaxesAmountCents sets field value
func (o *InvoiceObjectExtended) SetSubTotalIncludingTaxesAmountCents(v int32) {
	o.SubTotalIncludingTaxesAmountCents = v
}

// GetPrepaidCreditAmountCents returns the PrepaidCreditAmountCents field value
func (o *InvoiceObjectExtended) GetPrepaidCreditAmountCents() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PrepaidCreditAmountCents
}

// GetPrepaidCreditAmountCentsOk returns a tuple with the PrepaidCreditAmountCents field value
// and a boolean to check if the value has been set.
func (o *InvoiceObjectExtended) GetPrepaidCreditAmountCentsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrepaidCreditAmountCents, true
}

// SetPrepaidCreditAmountCents sets field value
func (o *InvoiceObjectExtended) SetPrepaidCreditAmountCents(v int32) {
	o.PrepaidCreditAmountCents = v
}

// GetTotalAmountCents returns the TotalAmountCents field value
func (o *InvoiceObjectExtended) GetTotalAmountCents() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalAmountCents
}

// GetTotalAmountCentsOk returns a tuple with the TotalAmountCents field value
// and a boolean to check if the value has been set.
func (o *InvoiceObjectExtended) GetTotalAmountCentsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalAmountCents, true
}

// SetTotalAmountCents sets field value
func (o *InvoiceObjectExtended) SetTotalAmountCents(v int32) {
	o.TotalAmountCents = v
}

// GetVersionNumber returns the VersionNumber field value
func (o *InvoiceObjectExtended) GetVersionNumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VersionNumber
}

// GetVersionNumberOk returns a tuple with the VersionNumber field value
// and a boolean to check if the value has been set.
func (o *InvoiceObjectExtended) GetVersionNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VersionNumber, true
}

// SetVersionNumber sets field value
func (o *InvoiceObjectExtended) SetVersionNumber(v int32) {
	o.VersionNumber = v
}

// GetFileUrl returns the FileUrl field value if set, zero value otherwise.
func (o *InvoiceObjectExtended) GetFileUrl() string {
	if o == nil || IsNil(o.FileUrl) {
		var ret string
		return ret
	}
	return *o.FileUrl
}

// GetFileUrlOk returns a tuple with the FileUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceObjectExtended) GetFileUrlOk() (*string, bool) {
	if o == nil || IsNil(o.FileUrl) {
		return nil, false
	}
	return o.FileUrl, true
}

// HasFileUrl returns a boolean if a field has been set.
func (o *InvoiceObjectExtended) HasFileUrl() bool {
	if o != nil && !IsNil(o.FileUrl) {
		return true
	}

	return false
}

// SetFileUrl gets a reference to the given string and assigns it to the FileUrl field.
func (o *InvoiceObjectExtended) SetFileUrl(v string) {
	o.FileUrl = &v
}

// GetCustomer returns the Customer field value if set, zero value otherwise.
func (o *InvoiceObjectExtended) GetCustomer() InvoiceObjectCustomer {
	if o == nil || IsNil(o.Customer) {
		var ret InvoiceObjectCustomer
		return ret
	}
	return *o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceObjectExtended) GetCustomerOk() (*InvoiceObjectCustomer, bool) {
	if o == nil || IsNil(o.Customer) {
		return nil, false
	}
	return o.Customer, true
}

// HasCustomer returns a boolean if a field has been set.
func (o *InvoiceObjectExtended) HasCustomer() bool {
	if o != nil && !IsNil(o.Customer) {
		return true
	}

	return false
}

// SetCustomer gets a reference to the given InvoiceObjectCustomer and assigns it to the Customer field.
func (o *InvoiceObjectExtended) SetCustomer(v InvoiceObjectCustomer) {
	o.Customer = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *InvoiceObjectExtended) GetMetadata() []InvoiceMetadataObject {
	if o == nil || IsNil(o.Metadata) {
		var ret []InvoiceMetadataObject
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceObjectExtended) GetMetadataOk() ([]InvoiceMetadataObject, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *InvoiceObjectExtended) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given []InvoiceMetadataObject and assigns it to the Metadata field.
func (o *InvoiceObjectExtended) SetMetadata(v []InvoiceMetadataObject) {
	o.Metadata = v
}

// GetAppliedTaxes returns the AppliedTaxes field value if set, zero value otherwise.
func (o *InvoiceObjectExtended) GetAppliedTaxes() []InvoiceAppliedTaxObject {
	if o == nil || IsNil(o.AppliedTaxes) {
		var ret []InvoiceAppliedTaxObject
		return ret
	}
	return o.AppliedTaxes
}

// GetAppliedTaxesOk returns a tuple with the AppliedTaxes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceObjectExtended) GetAppliedTaxesOk() ([]InvoiceAppliedTaxObject, bool) {
	if o == nil || IsNil(o.AppliedTaxes) {
		return nil, false
	}
	return o.AppliedTaxes, true
}

// HasAppliedTaxes returns a boolean if a field has been set.
func (o *InvoiceObjectExtended) HasAppliedTaxes() bool {
	if o != nil && !IsNil(o.AppliedTaxes) {
		return true
	}

	return false
}

// SetAppliedTaxes gets a reference to the given []InvoiceAppliedTaxObject and assigns it to the AppliedTaxes field.
func (o *InvoiceObjectExtended) SetAppliedTaxes(v []InvoiceAppliedTaxObject) {
	o.AppliedTaxes = v
}

// GetCredits returns the Credits field value if set, zero value otherwise.
func (o *InvoiceObjectExtended) GetCredits() []CreditObject {
	if o == nil || IsNil(o.Credits) {
		var ret []CreditObject
		return ret
	}
	return o.Credits
}

// GetCreditsOk returns a tuple with the Credits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceObjectExtended) GetCreditsOk() ([]CreditObject, bool) {
	if o == nil || IsNil(o.Credits) {
		return nil, false
	}
	return o.Credits, true
}

// HasCredits returns a boolean if a field has been set.
func (o *InvoiceObjectExtended) HasCredits() bool {
	if o != nil && !IsNil(o.Credits) {
		return true
	}

	return false
}

// SetCredits gets a reference to the given []CreditObject and assigns it to the Credits field.
func (o *InvoiceObjectExtended) SetCredits(v []CreditObject) {
	o.Credits = v
}

// GetFees returns the Fees field value if set, zero value otherwise.
func (o *InvoiceObjectExtended) GetFees() []FeeObject {
	if o == nil || IsNil(o.Fees) {
		var ret []FeeObject
		return ret
	}
	return o.Fees
}

// GetFeesOk returns a tuple with the Fees field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceObjectExtended) GetFeesOk() ([]FeeObject, bool) {
	if o == nil || IsNil(o.Fees) {
		return nil, false
	}
	return o.Fees, true
}

// HasFees returns a boolean if a field has been set.
func (o *InvoiceObjectExtended) HasFees() bool {
	if o != nil && !IsNil(o.Fees) {
		return true
	}

	return false
}

// SetFees gets a reference to the given []FeeObject and assigns it to the Fees field.
func (o *InvoiceObjectExtended) SetFees(v []FeeObject) {
	o.Fees = v
}

// GetSubscriptions returns the Subscriptions field value if set, zero value otherwise.
func (o *InvoiceObjectExtended) GetSubscriptions() []SubscriptionObject {
	if o == nil || IsNil(o.Subscriptions) {
		var ret []SubscriptionObject
		return ret
	}
	return o.Subscriptions
}

// GetSubscriptionsOk returns a tuple with the Subscriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceObjectExtended) GetSubscriptionsOk() ([]SubscriptionObject, bool) {
	if o == nil || IsNil(o.Subscriptions) {
		return nil, false
	}
	return o.Subscriptions, true
}

// HasSubscriptions returns a boolean if a field has been set.
func (o *InvoiceObjectExtended) HasSubscriptions() bool {
	if o != nil && !IsNil(o.Subscriptions) {
		return true
	}

	return false
}

// SetSubscriptions gets a reference to the given []SubscriptionObject and assigns it to the Subscriptions field.
func (o *InvoiceObjectExtended) SetSubscriptions(v []SubscriptionObject) {
	o.Subscriptions = v
}

func (o InvoiceObjectExtended) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvoiceObjectExtended) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["lago_id"] = o.LagoId
	toSerialize["sequential_id"] = o.SequentialId
	toSerialize["number"] = o.Number
	toSerialize["issuing_date"] = o.IssuingDate
	if !IsNil(o.PaymentDueDate) {
		toSerialize["payment_due_date"] = o.PaymentDueDate
	}
	if !IsNil(o.NetPaymentTerm) {
		toSerialize["net_payment_term"] = o.NetPaymentTerm
	}
	toSerialize["invoice_type"] = o.InvoiceType
	toSerialize["status"] = o.Status
	toSerialize["payment_status"] = o.PaymentStatus
	toSerialize["currency"] = o.Currency
	toSerialize["fees_amount_cents"] = o.FeesAmountCents
	toSerialize["coupons_amount_cents"] = o.CouponsAmountCents
	toSerialize["credit_notes_amount_cents"] = o.CreditNotesAmountCents
	toSerialize["sub_total_excluding_taxes_amount_cents"] = o.SubTotalExcludingTaxesAmountCents
	toSerialize["taxes_amount_cents"] = o.TaxesAmountCents
	toSerialize["sub_total_including_taxes_amount_cents"] = o.SubTotalIncludingTaxesAmountCents
	toSerialize["prepaid_credit_amount_cents"] = o.PrepaidCreditAmountCents
	toSerialize["total_amount_cents"] = o.TotalAmountCents
	toSerialize["version_number"] = o.VersionNumber
	if !IsNil(o.FileUrl) {
		toSerialize["file_url"] = o.FileUrl
	}
	if !IsNil(o.Customer) {
		toSerialize["customer"] = o.Customer
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.AppliedTaxes) {
		toSerialize["applied_taxes"] = o.AppliedTaxes
	}
	if !IsNil(o.Credits) {
		toSerialize["credits"] = o.Credits
	}
	if !IsNil(o.Fees) {
		toSerialize["fees"] = o.Fees
	}
	if !IsNil(o.Subscriptions) {
		toSerialize["subscriptions"] = o.Subscriptions
	}
	return toSerialize, nil
}

func (o *InvoiceObjectExtended) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"lago_id",
		"sequential_id",
		"number",
		"issuing_date",
		"invoice_type",
		"status",
		"payment_status",
		"currency",
		"fees_amount_cents",
		"coupons_amount_cents",
		"credit_notes_amount_cents",
		"sub_total_excluding_taxes_amount_cents",
		"taxes_amount_cents",
		"sub_total_including_taxes_amount_cents",
		"prepaid_credit_amount_cents",
		"total_amount_cents",
		"version_number",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varInvoiceObjectExtended := _InvoiceObjectExtended{}

	err = json.Unmarshal(bytes, &varInvoiceObjectExtended)

	if err != nil {
		return err
	}

	*o = InvoiceObjectExtended(varInvoiceObjectExtended)

	return err
}

type NullableInvoiceObjectExtended struct {
	value *InvoiceObjectExtended
	isSet bool
}

func (v NullableInvoiceObjectExtended) Get() *InvoiceObjectExtended {
	return v.value
}

func (v *NullableInvoiceObjectExtended) Set(val *InvoiceObjectExtended) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceObjectExtended) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceObjectExtended) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceObjectExtended(val *InvoiceObjectExtended) *NullableInvoiceObjectExtended {
	return &NullableInvoiceObjectExtended{value: val, isSet: true}
}

func (v NullableInvoiceObjectExtended) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceObjectExtended) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

