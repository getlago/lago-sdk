# WalletCreateInputWallet


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | The name of the wallet. | [optional] 
**rate_amount** | **str** | The rate of conversion between credits and the amount in the specified currency. It indicates the ratio or factor used to convert credits into the corresponding monetary value in the currency of the transaction. | 
**currency** | [**Currency**](Currency.md) |  | 
**paid_credits** | **str** | The number of paid credits. Required only if there is no granted credits. | [optional] 
**granted_credits** | **str** | The number of free granted credits. Required only if there is no paid credits. | [optional] 
**external_customer_id** | **str** | The customer external unique identifier (provided by your own application) | 
**expiration_at** | **datetime** | The date and time that determines when the wallet will expire. It follows the ISO 8601 datetime format and is expressed in Coordinated Universal Time (UTC). | [optional] 

## Example

```python
from lago_api.models.wallet_create_input_wallet import WalletCreateInputWallet

# TODO update the JSON string below
json = "{}"
# create an instance of WalletCreateInputWallet from a JSON string
wallet_create_input_wallet_instance = WalletCreateInputWallet.from_json(json)
# print the JSON string representation of the object
print WalletCreateInputWallet.to_json()

# convert the object into a dict
wallet_create_input_wallet_dict = wallet_create_input_wallet_instance.to_dict()
# create an instance of WalletCreateInputWallet from a dict
wallet_create_input_wallet_form_dict = wallet_create_input_wallet.from_dict(wallet_create_input_wallet_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


