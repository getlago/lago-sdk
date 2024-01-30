# # CustomerObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**lago_id** | **string** | Unique identifier assigned to the customer within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the customer&#39;s record within the Lago system |
**sequential_id** | **int** | The unique identifier assigned to the customer within the organization&#39;s scope. This identifier is used to track and reference the customer&#39;s order of creation within the organization&#39;s system. It ensures that each customer has a distinct &#x60;sequential_id&#x60;&#x60; associated with them, allowing for easy identification and sorting based on the order of creation |
**slug** | **string** | A concise and unique identifier for the customer, formed by combining the Organization&#39;s &#x60;name&#x60;, &#x60;id&#x60;, and customer&#39;s &#x60;sequential_id&#x60; |
**external_id** | **string** | The customer external unique identifier (provided by your own application) |
**address_line1** | **string** | The first line of the billing address | [optional]
**address_line2** | **string** | The second line of the billing address | [optional]
**applicable_timezone** | [**Timezone**](Timezone.md) |  |
**city** | **string** | The city of the customer&#39;s billing address | [optional]
**country** | [**Country**](Country.md) |  | [optional]
**currency** | [**Currency**](Currency.md) |  | [optional]
**email** | **string** | The email of the customer | [optional]
**legal_name** | **string** | The legal company name of the customer | [optional]
**legal_number** | **string** | The legal company number of the customer | [optional]
**logo_url** | **string** | The logo URL of the customer | [optional]
**name** | **string** | The full name of the customer | [optional]
**phone** | **string** | The phone number of the customer | [optional]
**state** | **string** | The state of the customer&#39;s billing address | [optional]
**tax_identification_number** | **string** | The tax identification number of the customer | [optional]
**timezone** | [**Timezone**](Timezone.md) |  | [optional]
**url** | **string** | The custom website URL of the customer | [optional]
**zipcode** | **string** | The zipcode of the customer&#39;s billing address | [optional]
**net_payment_term** | **int** | The net payment term, expressed in days, specifies the duration within which a customer is expected to remit payment after the invoice is finalized. | [optional]
**created_at** | **\DateTime** | The date of the customer creation, represented in ISO 8601 datetime format and expressed in Coordinated Universal Time (UTC). The creation_date provides a standardized and internationally recognized timestamp for when the customer object was created |
**updated_at** | **\DateTime** | The date of the customer update, represented in ISO 8601 datetime format and expressed in Coordinated Universal Time (UTC). The update_date provides a standardized and internationally recognized timestamp for when the customer object was updated | [optional]
**billing_configuration** | [**\OpenAPI\Client\Model\CustomerBillingConfiguration**](CustomerBillingConfiguration.md) |  | [optional]
**metadata** | [**\OpenAPI\Client\Model\CustomerMetadata[]**](CustomerMetadata.md) |  | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
