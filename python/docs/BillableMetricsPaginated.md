# BillableMetricsPaginated


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**billable_metrics** | [**List[BillableMetricObject]**](BillableMetricObject.md) |  | 
**meta** | [**PaginationMeta**](PaginationMeta.md) |  | 

## Example

```python
from lago_api.models.billable_metrics_paginated import BillableMetricsPaginated

# TODO update the JSON string below
json = "{}"
# create an instance of BillableMetricsPaginated from a JSON string
billable_metrics_paginated_instance = BillableMetricsPaginated.from_json(json)
# print the JSON string representation of the object
print BillableMetricsPaginated.to_json()

# convert the object into a dict
billable_metrics_paginated_dict = billable_metrics_paginated_instance.to_dict()
# create an instance of BillableMetricsPaginated from a dict
billable_metrics_paginated_form_dict = billable_metrics_paginated.from_dict(billable_metrics_paginated_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


