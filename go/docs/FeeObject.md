# FeeObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LagoId** | Pointer to **NullableString** | Unique identifier assigned to the fee within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the fee’s record within the Lago system. | [optional] 
**LagoGroupId** | Pointer to **NullableString** | Unique identifier assigned to the group that the fee belongs to | [optional] 
**LagoInvoiceId** | Pointer to **NullableString** | Unique identifier assigned to the invoice that the fee belongs to | [optional] 
**LagoTrueUpFeeId** | Pointer to **NullableString** | Unique identifier assigned to the true-up fee when a minimum has been set to the charge. This identifier helps to distinguish and manage the true-up fee associated with the charge, which may be applicable when a minimum threshold or limit is set for the charge amount. | [optional] 
**LagoTrueUpParentFeeId** | Pointer to **NullableString** | Unique identifier assigned to the parent fee on which the true-up fee is assigned. This identifier establishes the relationship between the parent fee and the associated true-up fee. | [optional] 
**LagoSubscriptionId** | Pointer to **NullableString** | Unique identifier assigned to the subscription, created by Lago. This field is specifically displayed when the fee type is charge or subscription. | [optional] 
**LagoCustomerId** | Pointer to **NullableString** | Unique identifier assigned to the customer, created by Lago. This field is specifically displayed when the fee type is charge or subscription. | [optional] 
**ExternalCustomerId** | Pointer to **NullableString** | Unique identifier assigned to the customer in your application. This field is specifically displayed when the fee type is charge or subscription. | [optional] 
**ExternalSubscriptionId** | Pointer to **NullableString** | Unique identifier assigned to the subscription in your application. This field is specifically displayed when the fee type is charge or subscription. | [optional] 
**InvoiceDisplayName** | Pointer to **string** | Specifies the name that will be displayed on an invoice. If no value is set for this field, the name of the actual charge will be used as the default display name. | [optional] 
**AmountCents** | **int32** | The cost of this specific fee, excluding any applicable taxes. | 
**AmountCurrency** | [**Currency**](Currency.md) |  | 
**TaxesAmountCents** | **int32** | The cost of the tax associated with this specific fee. | 
**TaxesRate** | **float32** | The tax rate associated with this specific fee. | 
**Units** | **string** | The number of units used to charge the customer. This field indicates the quantity or count of units consumed or utilized in the context of the charge. It helps in determining the basis for calculating the fee or cost associated with the usage of the service or product provided to the customer. | 
**PreciseUnitAmount** | **string** | The unit amount of the fee per unit, with precision. | 
**TotalAmountCents** | **int32** | The cost of this specific fee, including any applicable taxes. | 
**TotalAmountCurrency** | [**Currency**](Currency.md) |  | 
**EventsCount** | Pointer to **int32** | The number of events that have been sent and used to charge the customer. This field indicates the count or quantity of events that have been processed and considered in the charging process. | [optional] 
**PayInAdvance** | **bool** | Flag that indicates whether the fee was paid in advance. It serves as a boolean value, where &#x60;true&#x60; represents that the fee was paid in advance (straightaway), and &#x60;false&#x60; indicates that the fee was not paid in arrears (at the end of the period). | 
**Invoiceable** | **bool** | Flag that indicates whether the fee was included on the invoice. It serves as a boolean value, where &#x60;true&#x60; represents that the fee was included on the invoice, and &#x60;false&#x60; indicates that the fee was not included on the invoice. | 
**FromDate** | Pointer to **NullableTime** | The beginning date of the period that the fee covers. It is applicable only to &#x60;subscription&#x60; and &#x60;charge&#x60; fees. This field indicates the start date of the billing period or subscription period associated with the fee. | [optional] 
**ToDate** | Pointer to **NullableTime** | The ending date of the period that the fee covers. It is applicable only to &#x60;subscription&#x60; and &#x60;charge&#x60; fees. This field indicates the end date of the billing period or subscription period associated with the fee. | [optional] 
**PaymentStatus** | **string** | Indicates the payment status of the fee. It represents the current status of the payment associated with the fee. The possible values for this field are &#x60;pending&#x60;, &#x60;succeeded&#x60;, &#x60;failed&#x60; and &#x60;refunded&#x60;. | 
**CreatedAt** | Pointer to **NullableTime** | The date and time when the fee was created. It is provided in Coordinated Universal Time (UTC) format. | [optional] 
**SucceededAt** | Pointer to **NullableTime** | The date and time when the payment for the fee was successfully processed. It is provided in Coordinated Universal Time (UTC) format. | [optional] 
**FailedAt** | Pointer to **NullableTime** | The date and time when the payment for the fee failed to process. It is provided in Coordinated Universal Time (UTC) format. | [optional] 
**RefundedAt** | Pointer to **NullableTime** | The date and time when the payment for the fee was refunded. It is provided in Coordinated Universal Time (UTC) format | [optional] 
**EventTransactionId** | Pointer to **NullableString** | Unique identifier assigned to the transaction. This field is specifically displayed when the fee type is &#x60;charge&#x60; and the payment for the fee is made in advance (&#x60;pay_in_advance&#x60; is set to &#x60;true&#x60;). | [optional] 
**Item** | [**FeeObjectItem**](FeeObjectItem.md) |  | 
**AppliedTaxes** | Pointer to [**[]FeeAppliedTaxObject**](FeeAppliedTaxObject.md) | List of fee applied taxes | [optional] 

