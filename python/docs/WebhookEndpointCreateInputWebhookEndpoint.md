# WebhookEndpointCreateInputWebhookEndpoint


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**webhook_url** | **str** | The URL of the webhook endpoint. | 
**signature_algo** | **str** | The signature used for the webhook. If no value is passed, | [optional] 

## Example

```python
from lago_api.models.webhook_endpoint_create_input_webhook_endpoint import WebhookEndpointCreateInputWebhookEndpoint

# TODO update the JSON string below
json = "{}"
# create an instance of WebhookEndpointCreateInputWebhookEndpoint from a JSON string
webhook_endpoint_create_input_webhook_endpoint_instance = WebhookEndpointCreateInputWebhookEndpoint.from_json(json)
# print the JSON string representation of the object
print WebhookEndpointCreateInputWebhookEndpoint.to_json()

# convert the object into a dict
webhook_endpoint_create_input_webhook_endpoint_dict = webhook_endpoint_create_input_webhook_endpoint_instance.to_dict()
# create an instance of WebhookEndpointCreateInputWebhookEndpoint from a dict
webhook_endpoint_create_input_webhook_endpoint_form_dict = webhook_endpoint_create_input_webhook_endpoint.from_dict(webhook_endpoint_create_input_webhook_endpoint_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


