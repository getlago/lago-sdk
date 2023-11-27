# WalletTransactionsPaginated


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**wallet_transactions** | [**List[WalletTransactionObject]**](WalletTransactionObject.md) |  | 
**meta** | [**PaginationMeta**](PaginationMeta.md) |  | 

## Example

```python
from lago_api.models.wallet_transactions_paginated import WalletTransactionsPaginated

# TODO update the JSON string below
json = "{}"
# create an instance of WalletTransactionsPaginated from a JSON string
wallet_transactions_paginated_instance = WalletTransactionsPaginated.from_json(json)
# print the JSON string representation of the object
print WalletTransactionsPaginated.to_json()

# convert the object into a dict
wallet_transactions_paginated_dict = wallet_transactions_paginated_instance.to_dict()
# create an instance of WalletTransactionsPaginated from a dict
wallet_transactions_paginated_form_dict = wallet_transactions_paginated.from_dict(wallet_transactions_paginated_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


