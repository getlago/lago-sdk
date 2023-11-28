# GroupPropertiesObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | **string** | Unique identifier of a billable metric group, created by Lago. | 
**InvoiceDisplayName** | Pointer to **string** | Specifies the name that will be displayed on an invoice. If no value is set for this field, the name of the actual group will be used as the default display name. | [optional] 
**Values** | [**GroupPropertiesObjectValues**](GroupPropertiesObjectValues.md) |  | 

## Methods

### NewGroupPropertiesObject

`func NewGroupPropertiesObject(groupId string, values GroupPropertiesObjectValues, ) *GroupPropertiesObject`

NewGroupPropertiesObject instantiates a new GroupPropertiesObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupPropertiesObjectWithDefaults

`func NewGroupPropertiesObjectWithDefaults() *GroupPropertiesObject`

NewGroupPropertiesObjectWithDefaults instantiates a new GroupPropertiesObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *GroupPropertiesObject) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *GroupPropertiesObject) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *GroupPropertiesObject) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetInvoiceDisplayName

`func (o *GroupPropertiesObject) GetInvoiceDisplayName() string`

GetInvoiceDisplayName returns the InvoiceDisplayName field if non-nil, zero value otherwise.

### GetInvoiceDisplayNameOk

`func (o *GroupPropertiesObject) GetInvoiceDisplayNameOk() (*string, bool)`

GetInvoiceDisplayNameOk returns a tuple with the InvoiceDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceDisplayName

`func (o *GroupPropertiesObject) SetInvoiceDisplayName(v string)`

SetInvoiceDisplayName sets InvoiceDisplayName field to given value.

### HasInvoiceDisplayName

`func (o *GroupPropertiesObject) HasInvoiceDisplayName() bool`

HasInvoiceDisplayName returns a boolean if a field has been set.

### GetValues

`func (o *GroupPropertiesObject) GetValues() GroupPropertiesObjectValues`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GroupPropertiesObject) GetValuesOk() (*GroupPropertiesObjectValues, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GroupPropertiesObject) SetValues(v GroupPropertiesObjectValues)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


