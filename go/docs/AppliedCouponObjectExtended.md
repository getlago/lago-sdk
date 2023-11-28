# AppliedCouponObjectExtended

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
**Credits** | [**[]CreditObject**](CreditObject.md) |  | 

## Methods

### NewAppliedCouponObjectExtended

`func NewAppliedCouponObjectExtended(lagoId string, lagoCouponId string, couponCode string, couponName string, lagoCustomerId string, externalCustomerId string, status string, frequency string, createdAt time.Time, credits []CreditObject, ) *AppliedCouponObjectExtended`

NewAppliedCouponObjectExtended instantiates a new AppliedCouponObjectExtended object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppliedCouponObjectExtendedWithDefaults

`func NewAppliedCouponObjectExtendedWithDefaults() *AppliedCouponObjectExtended`

NewAppliedCouponObjectExtendedWithDefaults instantiates a new AppliedCouponObjectExtended object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLagoId

`func (o *AppliedCouponObjectExtended) GetLagoId() string`

GetLagoId returns the LagoId field if non-nil, zero value otherwise.

### GetLagoIdOk

`func (o *AppliedCouponObjectExtended) GetLagoIdOk() (*string, bool)`

GetLagoIdOk returns a tuple with the LagoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagoId

`func (o *AppliedCouponObjectExtended) SetLagoId(v string)`

SetLagoId sets LagoId field to given value.


### GetLagoCouponId

`func (o *AppliedCouponObjectExtended) GetLagoCouponId() string`

GetLagoCouponId returns the LagoCouponId field if non-nil, zero value otherwise.

### GetLagoCouponIdOk

`func (o *AppliedCouponObjectExtended) GetLagoCouponIdOk() (*string, bool)`

GetLagoCouponIdOk returns a tuple with the LagoCouponId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagoCouponId

`func (o *AppliedCouponObjectExtended) SetLagoCouponId(v string)`

SetLagoCouponId sets LagoCouponId field to given value.


### GetCouponCode

`func (o *AppliedCouponObjectExtended) GetCouponCode() string`

GetCouponCode returns the CouponCode field if non-nil, zero value otherwise.

### GetCouponCodeOk

`func (o *AppliedCouponObjectExtended) GetCouponCodeOk() (*string, bool)`

GetCouponCodeOk returns a tuple with the CouponCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponCode

`func (o *AppliedCouponObjectExtended) SetCouponCode(v string)`

SetCouponCode sets CouponCode field to given value.


### GetCouponName

`func (o *AppliedCouponObjectExtended) GetCouponName() string`

GetCouponName returns the CouponName field if non-nil, zero value otherwise.

### GetCouponNameOk

`func (o *AppliedCouponObjectExtended) GetCouponNameOk() (*string, bool)`

GetCouponNameOk returns a tuple with the CouponName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponName

`func (o *AppliedCouponObjectExtended) SetCouponName(v string)`

SetCouponName sets CouponName field to given value.


### GetLagoCustomerId

`func (o *AppliedCouponObjectExtended) GetLagoCustomerId() string`

GetLagoCustomerId returns the LagoCustomerId field if non-nil, zero value otherwise.

### GetLagoCustomerIdOk

`func (o *AppliedCouponObjectExtended) GetLagoCustomerIdOk() (*string, bool)`

GetLagoCustomerIdOk returns a tuple with the LagoCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagoCustomerId

`func (o *AppliedCouponObjectExtended) SetLagoCustomerId(v string)`

SetLagoCustomerId sets LagoCustomerId field to given value.


### GetExternalCustomerId

`func (o *AppliedCouponObjectExtended) GetExternalCustomerId() string`

GetExternalCustomerId returns the ExternalCustomerId field if non-nil, zero value otherwise.

### GetExternalCustomerIdOk

`func (o *AppliedCouponObjectExtended) GetExternalCustomerIdOk() (*string, bool)`

GetExternalCustomerIdOk returns a tuple with the ExternalCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalCustomerId

`func (o *AppliedCouponObjectExtended) SetExternalCustomerId(v string)`

SetExternalCustomerId sets ExternalCustomerId field to given value.


### GetStatus

`func (o *AppliedCouponObjectExtended) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AppliedCouponObjectExtended) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AppliedCouponObjectExtended) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetAmountCents

`func (o *AppliedCouponObjectExtended) GetAmountCents() int32`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *AppliedCouponObjectExtended) GetAmountCentsOk() (*int32, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *AppliedCouponObjectExtended) SetAmountCents(v int32)`

SetAmountCents sets AmountCents field to given value.

### HasAmountCents

`func (o *AppliedCouponObjectExtended) HasAmountCents() bool`

