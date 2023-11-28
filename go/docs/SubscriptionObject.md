# SubscriptionObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LagoId** | **string** | Unique identifier assigned to the subscription within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the subscription’s record within the Lago system | 
**ExternalId** | **string** | The subscription external unique identifier (provided by your own application). | 
**LagoCustomerId** | **string** | Unique identifier assigned to the customer within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the customer&#39;s record within the Lago system | 
**ExternalCustomerId** | **string** | The customer external unique identifier (provided by your own application). | 
**BillingTime** | **string** | The billing time for the subscription, which can be set as either &#x60;anniversary&#x60; or &#x60;calendar&#x60;. If not explicitly provided, it will default to &#x60;calendar&#x60;. The billing time determines the timing of recurring billing cycles for the subscription. By specifying &#x60;anniversary&#x60;, the billing cycle will be based on the specific date the subscription started (billed fully), while &#x60;calendar&#x60; sets the billing cycle at the first day of the week/month/year (billed with proration). | 
**Name** | Pointer to **NullableString** | The display name of the subscription on an invoice. This field allows for customization of the subscription&#39;s name for billing purposes, especially useful when a single customer has multiple subscriptions using the same plan. | [optional] 
**PlanCode** | **string** | The unique code representing the plan to be attached to the customer. This code must correspond to the &#x60;code&#x60; property of one of the active plans. | 
**Status** | **string** | The status of the subscription, which can have the following values: - &#x60;pending&#x60;: a previous subscription has been downgraded, and the current one is awaiting automatic activation at the end of the billing period. - &#x60;active&#x60;: the subscription is currently active and applied to the customer. - &#x60;terminated&#x60;: the subscription is no longer active. - &#x60;canceled&#x60;: the subscription has been stopped before its activation. This can occur when two consecutive downgrades have been applied to a customer or when a subscription with a pending status is terminated. | 
**CreatedAt** | **time.Time** | The creation date of the subscription, represented in ISO 8601 datetime format and expressed in Coordinated Universal Time (UTC). This date provides a timestamp indicating when the subscription was initially created. | 
**CanceledAt** | Pointer to **NullableTime** | The cancellation date of the subscription. This field is not null when the subscription is &#x60;canceled&#x60;. This date should be provided in ISO 8601 datetime format and expressed in Coordinated Universal Time (UTC). | [optional] 
**StartedAt** | Pointer to **NullableTime** | The effective start date of the subscription. This field can be null if the subscription is &#x60;pending&#x60; or &#x60;canceled&#x60;. This date should be provided in ISO 8601 datetime format and expressed in Coordinated Universal Time (UTC). | [optional] 
**EndingAt** | Pointer to **time.Time** | The effective end date of the subscription. If this field is set to null, the subscription will automatically renew. This date should be provided in ISO 8601 datetime format, and use Coordinated Universal Time (UTC). | [optional] 
**SubscriptionAt** | **time.Time** | The anniversary date and time of the initial subscription. This date serves as the basis for billing subscriptions with &#x60;anniversary&#x60; billing time. The &#x60;anniversary_date&#x60; should be provided in ISO 8601 datetime format and expressed in Coordinated Universal Time (UTC). | 
**TerminatedAt** | Pointer to **NullableTime** | The termination date of the subscription. This field is not null when the subscription is &#x60;terminated&#x60;. This date should be provided in ISO 8601 datetime format and expressed in Coordinated Universal Time (UTC) | [optional] 
**PreviousPlanCode** | Pointer to **NullableString** | The code identifying the previous plan associated with this subscription. | [optional] 
**NextPlanCode** | Pointer to **NullableString** | The code identifying the next plan in the case of a downgrade. | [optional] 
**DowngradePlanDate** | Pointer to **NullableTime** | The date when the plan will be downgraded, represented in ISO 8601 date format. | [optional] 

## Methods

### NewSubscriptionObject

`func NewSubscriptionObject(lagoId string, externalId string, lagoCustomerId string, externalCustomerId string, billingTime string, planCode string, status string, createdAt time.Time, subscriptionAt time.Time, ) *SubscriptionObject`

NewSubscriptionObject instantiates a new SubscriptionObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionObjectWithDefaults

`func NewSubscriptionObjectWithDefaults() *SubscriptionObject`

NewSubscriptionObjectWithDefaults instantiates a new SubscriptionObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLagoId

`func (o *SubscriptionObject) GetLagoId() string`

GetLagoId returns the LagoId field if non-nil, zero value otherwise.

### GetLagoIdOk

`func (o *SubscriptionObject) GetLagoIdOk() (*string, bool)`

GetLagoIdOk returns a tuple with the LagoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagoId

`func (o *SubscriptionObject) SetLagoId(v string)`

