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

// checks if the CreditNoteUpdateInputCreditNote type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreditNoteUpdateInputCreditNote{}

// CreditNoteUpdateInputCreditNote struct for CreditNoteUpdateInputCreditNote
type CreditNoteUpdateInputCreditNote struct {
	// The status of the refund portion of the credit note. It indicates the current state or condition of the refund associated with the credit note. The possible values for this field are:  - `pending`: this status indicates that the refund is pending execution. The refund request has been initiated but has not been processed or completed yet. - `succeeded`: this status indicates that the refund has been successfully executed. The refund amount has been processed and returned to the customer or the designated recipient. - `failed`: this status indicates that the refund failed to execute. The refund request encountered an error or unsuccessful processing, and the refund amount could not be returned.
	RefundStatus string `json:"refund_status"`
}

type _CreditNoteUpdateInputCreditNote CreditNoteUpdateInputCreditNote

// NewCreditNoteUpdateInputCreditNote instantiates a new CreditNoteUpdateInputCreditNote object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditNoteUpdateInputCreditNote(refundStatus string) *CreditNoteUpdateInputCreditNote {
	this := CreditNoteUpdateInputCreditNote{}
	this.RefundStatus = refundStatus
	return &this
}

// NewCreditNoteUpdateInputCreditNoteWithDefaults instantiates a new CreditNoteUpdateInputCreditNote object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditNoteUpdateInputCreditNoteWithDefaults() *CreditNoteUpdateInputCreditNote {
	this := CreditNoteUpdateInputCreditNote{}
	return &this
}

// GetRefundStatus returns the RefundStatus field value
func (o *CreditNoteUpdateInputCreditNote) GetRefundStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RefundStatus
}

// GetRefundStatusOk returns a tuple with the RefundStatus field value
// and a boolean to check if the value has been set.
func (o *CreditNoteUpdateInputCreditNote) GetRefundStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RefundStatus, true
}

// SetRefundStatus sets field value
func (o *CreditNoteUpdateInputCreditNote) SetRefundStatus(v string) {
	o.RefundStatus = v
}

func (o CreditNoteUpdateInputCreditNote) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreditNoteUpdateInputCreditNote) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["refund_status"] = o.RefundStatus
	return toSerialize, nil
}

func (o *CreditNoteUpdateInputCreditNote) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"refund_status",
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

	varCreditNoteUpdateInputCreditNote := _CreditNoteUpdateInputCreditNote{}

	err = json.Unmarshal(bytes, &varCreditNoteUpdateInputCreditNote)

	if err != nil {
		return err
	}

	*o = CreditNoteUpdateInputCreditNote(varCreditNoteUpdateInputCreditNote)

	return err
}

type NullableCreditNoteUpdateInputCreditNote struct {
	value *CreditNoteUpdateInputCreditNote
	isSet bool
}

func (v NullableCreditNoteUpdateInputCreditNote) Get() *CreditNoteUpdateInputCreditNote {
	return v.value
}

func (v *NullableCreditNoteUpdateInputCreditNote) Set(val *CreditNoteUpdateInputCreditNote) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditNoteUpdateInputCreditNote) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditNoteUpdateInputCreditNote) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditNoteUpdateInputCreditNote(val *CreditNoteUpdateInputCreditNote) *NullableCreditNoteUpdateInputCreditNote {
	return &NullableCreditNoteUpdateInputCreditNote{value: val, isSet: true}
}

func (v NullableCreditNoteUpdateInputCreditNote) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditNoteUpdateInputCreditNote) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


