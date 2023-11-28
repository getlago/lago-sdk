# GroupPropertiesObjectValues

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

### NewGroupPropertiesObjectValues

`func NewGroupPropertiesObjectValues() *GroupPropertiesObjectValues`

NewGroupPropertiesObjectValues instantiates a new GroupPropertiesObjectValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupPropertiesObjectValuesWithDefaults

`func NewGroupPropertiesObjectValuesWithDefaults() *GroupPropertiesObjectValues`

NewGroupPropertiesObjectValuesWithDefaults instantiates a new GroupPropertiesObjectValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGraduatedRanges

`func (o *GroupPropertiesObjectValues) GetGraduatedRanges() []ChargePropertiesGraduatedRangesInner`

GetGraduatedRanges returns the GraduatedRanges field if non-nil, zero value otherwise.

### GetGraduatedRangesOk

`func (o *GroupPropertiesObjectValues) GetGraduatedRangesOk() (*[]ChargePropertiesGraduatedRangesInner, bool)`

GetGraduatedRangesOk returns a tuple with the GraduatedRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraduatedRanges

`func (o *GroupPropertiesObjectValues) SetGraduatedRanges(v []ChargePropertiesGraduatedRangesInner)`

SetGraduatedRanges sets GraduatedRanges field to given value.

### HasGraduatedRanges

`func (o *GroupPropertiesObjectValues) HasGraduatedRanges() bool`

HasGraduatedRanges returns a boolean if a field has been set.

### GetGraduatedPercentageRanges

`func (o *GroupPropertiesObjectValues) GetGraduatedPercentageRanges() []ChargePropertiesGraduatedPercentageRangesInner`

GetGraduatedPercentageRanges returns the GraduatedPercentageRanges field if non-nil, zero value otherwise.

### GetGraduatedPercentageRangesOk

`func (o *GroupPropertiesObjectValues) GetGraduatedPercentageRangesOk() (*[]ChargePropertiesGraduatedPercentageRangesInner, bool)`

GetGraduatedPercentageRangesOk returns a tuple with the GraduatedPercentageRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraduatedPercentageRanges

`func (o *GroupPropertiesObjectValues) SetGraduatedPercentageRanges(v []ChargePropertiesGraduatedPercentageRangesInner)`

SetGraduatedPercentageRanges sets GraduatedPercentageRanges field to given value.

### HasGraduatedPercentageRanges

`func (o *GroupPropertiesObjectValues) HasGraduatedPercentageRanges() bool`

HasGraduatedPercentageRanges returns a boolean if a field has been set.

### GetAmount

`func (o *GroupPropertiesObjectValues) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *GroupPropertiesObjectValues) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *GroupPropertiesObjectValues) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *GroupPropertiesObjectValues) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetFreeUnits

`func (o *GroupPropertiesObjectValues) GetFreeUnits() int32`

GetFreeUnits returns the FreeUnits field if non-nil, zero value otherwise.

### GetFreeUnitsOk

`func (o *GroupPropertiesObjectValues) GetFreeUnitsOk() (*int32, bool)`

GetFreeUnitsOk returns a tuple with the FreeUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeUnits

`func (o *GroupPropertiesObjectValues) SetFreeUnits(v int32)`

SetFreeUnits sets FreeUnits field to given value.

### HasFreeUnits

`func (o *GroupPropertiesObjectValues) HasFreeUnits() bool`

HasFreeUnits returns a boolean if a field has been set.

### GetPackageSize

`func (o *GroupPropertiesObjectValues) GetPackageSize() int32`

GetPackageSize returns the PackageSize field if non-nil, zero value otherwise.

### GetPackageSizeOk

`func (o *GroupPropertiesObjectValues) GetPackageSizeOk() (*int32, bool)`

GetPackageSizeOk returns a tuple with the PackageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageSize

`func (o *GroupPropertiesObjectValues) SetPackageSize(v int32)`

SetPackageSize sets PackageSize field to given value.

### HasPackageSize

`func (o *GroupPropertiesObjectValues) HasPackageSize() bool`

HasPackageSize returns a boolean if a field has been set.

### GetRate

`func (o *GroupPropertiesObjectValues) GetRate() string`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *GroupPropertiesObjectValues) GetRateOk() (*string, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *GroupPropertiesObjectValues) SetRate(v string)`

