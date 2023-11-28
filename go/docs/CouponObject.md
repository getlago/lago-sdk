# CouponObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LagoId** | **string** | Unique identifier of the coupon, created by Lago. | 
**Name** | **string** | The name of the coupon. | 
**Code** | **string** | Unique code used to identify the coupon. | 
**Description** | Pointer to **NullableString** | Description of the coupon. | [optional] 
**CouponType** | **string** | The type of the coupon. It can have two possible values: &#x60;fixed_amount&#x60; or &#x60;percentage&#x60;.  - If set to &#x60;fixed_amount&#x60;, the coupon represents a fixed amount discount. - If set to &#x60;percentage&#x60;, the coupon represents a percentage-based discount. | 
**AmountCents** | Pointer to **NullableInt32** | The amount of the coupon in cents. This field is required only for coupon with &#x60;fixed_amount&#x60; type. | [optional] 
**AmountCurrency** | Pointer to [**Currency**](Currency.md) |  | [optional] 
**Reusable** | **bool** | Indicates whether the coupon can be reused or not. If set to &#x60;true&#x60;, the coupon is reusable, meaning it can be applied multiple times to the same customer. If set to &#x60;false&#x60;, the coupon can only be used once and is not reusable. If not specified, this field is set to &#x60;true&#x60; by default. | 
**LimitedPlans** | **bool** | The coupon is limited to specific plans. The possible values can be &#x60;true&#x60; or &#x60;false&#x60;. | 
**PlanCodes** | Pointer to **[]string** | An array of plan codes to which the coupon is applicable. By specifying the plan codes in this field, you can restrict the coupon&#39;s usage to specific plans only. | [optional] 
**LimitedBillableMetrics** | **bool** | The coupon is limited to specific billable metrics. The possible values can be &#x60;true&#x60; or &#x60;false&#x60;. | 
**BillableMetricCodes** | Pointer to **[]string** | An array of billable metric codes to which the coupon is applicable. By specifying the billable metric codes in this field, you can restrict the coupon&#39;s usage to specific metrics only. | [optional] 
**PercentageRate** | Pointer to **NullableString** | The percentage rate of the coupon. This field is required only for coupons with a &#x60;percentage&#x60; coupon type. | [optional] 
**Frequency** | **string** | The type of frequency for the coupon. It can have three possible values: &#x60;once&#x60;, &#x60;recurring&#x60;, or &#x60;forever&#x60;.  - If set to &#x60;once&#x60;, the coupon is applicable only for a single use. - If set to &#x60;recurring&#x60;, the coupon can be used multiple times for recurring billing periods. - If set to &#x60;forever&#x60;, the coupon has unlimited usage and can be applied indefinitely. | 
**FrequencyDuration** | Pointer to **NullableInt32** | Specifies the number of billing periods to which the coupon applies. This field is required only for coupons with a &#x60;recurring&#x60; frequency type | [optional] 
**Expiration** | **string** | Specifies the type of expiration for the coupon. It can have two possible values: &#x60;time_limit&#x60; or &#x60;no_expiration&#x60;.  - If set to &#x60;time_limit&#x60;, the coupon has an expiration based on a specified time limit. - If set to &#x60;no_expiration&#x60;, the coupon does not have an expiration date and remains valid indefinitely. | 
**ExpirationAt** | Pointer to **NullableTime** | The expiration date and time of the coupon. This field is required only for coupons with &#x60;expiration&#x60; set to &#x60;time_limit&#x60;. The expiration date and time should be specified in UTC format according to the ISO 8601 datetime standard. It indicates the exact moment when the coupon will expire and is no longer valid. | [optional] 
**CreatedAt** | **time.Time** | The date and time when the coupon was created. It is expressed in UTC format according to the ISO 8601 datetime standard. This field provides the timestamp for the exact moment when the coupon was initially created. | 
**TerminatedAt** | Pointer to **time.Time** | This field indicates if the coupon has been terminated and is no longer usable. If it&#39;s not null, it won&#39;t be removed for existing customers using it, but it prevents the coupon from being applied in the future. | [optional] 

## Methods

### NewCouponObject

`func NewCouponObject(lagoId string, name string, code string, couponType string, reusable bool, limitedPlans bool, limitedBillableMetrics bool, frequency string, expiration string, createdAt time.Time, ) *CouponObject`

NewCouponObject instantiates a new CouponObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCouponObjectWithDefaults

`func NewCouponObjectWithDefaults() *CouponObject`

NewCouponObjectWithDefaults instantiates a new CouponObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLagoId

`func (o *CouponObject) GetLagoId() string`

GetLagoId returns the LagoId field if non-nil, zero value otherwise.

### GetLagoIdOk

