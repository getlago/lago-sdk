# # PlanOverridesObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**amount_cents** | **int** | The base cost of the plan, excluding any applicable taxes, that is billed on a recurring basis. This value is defined at 0 if your plan is a pay-as-you-go plan. | [optional]
**amount_currency** | [**Currency**](Currency.md) |  | [optional]
**description** | **string** | The description on the plan. | [optional]
**invoice_display_name** | **string** | Specifies the name that will be displayed on an invoice. If no value is set for this field, the name of the plan will be used as the default display name. | [optional]
**name** | **string** | The name of the plan. | [optional]
**tax_codes** | **string[]** | List of unique code used to identify the taxes. | [optional]
**trial_period** | **float** | The duration in days during which the base cost of the plan is offered for free. | [optional]
**charges** | [**\OpenAPI\Client\Model\PlanOverridesObjectChargesInner[]**](PlanOverridesObjectChargesInner.md) | Additional usage-based charges for this plan. | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
