# AppliedCouponInputAppliedCoupon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalCustomerId** | **string** | The customer external unique identifier (provided by your own application) | 
**CouponCode** | **string** | Unique code used to identify the coupon. | 
**Frequency** | Pointer to **NullableString** | The type of frequency for the coupon. It can have three possible values: &#x60;once&#x60;, &#x60;recurring&#x60; or &#x60;forever&#x60;.  - If set to &#x60;once&#x60;, the coupon is applicable only for a single use. - If set to &#x60;recurring&#x60;, the coupon can be used multiple times for recurring billing periods. - If set to &#x60;forever&#x60;, the coupon has unlimited usage and can be applied indefinitely. | [optional] 
**FrequencyDuration** | Pointer to **NullableInt32** | Specifies the number of billing periods to which the coupon applies. This field is required only for coupons with a &#x60;recurring&#x60; frequency type | [optional] 
**AmountCents** | Pointer to **NullableInt32** | The amount of the coupon in cents. This field is required only for coupon with &#x60;fixed_amount&#x60; type. | [optional] 
**AmountCurrency** | Pointer to [**Currency**](Currency.md) |  | [optional] 
**PercentageRate** | Pointer to **NullableString** | The percentage rate of the coupon. This field is required only for coupons with a &#x60;percentage&#x60; coupon type. | [optional] 

## Methods

### NewAppliedCouponInputAppliedCoupon

`func NewAppliedCouponInputAppliedCoupon(externalCustomerId string, couponCode string, ) *AppliedCouponInputAppliedCoupon`

NewAppliedCouponInputAppliedCoupon instantiates a new AppliedCouponInputAppliedCoupon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppliedCouponInputAppliedCouponWithDefaults

`func NewAppliedCouponInputAppliedCouponWithDefaults() *AppliedCouponInputAppliedCoupon`

NewAppliedCouponInputAppliedCouponWithDefaults instantiates a new AppliedCouponInputAppliedCoupon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalCustomerId

`func (o *AppliedCouponInputAppliedCoupon) GetExternalCustomerId() string`

GetExternalCustomerId returns the ExternalCustomerId field if non-nil, zero value otherwise.

### GetExternalCustomerIdOk

`func (o *AppliedCouponInputAppliedCoupon) GetExternalCustomerIdOk() (*string, bool)`

GetExternalCustomerIdOk returns a tuple with the ExternalCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalCustomerId

`func (o *AppliedCouponInputAppliedCoupon) SetExternalCustomerId(v string)`

SetExternalCustomerId sets ExternalCustomerId field to given value.


### GetCouponCode

`func (o *AppliedCouponInputAppliedCoupon) GetCouponCode() string`

GetCouponCode returns the CouponCode field if non-nil, zero value otherwise.

### GetCouponCodeOk

`func (o *AppliedCouponInputAppliedCoupon) GetCouponCodeOk() (*string, bool)`

GetCouponCodeOk returns a tuple with the CouponCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponCode

`func (o *AppliedCouponInputAppliedCoupon) SetCouponCode(v string)`

SetCouponCode sets CouponCode field to given value.


### GetFrequency

