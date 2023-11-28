# AppliedCouponObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LagoId** | **string** | Unique identifier of the applied coupon, created by Lago. | 
**LagoCouponId** | **string** | Unique identifier of the coupon, created by Lago. | 
**CouponCode** | **string** | Unique code used to identify the coupon. | 
**CouponName** | **string** | The name of the coupon. | 
**LagoCustomerId** | **string** | Unique identifier of the customer, created by Lago. | 
**ExternalCustomerId** | **string** | The customer external unique identifier (provided by your own application) | 
**Status** | **string** | The status of the coupon. Can be either &#x60;active&#x60; or &#x60;terminated&#x60;. | 
**AmountCents** | Pointer to **NullableInt32** | The amount of the coupon in cents. This field is required only for coupon with &#x60;fixed_amount&#x60; type. | [optional] 
**AmountCentsRemaining** | Pointer to **NullableInt32** | The remaining amount in cents for a &#x60;fixed_amount&#x60; coupon with a frequency set to &#x60;once&#x60;. This field indicates the remaining balance or value that can still be utilized from the coupon. | [optional] 
**AmountCurrency** | Pointer to [**Currency**](Currency.md) |  | [optional] 
**PercentageRate** | Pointer to **NullableString** | The percentage rate of the coupon. This field is required only for coupons with a &#x60;percentage&#x60; coupon type. | [optional] 
**Frequency** | **string** | The type of frequency for the coupon. It can have three possible values: &#x60;once&#x60;, &#x60;recurring&#x60; or &#x60;forever&#x60;.  - If set to &#x60;once&#x60;, the coupon is applicable only for a single use. - If set to &#x60;recurring&#x60;, the coupon can be used multiple times for recurring billing periods. - If set to &#x60;forever&#x60;, the coupon has unlimited usage and can be applied indefinitely. | 
**FrequencyDuration** | Pointer to **NullableInt32** | Specifies the number of billing periods to which the coupon applies. This field is required only for coupons with a &#x60;recurring&#x60; frequency type | [optional] 
**FrequencyDurationRemaining** | Pointer to **NullableInt32** | The remaining number of billing periods to which the coupon is applicable. This field determines the remaining usage or availability of the coupon based on the remaining billing periods. | [optional] 
**ExpirationAt** | Pointer to **NullableTime** | The date and time after which the coupon will stop applying to customer&#39;s invoices. Once the expiration date is reached, the coupon will no longer be applicable, and any further invoices generated for the customer will not include the coupon discount. | [optional] 
**CreatedAt** | **time.Time** | The date and time when the coupon was assigned to a customer. It is expressed in UTC format according to the ISO 8601 datetime standard. | 
**TerminatedAt** | Pointer to **NullableTime** | This field indicates the specific moment when the coupon amount is fully utilized or when the coupon is removed from the customer&#39;s coupon list. It is expressed in UTC format according to the ISO 8601 datetime standard. | [optional] 

## Methods

### NewAppliedCouponObject

`func NewAppliedCouponObject(lagoId string, lagoCouponId string, couponCode string, couponName string, lagoCustomerId string, externalCustomerId string, status string, frequency string, createdAt time.Time, ) *AppliedCouponObject`

NewAppliedCouponObject instantiates a new AppliedCouponObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppliedCouponObjectWithDefaults

`func NewAppliedCouponObjectWithDefaults() *AppliedCouponObject`

NewAppliedCouponObjectWithDefaults instantiates a new AppliedCouponObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLagoId

`func (o *AppliedCouponObject) GetLagoId() string`

GetLagoId returns the LagoId field if non-nil, zero value otherwise.

### GetLagoIdOk

`func (o *AppliedCouponObject) GetLagoIdOk() (*string, bool)`

GetLagoIdOk returns a tuple with the LagoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagoId

`func (o *AppliedCouponObject) SetLagoId(v string)`

SetLagoId sets LagoId field to given value.


### GetLagoCouponId

`func (o *AppliedCouponObject) GetLagoCouponId() string`

GetLagoCouponId returns the LagoCouponId field if non-nil, zero value otherwise.

### GetLagoCouponIdOk

`func (o *AppliedCouponObject) GetLagoCouponIdOk() (*string, bool)`

GetLagoCouponIdOk returns a tuple with the LagoCouponId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagoCouponId

`func (o *AppliedCouponObject) SetLagoCouponId(v string)`

SetLagoCouponId sets LagoCouponId field to given value.


### GetCouponCode

`func (o *AppliedCouponObject) GetCouponCode() string`

GetCouponCode returns the CouponCode field if non-nil, zero value otherwise.

### GetCouponCodeOk

`func (o *AppliedCouponObject) GetCouponCodeOk() (*string, bool)`

GetCouponCodeOk returns a tuple with the CouponCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponCode

`func (o *AppliedCouponObject) SetCouponCode(v string)`

SetCouponCode sets CouponCode field to given value.


### GetCouponName

`func (o *AppliedCouponObject) GetCouponName() string`

GetCouponName returns the CouponName field if non-nil, zero value otherwise.

