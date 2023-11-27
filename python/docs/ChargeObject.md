# ChargeObject


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**lago_id** | **str** | Unique identifier of charge, created by Lago. | 
**lago_billable_metric_id** | **str** | Unique identifier of the billable metric created by Lago. | 
**billable_metric_code** | **str** | Unique code identifying a billable metric. | 
**invoice_display_name** | **str** | Specifies the name that will be displayed on an invoice. If no value is set for this field, the name of the actual charge will be used as the default display name. | [optional] 
**created_at** | **datetime** | The date and time when the charge was created. It is expressed in UTC format according to the ISO 8601 datetime standard. | 
**charge_model** | **str** | Specifies the pricing model used for the calculation of the final fee. It can be &#x60;standard&#x60;, &#x60;graduated&#x60;, &#x60;graduated_percentage&#x60;, &#x60;package&#x60;, &#x60;percentage&#x60; or &#x60;volume&#x60;. | 
**pay_in_advance** | **bool** | This field determines the billing timing for this specific usage-based charge. When set to &#x60;true&#x60;, the charge is due and invoiced immediately. Conversely, when set to &#x60;false&#x60;, the charge is due and invoiced at the end of each billing period. | [optional] 
**invoiceable** | **bool** | This field specifies whether the charge should be included in a proper invoice. If set to &#x60;false&#x60;, no invoice will be issued for this charge. You can only set it to &#x60;false&#x60; when &#x60;pay_in_advance&#x60; is &#x60;true&#x60;. | [optional] 
**prorated** | **bool** | Specifies whether a charge is prorated based on the remaining number of days in the billing period or billed fully.  - If set to &#x60;true&#x60;, the charge is prorated based on the remaining days in the current billing period. - If set to &#x60;false&#x60;, the charge is billed in full. - If not defined in the request, default value is &#x60;false&#x60;. | [optional] 
**min_amount_cents** | **int** | The minimum spending amount required for the charge, measured in cents and excluding any applicable taxes. It indicates the minimum amount that needs to be charged for each billing period. | [optional] 
**properties** | [**ChargeObjectProperties**](ChargeObjectProperties.md) |  | [optional] 
**group_properties** | [**List[GroupPropertiesObject]**](GroupPropertiesObject.md) | All charge information, sorted by groups. | [optional] 
**taxes** | [**List[TaxObject]**](TaxObject.md) | All taxes applied to the charge. | [optional] 

## Example

```python
from lago_api.models.charge_object import ChargeObject

# TODO update the JSON string below
json = "{}"
# create an instance of ChargeObject from a JSON string
charge_object_instance = ChargeObject.from_json(json)
# print the JSON string representation of the object
print ChargeObject.to_json()

# convert the object into a dict
charge_object_dict = charge_object_instance.to_dict()
# create an instance of ChargeObject from a dict
charge_object_form_dict = charge_object.from_dict(charge_object_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


