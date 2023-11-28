# CouponCreateInputCoupon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the coupon. | 
**Code** | **string** | Unique code used to identify the coupon. | 
**Description** | Pointer to **NullableString** | Description of the coupon. | [optional] 
**CouponType** | **string** | The type of the coupon. It can have two possible values: &#x60;fixed_amount&#x60; or &#x60;percentage&#x60;.  - If set to &#x60;fixed_amount&#x60;, the coupon represents a fixed amount discount. - If set to &#x60;percentage&#x60;, the coupon represents a percentage-based discount. | 
**AmountCents** | Pointer to **NullableInt32** | The amount of the coupon in cents. This field is required only for coupon with &#x60;fixed_amount&#x60; type. | [optional] 
**AmountCurrency** | Pointer to [**Currency**](Currency.md) |  | [optional] 
**Reusable** | Pointer to **bool** | Indicates whether the coupon can be reused or not. If set to &#x60;true&#x60;, the coupon is reusable, meaning it can be applied multiple times to the same customer. If set to &#x60;false&#x60;, the coupon can only be used once and is not reusable. If not specified, this field is set to &#x60;true&#x60; by default. | [optional] 
**PercentageRate** | Pointer to **NullableString** | The percentage rate of the coupon. This field is required only for coupons with a &#x60;percentage&#x60; coupon type. | [optional] 
**Frequency** | **string** | The type of frequency for the coupon. It can have three possible values: &#x60;once&#x60;, &#x60;recurring&#x60; or &#x60;forever&#x60;.  - If set to &#x60;once&#x60;, the coupon is applicable only for a single use. - If set to &#x60;recurring&#x60;, the coupon can be used multiple times for recurring billing periods. - If set to &#x60;forever&#x60;, the coupon has unlimited usage and can be applied indefinitely. | 
**FrequencyDuration** | Pointer to **NullableInt32** | Specifies the number of billing periods to which the coupon applies. This field is required only for coupons with a &#x60;recurring&#x60; frequency type | [optional] 
**Expiration** | Pointer to **string** | Specifies the type of expiration for the coupon. It can have two possible values: &#x60;time_limit&#x60; or &#x60;no_expiration&#x60;.  - If set to &#x60;time_limit&#x60;, the coupon has an expiration based on a specified time limit. - If set to &#x60;no_expiration&#x60;, the coupon does not have an expiration date and remains valid indefinitely. | [optional] 
**ExpirationAt** | Pointer to **NullableTime** | The expiration date and time of the coupon. This field is required only for coupons with &#x60;expiration&#x60; set to &#x60;time_limit&#x60;. The expiration date and time should be specified in UTC format according to the ISO 8601 datetime standard. It indicates the exact moment when the coupon will expire and is no longer valid. | [optional] 
**AppliesTo** | Pointer to [**NullableCouponBaseInputAppliesTo**](CouponBaseInputAppliesTo.md) |  | [optional] 

## Methods

### NewCouponCreateInputCoupon

`func NewCouponCreateInputCoupon(name string, code string, couponType string, frequency string, ) *CouponCreateInputCoupon`

NewCouponCreateInputCoupon instantiates a new CouponCreateInputCoupon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCouponCreateInputCouponWithDefaults

`func NewCouponCreateInputCouponWithDefaults() *CouponCreateInputCoupon`

NewCouponCreateInputCouponWithDefaults instantiates a new CouponCreateInputCoupon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CouponCreateInputCoupon) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CouponCreateInputCoupon) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CouponCreateInputCoupon) SetName(v string)`

SetName sets Name field to given value.


### GetCode

`func (o *CouponCreateInputCoupon) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CouponCreateInputCoupon) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CouponCreateInputCoupon) SetCode(v string)`

SetCode sets Code field to given value.


### GetDescription

`func (o *CouponCreateInputCoupon) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CouponCreateInputCoupon) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CouponCreateInputCoupon) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CouponCreateInputCoupon) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CouponCreateInputCoupon) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CouponCreateInputCoupon) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCouponType

`func (o *CouponCreateInputCoupon) GetCouponType() string`

GetCouponType returns the CouponType field if non-nil, zero value otherwise.

### GetCouponTypeOk

`func (o *CouponCreateInputCoupon) GetCouponTypeOk() (*string, bool)`

GetCouponTypeOk returns a tuple with the CouponType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponType

`func (o *CouponCreateInputCoupon) SetCouponType(v string)`

SetCouponType sets CouponType field to given value.


### GetAmountCents

