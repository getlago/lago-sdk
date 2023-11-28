# CustomerChargeUsageObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Units** | **string** | The number of units consumed by the customer for a specific charge item. | 
**EventsCount** | **int32** | The quantity of usage events that have been recorded for a particular charge during the specified time period. These events may also be referred to as the number of transactions in some contexts. | 
**AmountCents** | **int32** | The amount in cents, tax excluded, consumed by the customer for a specific charge item. | 
**AmountCurrency** | [**Currency**](Currency.md) |  | 
**Charge** | [**CustomerChargeUsageObjectCharge**](CustomerChargeUsageObjectCharge.md) |  | 
**BillableMetric** | [**CustomerChargeUsageObjectBillableMetric**](CustomerChargeUsageObjectBillableMetric.md) |  | 
**Groups** | [**[]CustomerChargeUsageObjectGroupsInner**](CustomerChargeUsageObjectGroupsInner.md) | Array of group object, representing multiple dimensions for a charge item. | 

## Methods

### NewCustomerChargeUsageObject

`func NewCustomerChargeUsageObject(units string, eventsCount int32, amountCents int32, amountCurrency Currency, charge CustomerChargeUsageObjectCharge, billableMetric CustomerChargeUsageObjectBillableMetric, groups []CustomerChargeUsageObjectGroupsInner, ) *CustomerChargeUsageObject`

NewCustomerChargeUsageObject instantiates a new CustomerChargeUsageObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerChargeUsageObjectWithDefaults

`func NewCustomerChargeUsageObjectWithDefaults() *CustomerChargeUsageObject`

NewCustomerChargeUsageObjectWithDefaults instantiates a new CustomerChargeUsageObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnits

`func (o *CustomerChargeUsageObject) GetUnits() string`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *CustomerChargeUsageObject) GetUnitsOk() (*string, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *CustomerChargeUsageObject) SetUnits(v string)`

SetUnits sets Units field to given value.


### GetEventsCount

`func (o *CustomerChargeUsageObject) GetEventsCount() int32`

GetEventsCount returns the EventsCount field if non-nil, zero value otherwise.

### GetEventsCountOk

`func (o *CustomerChargeUsageObject) GetEventsCountOk() (*int32, bool)`

GetEventsCountOk returns a tuple with the EventsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsCount

`func (o *CustomerChargeUsageObject) SetEventsCount(v int32)`

SetEventsCount sets EventsCount field to given value.


### GetAmountCents

`func (o *CustomerChargeUsageObject) GetAmountCents() int32`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *CustomerChargeUsageObject) GetAmountCentsOk() (*int32, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *CustomerChargeUsageObject) SetAmountCents(v int32)`

SetAmountCents sets AmountCents field to given value.


### GetAmountCurrency

`func (o *CustomerChargeUsageObject) GetAmountCurrency() Currency`

GetAmountCurrency returns the AmountCurrency field if non-nil, zero value otherwise.

### GetAmountCurrencyOk

`func (o *CustomerChargeUsageObject) GetAmountCurrencyOk() (*Currency, bool)`

GetAmountCurrencyOk returns a tuple with the AmountCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCurrency

`func (o *CustomerChargeUsageObject) SetAmountCurrency(v Currency)`

SetAmountCurrency sets AmountCurrency field to given value.


### GetCharge

`func (o *CustomerChargeUsageObject) GetCharge() CustomerChargeUsageObjectCharge`

GetCharge returns the Charge field if non-nil, zero value otherwise.

### GetChargeOk

`func (o *CustomerChargeUsageObject) GetChargeOk() (*CustomerChargeUsageObjectCharge, bool)`

GetChargeOk returns a tuple with the Charge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharge

`func (o *CustomerChargeUsageObject) SetCharge(v CustomerChargeUsageObjectCharge)`

SetCharge sets Charge field to given value.


### GetBillableMetric

`func (o *CustomerChargeUsageObject) GetBillableMetric() CustomerChargeUsageObjectBillableMetric`

GetBillableMetric returns the BillableMetric field if non-nil, zero value otherwise.

### GetBillableMetricOk

`func (o *CustomerChargeUsageObject) GetBillableMetricOk() (*CustomerChargeUsageObjectBillableMetric, bool)`

GetBillableMetricOk returns a tuple with the BillableMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillableMetric

`func (o *CustomerChargeUsageObject) SetBillableMetric(v CustomerChargeUsageObjectBillableMetric)`

SetBillableMetric sets BillableMetric field to given value.


### GetGroups

`func (o *CustomerChargeUsageObject) GetGroups() []CustomerChargeUsageObjectGroupsInner`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *CustomerChargeUsageObject) GetGroupsOk() (*[]CustomerChargeUsageObjectGroupsInner, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *CustomerChargeUsageObject) SetGroups(v []CustomerChargeUsageObjectGroupsInner)`

SetGroups sets Groups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


