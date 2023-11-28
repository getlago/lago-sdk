# ApiErrorForbidden

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **int32** |  | 
**Error** | **string** |  | 
**Code** | **string** |  | 

## Methods

### NewApiErrorForbidden

`func NewApiErrorForbidden(status int32, error_ string, code string, ) *ApiErrorForbidden`

NewApiErrorForbidden instantiates a new ApiErrorForbidden object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiErrorForbiddenWithDefaults

`func NewApiErrorForbiddenWithDefaults() *ApiErrorForbidden`

NewApiErrorForbiddenWithDefaults instantiates a new ApiErrorForbidden object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ApiErrorForbidden) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiErrorForbidden) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiErrorForbidden) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetError

`func (o *ApiErrorForbidden) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ApiErrorForbidden) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ApiErrorForbidden) SetError(v string)`

SetError sets Error field to given value.


### GetCode

`func (o *ApiErrorForbidden) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ApiErrorForbidden) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ApiErrorForbidden) SetCode(v string)`

SetCode sets Code field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