SetLagoId sets LagoId field to given value.


### GetExternalId

`func (o *SubscriptionObject) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *SubscriptionObject) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *SubscriptionObject) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetLagoCustomerId

`func (o *SubscriptionObject) GetLagoCustomerId() string`

GetLagoCustomerId returns the LagoCustomerId field if non-nil, zero value otherwise.

### GetLagoCustomerIdOk

`func (o *SubscriptionObject) GetLagoCustomerIdOk() (*string, bool)`

GetLagoCustomerIdOk returns a tuple with the LagoCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagoCustomerId

`func (o *SubscriptionObject) SetLagoCustomerId(v string)`

SetLagoCustomerId sets LagoCustomerId field to given value.


### GetExternalCustomerId

`func (o *SubscriptionObject) GetExternalCustomerId() string`

GetExternalCustomerId returns the ExternalCustomerId field if non-nil, zero value otherwise.

### GetExternalCustomerIdOk

`func (o *SubscriptionObject) GetExternalCustomerIdOk() (*string, bool)`

GetExternalCustomerIdOk returns a tuple with the ExternalCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalCustomerId

`func (o *SubscriptionObject) SetExternalCustomerId(v string)`

SetExternalCustomerId sets ExternalCustomerId field to given value.


### GetBillingTime

`func (o *SubscriptionObject) GetBillingTime() string`

GetBillingTime returns the BillingTime field if non-nil, zero value otherwise.

### GetBillingTimeOk

`func (o *SubscriptionObject) GetBillingTimeOk() (*string, bool)`

GetBillingTimeOk returns a tuple with the BillingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingTime

`func (o *SubscriptionObject) SetBillingTime(v string)`

SetBillingTime sets BillingTime field to given value.


### GetName

`func (o *SubscriptionObject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SubscriptionObject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SubscriptionObject) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SubscriptionObject) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *SubscriptionObject) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SubscriptionObject) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPlanCode

`func (o *SubscriptionObject) GetPlanCode() string`

GetPlanCode returns the PlanCode field if non-nil, zero value otherwise.

### GetPlanCodeOk

`func (o *SubscriptionObject) GetPlanCodeOk() (*string, bool)`

GetPlanCodeOk returns a tuple with the PlanCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanCode

`func (o *SubscriptionObject) SetPlanCode(v string)`

SetPlanCode sets PlanCode field to given value.


### GetStatus

`func (o *SubscriptionObject) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SubscriptionObject) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SubscriptionObject) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *SubscriptionObject) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SubscriptionObject) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SubscriptionObject) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCanceledAt

`func (o *SubscriptionObject) GetCanceledAt() time.Time`

GetCanceledAt returns the CanceledAt field if non-nil, zero value otherwise.

### GetCanceledAtOk

`func (o *SubscriptionObject) GetCanceledAtOk() (*time.Time, bool)`

GetCanceledAtOk returns a tuple with the CanceledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanceledAt

`func (o *SubscriptionObject) SetCanceledAt(v time.Time)`

SetCanceledAt sets CanceledAt field to given value.

### HasCanceledAt

`func (o *SubscriptionObject) HasCanceledAt() bool`

HasCanceledAt returns a boolean if a field has been set.

### SetCanceledAtNil

`func (o *SubscriptionObject) SetCanceledAtNil(b bool)`

 SetCanceledAtNil sets the value for CanceledAt to be an explicit nil

### UnsetCanceledAt
`func (o *SubscriptionObject) UnsetCanceledAt()`

UnsetCanceledAt ensures that no value is present for CanceledAt, not even an explicit nil
### GetStartedAt

`func (o *SubscriptionObject) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *SubscriptionObject) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *SubscriptionObject) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *SubscriptionObject) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *SubscriptionObject) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *SubscriptionObject) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetEndingAt

`func (o *SubscriptionObject) GetEndingAt() time.Time`

GetEndingAt returns the EndingAt field if non-nil, zero value otherwise.

### GetEndingAtOk

`func (o *SubscriptionObject) GetEndingAtOk() (*time.Time, bool)`

GetEndingAtOk returns a tuple with the EndingAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndingAt

`func (o *SubscriptionObject) SetEndingAt(v time.Time)`

SetEndingAt sets EndingAt field to given value.

### HasEndingAt

`func (o *SubscriptionObject) HasEndingAt() bool`

HasEndingAt returns a boolean if a field has been set.

### GetSubscriptionAt

`func (o *SubscriptionObject) GetSubscriptionAt() time.Time`

GetSubscriptionAt returns the SubscriptionAt field if non-nil, zero value otherwise.

### GetSubscriptionAtOk

`func (o *SubscriptionObject) GetSubscriptionAtOk() (*time.Time, bool)`

