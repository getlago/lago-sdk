# CustomerObjectExtended

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LagoId** | **string** | Unique identifier assigned to the customer within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the customer&#39;s record within the Lago system | 
**SequentialId** | **int32** | The unique identifier assigned to the customer within the organization&#39;s scope. This identifier is used to track and reference the customer&#39;s order of creation within the organization&#39;s system. It ensures that each customer has a distinct &#x60;sequential_id&#x60;&#x60; associated with them, allowing for easy identification and sorting based on the order of creation | 
**Slug** | **string** | A concise and unique identifier for the customer, formed by combining the Organization&#39;s &#x60;name&#x60;, &#x60;id&#x60;, and customer&#39;s &#x60;sequential_id&#x60; | 
**ExternalId** | **string** | The customer external unique identifier (provided by your own application) | 
**AddressLine1** | Pointer to **NullableString** | The first line of the billing address | [optional] 
**AddressLine2** | Pointer to **NullableString** | The second line of the billing address | [optional] 
**ApplicableTimezone** | [**Timezone**](Timezone.md) |  | 
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
**TaxIdentificationNumber** | Pointer to **NullableString** | The tax identification number of the customer | [optional] 
**Timezone** | Pointer to [**Timezone**](Timezone.md) |  | [optional] 
**Url** | Pointer to **NullableString** | The custom website URL of the customer | [optional] 
**Zipcode** | Pointer to **NullableString** | The zipcode of the customer&#39;s billing address | [optional] 
**NetPaymentTerm** | Pointer to **NullableInt32** | The net payment term, expressed in days, specifies the duration within which a customer is expected to remit payment after the invoice is finalized. | [optional] 
**CreatedAt** | **time.Time** | The date of the customer creation, represented in ISO 8601 datetime format and expressed in Coordinated Universal Time (UTC). The creation_date provides a standardized and internationally recognized timestamp for when the customer object was created | 
**UpdatedAt** | Pointer to **time.Time** | The date of the customer update, represented in ISO 8601 datetime format and expressed in Coordinated Universal Time (UTC). The update_date provides a standardized and internationally recognized timestamp for when the customer object was updated | [optional] 
**BillingConfiguration** | Pointer to [**CustomerBillingConfiguration**](CustomerBillingConfiguration.md) |  | [optional] 
**Metadata** | Pointer to [**[]CustomerMetadata**](CustomerMetadata.md) |  | [optional] 
**Taxes** | Pointer to [**[]TaxObject**](TaxObject.md) | List of customer taxes | [optional] 

## Methods

### NewCustomerObjectExtended

`func NewCustomerObjectExtended(lagoId string, sequentialId int32, slug string, externalId string, applicableTimezone Timezone, createdAt time.Time, ) *CustomerObjectExtended`

NewCustomerObjectExtended instantiates a new CustomerObjectExtended object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerObjectExtendedWithDefaults

`func NewCustomerObjectExtendedWithDefaults() *CustomerObjectExtended`

NewCustomerObjectExtendedWithDefaults instantiates a new CustomerObjectExtended object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLagoId

`func (o *CustomerObjectExtended) GetLagoId() string`

GetLagoId returns the LagoId field if non-nil, zero value otherwise.

### GetLagoIdOk

`func (o *CustomerObjectExtended) GetLagoIdOk() (*string, bool)`

GetLagoIdOk returns a tuple with the LagoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagoId

`func (o *CustomerObjectExtended) SetLagoId(v string)`

SetLagoId sets LagoId field to given value.


### GetSequentialId

`func (o *CustomerObjectExtended) GetSequentialId() int32`

GetSequentialId returns the SequentialId field if non-nil, zero value otherwise.

### GetSequentialIdOk

`func (o *CustomerObjectExtended) GetSequentialIdOk() (*int32, bool)`

GetSequentialIdOk returns a tuple with the SequentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequentialId

`func (o *CustomerObjectExtended) SetSequentialId(v int32)`

