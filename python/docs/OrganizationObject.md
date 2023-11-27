# OrganizationObject


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**lago_id** | **str** | Unique identifier assigned to the organization within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the organization&#39;s record within the Lago system | 
**name** | **str** | The name of your organization. | 
**created_at** | **datetime** | The date of creation of your organization, represented in ISO 8601 datetime format and expressed in Coordinated Universal Time (UTC). | 
**webhook_url** | **str** | The URL of your newest updated webhook endpoint. This URL allows your organization to receive important messages, notifications, or data from the Lago system. By configuring your webhook endpoint to this URL, you can ensure that your organization stays informed and receives relevant information in a timely manner. | [optional] 
**webhook_urls** | **List[str]** | The array containing your webhooks URLs. | [optional] 
**country** | [**Country**](Country.md) |  | [optional] 
**default_currency** | [**Currency**](Currency.md) |  | [optional] 
**address_line1** | **str** | The first line of your organization’s billing address. | [optional] 
**address_line2** | **str** | The second line of your organization’s billing address. | [optional] 
**state** | **str** | The state of your organization’s billing address. | [optional] 
**zipcode** | **str** | The zipcode of your organization’s billing address. | [optional] 
**email** | **str** | The email address of your organization used to bill your customers. | [optional] 
**city** | **str** | The city of your organization’s billing address. | [optional] 
**legal_name** | **str** | The legal name of your organization. | [optional] 
**legal_number** | **str** | The legal number of your organization. | [optional] 
**net_payment_term** | **int** | The net payment term, expressed in days, specifies the duration within which a customer is expected to remit payment after the invoice is finalized. | [optional] 
**tax_identification_number** | **str** | The tax identification number of your organization. | [optional] 
**timezone** | [**Timezone**](Timezone.md) |  | [optional] 
**billing_configuration** | [**OrganizationBillingConfiguration**](OrganizationBillingConfiguration.md) |  | 
**taxes** | [**List[TaxObject]**](TaxObject.md) | List of default organization taxes | [optional] 

## Example

```python
from lago_api.models.organization_object import OrganizationObject

# TODO update the JSON string below
json = "{}"
# create an instance of OrganizationObject from a JSON string
organization_object_instance = OrganizationObject.from_json(json)
# print the JSON string representation of the object
print OrganizationObject.to_json()

# convert the object into a dict
organization_object_dict = organization_object_instance.to_dict()
# create an instance of OrganizationObject from a dict
organization_object_form_dict = organization_object.from_dict(organization_object_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


