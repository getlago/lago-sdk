# InvoiceMetadataObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LagoId** | Pointer to **string** | Unique identifier assigned to the invoice metadata within the Lago application. | [optional] 
**Key** | Pointer to **string** | Represents the key of the metadata’s key-value pair. | [optional] 
**Value** | Pointer to **string** | Represents the value of the metadata’s key-value pair. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date and time when the metadata object was created. It follows the ISO 8601 datetime format and is expressed in Coordinated Universal Time (UTC). | [optional] 

## Methods

### NewInvoiceMetadataObject

`func NewInvoiceMetadataObject() *InvoiceMetadataObject`

NewInvoiceMetadataObject instantiates a new InvoiceMetadataObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceMetadataObjectWithDefaults

`func NewInvoiceMetadataObjectWithDefaults() *InvoiceMetadataObject`

NewInvoiceMetadataObjectWithDefaults instantiates a new InvoiceMetadataObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLagoId

`func (o *InvoiceMetadataObject) GetLagoId() string`

GetLagoId returns the LagoId field if non-nil, zero value otherwise.

### GetLagoIdOk

`func (o *InvoiceMetadataObject) GetLagoIdOk() (*string, bool)`

GetLagoIdOk returns a tuple with the LagoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagoId

`func (o *InvoiceMetadataObject) SetLagoId(v string)`

SetLagoId sets LagoId field to given value.

### HasLagoId

`func (o *InvoiceMetadataObject) HasLagoId() bool`

HasLagoId returns a boolean if a field has been set.

### GetKey

`func (o *InvoiceMetadataObject) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *InvoiceMetadataObject) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *InvoiceMetadataObject) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *InvoiceMetadataObject) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *InvoiceMetadataObject) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *InvoiceMetadataObject) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *InvoiceMetadataObject) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *InvoiceMetadataObject) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InvoiceMetadataObject) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InvoiceMetadataObject) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InvoiceMetadataObject) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InvoiceMetadataObject) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


