# TaxesPaginated


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**taxes** | [**List[TaxObject]**](TaxObject.md) |  | 
**meta** | [**PaginationMeta**](PaginationMeta.md) |  | 

## Example

```python
from lago_api.models.taxes_paginated import TaxesPaginated

# TODO update the JSON string below
json = "{}"
# create an instance of TaxesPaginated from a JSON string
taxes_paginated_instance = TaxesPaginated.from_json(json)
# print the JSON string representation of the object
print TaxesPaginated.to_json()

# convert the object into a dict
taxes_paginated_dict = taxes_paginated_instance.to_dict()
# create an instance of TaxesPaginated from a dict
taxes_paginated_form_dict = taxes_paginated.from_dict(taxes_paginated_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


