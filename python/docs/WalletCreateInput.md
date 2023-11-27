# WalletCreateInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**wallet** | [**WalletCreateInputWallet**](WalletCreateInputWallet.md) |  | [optional] 

## Example

```python
from lago_api.models.wallet_create_input import WalletCreateInput

# TODO update the JSON string below
json = "{}"
# create an instance of WalletCreateInput from a JSON string
wallet_create_input_instance = WalletCreateInput.from_json(json)
# print the JSON string representation of the object
print WalletCreateInput.to_json()

# convert the object into a dict
wallet_create_input_dict = wallet_create_input_instance.to_dict()
# create an instance of WalletCreateInput from a dict
wallet_create_input_form_dict = wallet_create_input.from_dict(wallet_create_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


