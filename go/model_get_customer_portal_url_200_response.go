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

// checks if the GetCustomerPortalUrl200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCustomerPortalUrl200Response{}

// GetCustomerPortalUrl200Response struct for GetCustomerPortalUrl200Response
type GetCustomerPortalUrl200Response struct {
	Customer GetCustomerPortalUrl200ResponseCustomer `json:"customer"`
}

type _GetCustomerPortalUrl200Response GetCustomerPortalUrl200Response

// NewGetCustomerPortalUrl200Response instantiates a new GetCustomerPortalUrl200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCustomerPortalUrl200Response(customer GetCustomerPortalUrl200ResponseCustomer) *GetCustomerPortalUrl200Response {
	this := GetCustomerPortalUrl200Response{}
	this.Customer = customer
	return &this
}

// NewGetCustomerPortalUrl200ResponseWithDefaults instantiates a new GetCustomerPortalUrl200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCustomerPortalUrl200ResponseWithDefaults() *GetCustomerPortalUrl200Response {
	this := GetCustomerPortalUrl200Response{}
	return &this
}

// GetCustomer returns the Customer field value
func (o *GetCustomerPortalUrl200Response) GetCustomer() GetCustomerPortalUrl200ResponseCustomer {
	if o == nil {
		var ret GetCustomerPortalUrl200ResponseCustomer
		return ret
	}

	return o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value
// and a boolean to check if the value has been set.
func (o *GetCustomerPortalUrl200Response) GetCustomerOk() (*GetCustomerPortalUrl200ResponseCustomer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Customer, true
}

// SetCustomer sets field value
func (o *GetCustomerPortalUrl200Response) SetCustomer(v GetCustomerPortalUrl200ResponseCustomer) {
	o.Customer = v
}

func (o GetCustomerPortalUrl200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCustomerPortalUrl200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["customer"] = o.Customer
	return toSerialize, nil
}

func (o *GetCustomerPortalUrl200Response) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"customer",
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

	varGetCustomerPortalUrl200Response := _GetCustomerPortalUrl200Response{}

	err = json.Unmarshal(bytes, &varGetCustomerPortalUrl200Response)

	if err != nil {
		return err
	}

	*o = GetCustomerPortalUrl200Response(varGetCustomerPortalUrl200Response)

	return err
}

type NullableGetCustomerPortalUrl200Response struct {
	value *GetCustomerPortalUrl200Response
	isSet bool
}

func (v NullableGetCustomerPortalUrl200Response) Get() *GetCustomerPortalUrl200Response {
	return v.value
}

func (v *NullableGetCustomerPortalUrl200Response) Set(val *GetCustomerPortalUrl200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCustomerPortalUrl200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCustomerPortalUrl200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCustomerPortalUrl200Response(val *GetCustomerPortalUrl200Response) *NullableGetCustomerPortalUrl200Response {
	return &NullableGetCustomerPortalUrl200Response{value: val, isSet: true}
}

func (v NullableGetCustomerPortalUrl200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCustomerPortalUrl200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

