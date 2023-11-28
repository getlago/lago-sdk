# OrganizationObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LagoId** | **string** | Unique identifier assigned to the organization within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the organization&#39;s record within the Lago system | 
**Name** | **string** | The name of your organization. | 
**CreatedAt** | **NullableTime** | The date of creation of your organization, represented in ISO 8601 datetime format and expressed in Coordinated Universal Time (UTC). | 
**WebhookUrl** | Pointer to **NullableString** | The URL of your newest updated webhook endpoint. This URL allows your organization to receive important messages, notifications, or data from the Lago system. By configuring your webhook endpoint to this URL, you can ensure that your organization stays informed and receives relevant information in a timely manner. | [optional] 
**WebhookUrls** | Pointer to **[]string** | The array containing your webhooks URLs. | [optional] 
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
**BillingConfiguration** | [**OrganizationBillingConfiguration**](OrganizationBillingConfiguration.md) |  | 
**Taxes** | Pointer to [**[]TaxObject**](TaxObject.md) | List of default organization taxes | [optional] 

## Methods

### NewOrganizationObject

`func NewOrganizationObject(lagoId string, name string, createdAt NullableTime, billingConfiguration OrganizationBillingConfiguration, ) *OrganizationObject`

NewOrganizationObject instantiates a new OrganizationObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationObjectWithDefaults

`func NewOrganizationObjectWithDefaults() *OrganizationObject`

NewOrganizationObjectWithDefaults instantiates a new OrganizationObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLagoId

`func (o *OrganizationObject) GetLagoId() string`

GetLagoId returns the LagoId field if non-nil, zero value otherwise.

### GetLagoIdOk

`func (o *OrganizationObject) GetLagoIdOk() (*string, bool)`

GetLagoIdOk returns a tuple with the LagoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagoId

`func (o *OrganizationObject) SetLagoId(v string)`

SetLagoId sets LagoId field to given value.


### GetName

`func (o *OrganizationObject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationObject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationObject) SetName(v string)`

SetName sets Name field to given value.


### GetCreatedAt

`func (o *OrganizationObject) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrganizationObject) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrganizationObject) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *OrganizationObject) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *OrganizationObject) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetWebhookUrl

`func (o *OrganizationObject) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *OrganizationObject) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *OrganizationObject) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *OrganizationObject) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### SetWebhookUrlNil

`func (o *OrganizationObject) SetWebhookUrlNil(b bool)`

 SetWebhookUrlNil sets the value for WebhookUrl to be an explicit nil

### UnsetWebhookUrl
`func (o *OrganizationObject) UnsetWebhookUrl()`

UnsetWebhookUrl ensures that no value is present for WebhookUrl, not even an explicit nil
### GetWebhookUrls

`func (o *OrganizationObject) GetWebhookUrls() []string`

GetWebhookUrls returns the WebhookUrls field if non-nil, zero value otherwise.

### GetWebhookUrlsOk

`func (o *OrganizationObject) GetWebhookUrlsOk() (*[]string, bool)`

GetWebhookUrlsOk returns a tuple with the WebhookUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrls

`func (o *OrganizationObject) SetWebhookUrls(v []string)`

SetWebhookUrls sets WebhookUrls field to given value.

### HasWebhookUrls

`func (o *OrganizationObject) HasWebhookUrls() bool`

HasWebhookUrls returns a boolean if a field has been set.

### SetWebhookUrlsNil

`func (o *OrganizationObject) SetWebhookUrlsNil(b bool)`

 SetWebhookUrlsNil sets the value for WebhookUrls to be an explicit nil

### UnsetWebhookUrls
`func (o *OrganizationObject) UnsetWebhookUrls()`

UnsetWebhookUrls ensures that no value is present for WebhookUrls, not even an explicit nil
### GetCountry

`func (o *OrganizationObject) GetCountry() Country`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *OrganizationObject) GetCountryOk() (*Country, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *OrganizationObject) SetCountry(v Country)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *OrganizationObject) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetDefaultCurrency

`func (o *OrganizationObject) GetDefaultCurrency() Currency`

GetDefaultCurrency returns the DefaultCurrency field if non-nil, zero value otherwise.

