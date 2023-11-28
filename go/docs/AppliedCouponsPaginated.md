# AppliedCouponsPaginated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppliedCoupons** | [**[]AppliedCouponObjectExtended**](AppliedCouponObjectExtended.md) |  | 
**Meta** | [**PaginationMeta**](PaginationMeta.md) |  | 

## Methods

### NewAppliedCouponsPaginated

`func NewAppliedCouponsPaginated(appliedCoupons []AppliedCouponObjectExtended, meta PaginationMeta, ) *AppliedCouponsPaginated`

NewAppliedCouponsPaginated instantiates a new AppliedCouponsPaginated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppliedCouponsPaginatedWithDefaults

`func NewAppliedCouponsPaginatedWithDefaults() *AppliedCouponsPaginated`

NewAppliedCouponsPaginatedWithDefaults instantiates a new AppliedCouponsPaginated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppliedCoupons

`func (o *AppliedCouponsPaginated) GetAppliedCoupons() []AppliedCouponObjectExtended`

GetAppliedCoupons returns the AppliedCoupons field if non-nil, zero value otherwise.

### GetAppliedCouponsOk

`func (o *AppliedCouponsPaginated) GetAppliedCouponsOk() (*[]AppliedCouponObjectExtended, bool)`

GetAppliedCouponsOk returns a tuple with the AppliedCoupons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedCoupons

`func (o *AppliedCouponsPaginated) SetAppliedCoupons(v []AppliedCouponObjectExtended)`

SetAppliedCoupons sets AppliedCoupons field to given value.


### GetMeta

`func (o *AppliedCouponsPaginated) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AppliedCouponsPaginated) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AppliedCouponsPaginated) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


