# ChargeProperties


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**graduated_ranges** | [**List[ChargePropertiesGraduatedRangesInner]**](ChargePropertiesGraduatedRangesInner.md) | Graduated ranges, sorted from bottom to top tiers, used for a &#x60;graduated&#x60; charge model. | [optional] 
**graduated_percentage_ranges** | [**List[ChargePropertiesGraduatedPercentageRangesInner]**](ChargePropertiesGraduatedPercentageRangesInner.md) | Graduated percentage ranges, sorted from bottom to top tiers, used for a &#x60;graduated_percentage&#x60; charge model. | [optional] 
**amount** | **str** | - The unit price, excluding tax, for a &#x60;standard&#x60; charge model. It is expressed as a decimal value. - The amount, excluding tax, for a complete set of units in a &#x60;package&#x60; charge model. It is expressed as a decimal value. | [optional] 
**free_units** | **int** | The quantity of units that are provided free of charge for each billing period in a &#x60;package&#x60; charge model. This field specifies the number of units that customers can use without incurring any additional cost during each billing cycle. | [optional] 
**package_size** | **int** | The quantity of units included in each pack or set for a &#x60;package&#x60; charge model. It indicates the number of units that are bundled together as a single package or set within the pricing structure. | [optional] 
**rate** | **str** | The percentage rate that is applied to the amount of each transaction for a &#x60;percentage&#x60; charge model. It is expressed as a decimal value. | [optional] 
**fixed_amount** | **str** | The fixed fee that is applied to each transaction for a &#x60;percentage&#x60; charge model. It is expressed as a decimal value. | [optional] 
**free_units_per_events** | **int** | The count of transactions that are not impacted by the &#x60;percentage&#x60; rate and fixed fee in a percentage charge model. This field indicates the number of transactions that are exempt from the calculation of charges based on the specified percentage rate and fixed fee. | [optional] 
**free_units_per_total_aggregation** | **str** | The transaction amount that is not impacted by the &#x60;percentage&#x60; rate and fixed fee in a percentage charge model. This field indicates the portion of the transaction amount that is exempt from the calculation of charges based on the specified percentage rate and fixed fee. | [optional] 
**per_transaction_max_amount** | **str** | Specifies the maximum allowable spending for a single transaction. Working as a transaction cap. | [optional] 
**per_transaction_min_amount** | **str** | Specifies the minimum allowable spending for a single transaction. Working as a transaction floor. | [optional] 
**volume_ranges** | [**List[ChargePropertiesVolumeRangesInner]**](ChargePropertiesVolumeRangesInner.md) | Volume ranges, sorted from bottom to top tiers, used for a &#x60;volume&#x60; charge model. | [optional] 

## Example

```python
from lago_api.models.charge_properties import ChargeProperties

# TODO update the JSON string below
json = "{}"
# create an instance of ChargeProperties from a JSON string
charge_properties_instance = ChargeProperties.from_json(json)
# print the JSON string representation of the object
print ChargeProperties.to_json()

# convert the object into a dict
charge_properties_dict = charge_properties_instance.to_dict()
# create an instance of ChargeProperties from a dict
charge_properties_form_dict = charge_properties.from_dict(charge_properties_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