SetRate sets Rate field to given value.

### HasRate

`func (o *GroupPropertiesObjectValues) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetFixedAmount

`func (o *GroupPropertiesObjectValues) GetFixedAmount() string`

GetFixedAmount returns the FixedAmount field if non-nil, zero value otherwise.

### GetFixedAmountOk

`func (o *GroupPropertiesObjectValues) GetFixedAmountOk() (*string, bool)`

GetFixedAmountOk returns a tuple with the FixedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedAmount

`func (o *GroupPropertiesObjectValues) SetFixedAmount(v string)`

SetFixedAmount sets FixedAmount field to given value.

### HasFixedAmount

`func (o *GroupPropertiesObjectValues) HasFixedAmount() bool`

HasFixedAmount returns a boolean if a field has been set.

### GetFreeUnitsPerEvents

`func (o *GroupPropertiesObjectValues) GetFreeUnitsPerEvents() int32`

GetFreeUnitsPerEvents returns the FreeUnitsPerEvents field if non-nil, zero value otherwise.

### GetFreeUnitsPerEventsOk

`func (o *GroupPropertiesObjectValues) GetFreeUnitsPerEventsOk() (*int32, bool)`

GetFreeUnitsPerEventsOk returns a tuple with the FreeUnitsPerEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeUnitsPerEvents

`func (o *GroupPropertiesObjectValues) SetFreeUnitsPerEvents(v int32)`

SetFreeUnitsPerEvents sets FreeUnitsPerEvents field to given value.

### HasFreeUnitsPerEvents

`func (o *GroupPropertiesObjectValues) HasFreeUnitsPerEvents() bool`

HasFreeUnitsPerEvents returns a boolean if a field has been set.

### SetFreeUnitsPerEventsNil

`func (o *GroupPropertiesObjectValues) SetFreeUnitsPerEventsNil(b bool)`

 SetFreeUnitsPerEventsNil sets the value for FreeUnitsPerEvents to be an explicit nil

### UnsetFreeUnitsPerEvents
`func (o *GroupPropertiesObjectValues) UnsetFreeUnitsPerEvents()`

UnsetFreeUnitsPerEvents ensures that no value is present for FreeUnitsPerEvents, not even an explicit nil
### GetFreeUnitsPerTotalAggregation

`func (o *GroupPropertiesObjectValues) GetFreeUnitsPerTotalAggregation() string`

GetFreeUnitsPerTotalAggregation returns the FreeUnitsPerTotalAggregation field if non-nil, zero value otherwise.

### GetFreeUnitsPerTotalAggregationOk

`func (o *GroupPropertiesObjectValues) GetFreeUnitsPerTotalAggregationOk() (*string, bool)`

GetFreeUnitsPerTotalAggregationOk returns a tuple with the FreeUnitsPerTotalAggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeUnitsPerTotalAggregation

`func (o *GroupPropertiesObjectValues) SetFreeUnitsPerTotalAggregation(v string)`

SetFreeUnitsPerTotalAggregation sets FreeUnitsPerTotalAggregation field to given value.

### HasFreeUnitsPerTotalAggregation

`func (o *GroupPropertiesObjectValues) HasFreeUnitsPerTotalAggregation() bool`

HasFreeUnitsPerTotalAggregation returns a boolean if a field has been set.

### SetFreeUnitsPerTotalAggregationNil

`func (o *GroupPropertiesObjectValues) SetFreeUnitsPerTotalAggregationNil(b bool)`

 SetFreeUnitsPerTotalAggregationNil sets the value for FreeUnitsPerTotalAggregation to be an explicit nil

### UnsetFreeUnitsPerTotalAggregation
`func (o *GroupPropertiesObjectValues) UnsetFreeUnitsPerTotalAggregation()`

UnsetFreeUnitsPerTotalAggregation ensures that no value is present for FreeUnitsPerTotalAggregation, not even an explicit nil
### GetPerTransactionMaxAmount

`func (o *GroupPropertiesObjectValues) GetPerTransactionMaxAmount() string`

GetPerTransactionMaxAmount returns the PerTransactionMaxAmount field if non-nil, zero value otherwise.

### GetPerTransactionMaxAmountOk

`func (o *GroupPropertiesObjectValues) GetPerTransactionMaxAmountOk() (*string, bool)`

GetPerTransactionMaxAmountOk returns a tuple with the PerTransactionMaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerTransactionMaxAmount

`func (o *GroupPropertiesObjectValues) SetPerTransactionMaxAmount(v string)`

SetPerTransactionMaxAmount sets PerTransactionMaxAmount field to given value.

### HasPerTransactionMaxAmount

`func (o *GroupPropertiesObjectValues) HasPerTransactionMaxAmount() bool`

HasPerTransactionMaxAmount returns a boolean if a field has been set.

### SetPerTransactionMaxAmountNil

`func (o *GroupPropertiesObjectValues) SetPerTransactionMaxAmountNil(b bool)`

 SetPerTransactionMaxAmountNil sets the value for PerTransactionMaxAmount to be an explicit nil

### UnsetPerTransactionMaxAmount
`func (o *GroupPropertiesObjectValues) UnsetPerTransactionMaxAmount()`

UnsetPerTransactionMaxAmount ensures that no value is present for PerTransactionMaxAmount, not even an explicit nil
### GetPerTransactionMinAmount

`func (o *GroupPropertiesObjectValues) GetPerTransactionMinAmount() string`

GetPerTransactionMinAmount returns the PerTransactionMinAmount field if non-nil, zero value otherwise.

### GetPerTransactionMinAmountOk

`func (o *GroupPropertiesObjectValues) GetPerTransactionMinAmountOk() (*string, bool)`

GetPerTransactionMinAmountOk returns a tuple with the PerTransactionMinAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerTransactionMinAmount

`func (o *GroupPropertiesObjectValues) SetPerTransactionMinAmount(v string)`

SetPerTransactionMinAmount sets PerTransactionMinAmount field to given value.

### HasPerTransactionMinAmount

`func (o *GroupPropertiesObjectValues) HasPerTransactionMinAmount() bool`

HasPerTransactionMinAmount returns a boolean if a field has been set.

### SetPerTransactionMinAmountNil

`func (o *GroupPropertiesObjectValues) SetPerTransactionMinAmountNil(b bool)`

 SetPerTransactionMinAmountNil sets the value for PerTransactionMinAmount to be an explicit nil

### UnsetPerTransactionMinAmount
`func (o *GroupPropertiesObjectValues) UnsetPerTransactionMinAmount()`

UnsetPerTransactionMinAmount ensures that no value is present for PerTransactionMinAmount, not even an explicit nil
### GetVolumeRanges

`func (o *GroupPropertiesObjectValues) GetVolumeRanges() []ChargePropertiesVolumeRangesInner`

GetVolumeRanges returns the VolumeRanges field if non-nil, zero value otherwise.

### GetVolumeRangesOk

`func (o *GroupPropertiesObjectValues) GetVolumeRangesOk() (*[]ChargePropertiesVolumeRangesInner, bool)`

GetVolumeRangesOk returns a tuple with the VolumeRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeRanges

`func (o *GroupPropertiesObjectValues) SetVolumeRanges(v []ChargePropertiesVolumeRangesInner)`

SetVolumeRanges sets VolumeRanges field to given value.

### HasVolumeRanges

`func (o *GroupPropertiesObjectValues) HasVolumeRanges() bool`

HasVolumeRanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