## Methods

### NewFeeObject

`func NewFeeObject(amountCents int32, amountCurrency Currency, taxesAmountCents int32, taxesRate float32, units string, preciseUnitAmount string, totalAmountCents int32, totalAmountCurrency Currency, payInAdvance bool, invoiceable bool, paymentStatus string, item FeeObjectItem, ) *FeeObject`

NewFeeObject instantiates a new FeeObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeeObjectWithDefaults

`func NewFeeObjectWithDefaults() *FeeObject`

NewFeeObjectWithDefaults instantiates a new FeeObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLagoId

`func (o *FeeObject) GetLagoId() string`

GetLagoId returns the LagoId field if non-nil, zero value otherwise.

### GetLagoIdOk

`func (o *FeeObject) GetLagoIdOk() (*string, bool)`

GetLagoIdOk returns a tuple with the LagoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagoId

`func (o *FeeObject) SetLagoId(v string)`

SetLagoId sets LagoId field to given value.

### HasLagoId

`func (o *FeeObject) HasLagoId() bool`

HasLagoId returns a boolean if a field has been set.

### SetLagoIdNil

`func (o *FeeObject) SetLagoIdNil(b bool)`

 SetLagoIdNil sets the value for LagoId to be an explicit nil

### UnsetLagoId
`func (o *FeeObject) UnsetLagoId()`

UnsetLagoId ensures that no value is present for LagoId, not even an explicit nil
### GetLagoGroupId

`func (o *FeeObject) GetLagoGroupId() string`

GetLagoGroupId returns the LagoGroupId field if non-nil, zero value otherwise.

### GetLagoGroupIdOk

`func (o *FeeObject) GetLagoGroupIdOk() (*string, bool)`

GetLagoGroupIdOk returns a tuple with the LagoGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagoGroupId

`func (o *FeeObject) SetLagoGroupId(v string)`

SetLagoGroupId sets LagoGroupId field to given value.

### HasLagoGroupId

`func (o *FeeObject) HasLagoGroupId() bool`

HasLagoGroupId returns a boolean if a field has been set.

### SetLagoGroupIdNil

`func (o *FeeObject) SetLagoGroupIdNil(b bool)`

 SetLagoGroupIdNil sets the value for LagoGroupId to be an explicit nil

### UnsetLagoGroupId
`func (o *FeeObject) UnsetLagoGroupId()`

UnsetLagoGroupId ensures that no value is present for LagoGroupId, not even an explicit nil
### GetLagoInvoiceId

`func (o *FeeObject) GetLagoInvoiceId() string`

GetLagoInvoiceId returns the LagoInvoiceId field if non-nil, zero value otherwise.

### GetLagoInvoiceIdOk

`func (o *FeeObject) GetLagoInvoiceIdOk() (*string, bool)`

GetLagoInvoiceIdOk returns a tuple with the LagoInvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagoInvoiceId

`func (o *FeeObject) SetLagoInvoiceId(v string)`

SetLagoInvoiceId sets LagoInvoiceId field to given value.

### HasLagoInvoiceId

