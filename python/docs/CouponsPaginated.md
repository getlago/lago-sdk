# CouponsPaginated


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**coupons** | [**List[CouponObject]**](CouponObject.md) |  | 
**meta** | [**PaginationMeta**](PaginationMeta.md) |  | 

## Example

```python
from lago_api.models.coupons_paginated import CouponsPaginated

# TODO update the JSON string below
json = "{}"
# create an instance of CouponsPaginated from a JSON string
coupons_paginated_instance = CouponsPaginated.from_json(json)
# print the JSON string representation of the object
print CouponsPaginated.to_json()

# convert the object into a dict
coupons_paginated_dict = coupons_paginated_instance.to_dict()
# create an instance of CouponsPaginated from a dict
coupons_paginated_form_dict = coupons_paginated.from_dict(coupons_paginated_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


