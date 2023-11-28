# OrganizationUpdateInputOrganization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebhookUrl** | Pointer to **NullableString** | The URL of your newest updated webhook endpoint. This URL allows your organization to receive important messages, notifications, or data from the Lago system. By configuring your webhook endpoint to this URL, you can ensure that your organization stays informed and receives relevant information in a timely manner. | [optional] 
**Country** | Pointer to [**Country**](Country.md) |  | [optional] 
**DefaultCurrency** | Pointer to [**Currency**](Currency.md) |  | [optional] 
**AddressLine1** | Pointer to **NullableString** | The first line of your organization’s billing address. | [optional] 
**AddressLine2** | Pointer to **NullableString** | The second line of your organization’s billing address. | [optional] 
**State** | Pointer to **NullableString** | The state of your organization’s billing address. | [optional] 
**Zipcode** | Pointer to **NullableString** | The zipcode of your organization’s billing address. | [optional] 
**Email** | Pointer to **NullableString** | The email address of your organization used to bill your customers. | [optional] 
**City** | Pointer to **NullableString** | The city of your organization’s billing address. | [optional] 
**LegalName** | Pointer to **NullableString** | The legal name of your organization. | [optional] 
**LegalNumber** | Pointer to **NullableString** | The legal number of your organization. | [optional] 
**NetPaymentTerm** | Pointer to **int32** | The net payment term, expressed in days, specifies the duration within which a customer is expected to remit payment after the invoice is finalized. | [optional] 
**TaxIdentificationNumber** | Pointer to **NullableString** | The tax identification number of your organization. | [optional] 
**Timezone** | Pointer to [**Timezone**](Timezone.md) |  | [optional] 
**EmailSettings** | Pointer to **[]string** | Represents the email settings of the organization. It allows you to define which documents are sent by email. The field value determines the types of documents that trigger email notifications. Possible values for are &#x60;invoice.finalized&#x60; and &#x60;credit_note.created&#x60;. By configuring this field, you can specify whether invoices, credit notes, or both should be sent to recipients via email. | [optional] 
**BillingConfiguration** | Pointer to [**OrganizationBillingConfiguration**](OrganizationBillingConfiguration.md) |  | [optional] 

## Methods

### NewOrganizationUpdateInputOrganization

`func NewOrganizationUpdateInputOrganization() *OrganizationUpdateInputOrganization`

NewOrganizationUpdateInputOrganization instantiates a new OrganizationUpdateInputOrganization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationUpdateInputOrganizationWithDefaults

`func NewOrganizationUpdateInputOrganizationWithDefaults() *OrganizationUpdateInputOrganization`

NewOrganizationUpdateInputOrganizationWithDefaults instantiates a new OrganizationUpdateInputOrganization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebhookUrl

`func (o *OrganizationUpdateInputOrganization) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *OrganizationUpdateInputOrganization) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *OrganizationUpdateInputOrganization) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *OrganizationUpdateInputOrganization) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### SetWebhookUrlNil

`func (o *OrganizationUpdateInputOrganization) SetWebhookUrlNil(b bool)`

 SetWebhookUrlNil sets the value for WebhookUrl to be an explicit nil

### UnsetWebhookUrl
`func (o *OrganizationUpdateInputOrganization) UnsetWebhookUrl()`

UnsetWebhookUrl ensures that no value is present for WebhookUrl, not even an explicit nil
### GetCountry

`func (o *OrganizationUpdateInputOrganization) GetCountry() Country`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *OrganizationUpdateInputOrganization) GetCountryOk() (*Country, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *OrganizationUpdateInputOrganization) SetCountry(v Country)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *OrganizationUpdateInputOrganization) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetDefaultCurrency

`func (o *OrganizationUpdateInputOrganization) GetDefaultCurrency() Currency`

GetDefaultCurrency returns the DefaultCurrency field if non-nil, zero value otherwise.

### GetDefaultCurrencyOk

`func (o *OrganizationUpdateInputOrganization) GetDefaultCurrencyOk() (*Currency, bool)`

GetDefaultCurrencyOk returns a tuple with the DefaultCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCurrency

`func (o *OrganizationUpdateInputOrganization) SetDefaultCurrency(v Currency)`

SetDefaultCurrency sets DefaultCurrency field to given value.

### HasDefaultCurrency

`func (o *OrganizationUpdateInputOrganization) HasDefaultCurrency() bool`

HasDefaultCurrency returns a boolean if a field has been set.

### GetAddressLine1

`func (o *OrganizationUpdateInputOrganization) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *OrganizationUpdateInputOrganization) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *OrganizationUpdateInputOrganization) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *OrganizationUpdateInputOrganization) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### SetAddressLine1Nil

