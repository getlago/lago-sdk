# GroupsPaginated


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**groups** | [**List[GroupObject]**](GroupObject.md) |  | 
**meta** | [**PaginationMeta**](PaginationMeta.md) |  | 

## Example

```python
from lago_api.models.groups_paginated import GroupsPaginated

# TODO update the JSON string below
json = "{}"
# create an instance of GroupsPaginated from a JSON string
groups_paginated_instance = GroupsPaginated.from_json(json)
# print the JSON string representation of the object
print GroupsPaginated.to_json()

# convert the object into a dict
groups_paginated_dict = groups_paginated_instance.to_dict()
# create an instance of GroupsPaginated from a dict
groups_paginated_form_dict = groups_paginated.from_dict(groups_paginated_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


