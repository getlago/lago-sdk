# LagoAPI::AddOnCreateInputAddOn

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **name** | **String** | The name of the add-on. |  |
| **invoice_display_name** | **String** | Specifies the name that will be displayed on an invoice. If no value is set for this field, the name of the actual charge will be used as the default display name. | [optional] |
| **code** | **String** | Unique code used to identify the add-on. |  |
| **amount_cents** | **Integer** | The cost of the add-on in cents, excluding any applicable taxes, that is billed to a customer. By creating a one-off invoice, you will be able to override this value. |  |
| **amount_currency** | [**Currency**](Currency.md) |  |  |
| **description** | **String** | The description of the add-on. | [optional] |
| **tax_codes** | **Array&lt;String&gt;** | List of unique code used to identify the taxes. | [optional] |

## Example

```ruby
require 'lago_ruby'

instance = LagoAPI::AddOnCreateInputAddOn.new(
  name: Setup Fee,
  invoice_display_name: Setup Fee (SF1),
  code: setup_fee,
  amount_cents: 50000,
  amount_currency: null,
  description: Implementation fee for new customers.,
  tax_codes: [french_standard_vat]
)
```

