# SubscriptionCreateInputSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalCustomerId** | **string** | The customer external unique identifier (provided by your own application) | 
**PlanCode** | **string** | The unique code representing the plan to be attached to the customer. This code must correspond to the &#x60;code&#x60; property of one of the active plans. | 
**Name** | Pointer to **string** | The display name of the subscription on an invoice. This field allows for customization of the subscription&#39;s name for billing purposes, especially useful when a single customer has multiple subscriptions using the same plan. | [optional] 
**ExternalId** | **string** | The unique external identifier for the subscription. This identifier serves as an idempotency key, ensuring that each subscription is unique. | 
**BillingTime** | Pointer to **string** | The billing time for the subscription, which can be set as either &#x60;anniversary&#x60; or &#x60;calendar&#x60;. If not explicitly provided, it will default to &#x60;calendar&#x60;. The billing time determines the timing of recurring billing cycles for the subscription. By specifying &#x60;anniversary&#x60;, the billing cycle will be based on the specific date the subscription started (billed fully), while &#x60;calendar&#x60; sets the billing cycle at the first day of the week/month/year (billed with proration). | [optional] 
**EndingAt** | Pointer to **time.Time** | The effective end date of the subscription. If this field is set to null, the subscription will automatically renew. This date should be provided in ISO 8601 datetime format, and use Coordinated Universal Time (UTC). | [optional] 
**SubscriptionAt** | Pointer to **time.Time** | The start date for the subscription, allowing for the creation of subscriptions that can begin in the past or future. Please note that it cannot be used to update the start date of a pending subscription or schedule an upgrade/downgrade. The start_date should be provided in ISO 8601 datetime format and expressed in Coordinated Universal Time (UTC). | [optional] 
**PlanOverrides** | Pointer to [**PlanOverridesObject**](PlanOverridesObject.md) |  | [optional] 

## Methods

### NewSubscriptionCreateInputSubscription

`func NewSubscriptionCreateInputSubscription(externalCustomerId string, planCode string, externalId string, ) *SubscriptionCreateInputSubscription`

NewSubscriptionCreateInputSubscription instantiates a new SubscriptionCreateInputSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionCreateInputSubscriptionWithDefaults

`func NewSubscriptionCreateInputSubscriptionWithDefaults() *SubscriptionCreateInputSubscription`

NewSubscriptionCreateInputSubscriptionWithDefaults instantiates a new SubscriptionCreateInputSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalCustomerId

`func (o *SubscriptionCreateInputSubscription) GetExternalCustomerId() string`

GetExternalCustomerId returns the ExternalCustomerId field if non-nil, zero value otherwise.

### GetExternalCustomerIdOk

`func (o *SubscriptionCreateInputSubscription) GetExternalCustomerIdOk() (*string, bool)`

GetExternalCustomerIdOk returns a tuple with the ExternalCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalCustomerId

`func (o *SubscriptionCreateInputSubscription) SetExternalCustomerId(v string)`

SetExternalCustomerId sets ExternalCustomerId field to given value.


### GetPlanCode

`func (o *SubscriptionCreateInputSubscription) GetPlanCode() string`

GetPlanCode returns the PlanCode field if non-nil, zero value otherwise.

### GetPlanCodeOk

`func (o *SubscriptionCreateInputSubscription) GetPlanCodeOk() (*string, bool)`

GetPlanCodeOk returns a tuple with the PlanCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanCode

`func (o *SubscriptionCreateInputSubscription) SetPlanCode(v string)`

SetPlanCode sets PlanCode field to given value.


### GetName

`func (o *SubscriptionCreateInputSubscription) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SubscriptionCreateInputSubscription) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SubscriptionCreateInputSubscription) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SubscriptionCreateInputSubscription) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExternalId

`func (o *SubscriptionCreateInputSubscription) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *SubscriptionCreateInputSubscription) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *SubscriptionCreateInputSubscription) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetBillingTime

`func (o *SubscriptionCreateInputSubscription) GetBillingTime() string`

GetBillingTime returns the BillingTime field if non-nil, zero value otherwise.

### GetBillingTimeOk

`func (o *SubscriptionCreateInputSubscription) GetBillingTimeOk() (*string, bool)`

GetBillingTimeOk returns a tuple with the BillingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingTime

`func (o *SubscriptionCreateInputSubscription) SetBillingTime(v string)`

SetBillingTime sets BillingTime field to given value.

### HasBillingTime

`func (o *SubscriptionCreateInputSubscription) HasBillingTime() bool`

HasBillingTime returns a boolean if a field has been set.

### GetEndingAt

`func (o *SubscriptionCreateInputSubscription) GetEndingAt() time.Time`

GetEndingAt returns the EndingAt field if non-nil, zero value otherwise.

### GetEndingAtOk

`func (o *SubscriptionCreateInputSubscription) GetEndingAtOk() (*time.Time, bool)`

GetEndingAtOk returns a tuple with the EndingAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndingAt

`func (o *SubscriptionCreateInputSubscription) SetEndingAt(v time.Time)`

SetEndingAt sets EndingAt field to given value.

### HasEndingAt

`func (o *SubscriptionCreateInputSubscription) HasEndingAt() bool`

HasEndingAt returns a boolean if a field has been set.

### GetSubscriptionAt

`func (o *SubscriptionCreateInputSubscription) GetSubscriptionAt() time.Time`

GetSubscriptionAt returns the SubscriptionAt field if non-nil, zero value otherwise.

### GetSubscriptionAtOk

`func (o *SubscriptionCreateInputSubscription) GetSubscriptionAtOk() (*time.Time, bool)`

GetSubscriptionAtOk returns a tuple with the SubscriptionAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionAt

`func (o *SubscriptionCreateInputSubscription) SetSubscriptionAt(v time.Time)`

SetSubscriptionAt sets SubscriptionAt field to given value.

### HasSubscriptionAt

`func (o *SubscriptionCreateInputSubscription) HasSubscriptionAt() bool`

HasSubscriptionAt returns a boolean if a field has been set.

### GetPlanOverrides

`func (o *SubscriptionCreateInputSubscription) GetPlanOverrides() PlanOverridesObject`

GetPlanOverrides returns the PlanOverrides field if non-nil, zero value otherwise.

### GetPlanOverridesOk

`func (o *SubscriptionCreateInputSubscription) GetPlanOverridesOk() (*PlanOverridesObject, bool)`

GetPlanOverridesOk returns a tuple with the PlanOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanOverrides

`func (o *SubscriptionCreateInputSubscription) SetPlanOverrides(v PlanOverridesObject)`

SetPlanOverrides sets PlanOverrides field to given value.

### HasPlanOverrides

`func (o *SubscriptionCreateInputSubscription) HasPlanOverrides() bool`

HasPlanOverrides returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