SetSequentialId sets SequentialId field to given value.


### GetSlug

`func (o *CustomerObjectExtended) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *CustomerObjectExtended) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *CustomerObjectExtended) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetExternalId

`func (o *CustomerObjectExtended) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CustomerObjectExtended) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CustomerObjectExtended) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetAddressLine1

`func (o *CustomerObjectExtended) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *CustomerObjectExtended) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *CustomerObjectExtended) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *CustomerObjectExtended) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### SetAddressLine1Nil

`func (o *CustomerObjectExtended) SetAddressLine1Nil(b bool)`

 SetAddressLine1Nil sets the value for AddressLine1 to be an explicit nil

### UnsetAddressLine1
`func (o *CustomerObjectExtended) UnsetAddressLine1()`

UnsetAddressLine1 ensures that no value is present for AddressLine1, not even an explicit nil
### GetAddressLine2

`func (o *CustomerObjectExtended) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *CustomerObjectExtended) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *CustomerObjectExtended) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *CustomerObjectExtended) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### SetAddressLine2Nil

`func (o *CustomerObjectExtended) SetAddressLine2Nil(b bool)`

 SetAddressLine2Nil sets the value for AddressLine2 to be an explicit nil

### UnsetAddressLine2
`func (o *CustomerObjectExtended) UnsetAddressLine2()`

UnsetAddressLine2 ensures that no value is present for AddressLine2, not even an explicit nil
### GetApplicableTimezone

`func (o *CustomerObjectExtended) GetApplicableTimezone() Timezone`

GetApplicableTimezone returns the ApplicableTimezone field if non-nil, zero value otherwise.

### GetApplicableTimezoneOk

`func (o *CustomerObjectExtended) GetApplicableTimezoneOk() (*Timezone, bool)`

GetApplicableTimezoneOk returns a tuple with the ApplicableTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicableTimezone

`func (o *CustomerObjectExtended) SetApplicableTimezone(v Timezone)`

SetApplicableTimezone sets ApplicableTimezone field to given value.


### GetCity

`func (o *CustomerObjectExtended) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *CustomerObjectExtended) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *CustomerObjectExtended) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *CustomerObjectExtended) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *CustomerObjectExtended) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *CustomerObjectExtended) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetCountry

`func (o *CustomerObjectExtended) GetCountry() Country`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CustomerObjectExtended) GetCountryOk() (*Country, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CustomerObjectExtended) SetCountry(v Country)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *CustomerObjectExtended) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCurrency

`func (o *CustomerObjectExtended) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CustomerObjectExtended) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CustomerObjectExtended) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CustomerObjectExtended) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetEmail

`func (o *CustomerObjectExtended) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CustomerObjectExtended) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CustomerObjectExtended) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CustomerObjectExtended) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *CustomerObjectExtended) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *CustomerObjectExtended) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetLegalName

`func (o *CustomerObjectExtended) GetLegalName() string`

GetLegalName returns the LegalName field if non-nil, zero value otherwise.

### GetLegalNameOk

`func (o *CustomerObjectExtended) GetLegalNameOk() (*string, bool)`

GetLegalNameOk returns a tuple with the LegalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalName

`func (o *CustomerObjectExtended) SetLegalName(v string)`

SetLegalName sets LegalName field to given value.

### HasLegalName

`func (o *CustomerObjectExtended) HasLegalName() bool`

HasLegalName returns a boolean if a field has been set.

### SetLegalNameNil

`func (o *CustomerObjectExtended) SetLegalNameNil(b bool)`

 SetLegalNameNil sets the value for LegalName to be an explicit nil

### UnsetLegalName
`func (o *CustomerObjectExtended) UnsetLegalName()`

UnsetLegalName ensures that no value is present for LegalName, not even an explicit nil
### GetLegalNumber

`func (o *CustomerObjectExtended) GetLegalNumber() string`

GetLegalNumber returns the LegalNumber field if non-nil, zero value otherwise.

### GetLegalNumberOk

`func (o *CustomerObjectExtended) GetLegalNumberOk() (*string, bool)`

GetLegalNumberOk returns a tuple with the LegalNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalNumber

`func (o *CustomerObjectExtended) SetLegalNumber(v string)`

SetLegalNumber sets LegalNumber field to given value.

### HasLegalNumber

`func (o *CustomerObjectExtended) HasLegalNumber() bool`

HasLegalNumber returns a boolean if a field has been set.

### SetLegalNumberNil

`func (o *CustomerObjectExtended) SetLegalNumberNil(b bool)`

 SetLegalNumberNil sets the value for LegalNumber to be an explicit nil

### UnsetLegalNumber
`func (o *CustomerObjectExtended) UnsetLegalNumber()`

UnsetLegalNumber ensures that no value is present for LegalNumber, not even an explicit nil
### GetLogoUrl

`func (o *CustomerObjectExtended) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *CustomerObjectExtended) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *CustomerObjectExtended) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *CustomerObjectExtended) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### SetLogoUrlNil

