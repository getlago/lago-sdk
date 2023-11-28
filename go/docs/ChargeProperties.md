# ChargeProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GraduatedRanges** | Pointer to [**[]ChargePropertiesGraduatedRangesInner**](ChargePropertiesGraduatedRangesInner.md) | Graduated ranges, sorted from bottom to top tiers, used for a &#x60;graduated&#x60; charge model. | [optional] 
**GraduatedPercentageRanges** | Pointer to [**[]ChargePropertiesGraduatedPercentageRangesInner**](ChargePropertiesGraduatedPercentageRangesInner.md) | Graduated percentage ranges, sorted from bottom to top tiers, used for a &#x60;graduated_percentage&#x60; charge model. | [optional] 
**Amount** | Pointer to **string** | - The unit price, excluding tax, for a &#x60;standard&#x60; charge model. It is expressed as a decimal value. - The amount, excluding tax, for a complete set of units in a &#x60;package&#x60; charge model. It is expressed as a decimal value. | [optional] 
**FreeUnits** | Pointer to **int32** | The quantity of units that are provided free of charge for each billing period in a &#x60;package&#x60; charge model. This field specifies the number of units that customers can use without incurring any additional cost during each billing cycle. | [optional] 
**PackageSize** | Pointer to **int32** | The quantity of units included in each pack or set for a &#x60;package&#x60; charge model. It indicates the number of units that are bundled together as a single package or set within the pricing structure. | [optional] 
**Rate** | Pointer to **string** | The percentage rate that is applied to the amount of each transaction for a &#x60;percentage&#x60; charge model. It is expressed as a decimal value. | [optional] 
**FixedAmount** | Pointer to **string** | The fixed fee that is applied to each transaction for a &#x60;percentage&#x60; charge model. It is expressed as a decimal value. | [optional] 
**FreeUnitsPerEvents** | Pointer to **NullableInt32** | The count of transactions that are not impacted by the &#x60;percentage&#x60; rate and fixed fee in a percentage charge model. This field indicates the number of transactions that are exempt from the calculation of charges based on the specified percentage rate and fixed fee. | [optional] 
**FreeUnitsPerTotalAggregation** | Pointer to **NullableString** | The transaction amount that is not impacted by the &#x60;percentage&#x60; rate and fixed fee in a percentage charge model. This field indicates the portion of the transaction amount that is exempt from the calculation of charges based on the specified percentage rate and fixed fee. | [optional] 
**PerTransactionMaxAmount** | Pointer to **NullableString** | Specifies the maximum allowable spending for a single transaction. Working as a transaction cap. | [optional] 
**PerTransactionMinAmount** | Pointer to **NullableString** | Specifies the minimum allowable spending for a single transaction. Working as a transaction floor. | [optional] 
**VolumeRanges** | Pointer to [**[]ChargePropertiesVolumeRangesInner**](ChargePropertiesVolumeRangesInner.md) | Volume ranges, sorted from bottom to top tiers, used for a &#x60;volume&#x60; charge model. | [optional] 

## Methods

### NewChargeProperties

`func NewChargeProperties() *ChargeProperties`

NewChargeProperties instantiates a new ChargeProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChargePropertiesWithDefaults

`func NewChargePropertiesWithDefaults() *ChargeProperties`

NewChargePropertiesWithDefaults instantiates a new ChargeProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGraduatedRanges

`func (o *ChargeProperties) GetGraduatedRanges() []ChargePropertiesGraduatedRangesInner`

GetGraduatedRanges returns the GraduatedRanges field if non-nil, zero value otherwise.

### GetGraduatedRangesOk

`func (o *ChargeProperties) GetGraduatedRangesOk() (*[]ChargePropertiesGraduatedRangesInner, bool)`

GetGraduatedRangesOk returns a tuple with the GraduatedRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraduatedRanges

`func (o *ChargeProperties) SetGraduatedRanges(v []ChargePropertiesGraduatedRangesInner)`

SetGraduatedRanges sets GraduatedRanges field to given value.

### HasGraduatedRanges

`func (o *ChargeProperties) HasGraduatedRanges() bool`

HasGraduatedRanges returns a boolean if a field has been set.

### GetGraduatedPercentageRanges

`func (o *ChargeProperties) GetGraduatedPercentageRanges() []ChargePropertiesGraduatedPercentageRangesInner`

GetGraduatedPercentageRanges returns the GraduatedPercentageRanges field if non-nil, zero value otherwise.

### GetGraduatedPercentageRangesOk

`func (o *ChargeProperties) GetGraduatedPercentageRangesOk() (*[]ChargePropertiesGraduatedPercentageRangesInner, bool)`

