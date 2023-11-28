# PlansPaginated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Plans** | [**[]PlanObject**](PlanObject.md) |  | 
**Meta** | [**PaginationMeta**](PaginationMeta.md) |  | 

## Methods

### NewPlansPaginated

`func NewPlansPaginated(plans []PlanObject, meta PaginationMeta, ) *PlansPaginated`

NewPlansPaginated instantiates a new PlansPaginated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlansPaginatedWithDefaults

`func NewPlansPaginatedWithDefaults() *PlansPaginated`

NewPlansPaginatedWithDefaults instantiates a new PlansPaginated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlans

`func (o *PlansPaginated) GetPlans() []PlanObject`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *PlansPaginated) GetPlansOk() (*[]PlanObject, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *PlansPaginated) SetPlans(v []PlanObject)`

SetPlans sets Plans field to given value.


### GetMeta

`func (o *PlansPaginated) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *PlansPaginated) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *PlansPaginated) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


