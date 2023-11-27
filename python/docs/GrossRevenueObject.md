# GrossRevenueObject


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**month** | **str** | Identifies the month to analyze revenue. | 
**amount_cents** | **int** | The total amount of revenue for a period, expressed in cents. | [optional] 
**currency** | [**Currency**](Currency.md) |  | [optional] 

## Example

```python
from lago_api.models.gross_revenue_object import GrossRevenueObject

# TODO update the JSON string below
json = "{}"
# create an instance of GrossRevenueObject from a JSON string
gross_revenue_object_instance = GrossRevenueObject.from_json(json)
# print the JSON string representation of the object
print GrossRevenueObject.to_json()

# convert the object into a dict
gross_revenue_object_dict = gross_revenue_object_instance.to_dict()
# create an instance of GrossRevenueObject from a dict
gross_revenue_object_form_dict = gross_revenue_object.from_dict(gross_revenue_object_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


