# # WalletTransactionObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**lago_id** | **string** | Unique identifier assigned to the wallet transaction within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the wallet transaction’s record within the Lago system. |
**lago_wallet_id** | **string** | Unique identifier assigned to the wallet within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the wallet’s record within the Lago system. |
**status** | **string** | The status of the wallet transaction. Possible values are &#x60;pending&#x60; or &#x60;settled&#x60;. |
**transaction_type** | **string** | The type of transaction. Possible values are &#x60;inbound&#x60; (increasing the balance) or &#x60;outbound&#x60; (decreasing the balance). |
**amount** | **string** | The amount of credits based on the rate and the currency. |
**credit_amount** | **string** | The number of credits used in the wallet transaction. |
**settled_at** | **\DateTime** | The date when wallet transaction is settled, represented in ISO 8601 datetime format and expressed in Coordinated Universal Time (UTC). | [optional]
**created_at** | **\DateTime** | The date of the wallet transaction creation, represented in ISO 8601 datetime format and expressed in Coordinated Universal Time (UTC). |

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
