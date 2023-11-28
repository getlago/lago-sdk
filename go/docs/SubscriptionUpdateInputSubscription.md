# SubscriptionUpdateInputSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | The display name of the subscription on an invoice. This field allows for customization of the subscription&#39;s name for billing purposes, especially useful when a single customer has multiple subscriptions using the same plan. | [optional] 
**EndingAt** | Pointer to **time.Time** | The effective end date of the subscription. If this field is set to null, the subscription will automatically renew. This date should be provided in ISO 8601 datetime format, and use Coordinated Universal Time (UTC). | [optional] 
**SubscriptionAt** | Pointer to **time.Time** | The start date and time of the subscription. This field can only be modified for pending subscriptions that have not yet started. This date should be provided in ISO 8601 datetime format and expressed in Coordinated Universal Time (UTC). | [optional] 
**PlanOverrides** | Pointer to [**PlanOverridesObject**](PlanOverridesObject.md) |  | [optional] 

## Methods

### NewSubscriptionUpdateInputSubscription

`func NewSubscriptionUpdateInputSubscription() *SubscriptionUpdateInputSubscription`

NewSubscriptionUpdateInputSubscription instantiates a new SubscriptionUpdateInputSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionUpdateInputSubscriptionWithDefaults

`func NewSubscriptionUpdateInputSubscriptionWithDefaults() *SubscriptionUpdateInputSubscription`

NewSubscriptionUpdateInputSubscriptionWithDefaults instantiates a new SubscriptionUpdateInputSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SubscriptionUpdateInputSubscription) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SubscriptionUpdateInputSubscription) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SubscriptionUpdateInputSubscription) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SubscriptionUpdateInputSubscription) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *SubscriptionUpdateInputSubscription) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SubscriptionUpdateInputSubscription) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetEndingAt

`func (o *SubscriptionUpdateInputSubscription) GetEndingAt() time.Time`

GetEndingAt returns the EndingAt field if non-nil, zero value otherwise.

### GetEndingAtOk

`func (o *SubscriptionUpdateInputSubscription) GetEndingAtOk() (*time.Time, bool)`

GetEndingAtOk returns a tuple with the EndingAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndingAt

`func (o *SubscriptionUpdateInputSubscription) SetEndingAt(v time.Time)`

SetEndingAt sets EndingAt field to given value.

### HasEndingAt

`func (o *SubscriptionUpdateInputSubscription) HasEndingAt() bool`

HasEndingAt returns a boolean if a field has been set.

### GetSubscriptionAt

`func (o *SubscriptionUpdateInputSubscription) GetSubscriptionAt() time.Time`

GetSubscriptionAt returns the SubscriptionAt field if non-nil, zero value otherwise.

### GetSubscriptionAtOk

`func (o *SubscriptionUpdateInputSubscription) GetSubscriptionAtOk() (*time.Time, bool)`

GetSubscriptionAtOk returns a tuple with the SubscriptionAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionAt

`func (o *SubscriptionUpdateInputSubscription) SetSubscriptionAt(v time.Time)`

SetSubscriptionAt sets SubscriptionAt field to given value.

### HasSubscriptionAt

`func (o *SubscriptionUpdateInputSubscription) HasSubscriptionAt() bool`

HasSubscriptionAt returns a boolean if a field has been set.

### GetPlanOverrides

`func (o *SubscriptionUpdateInputSubscription) GetPlanOverrides() PlanOverridesObject`

GetPlanOverrides returns the PlanOverrides field if non-nil, zero value otherwise.

### GetPlanOverridesOk

`func (o *SubscriptionUpdateInputSubscription) GetPlanOverridesOk() (*PlanOverridesObject, bool)`

GetPlanOverridesOk returns a tuple with the PlanOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanOverrides

`func (o *SubscriptionUpdateInputSubscription) SetPlanOverrides(v PlanOverridesObject)`

SetPlanOverrides sets PlanOverrides field to given value.

### HasPlanOverrides

`func (o *SubscriptionUpdateInputSubscription) HasPlanOverrides() bool`

HasPlanOverrides returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


