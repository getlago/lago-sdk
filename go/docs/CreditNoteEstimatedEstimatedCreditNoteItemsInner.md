# CreditNoteEstimatedEstimatedCreditNoteItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountCents** | **int32** | The credit note’s item amount, expressed in cents. | 
**LagoFeeId** | **NullableString** | Unique identifier assigned to the fee within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the fee’s record within the Lago system. | 

## Methods

### NewCreditNoteEstimatedEstimatedCreditNoteItemsInner

`func NewCreditNoteEstimatedEstimatedCreditNoteItemsInner(amountCents int32, lagoFeeId NullableString, ) *CreditNoteEstimatedEstimatedCreditNoteItemsInner`

NewCreditNoteEstimatedEstimatedCreditNoteItemsInner instantiates a new CreditNoteEstimatedEstimatedCreditNoteItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditNoteEstimatedEstimatedCreditNoteItemsInnerWithDefaults

`func NewCreditNoteEstimatedEstimatedCreditNoteItemsInnerWithDefaults() *CreditNoteEstimatedEstimatedCreditNoteItemsInner`

NewCreditNoteEstimatedEstimatedCreditNoteItemsInnerWithDefaults instantiates a new CreditNoteEstimatedEstimatedCreditNoteItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountCents

`func (o *CreditNoteEstimatedEstimatedCreditNoteItemsInner) GetAmountCents() int32`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *CreditNoteEstimatedEstimatedCreditNoteItemsInner) GetAmountCentsOk() (*int32, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *CreditNoteEstimatedEstimatedCreditNoteItemsInner) SetAmountCents(v int32)`

SetAmountCents sets AmountCents field to given value.


### GetLagoFeeId

`func (o *CreditNoteEstimatedEstimatedCreditNoteItemsInner) GetLagoFeeId() string`

GetLagoFeeId returns the LagoFeeId field if non-nil, zero value otherwise.

### GetLagoFeeIdOk

`func (o *CreditNoteEstimatedEstimatedCreditNoteItemsInner) GetLagoFeeIdOk() (*string, bool)`

GetLagoFeeIdOk returns a tuple with the LagoFeeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagoFeeId

`func (o *CreditNoteEstimatedEstimatedCreditNoteItemsInner) SetLagoFeeId(v string)`

SetLagoFeeId sets LagoFeeId field to given value.


### SetLagoFeeIdNil

`func (o *CreditNoteEstimatedEstimatedCreditNoteItemsInner) SetLagoFeeIdNil(b bool)`

 SetLagoFeeIdNil sets the value for LagoFeeId to be an explicit nil

### UnsetLagoFeeId
`func (o *CreditNoteEstimatedEstimatedCreditNoteItemsInner) UnsetLagoFeeId()`

UnsetLagoFeeId ensures that no value is present for LagoFeeId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


