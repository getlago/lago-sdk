# OrganizationUpdateInputOrganization


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**webhook_url** | **str** | The URL of your newest updated webhook endpoint. This URL allows your organization to receive important messages, notifications, or data from the Lago system. By configuring your webhook endpoint to this URL, you can ensure that your organization stays informed and receives relevant information in a timely manner. | [optional] 
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
**email_settings** | **List[str]** | Represents the email settings of the organization. It allows you to define which documents are sent by email. The field value determines the types of documents that trigger email notifications. Possible values for are &#x60;invoice.finalized&#x60; and &#x60;credit_note.created&#x60;. By configuring this field, you can specify whether invoices, credit notes, or both should be sent to recipients via email. | [optional] 
**billing_configuration** | [**OrganizationBillingConfiguration**](OrganizationBillingConfiguration.md) |  | [optional] 

## Example

```python
from lago_api.models.organization_update_input_organization import OrganizationUpdateInputOrganization

# TODO update the JSON string below
json = "{}"
# create an instance of OrganizationUpdateInputOrganization from a JSON string
organization_update_input_organization_instance = OrganizationUpdateInputOrganization.from_json(json)
# print the JSON string representation of the object
print OrganizationUpdateInputOrganization.to_json()

# convert the object into a dict
organization_update_input_organization_dict = organization_update_input_organization_instance.to_dict()
# create an instance of OrganizationUpdateInputOrganization from a dict
organization_update_input_organization_form_dict = organization_update_input_organization.from_dict(organization_update_input_organization_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


