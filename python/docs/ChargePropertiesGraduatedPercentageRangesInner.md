# ChargePropertiesGraduatedPercentageRangesInner


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**from_value** | **int** | Specifies the lower value of a tier for a &#x60;graduated_percentage&#x60; charge model. It must be either 0 or the previous range&#39;s &#x60;to_value + 1&#x60; to maintain the proper sequence of values. | 
**to_value** | **int** | Specifies the highest value of a tier for a &#x60;graduated_percentage&#x60; charge model. - This value must be higher than the from_value of the same tier. - This value must be null for the last tier. | 
**rate** | **str** | The percentage rate that is applied to the amount of each transaction in the tier for a &#x60;graduated_percentage&#x60; charge model. It is expressed as a decimal value. | 
**flat_amount** | **str** | The flat amount for a whole tier, excluding tax, for a &#x60;graduated_percentage&#x60; charge model. It is expressed as a decimal value. | 

## Example

```python
from lago_api.models.charge_properties_graduated_percentage_ranges_inner import ChargePropertiesGraduatedPercentageRangesInner

# TODO update the JSON string below
json = "{}"
# create an instance of ChargePropertiesGraduatedPercentageRangesInner from a JSON string
charge_properties_graduated_percentage_ranges_inner_instance = ChargePropertiesGraduatedPercentageRangesInner.from_json(json)
# print the JSON string representation of the object
print ChargePropertiesGraduatedPercentageRangesInner.to_json()

# convert the object into a dict
charge_properties_graduated_percentage_ranges_inner_dict = charge_properties_graduated_percentage_ranges_inner_instance.to_dict()
# create an instance of ChargePropertiesGraduatedPercentageRangesInner from a dict
charge_properties_graduated_percentage_ranges_inner_form_dict = charge_properties_graduated_percentage_ranges_inner.from_dict(charge_properties_graduated_percentage_ranges_inner_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


