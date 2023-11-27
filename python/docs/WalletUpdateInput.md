# WalletUpdateInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**wallet** | [**WalletUpdateInputWallet**](WalletUpdateInputWallet.md) |  | 

## Example

```python
from lago_api.models.wallet_update_input import WalletUpdateInput

# TODO update the JSON string below
json = "{}"
# create an instance of WalletUpdateInput from a JSON string
wallet_update_input_instance = WalletUpdateInput.from_json(json)
# print the JSON string representation of the object
print WalletUpdateInput.to_json()

# convert the object into a dict
wallet_update_input_dict = wallet_update_input_instance.to_dict()
# create an instance of WalletUpdateInput from a dict
wallet_update_input_form_dict = wallet_update_input.from_dict(wallet_update_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


