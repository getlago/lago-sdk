# FeeObjectItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The fee type. Possible values are &#x60;add-on&#x60;, &#x60;charge&#x60;, &#x60;credit&#x60; or &#x60;subscription&#x60;. | 
**Code** | **string** | The code of the fee item. It can be the code of the &#x60;add-o&#x60;n, the code of the &#x60;charge&#x60;, the code of the &#x60;credit&#x60; or the code of the &#x60;subscription&#x60;. | 
**Name** | **string** | The name of the fee item. It can be the name of the &#x60;add-on&#x60;, the name of the &#x60;charge&#x60;, the name of the &#x60;credit&#x60; or the name of the &#x60;subscription&#x60;. | 
**InvoiceDisplayName** | Pointer to **string** | Specifies the name that will be displayed on an invoice. If no value is set for this field, the name of the actual charge will be used as the default display name. | [optional] 
**GroupInvoiceDisplayName** | Pointer to **string** | Specifies the name that will be displayed on an invoice. If no value is set for this field, the name of the actual charge will be used as the default display name. | [optional] 
**LagoItemId** | **string** | Unique identifier of the fee item, created by Lago. It can be the identifier of the &#x60;add-on&#x60;, the identifier of the &#x60;charge&#x60;, the identifier of the &#x60;credit&#x60; or the identifier of the &#x60;subscription&#x60;. | 
**ItemType** | **string** | The type of the fee item. Possible values are &#x60;AddOn&#x60;, &#x60;BillableMetric&#x60;, &#x60;WalletTransaction&#x60; or &#x60;Subscription&#x60;. | 

## Methods

### NewFeeObjectItem

`func NewFeeObjectItem(type_ string, code string, name string, lagoItemId string, itemType string, ) *FeeObjectItem`

NewFeeObjectItem instantiates a new FeeObjectItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeeObjectItemWithDefaults

`func NewFeeObjectItemWithDefaults() *FeeObjectItem`

NewFeeObjectItemWithDefaults instantiates a new FeeObjectItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FeeObjectItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FeeObjectItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FeeObjectItem) SetType(v string)`

SetType sets Type field to given value.


### GetCode

`func (o *FeeObjectItem) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *FeeObjectItem) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *FeeObjectItem) SetCode(v string)`

SetCode sets Code field to given value.


### GetName

`func (o *FeeObjectItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FeeObjectItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FeeObjectItem) SetName(v string)`

SetName sets Name field to given value.


### GetInvoiceDisplayName

`func (o *FeeObjectItem) GetInvoiceDisplayName() string`

GetInvoiceDisplayName returns the InvoiceDisplayName field if non-nil, zero value otherwise.

### GetInvoiceDisplayNameOk

`func (o *FeeObjectItem) GetInvoiceDisplayNameOk() (*string, bool)`

GetInvoiceDisplayNameOk returns a tuple with the InvoiceDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceDisplayName

`func (o *FeeObjectItem) SetInvoiceDisplayName(v string)`

SetInvoiceDisplayName sets InvoiceDisplayName field to given value.

### HasInvoiceDisplayName

`func (o *FeeObjectItem) HasInvoiceDisplayName() bool`

HasInvoiceDisplayName returns a boolean if a field has been set.

### GetGroupInvoiceDisplayName

`func (o *FeeObjectItem) GetGroupInvoiceDisplayName() string`

GetGroupInvoiceDisplayName returns the GroupInvoiceDisplayName field if non-nil, zero value otherwise.

### GetGroupInvoiceDisplayNameOk

`func (o *FeeObjectItem) GetGroupInvoiceDisplayNameOk() (*string, bool)`

GetGroupInvoiceDisplayNameOk returns a tuple with the GroupInvoiceDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupInvoiceDisplayName

`func (o *FeeObjectItem) SetGroupInvoiceDisplayName(v string)`

SetGroupInvoiceDisplayName sets GroupInvoiceDisplayName field to given value.

### HasGroupInvoiceDisplayName

`func (o *FeeObjectItem) HasGroupInvoiceDisplayName() bool`

HasGroupInvoiceDisplayName returns a boolean if a field has been set.

### GetLagoItemId

`func (o *FeeObjectItem) GetLagoItemId() string`

GetLagoItemId returns the LagoItemId field if non-nil, zero value otherwise.

### GetLagoItemIdOk

`func (o *FeeObjectItem) GetLagoItemIdOk() (*string, bool)`

GetLagoItemIdOk returns a tuple with the LagoItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagoItemId

`func (o *FeeObjectItem) SetLagoItemId(v string)`

SetLagoItemId sets LagoItemId field to given value.


### GetItemType

`func (o *FeeObjectItem) GetItemType() string`

GetItemType returns the ItemType field if non-nil, zero value otherwise.

### GetItemTypeOk

`func (o *FeeObjectItem) GetItemTypeOk() (*string, bool)`

GetItemTypeOk returns a tuple with the ItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemType

`func (o *FeeObjectItem) SetItemType(v string)`

SetItemType sets ItemType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