`func (o *FeeObject) HasLagoInvoiceId() bool`

HasLagoInvoiceId returns a boolean if a field has been set.

### SetLagoInvoiceIdNil

`func (o *FeeObject) SetLagoInvoiceIdNil(b bool)`

 SetLagoInvoiceIdNil sets the value for LagoInvoiceId to be an explicit nil

### UnsetLagoInvoiceId
`func (o *FeeObject) UnsetLagoInvoiceId()`

UnsetLagoInvoiceId ensures that no value is present for LagoInvoiceId, not even an explicit nil
### GetLagoTrueUpFeeId

`func (o *FeeObject) GetLagoTrueUpFeeId() string`

GetLagoTrueUpFeeId returns the LagoTrueUpFeeId field if non-nil, zero value otherwise.

### GetLagoTrueUpFeeIdOk

`func (o *FeeObject) GetLagoTrueUpFeeIdOk() (*string, bool)`

GetLagoTrueUpFeeIdOk returns a tuple with the LagoTrueUpFeeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagoTrueUpFeeId

`func (o *FeeObject) SetLagoTrueUpFeeId(v string)`

SetLagoTrueUpFeeId sets LagoTrueUpFeeId field to given value.

### HasLagoTrueUpFeeId

`func (o *FeeObject) HasLagoTrueUpFeeId() bool`

HasLagoTrueUpFeeId returns a boolean if a field has been set.

### SetLagoTrueUpFeeIdNil

`func (o *FeeObject) SetLagoTrueUpFeeIdNil(b bool)`

 SetLagoTrueUpFeeIdNil sets the value for LagoTrueUpFeeId to be an explicit nil

### UnsetLagoTrueUpFeeId
`func (o *FeeObject) UnsetLagoTrueUpFeeId()`

UnsetLagoTrueUpFeeId ensures that no value is present for LagoTrueUpFeeId, not even an explicit nil
### GetLagoTrueUpParentFeeId

`func (o *FeeObject) GetLagoTrueUpParentFeeId() string`

GetLagoTrueUpParentFeeId returns the LagoTrueUpParentFeeId field if non-nil, zero value otherwise.

### GetLagoTrueUpParentFeeIdOk

`func (o *FeeObject) GetLagoTrueUpParentFeeIdOk() (*string, bool)`

GetLagoTrueUpParentFeeIdOk returns a tuple with the LagoTrueUpParentFeeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagoTrueUpParentFeeId

`func (o *FeeObject) SetLagoTrueUpParentFeeId(v string)`

SetLagoTrueUpParentFeeId sets LagoTrueUpParentFeeId field to given value.

### HasLagoTrueUpParentFeeId

`func (o *FeeObject) HasLagoTrueUpParentFeeId() bool`

HasLagoTrueUpParentFeeId returns a boolean if a field has been set.

### SetLagoTrueUpParentFeeIdNil

`func (o *FeeObject) SetLagoTrueUpParentFeeIdNil(b bool)`

 SetLagoTrueUpParentFeeIdNil sets the value for LagoTrueUpParentFeeId to be an explicit nil

### UnsetLagoTrueUpParentFeeId
`func (o *FeeObject) UnsetLagoTrueUpParentFeeId()`

UnsetLagoTrueUpParentFeeId ensures that no value is present for LagoTrueUpParentFeeId, not even an explicit nil
### GetLagoSubscriptionId

`func (o *FeeObject) GetLagoSubscriptionId() string`

GetLagoSubscriptionId returns the LagoSubscriptionId field if non-nil, zero value otherwise.

### GetLagoSubscriptionIdOk

`func (o *FeeObject) GetLagoSubscriptionIdOk() (*string, bool)`

GetLagoSubscriptionIdOk returns a tuple with the LagoSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagoSubscriptionId

`func (o *FeeObject) SetLagoSubscriptionId(v string)`

SetLagoSubscriptionId sets LagoSubscriptionId field to given value.

### HasLagoSubscriptionId

`func (o *FeeObject) HasLagoSubscriptionId() bool`

HasLagoSubscriptionId returns a boolean if a field has been set.

### SetLagoSubscriptionIdNil

`func (o *FeeObject) SetLagoSubscriptionIdNil(b bool)`

 SetLagoSubscriptionIdNil sets the value for LagoSubscriptionId to be an explicit nil

