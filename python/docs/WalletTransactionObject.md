# WalletTransactionObject


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**lago_id** | **str** | Unique identifier assigned to the wallet transaction within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the wallet transaction’s record within the Lago system. | 
**lago_wallet_id** | **str** | Unique identifier assigned to the wallet within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the wallet’s record within the Lago system. | 
**status** | **str** | The status of the wallet transaction. Possible values are &#x60;pending&#x60; or &#x60;settled&#x60;. | 
**transaction_type** | **str** | The type of transaction. Possible values are &#x60;inbound&#x60; (increasing the balance) or &#x60;outbound&#x60; (decreasing the balance). | 
**amount** | **str** | The amount of credits based on the rate and the currency. | 
**credit_amount** | **str** | The number of credits used in the wallet transaction. | 
**settled_at** | **datetime** | The date when wallet transaction is settled, represented in ISO 8601 datetime format and expressed in Coordinated Universal Time (UTC). | [optional] 
**created_at** | **datetime** | The date of the wallet transaction creation, represented in ISO 8601 datetime format and expressed in Coordinated Universal Time (UTC). | 

## Example

```python
from lago_api.models.wallet_transaction_object import WalletTransactionObject

# TODO update the JSON string below
json = "{}"
# create an instance of WalletTransactionObject from a JSON string
wallet_transaction_object_instance = WalletTransactionObject.from_json(json)
# print the JSON string representation of the object
print WalletTransactionObject.to_json()

# convert the object into a dict
wallet_transaction_object_dict = wallet_transaction_object_instance.to_dict()
# create an instance of WalletTransactionObject from a dict
wallet_transaction_object_form_dict = wallet_transaction_object.from_dict(wallet_transaction_object_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


