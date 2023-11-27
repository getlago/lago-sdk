# MrrObject


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**month** | **str** | Identifies the month to analyze MRR. | 
**amount_cents** | **int** | The total amount of MRR, expressed in cents. | [optional] 
**currency** | [**Currency**](Currency.md) |  | [optional] 

## Example

```python
from lago_api.models.mrr_object import MrrObject

# TODO update the JSON string below
json = "{}"
# create an instance of MrrObject from a JSON string
mrr_object_instance = MrrObject.from_json(json)
# print the JSON string representation of the object
print MrrObject.to_json()

# convert the object into a dict
mrr_object_dict = mrr_object_instance.to_dict()
# create an instance of MrrObject from a dict
mrr_object_form_dict = mrr_object.from_dict(mrr_object_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


