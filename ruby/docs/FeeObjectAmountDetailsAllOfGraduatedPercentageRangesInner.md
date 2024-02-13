# LagoAPI::FeeObjectAmountDetailsAllOfGraduatedPercentageRangesInner

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **units** | **String** | Total units received in Lago. |  |
| **from_value** | **Integer** | Lower value of a tier. It is either 0 or the previous range&#39;s &#x60;to_value + 1&#x60;. |  |
| **to_value** | **Integer** | Highest value of a tier. - This value is higher than the from_value of the same tier. - This value is null for the last tier. |  |
| **flat_unit_amount** | **String** | Flat unit amount within a specified tier. |  |
| **rate** | **String** | Percentage rate applied within a specified tier. |  |
| **per_unit_total_amount** | **String** | Total amount of received units to be charged within a specified tier. |  |
| **total_with_flat_amount** | **String** | Total amount to be charged for a specific tier, taking into account the flat_unit_amount and the per_unit_total_amount. |  |

## Example

```ruby
require 'lago_ruby'

instance = LagoAPI::FeeObjectAmountDetailsAllOfGraduatedPercentageRangesInner.new(
  units: 10.0,
  from_value: 0,
  to_value: 10,
  flat_unit_amount: 1.0,
  rate: 1.0,
  per_unit_total_amount: 10.0,
  total_with_flat_amount: 11.0
)
```

