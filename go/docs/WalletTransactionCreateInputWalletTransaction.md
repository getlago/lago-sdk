# WalletTransactionCreateInputWalletTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WalletId** | **string** | Unique identifier assigned to the wallet within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the wallet’s record within the Lago system. | 
**PaidCredits** | Pointer to **string** | The number of paid credits. Required only if there is no granted credits. | [optional] 
**GrantedCredits** | Pointer to **string** | The number of free granted credits. Required only if there is no paid credits. | [optional] 

## Methods

### NewWalletTransactionCreateInputWalletTransaction

`func NewWalletTransactionCreateInputWalletTransaction(walletId string, ) *WalletTransactionCreateInputWalletTransaction`

NewWalletTransactionCreateInputWalletTransaction instantiates a new WalletTransactionCreateInputWalletTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletTransactionCreateInputWalletTransactionWithDefaults

`func NewWalletTransactionCreateInputWalletTransactionWithDefaults() *WalletTransactionCreateInputWalletTransaction`

NewWalletTransactionCreateInputWalletTransactionWithDefaults instantiates a new WalletTransactionCreateInputWalletTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWalletId

`func (o *WalletTransactionCreateInputWalletTransaction) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *WalletTransactionCreateInputWalletTransaction) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *WalletTransactionCreateInputWalletTransaction) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetPaidCredits

`func (o *WalletTransactionCreateInputWalletTransaction) GetPaidCredits() string`

GetPaidCredits returns the PaidCredits field if non-nil, zero value otherwise.

### GetPaidCreditsOk

`func (o *WalletTransactionCreateInputWalletTransaction) GetPaidCreditsOk() (*string, bool)`

GetPaidCreditsOk returns a tuple with the PaidCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidCredits

`func (o *WalletTransactionCreateInputWalletTransaction) SetPaidCredits(v string)`

SetPaidCredits sets PaidCredits field to given value.

### HasPaidCredits

`func (o *WalletTransactionCreateInputWalletTransaction) HasPaidCredits() bool`

HasPaidCredits returns a boolean if a field has been set.

### GetGrantedCredits

`func (o *WalletTransactionCreateInputWalletTransaction) GetGrantedCredits() string`

GetGrantedCredits returns the GrantedCredits field if non-nil, zero value otherwise.

### GetGrantedCreditsOk

`func (o *WalletTransactionCreateInputWalletTransaction) GetGrantedCreditsOk() (*string, bool)`

GetGrantedCreditsOk returns a tuple with the GrantedCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantedCredits

`func (o *WalletTransactionCreateInputWalletTransaction) SetGrantedCredits(v string)`

SetGrantedCredits sets GrantedCredits field to given value.

### HasGrantedCredits

`func (o *WalletTransactionCreateInputWalletTransaction) HasGrantedCredits() bool`

HasGrantedCredits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


