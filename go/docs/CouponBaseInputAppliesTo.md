# CouponBaseInputAppliesTo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlanCodes** | Pointer to **[]string** | An array of plan codes to which the coupon is applicable. By specifying the plan codes in this field, you can restrict the coupon&#39;s usage to specific plans only. | [optional] 
**BillableMetricCodes** | Pointer to **[]string** | An array of billable metric codes to which the coupon is applicable. By specifying the billable metric codes in this field, you can restrict the coupon&#39;s usage to specific metrics only. | [optional] 

## Methods

### NewCouponBaseInputAppliesTo

`func NewCouponBaseInputAppliesTo() *CouponBaseInputAppliesTo`

NewCouponBaseInputAppliesTo instantiates a new CouponBaseInputAppliesTo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCouponBaseInputAppliesToWithDefaults

`func NewCouponBaseInputAppliesToWithDefaults() *CouponBaseInputAppliesTo`

NewCouponBaseInputAppliesToWithDefaults instantiates a new CouponBaseInputAppliesTo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlanCodes

`func (o *CouponBaseInputAppliesTo) GetPlanCodes() []string`

GetPlanCodes returns the PlanCodes field if non-nil, zero value otherwise.

### GetPlanCodesOk

`func (o *CouponBaseInputAppliesTo) GetPlanCodesOk() (*[]string, bool)`

GetPlanCodesOk returns a tuple with the PlanCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanCodes

`func (o *CouponBaseInputAppliesTo) SetPlanCodes(v []string)`

SetPlanCodes sets PlanCodes field to given value.

### HasPlanCodes

`func (o *CouponBaseInputAppliesTo) HasPlanCodes() bool`

HasPlanCodes returns a boolean if a field has been set.

### SetPlanCodesNil

`func (o *CouponBaseInputAppliesTo) SetPlanCodesNil(b bool)`

 SetPlanCodesNil sets the value for PlanCodes to be an explicit nil

### UnsetPlanCodes
`func (o *CouponBaseInputAppliesTo) UnsetPlanCodes()`

UnsetPlanCodes ensures that no value is present for PlanCodes, not even an explicit nil
### GetBillableMetricCodes

`func (o *CouponBaseInputAppliesTo) GetBillableMetricCodes() []string`

GetBillableMetricCodes returns the BillableMetricCodes field if non-nil, zero value otherwise.

### GetBillableMetricCodesOk

`func (o *CouponBaseInputAppliesTo) GetBillableMetricCodesOk() (*[]string, bool)`

GetBillableMetricCodesOk returns a tuple with the BillableMetricCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillableMetricCodes

`func (o *CouponBaseInputAppliesTo) SetBillableMetricCodes(v []string)`

SetBillableMetricCodes sets BillableMetricCodes field to given value.

### HasBillableMetricCodes

`func (o *CouponBaseInputAppliesTo) HasBillableMetricCodes() bool`

HasBillableMetricCodes returns a boolean if a field has been set.

### SetBillableMetricCodesNil

`func (o *CouponBaseInputAppliesTo) SetBillableMetricCodesNil(b bool)`

 SetBillableMetricCodesNil sets the value for BillableMetricCodes to be an explicit nil

### UnsetBillableMetricCodes
`func (o *CouponBaseInputAppliesTo) UnsetBillableMetricCodes()`

UnsetBillableMetricCodes ensures that no value is present for BillableMetricCodes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


