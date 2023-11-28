# CustomerChargeUsageObjectCharge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LagoId** | **string** | Unique identifier assigned to the charge within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the charge’s record within the Lago system. | 
**ChargeModel** | **string** | The pricing model applied to this charge. Possible values are standard, &#x60;graduated&#x60;, &#x60;percentage&#x60;, &#x60;package&#x60; or &#x60;volume&#x60;. | 
**InvoiceDisplayName** | Pointer to **string** | Specifies the name that will be displayed on an invoice. If no value is set for this field, the name of the actual charge will be used as the default display name. | [optional] 

## Methods

### NewCustomerChargeUsageObjectCharge

`func NewCustomerChargeUsageObjectCharge(lagoId string, chargeModel string, ) *CustomerChargeUsageObjectCharge`

NewCustomerChargeUsageObjectCharge instantiates a new CustomerChargeUsageObjectCharge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerChargeUsageObjectChargeWithDefaults

`func NewCustomerChargeUsageObjectChargeWithDefaults() *CustomerChargeUsageObjectCharge`

NewCustomerChargeUsageObjectChargeWithDefaults instantiates a new CustomerChargeUsageObjectCharge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLagoId

`func (o *CustomerChargeUsageObjectCharge) GetLagoId() string`

GetLagoId returns the LagoId field if non-nil, zero value otherwise.

### GetLagoIdOk

`func (o *CustomerChargeUsageObjectCharge) GetLagoIdOk() (*string, bool)`

GetLagoIdOk returns a tuple with the LagoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagoId

`func (o *CustomerChargeUsageObjectCharge) SetLagoId(v string)`

SetLagoId sets LagoId field to given value.


### GetChargeModel

`func (o *CustomerChargeUsageObjectCharge) GetChargeModel() string`

GetChargeModel returns the ChargeModel field if non-nil, zero value otherwise.

### GetChargeModelOk

`func (o *CustomerChargeUsageObjectCharge) GetChargeModelOk() (*string, bool)`

GetChargeModelOk returns a tuple with the ChargeModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeModel

`func (o *CustomerChargeUsageObjectCharge) SetChargeModel(v string)`

SetChargeModel sets ChargeModel field to given value.


### GetInvoiceDisplayName

`func (o *CustomerChargeUsageObjectCharge) GetInvoiceDisplayName() string`

GetInvoiceDisplayName returns the InvoiceDisplayName field if non-nil, zero value otherwise.

### GetInvoiceDisplayNameOk

`func (o *CustomerChargeUsageObjectCharge) GetInvoiceDisplayNameOk() (*string, bool)`

GetInvoiceDisplayNameOk returns a tuple with the InvoiceDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceDisplayName

`func (o *CustomerChargeUsageObjectCharge) SetInvoiceDisplayName(v string)`

SetInvoiceDisplayName sets InvoiceDisplayName field to given value.

### HasInvoiceDisplayName

`func (o *CustomerChargeUsageObjectCharge) HasInvoiceDisplayName() bool`

HasInvoiceDisplayName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


