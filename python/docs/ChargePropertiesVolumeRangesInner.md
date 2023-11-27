# ChargePropertiesVolumeRangesInner


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**from_value** | **int** | Specifies the lower value of a tier for a &#x60;volume&#x60; charge model. It must be either 0 or the previous range&#39;s &#x60;to_value + 1&#x60; to maintain the proper sequence of values. | 
**to_value** | **int** | Specifies the highest value of a tier for a &#x60;volume&#x60; charge model. - This value must be higher than the &#x60;from_value&#x60; of the same tier. - This value must be &#x60;null&#x60; for the last tier. | 
**flat_amount** | **str** | The unit price, excluding tax, for a specific tier of a &#x60;volume&#x60; charge model. It is expressed as a decimal value. | 
**per_unit_amount** | **str** | The flat amount for a whole tier, excluding tax, for a &#x60;volume&#x60; charge model. It is expressed as a decimal value. | 

## Example

```python
from lago_api.models.charge_properties_volume_ranges_inner import ChargePropertiesVolumeRangesInner

# TODO update the JSON string below
json = "{}"
# create an instance of ChargePropertiesVolumeRangesInner from a JSON string
charge_properties_volume_ranges_inner_instance = ChargePropertiesVolumeRangesInner.from_json(json)
# print the JSON string representation of the object
print ChargePropertiesVolumeRangesInner.to_json()

# convert the object into a dict
charge_properties_volume_ranges_inner_dict = charge_properties_volume_ranges_inner_instance.to_dict()
# create an instance of ChargePropertiesVolumeRangesInner from a dict
charge_properties_volume_ranges_inner_form_dict = charge_properties_volume_ranges_inner.from_dict(charge_properties_volume_ranges_inner_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


