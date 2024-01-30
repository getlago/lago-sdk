# # WalletObjectRecurringTransactionRulesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**lago_id** | **string** | A unique identifier for the recurring transaction rule in the Lago application. Can be used to update a recurring transaction rule. |
**rule_type** | **string** | The rule type. Possible values are &#x60;interval&#x60; or &#x60;threshold&#x60;. |
**interval** | **string** | The interval used for recurring top-up. It represents the frequency at which automatic top-up occurs. The interval can be one of the following values: &#x60;weekly&#x60;, &#x60;monthly&#x60;, &#x60;quarterly&#x60; or &#x60;yearly&#x60;. Required only if rule type is set to &#x60;interval&#x60;. | [optional]
**threshold_credits** | **string** | The threshold for recurring top-ups is the value at which an automatic top-up is triggered. For example, if this threshold is set at 10 credits, an automatic top-up will occur whenever the wallet balance falls to or below 10 credits. Required only when rule type is set to &#x60;threshold&#x60;. | [optional]
**paid_credits** | **string** | The number of paid credits. Required only if there is no granted credits. |
**granted_credits** | **string** | The number of free granted credits. Required only if there is no paid credits. |
**created_at** | **\DateTime** | The date of the metadata object creation, represented in ISO 8601 datetime format and expressed in Coordinated Universal Time (UTC). The creation_date provides a standardized and internationally recognized timestamp for when the metadata object was created |

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