`func (o *CouponObject) GetLagoIdOk() (*string, bool)`

GetLagoIdOk returns a tuple with the LagoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagoId

`func (o *CouponObject) SetLagoId(v string)`

SetLagoId sets LagoId field to given value.


### GetName

`func (o *CouponObject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CouponObject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CouponObject) SetName(v string)`

SetName sets Name field to given value.


### GetCode

`func (o *CouponObject) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CouponObject) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CouponObject) SetCode(v string)`

SetCode sets Code field to given value.


### GetDescription

`func (o *CouponObject) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CouponObject) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CouponObject) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CouponObject) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CouponObject) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CouponObject) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCouponType

`func (o *CouponObject) GetCouponType() string`

GetCouponType returns the CouponType field if non-nil, zero value otherwise.

### GetCouponTypeOk

`func (o *CouponObject) GetCouponTypeOk() (*string, bool)`

GetCouponTypeOk returns a tuple with the CouponType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponType

`func (o *CouponObject) SetCouponType(v string)`

SetCouponType sets CouponType field to given value.


### GetAmountCents

`func (o *CouponObject) GetAmountCents() int32`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *CouponObject) GetAmountCentsOk() (*int32, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *CouponObject) SetAmountCents(v int32)`

SetAmountCents sets AmountCents field to given value.

### HasAmountCents

`func (o *CouponObject) HasAmountCents() bool`

HasAmountCents returns a boolean if a field has been set.

### SetAmountCentsNil

`func (o *CouponObject) SetAmountCentsNil(b bool)`

 SetAmountCentsNil sets the value for AmountCents to be an explicit nil

### UnsetAmountCents
`func (o *CouponObject) UnsetAmountCents()`

UnsetAmountCents ensures that no value is present for AmountCents, not even an explicit nil
### GetAmountCurrency

`func (o *CouponObject) GetAmountCurrency() Currency`

GetAmountCurrency returns the AmountCurrency field if non-nil, zero value otherwise.

### GetAmountCurrencyOk

`func (o *CouponObject) GetAmountCurrencyOk() (*Currency, bool)`

GetAmountCurrencyOk returns a tuple with the AmountCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCurrency

`func (o *CouponObject) SetAmountCurrency(v Currency)`

SetAmountCurrency sets AmountCurrency field to given value.

### HasAmountCurrency

`func (o *CouponObject) HasAmountCurrency() bool`

HasAmountCurrency returns a boolean if a field has been set.

### GetReusable

`func (o *CouponObject) GetReusable() bool`

GetReusable returns the Reusable field if non-nil, zero value otherwise.

### GetReusableOk

`func (o *CouponObject) GetReusableOk() (*bool, bool)`

GetReusableOk returns a tuple with the Reusable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReusable

`func (o *CouponObject) SetReusable(v bool)`

SetReusable sets Reusable field to given value.


### GetLimitedPlans

`func (o *CouponObject) GetLimitedPlans() bool`

GetLimitedPlans returns the LimitedPlans field if non-nil, zero value otherwise.

### GetLimitedPlansOk

`func (o *CouponObject) GetLimitedPlansOk() (*bool, bool)`

GetLimitedPlansOk returns a tuple with the LimitedPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitedPlans

`func (o *CouponObject) SetLimitedPlans(v bool)`

SetLimitedPlans sets LimitedPlans field to given value.


### GetPlanCodes

`func (o *CouponObject) GetPlanCodes() []string`

GetPlanCodes returns the PlanCodes field if non-nil, zero value otherwise.

### GetPlanCodesOk

`func (o *CouponObject) GetPlanCodesOk() (*[]string, bool)`

GetPlanCodesOk returns a tuple with the PlanCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanCodes

`func (o *CouponObject) SetPlanCodes(v []string)`

SetPlanCodes sets PlanCodes field to given value.

### HasPlanCodes

`func (o *CouponObject) HasPlanCodes() bool`

HasPlanCodes returns a boolean if a field has been set.

### GetLimitedBillableMetrics

`func (o *CouponObject) GetLimitedBillableMetrics() bool`

GetLimitedBillableMetrics returns the LimitedBillableMetrics field if non-nil, zero value otherwise.

### GetLimitedBillableMetricsOk

`func (o *CouponObject) GetLimitedBillableMetricsOk() (*bool, bool)`

GetLimitedBillableMetricsOk returns a tuple with the LimitedBillableMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitedBillableMetrics

`func (o *CouponObject) SetLimitedBillableMetrics(v bool)`

SetLimitedBillableMetrics sets LimitedBillableMetrics field to given value.


### GetBillableMetricCodes

`func (o *CouponObject) GetBillableMetricCodes() []string`