`func (o *OrganizationUpdateInputOrganization) SetAddressLine1Nil(b bool)`

 SetAddressLine1Nil sets the value for AddressLine1 to be an explicit nil

### UnsetAddressLine1
`func (o *OrganizationUpdateInputOrganization) UnsetAddressLine1()`

UnsetAddressLine1 ensures that no value is present for AddressLine1, not even an explicit nil
### GetAddressLine2

`func (o *OrganizationUpdateInputOrganization) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *OrganizationUpdateInputOrganization) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *OrganizationUpdateInputOrganization) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *OrganizationUpdateInputOrganization) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### SetAddressLine2Nil

`func (o *OrganizationUpdateInputOrganization) SetAddressLine2Nil(b bool)`

 SetAddressLine2Nil sets the value for AddressLine2 to be an explicit nil

### UnsetAddressLine2
`func (o *OrganizationUpdateInputOrganization) UnsetAddressLine2()`

UnsetAddressLine2 ensures that no value is present for AddressLine2, not even an explicit nil
### GetState

`func (o *OrganizationUpdateInputOrganization) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *OrganizationUpdateInputOrganization) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *OrganizationUpdateInputOrganization) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *OrganizationUpdateInputOrganization) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *OrganizationUpdateInputOrganization) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *OrganizationUpdateInputOrganization) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetZipcode

`func (o *OrganizationUpdateInputOrganization) GetZipcode() string`

GetZipcode returns the Zipcode field if non-nil, zero value otherwise.

### GetZipcodeOk

`func (o *OrganizationUpdateInputOrganization) GetZipcodeOk() (*string, bool)`

GetZipcodeOk returns a tuple with the Zipcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipcode

`func (o *OrganizationUpdateInputOrganization) SetZipcode(v string)`

SetZipcode sets Zipcode field to given value.

### HasZipcode

`func (o *OrganizationUpdateInputOrganization) HasZipcode() bool`

HasZipcode returns a boolean if a field has been set.

### SetZipcodeNil

`func (o *OrganizationUpdateInputOrganization) SetZipcodeNil(b bool)`

 SetZipcodeNil sets the value for Zipcode to be an explicit nil

### UnsetZipcode
`func (o *OrganizationUpdateInputOrganization) UnsetZipcode()`

UnsetZipcode ensures that no value is present for Zipcode, not even an explicit nil
### GetEmail

`func (o *OrganizationUpdateInputOrganization) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OrganizationUpdateInputOrganization) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OrganizationUpdateInputOrganization) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *OrganizationUpdateInputOrganization) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *OrganizationUpdateInputOrganization) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *OrganizationUpdateInputOrganization) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetCity

`func (o *OrganizationUpdateInputOrganization) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *OrganizationUpdateInputOrganization) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *OrganizationUpdateInputOrganization) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *OrganizationUpdateInputOrganization) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *OrganizationUpdateInputOrganization) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *OrganizationUpdateInputOrganization) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetLegalName

`func (o *OrganizationUpdateInputOrganization) GetLegalName() string`

GetLegalName returns the LegalName field if non-nil, zero value otherwise.

### GetLegalNameOk

`func (o *OrganizationUpdateInputOrganization) GetLegalNameOk() (*string, bool)`

GetLegalNameOk returns a tuple with the LegalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalName

`func (o *OrganizationUpdateInputOrganization) SetLegalName(v string)`

SetLegalName sets LegalName field to given value.

### HasLegalName

`func (o *OrganizationUpdateInputOrganization) HasLegalName() bool`

HasLegalName returns a boolean if a field has been set.

### SetLegalNameNil

`func (o *OrganizationUpdateInputOrganization) SetLegalNameNil(b bool)`

 SetLegalNameNil sets the value for LegalName to be an explicit nil

### UnsetLegalName
`func (o *OrganizationUpdateInputOrganization) UnsetLegalName()`

UnsetLegalName ensures that no value is present for LegalName, not even an explicit nil
### GetLegalNumber

`func (o *OrganizationUpdateInputOrganization) GetLegalNumber() string`

GetLegalNumber returns the LegalNumber field if non-nil, zero value otherwise.

### GetLegalNumberOk

`func (o *OrganizationUpdateInputOrganization) GetLegalNumberOk() (*string, bool)`

GetLegalNumberOk returns a tuple with the LegalNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalNumber

`func (o *OrganizationUpdateInputOrganization) SetLegalNumber(v string)`

SetLegalNumber sets LegalNumber field to given value.

### HasLegalNumber