### GetDefaultCurrencyOk

`func (o *OrganizationObject) GetDefaultCurrencyOk() (*Currency, bool)`

GetDefaultCurrencyOk returns a tuple with the DefaultCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCurrency

`func (o *OrganizationObject) SetDefaultCurrency(v Currency)`

SetDefaultCurrency sets DefaultCurrency field to given value.

### HasDefaultCurrency

`func (o *OrganizationObject) HasDefaultCurrency() bool`

HasDefaultCurrency returns a boolean if a field has been set.

### GetAddressLine1

`func (o *OrganizationObject) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *OrganizationObject) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *OrganizationObject) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *OrganizationObject) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### SetAddressLine1Nil

`func (o *OrganizationObject) SetAddressLine1Nil(b bool)`

 SetAddressLine1Nil sets the value for AddressLine1 to be an explicit nil

### UnsetAddressLine1
`func (o *OrganizationObject) UnsetAddressLine1()`

UnsetAddressLine1 ensures that no value is present for AddressLine1, not even an explicit nil
### GetAddressLine2

`func (o *OrganizationObject) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *OrganizationObject) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *OrganizationObject) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *OrganizationObject) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### SetAddressLine2Nil

`func (o *OrganizationObject) SetAddressLine2Nil(b bool)`

 SetAddressLine2Nil sets the value for AddressLine2 to be an explicit nil

### UnsetAddressLine2
`func (o *OrganizationObject) UnsetAddressLine2()`

UnsetAddressLine2 ensures that no value is present for AddressLine2, not even an explicit nil
### GetState

`func (o *OrganizationObject) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *OrganizationObject) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *OrganizationObject) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *OrganizationObject) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *OrganizationObject) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *OrganizationObject) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetZipcode

`func (o *OrganizationObject) GetZipcode() string`

GetZipcode returns the Zipcode field if non-nil, zero value otherwise.

### GetZipcodeOk

`func (o *OrganizationObject) GetZipcodeOk() (*string, bool)`

GetZipcodeOk returns a tuple with the Zipcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipcode

`func (o *OrganizationObject) SetZipcode(v string)`

SetZipcode sets Zipcode field to given value.

### HasZipcode

`func (o *OrganizationObject) HasZipcode() bool`

HasZipcode returns a boolean if a field has been set.

### SetZipcodeNil

`func (o *OrganizationObject) SetZipcodeNil(b bool)`

 SetZipcodeNil sets the value for Zipcode to be an explicit nil

### UnsetZipcode
`func (o *OrganizationObject) UnsetZipcode()`

UnsetZipcode ensures that no value is present for Zipcode, not even an explicit nil
### GetEmail

`func (o *OrganizationObject) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OrganizationObject) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OrganizationObject) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *OrganizationObject) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *OrganizationObject) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *OrganizationObject) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetCity

