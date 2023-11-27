# CustomerUsageObject


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**from_datetime** | **datetime** | The lower bound of the billing period, expressed in the ISO 8601 datetime format in Coordinated Universal Time (UTC). | 
**to_datetime** | **datetime** | The upper bound of the billing period, expressed in the ISO 8601 datetime format in Coordinated Universal Time (UTC). | 
**issuing_date** | **datetime** | The date of creation of the invoice. | 
**lago_invoice_id** | **str** | A unique identifier associated with the invoice related to this particular usage record. | [optional] 
**currency** | [**Currency**](Currency.md) |  | [optional] 
**amount_cents** | **int** | The amount in cents, tax excluded. | 
**taxes_amount_cents** | **int** | The tax amount in cents. | 
**total_amount_cents** | **int** | The total amount in cents, tax included. | 
**charges_usage** | [**List[CustomerChargeUsageObject]**](CustomerChargeUsageObject.md) | Array of charges that comprise the current usage. It contains detailed information about individual charge items associated with the usage. | 

## Example

```python
from lago_api.models.customer_usage_object import CustomerUsageObject

# TODO update the JSON string below
json = "{}"
# create an instance of CustomerUsageObject from a JSON string
customer_usage_object_instance = CustomerUsageObject.from_json(json)
# print the JSON string representation of the object
print CustomerUsageObject.to_json()

# convert the object into a dict
customer_usage_object_dict = customer_usage_object_instance.to_dict()
# create an instance of CustomerUsageObject from a dict
customer_usage_object_form_dict = customer_usage_object.from_dict(customer_usage_object_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


