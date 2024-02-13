# LagoAPI::FeeObjectAmountDetails

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **graduated_ranges** | [**Array&lt;FeeObjectAmountDetailsAllOfGraduatedRangesInner&gt;**](FeeObjectAmountDetailsAllOfGraduatedRangesInner.md) | Graduated ranges, used for a &#x60;graduated&#x60; charge model. | [optional] |
| **graduated_percentage_ranges** | [**Array&lt;FeeObjectAmountDetailsAllOfGraduatedPercentageRangesInner&gt;**](FeeObjectAmountDetailsAllOfGraduatedPercentageRangesInner.md) | Graduated percentage ranges, used for a &#x60;graduated_percentage&#x60; charge model. | [optional] |
| **free_units** | **String** | The quantity of units that are provided free of charge for each billing period in a &#x60;package&#x60; charge model. | [optional] |
| **paid_units** | **String** | The quantity of units that are not provided free of charge for each billing period in a &#x60;package&#x60; charge model. | [optional] |
| **per_package_size** | **Integer** | The quantity of units included, defined for Package or Percentage charge model. | [optional] |
| **per_package_unit_amount** | **String** | Total amount to charge for received paid_units, defined for Package or Percentage charge model. | [optional] |
| **units** | **String** | The total units received in Lago for the Percentage charge model. | [optional] |
| **free_events** | **Integer** | Total number of free events allowed for the Percentage charge model. | [optional] |
| **rate** | **String** | Percentage rate applied for the Percentage charge model. | [optional] |
| **per_unit_total_amount** | **String** | Total amount of received units to be charged for the Percentage charge model. | [optional] |
| **paid_events** | **Integer** | Total number of paid events for the Percentage charge model. | [optional] |
| **fixed_fee_unit_amount** | **String** | Fixed fee unit price per received paid_event for the Percentage charge model. | [optional] |
| **fixed_fee_total_amount** | **String** | Total amount to charge for received paid_events for the Percentage charge model. | [optional] |
| **min_max_adjustment_total_amount** | **String** | Total adjustment amount linked to minimum and maximum spending per transaction for the Percentage charge model. | [optional] |
| **volume_ranges** | [**Array&lt;FeeObjectAmountDetailsAllOfVolumeRangesInner&gt;**](FeeObjectAmountDetailsAllOfVolumeRangesInner.md) | Volume ranges, used for a &#x60;volume&#x60; charge model. | [optional] |

## Example

```ruby
require 'lago_ruby'

instance = LagoAPI::FeeObjectAmountDetails.new(
  graduated_ranges: null,
  graduated_percentage_ranges: null,
  free_units: 10.0,
  paid_units: 40.0,
  per_package_size: 1000,
  per_package_unit_amount: 0.5,
  units: 20.0,
  free_events: 10,
  rate: 1.0,
  per_unit_total_amount: 10.0,
  paid_events: 20,
  fixed_fee_unit_amount: 1.0,
  fixed_fee_total_amount: 20.0,
  min_max_adjustment_total_amount: 20.0,
  volume_ranges: null
)
```

