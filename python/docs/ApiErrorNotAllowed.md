# ApiErrorNotAllowed


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**status** | **int** |  | 
**error** | **str** |  | 
**code** | **str** |  | 

## Example

```python
from lago_api.models.api_error_not_allowed import ApiErrorNotAllowed

# TODO update the JSON string below
json = "{}"
# create an instance of ApiErrorNotAllowed from a JSON string
api_error_not_allowed_instance = ApiErrorNotAllowed.from_json(json)
# print the JSON string representation of the object
print ApiErrorNotAllowed.to_json()

# convert the object into a dict
api_error_not_allowed_dict = api_error_not_allowed_instance.to_dict()
# create an instance of ApiErrorNotAllowed from a dict
api_error_not_allowed_form_dict = api_error_not_allowed.from_dict(api_error_not_allowed_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


