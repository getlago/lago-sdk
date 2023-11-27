# WalletTransactions


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**wallet_transactions** | [**List[WalletTransactionObject]**](WalletTransactionObject.md) |  | 

## Example

```python
from lago_api.models.wallet_transactions import WalletTransactions

# TODO update the JSON string below
json = "{}"
# create an instance of WalletTransactions from a JSON string
wallet_transactions_instance = WalletTransactions.from_json(json)
# print the JSON string representation of the object
print WalletTransactions.to_json()

# convert the object into a dict
wallet_transactions_dict = wallet_transactions_instance.to_dict()
# create an instance of WalletTransactions from a dict
wallet_transactions_form_dict = wallet_transactions.from_dict(wallet_transactions_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


