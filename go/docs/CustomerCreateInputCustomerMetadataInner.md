# CustomerCreateInputCustomerMetadataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Identifier for the metadata object, only required when updating a key-value pair | [optional] 
**Key** | **string** | The metadata object key | 
**Value** | **string** | The metadata object value | 
**DisplayInInvoice** | **bool** | Determines whether the item or information should be displayed in the invoice. If set to true, the item or information will be included and visible in the generated invoice. If set to false, the item or information will be excluded and not displayed in the invoice. | 

## Methods

### NewCustomerCreateInputCustomerMetadataInner

`func NewCustomerCreateInputCustomerMetadataInner(key string, value string, displayInInvoice bool, ) *CustomerCreateInputCustomerMetadataInner`

NewCustomerCreateInputCustomerMetadataInner instantiates a new CustomerCreateInputCustomerMetadataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerCreateInputCustomerMetadataInnerWithDefaults

`func NewCustomerCreateInputCustomerMetadataInnerWithDefaults() *CustomerCreateInputCustomerMetadataInner`

NewCustomerCreateInputCustomerMetadataInnerWithDefaults instantiates a new CustomerCreateInputCustomerMetadataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomerCreateInputCustomerMetadataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomerCreateInputCustomerMetadataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomerCreateInputCustomerMetadataInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustomerCreateInputCustomerMetadataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *CustomerCreateInputCustomerMetadataInner) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CustomerCreateInputCustomerMetadataInner) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CustomerCreateInputCustomerMetadataInner) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *CustomerCreateInputCustomerMetadataInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CustomerCreateInputCustomerMetadataInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CustomerCreateInputCustomerMetadataInner) SetValue(v string)`

SetValue sets Value field to given value.


### GetDisplayInInvoice

`func (o *CustomerCreateInputCustomerMetadataInner) GetDisplayInInvoice() bool`

GetDisplayInInvoice returns the DisplayInInvoice field if non-nil, zero value otherwise.

### GetDisplayInInvoiceOk

`func (o *CustomerCreateInputCustomerMetadataInner) GetDisplayInInvoiceOk() (*bool, bool)`

GetDisplayInInvoiceOk returns a tuple with the DisplayInInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayInInvoice

`func (o *CustomerCreateInputCustomerMetadataInner) SetDisplayInInvoice(v bool)`

SetDisplayInInvoice sets DisplayInInvoice field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


