# GrossRevenues


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**gross_revenues** | [**List[GrossRevenueObject]**](GrossRevenueObject.md) |  | 

## Example

```python
from lago_api.models.gross_revenues import GrossRevenues

# TODO update the JSON string below
json = "{}"
# create an instance of GrossRevenues from a JSON string
gross_revenues_instance = GrossRevenues.from_json(json)
# print the JSON string representation of the object
print GrossRevenues.to_json()

# convert the object into a dict
gross_revenues_dict = gross_revenues_instance.to_dict()
# create an instance of GrossRevenues from a dict
gross_revenues_form_dict = gross_revenues.from_dict(gross_revenues_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


