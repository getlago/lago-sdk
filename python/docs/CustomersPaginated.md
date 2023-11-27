# CustomersPaginated


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**customers** | [**List[CustomerObjectExtended]**](CustomerObjectExtended.md) |  | 
**meta** | [**PaginationMeta**](PaginationMeta.md) |  | 

## Example

```python
from lago_api.models.customers_paginated import CustomersPaginated

# TODO update the JSON string below
json = "{}"
# create an instance of CustomersPaginated from a JSON string
customers_paginated_instance = CustomersPaginated.from_json(json)
# print the JSON string representation of the object
print CustomersPaginated.to_json()

# convert the object into a dict
customers_paginated_dict = customers_paginated_instance.to_dict()
# create an instance of CustomersPaginated from a dict
customers_paginated_form_dict = customers_paginated.from_dict(customers_paginated_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


