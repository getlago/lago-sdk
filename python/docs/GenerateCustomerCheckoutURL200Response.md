# GenerateCustomerCheckoutURL200Response

.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**lago_customer_id** | **str** | Unique identifier assigned to the customer within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the customer&#39;s record within the Lago system | [optional] 
**external_customer_id** | **str** | The customer external unique identifier (provided by your own application) | [optional] 
**payment_provider** | **str** | The Payment Provider name linked to the Customer. | [optional] 
**checkout_url** | **str** | The new generated Payment Provider Checkout URL for the Customer. | [optional] 

## Example

```python
from lago_api.models.generate_customer_checkout_url200_response import GenerateCustomerCheckoutURL200Response

# TODO update the JSON string below
json = "{}"
# create an instance of GenerateCustomerCheckoutURL200Response from a JSON string
generate_customer_checkout_url200_response_instance = GenerateCustomerCheckoutURL200Response.from_json(json)
# print the JSON string representation of the object
print GenerateCustomerCheckoutURL200Response.to_json()

# convert the object into a dict
generate_customer_checkout_url200_response_dict = generate_customer_checkout_url200_response_instance.to_dict()
# create an instance of GenerateCustomerCheckoutURL200Response from a dict
generate_customer_checkout_url200_response_form_dict = generate_customer_checkout_url200_response.from_dict(generate_customer_checkout_url200_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


