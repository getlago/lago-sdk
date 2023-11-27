# PlansPaginated


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**plans** | [**List[PlanObject]**](PlanObject.md) |  | 
**meta** | [**PaginationMeta**](PaginationMeta.md) |  | 

## Example

```python
from lago_api.models.plans_paginated import PlansPaginated

# TODO update the JSON string below
json = "{}"
# create an instance of PlansPaginated from a JSON string
plans_paginated_instance = PlansPaginated.from_json(json)
# print the JSON string representation of the object
print PlansPaginated.to_json()

# convert the object into a dict
plans_paginated_dict = plans_paginated_instance.to_dict()
# create an instance of PlansPaginated from a dict
plans_paginated_form_dict = plans_paginated.from_dict(plans_paginated_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