### GetCouponNameOk

`func (o *AppliedCouponObject) GetCouponNameOk() (*string, bool)`

GetCouponNameOk returns a tuple with the CouponName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponName

`func (o *AppliedCouponObject) SetCouponName(v string)`

SetCouponName sets CouponName field to given value.


### GetLagoCustomerId

`func (o *AppliedCouponObject) GetLagoCustomerId() string`

GetLagoCustomerId returns the LagoCustomerId field if non-nil, zero value otherwise.

### GetLagoCustomerIdOk

`func (o *AppliedCouponObject) GetLagoCustomerIdOk() (*string, bool)`

GetLagoCustomerIdOk returns a tuple with the LagoCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagoCustomerId

`func (o *AppliedCouponObject) SetLagoCustomerId(v string)`

SetLagoCustomerId sets LagoCustomerId field to given value.


### GetExternalCustomerId

`func (o *AppliedCouponObject) GetExternalCustomerId() string`

GetExternalCustomerId returns the ExternalCustomerId field if non-nil, zero value otherwise.

### GetExternalCustomerIdOk

`func (o *AppliedCouponObject) GetExternalCustomerIdOk() (*string, bool)`

GetExternalCustomerIdOk returns a tuple with the ExternalCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalCustomerId

`func (o *AppliedCouponObject) SetExternalCustomerId(v string)`

SetExternalCustomerId sets ExternalCustomerId field to given value.


### GetStatus

`func (o *AppliedCouponObject) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AppliedCouponObject) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AppliedCouponObject) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetAmountCents

`func (o *AppliedCouponObject) GetAmountCents() int32`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *AppliedCouponObject) GetAmountCentsOk() (*int32, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *AppliedCouponObject) SetAmountCents(v int32)`

SetAmountCents sets AmountCents field to given value.

### HasAmountCents

`func (o *AppliedCouponObject) HasAmountCents() bool`

HasAmountCents returns a boolean if a field has been set.

### SetAmountCentsNil

`func (o *AppliedCouponObject) SetAmountCentsNil(b bool)`

 SetAmountCentsNil sets the value for AmountCents to be an explicit nil

### UnsetAmountCents
`func (o *AppliedCouponObject) UnsetAmountCents()`

UnsetAmountCents ensures that no value is present for AmountCents, not even an explicit nil
### GetAmountCentsRemaining

`func (o *AppliedCouponObject) GetAmountCentsRemaining() int32`

GetAmountCentsRemaining returns the AmountCentsRemaining field if non-nil, zero value otherwise.

### GetAmountCentsRemainingOk

`func (o *AppliedCouponObject) GetAmountCentsRemainingOk() (*int32, bool)`

GetAmountCentsRemainingOk returns a tuple with the AmountCentsRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCentsRemaining

`func (o *AppliedCouponObject) SetAmountCentsRemaining(v int32)`

SetAmountCentsRemaining sets AmountCentsRemaining field to given value.

### HasAmountCentsRemaining

`func (o *AppliedCouponObject) HasAmountCentsRemaining() bool`

HasAmountCentsRemaining returns a boolean if a field has been set.

### SetAmountCentsRemainingNil

`func (o *AppliedCouponObject) SetAmountCentsRemainingNil(b bool)`

 SetAmountCentsRemainingNil sets the value for AmountCentsRemaining to be an explicit nil

### UnsetAmountCentsRemaining
`func (o *AppliedCouponObject) UnsetAmountCentsRemaining()`

UnsetAmountCentsRemaining ensures that no value is present for AmountCentsRemaining, not even an explicit nil
### GetAmountCurrency

`func (o *AppliedCouponObject) GetAmountCurrency() Currency`

GetAmountCurrency returns the AmountCurrency field if non-nil, zero value otherwise.

### GetAmountCurrencyOk

`func (o *AppliedCouponObject) GetAmountCurrencyOk() (*Currency, bool)`

GetAmountCurrencyOk returns a tuple with the AmountCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCurrency

`func (o *AppliedCouponObject) SetAmountCurrency(v Currency)`

SetAmountCurrency sets AmountCurrency field to given value.

### HasAmountCurrency

`func (o *AppliedCouponObject) HasAmountCurrency() bool`

HasAmountCurrency returns a boolean if a field has been set.

### GetPercentageRate

`func (o *AppliedCouponObject) GetPercentageRate() string`

GetPercentageRate returns the PercentageRate field if non-nil, zero value otherwise.

### GetPercentageRateOk

`func (o *AppliedCouponObject) GetPercentageRateOk() (*string, bool)`

GetPercentageRateOk returns a tuple with the PercentageRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageRate

`func (o *AppliedCouponObject) SetPercentageRate(v string)`

SetPercentageRate sets PercentageRate field to given value.

### HasPercentageRate

`func (o *AppliedCouponObject) HasPercentageRate() bool`

HasPercentageRate returns a boolean if a field has been set.

### SetPercentageRateNil

`func (o *AppliedCouponObject) SetPercentageRateNil(b bool)`

 SetPercentageRateNil sets the value for PercentageRate to be an explicit nil

