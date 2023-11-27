# SubscriptionUpdateInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**subscription** | [**SubscriptionUpdateInputSubscription**](SubscriptionUpdateInputSubscription.md) |  | 

## Example

```python
from lago_api.models.subscription_update_input import SubscriptionUpdateInput

# TODO update the JSON string below
json = "{}"
# create an instance of SubscriptionUpdateInput from a JSON string
subscription_update_input_instance = SubscriptionUpdateInput.from_json(json)
# print the JSON string representation of the object
print SubscriptionUpdateInput.to_json()

# convert the object into a dict
subscription_update_input_dict = subscription_update_input_instance.to_dict()
# create an instance of SubscriptionUpdateInput from a dict
subscription_update_input_form_dict = subscription_update_input.from_dict(subscription_update_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


