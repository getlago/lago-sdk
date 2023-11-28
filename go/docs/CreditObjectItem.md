# CreditObjectItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LagoItemId** | **string** | Unique identifier assigned to the credit item within the Lago application. | 
**Type** | **string** | The type of credit applied. Possible values are &#x60;coupon&#x60; or &#x60;credit_note&#x60;. | 
**Code** | **string** | The code of the credit applied. It can be the code of the coupon attached to the credit or the credit note&#39;s number. | 
**Name** | **string** | The name of the credit applied. It can be the name of the coupon attached to the credit or the initial invoice&#39;s number linked to the credit note. | 

## Methods

### NewCreditObjectItem

`func NewCreditObjectItem(lagoItemId string, type_ string, code string, name string, ) *CreditObjectItem`

NewCreditObjectItem instantiates a new CreditObjectItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditObjectItemWithDefaults

`func NewCreditObjectItemWithDefaults() *CreditObjectItem`

NewCreditObjectItemWithDefaults instantiates a new CreditObjectItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLagoItemId

`func (o *CreditObjectItem) GetLagoItemId() string`

GetLagoItemId returns the LagoItemId field if non-nil, zero value otherwise.

### GetLagoItemIdOk

`func (o *CreditObjectItem) GetLagoItemIdOk() (*string, bool)`

GetLagoItemIdOk returns a tuple with the LagoItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagoItemId

`func (o *CreditObjectItem) SetLagoItemId(v string)`

SetLagoItemId sets LagoItemId field to given value.


### GetType

`func (o *CreditObjectItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreditObjectItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreditObjectItem) SetType(v string)`

SetType sets Type field to given value.


### GetCode

`func (o *CreditObjectItem) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CreditObjectItem) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CreditObjectItem) SetCode(v string)`

SetCode sets Code field to given value.


### GetName

`func (o *CreditObjectItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreditObjectItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreditObjectItem) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


