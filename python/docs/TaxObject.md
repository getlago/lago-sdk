# TaxObject


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**lago_id** | **str** | Unique identifier of the tax, created by Lago. | 
**name** | **str** | Name of the tax. | 
**code** | **str** | Unique code used to identify the tax associated with the API request. | 
**description** | **str** | Internal description of the taxe | [optional] 
**rate** | **float** | The percentage rate of the tax | 
**applied_to_organization** | **bool** | Set to &#x60;true&#x60; if the tax is used as one of the organization&#39;s default | 
**add_ons_count** | **int** | Number of add-ons this tax is applied to. | [optional] 
**charges_count** | **int** | Number of charges this tax is applied to. | [optional] 
**customers_count** | **int** | Number of customers this tax is applied to (directly or via the organization&#39;s default). | 
**plans_count** | **int** | Number of plans this tax is applied to. | [optional] 
**created_at** | **datetime** | Creation date of the tax. | 

## Example

```python
from lago_api.models.tax_object import TaxObject

# TODO update the JSON string below
json = "{}"
# create an instance of TaxObject from a JSON string
tax_object_instance = TaxObject.from_json(json)
# print the JSON string representation of the object
print TaxObject.to_json()

# convert the object into a dict
tax_object_dict = tax_object_instance.to_dict()
# create an instance of TaxObject from a dict
tax_object_form_dict = tax_object.from_dict(tax_object_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