### UnsetLagoSubscriptionId
`func (o *FeeObject) UnsetLagoSubscriptionId()`

UnsetLagoSubscriptionId ensures that no value is present for LagoSubscriptionId, not even an explicit nil
### GetLagoCustomerId

`func (o *FeeObject) GetLagoCustomerId() string`

GetLagoCustomerId returns the LagoCustomerId field if non-nil, zero value otherwise.

### GetLagoCustomerIdOk

`func (o *FeeObject) GetLagoCustomerIdOk() (*string, bool)`

GetLagoCustomerIdOk returns a tuple with the LagoCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagoCustomerId

`func (o *FeeObject) SetLagoCustomerId(v string)`

SetLagoCustomerId sets LagoCustomerId field to given value.

### HasLagoCustomerId

`func (o *FeeObject) HasLagoCustomerId() bool`

HasLagoCustomerId returns a boolean if a field has been set.

### SetLagoCustomerIdNil

`func (o *FeeObject) SetLagoCustomerIdNil(b bool)`

 SetLagoCustomerIdNil sets the value for LagoCustomerId to be an explicit nil

### UnsetLagoCustomerId
`func (o *FeeObject) UnsetLagoCustomerId()`

UnsetLagoCustomerId ensures that no value is present for LagoCustomerId, not even an explicit nil
### GetExternalCustomerId

`func (o *FeeObject) GetExternalCustomerId() string`

GetExternalCustomerId returns the ExternalCustomerId field if non-nil, zero value otherwise.

### GetExternalCustomerIdOk

`func (o *FeeObject) GetExternalCustomerIdOk() (*string, bool)`

GetExternalCustomerIdOk returns a tuple with the ExternalCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalCustomerId

`func (o *FeeObject) SetExternalCustomerId(v string)`

SetExternalCustomerId sets ExternalCustomerId field to given value.

### HasExternalCustomerId

`func (o *FeeObject) HasExternalCustomerId() bool`

HasExternalCustomerId returns a boolean if a field has been set.

### SetExternalCustomerIdNil

`func (o *FeeObject) SetExternalCustomerIdNil(b bool)`

 SetExternalCustomerIdNil sets the value for ExternalCustomerId to be an explicit nil

### UnsetExternalCustomerId
`func (o *FeeObject) UnsetExternalCustomerId()`

UnsetExternalCustomerId ensures that no value is present for ExternalCustomerId, not even an explicit nil
### GetExternalSubscriptionId

`func (o *FeeObject) GetExternalSubscriptionId() string`

GetExternalSubscriptionId returns the ExternalSubscriptionId field if non-nil, zero value otherwise.

### GetExternalSubscriptionIdOk

`func (o *FeeObject) GetExternalSubscriptionIdOk() (*string, bool)`

GetExternalSubscriptionIdOk returns a tuple with the ExternalSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSubscriptionId

`func (o *FeeObject) SetExternalSubscriptionId(v string)`

SetExternalSubscriptionId sets ExternalSubscriptionId field to given value.

### HasExternalSubscriptionId

`func (o *FeeObject) HasExternalSubscriptionId() bool`

HasExternalSubscriptionId returns a boolean if a field has been set.

### SetExternalSubscriptionIdNil

`func (o *FeeObject) SetExternalSubscriptionIdNil(b bool)`

 SetExternalSubscriptionIdNil sets the value for ExternalSubscriptionId to be an explicit nil

### UnsetExternalSubscriptionId
`func (o *FeeObject) UnsetExternalSubscriptionId()`

UnsetExternalSubscriptionId ensures that no value is present for ExternalSubscriptionId, not even an explicit nil
### GetInvoiceDisplayName

`func (o *FeeObject) GetInvoiceDisplayName() string`

GetInvoiceDisplayName returns the InvoiceDisplayName field if non-nil, zero value otherwise.

### GetInvoiceDisplayNameOk

`func (o *FeeObject) GetInvoiceDisplayNameOk() (*string, bool)`

GetInvoiceDisplayNameOk returns a tuple with the InvoiceDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceDisplayName

`func (o *FeeObject) SetInvoiceDisplayName(v string)`

SetInvoiceDisplayName sets InvoiceDisplayName field to given value.

