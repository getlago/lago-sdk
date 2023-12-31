/*
Lago API documentation

Lago API allows your application to push customer information and metrics (events) from your application to the billing application.

API version: 0.52.0-beta
Contact: tech@getlago.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lagoapi

import (
	"encoding/json"
	"fmt"
)

// checks if the ApiErrorNotFound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiErrorNotFound{}

// ApiErrorNotFound struct for ApiErrorNotFound
type ApiErrorNotFound struct {
	Status int32 `json:"status"`
	Error string `json:"error"`
	Code string `json:"code"`
}

type _ApiErrorNotFound ApiErrorNotFound

// NewApiErrorNotFound instantiates a new ApiErrorNotFound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiErrorNotFound(status int32, error_ string, code string) *ApiErrorNotFound {
	this := ApiErrorNotFound{}
	this.Status = status
	this.Error = error_
	this.Code = code
	return &this
}

// NewApiErrorNotFoundWithDefaults instantiates a new ApiErrorNotFound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiErrorNotFoundWithDefaults() *ApiErrorNotFound {
	this := ApiErrorNotFound{}
	return &this
}

// GetStatus returns the Status field value
func (o *ApiErrorNotFound) GetStatus() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ApiErrorNotFound) GetStatusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ApiErrorNotFound) SetStatus(v int32) {
	o.Status = v
}

// GetError returns the Error field value
func (o *ApiErrorNotFound) GetError() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *ApiErrorNotFound) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *ApiErrorNotFound) SetError(v string) {
	o.Error = v
}

// GetCode returns the Code field value
func (o *ApiErrorNotFound) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *ApiErrorNotFound) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *ApiErrorNotFound) SetCode(v string) {
	o.Code = v
}

func (o ApiErrorNotFound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiErrorNotFound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	toSerialize["error"] = o.Error
	toSerialize["code"] = o.Code
	return toSerialize, nil
}

func (o *ApiErrorNotFound) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
		"error",
		"code",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varApiErrorNotFound := _ApiErrorNotFound{}

	err = json.Unmarshal(bytes, &varApiErrorNotFound)

	if err != nil {
		return err
	}

	*o = ApiErrorNotFound(varApiErrorNotFound)

	return err
}

type NullableApiErrorNotFound struct {
	value *ApiErrorNotFound
	isSet bool
}

func (v NullableApiErrorNotFound) Get() *ApiErrorNotFound {
	return v.value
}

func (v *NullableApiErrorNotFound) Set(val *ApiErrorNotFound) {
	v.value = val
	v.isSet = true
}

func (v NullableApiErrorNotFound) IsSet() bool {
	return v.isSet
}

func (v *NullableApiErrorNotFound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiErrorNotFound(val *ApiErrorNotFound) *NullableApiErrorNotFound {
	return &NullableApiErrorNotFound{value: val, isSet: true}
}

func (v NullableApiErrorNotFound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiErrorNotFound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


