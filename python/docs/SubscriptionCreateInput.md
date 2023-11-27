# SubscriptionCreateInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**subscription** | [**SubscriptionCreateInputSubscription**](SubscriptionCreateInputSubscription.md) |  | 

## Example

```python
from lago_api.models.subscription_create_input import SubscriptionCreateInput

# TODO update the JSON string below
json = "{}"
# create an instance of SubscriptionCreateInput from a JSON string
subscription_create_input_instance = SubscriptionCreateInput.from_json(json)
# print the JSON string representation of the object
print SubscriptionCreateInput.to_json()

# convert the object into a dict
subscription_create_input_dict = subscription_create_input_instance.to_dict()
# create an instance of SubscriptionCreateInput from a dict
subscription_create_input_form_dict = subscription_create_input.from_dict(subscription_create_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