`func (o *CustomerObjectExtended) SetLogoUrlNil(b bool)`

 SetLogoUrlNil sets the value for LogoUrl to be an explicit nil

### UnsetLogoUrl
`func (o *CustomerObjectExtended) UnsetLogoUrl()`

UnsetLogoUrl ensures that no value is present for LogoUrl, not even an explicit nil
### GetName

`func (o *CustomerObjectExtended) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomerObjectExtended) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomerObjectExtended) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CustomerObjectExtended) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CustomerObjectExtended) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CustomerObjectExtended) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPhone

`func (o *CustomerObjectExtended) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *CustomerObjectExtended) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *CustomerObjectExtended) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *CustomerObjectExtended) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *CustomerObjectExtended) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *CustomerObjectExtended) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetState

`func (o *CustomerObjectExtended) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CustomerObjectExtended) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CustomerObjectExtended) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CustomerObjectExtended) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *CustomerObjectExtended) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *CustomerObjectExtended) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetTaxIdentificationNumber

`func (o *CustomerObjectExtended) GetTaxIdentificationNumber() string`

GetTaxIdentificationNumber returns the TaxIdentificationNumber field if non-nil, zero value otherwise.

### GetTaxIdentificationNumberOk

`func (o *CustomerObjectExtended) GetTaxIdentificationNumberOk() (*string, bool)`

GetTaxIdentificationNumberOk returns a tuple with the TaxIdentificationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxIdentificationNumber

`func (o *CustomerObjectExtended) SetTaxIdentificationNumber(v string)`

SetTaxIdentificationNumber sets TaxIdentificationNumber field to given value.

### HasTaxIdentificationNumber

`func (o *CustomerObjectExtended) HasTaxIdentificationNumber() bool`

HasTaxIdentificationNumber returns a boolean if a field has been set.

### SetTaxIdentificationNumberNil

`func (o *CustomerObjectExtended) SetTaxIdentificationNumberNil(b bool)`

 SetTaxIdentificationNumberNil sets the value for TaxIdentificationNumber to be an explicit nil

### UnsetTaxIdentificationNumber
`func (o *CustomerObjectExtended) UnsetTaxIdentificationNumber()`

UnsetTaxIdentificationNumber ensures that no value is present for TaxIdentificationNumber, not even an explicit nil
### GetTimezone

`func (o *CustomerObjectExtended) GetTimezone() Timezone`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *CustomerObjectExtended) GetTimezoneOk() (*Timezone, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *CustomerObjectExtended) SetTimezone(v Timezone)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *CustomerObjectExtended) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetUrl

`func (o *CustomerObjectExtended) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CustomerObjectExtended) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CustomerObjectExtended) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CustomerObjectExtended) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *CustomerObjectExtended) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *CustomerObjectExtended) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetZipcode

`func (o *CustomerObjectExtended) GetZipcode() string`

GetZipcode returns the Zipcode field if non-nil, zero value otherwise.