GetGraduatedPercentageRangesOk returns a tuple with the GraduatedPercentageRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraduatedPercentageRanges

`func (o *ChargeProperties) SetGraduatedPercentageRanges(v []ChargePropertiesGraduatedPercentageRangesInner)`

SetGraduatedPercentageRanges sets GraduatedPercentageRanges field to given value.

### HasGraduatedPercentageRanges

`func (o *ChargeProperties) HasGraduatedPercentageRanges() bool`

HasGraduatedPercentageRanges returns a boolean if a field has been set.

### GetAmount

`func (o *ChargeProperties) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ChargeProperties) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ChargeProperties) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *ChargeProperties) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetFreeUnits

`func (o *ChargeProperties) GetFreeUnits() int32`

GetFreeUnits returns the FreeUnits field if non-nil, zero value otherwise.

### GetFreeUnitsOk

`func (o *ChargeProperties) GetFreeUnitsOk() (*int32, bool)`

GetFreeUnitsOk returns a tuple with the FreeUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeUnits

`func (o *ChargeProperties) SetFreeUnits(v int32)`

SetFreeUnits sets FreeUnits field to given value.

### HasFreeUnits

`func (o *ChargeProperties) HasFreeUnits() bool`

HasFreeUnits returns a boolean if a field has been set.

### GetPackageSize

`func (o *ChargeProperties) GetPackageSize() int32`

GetPackageSize returns the PackageSize field if non-nil, zero value otherwise.

### GetPackageSizeOk

`func (o *ChargeProperties) GetPackageSizeOk() (*int32, bool)`

GetPackageSizeOk returns a tuple with the PackageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageSize

`func (o *ChargeProperties) SetPackageSize(v int32)`

SetPackageSize sets PackageSize field to given value.

### HasPackageSize

`func (o *ChargeProperties) HasPackageSize() bool`

HasPackageSize returns a boolean if a field has been set.

### GetRate

`func (o *ChargeProperties) GetRate() string`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *ChargeProperties) GetRateOk() (*string, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *ChargeProperties) SetRate(v string)`

SetRate sets Rate field to given value.

### HasRate

