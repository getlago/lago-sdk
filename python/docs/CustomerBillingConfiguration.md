# CustomerBillingConfiguration

Configuration specific to the payment provider, utilized for billing the customer. This object contains settings and parameters necessary for processing payments and invoicing the customer.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**invoice_grace_period** | **int** | The grace period, expressed in days, for the invoice. This period refers to the additional time granted to the customer beyond the invoice due date to adjust usage and line items | [optional] 
**payment_provider** | **str** | The payment provider utilized to initiate payments for invoices issued by Lago. Accepted values: &#x60;stripe&#x60;, &#x60;adyen&#x60;, &#x60;gocardless&#x60; or null. This field is required if you intend to assign a &#x60;provider_customer_id&#x60;. | [optional] 
**provider_customer_id** | **str** | The customer ID within the payment provider&#39;s system. If this field is not provided, Lago has the option to create a new customer record within the payment provider&#39;s system on behalf of the customer | [optional] 
**sync** | **bool** | Set this field to &#x60;true&#x60; if you want to create the customer in the payment provider synchronously with the customer creation process in Lago. This option is applicable only when the &#x60;provider_customer_id&#x60; is &#x60;null&#x60; and the customer is automatically created in the payment provider through Lago. By default, the value is set to &#x60;false&#x60; | [optional] 
**sync_with_provider** | **bool** | Set this field to &#x60;true&#x60; if you want to create a customer record in the payment provider&#39;s system. This option is applicable only when the &#x60;provider_customer_id&#x60; is null and the &#x60;sync_with_provider&#x60; field is set to &#x60;true&#x60;. By default, the value is set to &#x60;false&#x60; | [optional] 
**document_locale** | **str** | The document locale, specified in the ISO 639-1 format. This field represents the language or locale used for the documents issued by Lago | [optional] 
**provider_payment_methods** | **List[str]** | Specifies the available payment methods that can be used for this customer when &#x60;payment_provider&#x60; is set to &#x60;stripe&#x60;. The &#x60;provider_payment_methods&#x60; field is an array that allows multiple payment options to be defined. If this field is not explicitly set, all the payment methods are selected. For now, possible values are &#x60;card&#x60; and &#x60;sepa_debit&#x60;. | [optional] 

## Example

```python
from lago_api.models.customer_billing_configuration import CustomerBillingConfiguration

# TODO update the JSON string below
json = "{}"
# create an instance of CustomerBillingConfiguration from a JSON string
customer_billing_configuration_instance = CustomerBillingConfiguration.from_json(json)
# print the JSON string representation of the object
print CustomerBillingConfiguration.to_json()

# convert the object into a dict
customer_billing_configuration_dict = customer_billing_configuration_instance.to_dict()
# create an instance of CustomerBillingConfiguration from a dict
customer_billing_configuration_form_dict = customer_billing_configuration.from_dict(customer_billing_configuration_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