`func (o *OrganizationUpdateInputOrganization) HasLegalNumber() bool`

HasLegalNumber returns a boolean if a field has been set.

### SetLegalNumberNil

`func (o *OrganizationUpdateInputOrganization) SetLegalNumberNil(b bool)`

 SetLegalNumberNil sets the value for LegalNumber to be an explicit nil

### UnsetLegalNumber
`func (o *OrganizationUpdateInputOrganization) UnsetLegalNumber()`

UnsetLegalNumber ensures that no value is present for LegalNumber, not even an explicit nil
### GetNetPaymentTerm

`func (o *OrganizationUpdateInputOrganization) GetNetPaymentTerm() int32`

GetNetPaymentTerm returns the NetPaymentTerm field if non-nil, zero value otherwise.

### GetNetPaymentTermOk

`func (o *OrganizationUpdateInputOrganization) GetNetPaymentTermOk() (*int32, bool)`

GetNetPaymentTermOk returns a tuple with the NetPaymentTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetPaymentTerm

`func (o *OrganizationUpdateInputOrganization) SetNetPaymentTerm(v int32)`

SetNetPaymentTerm sets NetPaymentTerm field to given value.

### HasNetPaymentTerm

`func (o *OrganizationUpdateInputOrganization) HasNetPaymentTerm() bool`

HasNetPaymentTerm returns a boolean if a field has been set.

### GetTaxIdentificationNumber

`func (o *OrganizationUpdateInputOrganization) GetTaxIdentificationNumber() string`

GetTaxIdentificationNumber returns the TaxIdentificationNumber field if non-nil, zero value otherwise.

### GetTaxIdentificationNumberOk

`func (o *OrganizationUpdateInputOrganization) GetTaxIdentificationNumberOk() (*string, bool)`

GetTaxIdentificationNumberOk returns a tuple with the TaxIdentificationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxIdentificationNumber

`func (o *OrganizationUpdateInputOrganization) SetTaxIdentificationNumber(v string)`

SetTaxIdentificationNumber sets TaxIdentificationNumber field to given value.

### HasTaxIdentificationNumber

`func (o *OrganizationUpdateInputOrganization) HasTaxIdentificationNumber() bool`

HasTaxIdentificationNumber returns a boolean if a field has been set.

### SetTaxIdentificationNumberNil

`func (o *OrganizationUpdateInputOrganization) SetTaxIdentificationNumberNil(b bool)`

 SetTaxIdentificationNumberNil sets the value for TaxIdentificationNumber to be an explicit nil

### UnsetTaxIdentificationNumber
`func (o *OrganizationUpdateInputOrganization) UnsetTaxIdentificationNumber()`

UnsetTaxIdentificationNumber ensures that no value is present for TaxIdentificationNumber, not even an explicit nil
### GetTimezone

`func (o *OrganizationUpdateInputOrganization) GetTimezone() Timezone`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *OrganizationUpdateInputOrganization) GetTimezoneOk() (*Timezone, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *OrganizationUpdateInputOrganization) SetTimezone(v Timezone)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *OrganizationUpdateInputOrganization) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetEmailSettings

`func (o *OrganizationUpdateInputOrganization) GetEmailSettings() []string`

GetEmailSettings returns the EmailSettings field if non-nil, zero value otherwise.

### GetEmailSettingsOk

`func (o *OrganizationUpdateInputOrganization) GetEmailSettingsOk() (*[]string, bool)`

GetEmailSettingsOk returns a tuple with the EmailSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailSettings

`func (o *OrganizationUpdateInputOrganization) SetEmailSettings(v []string)`

SetEmailSettings sets EmailSettings field to given value.

### HasEmailSettings

`func (o *OrganizationUpdateInputOrganization) HasEmailSettings() bool`

HasEmailSettings returns a boolean if a field has been set.

### GetBillingConfiguration

`func (o *OrganizationUpdateInputOrganization) GetBillingConfiguration() OrganizationBillingConfiguration`

GetBillingConfiguration returns the BillingConfiguration field if non-nil, zero value otherwise.

### GetBillingConfigurationOk

`func (o *OrganizationUpdateInputOrganization) GetBillingConfigurationOk() (*OrganizationBillingConfiguration, bool)`

GetBillingConfigurationOk returns a tuple with the BillingConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingConfiguration

`func (o *OrganizationUpdateInputOrganization) SetBillingConfiguration(v OrganizationBillingConfiguration)`

SetBillingConfiguration sets BillingConfiguration field to given value.

### HasBillingConfiguration

`func (o *OrganizationUpdateInputOrganization) HasBillingConfiguration() bool`

HasBillingConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


