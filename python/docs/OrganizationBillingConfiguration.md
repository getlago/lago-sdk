# OrganizationBillingConfiguration

The custom billing settings for your organization.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**invoice_footer** | **str** | The customer invoice message that appears at the bottom of each billing documents. | [optional] 
**invoice_grace_period** | **int** | The grace period, expressed in days, for finalizing the invoice. This period refers to the additional time granted to your customers beyond the invoice due date to adjust usage and line items. Can be overwritten by the customer’s grace period. | [optional] 
**document_locale** | **str** | The locale of the billing documents, expressed in the ISO 639-1 format. This field indicates the language or regional variant used for the documents content issued or the embeddable customer portal. | [optional] 

## Example

```python
from lago_api.models.organization_billing_configuration import OrganizationBillingConfiguration

# TODO update the JSON string below
json = "{}"
# create an instance of OrganizationBillingConfiguration from a JSON string
organization_billing_configuration_instance = OrganizationBillingConfiguration.from_json(json)
# print the JSON string representation of the object
print OrganizationBillingConfiguration.to_json()

# convert the object into a dict
organization_billing_configuration_dict = organization_billing_configuration_instance.to_dict()
# create an instance of OrganizationBillingConfiguration from a dict
organization_billing_configuration_form_dict = organization_billing_configuration.from_dict(organization_billing_configuration_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


