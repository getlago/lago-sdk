# CreditObjectItem

The item attached to the credit.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**lago_item_id** | **str** | Unique identifier assigned to the credit item within the Lago application. | 
**type** | **str** | The type of credit applied. Possible values are &#x60;coupon&#x60; or &#x60;credit_note&#x60;. | 
**code** | **str** | The code of the credit applied. It can be the code of the coupon attached to the credit or the credit note&#39;s number. | 
**name** | **str** | The name of the credit applied. It can be the name of the coupon attached to the credit or the initial invoice&#39;s number linked to the credit note. | 

## Example

```python
from lago_api.models.credit_object_item import CreditObjectItem

# TODO update the JSON string below
json = "{}"
# create an instance of CreditObjectItem from a JSON string
credit_object_item_instance = CreditObjectItem.from_json(json)
# print the JSON string representation of the object
print CreditObjectItem.to_json()

# convert the object into a dict
credit_object_item_dict = credit_object_item_instance.to_dict()
# create an instance of CreditObjectItem from a dict
credit_object_item_form_dict = credit_object_item.from_dict(credit_object_item_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