`func (o *OrganizationObject) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *OrganizationObject) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *OrganizationObject) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *OrganizationObject) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *OrganizationObject) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *OrganizationObject) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetLegalName

`func (o *OrganizationObject) GetLegalName() string`

GetLegalName returns the LegalName field if non-nil, zero value otherwise.

### GetLegalNameOk

`func (o *OrganizationObject) GetLegalNameOk() (*string, bool)`

GetLegalNameOk returns a tuple with the LegalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalName

`func (o *OrganizationObject) SetLegalName(v string)`

SetLegalName sets LegalName field to given value.

### HasLegalName

`func (o *OrganizationObject) HasLegalName() bool`

HasLegalName returns a boolean if a field has been set.

### SetLegalNameNil

`func (o *OrganizationObject) SetLegalNameNil(b bool)`

 SetLegalNameNil sets the value for LegalName to be an explicit nil

### UnsetLegalName
`func (o *OrganizationObject) UnsetLegalName()`

UnsetLegalName ensures that no value is present for LegalName, not even an explicit nil
### GetLegalNumber

`func (o *OrganizationObject) GetLegalNumber() string`

GetLegalNumber returns the LegalNumber field if non-nil, zero value otherwise.

### GetLegalNumberOk

`func (o *OrganizationObject) GetLegalNumberOk() (*string, bool)`

GetLegalNumberOk returns a tuple with the LegalNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalNumber

`func (o *OrganizationObject) SetLegalNumber(v string)`

SetLegalNumber sets LegalNumber field to given value.

### HasLegalNumber

`func (o *OrganizationObject) HasLegalNumber() bool`

HasLegalNumber returns a boolean if a field has been set.

### SetLegalNumberNil

`func (o *OrganizationObject) SetLegalNumberNil(b bool)`

 SetLegalNumberNil sets the value for LegalNumber to be an explicit nil

### UnsetLegalNumber
`func (o *OrganizationObject) UnsetLegalNumber()`

UnsetLegalNumber ensures that no value is present for LegalNumber, not even an explicit nil
### GetNetPaymentTerm

`func (o *OrganizationObject) GetNetPaymentTerm() int32`

GetNetPaymentTerm returns the NetPaymentTerm field if non-nil, zero value otherwise.

### GetNetPaymentTermOk

`func (o *OrganizationObject) GetNetPaymentTermOk() (*int32, bool)`

GetNetPaymentTermOk returns a tuple with the NetPaymentTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetPaymentTerm

`func (o *OrganizationObject) SetNetPaymentTerm(v int32)`

SetNetPaymentTerm sets NetPaymentTerm field to given value.

### HasNetPaymentTerm

`func (o *OrganizationObject) HasNetPaymentTerm() bool`

HasNetPaymentTerm returns a boolean if a field has been set.

### GetTaxIdentificationNumber

`func (o *OrganizationObject) GetTaxIdentificationNumber() string`

GetTaxIdentificationNumber returns the TaxIdentificationNumber field if non-nil, zero value otherwise.

### GetTaxIdentificationNumberOk

`func (o *OrganizationObject) GetTaxIdentificationNumberOk() (*string, bool)`

GetTaxIdentificationNumberOk returns a tuple with the TaxIdentificationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxIdentificationNumber

`func (o *OrganizationObject) SetTaxIdentificationNumber(v string)`

SetTaxIdentificationNumber sets TaxIdentificationNumber field to given value.

### HasTaxIdentificationNumber

`func (o *OrganizationObject) HasTaxIdentificationNumber() bool`

HasTaxIdentificationNumber returns a boolean if a field has been set.

### SetTaxIdentificationNumberNil

`func (o *OrganizationObject) SetTaxIdentificationNumberNil(b bool)`

 SetTaxIdentificationNumberNil sets the value for TaxIdentificationNumber to be an explicit nil

### UnsetTaxIdentificationNumber
`func (o *OrganizationObject) UnsetTaxIdentificationNumber()`

UnsetTaxIdentificationNumber ensures that no value is present for TaxIdentificationNumber, not even an explicit nil
### GetTimezone

`func (o *OrganizationObject) GetTimezone() Timezone`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *OrganizationObject) GetTimezoneOk() (*Timezone, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *OrganizationObject) SetTimezone(v Timezone)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *OrganizationObject) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetBillingConfiguration

`func (o *OrganizationObject) GetBillingConfiguration() OrganizationBillingConfiguration`

GetBillingConfiguration returns the BillingConfiguration field if non-nil, zero value otherwise.

### GetBillingConfigurationOk

`func (o *OrganizationObject) GetBillingConfigurationOk() (*OrganizationBillingConfiguration, bool)`

GetBillingConfigurationOk returns a tuple with the BillingConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingConfiguration

`func (o *OrganizationObject) SetBillingConfiguration(v OrganizationBillingConfiguration)`

SetBillingConfiguration sets BillingConfiguration field to given value.


### GetTaxes

`func (o *OrganizationObject) GetTaxes() []TaxObject`

GetTaxes returns the Taxes field if non-nil, zero value otherwise.

### GetTaxesOk

`func (o *OrganizationObject) GetTaxesOk() (*[]TaxObject, bool)`

GetTaxesOk returns a tuple with the Taxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxes

`func (o *OrganizationObject) SetTaxes(v []TaxObject)`

SetTaxes sets Taxes field to given value.

### HasTaxes

`func (o *OrganizationObject) HasTaxes() bool`

HasTaxes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


