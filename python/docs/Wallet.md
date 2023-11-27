# Wallet


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**wallet** | [**WalletObject**](WalletObject.md) |  | 

## Example

```python
from lago_api.models.wallet import Wallet

# TODO update the JSON string below
json = "{}"
# create an instance of Wallet from a JSON string
wallet_instance = Wallet.from_json(json)
# print the JSON string representation of the object
print Wallet.to_json()

# convert the object into a dict
wallet_dict = wallet_instance.to_dict()
# create an instance of Wallet from a dict
wallet_form_dict = wallet.from_dict(wallet_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


