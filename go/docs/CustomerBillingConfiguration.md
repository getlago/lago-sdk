# CustomerBillingConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvoiceGracePeriod** | Pointer to **int32** | The grace period, expressed in days, for the invoice. This period refers to the additional time granted to the customer beyond the invoice due date to adjust usage and line items | [optional] 
**PaymentProvider** | Pointer to **string** | The payment provider utilized to initiate payments for invoices issued by Lago. Accepted values: &#x60;stripe&#x60;, &#x60;adyen&#x60;, &#x60;gocardless&#x60; or null. This field is required if you intend to assign a &#x60;provider_customer_id&#x60;. | [optional] 
**ProviderCustomerId** | Pointer to **string** | The customer ID within the payment provider&#39;s system. If this field is not provided, Lago has the option to create a new customer record within the payment provider&#39;s system on behalf of the customer | [optional] 
**Sync** | Pointer to **bool** | Set this field to &#x60;true&#x60; if you want to create the customer in the payment provider synchronously with the customer creation process in Lago. This option is applicable only when the &#x60;provider_customer_id&#x60; is &#x60;null&#x60; and the customer is automatically created in the payment provider through Lago. By default, the value is set to &#x60;false&#x60; | [optional] 
**SyncWithProvider** | Pointer to **bool** | Set this field to &#x60;true&#x60; if you want to create a customer record in the payment provider&#39;s system. This option is applicable only when the &#x60;provider_customer_id&#x60; is null and the &#x60;sync_with_provider&#x60; field is set to &#x60;true&#x60;. By default, the value is set to &#x60;false&#x60; | [optional] 
**DocumentLocale** | Pointer to **string** | The document locale, specified in the ISO 639-1 format. This field represents the language or locale used for the documents issued by Lago | [optional] 
**ProviderPaymentMethods** | Pointer to **[]string** | Specifies the available payment methods that can be used for this customer when &#x60;payment_provider&#x60; is set to &#x60;stripe&#x60;. The &#x60;provider_payment_methods&#x60; field is an array that allows multiple payment options to be defined. If this field is not explicitly set, all the payment methods are selected. For now, possible values are &#x60;card&#x60; and &#x60;sepa_debit&#x60;. | [optional] 

## Methods

### NewCustomerBillingConfiguration

`func NewCustomerBillingConfiguration() *CustomerBillingConfiguration`

NewCustomerBillingConfiguration instantiates a new CustomerBillingConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerBillingConfigurationWithDefaults

`func NewCustomerBillingConfigurationWithDefaults() *CustomerBillingConfiguration`

NewCustomerBillingConfigurationWithDefaults instantiates a new CustomerBillingConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoiceGracePeriod

`func (o *CustomerBillingConfiguration) GetInvoiceGracePeriod() int32`

GetInvoiceGracePeriod returns the InvoiceGracePeriod field if non-nil, zero value otherwise.

### GetInvoiceGracePeriodOk

`func (o *CustomerBillingConfiguration) GetInvoiceGracePeriodOk() (*int32, bool)`

GetInvoiceGracePeriodOk returns a tuple with the InvoiceGracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceGracePeriod

`func (o *CustomerBillingConfiguration) SetInvoiceGracePeriod(v int32)`

SetInvoiceGracePeriod sets InvoiceGracePeriod field to given value.

### HasInvoiceGracePeriod

`func (o *CustomerBillingConfiguration) HasInvoiceGracePeriod() bool`

HasInvoiceGracePeriod returns a boolean if a field has been set.

### GetPaymentProvider

`func (o *CustomerBillingConfiguration) GetPaymentProvider() string`

GetPaymentProvider returns the PaymentProvider field if non-nil, zero value otherwise.

### GetPaymentProviderOk

`func (o *CustomerBillingConfiguration) GetPaymentProviderOk() (*string, bool)`

GetPaymentProviderOk returns a tuple with the PaymentProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentProvider

`func (o *CustomerBillingConfiguration) SetPaymentProvider(v string)`

SetPaymentProvider sets PaymentProvider field to given value.

### HasPaymentProvider

`func (o *CustomerBillingConfiguration) HasPaymentProvider() bool`