### HasInvoiceDisplayName

`func (o *FeeObject) HasInvoiceDisplayName() bool`

HasInvoiceDisplayName returns a boolean if a field has been set.

### GetAmountCents

`func (o *FeeObject) GetAmountCents() int32`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *FeeObject) GetAmountCentsOk() (*int32, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *FeeObject) SetAmountCents(v int32)`

SetAmountCents sets AmountCents field to given value.


### GetAmountCurrency

`func (o *FeeObject) GetAmountCurrency() Currency`

GetAmountCurrency returns the AmountCurrency field if non-nil, zero value otherwise.

### GetAmountCurrencyOk

`func (o *FeeObject) GetAmountCurrencyOk() (*Currency, bool)`

GetAmountCurrencyOk returns a tuple with the AmountCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCurrency

`func (o *FeeObject) SetAmountCurrency(v Currency)`

SetAmountCurrency sets AmountCurrency field to given value.


### GetTaxesAmountCents

`func (o *FeeObject) GetTaxesAmountCents() int32`

GetTaxesAmountCents returns the TaxesAmountCents field if non-nil, zero value otherwise.

### GetTaxesAmountCentsOk

`func (o *FeeObject) GetTaxesAmountCentsOk() (*int32, bool)`

GetTaxesAmountCentsOk returns a tuple with the TaxesAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxesAmountCents

`func (o *FeeObject) SetTaxesAmountCents(v int32)`

SetTaxesAmountCents sets TaxesAmountCents field to given value.


### GetTaxesRate

`func (o *FeeObject) GetTaxesRate() float32`

GetTaxesRate returns the TaxesRate field if non-nil, zero value otherwise.

### GetTaxesRateOk

`func (o *FeeObject) GetTaxesRateOk() (*float32, bool)`

GetTaxesRateOk returns a tuple with the TaxesRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxesRate

`func (o *FeeObject) SetTaxesRate(v float32)`

SetTaxesRate sets TaxesRate field to given value.


### GetUnits

`func (o *FeeObject) GetUnits() string`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *FeeObject) GetUnitsOk() (*string, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *FeeObject) SetUnits(v string)`

SetUnits sets Units field to given value.


### GetPreciseUnitAmount

`func (o *FeeObject) GetPreciseUnitAmount() string`

GetPreciseUnitAmount returns the PreciseUnitAmount field if non-nil, zero value otherwise.

### GetPreciseUnitAmountOk

`func (o *FeeObject) GetPreciseUnitAmountOk() (*string, bool)`

GetPreciseUnitAmountOk returns a tuple with the PreciseUnitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreciseUnitAmount

`func (o *FeeObject) SetPreciseUnitAmount(v string)`

SetPreciseUnitAmount sets PreciseUnitAmount field to given value.


### GetTotalAmountCents

`func (o *FeeObject) GetTotalAmountCents() int32`

GetTotalAmountCents returns the TotalAmountCents field if non-nil, zero value otherwise.

### GetTotalAmountCentsOk

`func (o *FeeObject) GetTotalAmountCentsOk() (*int32, bool)`

GetTotalAmountCentsOk returns a tuple with the TotalAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountCents

`func (o *FeeObject) SetTotalAmountCents(v int32)`

SetTotalAmountCents sets TotalAmountCents field to given value.


### GetTotalAmountCurrency

`func (o *FeeObject) GetTotalAmountCurrency() Currency`

GetTotalAmountCurrency returns the TotalAmountCurrency field if non-nil, zero value otherwise.

### GetTotalAmountCurrencyOk

`func (o *FeeObject) GetTotalAmountCurrencyOk() (*Currency, bool)`

GetTotalAmountCurrencyOk returns a tuple with the TotalAmountCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountCurrency

`func (o *FeeObject) SetTotalAmountCurrency(v Currency)`

SetTotalAmountCurrency sets TotalAmountCurrency field to given value.


### GetEventsCount

`func (o *FeeObject) GetEventsCount() int32`

GetEventsCount returns the EventsCount field if non-nil, zero value otherwise.

### GetEventsCountOk

`func (o *FeeObject) GetEventsCountOk() (*int32, bool)`

GetEventsCountOk returns a tuple with the EventsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsCount

