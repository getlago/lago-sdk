# LagoAPI::WalletUpdateInputWalletRecurringTransactionRulesInner

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **lago_id** | **String** | A unique identifier for the recurring transaction rule in the Lago application. Can be used to update a recurring transaction rule. | [optional] |
| **rule_type** | **String** | The rule type. Possible values are &#x60;interval&#x60; or &#x60;threshold&#x60;. | [optional] |
| **interval** | **String** | The interval used for recurring top-up. It represents the frequency at which automatic top-up occurs. The interval can be one of the following values: &#x60;weekly&#x60;, &#x60;monthly&#x60;, &#x60;quarterly&#x60; or &#x60;yearly&#x60;. Required only when rule type is set to &#x60;interval&#x60;. | [optional] |
| **threshold_credits** | **String** | The threshold for recurring top-ups is the value at which an automatic top-up is triggered. For example, if this threshold is set at 10 credits, an automatic top-up will occur whenever the wallet balance falls to or below 10 credits. Required only when rule type is set to &#x60;threshold&#x60;. | [optional] |
| **paid_credits** | **String** | The number of paid credits. Required only if there is no granted credits. | [optional] |
| **granted_credits** | **String** | The number of free granted credits. Required only if there is no paid credits. | [optional] |

## Example

```ruby
require 'lago_ruby'

instance = LagoAPI::WalletUpdateInputWalletRecurringTransactionRulesInner.new(
  lago_id: 1a901a90-1a90-1a90-1a90-1a901a901a90,
  rule_type: interval,
  interval: monthly,
  threshold_credits: 20.0,
  paid_credits: 20.0,
  granted_credits: 10.0
)
```

