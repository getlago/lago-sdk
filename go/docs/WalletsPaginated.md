# WalletsPaginated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Wallets** | [**[]WalletObject**](WalletObject.md) |  | 
**Meta** | [**PaginationMeta**](PaginationMeta.md) |  | 

## Methods

### NewWalletsPaginated

`func NewWalletsPaginated(wallets []WalletObject, meta PaginationMeta, ) *WalletsPaginated`

NewWalletsPaginated instantiates a new WalletsPaginated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletsPaginatedWithDefaults

`func NewWalletsPaginatedWithDefaults() *WalletsPaginated`

NewWalletsPaginatedWithDefaults instantiates a new WalletsPaginated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWallets

`func (o *WalletsPaginated) GetWallets() []WalletObject`

GetWallets returns the Wallets field if non-nil, zero value otherwise.

### GetWalletsOk

`func (o *WalletsPaginated) GetWalletsOk() (*[]WalletObject, bool)`

GetWalletsOk returns a tuple with the Wallets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWallets

`func (o *WalletsPaginated) SetWallets(v []WalletObject)`

SetWallets sets Wallets field to given value.


### GetMeta

`func (o *WalletsPaginated) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *WalletsPaginated) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *WalletsPaginated) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


