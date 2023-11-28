# CreditNoteUpdateInputCreditNote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RefundStatus** | **string** | The status of the refund portion of the credit note. It indicates the current state or condition of the refund associated with the credit note. The possible values for this field are:  - &#x60;pending&#x60;: this status indicates that the refund is pending execution. The refund request has been initiated but has not been processed or completed yet. - &#x60;succeeded&#x60;: this status indicates that the refund has been successfully executed. The refund amount has been processed and returned to the customer or the designated recipient. - &#x60;failed&#x60;: this status indicates that the refund failed to execute. The refund request encountered an error or unsuccessful processing, and the refund amount could not be returned. | 

## Methods

### NewCreditNoteUpdateInputCreditNote

`func NewCreditNoteUpdateInputCreditNote(refundStatus string, ) *CreditNoteUpdateInputCreditNote`

NewCreditNoteUpdateInputCreditNote instantiates a new CreditNoteUpdateInputCreditNote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditNoteUpdateInputCreditNoteWithDefaults

`func NewCreditNoteUpdateInputCreditNoteWithDefaults() *CreditNoteUpdateInputCreditNote`

NewCreditNoteUpdateInputCreditNoteWithDefaults instantiates a new CreditNoteUpdateInputCreditNote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRefundStatus

`func (o *CreditNoteUpdateInputCreditNote) GetRefundStatus() string`

GetRefundStatus returns the RefundStatus field if non-nil, zero value otherwise.

### GetRefundStatusOk

`func (o *CreditNoteUpdateInputCreditNote) GetRefundStatusOk() (*string, bool)`

GetRefundStatusOk returns a tuple with the RefundStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundStatus

`func (o *CreditNoteUpdateInputCreditNote) SetRefundStatus(v string)`

SetRefundStatus sets RefundStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


