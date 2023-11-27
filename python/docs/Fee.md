# Fee


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**fee** | [**FeeObject**](FeeObject.md) |  | 

## Example

```python
from lago_api.models.fee import Fee

# TODO update the JSON string below
json = "{}"
# create an instance of Fee from a JSON string
fee_instance = Fee.from_json(json)
# print the JSON string representation of the object
print Fee.to_json()

# convert the object into a dict
fee_dict = fee_instance.to_dict()
# create an instance of Fee from a dict
fee_form_dict = fee.from_dict(fee_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


