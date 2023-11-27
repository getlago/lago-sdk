# FeesPaginated


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**fees** | [**List[FeeObject]**](FeeObject.md) |  | 
**meta** | [**PaginationMeta**](PaginationMeta.md) |  | 

## Example

```python
from lago_api.models.fees_paginated import FeesPaginated

# TODO update the JSON string below
json = "{}"
# create an instance of FeesPaginated from a JSON string
fees_paginated_instance = FeesPaginated.from_json(json)
# print the JSON string representation of the object
print FeesPaginated.to_json()

# convert the object into a dict
fees_paginated_dict = fees_paginated_instance.to_dict()
# create an instance of FeesPaginated from a dict
fees_paginated_form_dict = fees_paginated.from_dict(fees_paginated_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


