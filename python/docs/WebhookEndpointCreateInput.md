# WebhookEndpointCreateInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**webhook_endpoint** | [**WebhookEndpointCreateInputWebhookEndpoint**](WebhookEndpointCreateInputWebhookEndpoint.md) |  | [optional] 

## Example

```python
from lago_api.models.webhook_endpoint_create_input import WebhookEndpointCreateInput

# TODO update the JSON string below
json = "{}"
# create an instance of WebhookEndpointCreateInput from a JSON string
webhook_endpoint_create_input_instance = WebhookEndpointCreateInput.from_json(json)
# print the JSON string representation of the object
print WebhookEndpointCreateInput.to_json()

# convert the object into a dict
webhook_endpoint_create_input_dict = webhook_endpoint_create_input_instance.to_dict()
# create an instance of WebhookEndpointCreateInput from a dict
webhook_endpoint_create_input_form_dict = webhook_endpoint_create_input.from_dict(webhook_endpoint_create_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


