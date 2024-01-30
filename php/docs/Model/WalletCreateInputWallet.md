# # WalletCreateInputWallet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **string** | The name of the wallet. | [optional]
**rate_amount** | **string** | The rate of conversion between credits and the amount in the specified currency. It indicates the ratio or factor used to convert credits into the corresponding monetary value in the currency of the transaction. |
**currency** | [**Currency**](Currency.md) |  |
**paid_credits** | **string** | The number of paid credits. Required only if there is no granted credits. | [optional]
**granted_credits** | **string** | The number of free granted credits. Required only if there is no paid credits. | [optional]
**external_customer_id** | **string** | The customer external unique identifier (provided by your own application) |
**expiration_at** | **\DateTime** | The date and time that determines when the wallet will expire. It follows the ISO 8601 datetime format and is expressed in Coordinated Universal Time (UTC). | [optional]
**recurring_transaction_rules** | [**\OpenAPI\Client\Model\WalletCreateInputWalletRecurringTransactionRulesInner[]**](WalletCreateInputWalletRecurringTransactionRulesInner.md) | List of recurring transaction rules. Currently, we only allow one recurring rule per wallet. | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
