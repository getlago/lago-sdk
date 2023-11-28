# ApiErrorUnauthorized

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **int32** |  | 
**Error** | **string** |  | 

## Methods

### NewApiErrorUnauthorized

`func NewApiErrorUnauthorized(status int32, error_ string, ) *ApiErrorUnauthorized`

NewApiErrorUnauthorized instantiates a new ApiErrorUnauthorized object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiErrorUnauthorizedWithDefaults

`func NewApiErrorUnauthorizedWithDefaults() *ApiErrorUnauthorized`

NewApiErrorUnauthorizedWithDefaults instantiates a new ApiErrorUnauthorized object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ApiErrorUnauthorized) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiErrorUnauthorized) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiErrorUnauthorized) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetError

`func (o *ApiErrorUnauthorized) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ApiErrorUnauthorized) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ApiErrorUnauthorized) SetError(v string)`

SetError sets Error field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


