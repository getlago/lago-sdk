# WalletCreateInputWallet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the wallet. | [optional] 
**RateAmount** | **string** | The rate of conversion between credits and the amount in the specified currency. It indicates the ratio or factor used to convert credits into the corresponding monetary value in the currency of the transaction. | 
**Currency** | [**Currency**](Currency.md) |  | 
**PaidCredits** | Pointer to **NullableString** | The number of paid credits. Required only if there is no granted credits. | [optional] 
**GrantedCredits** | Pointer to **NullableString** | The number of free granted credits. Required only if there is no paid credits. | [optional] 
**ExternalCustomerId** | **string** | The customer external unique identifier (provided by your own application) | 
**ExpirationAt** | Pointer to **NullableTime** | The date and time that determines when the wallet will expire. It follows the ISO 8601 datetime format and is expressed in Coordinated Universal Time (UTC). | [optional] 

## Methods

### NewWalletCreateInputWallet

`func NewWalletCreateInputWallet(rateAmount string, currency Currency, externalCustomerId string, ) *WalletCreateInputWallet`

NewWalletCreateInputWallet instantiates a new WalletCreateInputWallet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletCreateInputWalletWithDefaults

`func NewWalletCreateInputWalletWithDefaults() *WalletCreateInputWallet`

NewWalletCreateInputWalletWithDefaults instantiates a new WalletCreateInputWallet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WalletCreateInputWallet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WalletCreateInputWallet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WalletCreateInputWallet) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WalletCreateInputWallet) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRateAmount

`func (o *WalletCreateInputWallet) GetRateAmount() string`

GetRateAmount returns the RateAmount field if non-nil, zero value otherwise.

### GetRateAmountOk

`func (o *WalletCreateInputWallet) GetRateAmountOk() (*string, bool)`

GetRateAmountOk returns a tuple with the RateAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateAmount

`func (o *WalletCreateInputWallet) SetRateAmount(v string)`

SetRateAmount sets RateAmount field to given value.


### GetCurrency

`func (o *WalletCreateInputWallet) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *WalletCreateInputWallet) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *WalletCreateInputWallet) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.


### GetPaidCredits

`func (o *WalletCreateInputWallet) GetPaidCredits() string`

GetPaidCredits returns the PaidCredits field if non-nil, zero value otherwise.

### GetPaidCreditsOk

`func (o *WalletCreateInputWallet) GetPaidCreditsOk() (*string, bool)`

GetPaidCreditsOk returns a tuple with the PaidCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidCredits

`func (o *WalletCreateInputWallet) SetPaidCredits(v string)`

SetPaidCredits sets PaidCredits field to given value.

### HasPaidCredits

`func (o *WalletCreateInputWallet) HasPaidCredits() bool`

HasPaidCredits returns a boolean if a field has been set.

### SetPaidCreditsNil

`func (o *WalletCreateInputWallet) SetPaidCreditsNil(b bool)`

 SetPaidCreditsNil sets the value for PaidCredits to be an explicit nil

### UnsetPaidCredits
`func (o *WalletCreateInputWallet) UnsetPaidCredits()`

UnsetPaidCredits ensures that no value is present for PaidCredits, not even an explicit nil
### GetGrantedCredits

`func (o *WalletCreateInputWallet) GetGrantedCredits() string`

GetGrantedCredits returns the GrantedCredits field if non-nil, zero value otherwise.

### GetGrantedCreditsOk

`func (o *WalletCreateInputWallet) GetGrantedCreditsOk() (*string, bool)`

GetGrantedCreditsOk returns a tuple with the GrantedCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantedCredits

`func (o *WalletCreateInputWallet) SetGrantedCredits(v string)`

SetGrantedCredits sets GrantedCredits field to given value.

### HasGrantedCredits

`func (o *WalletCreateInputWallet) HasGrantedCredits() bool`

HasGrantedCredits returns a boolean if a field has been set.

### SetGrantedCreditsNil

`func (o *WalletCreateInputWallet) SetGrantedCreditsNil(b bool)`

 SetGrantedCreditsNil sets the value for GrantedCredits to be an explicit nil

### UnsetGrantedCredits
`func (o *WalletCreateInputWallet) UnsetGrantedCredits()`

UnsetGrantedCredits ensures that no value is present for GrantedCredits, not even an explicit nil
### GetExternalCustomerId

`func (o *WalletCreateInputWallet) GetExternalCustomerId() string`

GetExternalCustomerId returns the ExternalCustomerId field if non-nil, zero value otherwise.

### GetExternalCustomerIdOk

`func (o *WalletCreateInputWallet) GetExternalCustomerIdOk() (*string, bool)`

GetExternalCustomerIdOk returns a tuple with the ExternalCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalCustomerId

`func (o *WalletCreateInputWallet) SetExternalCustomerId(v string)`

SetExternalCustomerId sets ExternalCustomerId field to given value.


### GetExpirationAt

`func (o *WalletCreateInputWallet) GetExpirationAt() time.Time`

GetExpirationAt returns the ExpirationAt field if non-nil, zero value otherwise.

### GetExpirationAtOk

`func (o *WalletCreateInputWallet) GetExpirationAtOk() (*time.Time, bool)`

GetExpirationAtOk returns a tuple with the ExpirationAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationAt

`func (o *WalletCreateInputWallet) SetExpirationAt(v time.Time)`

SetExpirationAt sets ExpirationAt field to given value.

### HasExpirationAt

`func (o *WalletCreateInputWallet) HasExpirationAt() bool`

HasExpirationAt returns a boolean if a field has been set.

### SetExpirationAtNil

`func (o *WalletCreateInputWallet) SetExpirationAtNil(b bool)`

 SetExpirationAtNil sets the value for ExpirationAt to be an explicit nil

### UnsetExpirationAt
`func (o *WalletCreateInputWallet) UnsetExpirationAt()`

UnsetExpirationAt ensures that no value is present for ExpirationAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


