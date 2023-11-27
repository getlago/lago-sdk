# FeeObjectItem

Item attached to the fee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**type** | **str** | The fee type. Possible values are &#x60;add-on&#x60;, &#x60;charge&#x60;, &#x60;credit&#x60; or &#x60;subscription&#x60;. | 
**code** | **str** | The code of the fee item. It can be the code of the &#x60;add-o&#x60;n, the code of the &#x60;charge&#x60;, the code of the &#x60;credit&#x60; or the code of the &#x60;subscription&#x60;. | 
**name** | **str** | The name of the fee item. It can be the name of the &#x60;add-on&#x60;, the name of the &#x60;charge&#x60;, the name of the &#x60;credit&#x60; or the name of the &#x60;subscription&#x60;. | 
**invoice_display_name** | **str** | Specifies the name that will be displayed on an invoice. If no value is set for this field, the name of the actual charge will be used as the default display name. | [optional] 
**group_invoice_display_name** | **str** | Specifies the name that will be displayed on an invoice. If no value is set for this field, the name of the actual charge will be used as the default display name. | [optional] 
**lago_item_id** | **str** | Unique identifier of the fee item, created by Lago. It can be the identifier of the &#x60;add-on&#x60;, the identifier of the &#x60;charge&#x60;, the identifier of the &#x60;credit&#x60; or the identifier of the &#x60;subscription&#x60;. | 
**item_type** | **str** | The type of the fee item. Possible values are &#x60;AddOn&#x60;, &#x60;BillableMetric&#x60;, &#x60;WalletTransaction&#x60; or &#x60;Subscription&#x60;. | 

## Example

```python
from lago_api.models.fee_object_item import FeeObjectItem

# TODO update the JSON string below
json = "{}"
# create an instance of FeeObjectItem from a JSON string
fee_object_item_instance = FeeObjectItem.from_json(json)
# print the JSON string representation of the object
print FeeObjectItem.to_json()

# convert the object into a dict
fee_object_item_dict = fee_object_item_instance.to_dict()
# create an instance of FeeObjectItem from a dict
fee_object_item_form_dict = fee_object_item.from_dict(fee_object_item_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


