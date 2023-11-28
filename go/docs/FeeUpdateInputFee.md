# FeeUpdateInputFee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentStatus** | **string** | The payment status of the fee. Possible values are &#x60;pending&#x60;, &#x60;succeeded&#x60;, &#x60;failed&#x60; or &#x60;refunded&#x60;. | 

## Methods

### NewFeeUpdateInputFee

`func NewFeeUpdateInputFee(paymentStatus string, ) *FeeUpdateInputFee`

NewFeeUpdateInputFee instantiates a new FeeUpdateInputFee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeeUpdateInputFeeWithDefaults

`func NewFeeUpdateInputFeeWithDefaults() *FeeUpdateInputFee`

NewFeeUpdateInputFeeWithDefaults instantiates a new FeeUpdateInputFee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentStatus

`func (o *FeeUpdateInputFee) GetPaymentStatus() string`

GetPaymentStatus returns the PaymentStatus field if non-nil, zero value otherwise.

### GetPaymentStatusOk

`func (o *FeeUpdateInputFee) GetPaymentStatusOk() (*string, bool)`

GetPaymentStatusOk returns a tuple with the PaymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentStatus

`func (o *FeeUpdateInputFee) SetPaymentStatus(v string)`

SetPaymentStatus sets PaymentStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


