# InvoiceUpdateInputInvoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentStatus** | Pointer to **string** | The payment status of the invoice. Possible values are &#x60;pending&#x60;, &#x60;failed&#x60; or &#x60;succeeded&#x60;. | [optional] 
**Metadata** | Pointer to [**[]InvoiceUpdateInputInvoiceMetadataInner**](InvoiceUpdateInputInvoiceMetadataInner.md) |  | [optional] 

## Methods

### NewInvoiceUpdateInputInvoice

`func NewInvoiceUpdateInputInvoice() *InvoiceUpdateInputInvoice`

NewInvoiceUpdateInputInvoice instantiates a new InvoiceUpdateInputInvoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceUpdateInputInvoiceWithDefaults

`func NewInvoiceUpdateInputInvoiceWithDefaults() *InvoiceUpdateInputInvoice`

NewInvoiceUpdateInputInvoiceWithDefaults instantiates a new InvoiceUpdateInputInvoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentStatus

`func (o *InvoiceUpdateInputInvoice) GetPaymentStatus() string`

GetPaymentStatus returns the PaymentStatus field if non-nil, zero value otherwise.

### GetPaymentStatusOk

`func (o *InvoiceUpdateInputInvoice) GetPaymentStatusOk() (*string, bool)`

GetPaymentStatusOk returns a tuple with the PaymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentStatus

`func (o *InvoiceUpdateInputInvoice) SetPaymentStatus(v string)`

SetPaymentStatus sets PaymentStatus field to given value.

### HasPaymentStatus

`func (o *InvoiceUpdateInputInvoice) HasPaymentStatus() bool`

HasPaymentStatus returns a boolean if a field has been set.

### GetMetadata

`func (o *InvoiceUpdateInputInvoice) GetMetadata() []InvoiceUpdateInputInvoiceMetadataInner`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *InvoiceUpdateInputInvoice) GetMetadataOk() (*[]InvoiceUpdateInputInvoiceMetadataInner, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *InvoiceUpdateInputInvoice) SetMetadata(v []InvoiceUpdateInputInvoiceMetadataInner)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *InvoiceUpdateInputInvoice) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