HasPaymentProvider returns a boolean if a field has been set.

### GetProviderCustomerId

`func (o *CustomerBillingConfiguration) GetProviderCustomerId() string`

GetProviderCustomerId returns the ProviderCustomerId field if non-nil, zero value otherwise.

### GetProviderCustomerIdOk

`func (o *CustomerBillingConfiguration) GetProviderCustomerIdOk() (*string, bool)`

GetProviderCustomerIdOk returns a tuple with the ProviderCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderCustomerId

`func (o *CustomerBillingConfiguration) SetProviderCustomerId(v string)`

SetProviderCustomerId sets ProviderCustomerId field to given value.

### HasProviderCustomerId

`func (o *CustomerBillingConfiguration) HasProviderCustomerId() bool`

HasProviderCustomerId returns a boolean if a field has been set.

### GetSync

`func (o *CustomerBillingConfiguration) GetSync() bool`

GetSync returns the Sync field if non-nil, zero value otherwise.

### GetSyncOk

`func (o *CustomerBillingConfiguration) GetSyncOk() (*bool, bool)`

GetSyncOk returns a tuple with the Sync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSync

`func (o *CustomerBillingConfiguration) SetSync(v bool)`

SetSync sets Sync field to given value.

### HasSync

`func (o *CustomerBillingConfiguration) HasSync() bool`

HasSync returns a boolean if a field has been set.

### GetSyncWithProvider

`func (o *CustomerBillingConfiguration) GetSyncWithProvider() bool`

GetSyncWithProvider returns the SyncWithProvider field if non-nil, zero value otherwise.

### GetSyncWithProviderOk

`func (o *CustomerBillingConfiguration) GetSyncWithProviderOk() (*bool, bool)`

GetSyncWithProviderOk returns a tuple with the SyncWithProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncWithProvider

`func (o *CustomerBillingConfiguration) SetSyncWithProvider(v bool)`

SetSyncWithProvider sets SyncWithProvider field to given value.

### HasSyncWithProvider

`func (o *CustomerBillingConfiguration) HasSyncWithProvider() bool`

HasSyncWithProvider returns a boolean if a field has been set.

### GetDocumentLocale

`func (o *CustomerBillingConfiguration) GetDocumentLocale() string`

GetDocumentLocale returns the DocumentLocale field if non-nil, zero value otherwise.

### GetDocumentLocaleOk

`func (o *CustomerBillingConfiguration) GetDocumentLocaleOk() (*string, bool)`

GetDocumentLocaleOk returns a tuple with the DocumentLocale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentLocale

`func (o *CustomerBillingConfiguration) SetDocumentLocale(v string)`

SetDocumentLocale sets DocumentLocale field to given value.

### HasDocumentLocale

`func (o *CustomerBillingConfiguration) HasDocumentLocale() bool`

HasDocumentLocale returns a boolean if a field has been set.

### GetProviderPaymentMethods

`func (o *CustomerBillingConfiguration) GetProviderPaymentMethods() []string`

GetProviderPaymentMethods returns the ProviderPaymentMethods field if non-nil, zero value otherwise.

### GetProviderPaymentMethodsOk

`func (o *CustomerBillingConfiguration) GetProviderPaymentMethodsOk() (*[]string, bool)`

GetProviderPaymentMethodsOk returns a tuple with the ProviderPaymentMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderPaymentMethods

`func (o *CustomerBillingConfiguration) SetProviderPaymentMethods(v []string)`

SetProviderPaymentMethods sets ProviderPaymentMethods field to given value.

### HasProviderPaymentMethods

`func (o *CustomerBillingConfiguration) HasProviderPaymentMethods() bool`

HasProviderPaymentMethods returns a boolean if a field has been set.

### SetProviderPaymentMethodsNil

`func (o *CustomerBillingConfiguration) SetProviderPaymentMethodsNil(b bool)`

 SetProviderPaymentMethodsNil sets the value for ProviderPaymentMethods to be an explicit nil

### UnsetProviderPaymentMethods
`func (o *CustomerBillingConfiguration) UnsetProviderPaymentMethods()`

UnsetProviderPaymentMethods ensures that no value is present for ProviderPaymentMethods, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


