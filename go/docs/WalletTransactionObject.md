# WalletTransactionObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LagoId** | **string** | Unique identifier assigned to the wallet transaction within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the wallet transaction’s record within the Lago system. | 
**LagoWalletId** | **string** | Unique identifier assigned to the wallet within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the wallet’s record within the Lago system. | 
**Status** | **string** | The status of the wallet transaction. Possible values are &#x60;pending&#x60; or &#x60;settled&#x60;. | 
**TransactionType** | **string** | The type of transaction. Possible values are &#x60;inbound&#x60; (increasing the balance) or &#x60;outbound&#x60; (decreasing the balance). | 
**Amount** | **string** | The amount of credits based on the rate and the currency. | 
**CreditAmount** | **string** | The number of credits used in the wallet transaction. | 
**SettledAt** | Pointer to **time.Time** | The date when wallet transaction is settled, represented in ISO 8601 datetime format and expressed in Coordinated Universal Time (UTC). | [optional] 
**CreatedAt** | **time.Time** | The date of the wallet transaction creation, represented in ISO 8601 datetime format and expressed in Coordinated Universal Time (UTC). | 

## Methods

### NewWalletTransactionObject

`func NewWalletTransactionObject(lagoId string, lagoWalletId string, status string, transactionType string, amount string, creditAmount string, createdAt time.Time, ) *WalletTransactionObject`

NewWalletTransactionObject instantiates a new WalletTransactionObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletTransactionObjectWithDefaults

`func NewWalletTransactionObjectWithDefaults() *WalletTransactionObject`

NewWalletTransactionObjectWithDefaults instantiates a new WalletTransactionObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLagoId

`func (o *WalletTransactionObject) GetLagoId() string`

GetLagoId returns the LagoId field if non-nil, zero value otherwise.

### GetLagoIdOk

`func (o *WalletTransactionObject) GetLagoIdOk() (*string, bool)`

GetLagoIdOk returns a tuple with the LagoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagoId

`func (o *WalletTransactionObject) SetLagoId(v string)`

SetLagoId sets LagoId field to given value.


### GetLagoWalletId

`func (o *WalletTransactionObject) GetLagoWalletId() string`

GetLagoWalletId returns the LagoWalletId field if non-nil, zero value otherwise.

### GetLagoWalletIdOk

`func (o *WalletTransactionObject) GetLagoWalletIdOk() (*string, bool)`

GetLagoWalletIdOk returns a tuple with the LagoWalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagoWalletId

`func (o *WalletTransactionObject) SetLagoWalletId(v string)`

SetLagoWalletId sets LagoWalletId field to given value.


### GetStatus

`func (o *WalletTransactionObject) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WalletTransactionObject) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WalletTransactionObject) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTransactionType

`func (o *WalletTransactionObject) GetTransactionType() string`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *WalletTransactionObject) GetTransactionTypeOk() (*string, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *WalletTransactionObject) SetTransactionType(v string)`

SetTransactionType sets TransactionType field to given value.


### GetAmount

`func (o *WalletTransactionObject) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *WalletTransactionObject) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *WalletTransactionObject) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetCreditAmount

`func (o *WalletTransactionObject) GetCreditAmount() string`

GetCreditAmount returns the CreditAmount field if non-nil, zero value otherwise.

### GetCreditAmountOk

`func (o *WalletTransactionObject) GetCreditAmountOk() (*string, bool)`

GetCreditAmountOk returns a tuple with the CreditAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditAmount

`func (o *WalletTransactionObject) SetCreditAmount(v string)`

SetCreditAmount sets CreditAmount field to given value.


### GetSettledAt

`func (o *WalletTransactionObject) GetSettledAt() time.Time`

GetSettledAt returns the SettledAt field if non-nil, zero value otherwise.

### GetSettledAtOk

`func (o *WalletTransactionObject) GetSettledAtOk() (*time.Time, bool)`

GetSettledAtOk returns a tuple with the SettledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettledAt

`func (o *WalletTransactionObject) SetSettledAt(v time.Time)`

SetSettledAt sets SettledAt field to given value.

### HasSettledAt

`func (o *WalletTransactionObject) HasSettledAt() bool`

HasSettledAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *WalletTransactionObject) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WalletTransactionObject) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WalletTransactionObject) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


