# WalletObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LagoId** | **string** | Unique identifier assigned to the wallet within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the wallet’s record within the Lago system. | 
**LagoCustomerId** | **string** | Unique identifier assigned to the customer within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the customer’s record within the Lago system. | 
**ExternalCustomerId** | **string** | The customer external unique identifier (provided by your own application) | 
**Status** | **string** | The status of the wallet. Possible values are &#x60;active&#x60; or &#x60;terminated&#x60;. | 
**Currency** | [**Currency**](Currency.md) |  | 
**Name** | Pointer to **string** | The name of the wallet. | [optional] 
**RateAmount** | **string** | The rate of conversion between credits and the amount in the specified currency. It indicates the ratio or factor used to convert credits into the corresponding monetary value in the currency of the transaction. | 
**CreditsBalance** | **string** | The current wallet balance expressed in credits. | 
**BalanceCents** | **int32** | The current wallet balance expressed in cents. | 
**ConsumedCredits** | **string** | The number of consumed credits. | 
**CreatedAt** | **time.Time** | The date of the wallet creation, represented in ISO 8601 datetime format and expressed in Coordinated Universal Time (UTC). | 
**ExpirationAt** | Pointer to **NullableTime** | The date and time that determines when the wallet will expire. It follows the ISO 8601 datetime format and is expressed in Coordinated Universal Time (UTC). | [optional] 
**LastBalanceSyncAt** | Pointer to **NullableTime** | The date and time of the last balance top-up. It follows the ISO 8601 datetime format and is expressed in Coordinated Universal Time (UTC). | [optional] 
**LastConsumedCreditAt** | Pointer to **NullableTime** | The date and time of the last credits consumption. It follows the ISO 8601 datetime format and is expressed in Coordinated Universal Time (UTC). | [optional] 
**TerminatedAt** | Pointer to **NullableTime** | The date of terminaison of the wallet. It follows the ISO 8601 datetime format and is expressed in Coordinated Universal Time (UTC). | [optional] 

## Methods

### NewWalletObject

`func NewWalletObject(lagoId string, lagoCustomerId string, externalCustomerId string, status string, currency Currency, rateAmount string, creditsBalance string, balanceCents int32, consumedCredits string, createdAt time.Time, ) *WalletObject`

NewWalletObject instantiates a new WalletObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletObjectWithDefaults

`func NewWalletObjectWithDefaults() *WalletObject`

NewWalletObjectWithDefaults instantiates a new WalletObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLagoId

`func (o *WalletObject) GetLagoId() string`

GetLagoId returns the LagoId field if non-nil, zero value otherwise.

### GetLagoIdOk

`func (o *WalletObject) GetLagoIdOk() (*string, bool)`

GetLagoIdOk returns a tuple with the LagoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagoId

`func (o *WalletObject) SetLagoId(v string)`

SetLagoId sets LagoId field to given value.


### GetLagoCustomerId

`func (o *WalletObject) GetLagoCustomerId() string`

GetLagoCustomerId returns the LagoCustomerId field if non-nil, zero value otherwise.

### GetLagoCustomerIdOk

`func (o *WalletObject) GetLagoCustomerIdOk() (*string, bool)`

GetLagoCustomerIdOk returns a tuple with the LagoCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagoCustomerId

`func (o *WalletObject) SetLagoCustomerId(v string)`

SetLagoCustomerId sets LagoCustomerId field to given value.


### GetExternalCustomerId

`func (o *WalletObject) GetExternalCustomerId() string`

GetExternalCustomerId returns the ExternalCustomerId field if non-nil, zero value otherwise.

### GetExternalCustomerIdOk

`func (o *WalletObject) GetExternalCustomerIdOk() (*string, bool)`

GetExternalCustomerIdOk returns a tuple with the ExternalCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalCustomerId

`func (o *WalletObject) SetExternalCustomerId(v string)`

SetExternalCustomerId sets ExternalCustomerId field to given value.


### GetStatus

`func (o *WalletObject) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WalletObject) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WalletObject) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCurrency

`func (o *WalletObject) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *WalletObject) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *WalletObject) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.


### GetName

`func (o *WalletObject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WalletObject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WalletObject) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WalletObject) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRateAmount

`func (o *WalletObject) GetRateAmount() string`

GetRateAmount returns the RateAmount field if non-nil, zero value otherwise.

### GetRateAmountOk

`func (o *WalletObject) GetRateAmountOk() (*string, bool)`

GetRateAmountOk returns a tuple with the RateAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateAmount

`func (o *WalletObject) SetRateAmount(v string)`

SetRateAmount sets RateAmount field to given value.


### GetCreditsBalance

`func (o *WalletObject) GetCreditsBalance() string`

GetCreditsBalance returns the CreditsBalance field if non-nil, zero value otherwise.

### GetCreditsBalanceOk

`func (o *WalletObject) GetCreditsBalanceOk() (*string, bool)`

GetCreditsBalanceOk returns a tuple with the CreditsBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditsBalance

`func (o *WalletObject) SetCreditsBalance(v string)`

SetCreditsBalance sets CreditsBalance field to given value.


### GetBalanceCents

`func (o *WalletObject) GetBalanceCents() int32`

GetBalanceCents returns the BalanceCents field if non-nil, zero value otherwise.

### GetBalanceCentsOk

`func (o *WalletObject) GetBalanceCentsOk() (*int32, bool)`

GetBalanceCentsOk returns a tuple with the BalanceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceCents

`func (o *WalletObject) SetBalanceCents(v int32)`

SetBalanceCents sets BalanceCents field to given value.


### GetConsumedCredits

`func (o *WalletObject) GetConsumedCredits() string`

GetConsumedCredits returns the ConsumedCredits field if non-nil, zero value otherwise.

### GetConsumedCreditsOk

`func (o *WalletObject) GetConsumedCreditsOk() (*string, bool)`

GetConsumedCreditsOk returns a tuple with the ConsumedCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumedCredits

`func (o *WalletObject) SetConsumedCredits(v string)`

SetConsumedCredits sets ConsumedCredits field to given value.


### GetCreatedAt

`func (o *WalletObject) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WalletObject) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WalletObject) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetExpirationAt

