# WalletsPaginated


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**wallets** | [**List[WalletObject]**](WalletObject.md) |  | 
**meta** | [**PaginationMeta**](PaginationMeta.md) |  | 

## Example

```python
from lago_api.models.wallets_paginated import WalletsPaginated

# TODO update the JSON string below
json = "{}"
# create an instance of WalletsPaginated from a JSON string
wallets_paginated_instance = WalletsPaginated.from_json(json)
# print the JSON string representation of the object
print WalletsPaginated.to_json()

# convert the object into a dict
wallets_paginated_dict = wallets_paginated_instance.to_dict()
# create an instance of WalletsPaginated from a dict
wallets_paginated_form_dict = wallets_paginated.from_dict(wallets_paginated_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