`func (o *FeeObject) SetEventsCount(v int32)`

SetEventsCount sets EventsCount field to given value.

### HasEventsCount

`func (o *FeeObject) HasEventsCount() bool`

HasEventsCount returns a boolean if a field has been set.

### GetPayInAdvance

`func (o *FeeObject) GetPayInAdvance() bool`

GetPayInAdvance returns the PayInAdvance field if non-nil, zero value otherwise.

### GetPayInAdvanceOk

`func (o *FeeObject) GetPayInAdvanceOk() (*bool, bool)`

GetPayInAdvanceOk returns a tuple with the PayInAdvance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayInAdvance

`func (o *FeeObject) SetPayInAdvance(v bool)`

SetPayInAdvance sets PayInAdvance field to given value.


### GetInvoiceable

`func (o *FeeObject) GetInvoiceable() bool`

GetInvoiceable returns the Invoiceable field if non-nil, zero value otherwise.

### GetInvoiceableOk

`func (o *FeeObject) GetInvoiceableOk() (*bool, bool)`

GetInvoiceableOk returns a tuple with the Invoiceable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceable

`func (o *FeeObject) SetInvoiceable(v bool)`

SetInvoiceable sets Invoiceable field to given value.


### GetFromDate

`func (o *FeeObject) GetFromDate() time.Time`

GetFromDate returns the FromDate field if non-nil, zero value otherwise.

### GetFromDateOk

`func (o *FeeObject) GetFromDateOk() (*time.Time, bool)`

GetFromDateOk returns a tuple with the FromDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDate

`func (o *FeeObject) SetFromDate(v time.Time)`

SetFromDate sets FromDate field to given value.

### HasFromDate

`func (o *FeeObject) HasFromDate() bool`

HasFromDate returns a boolean if a field has been set.

### SetFromDateNil

`func (o *FeeObject) SetFromDateNil(b bool)`

 SetFromDateNil sets the value for FromDate to be an explicit nil

### UnsetFromDate
`func (o *FeeObject) UnsetFromDate()`

UnsetFromDate ensures that no value is present for FromDate, not even an explicit nil
### GetToDate

`func (o *FeeObject) GetToDate() time.Time`

GetToDate returns the ToDate field if non-nil, zero value otherwise.

### GetToDateOk

`func (o *FeeObject) GetToDateOk() (*time.Time, bool)`

GetToDateOk returns a tuple with the ToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToDate

`func (o *FeeObject) SetToDate(v time.Time)`

SetToDate sets ToDate field to given value.

### HasToDate

`func (o *FeeObject) HasToDate() bool`

HasToDate returns a boolean if a field has been set.

### SetToDateNil

`func (o *FeeObject) SetToDateNil(b bool)`

 SetToDateNil sets the value for ToDate to be an explicit nil

### UnsetToDate
`func (o *FeeObject) UnsetToDate()`

UnsetToDate ensures that no value is present for ToDate, not even an explicit nil
### GetPaymentStatus

`func (o *FeeObject) GetPaymentStatus() string`

GetPaymentStatus returns the PaymentStatus field if non-nil, zero value otherwise.

### GetPaymentStatusOk

`func (o *FeeObject) GetPaymentStatusOk() (*string, bool)`

GetPaymentStatusOk returns a tuple with the PaymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentStatus

`func (o *FeeObject) SetPaymentStatus(v string)`

SetPaymentStatus sets PaymentStatus field to given value.


### GetCreatedAt

