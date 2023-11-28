# ChargePropertiesVolumeRangesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromValue** | **int32** | Specifies the lower value of a tier for a &#x60;volume&#x60; charge model. It must be either 0 or the previous range&#39;s &#x60;to_value + 1&#x60; to maintain the proper sequence of values. | 
**ToValue** | **NullableInt32** | Specifies the highest value of a tier for a &#x60;volume&#x60; charge model. - This value must be higher than the &#x60;from_value&#x60; of the same tier. - This value must be &#x60;null&#x60; for the last tier. | 
**FlatAmount** | **string** | The unit price, excluding tax, for a specific tier of a &#x60;volume&#x60; charge model. It is expressed as a decimal value. | 
**PerUnitAmount** | **string** | The flat amount for a whole tier, excluding tax, for a &#x60;volume&#x60; charge model. It is expressed as a decimal value. | 

## Methods

### NewChargePropertiesVolumeRangesInner

`func NewChargePropertiesVolumeRangesInner(fromValue int32, toValue NullableInt32, flatAmount string, perUnitAmount string, ) *ChargePropertiesVolumeRangesInner`

NewChargePropertiesVolumeRangesInner instantiates a new ChargePropertiesVolumeRangesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChargePropertiesVolumeRangesInnerWithDefaults

`func NewChargePropertiesVolumeRangesInnerWithDefaults() *ChargePropertiesVolumeRangesInner`

NewChargePropertiesVolumeRangesInnerWithDefaults instantiates a new ChargePropertiesVolumeRangesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromValue

`func (o *ChargePropertiesVolumeRangesInner) GetFromValue() int32`

GetFromValue returns the FromValue field if non-nil, zero value otherwise.

### GetFromValueOk

`func (o *ChargePropertiesVolumeRangesInner) GetFromValueOk() (*int32, bool)`

GetFromValueOk returns a tuple with the FromValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromValue

`func (o *ChargePropertiesVolumeRangesInner) SetFromValue(v int32)`

SetFromValue sets FromValue field to given value.


### GetToValue

`func (o *ChargePropertiesVolumeRangesInner) GetToValue() int32`

GetToValue returns the ToValue field if non-nil, zero value otherwise.

### GetToValueOk

`func (o *ChargePropertiesVolumeRangesInner) GetToValueOk() (*int32, bool)`

GetToValueOk returns a tuple with the ToValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToValue

`func (o *ChargePropertiesVolumeRangesInner) SetToValue(v int32)`

SetToValue sets ToValue field to given value.


### SetToValueNil

`func (o *ChargePropertiesVolumeRangesInner) SetToValueNil(b bool)`

 SetToValueNil sets the value for ToValue to be an explicit nil

### UnsetToValue
`func (o *ChargePropertiesVolumeRangesInner) UnsetToValue()`

UnsetToValue ensures that no value is present for ToValue, not even an explicit nil
### GetFlatAmount

`func (o *ChargePropertiesVolumeRangesInner) GetFlatAmount() string`

GetFlatAmount returns the FlatAmount field if non-nil, zero value otherwise.

### GetFlatAmountOk

`func (o *ChargePropertiesVolumeRangesInner) GetFlatAmountOk() (*string, bool)`

GetFlatAmountOk returns a tuple with the FlatAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlatAmount

`func (o *ChargePropertiesVolumeRangesInner) SetFlatAmount(v string)`

SetFlatAmount sets FlatAmount field to given value.


### GetPerUnitAmount

`func (o *ChargePropertiesVolumeRangesInner) GetPerUnitAmount() string`

GetPerUnitAmount returns the PerUnitAmount field if non-nil, zero value otherwise.

### GetPerUnitAmountOk

`func (o *ChargePropertiesVolumeRangesInner) GetPerUnitAmountOk() (*string, bool)`

GetPerUnitAmountOk returns a tuple with the PerUnitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerUnitAmount

`func (o *ChargePropertiesVolumeRangesInner) SetPerUnitAmount(v string)`

SetPerUnitAmount sets PerUnitAmount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


