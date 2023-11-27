# WalletTransactionCreateInputWalletTransaction


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**wallet_id** | **str** | Unique identifier assigned to the wallet within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the wallet’s record within the Lago system. | 
**paid_credits** | **str** | The number of paid credits. Required only if there is no granted credits. | [optional] 
**granted_credits** | **str** | The number of free granted credits. Required only if there is no paid credits. | [optional] 

## Example

```python
from lago_api.models.wallet_transaction_create_input_wallet_transaction import WalletTransactionCreateInputWalletTransaction

# TODO update the JSON string below
json = "{}"
# create an instance of WalletTransactionCreateInputWalletTransaction from a JSON string
wallet_transaction_create_input_wallet_transaction_instance = WalletTransactionCreateInputWalletTransaction.from_json(json)
# print the JSON string representation of the object
print WalletTransactionCreateInputWalletTransaction.to_json()

# convert the object into a dict
wallet_transaction_create_input_wallet_transaction_dict = wallet_transaction_create_input_wallet_transaction_instance.to_dict()
# create an instance of WalletTransactionCreateInputWalletTransaction from a dict
wallet_transaction_create_input_wallet_transaction_form_dict = wallet_transaction_create_input_wallet_transaction.from_dict(wallet_transaction_create_input_wallet_transaction_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


