# AddOnObject


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**lago_id** | **str** | Unique identifier of the add-on, created by Lago. | 
**name** | **str** | The name of the add-on. | 
**invoice_display_name** | **str** | Specifies the name that will be displayed on an invoice. If no value is set for this field, the name of the actual charge will be used as the default display name. | [optional] 
**code** | **str** | Unique code used to identify the add-on. | 
**amount_cents** | **int** | The cost of the add-on in cents, excluding any applicable taxes, that is billed to a customer. By creating a one-off invoice, you will be able to override this value. | 
**amount_currency** | [**Currency**](Currency.md) |  | 
**description** | **str** | The description of the add-on. | [optional] 
**created_at** | **datetime** | The date and time when the add-on was created. It is expressed in UTC format according to the ISO 8601 datetime standard. This field provides the timestamp for the exact moment when the add-on was initially created. | 
**taxes** | [**List[TaxObject]**](TaxObject.md) | All taxes applied to the add-on. | [optional] 

## Example

```python
from lago_api.models.add_on_object import AddOnObject

# TODO update the JSON string below
json = "{}"
# create an instance of AddOnObject from a JSON string
add_on_object_instance = AddOnObject.from_json(json)
# print the JSON string representation of the object
print AddOnObject.to_json()

# convert the object into a dict
add_on_object_dict = add_on_object_instance.to_dict()
# create an instance of AddOnObject from a dict
add_on_object_form_dict = add_on_object.from_dict(add_on_object_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


