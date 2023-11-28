# CouponsPaginated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Coupons** | [**[]CouponObject**](CouponObject.md) |  | 
**Meta** | [**PaginationMeta**](PaginationMeta.md) |  | 

## Methods

### NewCouponsPaginated

`func NewCouponsPaginated(coupons []CouponObject, meta PaginationMeta, ) *CouponsPaginated`

NewCouponsPaginated instantiates a new CouponsPaginated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCouponsPaginatedWithDefaults

`func NewCouponsPaginatedWithDefaults() *CouponsPaginated`

NewCouponsPaginatedWithDefaults instantiates a new CouponsPaginated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoupons

`func (o *CouponsPaginated) GetCoupons() []CouponObject`

GetCoupons returns the Coupons field if non-nil, zero value otherwise.

### GetCouponsOk

`func (o *CouponsPaginated) GetCouponsOk() (*[]CouponObject, bool)`

GetCouponsOk returns a tuple with the Coupons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoupons

`func (o *CouponsPaginated) SetCoupons(v []CouponObject)`

SetCoupons sets Coupons field to given value.


### GetMeta

`func (o *CouponsPaginated) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CouponsPaginated) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CouponsPaginated) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


