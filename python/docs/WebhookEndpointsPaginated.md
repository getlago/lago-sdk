# WebhookEndpointsPaginated


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**webhook_endpoints** | [**List[WebhookEndpointObject]**](WebhookEndpointObject.md) |  | 
**meta** | [**PaginationMeta**](PaginationMeta.md) |  | 

## Example

```python
from lago_api.models.webhook_endpoints_paginated import WebhookEndpointsPaginated

# TODO update the JSON string below
json = "{}"
# create an instance of WebhookEndpointsPaginated from a JSON string
webhook_endpoints_paginated_instance = WebhookEndpointsPaginated.from_json(json)
# print the JSON string representation of the object
print WebhookEndpointsPaginated.to_json()

# convert the object into a dict
webhook_endpoints_paginated_dict = webhook_endpoints_paginated_instance.to_dict()
# create an instance of WebhookEndpointsPaginated from a dict
webhook_endpoints_paginated_form_dict = webhook_endpoints_paginated.from_dict(webhook_endpoints_paginated_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


