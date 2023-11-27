# WalletUpdateInputWallet


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | The name of the wallet. | [optional] 
**expiration_at** | **datetime** | The date and time that determines when the wallet will expire. It follows the ISO 8601 datetime format and is expressed in Coordinated Universal Time (UTC). | [optional] 

## Example

```python
from lago_api.models.wallet_update_input_wallet import WalletUpdateInputWallet

# TODO update the JSON string below
json = "{}"
# create an instance of WalletUpdateInputWallet from a JSON string
wallet_update_input_wallet_instance = WalletUpdateInputWallet.from_json(json)
# print the JSON string representation of the object
print WalletUpdateInputWallet.to_json()

# convert the object into a dict
wallet_update_input_wallet_dict = wallet_update_input_wallet_instance.to_dict()
# create an instance of WalletUpdateInputWallet from a dict
wallet_update_input_wallet_form_dict = wallet_update_input_wallet.from_dict(wallet_update_input_wallet_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


