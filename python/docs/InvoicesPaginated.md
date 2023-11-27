# InvoicesPaginated


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**invoices** | [**List[InvoiceObject]**](InvoiceObject.md) |  | 
**meta** | [**PaginationMeta**](PaginationMeta.md) |  | 

## Example

```python
from lago_api.models.invoices_paginated import InvoicesPaginated

# TODO update the JSON string below
json = "{}"
# create an instance of InvoicesPaginated from a JSON string
invoices_paginated_instance = InvoicesPaginated.from_json(json)
# print the JSON string representation of the object
print InvoicesPaginated.to_json()

# convert the object into a dict
invoices_paginated_dict = invoices_paginated_instance.to_dict()
# create an instance of InvoicesPaginated from a dict
invoices_paginated_form_dict = invoices_paginated.from_dict(invoices_paginated_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


