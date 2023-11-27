# GetCustomerPortalUrl200Response


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**customer** | [**GetCustomerPortalUrl200ResponseCustomer**](GetCustomerPortalUrl200ResponseCustomer.md) |  | 

## Example

```python
from lago_api.models.get_customer_portal_url200_response import GetCustomerPortalUrl200Response

# TODO update the JSON string below
json = "{}"
# create an instance of GetCustomerPortalUrl200Response from a JSON string
get_customer_portal_url200_response_instance = GetCustomerPortalUrl200Response.from_json(json)
# print the JSON string representation of the object
print GetCustomerPortalUrl200Response.to_json()

# convert the object into a dict
get_customer_portal_url200_response_dict = get_customer_portal_url200_response_instance.to_dict()
# create an instance of GetCustomerPortalUrl200Response from a dict
get_customer_portal_url200_response_form_dict = get_customer_portal_url200_response.from_dict(get_customer_portal_url200_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