GetBillableMetricCodes returns the BillableMetricCodes field if non-nil, zero value otherwise.

### GetBillableMetricCodesOk

`func (o *CouponObject) GetBillableMetricCodesOk() (*[]string, bool)`

GetBillableMetricCodesOk returns a tuple with the BillableMetricCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillableMetricCodes

`func (o *CouponObject) SetBillableMetricCodes(v []string)`

SetBillableMetricCodes sets BillableMetricCodes field to given value.

### HasBillableMetricCodes

`func (o *CouponObject) HasBillableMetricCodes() bool`

HasBillableMetricCodes returns a boolean if a field has been set.

### GetPercentageRate

`func (o *CouponObject) GetPercentageRate() string`

GetPercentageRate returns the PercentageRate field if non-nil, zero value otherwise.

### GetPercentageRateOk

`func (o *CouponObject) GetPercentageRateOk() (*string, bool)`

GetPercentageRateOk returns a tuple with the PercentageRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageRate

`func (o *CouponObject) SetPercentageRate(v string)`

SetPercentageRate sets PercentageRate field to given value.

### HasPercentageRate

`func (o *CouponObject) HasPercentageRate() bool`

HasPercentageRate returns a boolean if a field has been set.

### SetPercentageRateNil

`func (o *CouponObject) SetPercentageRateNil(b bool)`

 SetPercentageRateNil sets the value for PercentageRate to be an explicit nil

### UnsetPercentageRate
`func (o *CouponObject) UnsetPercentageRate()`

UnsetPercentageRate ensures that no value is present for PercentageRate, not even an explicit nil
### GetFrequency

`func (o *CouponObject) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *CouponObject) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *CouponObject) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.


### GetFrequencyDuration

`func (o *CouponObject) GetFrequencyDuration() int32`

GetFrequencyDuration returns the FrequencyDuration field if non-nil, zero value otherwise.

### GetFrequencyDurationOk

`func (o *CouponObject) GetFrequencyDurationOk() (*int32, bool)`

GetFrequencyDurationOk returns a tuple with the FrequencyDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencyDuration

`func (o *CouponObject) SetFrequencyDuration(v int32)`

SetFrequencyDuration sets FrequencyDuration field to given value.

### HasFrequencyDuration

`func (o *CouponObject) HasFrequencyDuration() bool`

HasFrequencyDuration returns a boolean if a field has been set.

### SetFrequencyDurationNil

`func (o *CouponObject) SetFrequencyDurationNil(b bool)`

 SetFrequencyDurationNil sets the value for FrequencyDuration to be an explicit nil

### UnsetFrequencyDuration
`func (o *CouponObject) UnsetFrequencyDuration()`

UnsetFrequencyDuration ensures that no value is present for FrequencyDuration, not even an explicit nil
### GetExpiration

`func (o *CouponObject) GetExpiration() string`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *CouponObject) GetExpirationOk() (*string, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *CouponObject) SetExpiration(v string)`

SetExpiration sets Expiration field to given value.


### GetExpirationAt

`func (o *CouponObject) GetExpirationAt() time.Time`

GetExpirationAt returns the ExpirationAt field if non-nil, zero value otherwise.

### GetExpirationAtOk

`func (o *CouponObject) GetExpirationAtOk() (*time.Time, bool)`

GetExpirationAtOk returns a tuple with the ExpirationAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationAt

`func (o *CouponObject) SetExpirationAt(v time.Time)`

SetExpirationAt sets ExpirationAt field to given value.

### HasExpirationAt

`func (o *CouponObject) HasExpirationAt() bool`

HasExpirationAt returns a boolean if a field has been set.

### SetExpirationAtNil

`func (o *CouponObject) SetExpirationAtNil(b bool)`

 SetExpirationAtNil sets the value for ExpirationAt to be an explicit nil

### UnsetExpirationAt
`func (o *CouponObject) UnsetExpirationAt()`

UnsetExpirationAt ensures that no value is present for ExpirationAt, not even an explicit nil
### GetCreatedAt

`func (o *CouponObject) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CouponObject) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CouponObject) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetTerminatedAt

`func (o *CouponObject) GetTerminatedAt() time.Time`

GetTerminatedAt returns the TerminatedAt field if non-nil, zero value otherwise.

### GetTerminatedAtOk

`func (o *CouponObject) GetTerminatedAtOk() (*time.Time, bool)`

GetTerminatedAtOk returns a tuple with the TerminatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminatedAt

`func (o *CouponObject) SetTerminatedAt(v time.Time)`

SetTerminatedAt sets TerminatedAt field to given value.

### HasTerminatedAt

`func (o *CouponObject) HasTerminatedAt() bool`

HasTerminatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


