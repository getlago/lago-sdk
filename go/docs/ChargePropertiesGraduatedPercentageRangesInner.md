# ChargePropertiesGraduatedPercentageRangesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromValue** | **int32** | Specifies the lower value of a tier for a &#x60;graduated_percentage&#x60; charge model. It must be either 0 or the previous range&#39;s &#x60;to_value + 1&#x60; to maintain the proper sequence of values. | 
**ToValue** | **NullableInt32** | Specifies the highest value of a tier for a &#x60;graduated_percentage&#x60; charge model. - This value must be higher than the from_value of the same tier. - This value must be null for the last tier. | 
**Rate** | **string** | The percentage rate that is applied to the amount of each transaction in the tier for a &#x60;graduated_percentage&#x60; charge model. It is expressed as a decimal value. | 
**FlatAmount** | **string** | The flat amount for a whole tier, excluding tax, for a &#x60;graduated_percentage&#x60; charge model. It is expressed as a decimal value. | 

## Methods

### NewChargePropertiesGraduatedPercentageRangesInner

`func NewChargePropertiesGraduatedPercentageRangesInner(fromValue int32, toValue NullableInt32, rate string, flatAmount string, ) *ChargePropertiesGraduatedPercentageRangesInner`

NewChargePropertiesGraduatedPercentageRangesInner instantiates a new ChargePropertiesGraduatedPercentageRangesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChargePropertiesGraduatedPercentageRangesInnerWithDefaults

`func NewChargePropertiesGraduatedPercentageRangesInnerWithDefaults() *ChargePropertiesGraduatedPercentageRangesInner`

NewChargePropertiesGraduatedPercentageRangesInnerWithDefaults instantiates a new ChargePropertiesGraduatedPercentageRangesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromValue

`func (o *ChargePropertiesGraduatedPercentageRangesInner) GetFromValue() int32`

GetFromValue returns the FromValue field if non-nil, zero value otherwise.

### GetFromValueOk

`func (o *ChargePropertiesGraduatedPercentageRangesInner) GetFromValueOk() (*int32, bool)`

GetFromValueOk returns a tuple with the FromValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromValue

`func (o *ChargePropertiesGraduatedPercentageRangesInner) SetFromValue(v int32)`

SetFromValue sets FromValue field to given value.


### GetToValue

`func (o *ChargePropertiesGraduatedPercentageRangesInner) GetToValue() int32`

GetToValue returns the ToValue field if non-nil, zero value otherwise.

### GetToValueOk

`func (o *ChargePropertiesGraduatedPercentageRangesInner) GetToValueOk() (*int32, bool)`

GetToValueOk returns a tuple with the ToValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToValue

`func (o *ChargePropertiesGraduatedPercentageRangesInner) SetToValue(v int32)`

SetToValue sets ToValue field to given value.


### SetToValueNil

`func (o *ChargePropertiesGraduatedPercentageRangesInner) SetToValueNil(b bool)`

 SetToValueNil sets the value for ToValue to be an explicit nil

### UnsetToValue
`func (o *ChargePropertiesGraduatedPercentageRangesInner) UnsetToValue()`

UnsetToValue ensures that no value is present for ToValue, not even an explicit nil
### GetRate

`func (o *ChargePropertiesGraduatedPercentageRangesInner) GetRate() string`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *ChargePropertiesGraduatedPercentageRangesInner) GetRateOk() (*string, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *ChargePropertiesGraduatedPercentageRangesInner) SetRate(v string)`

SetRate sets Rate field to given value.


### GetFlatAmount

`func (o *ChargePropertiesGraduatedPercentageRangesInner) GetFlatAmount() string`

GetFlatAmount returns the FlatAmount field if non-nil, zero value otherwise.

### GetFlatAmountOk

`func (o *ChargePropertiesGraduatedPercentageRangesInner) GetFlatAmountOk() (*string, bool)`

GetFlatAmountOk returns a tuple with the FlatAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlatAmount

`func (o *ChargePropertiesGraduatedPercentageRangesInner) SetFlatAmount(v string)`

SetFlatAmount sets FlatAmount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


