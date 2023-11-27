# WalletObject


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**lago_id** | **str** | Unique identifier assigned to the wallet within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the wallet’s record within the Lago system. | 
**lago_customer_id** | **str** | Unique identifier assigned to the customer within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the customer’s record within the Lago system. | 
**external_customer_id** | **str** | The customer external unique identifier (provided by your own application) | 
**status** | **str** | The status of the wallet. Possible values are &#x60;active&#x60; or &#x60;terminated&#x60;. | 
**currency** | [**Currency**](Currency.md) |  | 
**name** | **str** | The name of the wallet. | [optional] 
**rate_amount** | **str** | The rate of conversion between credits and the amount in the specified currency. It indicates the ratio or factor used to convert credits into the corresponding monetary value in the currency of the transaction. | 
**credits_balance** | **str** | The current wallet balance expressed in credits. | 
**balance_cents** | **int** | The current wallet balance expressed in cents. | 
**consumed_credits** | **str** | The number of consumed credits. | 
**created_at** | **datetime** | The date of the wallet creation, represented in ISO 8601 datetime format and expressed in Coordinated Universal Time (UTC). | 
**expiration_at** | **datetime** | The date and time that determines when the wallet will expire. It follows the ISO 8601 datetime format and is expressed in Coordinated Universal Time (UTC). | [optional] 
**last_balance_sync_at** | **datetime** | The date and time of the last balance top-up. It follows the ISO 8601 datetime format and is expressed in Coordinated Universal Time (UTC). | [optional] 
**last_consumed_credit_at** | **datetime** | The date and time of the last credits consumption. It follows the ISO 8601 datetime format and is expressed in Coordinated Universal Time (UTC). | [optional] 
**terminated_at** | **datetime** | The date of terminaison of the wallet. It follows the ISO 8601 datetime format and is expressed in Coordinated Universal Time (UTC). | [optional] 

## Example

```python
from lago_api.models.wallet_object import WalletObject

# TODO update the JSON string below
json = "{}"
# create an instance of WalletObject from a JSON string
wallet_object_instance = WalletObject.from_json(json)
# print the JSON string representation of the object
print WalletObject.to_json()

# convert the object into a dict
wallet_object_dict = wallet_object_instance.to_dict()
# create an instance of WalletObject from a dict
wallet_object_form_dict = wallet_object.from_dict(wallet_object_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


