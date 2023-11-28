# EventObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LagoId** | **string** | Unique identifier assigned to the event within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the event&#39;s record within the Lago system | 
**TransactionId** | **string** | This field represents a unique identifier for the event. It is crucial for ensuring idempotency, meaning that each event can be uniquely identified and processed without causing any unintended side effects. | 
**LagoCustomerId** | **NullableString** | Unique identifier assigned to the customer within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the customer&#39;s record within the Lago system | 
**ExternalCustomerId** | **NullableString** | The customer external unique identifier (provided by your own application). This field is optional if you send the &#x60;external_subscription_id&#x60;, targeting a specific subscription. | 
**Code** | **string** | The code that identifies a targeted billable metric. It is essential that this code matches the &#x60;code&#x60; property of one of your active billable metrics. If the provided code does not correspond to any active billable metric, it will be ignored during the process. | 
**Timestamp** | **time.Time** | This field captures the Unix timestamp in seconds indicating the occurrence of the event in Coordinated Universal Time (UTC). If this timestamp is not provided, the API will automatically set it to the time of event reception. | 
**Properties** | Pointer to [**EventObjectProperties**](EventObjectProperties.md) |  | [optional] 
**LagoSubscriptionId** | **NullableString** | Unique identifier assigned to the subscription within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the subscription’s record within the Lago system | 
**ExternalSubscriptionId** | **NullableString** | The unique identifier of the subscription within your application. It is a mandatory field when the customer possesses multiple subscriptions or when the &#x60;external_customer_id&#x60; is not provided. | 
**CreatedAt** | **time.Time** | The creation date of the event&#39;s record in the Lago application, presented in the ISO 8601 datetime format, specifically in Coordinated Universal Time (UTC). It provides the precise timestamp of when the event&#39;s record was created within the Lago application | 

## Methods

### NewEventObject

`func NewEventObject(lagoId string, transactionId string, lagoCustomerId NullableString, externalCustomerId NullableString, code string, timestamp time.Time, lagoSubscriptionId NullableString, externalSubscriptionId NullableString, createdAt time.Time, ) *EventObject`

NewEventObject instantiates a new EventObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventObjectWithDefaults

`func NewEventObjectWithDefaults() *EventObject`

NewEventObjectWithDefaults instantiates a new EventObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLagoId

`func (o *EventObject) GetLagoId() string`

GetLagoId returns the LagoId field if non-nil, zero value otherwise.

### GetLagoIdOk

`func (o *EventObject) GetLagoIdOk() (*string, bool)`

GetLagoIdOk returns a tuple with the LagoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagoId

`func (o *EventObject) SetLagoId(v string)`

SetLagoId sets LagoId field to given value.


### GetTransactionId

`func (o *EventObject) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *EventObject) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *EventObject) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetLagoCustomerId

`func (o *EventObject) GetLagoCustomerId() string`

GetLagoCustomerId returns the LagoCustomerId field if non-nil, zero value otherwise.

### GetLagoCustomerIdOk

`func (o *EventObject) GetLagoCustomerIdOk() (*string, bool)`

GetLagoCustomerIdOk returns a tuple with the LagoCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagoCustomerId

`func (o *EventObject) SetLagoCustomerId(v string)`

SetLagoCustomerId sets LagoCustomerId field to given value.


### SetLagoCustomerIdNil

`func (o *EventObject) SetLagoCustomerIdNil(b bool)`

 SetLagoCustomerIdNil sets the value for LagoCustomerId to be an explicit nil

### UnsetLagoCustomerId
`func (o *EventObject) UnsetLagoCustomerId()`

UnsetLagoCustomerId ensures that no value is present for LagoCustomerId, not even an explicit nil
### GetExternalCustomerId

`func (o *EventObject) GetExternalCustomerId() string`

GetExternalCustomerId returns the ExternalCustomerId field if non-nil, zero value otherwise.

### GetExternalCustomerIdOk

`func (o *EventObject) GetExternalCustomerIdOk() (*string, bool)`

GetExternalCustomerIdOk returns a tuple with the ExternalCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalCustomerId

`func (o *EventObject) SetExternalCustomerId(v string)`

SetExternalCustomerId sets ExternalCustomerId field to given value.


### SetExternalCustomerIdNil

`func (o *EventObject) SetExternalCustomerIdNil(b bool)`

 SetExternalCustomerIdNil sets the value for ExternalCustomerId to be an explicit nil

### UnsetExternalCustomerId
`func (o *EventObject) UnsetExternalCustomerId()`

UnsetExternalCustomerId ensures that no value is present for ExternalCustomerId, not even an explicit nil
### GetCode

`func (o *EventObject) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *EventObject) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *EventObject) SetCode(v string)`

SetCode sets Code field to given value.


### GetTimestamp

`func (o *EventObject) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *EventObject) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *EventObject) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetProperties

`func (o *EventObject) GetProperties() EventObjectProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *EventObject) GetPropertiesOk() (*EventObjectProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *EventObject) SetProperties(v EventObjectProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *EventObject) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetLagoSubscriptionId

`func (o *EventObject) GetLagoSubscriptionId() string`

GetLagoSubscriptionId returns the LagoSubscriptionId field if non-nil, zero value otherwise.

### GetLagoSubscriptionIdOk

`func (o *EventObject) GetLagoSubscriptionIdOk() (*string, bool)`

GetLagoSubscriptionIdOk returns a tuple with the LagoSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagoSubscriptionId

`func (o *EventObject) SetLagoSubscriptionId(v string)`

SetLagoSubscriptionId sets LagoSubscriptionId field to given value.


### SetLagoSubscriptionIdNil

`func (o *EventObject) SetLagoSubscriptionIdNil(b bool)`

 SetLagoSubscriptionIdNil sets the value for LagoSubscriptionId to be an explicit nil

### UnsetLagoSubscriptionId
`func (o *EventObject) UnsetLagoSubscriptionId()`

UnsetLagoSubscriptionId ensures that no value is present for LagoSubscriptionId, not even an explicit nil
### GetExternalSubscriptionId

`func (o *EventObject) GetExternalSubscriptionId() string`

GetExternalSubscriptionId returns the ExternalSubscriptionId field if non-nil, zero value otherwise.

### GetExternalSubscriptionIdOk

`func (o *EventObject) GetExternalSubscriptionIdOk() (*string, bool)`

GetExternalSubscriptionIdOk returns a tuple with the ExternalSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSubscriptionId

`func (o *EventObject) SetExternalSubscriptionId(v string)`

SetExternalSubscriptionId sets ExternalSubscriptionId field to given value.


### SetExternalSubscriptionIdNil

`func (o *EventObject) SetExternalSubscriptionIdNil(b bool)`

 SetExternalSubscriptionIdNil sets the value for ExternalSubscriptionId to be an explicit nil

### UnsetExternalSubscriptionId
`func (o *EventObject) UnsetExternalSubscriptionId()`

UnsetExternalSubscriptionId ensures that no value is present for ExternalSubscriptionId, not even an explicit nil
### GetCreatedAt

`func (o *EventObject) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EventObject) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EventObject) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


