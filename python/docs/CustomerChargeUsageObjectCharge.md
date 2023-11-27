# CustomerChargeUsageObjectCharge

Object listing all the properties for a specific charge item.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**lago_id** | **str** | Unique identifier assigned to the charge within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the charge’s record within the Lago system. | 
**charge_model** | **str** | The pricing model applied to this charge. Possible values are standard, &#x60;graduated&#x60;, &#x60;percentage&#x60;, &#x60;package&#x60; or &#x60;volume&#x60;. | 
**invoice_display_name** | **str** | Specifies the name that will be displayed on an invoice. If no value is set for this field, the name of the actual charge will be used as the default display name. | [optional] 

## Example

```python
from lago_api.models.customer_charge_usage_object_charge import CustomerChargeUsageObjectCharge

# TODO update the JSON string below
json = "{}"
# create an instance of CustomerChargeUsageObjectCharge from a JSON string
customer_charge_usage_object_charge_instance = CustomerChargeUsageObjectCharge.from_json(json)
# print the JSON string representation of the object
print CustomerChargeUsageObjectCharge.to_json()

# convert the object into a dict
customer_charge_usage_object_charge_dict = customer_charge_usage_object_charge_instance.to_dict()
# create an instance of CustomerChargeUsageObjectCharge from a dict
customer_charge_usage_object_charge_form_dict = customer_charge_usage_object_charge.from_dict(customer_charge_usage_object_charge_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