### GetZipcodeOk

`func (o *CustomerObjectExtended) GetZipcodeOk() (*string, bool)`

GetZipcodeOk returns a tuple with the Zipcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipcode

`func (o *CustomerObjectExtended) SetZipcode(v string)`

SetZipcode sets Zipcode field to given value.

### HasZipcode

`func (o *CustomerObjectExtended) HasZipcode() bool`

HasZipcode returns a boolean if a field has been set.

### SetZipcodeNil

`func (o *CustomerObjectExtended) SetZipcodeNil(b bool)`

 SetZipcodeNil sets the value for Zipcode to be an explicit nil

### UnsetZipcode
`func (o *CustomerObjectExtended) UnsetZipcode()`

UnsetZipcode ensures that no value is present for Zipcode, not even an explicit nil
### GetNetPaymentTerm

`func (o *CustomerObjectExtended) GetNetPaymentTerm() int32`

GetNetPaymentTerm returns the NetPaymentTerm field if non-nil, zero value otherwise.

### GetNetPaymentTermOk

`func (o *CustomerObjectExtended) GetNetPaymentTermOk() (*int32, bool)`

GetNetPaymentTermOk returns a tuple with the NetPaymentTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetPaymentTerm

`func (o *CustomerObjectExtended) SetNetPaymentTerm(v int32)`

SetNetPaymentTerm sets NetPaymentTerm field to given value.

### HasNetPaymentTerm

`func (o *CustomerObjectExtended) HasNetPaymentTerm() bool`

HasNetPaymentTerm returns a boolean if a field has been set.

### SetNetPaymentTermNil

`func (o *CustomerObjectExtended) SetNetPaymentTermNil(b bool)`

 SetNetPaymentTermNil sets the value for NetPaymentTerm to be an explicit nil

### UnsetNetPaymentTerm
`func (o *CustomerObjectExtended) UnsetNetPaymentTerm()`

UnsetNetPaymentTerm ensures that no value is present for NetPaymentTerm, not even an explicit nil
### GetCreatedAt

`func (o *CustomerObjectExtended) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CustomerObjectExtended) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CustomerObjectExtended) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *CustomerObjectExtended) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CustomerObjectExtended) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CustomerObjectExtended) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CustomerObjectExtended) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetBillingConfiguration

`func (o *CustomerObjectExtended) GetBillingConfiguration() CustomerBillingConfiguration`

GetBillingConfiguration returns the BillingConfiguration field if non-nil, zero value otherwise.

### GetBillingConfigurationOk

`func (o *CustomerObjectExtended) GetBillingConfigurationOk() (*CustomerBillingConfiguration, bool)`

GetBillingConfigurationOk returns a tuple with the BillingConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingConfiguration

`func (o *CustomerObjectExtended) SetBillingConfiguration(v CustomerBillingConfiguration)`

SetBillingConfiguration sets BillingConfiguration field to given value.

### HasBillingConfiguration

`func (o *CustomerObjectExtended) HasBillingConfiguration() bool`

HasBillingConfiguration returns a boolean if a field has been set.

### GetMetadata

`func (o *CustomerObjectExtended) GetMetadata() []CustomerMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CustomerObjectExtended) GetMetadataOk() (*[]CustomerMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CustomerObjectExtended) SetMetadata(v []CustomerMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CustomerObjectExtended) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetTaxes

`func (o *CustomerObjectExtended) GetTaxes() []TaxObject`

GetTaxes returns the Taxes field if non-nil, zero value otherwise.

### GetTaxesOk

`func (o *CustomerObjectExtended) GetTaxesOk() (*[]TaxObject, bool)`

GetTaxesOk returns a tuple with the Taxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxes

`func (o *CustomerObjectExtended) SetTaxes(v []TaxObject)`

SetTaxes sets Taxes field to given value.

### HasTaxes

`func (o *CustomerObjectExtended) HasTaxes() bool`

HasTaxes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


