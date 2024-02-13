# LagoAPI::FeeObjectAmountDetailsAllOfVolumeRangesInner

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **per_unit_amount** | **String** | The flat amount for a whole tier, excluding tax, for a &#x60;volume&#x60; charge model. |  |
| **flat_unit_amount** | **String** | The unit price, excluding tax, for a specific tier of a &#x60;volume&#x60; charge model. |  |
| **per_unit_total_amount** | **String** | Total amount of received units to be charged. |  |

## Example

```ruby
require 'lago_ruby'

instance = LagoAPI::FeeObjectAmountDetailsAllOfVolumeRangesInner.new(
  per_unit_amount: 0.5,
  flat_unit_amount: 10.0,
  per_unit_total_amount: 10.0
)
```