`func (o *ChargeProperties) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetFixedAmount

`func (o *ChargeProperties) GetFixedAmount() string`

GetFixedAmount returns the FixedAmount field if non-nil, zero value otherwise.

### GetFixedAmountOk

`func (o *ChargeProperties) GetFixedAmountOk() (*string, bool)`

GetFixedAmountOk returns a tuple with the FixedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedAmount

`func (o *ChargeProperties) SetFixedAmount(v string)`

SetFixedAmount sets FixedAmount field to given value.

### HasFixedAmount

`func (o *ChargeProperties) HasFixedAmount() bool`

HasFixedAmount returns a boolean if a field has been set.

### GetFreeUnitsPerEvents

`func (o *ChargeProperties) GetFreeUnitsPerEvents() int32`

GetFreeUnitsPerEvents returns the FreeUnitsPerEvents field if non-nil, zero value otherwise.

### GetFreeUnitsPerEventsOk

`func (o *ChargeProperties) GetFreeUnitsPerEventsOk() (*int32, bool)`

GetFreeUnitsPerEventsOk returns a tuple with the FreeUnitsPerEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeUnitsPerEvents

`func (o *ChargeProperties) SetFreeUnitsPerEvents(v int32)`

SetFreeUnitsPerEvents sets FreeUnitsPerEvents field to given value.

### HasFreeUnitsPerEvents

`func (o *ChargeProperties) HasFreeUnitsPerEvents() bool`

HasFreeUnitsPerEvents returns a boolean if a field has been set.

### SetFreeUnitsPerEventsNil

`func (o *ChargeProperties) SetFreeUnitsPerEventsNil(b bool)`

 SetFreeUnitsPerEventsNil sets the value for FreeUnitsPerEvents to be an explicit nil

### UnsetFreeUnitsPerEvents
`func (o *ChargeProperties) UnsetFreeUnitsPerEvents()`

UnsetFreeUnitsPerEvents ensures that no value is present for FreeUnitsPerEvents, not even an explicit nil
### GetFreeUnitsPerTotalAggregation

`func (o *ChargeProperties) GetFreeUnitsPerTotalAggregation() string`

GetFreeUnitsPerTotalAggregation returns the FreeUnitsPerTotalAggregation field if non-nil, zero value otherwise.

### GetFreeUnitsPerTotalAggregationOk

`func (o *ChargeProperties) GetFreeUnitsPerTotalAggregationOk() (*string, bool)`

GetFreeUnitsPerTotalAggregationOk returns a tuple with the FreeUnitsPerTotalAggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeUnitsPerTotalAggregation

`func (o *ChargeProperties) SetFreeUnitsPerTotalAggregation(v string)`

SetFreeUnitsPerTotalAggregation sets FreeUnitsPerTotalAggregation field to given value.

### HasFreeUnitsPerTotalAggregation

`func (o *ChargeProperties) HasFreeUnitsPerTotalAggregation() bool`

HasFreeUnitsPerTotalAggregation returns a boolean if a field has been set.

### SetFreeUnitsPerTotalAggregationNil

`func (o *ChargeProperties) SetFreeUnitsPerTotalAggregationNil(b bool)`

 SetFreeUnitsPerTotalAggregationNil sets the value for FreeUnitsPerTotalAggregation to be an explicit nil

### UnsetFreeUnitsPerTotalAggregation
`func (o *ChargeProperties) UnsetFreeUnitsPerTotalAggregation()`

UnsetFreeUnitsPerTotalAggregation ensures that no value is present for FreeUnitsPerTotalAggregation, not even an explicit nil
### GetPerTransactionMaxAmount

`func (o *ChargeProperties) GetPerTransactionMaxAmount() string`

GetPerTransactionMaxAmount returns the PerTransactionMaxAmount field if non-nil, zero value otherwise.

### GetPerTransactionMaxAmountOk

`func (o *ChargeProperties) GetPerTransactionMaxAmountOk() (*string, bool)`

GetPerTransactionMaxAmountOk returns a tuple with the PerTransactionMaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerTransactionMaxAmount

`func (o *ChargeProperties) SetPerTransactionMaxAmount(v string)`

SetPerTransactionMaxAmount sets PerTransactionMaxAmount field to given value.

### HasPerTransactionMaxAmount

`func (o *ChargeProperties) HasPerTransactionMaxAmount() bool`

HasPerTransactionMaxAmount returns a boolean if a field has been set.

### SetPerTransactionMaxAmountNil

`func (o *ChargeProperties) SetPerTransactionMaxAmountNil(b bool)`

 SetPerTransactionMaxAmountNil sets the value for PerTransactionMaxAmount to be an explicit nil

### UnsetPerTransactionMaxAmount
`func (o *ChargeProperties) UnsetPerTransactionMaxAmount()`

UnsetPerTransactionMaxAmount ensures that no value is present for PerTransactionMaxAmount, not even an explicit nil
### GetPerTransactionMinAmount

`func (o *ChargeProperties) GetPerTransactionMinAmount() string`

GetPerTransactionMinAmount returns the PerTransactionMinAmount field if non-nil, zero value otherwise.

### GetPerTransactionMinAmountOk

`func (o *ChargeProperties) GetPerTransactionMinAmountOk() (*string, bool)`

GetPerTransactionMinAmountOk returns a tuple with the PerTransactionMinAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerTransactionMinAmount

`func (o *ChargeProperties) SetPerTransactionMinAmount(v string)`

SetPerTransactionMinAmount sets PerTransactionMinAmount field to given value.

### HasPerTransactionMinAmount

`func (o *ChargeProperties) HasPerTransactionMinAmount() bool`

HasPerTransactionMinAmount returns a boolean if a field has been set.

### SetPerTransactionMinAmountNil

`func (o *ChargeProperties) SetPerTransactionMinAmountNil(b bool)`

 SetPerTransactionMinAmountNil sets the value for PerTransactionMinAmount to be an explicit nil

### UnsetPerTransactionMinAmount
`func (o *ChargeProperties) UnsetPerTransactionMinAmount()`

UnsetPerTransactionMinAmount ensures that no value is present for PerTransactionMinAmount, not even an explicit nil
### GetVolumeRanges

`func (o *ChargeProperties) GetVolumeRanges() []ChargePropertiesVolumeRangesInner`

GetVolumeRanges returns the VolumeRanges field if non-nil, zero value otherwise.

### GetVolumeRangesOk

`func (o *ChargeProperties) GetVolumeRangesOk() (*[]ChargePropertiesVolumeRangesInner, bool)`

GetVolumeRangesOk returns a tuple with the VolumeRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeRanges

`func (o *ChargeProperties) SetVolumeRanges(v []ChargePropertiesVolumeRangesInner)`

SetVolumeRanges sets VolumeRanges field to given value.

### HasVolumeRanges

`func (o *ChargeProperties) HasVolumeRanges() bool`

HasVolumeRanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