GetSubscriptionAtOk returns a tuple with the SubscriptionAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionAt

`func (o *SubscriptionObject) SetSubscriptionAt(v time.Time)`

SetSubscriptionAt sets SubscriptionAt field to given value.


### GetTerminatedAt

`func (o *SubscriptionObject) GetTerminatedAt() time.Time`

GetTerminatedAt returns the TerminatedAt field if non-nil, zero value otherwise.

### GetTerminatedAtOk

`func (o *SubscriptionObject) GetTerminatedAtOk() (*time.Time, bool)`

GetTerminatedAtOk returns a tuple with the TerminatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminatedAt

`func (o *SubscriptionObject) SetTerminatedAt(v time.Time)`

SetTerminatedAt sets TerminatedAt field to given value.

### HasTerminatedAt

`func (o *SubscriptionObject) HasTerminatedAt() bool`

HasTerminatedAt returns a boolean if a field has been set.

### SetTerminatedAtNil

`func (o *SubscriptionObject) SetTerminatedAtNil(b bool)`

 SetTerminatedAtNil sets the value for TerminatedAt to be an explicit nil

### UnsetTerminatedAt
`func (o *SubscriptionObject) UnsetTerminatedAt()`

UnsetTerminatedAt ensures that no value is present for TerminatedAt, not even an explicit nil
### GetPreviousPlanCode

`func (o *SubscriptionObject) GetPreviousPlanCode() string`

GetPreviousPlanCode returns the PreviousPlanCode field if non-nil, zero value otherwise.

### GetPreviousPlanCodeOk

`func (o *SubscriptionObject) GetPreviousPlanCodeOk() (*string, bool)`

GetPreviousPlanCodeOk returns a tuple with the PreviousPlanCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousPlanCode

`func (o *SubscriptionObject) SetPreviousPlanCode(v string)`

SetPreviousPlanCode sets PreviousPlanCode field to given value.

### HasPreviousPlanCode

`func (o *SubscriptionObject) HasPreviousPlanCode() bool`

HasPreviousPlanCode returns a boolean if a field has been set.

### SetPreviousPlanCodeNil

`func (o *SubscriptionObject) SetPreviousPlanCodeNil(b bool)`

 SetPreviousPlanCodeNil sets the value for PreviousPlanCode to be an explicit nil

### UnsetPreviousPlanCode
`func (o *SubscriptionObject) UnsetPreviousPlanCode()`

UnsetPreviousPlanCode ensures that no value is present for PreviousPlanCode, not even an explicit nil
### GetNextPlanCode

`func (o *SubscriptionObject) GetNextPlanCode() string`

GetNextPlanCode returns the NextPlanCode field if non-nil, zero value otherwise.

### GetNextPlanCodeOk

`func (o *SubscriptionObject) GetNextPlanCodeOk() (*string, bool)`

GetNextPlanCodeOk returns a tuple with the NextPlanCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPlanCode

`func (o *SubscriptionObject) SetNextPlanCode(v string)`

SetNextPlanCode sets NextPlanCode field to given value.

### HasNextPlanCode

`func (o *SubscriptionObject) HasNextPlanCode() bool`

HasNextPlanCode returns a boolean if a field has been set.

### SetNextPlanCodeNil

`func (o *SubscriptionObject) SetNextPlanCodeNil(b bool)`

 SetNextPlanCodeNil sets the value for NextPlanCode to be an explicit nil

### UnsetNextPlanCode
`func (o *SubscriptionObject) UnsetNextPlanCode()`

UnsetNextPlanCode ensures that no value is present for NextPlanCode, not even an explicit nil
### GetDowngradePlanDate

`func (o *SubscriptionObject) GetDowngradePlanDate() time.Time`

GetDowngradePlanDate returns the DowngradePlanDate field if non-nil, zero value otherwise.

### GetDowngradePlanDateOk

`func (o *SubscriptionObject) GetDowngradePlanDateOk() (*time.Time, bool)`

GetDowngradePlanDateOk returns a tuple with the DowngradePlanDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDowngradePlanDate

`func (o *SubscriptionObject) SetDowngradePlanDate(v time.Time)`

SetDowngradePlanDate sets DowngradePlanDate field to given value.

### HasDowngradePlanDate

`func (o *SubscriptionObject) HasDowngradePlanDate() bool`

HasDowngradePlanDate returns a boolean if a field has been set.

### SetDowngradePlanDateNil

`func (o *SubscriptionObject) SetDowngradePlanDateNil(b bool)`

 SetDowngradePlanDateNil sets the value for DowngradePlanDate to be an explicit nil

### UnsetDowngradePlanDate
`func (o *SubscriptionObject) UnsetDowngradePlanDate()`

UnsetDowngradePlanDate ensures that no value is present for DowngradePlanDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


