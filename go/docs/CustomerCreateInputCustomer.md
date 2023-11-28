# CustomerCreateInputCustomer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalId** | **string** | The customer external unique identifier (provided by your own application) | 
**AddressLine1** | Pointer to **NullableString** | The first line of the billing address | [optional] 
**AddressLine2** | Pointer to **NullableString** | The second line of the billing address | [optional] 
**City** | Pointer to **NullableString** | The city of the customer&#39;s billing address | [optional] 
**Country** | Pointer to [**Country**](Country.md) |  | [optional] 
**Currency** | Pointer to [**Currency**](Currency.md) |  | [optional] 
**Email** | Pointer to **NullableString** | The email of the customer | [optional] 
**LegalName** | Pointer to **NullableString** | The legal company name of the customer | [optional] 
**LegalNumber** | Pointer to **NullableString** | The legal company number of the customer | [optional] 
**LogoUrl** | Pointer to **NullableString** | The logo URL of the customer | [optional] 
**Name** | Pointer to **NullableString** | The full name of the customer | [optional] 
**Phone** | Pointer to **NullableString** | The phone number of the customer | [optional] 
**State** | Pointer to **NullableString** | The state of the customer&#39;s billing address | [optional] 
**TaxCodes** | Pointer to **[]string** | List of unique code used to identify the taxes. | [optional] 
**TaxIdentificationNumber** | Pointer to **NullableString** | The tax identification number of the customer | [optional] 
**Timezone** | Pointer to [**Timezone**](Timezone.md) |  | [optional] 
**Url** | Pointer to **NullableString** | The custom website URL of the customer | [optional] 
**Zipcode** | Pointer to **NullableString** | The zipcode of the customer&#39;s billing address | [optional] 
**NetPaymentTerm** | Pointer to **NullableInt32** | The net payment term, expressed in days, specifies the duration within which a customer is expected to remit payment after the invoice is finalized. | [optional] 
**BillingConfiguration** | Pointer to [**CustomerBillingConfiguration**](CustomerBillingConfiguration.md) |  | [optional] 
**Metadata** | Pointer to [**[]CustomerCreateInputCustomerMetadataInner**](CustomerCreateInputCustomerMetadataInner.md) | Set of key-value pairs that you can attach to a customer. This can be useful for storing additional information about the customer in a structured format | [optional] 

## Methods

### NewCustomerCreateInputCustomer

`func NewCustomerCreateInputCustomer(externalId string, ) *CustomerCreateInputCustomer`

NewCustomerCreateInputCustomer instantiates a new CustomerCreateInputCustomer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerCreateInputCustomerWithDefaults

`func NewCustomerCreateInputCustomerWithDefaults() *CustomerCreateInputCustomer`

NewCustomerCreateInputCustomerWithDefaults instantiates a new CustomerCreateInputCustomer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalId

`func (o *CustomerCreateInputCustomer) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CustomerCreateInputCustomer) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CustomerCreateInputCustomer) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetAddressLine1

`func (o *CustomerCreateInputCustomer) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *CustomerCreateInputCustomer) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *CustomerCreateInputCustomer) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *CustomerCreateInputCustomer) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### SetAddressLine1Nil

`func (o *CustomerCreateInputCustomer) SetAddressLine1Nil(b bool)`

 SetAddressLine1Nil sets the value for AddressLine1 to be an explicit nil

### UnsetAddressLine1
`func (o *CustomerCreateInputCustomer) UnsetAddressLine1()`

UnsetAddressLine1 ensures that no value is present for AddressLine1, not even an explicit nil
### GetAddressLine2

`func (o *CustomerCreateInputCustomer) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *CustomerCreateInputCustomer) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *CustomerCreateInputCustomer) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *CustomerCreateInputCustomer) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### SetAddressLine2Nil

`func (o *CustomerCreateInputCustomer) SetAddressLine2Nil(b bool)`

 SetAddressLine2Nil sets the value for AddressLine2 to be an explicit nil

### UnsetAddressLine2
`func (o *CustomerCreateInputCustomer) UnsetAddressLine2()`

UnsetAddressLine2 ensures that no value is present for AddressLine2, not even an explicit nil
### GetCity

`func (o *CustomerCreateInputCustomer) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *CustomerCreateInputCustomer) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *CustomerCreateInputCustomer) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *CustomerCreateInputCustomer) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *CustomerCreateInputCustomer) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *CustomerCreateInputCustomer) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetCountry

`func (o *CustomerCreateInputCustomer) GetCountry() Country`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CustomerCreateInputCustomer) GetCountryOk() (*Country, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CustomerCreateInputCustomer) SetCountry(v Country)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *CustomerCreateInputCustomer) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCurrency

`func (o *CustomerCreateInputCustomer) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CustomerCreateInputCustomer) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CustomerCreateInputCustomer) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CustomerCreateInputCustomer) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetEmail

`func (o *CustomerCreateInputCustomer) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CustomerCreateInputCustomer) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CustomerCreateInputCustomer) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CustomerCreateInputCustomer) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *CustomerCreateInputCustomer) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *CustomerCreateInputCustomer) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetLegalName

`func (o *CustomerCreateInputCustomer) GetLegalName() string`

GetLegalName returns the LegalName field if non-nil, zero value otherwise.

### GetLegalNameOk

`func (o *CustomerCreateInputCustomer) GetLegalNameOk() (*string, bool)`

GetLegalNameOk returns a tuple with the LegalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalName

`func (o *CustomerCreateInputCustomer) SetLegalName(v string)`

SetLegalName sets LegalName field to given value.

### HasLegalName

`func (o *CustomerCreateInputCustomer) HasLegalName() bool`

HasLegalName returns a boolean if a field has been set.

### SetLegalNameNil

`func (o *CustomerCreateInputCustomer) SetLegalNameNil(b bool)`

 SetLegalNameNil sets the value for LegalName to be an explicit nil

### UnsetLegalName
`func (o *CustomerCreateInputCustomer) UnsetLegalName()`

UnsetLegalName ensures that no value is present for LegalName, not even an explicit nil
### GetLegalNumber

`func (o *CustomerCreateInputCustomer) GetLegalNumber() string`

GetLegalNumber returns the LegalNumber field if non-nil, zero value otherwise.

### GetLegalNumberOk

`func (o *CustomerCreateInputCustomer) GetLegalNumberOk() (*string, bool)`

GetLegalNumberOk returns a tuple with the LegalNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalNumber

`func (o *CustomerCreateInputCustomer) SetLegalNumber(v string)`

SetLegalNumber sets LegalNumber field to given value.

### HasLegalNumber

`func (o *CustomerCreateInputCustomer) HasLegalNumber() bool`

HasLegalNumber returns a boolean if a field has been set.

### SetLegalNumberNil

`func (o *CustomerCreateInputCustomer) SetLegalNumberNil(b bool)`

 SetLegalNumberNil sets the value for LegalNumber to be an explicit nil

### UnsetLegalNumber
`func (o *CustomerCreateInputCustomer) UnsetLegalNumber()`

UnsetLegalNumber ensures that no value is present for LegalNumber, not even an explicit nil
### GetLogoUrl

`func (o *CustomerCreateInputCustomer) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *CustomerCreateInputCustomer) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *CustomerCreateInputCustomer) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *CustomerCreateInputCustomer) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### SetLogoUrlNil

`func (o *CustomerCreateInputCustomer) SetLogoUrlNil(b bool)`

 SetLogoUrlNil sets the value for LogoUrl to be an explicit nil

### UnsetLogoUrl
`func (o *CustomerCreateInputCustomer) UnsetLogoUrl()`

UnsetLogoUrl ensures that no value is present for LogoUrl, not even an explicit nil
### GetName

`func (o *CustomerCreateInputCustomer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomerCreateInputCustomer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomerCreateInputCustomer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CustomerCreateInputCustomer) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CustomerCreateInputCustomer) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CustomerCreateInputCustomer) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPhone

`func (o *CustomerCreateInputCustomer) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *CustomerCreateInputCustomer) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *CustomerCreateInputCustomer) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *CustomerCreateInputCustomer) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *CustomerCreateInputCustomer) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *CustomerCreateInputCustomer) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetState

`func (o *CustomerCreateInputCustomer) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CustomerCreateInputCustomer) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CustomerCreateInputCustomer) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CustomerCreateInputCustomer) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *CustomerCreateInputCustomer) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *CustomerCreateInputCustomer) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetTaxCodes

`func (o *CustomerCreateInputCustomer) GetTaxCodes() []string`

GetTaxCodes returns the TaxCodes field if non-nil, zero value otherwise.

### GetTaxCodesOk

`func (o *CustomerCreateInputCustomer) GetTaxCodesOk() (*[]string, bool)`

GetTaxCodesOk returns a tuple with the TaxCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCodes

`func (o *CustomerCreateInputCustomer) SetTaxCodes(v []string)`

SetTaxCodes sets TaxCodes field to given value.

### HasTaxCodes

`func (o *CustomerCreateInputCustomer) HasTaxCodes() bool`

HasTaxCodes returns a boolean if a field has been set.

### GetTaxIdentificationNumber

`func (o *CustomerCreateInputCustomer) GetTaxIdentificationNumber() string`

GetTaxIdentificationNumber returns the TaxIdentificationNumber field if non-nil, zero value otherwise.

### GetTaxIdentificationNumberOk

`func (o *CustomerCreateInputCustomer) GetTaxIdentificationNumberOk() (*string, bool)`

GetTaxIdentificationNumberOk returns a tuple with the TaxIdentificationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxIdentificationNumber

`func (o *CustomerCreateInputCustomer) SetTaxIdentificationNumber(v string)`

