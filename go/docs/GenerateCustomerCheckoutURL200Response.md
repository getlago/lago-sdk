# GenerateCustomerCheckoutURL200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LagoCustomerId** | Pointer to **string** | Unique identifier assigned to the customer within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the customer&#39;s record within the Lago system | [optional] 
**ExternalCustomerId** | Pointer to **string** | The customer external unique identifier (provided by your own application) | [optional] 
**PaymentProvider** | Pointer to **string** | The Payment Provider name linked to the Customer. | [optional] 
**CheckoutUrl** | Pointer to **string** | The new generated Payment Provider Checkout URL for the Customer. | [optional] 

## Methods

### NewGenerateCustomerCheckoutURL200Response

`func NewGenerateCustomerCheckoutURL200Response() *GenerateCustomerCheckoutURL200Response`

NewGenerateCustomerCheckoutURL200Response instantiates a new GenerateCustomerCheckoutURL200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateCustomerCheckoutURL200ResponseWithDefaults

`func NewGenerateCustomerCheckoutURL200ResponseWithDefaults() *GenerateCustomerCheckoutURL200Response`

NewGenerateCustomerCheckoutURL200ResponseWithDefaults instantiates a new GenerateCustomerCheckoutURL200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLagoCustomerId

`func (o *GenerateCustomerCheckoutURL200Response) GetLagoCustomerId() string`

GetLagoCustomerId returns the LagoCustomerId field if non-nil, zero value otherwise.

### GetLagoCustomerIdOk

`func (o *GenerateCustomerCheckoutURL200Response) GetLagoCustomerIdOk() (*string, bool)`

GetLagoCustomerIdOk returns a tuple with the LagoCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagoCustomerId

`func (o *GenerateCustomerCheckoutURL200Response) SetLagoCustomerId(v string)`

SetLagoCustomerId sets LagoCustomerId field to given value.

### HasLagoCustomerId

`func (o *GenerateCustomerCheckoutURL200Response) HasLagoCustomerId() bool`

HasLagoCustomerId returns a boolean if a field has been set.

### GetExternalCustomerId

`func (o *GenerateCustomerCheckoutURL200Response) GetExternalCustomerId() string`

GetExternalCustomerId returns the ExternalCustomerId field if non-nil, zero value otherwise.

### GetExternalCustomerIdOk

`func (o *GenerateCustomerCheckoutURL200Response) GetExternalCustomerIdOk() (*string, bool)`

GetExternalCustomerIdOk returns a tuple with the ExternalCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalCustomerId

`func (o *GenerateCustomerCheckoutURL200Response) SetExternalCustomerId(v string)`

SetExternalCustomerId sets ExternalCustomerId field to given value.

### HasExternalCustomerId

`func (o *GenerateCustomerCheckoutURL200Response) HasExternalCustomerId() bool`

HasExternalCustomerId returns a boolean if a field has been set.

### GetPaymentProvider

`func (o *GenerateCustomerCheckoutURL200Response) GetPaymentProvider() string`

GetPaymentProvider returns the PaymentProvider field if non-nil, zero value otherwise.

### GetPaymentProviderOk

`func (o *GenerateCustomerCheckoutURL200Response) GetPaymentProviderOk() (*string, bool)`

GetPaymentProviderOk returns a tuple with the PaymentProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentProvider

`func (o *GenerateCustomerCheckoutURL200Response) SetPaymentProvider(v string)`

SetPaymentProvider sets PaymentProvider field to given value.

### HasPaymentProvider

`func (o *GenerateCustomerCheckoutURL200Response) HasPaymentProvider() bool`

HasPaymentProvider returns a boolean if a field has been set.

### GetCheckoutUrl

`func (o *GenerateCustomerCheckoutURL200Response) GetCheckoutUrl() string`

GetCheckoutUrl returns the CheckoutUrl field if non-nil, zero value otherwise.

### GetCheckoutUrlOk

`func (o *GenerateCustomerCheckoutURL200Response) GetCheckoutUrlOk() (*string, bool)`

GetCheckoutUrlOk returns a tuple with the CheckoutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutUrl

`func (o *GenerateCustomerCheckoutURL200Response) SetCheckoutUrl(v string)`

SetCheckoutUrl sets CheckoutUrl field to given value.

### HasCheckoutUrl

`func (o *GenerateCustomerCheckoutURL200Response) HasCheckoutUrl() bool`

HasCheckoutUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