HasAmountCents returns a boolean if a field has been set.

### SetAmountCentsNil

`func (o *AppliedCouponObjectExtended) SetAmountCentsNil(b bool)`

 SetAmountCentsNil sets the value for AmountCents to be an explicit nil

### UnsetAmountCents
`func (o *AppliedCouponObjectExtended) UnsetAmountCents()`

UnsetAmountCents ensures that no value is present for AmountCents, not even an explicit nil
### GetAmountCentsRemaining

`func (o *AppliedCouponObjectExtended) GetAmountCentsRemaining() int32`

GetAmountCentsRemaining returns the AmountCentsRemaining field if non-nil, zero value otherwise.

### GetAmountCentsRemainingOk

`func (o *AppliedCouponObjectExtended) GetAmountCentsRemainingOk() (*int32, bool)`

GetAmountCentsRemainingOk returns a tuple with the AmountCentsRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCentsRemaining

`func (o *AppliedCouponObjectExtended) SetAmountCentsRemaining(v int32)`

SetAmountCentsRemaining sets AmountCentsRemaining field to given value.

### HasAmountCentsRemaining

`func (o *AppliedCouponObjectExtended) HasAmountCentsRemaining() bool`

HasAmountCentsRemaining returns a boolean if a field has been set.

### SetAmountCentsRemainingNil

`func (o *AppliedCouponObjectExtended) SetAmountCentsRemainingNil(b bool)`

 SetAmountCentsRemainingNil sets the value for AmountCentsRemaining to be an explicit nil

### UnsetAmountCentsRemaining
`func (o *AppliedCouponObjectExtended) UnsetAmountCentsRemaining()`

UnsetAmountCentsRemaining ensures that no value is present for AmountCentsRemaining, not even an explicit nil
### GetAmountCurrency

`func (o *AppliedCouponObjectExtended) GetAmountCurrency() Currency`

GetAmountCurrency returns the AmountCurrency field if non-nil, zero value otherwise.

### GetAmountCurrencyOk

`func (o *AppliedCouponObjectExtended) GetAmountCurrencyOk() (*Currency, bool)`

GetAmountCurrencyOk returns a tuple with the AmountCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCurrency

`func (o *AppliedCouponObjectExtended) SetAmountCurrency(v Currency)`

SetAmountCurrency sets AmountCurrency field to given value.

### HasAmountCurrency

`func (o *AppliedCouponObjectExtended) HasAmountCurrency() bool`

HasAmountCurrency returns a boolean if a field has been set.

### GetPercentageRate

`func (o *AppliedCouponObjectExtended) GetPercentageRate() string`

GetPercentageRate returns the PercentageRate field if non-nil, zero value otherwise.

### GetPercentageRateOk

`func (o *AppliedCouponObjectExtended) GetPercentageRateOk() (*string, bool)`

GetPercentageRateOk returns a tuple with the PercentageRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageRate

`func (o *AppliedCouponObjectExtended) SetPercentageRate(v string)`

SetPercentageRate sets PercentageRate field to given value.

### HasPercentageRate

`func (o *AppliedCouponObjectExtended) HasPercentageRate() bool`

HasPercentageRate returns a boolean if a field has been set.

### SetPercentageRateNil

`func (o *AppliedCouponObjectExtended) SetPercentageRateNil(b bool)`

 SetPercentageRateNil sets the value for PercentageRate to be an explicit nil

### UnsetPercentageRate
`func (o *AppliedCouponObjectExtended) UnsetPercentageRate()`

UnsetPercentageRate ensures that no value is present for PercentageRate, not even an explicit nil
### GetFrequency

