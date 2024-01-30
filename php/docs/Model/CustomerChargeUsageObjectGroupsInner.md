# # CustomerChargeUsageObjectGroupsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**lago_id** | **string** | Unique identifier assigned to the group within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the group record within the Lago system. | [optional]
**key** | **string** | The group key, only returned for groups with two dimensions. | [optional]
**value** | **string** | The group value. | [optional]
**units** | **string** | The number of units consumed for a specific group related to a charge item. | [optional]
**events_count** | **int** | The quantity of usage events that have been recorded for a particular charge during the specified time period. These events may also be referred to as the number of transactions in some contexts. | [optional]
**amount_cents** | **int** | The amount in cents, tax excluded, consumed for a specific group related to a charge item. | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
