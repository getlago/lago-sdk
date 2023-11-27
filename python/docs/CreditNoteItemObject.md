# CreditNoteItemObject


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**lago_id** | **str** | The credit note’s item unique identifier, created by Lago. | 
**amount_cents** | **int** | The credit note’s item amount, expressed in cents. | 
**amount_currency** | [**Currency**](Currency.md) |  | 
**fee** | [**CreditNoteItemObjectFee**](CreditNoteItemObjectFee.md) |  | 

## Example

```python
from lago_api.models.credit_note_item_object import CreditNoteItemObject

# TODO update the JSON string below
json = "{}"
# create an instance of CreditNoteItemObject from a JSON string
credit_note_item_object_instance = CreditNoteItemObject.from_json(json)
# print the JSON string representation of the object
print CreditNoteItemObject.to_json()

# convert the object into a dict
credit_note_item_object_dict = credit_note_item_object_instance.to_dict()
# create an instance of CreditNoteItemObject from a dict
credit_note_item_object_form_dict = credit_note_item_object.from_dict(credit_note_item_object_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