`func (o *WalletObject) GetExpirationAt() time.Time`

GetExpirationAt returns the ExpirationAt field if non-nil, zero value otherwise.

### GetExpirationAtOk

`func (o *WalletObject) GetExpirationAtOk() (*time.Time, bool)`

GetExpirationAtOk returns a tuple with the ExpirationAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationAt

`func (o *WalletObject) SetExpirationAt(v time.Time)`

SetExpirationAt sets ExpirationAt field to given value.

### HasExpirationAt

`func (o *WalletObject) HasExpirationAt() bool`

HasExpirationAt returns a boolean if a field has been set.

### SetExpirationAtNil

`func (o *WalletObject) SetExpirationAtNil(b bool)`

 SetExpirationAtNil sets the value for ExpirationAt to be an explicit nil

### UnsetExpirationAt
`func (o *WalletObject) UnsetExpirationAt()`

UnsetExpirationAt ensures that no value is present for ExpirationAt, not even an explicit nil
### GetLastBalanceSyncAt

`func (o *WalletObject) GetLastBalanceSyncAt() time.Time`

GetLastBalanceSyncAt returns the LastBalanceSyncAt field if non-nil, zero value otherwise.

### GetLastBalanceSyncAtOk

`func (o *WalletObject) GetLastBalanceSyncAtOk() (*time.Time, bool)`

GetLastBalanceSyncAtOk returns a tuple with the LastBalanceSyncAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBalanceSyncAt

`func (o *WalletObject) SetLastBalanceSyncAt(v time.Time)`

SetLastBalanceSyncAt sets LastBalanceSyncAt field to given value.

### HasLastBalanceSyncAt

`func (o *WalletObject) HasLastBalanceSyncAt() bool`

HasLastBalanceSyncAt returns a boolean if a field has been set.

### SetLastBalanceSyncAtNil

`func (o *WalletObject) SetLastBalanceSyncAtNil(b bool)`

 SetLastBalanceSyncAtNil sets the value for LastBalanceSyncAt to be an explicit nil

### UnsetLastBalanceSyncAt
`func (o *WalletObject) UnsetLastBalanceSyncAt()`

UnsetLastBalanceSyncAt ensures that no value is present for LastBalanceSyncAt, not even an explicit nil
### GetLastConsumedCreditAt

`func (o *WalletObject) GetLastConsumedCreditAt() time.Time`

GetLastConsumedCreditAt returns the LastConsumedCreditAt field if non-nil, zero value otherwise.

### GetLastConsumedCreditAtOk

`func (o *WalletObject) GetLastConsumedCreditAtOk() (*time.Time, bool)`

GetLastConsumedCreditAtOk returns a tuple with the LastConsumedCreditAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastConsumedCreditAt

`func (o *WalletObject) SetLastConsumedCreditAt(v time.Time)`

SetLastConsumedCreditAt sets LastConsumedCreditAt field to given value.

### HasLastConsumedCreditAt

`func (o *WalletObject) HasLastConsumedCreditAt() bool`

HasLastConsumedCreditAt returns a boolean if a field has been set.

### SetLastConsumedCreditAtNil

`func (o *WalletObject) SetLastConsumedCreditAtNil(b bool)`

 SetLastConsumedCreditAtNil sets the value for LastConsumedCreditAt to be an explicit nil

### UnsetLastConsumedCreditAt
`func (o *WalletObject) UnsetLastConsumedCreditAt()`

UnsetLastConsumedCreditAt ensures that no value is present for LastConsumedCreditAt, not even an explicit nil
### GetTerminatedAt

`func (o *WalletObject) GetTerminatedAt() time.Time`

GetTerminatedAt returns the TerminatedAt field if non-nil, zero value otherwise.

### GetTerminatedAtOk

`func (o *WalletObject) GetTerminatedAtOk() (*time.Time, bool)`

GetTerminatedAtOk returns a tuple with the TerminatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminatedAt

`func (o *WalletObject) SetTerminatedAt(v time.Time)`

SetTerminatedAt sets TerminatedAt field to given value.

### HasTerminatedAt

`func (o *WalletObject) HasTerminatedAt() bool`

HasTerminatedAt returns a boolean if a field has been set.

### SetTerminatedAtNil

`func (o *WalletObject) SetTerminatedAtNil(b bool)`

 SetTerminatedAtNil sets the value for TerminatedAt to be an explicit nil

### UnsetTerminatedAt
`func (o *WalletObject) UnsetTerminatedAt()`

UnsetTerminatedAt ensures that no value is present for TerminatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


