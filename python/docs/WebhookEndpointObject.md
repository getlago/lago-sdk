# WebhookEndpointObject


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**lago_id** | **str** | Unique identifier assigned to the wallet within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the webhook endpoint&#39;s record within the Lago system. | 
**lago_organization_id** | **str** | Unique identifier assigned to the organization attached to the webhook endpoint within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the organization’s record within the Lago system. | 
**webhook_url** | **str** | The name of the wallet. | 
**signature_algo** | **str** | The signature algo for the webhook. | [optional] 
**created_at** | **datetime** | The date of the webhook endpoint creation, represented in ISO 8601 datetime format and expressed in Coordinated Universal Time (UTC). | 

## Example

```python
from lago_api.models.webhook_endpoint_object import WebhookEndpointObject

# TODO update the JSON string below
json = "{}"
# create an instance of WebhookEndpointObject from a JSON string
webhook_endpoint_object_instance = WebhookEndpointObject.from_json(json)
# print the JSON string representation of the object
print WebhookEndpointObject.to_json()

# convert the object into a dict
webhook_endpoint_object_dict = webhook_endpoint_object_instance.to_dict()
# create an instance of WebhookEndpointObject from a dict
webhook_endpoint_object_form_dict = webhook_endpoint_object.from_dict(webhook_endpoint_object_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


