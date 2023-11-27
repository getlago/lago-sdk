# AddOnsPaginated


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**add_ons** | [**List[AddOnObject]**](AddOnObject.md) |  | 
**meta** | [**PaginationMeta**](PaginationMeta.md) |  | 

## Example

```python
from lago_api.models.add_ons_paginated import AddOnsPaginated

# TODO update the JSON string below
json = "{}"
# create an instance of AddOnsPaginated from a JSON string
add_ons_paginated_instance = AddOnsPaginated.from_json(json)
# print the JSON string representation of the object
print AddOnsPaginated.to_json()

# convert the object into a dict
add_ons_paginated_dict = add_ons_paginated_instance.to_dict()
# create an instance of AddOnsPaginated from a dict
add_ons_paginated_form_dict = add_ons_paginated.from_dict(add_ons_paginated_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


