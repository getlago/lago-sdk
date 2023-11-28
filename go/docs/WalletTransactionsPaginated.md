# WalletTransactionsPaginated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WalletTransactions** | [**[]WalletTransactionObject**](WalletTransactionObject.md) |  | 
**Meta** | [**PaginationMeta**](PaginationMeta.md) |  | 

## Methods

### NewWalletTransactionsPaginated

`func NewWalletTransactionsPaginated(walletTransactions []WalletTransactionObject, meta PaginationMeta, ) *WalletTransactionsPaginated`

NewWalletTransactionsPaginated instantiates a new WalletTransactionsPaginated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletTransactionsPaginatedWithDefaults

`func NewWalletTransactionsPaginatedWithDefaults() *WalletTransactionsPaginated`

NewWalletTransactionsPaginatedWithDefaults instantiates a new WalletTransactionsPaginated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWalletTransactions

`func (o *WalletTransactionsPaginated) GetWalletTransactions() []WalletTransactionObject`

GetWalletTransactions returns the WalletTransactions field if non-nil, zero value otherwise.

### GetWalletTransactionsOk

`func (o *WalletTransactionsPaginated) GetWalletTransactionsOk() (*[]WalletTransactionObject, bool)`

GetWalletTransactionsOk returns a tuple with the WalletTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletTransactions

`func (o *WalletTransactionsPaginated) SetWalletTransactions(v []WalletTransactionObject)`

SetWalletTransactions sets WalletTransactions field to given value.


### GetMeta

`func (o *WalletTransactionsPaginated) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *WalletTransactionsPaginated) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *WalletTransactionsPaginated) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