`func (o *FeeObject) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FeeObject) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FeeObject) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FeeObject) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *FeeObject) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *FeeObject) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetSucceededAt

`func (o *FeeObject) GetSucceededAt() time.Time`

GetSucceededAt returns the SucceededAt field if non-nil, zero value otherwise.

### GetSucceededAtOk

`func (o *FeeObject) GetSucceededAtOk() (*time.Time, bool)`

GetSucceededAtOk returns a tuple with the SucceededAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSucceededAt

`func (o *FeeObject) SetSucceededAt(v time.Time)`

SetSucceededAt sets SucceededAt field to given value.

### HasSucceededAt

`func (o *FeeObject) HasSucceededAt() bool`

HasSucceededAt returns a boolean if a field has been set.

### SetSucceededAtNil

`func (o *FeeObject) SetSucceededAtNil(b bool)`

 SetSucceededAtNil sets the value for SucceededAt to be an explicit nil

### UnsetSucceededAt
`func (o *FeeObject) UnsetSucceededAt()`

UnsetSucceededAt ensures that no value is present for SucceededAt, not even an explicit nil
### GetFailedAt

`func (o *FeeObject) GetFailedAt() time.Time`

GetFailedAt returns the FailedAt field if non-nil, zero value otherwise.

### GetFailedAtOk

`func (o *FeeObject) GetFailedAtOk() (*time.Time, bool)`

GetFailedAtOk returns a tuple with the FailedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedAt

`func (o *FeeObject) SetFailedAt(v time.Time)`

SetFailedAt sets FailedAt field to given value.

### HasFailedAt

`func (o *FeeObject) HasFailedAt() bool`

HasFailedAt returns a boolean if a field has been set.

### SetFailedAtNil

`func (o *FeeObject) SetFailedAtNil(b bool)`

 SetFailedAtNil sets the value for FailedAt to be an explicit nil

### UnsetFailedAt
`func (o *FeeObject) UnsetFailedAt()`

UnsetFailedAt ensures that no value is present for FailedAt, not even an explicit nil
### GetRefundedAt

`func (o *FeeObject) GetRefundedAt() time.Time`

GetRefundedAt returns the RefundedAt field if non-nil, zero value otherwise.

### GetRefundedAtOk

`func (o *FeeObject) GetRefundedAtOk() (*time.Time, bool)`

GetRefundedAtOk returns a tuple with the RefundedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundedAt

`func (o *FeeObject) SetRefundedAt(v time.Time)`

SetRefundedAt sets RefundedAt field to given value.

### HasRefundedAt

`func (o *FeeObject) HasRefundedAt() bool`

HasRefundedAt returns a boolean if a field has been set.

### SetRefundedAtNil

`func (o *FeeObject) SetRefundedAtNil(b bool)`

 SetRefundedAtNil sets the value for RefundedAt to be an explicit nil

### UnsetRefundedAt
`func (o *FeeObject) UnsetRefundedAt()`

UnsetRefundedAt ensures that no value is present for RefundedAt, not even an explicit nil
### GetEventTransactionId

`func (o *FeeObject) GetEventTransactionId() string`

GetEventTransactionId returns the EventTransactionId field if non-nil, zero value otherwise.

### GetEventTransactionIdOk

`func (o *FeeObject) GetEventTransactionIdOk() (*string, bool)`

GetEventTransactionIdOk returns a tuple with the EventTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTransactionId

`func (o *FeeObject) SetEventTransactionId(v string)`

SetEventTransactionId sets EventTransactionId field to given value.

### HasEventTransactionId

`func (o *FeeObject) HasEventTransactionId() bool`

HasEventTransactionId returns a boolean if a field has been set.

### SetEventTransactionIdNil

`func (o *FeeObject) SetEventTransactionIdNil(b bool)`

 SetEventTransactionIdNil sets the value for EventTransactionId to be an explicit nil

### UnsetEventTransactionId
`func (o *FeeObject) UnsetEventTransactionId()`

UnsetEventTransactionId ensures that no value is present for EventTransactionId, not even an explicit nil
### GetItem

`func (o *FeeObject) GetItem() FeeObjectItem`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *FeeObject) GetItemOk() (*FeeObjectItem, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *FeeObject) SetItem(v FeeObjectItem)`

SetItem sets Item field to given value.


### GetAppliedTaxes

`func (o *FeeObject) GetAppliedTaxes() []FeeAppliedTaxObject`

GetAppliedTaxes returns the AppliedTaxes field if non-nil, zero value otherwise.

### GetAppliedTaxesOk

`func (o *FeeObject) GetAppliedTaxesOk() (*[]FeeAppliedTaxObject, bool)`

GetAppliedTaxesOk returns a tuple with the AppliedTaxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedTaxes

`func (o *FeeObject) SetAppliedTaxes(v []FeeAppliedTaxObject)`

SetAppliedTaxes sets AppliedTaxes field to given value.

### HasAppliedTaxes

`func (o *FeeObject) HasAppliedTaxes() bool`

HasAppliedTaxes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