### UnsetPercentageRate
`func (o *AppliedCouponObject) UnsetPercentageRate()`

UnsetPercentageRate ensures that no value is present for PercentageRate, not even an explicit nil
### GetFrequency

`func (o *AppliedCouponObject) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *AppliedCouponObject) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *AppliedCouponObject) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.


### GetFrequencyDuration

`func (o *AppliedCouponObject) GetFrequencyDuration() int32`

GetFrequencyDuration returns the FrequencyDuration field if non-nil, zero value otherwise.

### GetFrequencyDurationOk

`func (o *AppliedCouponObject) GetFrequencyDurationOk() (*int32, bool)`

GetFrequencyDurationOk returns a tuple with the FrequencyDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencyDuration

`func (o *AppliedCouponObject) SetFrequencyDuration(v int32)`

SetFrequencyDuration sets FrequencyDuration field to given value.

### HasFrequencyDuration

`func (o *AppliedCouponObject) HasFrequencyDuration() bool`

HasFrequencyDuration returns a boolean if a field has been set.

### SetFrequencyDurationNil

`func (o *AppliedCouponObject) SetFrequencyDurationNil(b bool)`

 SetFrequencyDurationNil sets the value for FrequencyDuration to be an explicit nil

### UnsetFrequencyDuration
`func (o *AppliedCouponObject) UnsetFrequencyDuration()`

UnsetFrequencyDuration ensures that no value is present for FrequencyDuration, not even an explicit nil
### GetFrequencyDurationRemaining

`func (o *AppliedCouponObject) GetFrequencyDurationRemaining() int32`

GetFrequencyDurationRemaining returns the FrequencyDurationRemaining field if non-nil, zero value otherwise.

### GetFrequencyDurationRemainingOk

`func (o *AppliedCouponObject) GetFrequencyDurationRemainingOk() (*int32, bool)`

GetFrequencyDurationRemainingOk returns a tuple with the FrequencyDurationRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencyDurationRemaining

`func (o *AppliedCouponObject) SetFrequencyDurationRemaining(v int32)`

SetFrequencyDurationRemaining sets FrequencyDurationRemaining field to given value.

### HasFrequencyDurationRemaining

`func (o *AppliedCouponObject) HasFrequencyDurationRemaining() bool`

HasFrequencyDurationRemaining returns a boolean if a field has been set.

### SetFrequencyDurationRemainingNil

`func (o *AppliedCouponObject) SetFrequencyDurationRemainingNil(b bool)`

 SetFrequencyDurationRemainingNil sets the value for FrequencyDurationRemaining to be an explicit nil

### UnsetFrequencyDurationRemaining
`func (o *AppliedCouponObject) UnsetFrequencyDurationRemaining()`

UnsetFrequencyDurationRemaining ensures that no value is present for FrequencyDurationRemaining, not even an explicit nil
### GetExpirationAt

`func (o *AppliedCouponObject) GetExpirationAt() time.Time`

GetExpirationAt returns the ExpirationAt field if non-nil, zero value otherwise.

### GetExpirationAtOk

`func (o *AppliedCouponObject) GetExpirationAtOk() (*time.Time, bool)`

GetExpirationAtOk returns a tuple with the ExpirationAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationAt

`func (o *AppliedCouponObject) SetExpirationAt(v time.Time)`

SetExpirationAt sets ExpirationAt field to given value.

### HasExpirationAt

`func (o *AppliedCouponObject) HasExpirationAt() bool`

HasExpirationAt returns a boolean if a field has been set.

### SetExpirationAtNil

`func (o *AppliedCouponObject) SetExpirationAtNil(b bool)`

 SetExpirationAtNil sets the value for ExpirationAt to be an explicit nil

### UnsetExpirationAt
`func (o *AppliedCouponObject) UnsetExpirationAt()`

UnsetExpirationAt ensures that no value is present for ExpirationAt, not even an explicit nil
### GetCreatedAt

`func (o *AppliedCouponObject) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AppliedCouponObject) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AppliedCouponObject) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetTerminatedAt

`func (o *AppliedCouponObject) GetTerminatedAt() time.Time`

GetTerminatedAt returns the TerminatedAt field if non-nil, zero value otherwise.

### GetTerminatedAtOk

`func (o *AppliedCouponObject) GetTerminatedAtOk() (*time.Time, bool)`

GetTerminatedAtOk returns a tuple with the TerminatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminatedAt

`func (o *AppliedCouponObject) SetTerminatedAt(v time.Time)`

SetTerminatedAt sets TerminatedAt field to given value.

### HasTerminatedAt

`func (o *AppliedCouponObject) HasTerminatedAt() bool`

HasTerminatedAt returns a boolean if a field has been set.

### SetTerminatedAtNil

`func (o *AppliedCouponObject) SetTerminatedAtNil(b bool)`

 SetTerminatedAtNil sets the value for TerminatedAt to be an explicit nil

### UnsetTerminatedAt
`func (o *AppliedCouponObject) UnsetTerminatedAt()`

UnsetTerminatedAt ensures that no value is present for TerminatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


