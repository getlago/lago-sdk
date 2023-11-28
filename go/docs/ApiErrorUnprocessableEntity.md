# ApiErrorUnprocessableEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **int32** |  | 
**Error** | **string** |  | 
**Code** | **string** |  | 
**ErrorDetails** | **map[string]interface{}** |  | 

## Methods

### NewApiErrorUnprocessableEntity

`func NewApiErrorUnprocessableEntity(status int32, error_ string, code string, errorDetails map[string]interface{}, ) *ApiErrorUnprocessableEntity`

NewApiErrorUnprocessableEntity instantiates a new ApiErrorUnprocessableEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiErrorUnprocessableEntityWithDefaults

`func NewApiErrorUnprocessableEntityWithDefaults() *ApiErrorUnprocessableEntity`

NewApiErrorUnprocessableEntityWithDefaults instantiates a new ApiErrorUnprocessableEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ApiErrorUnprocessableEntity) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiErrorUnprocessableEntity) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiErrorUnprocessableEntity) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetError

`func (o *ApiErrorUnprocessableEntity) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ApiErrorUnprocessableEntity) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ApiErrorUnprocessableEntity) SetError(v string)`

SetError sets Error field to given value.


### GetCode

`func (o *ApiErrorUnprocessableEntity) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ApiErrorUnprocessableEntity) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ApiErrorUnprocessableEntity) SetCode(v string)`

SetCode sets Code field to given value.


### GetErrorDetails

`func (o *ApiErrorUnprocessableEntity) GetErrorDetails() map[string]interface{}`

GetErrorDetails returns the ErrorDetails field if non-nil, zero value otherwise.

### GetErrorDetailsOk

`func (o *ApiErrorUnprocessableEntity) GetErrorDetailsOk() (*map[string]interface{}, bool)`

GetErrorDetailsOk returns a tuple with the ErrorDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDetails

`func (o *ApiErrorUnprocessableEntity) SetErrorDetails(v map[string]interface{})`

SetErrorDetails sets ErrorDetails field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


