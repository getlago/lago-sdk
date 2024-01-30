# # WalletObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**lago_id** | **string** | Unique identifier assigned to the wallet within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the wallet’s record within the Lago system. |
**lago_customer_id** | **string** | Unique identifier assigned to the customer within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the customer’s record within the Lago system. |
**external_customer_id** | **string** | The customer external unique identifier (provided by your own application) |
**status** | **string** | The status of the wallet. Possible values are &#x60;active&#x60; or &#x60;terminated&#x60;. |
**currency** | [**Currency**](Currency.md) |  |
**name** | **string** | The name of the wallet. | [optional]
**rate_amount** | **string** | The rate of conversion between credits and the amount in the specified currency. It indicates the ratio or factor used to convert credits into the corresponding monetary value in the currency of the transaction. |
**credits_balance** | **string** | The current wallet balance expressed in credits. |
**balance_cents** | **int** | The current wallet balance expressed in cents. |
**consumed_credits** | **string** | The number of consumed credits. |
**created_at** | **\DateTime** | The date of the wallet creation, represented in ISO 8601 datetime format and expressed in Coordinated Universal Time (UTC). |
**expiration_at** | **\DateTime** | The date and time that determines when the wallet will expire. It follows the ISO 8601 datetime format and is expressed in Coordinated Universal Time (UTC). | [optional]
**last_balance_sync_at** | **\DateTime** | The date and time of the last balance top-up. It follows the ISO 8601 datetime format and is expressed in Coordinated Universal Time (UTC). | [optional]
**last_consumed_credit_at** | **\DateTime** | The date and time of the last credits consumption. It follows the ISO 8601 datetime format and is expressed in Coordinated Universal Time (UTC). | [optional]
**terminated_at** | **\DateTime** | The date of terminaison of the wallet. It follows the ISO 8601 datetime format and is expressed in Coordinated Universal Time (UTC). | [optional]
**recurring_transaction_rules** | [**\OpenAPI\Client\Model\WalletObjectRecurringTransactionRulesInner[]**](WalletObjectRecurringTransactionRulesInner.md) | List of recurring transaction rules. Currently, we only allow one recurring rule per wallet. | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
