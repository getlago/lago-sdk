# ApiErrorForbidden


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**status** | **int** |  | 
**error** | **str** |  | 
**code** | **str** |  | 

## Example

```python
from lago_api.models.api_error_forbidden import ApiErrorForbidden

# TODO update the JSON string below
json = "{}"
# create an instance of ApiErrorForbidden from a JSON string
api_error_forbidden_instance = ApiErrorForbidden.from_json(json)
# print the JSON string representation of the object
print ApiErrorForbidden.to_json()

# convert the object into a dict
api_error_forbidden_dict = api_error_forbidden_instance.to_dict()
# create an instance of ApiErrorForbidden from a dict
api_error_forbidden_form_dict = api_error_forbidden.from_dict(api_error_forbidden_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


