# OrganizationUpdateInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**organization** | [**OrganizationUpdateInputOrganization**](OrganizationUpdateInputOrganization.md) |  | 

## Example

```python
from lago_api.models.organization_update_input import OrganizationUpdateInput

# TODO update the JSON string below
json = "{}"
# create an instance of OrganizationUpdateInput from a JSON string
organization_update_input_instance = OrganizationUpdateInput.from_json(json)
# print the JSON string representation of the object
print OrganizationUpdateInput.to_json()

# convert the object into a dict
organization_update_input_dict = organization_update_input_instance.to_dict()
# create an instance of OrganizationUpdateInput from a dict
organization_update_input_form_dict = organization_update_input.from_dict(organization_update_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


