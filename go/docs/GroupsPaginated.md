# GroupsPaginated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Groups** | [**[]GroupObject**](GroupObject.md) |  | 
**Meta** | [**PaginationMeta**](PaginationMeta.md) |  | 

## Methods

### NewGroupsPaginated

`func NewGroupsPaginated(groups []GroupObject, meta PaginationMeta, ) *GroupsPaginated`

NewGroupsPaginated instantiates a new GroupsPaginated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupsPaginatedWithDefaults

`func NewGroupsPaginatedWithDefaults() *GroupsPaginated`

NewGroupsPaginatedWithDefaults instantiates a new GroupsPaginated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroups

`func (o *GroupsPaginated) GetGroups() []GroupObject`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *GroupsPaginated) GetGroupsOk() (*[]GroupObject, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *GroupsPaginated) SetGroups(v []GroupObject)`

SetGroups sets Groups field to given value.


### GetMeta

`func (o *GroupsPaginated) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GroupsPaginated) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GroupsPaginated) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


