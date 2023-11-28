# BillableMetricGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | Name of the event property used to group values. | 
**Values** | [**[]BillableMetricGroupValuesInner**](BillableMetricGroupValuesInner.md) | Array of strings or objects representing all possible values. | 

## Methods

### NewBillableMetricGroup

`func NewBillableMetricGroup(key string, values []BillableMetricGroupValuesInner, ) *BillableMetricGroup`

NewBillableMetricGroup instantiates a new BillableMetricGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillableMetricGroupWithDefaults

`func NewBillableMetricGroupWithDefaults() *BillableMetricGroup`

NewBillableMetricGroupWithDefaults instantiates a new BillableMetricGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *BillableMetricGroup) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *BillableMetricGroup) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *BillableMetricGroup) SetKey(v string)`

SetKey sets Key field to given value.


### GetValues

`func (o *BillableMetricGroup) GetValues() []BillableMetricGroupValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *BillableMetricGroup) GetValuesOk() (*[]BillableMetricGroupValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *BillableMetricGroup) SetValues(v []BillableMetricGroupValuesInner)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


