# LagoAPI::CustomerChargeUsageObject

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **units** | **String** | The number of units consumed by the customer for a specific charge item. |  |
| **events_count** | **Integer** | The quantity of usage events that have been recorded for a particular charge during the specified time period. These events may also be referred to as the number of transactions in some contexts. |  |
| **amount_cents** | **Integer** | The amount in cents, tax excluded, consumed by the customer for a specific charge item. |  |
| **amount_currency** | [**Currency**](Currency.md) |  |  |
| **charge** | [**CustomerChargeUsageObjectCharge**](CustomerChargeUsageObjectCharge.md) |  |  |
| **billable_metric** | [**CustomerChargeUsageObjectBillableMetric**](CustomerChargeUsageObjectBillableMetric.md) |  |  |
| **groups** | [**Array&lt;CustomerChargeGroupsUsageObjectInner&gt;**](CustomerChargeGroupsUsageObjectInner.md) | Array of group object, representing multiple dimensions for a charge item. |  |
| **grouped_usage** | [**Array&lt;CustomerChargeGroupedUsageObjectInner&gt;**](CustomerChargeGroupedUsageObjectInner.md) | Array of aggregated fees, grouped by the event properties defined in a &#x60;standard&#x60; charge model. | [optional] |

## Example

```ruby
require 'lago_ruby'

instance = LagoAPI::CustomerChargeUsageObject.new(
  units: 1.0,
  events_count: 10,
  amount_cents: 123,
  amount_currency: null,
  charge: null,
  billable_metric: null,
  groups: null,
  grouped_usage: null
)
```

