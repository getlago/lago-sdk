# GroupObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LagoId** | **string** | Unique identifier of a specific group associated with the billable metric. | 
**Key** | **string** | Key of a specific group associated with the billable metric. | 
**Value** | **string** | One of the values for a specific group associated with the billable metric. | 

## Methods

### NewGroupObject

`func NewGroupObject(lagoId string, key string, value string, ) *GroupObject`

NewGroupObject instantiates a new GroupObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupObjectWithDefaults

`func NewGroupObjectWithDefaults() *GroupObject`

NewGroupObjectWithDefaults instantiates a new GroupObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLagoId

`func (o *GroupObject) GetLagoId() string`

GetLagoId returns the LagoId field if non-nil, zero value otherwise.

### GetLagoIdOk

`func (o *GroupObject) GetLagoIdOk() (*string, bool)`

GetLagoIdOk returns a tuple with the LagoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagoId

`func (o *GroupObject) SetLagoId(v string)`

SetLagoId sets LagoId field to given value.


### GetKey

`func (o *GroupObject) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *GroupObject) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *GroupObject) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *GroupObject) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GroupObject) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GroupObject) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


