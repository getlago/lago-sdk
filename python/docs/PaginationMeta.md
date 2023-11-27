# PaginationMeta


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**current_page** | **int** | Current page. | 
**next_page** | **int** | Next page. | [optional] 
**prev_page** | **int** | Previous page. | [optional] 
**total_pages** | **int** | Total number of pages. | 
**total_count** | **int** | Total number of records. | 

## Example

```python
from lago_api.models.pagination_meta import PaginationMeta

# TODO update the JSON string below
json = "{}"
# create an instance of PaginationMeta from a JSON string
pagination_meta_instance = PaginationMeta.from_json(json)
# print the JSON string representation of the object
print PaginationMeta.to_json()

# convert the object into a dict
pagination_meta_dict = pagination_meta_instance.to_dict()
# create an instance of PaginationMeta from a dict
pagination_meta_form_dict = pagination_meta.from_dict(pagination_meta_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