`func (o *CouponCreateInputCoupon) GetAmountCents() int32`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *CouponCreateInputCoupon) GetAmountCentsOk() (*int32, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *CouponCreateInputCoupon) SetAmountCents(v int32)`

SetAmountCents sets AmountCents field to given value.

### HasAmountCents

`func (o *CouponCreateInputCoupon) HasAmountCents() bool`

HasAmountCents returns a boolean if a field has been set.

### SetAmountCentsNil

`func (o *CouponCreateInputCoupon) SetAmountCentsNil(b bool)`

 SetAmountCentsNil sets the value for AmountCents to be an explicit nil

### UnsetAmountCents
`func (o *CouponCreateInputCoupon) UnsetAmountCents()`

UnsetAmountCents ensures that no value is present for AmountCents, not even an explicit nil
### GetAmountCurrency

`func (o *CouponCreateInputCoupon) GetAmountCurrency() Currency`

GetAmountCurrency returns the AmountCurrency field if non-nil, zero value otherwise.

### GetAmountCurrencyOk

`func (o *CouponCreateInputCoupon) GetAmountCurrencyOk() (*Currency, bool)`

GetAmountCurrencyOk returns a tuple with the AmountCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCurrency

`func (o *CouponCreateInputCoupon) SetAmountCurrency(v Currency)`

SetAmountCurrency sets AmountCurrency field to given value.

### HasAmountCurrency

`func (o *CouponCreateInputCoupon) HasAmountCurrency() bool`

HasAmountCurrency returns a boolean if a field has been set.

### GetReusable

`func (o *CouponCreateInputCoupon) GetReusable() bool`

GetReusable returns the Reusable field if non-nil, zero value otherwise.

### GetReusableOk

`func (o *CouponCreateInputCoupon) GetReusableOk() (*bool, bool)`

GetReusableOk returns a tuple with the Reusable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReusable

`func (o *CouponCreateInputCoupon) SetReusable(v bool)`

SetReusable sets Reusable field to given value.

### HasReusable

`func (o *CouponCreateInputCoupon) HasReusable() bool`

HasReusable returns a boolean if a field has been set.

### GetPercentageRate

`func (o *CouponCreateInputCoupon) GetPercentageRate() string`

GetPercentageRate returns the PercentageRate field if non-nil, zero value otherwise.

### GetPercentageRateOk

`func (o *CouponCreateInputCoupon) GetPercentageRateOk() (*string, bool)`

GetPercentageRateOk returns a tuple with the PercentageRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageRate

`func (o *CouponCreateInputCoupon) SetPercentageRate(v string)`

SetPercentageRate sets PercentageRate field to given value.

### HasPercentageRate

`func (o *CouponCreateInputCoupon) HasPercentageRate() bool`

HasPercentageRate returns a boolean if a field has been set.

### SetPercentageRateNil

`func (o *CouponCreateInputCoupon) SetPercentageRateNil(b bool)`

 SetPercentageRateNil sets the value for PercentageRate to be an explicit nil

### UnsetPercentageRate
`func (o *CouponCreateInputCoupon) UnsetPercentageRate()`

UnsetPercentageRate ensures that no value is present for PercentageRate, not even an explicit nil
### GetFrequency

`func (o *CouponCreateInputCoupon) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *CouponCreateInputCoupon) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *CouponCreateInputCoupon) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.


### GetFrequencyDuration

`func (o *CouponCreateInputCoupon) GetFrequencyDuration() int32`

GetFrequencyDuration returns the FrequencyDuration field if non-nil, zero value otherwise.

### GetFrequencyDurationOk

`func (o *CouponCreateInputCoupon) GetFrequencyDurationOk() (*int32, bool)`

GetFrequencyDurationOk returns a tuple with the FrequencyDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencyDuration

`func (o *CouponCreateInputCoupon) SetFrequencyDuration(v int32)`

SetFrequencyDuration sets FrequencyDuration field to given value.

### HasFrequencyDuration

`func (o *CouponCreateInputCoupon) HasFrequencyDuration() bool`

HasFrequencyDuration returns a boolean if a field has been set.

### SetFrequencyDurationNil

`func (o *CouponCreateInputCoupon) SetFrequencyDurationNil(b bool)`

 SetFrequencyDurationNil sets the value for FrequencyDuration to be an explicit nil

### UnsetFrequencyDuration
`func (o *CouponCreateInputCoupon) UnsetFrequencyDuration()`

UnsetFrequencyDuration ensures that no value is present for FrequencyDuration, not even an explicit nil
### GetExpiration

`func (o *CouponCreateInputCoupon) GetExpiration() string`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *CouponCreateInputCoupon) GetExpirationOk() (*string, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *CouponCreateInputCoupon) SetExpiration(v string)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *CouponCreateInputCoupon) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetExpirationAt

`func (o *CouponCreateInputCoupon) GetExpirationAt() time.Time`

GetExpirationAt returns the ExpirationAt field if non-nil, zero value otherwise.

### GetExpirationAtOk

`func (o *CouponCreateInputCoupon) GetExpirationAtOk() (*time.Time, bool)`

GetExpirationAtOk returns a tuple with the ExpirationAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationAt

`func (o *CouponCreateInputCoupon) SetExpirationAt(v time.Time)`

SetExpirationAt sets ExpirationAt field to given value.

### HasExpirationAt

`func (o *CouponCreateInputCoupon) HasExpirationAt() bool`

HasExpirationAt returns a boolean if a field has been set.

### SetExpirationAtNil

`func (o *CouponCreateInputCoupon) SetExpirationAtNil(b bool)`

 SetExpirationAtNil sets the value for ExpirationAt to be an explicit nil

### UnsetExpirationAt
`func (o *CouponCreateInputCoupon) UnsetExpirationAt()`

UnsetExpirationAt ensures that no value is present for ExpirationAt, not even an explicit nil
### GetAppliesTo

`func (o *CouponCreateInputCoupon) GetAppliesTo() CouponBaseInputAppliesTo`

GetAppliesTo returns the AppliesTo field if non-nil, zero value otherwise.

### GetAppliesToOk

`func (o *CouponCreateInputCoupon) GetAppliesToOk() (*CouponBaseInputAppliesTo, bool)`

GetAppliesToOk returns a tuple with the AppliesTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliesTo

`func (o *CouponCreateInputCoupon) SetAppliesTo(v CouponBaseInputAppliesTo)`

SetAppliesTo sets AppliesTo field to given value.

### HasAppliesTo

`func (o *CouponCreateInputCoupon) HasAppliesTo() bool`

HasAppliesTo returns a boolean if a field has been set.

### SetAppliesToNil

`func (o *CouponCreateInputCoupon) SetAppliesToNil(b bool)`

 SetAppliesToNil sets the value for AppliesTo to be an explicit nil

### UnsetAppliesTo
`func (o *CouponCreateInputCoupon) UnsetAppliesTo()`

UnsetAppliesTo ensures that no value is present for AppliesTo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


