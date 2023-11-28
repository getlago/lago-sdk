# WalletUpdateInputWallet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | The name of the wallet. | [optional] 
**ExpirationAt** | Pointer to **NullableTime** | The date and time that determines when the wallet will expire. It follows the ISO 8601 datetime format and is expressed in Coordinated Universal Time (UTC). | [optional] 

## Methods

### NewWalletUpdateInputWallet

`func NewWalletUpdateInputWallet() *WalletUpdateInputWallet`

NewWalletUpdateInputWallet instantiates a new WalletUpdateInputWallet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletUpdateInputWalletWithDefaults

`func NewWalletUpdateInputWalletWithDefaults() *WalletUpdateInputWallet`

NewWalletUpdateInputWalletWithDefaults instantiates a new WalletUpdateInputWallet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WalletUpdateInputWallet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WalletUpdateInputWallet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WalletUpdateInputWallet) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WalletUpdateInputWallet) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *WalletUpdateInputWallet) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *WalletUpdateInputWallet) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetExpirationAt

`func (o *WalletUpdateInputWallet) GetExpirationAt() time.Time`

GetExpirationAt returns the ExpirationAt field if non-nil, zero value otherwise.

### GetExpirationAtOk

`func (o *WalletUpdateInputWallet) GetExpirationAtOk() (*time.Time, bool)`

GetExpirationAtOk returns a tuple with the ExpirationAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationAt

`func (o *WalletUpdateInputWallet) SetExpirationAt(v time.Time)`

SetExpirationAt sets ExpirationAt field to given value.

### HasExpirationAt

`func (o *WalletUpdateInputWallet) HasExpirationAt() bool`

HasExpirationAt returns a boolean if a field has been set.

### SetExpirationAtNil

`func (o *WalletUpdateInputWallet) SetExpirationAtNil(b bool)`

 SetExpirationAtNil sets the value for ExpirationAt to be an explicit nil

### UnsetExpirationAt
`func (o *WalletUpdateInputWallet) UnsetExpirationAt()`

UnsetExpirationAt ensures that no value is present for ExpirationAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


