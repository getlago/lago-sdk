# CustomerChargeUsageObject


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**units** | **str** | The number of units consumed by the customer for a specific charge item. | 
**events_count** | **int** | The quantity of usage events that have been recorded for a particular charge during the specified time period. These events may also be referred to as the number of transactions in some contexts. | 
**amount_cents** | **int** | The amount in cents, tax excluded, consumed by the customer for a specific charge item. | 
**amount_currency** | [**Currency**](Currency.md) |  | 
**charge** | [**CustomerChargeUsageObjectCharge**](CustomerChargeUsageObjectCharge.md) |  | 
**billable_metric** | [**CustomerChargeUsageObjectBillableMetric**](CustomerChargeUsageObjectBillableMetric.md) |  | 
**groups** | [**List[CustomerChargeUsageObjectGroupsInner]**](CustomerChargeUsageObjectGroupsInner.md) | Array of group object, representing multiple dimensions for a charge item. | 

## Example

```python
from lago_api.models.customer_charge_usage_object import CustomerChargeUsageObject

# TODO update the JSON string below
json = "{}"
# create an instance of CustomerChargeUsageObject from a JSON string
customer_charge_usage_object_instance = CustomerChargeUsageObject.from_json(json)
# print the JSON string representation of the object
print CustomerChargeUsageObject.to_json()

# convert the object into a dict
customer_charge_usage_object_dict = customer_charge_usage_object_instance.to_dict()
# create an instance of CustomerChargeUsageObject from a dict
customer_charge_usage_object_form_dict = customer_charge_usage_object.from_dict(customer_charge_usage_object_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


