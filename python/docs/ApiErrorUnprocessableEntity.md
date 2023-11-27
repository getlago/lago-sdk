# ApiErrorUnprocessableEntity


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**status** | **int** |  | 
**error** | **str** |  | 
**code** | **str** |  | 
**error_details** | **object** |  | 

## Example

```python
from lago_api.models.api_error_unprocessable_entity import ApiErrorUnprocessableEntity

# TODO update the JSON string below
json = "{}"
# create an instance of ApiErrorUnprocessableEntity from a JSON string
api_error_unprocessable_entity_instance = ApiErrorUnprocessableEntity.from_json(json)
# print the JSON string representation of the object
print ApiErrorUnprocessableEntity.to_json()

# convert the object into a dict
api_error_unprocessable_entity_dict = api_error_unprocessable_entity_instance.to_dict()
# create an instance of ApiErrorUnprocessableEntity from a dict
api_error_unprocessable_entity_form_dict = api_error_unprocessable_entity.from_dict(api_error_unprocessable_entity_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


