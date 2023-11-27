# CustomerObjectExtended


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**lago_id** | **str** | Unique identifier assigned to the customer within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the customer&#39;s record within the Lago system | 
**sequential_id** | **int** | The unique identifier assigned to the customer within the organization&#39;s scope. This identifier is used to track and reference the customer&#39;s order of creation within the organization&#39;s system. It ensures that each customer has a distinct &#x60;sequential_id&#x60;&#x60; associated with them, allowing for easy identification and sorting based on the order of creation | 
**slug** | **str** | A concise and unique identifier for the customer, formed by combining the Organization&#39;s &#x60;name&#x60;, &#x60;id&#x60;, and customer&#39;s &#x60;sequential_id&#x60; | 
**external_id** | **str** | The customer external unique identifier (provided by your own application) | 
**address_line1** | **str** | The first line of the billing address | [optional] 
**address_line2** | **str** | The second line of the billing address | [optional] 
**applicable_timezone** | [**Timezone**](Timezone.md) |  | 
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
**tax_identification_number** | **str** | The tax identification number of the customer | [optional] 
**timezone** | [**Timezone**](Timezone.md) |  | [optional] 
**url** | **str** | The custom website URL of the customer | [optional] 
**zipcode** | **str** | The zipcode of the customer&#39;s billing address | [optional] 
**net_payment_term** | **int** | The net payment term, expressed in days, specifies the duration within which a customer is expected to remit payment after the invoice is finalized. | [optional] 
**created_at** | **datetime** | The date of the customer creation, represented in ISO 8601 datetime format and expressed in Coordinated Universal Time (UTC). The creation_date provides a standardized and internationally recognized timestamp for when the customer object was created | 
**updated_at** | **datetime** | The date of the customer update, represented in ISO 8601 datetime format and expressed in Coordinated Universal Time (UTC). The update_date provides a standardized and internationally recognized timestamp for when the customer object was updated | [optional] 
**billing_configuration** | [**CustomerBillingConfiguration**](CustomerBillingConfiguration.md) |  | [optional] 
**metadata** | [**List[CustomerMetadata]**](CustomerMetadata.md) |  | [optional] 
**taxes** | [**List[TaxObject]**](TaxObject.md) | List of customer taxes | [optional] 

## Example

```python
from lago_api.models.customer_object_extended import CustomerObjectExtended

# TODO update the JSON string below
json = "{}"
# create an instance of CustomerObjectExtended from a JSON string
customer_object_extended_instance = CustomerObjectExtended.from_json(json)
# print the JSON string representation of the object
print CustomerObjectExtended.to_json()

# convert the object into a dict
customer_object_extended_dict = customer_object_extended_instance.to_dict()
# create an instance of CustomerObjectExtended from a dict
customer_object_extended_form_dict = customer_object_extended.from_dict(customer_object_extended_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