`func (o *AppliedCouponInputAppliedCoupon) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *AppliedCouponInputAppliedCoupon) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *AppliedCouponInputAppliedCoupon) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *AppliedCouponInputAppliedCoupon) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### SetFrequencyNil

`func (o *AppliedCouponInputAppliedCoupon) SetFrequencyNil(b bool)`

 SetFrequencyNil sets the value for Frequency to be an explicit nil

### UnsetFrequency
`func (o *AppliedCouponInputAppliedCoupon) UnsetFrequency()`

UnsetFrequency ensures that no value is present for Frequency, not even an explicit nil
### GetFrequencyDuration

`func (o *AppliedCouponInputAppliedCoupon) GetFrequencyDuration() int32`

GetFrequencyDuration returns the FrequencyDuration field if non-nil, zero value otherwise.

### GetFrequencyDurationOk

`func (o *AppliedCouponInputAppliedCoupon) GetFrequencyDurationOk() (*int32, bool)`

GetFrequencyDurationOk returns a tuple with the FrequencyDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencyDuration

`func (o *AppliedCouponInputAppliedCoupon) SetFrequencyDuration(v int32)`

SetFrequencyDuration sets FrequencyDuration field to given value.

### HasFrequencyDuration

`func (o *AppliedCouponInputAppliedCoupon) HasFrequencyDuration() bool`

HasFrequencyDuration returns a boolean if a field has been set.

### SetFrequencyDurationNil

`func (o *AppliedCouponInputAppliedCoupon) SetFrequencyDurationNil(b bool)`

 SetFrequencyDurationNil sets the value for FrequencyDuration to be an explicit nil

### UnsetFrequencyDuration
`func (o *AppliedCouponInputAppliedCoupon) UnsetFrequencyDuration()`

UnsetFrequencyDuration ensures that no value is present for FrequencyDuration, not even an explicit nil
### GetAmountCents

`func (o *AppliedCouponInputAppliedCoupon) GetAmountCents() int32`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *AppliedCouponInputAppliedCoupon) GetAmountCentsOk() (*int32, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *AppliedCouponInputAppliedCoupon) SetAmountCents(v int32)`

SetAmountCents sets AmountCents field to given value.

### HasAmountCents

`func (o *AppliedCouponInputAppliedCoupon) HasAmountCents() bool`

HasAmountCents returns a boolean if a field has been set.

### SetAmountCentsNil

`func (o *AppliedCouponInputAppliedCoupon) SetAmountCentsNil(b bool)`

 SetAmountCentsNil sets the value for AmountCents to be an explicit nil

### UnsetAmountCents
`func (o *AppliedCouponInputAppliedCoupon) UnsetAmountCents()`

UnsetAmountCents ensures that no value is present for AmountCents, not even an explicit nil
### GetAmountCurrency

`func (o *AppliedCouponInputAppliedCoupon) GetAmountCurrency() Currency`

GetAmountCurrency returns the AmountCurrency field if non-nil, zero value otherwise.

### GetAmountCurrencyOk

`func (o *AppliedCouponInputAppliedCoupon) GetAmountCurrencyOk() (*Currency, bool)`

GetAmountCurrencyOk returns a tuple with the AmountCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCurrency

`func (o *AppliedCouponInputAppliedCoupon) SetAmountCurrency(v Currency)`

SetAmountCurrency sets AmountCurrency field to given value.

### HasAmountCurrency

`func (o *AppliedCouponInputAppliedCoupon) HasAmountCurrency() bool`

HasAmountCurrency returns a boolean if a field has been set.

### GetPercentageRate

`func (o *AppliedCouponInputAppliedCoupon) GetPercentageRate() string`

GetPercentageRate returns the PercentageRate field if non-nil, zero value otherwise.

### GetPercentageRateOk

`func (o *AppliedCouponInputAppliedCoupon) GetPercentageRateOk() (*string, bool)`

GetPercentageRateOk returns a tuple with the PercentageRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageRate

`func (o *AppliedCouponInputAppliedCoupon) SetPercentageRate(v string)`

SetPercentageRate sets PercentageRate field to given value.

### HasPercentageRate

`func (o *AppliedCouponInputAppliedCoupon) HasPercentageRate() bool`

HasPercentageRate returns a boolean if a field has been set.

### SetPercentageRateNil

`func (o *AppliedCouponInputAppliedCoupon) SetPercentageRateNil(b bool)`

 SetPercentageRateNil sets the value for PercentageRate to be an explicit nil

### UnsetPercentageRate
`func (o *AppliedCouponInputAppliedCoupon) UnsetPercentageRate()`

UnsetPercentageRate ensures that no value is present for PercentageRate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


