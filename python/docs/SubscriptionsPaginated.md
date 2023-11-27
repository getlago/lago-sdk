# SubscriptionsPaginated


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**subscriptions** | [**List[SubscriptionObject]**](SubscriptionObject.md) |  | 
**meta** | [**PaginationMeta**](PaginationMeta.md) |  | 

## Example

```python
from lago_api.models.subscriptions_paginated import SubscriptionsPaginated

# TODO update the JSON string below
json = "{}"
# create an instance of SubscriptionsPaginated from a JSON string
subscriptions_paginated_instance = SubscriptionsPaginated.from_json(json)
# print the JSON string representation of the object
print SubscriptionsPaginated.to_json()

# convert the object into a dict
subscriptions_paginated_dict = subscriptions_paginated_instance.to_dict()
# create an instance of SubscriptionsPaginated from a dict
subscriptions_paginated_form_dict = subscriptions_paginated.from_dict(subscriptions_paginated_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