`func (o *AppliedCouponObjectExtended) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *AppliedCouponObjectExtended) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *AppliedCouponObjectExtended) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.


### GetFrequencyDuration

`func (o *AppliedCouponObjectExtended) GetFrequencyDuration() int32`

GetFrequencyDuration returns the FrequencyDuration field if non-nil, zero value otherwise.

### GetFrequencyDurationOk

`func (o *AppliedCouponObjectExtended) GetFrequencyDurationOk() (*int32, bool)`

GetFrequencyDurationOk returns a tuple with the FrequencyDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencyDuration

`func (o *AppliedCouponObjectExtended) SetFrequencyDuration(v int32)`

SetFrequencyDuration sets FrequencyDuration field to given value.

### HasFrequencyDuration

`func (o *AppliedCouponObjectExtended) HasFrequencyDuration() bool`

HasFrequencyDuration returns a boolean if a field has been set.

### SetFrequencyDurationNil

`func (o *AppliedCouponObjectExtended) SetFrequencyDurationNil(b bool)`

 SetFrequencyDurationNil sets the value for FrequencyDuration to be an explicit nil

### UnsetFrequencyDuration
`func (o *AppliedCouponObjectExtended) UnsetFrequencyDuration()`

UnsetFrequencyDuration ensures that no value is present for FrequencyDuration, not even an explicit nil
### GetFrequencyDurationRemaining

`func (o *AppliedCouponObjectExtended) GetFrequencyDurationRemaining() int32`

GetFrequencyDurationRemaining returns the FrequencyDurationRemaining field if non-nil, zero value otherwise.

### GetFrequencyDurationRemainingOk

`func (o *AppliedCouponObjectExtended) GetFrequencyDurationRemainingOk() (*int32, bool)`

GetFrequencyDurationRemainingOk returns a tuple with the FrequencyDurationRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencyDurationRemaining

`func (o *AppliedCouponObjectExtended) SetFrequencyDurationRemaining(v int32)`

SetFrequencyDurationRemaining sets FrequencyDurationRemaining field to given value.

### HasFrequencyDurationRemaining

`func (o *AppliedCouponObjectExtended) HasFrequencyDurationRemaining() bool`

HasFrequencyDurationRemaining returns a boolean if a field has been set.

### SetFrequencyDurationRemainingNil

`func (o *AppliedCouponObjectExtended) SetFrequencyDurationRemainingNil(b bool)`

 SetFrequencyDurationRemainingNil sets the value for FrequencyDurationRemaining to be an explicit nil

### UnsetFrequencyDurationRemaining
`func (o *AppliedCouponObjectExtended) UnsetFrequencyDurationRemaining()`

UnsetFrequencyDurationRemaining ensures that no value is present for FrequencyDurationRemaining, not even an explicit nil
### GetExpirationAt

`func (o *AppliedCouponObjectExtended) GetExpirationAt() time.Time`

GetExpirationAt returns the ExpirationAt field if non-nil, zero value otherwise.

### GetExpirationAtOk

`func (o *AppliedCouponObjectExtended) GetExpirationAtOk() (*time.Time, bool)`

GetExpirationAtOk returns a tuple with the ExpirationAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationAt

`func (o *AppliedCouponObjectExtended) SetExpirationAt(v time.Time)`

SetExpirationAt sets ExpirationAt field to given value.

### HasExpirationAt

`func (o *AppliedCouponObjectExtended) HasExpirationAt() bool`

HasExpirationAt returns a boolean if a field has been set.

### SetExpirationAtNil

`func (o *AppliedCouponObjectExtended) SetExpirationAtNil(b bool)`

 SetExpirationAtNil sets the value for ExpirationAt to be an explicit nil

### UnsetExpirationAt
`func (o *AppliedCouponObjectExtended) UnsetExpirationAt()`

UnsetExpirationAt ensures that no value is present for ExpirationAt, not even an explicit nil
### GetCreatedAt

`func (o *AppliedCouponObjectExtended) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AppliedCouponObjectExtended) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AppliedCouponObjectExtended) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetTerminatedAt

`func (o *AppliedCouponObjectExtended) GetTerminatedAt() time.Time`

GetTerminatedAt returns the TerminatedAt field if non-nil, zero value otherwise.

### GetTerminatedAtOk

`func (o *AppliedCouponObjectExtended) GetTerminatedAtOk() (*time.Time, bool)`

GetTerminatedAtOk returns a tuple with the TerminatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminatedAt

`func (o *AppliedCouponObjectExtended) SetTerminatedAt(v time.Time)`

SetTerminatedAt sets TerminatedAt field to given value.

### HasTerminatedAt

`func (o *AppliedCouponObjectExtended) HasTerminatedAt() bool`

HasTerminatedAt returns a boolean if a field has been set.

### SetTerminatedAtNil

`func (o *AppliedCouponObjectExtended) SetTerminatedAtNil(b bool)`

 SetTerminatedAtNil sets the value for TerminatedAt to be an explicit nil

### UnsetTerminatedAt
`func (o *AppliedCouponObjectExtended) UnsetTerminatedAt()`

UnsetTerminatedAt ensures that no value is present for TerminatedAt, not even an explicit nil
### GetCredits

`func (o *AppliedCouponObjectExtended) GetCredits() []CreditObject`

GetCredits returns the Credits field if non-nil, zero value otherwise.

### GetCreditsOk

`func (o *AppliedCouponObjectExtended) GetCreditsOk() (*[]CreditObject, bool)`

GetCreditsOk returns a tuple with the Credits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredits

`func (o *AppliedCouponObjectExtended) SetCredits(v []CreditObject)`

SetCredits sets Credits field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


