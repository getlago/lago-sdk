# CustomerMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LagoId** | **string** | A unique identifier for the customer metadata object in the Lago application. Can be used to update a key-value pair | 
**Key** | **string** | The metadata object key | 
**Value** | **string** | The metadata object value | 
**DisplayInInvoice** | **bool** | Determines whether the item or information should be displayed in the invoice. If set to true, the item or information will be included and visible in the generated invoice. If set to false, the item or information will be excluded and not displayed in the invoice. | 
**CreatedAt** | **time.Time** | The date of the metadata object creation, represented in ISO 8601 datetime format and expressed in Coordinated Universal Time (UTC). The creation_date provides a standardized and internationally recognized timestamp for when the metadata object was created | 

## Methods

### NewCustomerMetadata

`func NewCustomerMetadata(lagoId string, key string, value string, displayInInvoice bool, createdAt time.Time, ) *CustomerMetadata`

NewCustomerMetadata instantiates a new CustomerMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerMetadataWithDefaults

`func NewCustomerMetadataWithDefaults() *CustomerMetadata`

NewCustomerMetadataWithDefaults instantiates a new CustomerMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLagoId

`func (o *CustomerMetadata) GetLagoId() string`

GetLagoId returns the LagoId field if non-nil, zero value otherwise.

### GetLagoIdOk

`func (o *CustomerMetadata) GetLagoIdOk() (*string, bool)`

GetLagoIdOk returns a tuple with the LagoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagoId

`func (o *CustomerMetadata) SetLagoId(v string)`

SetLagoId sets LagoId field to given value.


### GetKey

`func (o *CustomerMetadata) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CustomerMetadata) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CustomerMetadata) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *CustomerMetadata) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CustomerMetadata) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CustomerMetadata) SetValue(v string)`

SetValue sets Value field to given value.


### GetDisplayInInvoice

`func (o *CustomerMetadata) GetDisplayInInvoice() bool`

GetDisplayInInvoice returns the DisplayInInvoice field if non-nil, zero value otherwise.

### GetDisplayInInvoiceOk

`func (o *CustomerMetadata) GetDisplayInInvoiceOk() (*bool, bool)`

GetDisplayInInvoiceOk returns a tuple with the DisplayInInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayInInvoice

`func (o *CustomerMetadata) SetDisplayInInvoice(v bool)`

SetDisplayInInvoice sets DisplayInInvoice field to given value.


### GetCreatedAt

`func (o *CustomerMetadata) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CustomerMetadata) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CustomerMetadata) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


