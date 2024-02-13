# LagoAPI::CustomerChargeGroupedUsageObjectInner

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **amount_cents** | **Integer** | The amount in cents, tax excluded, consumed for a specific group related to a charge item. | [optional] |
| **events_count** | **Integer** | The quantity of usage events that have been recorded for a particular charge during the specified time period. These events may also be referred to as the number of transactions in some contexts. | [optional] |
| **units** | **String** | The number of units consumed for a specific group related to a charge item. | [optional] |
| **grouped_by** | **Hash&lt;String, String&gt;** | Key value list of event properties aggregated by the charge model | [optional] |
| **groups** | [**Array&lt;CustomerChargeGroupsUsageObjectInner&gt;**](CustomerChargeGroupsUsageObjectInner.md) | Array of group object, representing multiple dimensions for a charge item. | [optional] |

## Example

```ruby
require 'lago_ruby'

instance = LagoAPI::CustomerChargeGroupedUsageObjectInner.new(
  amount_cents: 1000,
  events_count: 10,
  units: 0.9,
  grouped_by: null,
  groups: null
)
```