SetTaxIdentificationNumber sets TaxIdentificationNumber field to given value.

### HasTaxIdentificationNumber

`func (o *CustomerCreateInputCustomer) HasTaxIdentificationNumber() bool`

HasTaxIdentificationNumber returns a boolean if a field has been set.

### SetTaxIdentificationNumberNil

`func (o *CustomerCreateInputCustomer) SetTaxIdentificationNumberNil(b bool)`

 SetTaxIdentificationNumberNil sets the value for TaxIdentificationNumber to be an explicit nil

### UnsetTaxIdentificationNumber
`func (o *CustomerCreateInputCustomer) UnsetTaxIdentificationNumber()`

UnsetTaxIdentificationNumber ensures that no value is present for TaxIdentificationNumber, not even an explicit nil
### GetTimezone

`func (o *CustomerCreateInputCustomer) GetTimezone() Timezone`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *CustomerCreateInputCustomer) GetTimezoneOk() (*Timezone, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *CustomerCreateInputCustomer) SetTimezone(v Timezone)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *CustomerCreateInputCustomer) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetUrl

`func (o *CustomerCreateInputCustomer) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CustomerCreateInputCustomer) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CustomerCreateInputCustomer) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CustomerCreateInputCustomer) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *CustomerCreateInputCustomer) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *CustomerCreateInputCustomer) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetZipcode

`func (o *CustomerCreateInputCustomer) GetZipcode() string`

GetZipcode returns the Zipcode field if non-nil, zero value otherwise.

### GetZipcodeOk

`func (o *CustomerCreateInputCustomer) GetZipcodeOk() (*string, bool)`

GetZipcodeOk returns a tuple with the Zipcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipcode

`func (o *CustomerCreateInputCustomer) SetZipcode(v string)`

SetZipcode sets Zipcode field to given value.

### HasZipcode

`func (o *CustomerCreateInputCustomer) HasZipcode() bool`

HasZipcode returns a boolean if a field has been set.

### SetZipcodeNil

`func (o *CustomerCreateInputCustomer) SetZipcodeNil(b bool)`

 SetZipcodeNil sets the value for Zipcode to be an explicit nil

### UnsetZipcode
`func (o *CustomerCreateInputCustomer) UnsetZipcode()`

UnsetZipcode ensures that no value is present for Zipcode, not even an explicit nil
### GetNetPaymentTerm

`func (o *CustomerCreateInputCustomer) GetNetPaymentTerm() int32`

GetNetPaymentTerm returns the NetPaymentTerm field if non-nil, zero value otherwise.

### GetNetPaymentTermOk

`func (o *CustomerCreateInputCustomer) GetNetPaymentTermOk() (*int32, bool)`

GetNetPaymentTermOk returns a tuple with the NetPaymentTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetPaymentTerm

`func (o *CustomerCreateInputCustomer) SetNetPaymentTerm(v int32)`

SetNetPaymentTerm sets NetPaymentTerm field to given value.

### HasNetPaymentTerm

`func (o *CustomerCreateInputCustomer) HasNetPaymentTerm() bool`

HasNetPaymentTerm returns a boolean if a field has been set.

### SetNetPaymentTermNil

`func (o *CustomerCreateInputCustomer) SetNetPaymentTermNil(b bool)`

 SetNetPaymentTermNil sets the value for NetPaymentTerm to be an explicit nil

### UnsetNetPaymentTerm
`func (o *CustomerCreateInputCustomer) UnsetNetPaymentTerm()`

UnsetNetPaymentTerm ensures that no value is present for NetPaymentTerm, not even an explicit nil
### GetBillingConfiguration

`func (o *CustomerCreateInputCustomer) GetBillingConfiguration() CustomerBillingConfiguration`

GetBillingConfiguration returns the BillingConfiguration field if non-nil, zero value otherwise.

### GetBillingConfigurationOk

`func (o *CustomerCreateInputCustomer) GetBillingConfigurationOk() (*CustomerBillingConfiguration, bool)`

GetBillingConfigurationOk returns a tuple with the BillingConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingConfiguration

`func (o *CustomerCreateInputCustomer) SetBillingConfiguration(v CustomerBillingConfiguration)`

SetBillingConfiguration sets BillingConfiguration field to given value.

### HasBillingConfiguration

`func (o *CustomerCreateInputCustomer) HasBillingConfiguration() bool`

HasBillingConfiguration returns a boolean if a field has been set.

### GetMetadata

`func (o *CustomerCreateInputCustomer) GetMetadata() []CustomerCreateInputCustomerMetadataInner`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CustomerCreateInputCustomer) GetMetadataOk() (*[]CustomerCreateInputCustomerMetadataInner, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CustomerCreateInputCustomer) SetMetadata(v []CustomerCreateInputCustomerMetadataInner)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CustomerCreateInputCustomer) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


