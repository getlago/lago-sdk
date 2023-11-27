# CustomerCreateInputCustomer


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**external_id** | **str** | The customer external unique identifier (provided by your own application) | 
**address_line1** | **str** | The first line of the billing address | [optional] 
**address_line2** | **str** | The second line of the billing address | [optional] 
**city** | **str** | The city of the customer&#39;s billing address | [optional] 
**country** | [**Country**](Country.md) |  | [optional] 
**currency** | [**Currency**](Currency.md) |  | [optional] 
**email** | **str** | The email of the customer | [optional] 
**legal_name** | **str** | The legal company name of the customer | [optional] 
**legal_number** | **str** | The legal company number of the customer | [optional] 
**logo_url** | **str** | The logo URL of the customer | [optional] 
**name** | **str** | The full name of the customer | [optional] 
**phone** | **str** | The phone number of the customer | [optional] 
**state** | **str** | The state of the customer&#39;s billing address | [optional] 
**tax_codes** | **List[str]** | List of unique code used to identify the taxes. | [optional] 
**tax_identification_number** | **str** | The tax identification number of the customer | [optional] 
**timezone** | [**Timezone**](Timezone.md) |  | [optional] 
**url** | **str** | The custom website URL of the customer | [optional] 
**zipcode** | **str** | The zipcode of the customer&#39;s billing address | [optional] 
**net_payment_term** | **int** | The net payment term, expressed in days, specifies the duration within which a customer is expected to remit payment after the invoice is finalized. | [optional] 
**billing_configuration** | [**CustomerBillingConfiguration**](CustomerBillingConfiguration.md) |  | [optional] 
**metadata** | [**List[CustomerCreateInputCustomerMetadataInner]**](CustomerCreateInputCustomerMetadataInner.md) | Set of key-value pairs that you can attach to a customer. This can be useful for storing additional information about the customer in a structured format | [optional] 

## Example

```python
from lago_api.models.customer_create_input_customer import CustomerCreateInputCustomer

# TODO update the JSON string below
json = "{}"
# create an instance of CustomerCreateInputCustomer from a JSON string
customer_create_input_customer_instance = CustomerCreateInputCustomer.from_json(json)
# print the JSON string representation of the object
print CustomerCreateInputCustomer.to_json()

# convert the object into a dict
customer_create_input_customer_dict = customer_create_input_customer_instance.to_dict()
# create an instance of CustomerCreateInputCustomer from a dict
customer_create_input_customer_form_dict = customer_create_input_customer.from_dict(customer_create_input_customer_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


