# WalletTransactionCreateInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**wallet_transaction** | [**WalletTransactionCreateInputWalletTransaction**](WalletTransactionCreateInputWalletTransaction.md) |  | 

## Example

```python
from lago_api.models.wallet_transaction_create_input import WalletTransactionCreateInput

# TODO update the JSON string below
json = "{}"
# create an instance of WalletTransactionCreateInput from a JSON string
wallet_transaction_create_input_instance = WalletTransactionCreateInput.from_json(json)
# print the JSON string representation of the object
print WalletTransactionCreateInput.to_json()

# convert the object into a dict
wallet_transaction_create_input_dict = wallet_transaction_create_input_instance.to_dict()
# create an instance of WalletTransactionCreateInput from a dict
wallet_transaction_create_input_form_dict = wallet_transaction_create_input.from_dict(wallet_transaction_create_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


