# OrganizationBillingConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvoiceFooter** | Pointer to **NullableString** | The customer invoice message that appears at the bottom of each billing documents. | [optional] 
**InvoiceGracePeriod** | Pointer to **int32** | The grace period, expressed in days, for finalizing the invoice. This period refers to the additional time granted to your customers beyond the invoice due date to adjust usage and line items. Can be overwritten by the customer’s grace period. | [optional] 
**DocumentLocale** | Pointer to **string** | The locale of the billing documents, expressed in the ISO 639-1 format. This field indicates the language or regional variant used for the documents content issued or the embeddable customer portal. | [optional] 

## Methods

### NewOrganizationBillingConfiguration

`func NewOrganizationBillingConfiguration() *OrganizationBillingConfiguration`

NewOrganizationBillingConfiguration instantiates a new OrganizationBillingConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationBillingConfigurationWithDefaults

`func NewOrganizationBillingConfigurationWithDefaults() *OrganizationBillingConfiguration`

NewOrganizationBillingConfigurationWithDefaults instantiates a new OrganizationBillingConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoiceFooter

`func (o *OrganizationBillingConfiguration) GetInvoiceFooter() string`

GetInvoiceFooter returns the InvoiceFooter field if non-nil, zero value otherwise.

### GetInvoiceFooterOk

`func (o *OrganizationBillingConfiguration) GetInvoiceFooterOk() (*string, bool)`

GetInvoiceFooterOk returns a tuple with the InvoiceFooter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceFooter

`func (o *OrganizationBillingConfiguration) SetInvoiceFooter(v string)`

SetInvoiceFooter sets InvoiceFooter field to given value.

### HasInvoiceFooter

`func (o *OrganizationBillingConfiguration) HasInvoiceFooter() bool`

HasInvoiceFooter returns a boolean if a field has been set.

### SetInvoiceFooterNil

`func (o *OrganizationBillingConfiguration) SetInvoiceFooterNil(b bool)`

 SetInvoiceFooterNil sets the value for InvoiceFooter to be an explicit nil

### UnsetInvoiceFooter
`func (o *OrganizationBillingConfiguration) UnsetInvoiceFooter()`

UnsetInvoiceFooter ensures that no value is present for InvoiceFooter, not even an explicit nil
### GetInvoiceGracePeriod

`func (o *OrganizationBillingConfiguration) GetInvoiceGracePeriod() int32`

GetInvoiceGracePeriod returns the InvoiceGracePeriod field if non-nil, zero value otherwise.

### GetInvoiceGracePeriodOk

`func (o *OrganizationBillingConfiguration) GetInvoiceGracePeriodOk() (*int32, bool)`

GetInvoiceGracePeriodOk returns a tuple with the InvoiceGracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceGracePeriod

`func (o *OrganizationBillingConfiguration) SetInvoiceGracePeriod(v int32)`

SetInvoiceGracePeriod sets InvoiceGracePeriod field to given value.

### HasInvoiceGracePeriod

`func (o *OrganizationBillingConfiguration) HasInvoiceGracePeriod() bool`

HasInvoiceGracePeriod returns a boolean if a field has been set.

### GetDocumentLocale

`func (o *OrganizationBillingConfiguration) GetDocumentLocale() string`

GetDocumentLocale returns the DocumentLocale field if non-nil, zero value otherwise.

### GetDocumentLocaleOk

`func (o *OrganizationBillingConfiguration) GetDocumentLocaleOk() (*string, bool)`

GetDocumentLocaleOk returns a tuple with the DocumentLocale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentLocale

`func (o *OrganizationBillingConfiguration) SetDocumentLocale(v string)`

SetDocumentLocale sets DocumentLocale field to given value.

### HasDocumentLocale

`func (o *OrganizationBillingConfiguration) HasDocumentLocale() bool`

HasDocumentLocale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


