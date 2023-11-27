# AppliedCouponsPaginated


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**applied_coupons** | [**List[AppliedCouponObjectExtended]**](AppliedCouponObjectExtended.md) |  | 
**meta** | [**PaginationMeta**](PaginationMeta.md) |  | 

## Example

```python
from lago_api.models.applied_coupons_paginated import AppliedCouponsPaginated

# TODO update the JSON string below
json = "{}"
# create an instance of AppliedCouponsPaginated from a JSON string
applied_coupons_paginated_instance = AppliedCouponsPaginated.from_json(json)
# print the JSON string representation of the object
print AppliedCouponsPaginated.to_json()

# convert the object into a dict
applied_coupons_paginated_dict = applied_coupons_paginated_instance.to_dict()
# create an instance of AppliedCouponsPaginated from a dict
applied_coupons_paginated_form_dict = applied_coupons_paginated.from_dict(applied_coupons_paginated_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


