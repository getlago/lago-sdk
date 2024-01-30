# # CustomerChargeUsageObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**units** | **string** | The number of units consumed by the customer for a specific charge item. |
**events_count** | **int** | The quantity of usage events that have been recorded for a particular charge during the specified time period. These events may also be referred to as the number of transactions in some contexts. |
**amount_cents** | **int** | The amount in cents, tax excluded, consumed by the customer for a specific charge item. |
**amount_currency** | [**Currency**](Currency.md) |  |
**charge** | [**\OpenAPI\Client\Model\CustomerChargeUsageObjectCharge**](CustomerChargeUsageObjectCharge.md) |  |
**billable_metric** | [**\OpenAPI\Client\Model\CustomerChargeUsageObjectBillableMetric**](CustomerChargeUsageObjectBillableMetric.md) |  |
**groups** | [**\OpenAPI\Client\Model\CustomerChargeUsageObjectGroupsInner[]**](CustomerChargeUsageObjectGroupsInner.md) | Array of group object, representing multiple dimensions for a charge item. |

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
