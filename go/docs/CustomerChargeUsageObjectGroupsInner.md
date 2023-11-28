# CustomerChargeUsageObjectGroupsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LagoId** | Pointer to **string** | Unique identifier assigned to the group within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the group record within the Lago system. | [optional] 
**Key** | Pointer to **NullableString** | The group key, only returned for groups with two dimensions. | [optional] 
**Value** | Pointer to **string** | The group value. | [optional] 
**Units** | Pointer to **string** | The number of units consumed for a specific group related to a charge item. | [optional] 
**EventsCount** | Pointer to **int32** | The quantity of usage events that have been recorded for a particular charge during the specified time period. These events may also be referred to as the number of transactions in some contexts. | [optional] 
**AmountCents** | Pointer to **int32** | The amount in cents, tax excluded, consumed for a specific group related to a charge item. | [optional] 

## Methods

### NewCustomerChargeUsageObjectGroupsInner

`func NewCustomerChargeUsageObjectGroupsInner() *CustomerChargeUsageObjectGroupsInner`

NewCustomerChargeUsageObjectGroupsInner instantiates a new CustomerChargeUsageObjectGroupsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerChargeUsageObjectGroupsInnerWithDefaults

`func NewCustomerChargeUsageObjectGroupsInnerWithDefaults() *CustomerChargeUsageObjectGroupsInner`

NewCustomerChargeUsageObjectGroupsInnerWithDefaults instantiates a new CustomerChargeUsageObjectGroupsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLagoId

`func (o *CustomerChargeUsageObjectGroupsInner) GetLagoId() string`

GetLagoId returns the LagoId field if non-nil, zero value otherwise.

### GetLagoIdOk

`func (o *CustomerChargeUsageObjectGroupsInner) GetLagoIdOk() (*string, bool)`

GetLagoIdOk returns a tuple with the LagoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagoId

`func (o *CustomerChargeUsageObjectGroupsInner) SetLagoId(v string)`

SetLagoId sets LagoId field to given value.

### HasLagoId

`func (o *CustomerChargeUsageObjectGroupsInner) HasLagoId() bool`

HasLagoId returns a boolean if a field has been set.

### GetKey

`func (o *CustomerChargeUsageObjectGroupsInner) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CustomerChargeUsageObjectGroupsInner) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CustomerChargeUsageObjectGroupsInner) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CustomerChargeUsageObjectGroupsInner) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *CustomerChargeUsageObjectGroupsInner) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *CustomerChargeUsageObjectGroupsInner) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetValue

`func (o *CustomerChargeUsageObjectGroupsInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CustomerChargeUsageObjectGroupsInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CustomerChargeUsageObjectGroupsInner) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *CustomerChargeUsageObjectGroupsInner) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetUnits

`func (o *CustomerChargeUsageObjectGroupsInner) GetUnits() string`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *CustomerChargeUsageObjectGroupsInner) GetUnitsOk() (*string, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *CustomerChargeUsageObjectGroupsInner) SetUnits(v string)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *CustomerChargeUsageObjectGroupsInner) HasUnits() bool`

HasUnits returns a boolean if a field has been set.

### GetEventsCount

`func (o *CustomerChargeUsageObjectGroupsInner) GetEventsCount() int32`

GetEventsCount returns the EventsCount field if non-nil, zero value otherwise.

### GetEventsCountOk

`func (o *CustomerChargeUsageObjectGroupsInner) GetEventsCountOk() (*int32, bool)`

GetEventsCountOk returns a tuple with the EventsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsCount

`func (o *CustomerChargeUsageObjectGroupsInner) SetEventsCount(v int32)`

SetEventsCount sets EventsCount field to given value.

### HasEventsCount

`func (o *CustomerChargeUsageObjectGroupsInner) HasEventsCount() bool`

HasEventsCount returns a boolean if a field has been set.

### GetAmountCents

`func (o *CustomerChargeUsageObjectGroupsInner) GetAmountCents() int32`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *CustomerChargeUsageObjectGroupsInner) GetAmountCentsOk() (*int32, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *CustomerChargeUsageObjectGroupsInner) SetAmountCents(v int32)`

SetAmountCents sets AmountCents field to given value.

### HasAmountCents

`func (o *CustomerChargeUsageObjectGroupsInner) HasAmountCents() bool`

HasAmountCents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


