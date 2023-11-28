# CreditNoteItemObjectFee

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

### NewCreditNoteItemObjectFee

`func NewCreditNoteItemObjectFee(amountCents int32, amountCurrency Currency, taxesAmountCents int32, taxesRate float32, units string, preciseUnitAmount string, totalAmountCents int32, totalAmountCurrency Currency, payInAdvance bool, invoiceable bool, paymentStatus string, item FeeObjectItem, ) *CreditNoteItemObjectFee`

NewCreditNoteItemObjectFee instantiates a new CreditNoteItemObjectFee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditNoteItemObjectFeeWithDefaults

`func NewCreditNoteItemObjectFeeWithDefaults() *CreditNoteItemObjectFee`

NewCreditNoteItemObjectFeeWithDefaults instantiates a new CreditNoteItemObjectFee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLagoId

`func (o *CreditNoteItemObjectFee) GetLagoId() string`

GetLagoId returns the LagoId field if non-nil, zero value otherwise.

### GetLagoIdOk

`func (o *CreditNoteItemObjectFee) GetLagoIdOk() (*string, bool)`

GetLagoIdOk returns a tuple with the LagoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagoId

`func (o *CreditNoteItemObjectFee) SetLagoId(v string)`

SetLagoId sets LagoId field to given value.

### HasLagoId

`func (o *CreditNoteItemObjectFee) HasLagoId() bool`

HasLagoId returns a boolean if a field has been set.

### SetLagoIdNil

`func (o *CreditNoteItemObjectFee) SetLagoIdNil(b bool)`

 SetLagoIdNil sets the value for LagoId to be an explicit nil

### UnsetLagoId
`func (o *CreditNoteItemObjectFee) UnsetLagoId()`

UnsetLagoId ensures that no value is present for LagoId, not even an explicit nil
### GetLagoGroupId

`func (o *CreditNoteItemObjectFee) GetLagoGroupId() string`

GetLagoGroupId returns the LagoGroupId field if non-nil, zero value otherwise.

### GetLagoGroupIdOk

`func (o *CreditNoteItemObjectFee) GetLagoGroupIdOk() (*string, bool)`

GetLagoGroupIdOk returns a tuple with the LagoGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagoGroupId

`func (o *CreditNoteItemObjectFee) SetLagoGroupId(v string)`

SetLagoGroupId sets LagoGroupId field to given value.

### HasLagoGroupId

`func (o *CreditNoteItemObjectFee) HasLagoGroupId() bool`

HasLagoGroupId returns a boolean if a field has been set.

### SetLagoGroupIdNil

`func (o *CreditNoteItemObjectFee) SetLagoGroupIdNil(b bool)`

 SetLagoGroupIdNil sets the value for LagoGroupId to be an explicit nil

### UnsetLagoGroupId
`func (o *CreditNoteItemObjectFee) UnsetLagoGroupId()`

UnsetLagoGroupId ensures that no value is present for LagoGroupId, not even an explicit nil
### GetLagoInvoiceId

`func (o *CreditNoteItemObjectFee) GetLagoInvoiceId() string`

GetLagoInvoiceId returns the LagoInvoiceId field if non-nil, zero value otherwise.

### GetLagoInvoiceIdOk

`func (o *CreditNoteItemObjectFee) GetLagoInvoiceIdOk() (*string, bool)`

GetLagoInvoiceIdOk returns a tuple with the LagoInvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagoInvoiceId

`func (o *CreditNoteItemObjectFee) SetLagoInvoiceId(v string)`

SetLagoInvoiceId sets LagoInvoiceId field to given value.

### HasLagoInvoiceId

`func (o *CreditNoteItemObjectFee) HasLagoInvoiceId() bool`

HasLagoInvoiceId returns a boolean if a field has been set.

### SetLagoInvoiceIdNil

`func (o *CreditNoteItemObjectFee) SetLagoInvoiceIdNil(b bool)`

 SetLagoInvoiceIdNil sets the value for LagoInvoiceId to be an explicit nil

### UnsetLagoInvoiceId
`func (o *CreditNoteItemObjectFee) UnsetLagoInvoiceId()`

UnsetLagoInvoiceId ensures that no value is present for LagoInvoiceId, not even an explicit nil
### GetLagoTrueUpFeeId

`func (o *CreditNoteItemObjectFee) GetLagoTrueUpFeeId() string`

GetLagoTrueUpFeeId returns the LagoTrueUpFeeId field if non-nil, zero value otherwise.

### GetLagoTrueUpFeeIdOk

`func (o *CreditNoteItemObjectFee) GetLagoTrueUpFeeIdOk() (*string, bool)`

GetLagoTrueUpFeeIdOk returns a tuple with the LagoTrueUpFeeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagoTrueUpFeeId

`func (o *CreditNoteItemObjectFee) SetLagoTrueUpFeeId(v string)`

SetLagoTrueUpFeeId sets LagoTrueUpFeeId field to given value.

### HasLagoTrueUpFeeId

`func (o *CreditNoteItemObjectFee) HasLagoTrueUpFeeId() bool`

HasLagoTrueUpFeeId returns a boolean if a field has been set.

### SetLagoTrueUpFeeIdNil

`func (o *CreditNoteItemObjectFee) SetLagoTrueUpFeeIdNil(b bool)`

 SetLagoTrueUpFeeIdNil sets the value for LagoTrueUpFeeId to be an explicit nil

### UnsetLagoTrueUpFeeId
`func (o *CreditNoteItemObjectFee) UnsetLagoTrueUpFeeId()`

UnsetLagoTrueUpFeeId ensures that no value is present for LagoTrueUpFeeId, not even an explicit nil
### GetLagoTrueUpParentFeeId

`func (o *CreditNoteItemObjectFee) GetLagoTrueUpParentFeeId() string`

GetLagoTrueUpParentFeeId returns the LagoTrueUpParentFeeId field if non-nil, zero value otherwise.

### GetLagoTrueUpParentFeeIdOk

`func (o *CreditNoteItemObjectFee) GetLagoTrueUpParentFeeIdOk() (*string, bool)`

GetLagoTrueUpParentFeeIdOk returns a tuple with the LagoTrueUpParentFeeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagoTrueUpParentFeeId

`func (o *CreditNoteItemObjectFee) SetLagoTrueUpParentFeeId(v string)`

SetLagoTrueUpParentFeeId sets LagoTrueUpParentFeeId field to given value.

### HasLagoTrueUpParentFeeId

`func (o *CreditNoteItemObjectFee) HasLagoTrueUpParentFeeId() bool`

HasLagoTrueUpParentFeeId returns a boolean if a field has been set.

### SetLagoTrueUpParentFeeIdNil

`func (o *CreditNoteItemObjectFee) SetLagoTrueUpParentFeeIdNil(b bool)`

 SetLagoTrueUpParentFeeIdNil sets the value for LagoTrueUpParentFeeId to be an explicit nil

### UnsetLagoTrueUpParentFeeId
`func (o *CreditNoteItemObjectFee) UnsetLagoTrueUpParentFeeId()`

UnsetLagoTrueUpParentFeeId ensures that no value is present for LagoTrueUpParentFeeId, not even an explicit nil
### GetLagoSubscriptionId

`func (o *CreditNoteItemObjectFee) GetLagoSubscriptionId() string`

GetLagoSubscriptionId returns the LagoSubscriptionId field if non-nil, zero value otherwise.

### GetLagoSubscriptionIdOk

`func (o *CreditNoteItemObjectFee) GetLagoSubscriptionIdOk() (*string, bool)`

GetLagoSubscriptionIdOk returns a tuple with the LagoSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagoSubscriptionId

`func (o *CreditNoteItemObjectFee) SetLagoSubscriptionId(v string)`

SetLagoSubscriptionId sets LagoSubscriptionId field to given value.

### HasLagoSubscriptionId

`func (o *CreditNoteItemObjectFee) HasLagoSubscriptionId() bool`

HasLagoSubscriptionId returns a boolean if a field has been set.

### SetLagoSubscriptionIdNil

`func (o *CreditNoteItemObjectFee) SetLagoSubscriptionIdNil(b bool)`

 SetLagoSubscriptionIdNil sets the value for LagoSubscriptionId to be an explicit nil

### UnsetLagoSubscriptionId
`func (o *CreditNoteItemObjectFee) UnsetLagoSubscriptionId()`

UnsetLagoSubscriptionId ensures that no value is present for LagoSubscriptionId, not even an explicit nil
### GetLagoCustomerId

`func (o *CreditNoteItemObjectFee) GetLagoCustomerId() string`

GetLagoCustomerId returns the LagoCustomerId field if non-nil, zero value otherwise.

### GetLagoCustomerIdOk

`func (o *CreditNoteItemObjectFee) GetLagoCustomerIdOk() (*string, bool)`

GetLagoCustomerIdOk returns a tuple with the LagoCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagoCustomerId

`func (o *CreditNoteItemObjectFee) SetLagoCustomerId(v string)`

SetLagoCustomerId sets LagoCustomerId field to given value.

### HasLagoCustomerId

`func (o *CreditNoteItemObjectFee) HasLagoCustomerId() bool`

HasLagoCustomerId returns a boolean if a field has been set.

### SetLagoCustomerIdNil

`func (o *CreditNoteItemObjectFee) SetLagoCustomerIdNil(b bool)`

 SetLagoCustomerIdNil sets the value for LagoCustomerId to be an explicit nil

### UnsetLagoCustomerId
`func (o *CreditNoteItemObjectFee) UnsetLagoCustomerId()`

UnsetLagoCustomerId ensures that no value is present for LagoCustomerId, not even an explicit nil
### GetExternalCustomerId

`func (o *CreditNoteItemObjectFee) GetExternalCustomerId() string`

GetExternalCustomerId returns the ExternalCustomerId field if non-nil, zero value otherwise.

### GetExternalCustomerIdOk

`func (o *CreditNoteItemObjectFee) GetExternalCustomerIdOk() (*string, bool)`

GetExternalCustomerIdOk returns a tuple with the ExternalCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalCustomerId

`func (o *CreditNoteItemObjectFee) SetExternalCustomerId(v string)`

SetExternalCustomerId sets ExternalCustomerId field to given value.

### HasExternalCustomerId

`func (o *CreditNoteItemObjectFee) HasExternalCustomerId() bool`

HasExternalCustomerId returns a boolean if a field has been set.

### SetExternalCustomerIdNil

`func (o *CreditNoteItemObjectFee) SetExternalCustomerIdNil(b bool)`

 SetExternalCustomerIdNil sets the value for ExternalCustomerId to be an explicit nil

### UnsetExternalCustomerId
`func (o *CreditNoteItemObjectFee) UnsetExternalCustomerId()`

UnsetExternalCustomerId ensures that no value is present for ExternalCustomerId, not even an explicit nil
### GetExternalSubscriptionId

`func (o *CreditNoteItemObjectFee) GetExternalSubscriptionId() string`

GetExternalSubscriptionId returns the ExternalSubscriptionId field if non-nil, zero value otherwise.

### GetExternalSubscriptionIdOk

`func (o *CreditNoteItemObjectFee) GetExternalSubscriptionIdOk() (*string, bool)`

GetExternalSubscriptionIdOk returns a tuple with the ExternalSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSubscriptionId

`func (o *CreditNoteItemObjectFee) SetExternalSubscriptionId(v string)`

SetExternalSubscriptionId sets ExternalSubscriptionId field to given value.

### HasExternalSubscriptionId

`func (o *CreditNoteItemObjectFee) HasExternalSubscriptionId() bool`

HasExternalSubscriptionId returns a boolean if a field has been set.

### SetExternalSubscriptionIdNil

`func (o *CreditNoteItemObjectFee) SetExternalSubscriptionIdNil(b bool)`

 SetExternalSubscriptionIdNil sets the value for ExternalSubscriptionId to be an explicit nil

### UnsetExternalSubscriptionId
`func (o *CreditNoteItemObjectFee) UnsetExternalSubscriptionId()`

UnsetExternalSubscriptionId ensures that no value is present for ExternalSubscriptionId, not even an explicit nil
### GetInvoiceDisplayName

`func (o *CreditNoteItemObjectFee) GetInvoiceDisplayName() string`

GetInvoiceDisplayName returns the InvoiceDisplayName field if non-nil, zero value otherwise.

### GetInvoiceDisplayNameOk

`func (o *CreditNoteItemObjectFee) GetInvoiceDisplayNameOk() (*string, bool)`

GetInvoiceDisplayNameOk returns a tuple with the InvoiceDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceDisplayName

`func (o *CreditNoteItemObjectFee) SetInvoiceDisplayName(v string)`

SetInvoiceDisplayName sets InvoiceDisplayName field to given value.

### HasInvoiceDisplayName

`func (o *CreditNoteItemObjectFee) HasInvoiceDisplayName() bool`

HasInvoiceDisplayName returns a boolean if a field has been set.

### GetAmountCents

`func (o *CreditNoteItemObjectFee) GetAmountCents() int32`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *CreditNoteItemObjectFee) GetAmountCentsOk() (*int32, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *CreditNoteItemObjectFee) SetAmountCents(v int32)`

SetAmountCents sets AmountCents field to given value.


### GetAmountCurrency

`func (o *CreditNoteItemObjectFee) GetAmountCurrency() Currency`

GetAmountCurrency returns the AmountCurrency field if non-nil, zero value otherwise.

### GetAmountCurrencyOk

`func (o *CreditNoteItemObjectFee) GetAmountCurrencyOk() (*Currency, bool)`

GetAmountCurrencyOk returns a tuple with the AmountCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCurrency

`func (o *CreditNoteItemObjectFee) SetAmountCurrency(v Currency)`

SetAmountCurrency sets AmountCurrency field to given value.


### GetTaxesAmountCents

`func (o *CreditNoteItemObjectFee) GetTaxesAmountCents() int32`

GetTaxesAmountCents returns the TaxesAmountCents field if non-nil, zero value otherwise.

### GetTaxesAmountCentsOk

`func (o *CreditNoteItemObjectFee) GetTaxesAmountCentsOk() (*int32, bool)`

GetTaxesAmountCentsOk returns a tuple with the TaxesAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxesAmountCents

`func (o *CreditNoteItemObjectFee) SetTaxesAmountCents(v int32)`

SetTaxesAmountCents sets TaxesAmountCents field to given value.


### GetTaxesRate

`func (o *CreditNoteItemObjectFee) GetTaxesRate() float32`

GetTaxesRate returns the TaxesRate field if non-nil, zero value otherwise.

### GetTaxesRateOk

`func (o *CreditNoteItemObjectFee) GetTaxesRateOk() (*float32, bool)`

GetTaxesRateOk returns a tuple with the TaxesRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxesRate

`func (o *CreditNoteItemObjectFee) SetTaxesRate(v float32)`

SetTaxesRate sets TaxesRate field to given value.


### GetUnits

`func (o *CreditNoteItemObjectFee) GetUnits() string`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *CreditNoteItemObjectFee) GetUnitsOk() (*string, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *CreditNoteItemObjectFee) SetUnits(v string)`

SetUnits sets Units field to given value.


### GetPreciseUnitAmount

`func (o *CreditNoteItemObjectFee) GetPreciseUnitAmount() string`

GetPreciseUnitAmount returns the PreciseUnitAmount field if non-nil, zero value otherwise.

### GetPreciseUnitAmountOk

`func (o *CreditNoteItemObjectFee) GetPreciseUnitAmountOk() (*string, bool)`

GetPreciseUnitAmountOk returns a tuple with the PreciseUnitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreciseUnitAmount

`func (o *CreditNoteItemObjectFee) SetPreciseUnitAmount(v string)`

SetPreciseUnitAmount sets PreciseUnitAmount field to given value.


### GetTotalAmountCents

`func (o *CreditNoteItemObjectFee) GetTotalAmountCents() int32`

GetTotalAmountCents returns the TotalAmountCents field if non-nil, zero value otherwise.

### GetTotalAmountCentsOk

`func (o *CreditNoteItemObjectFee) GetTotalAmountCentsOk() (*int32, bool)`

GetTotalAmountCentsOk returns a tuple with the TotalAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountCents

`func (o *CreditNoteItemObjectFee) SetTotalAmountCents(v int32)`

SetTotalAmountCents sets TotalAmountCents field to given value.


### GetTotalAmountCurrency

`func (o *CreditNoteItemObjectFee) GetTotalAmountCurrency() Currency`

GetTotalAmountCurrency returns the TotalAmountCurrency field if non-nil, zero value otherwise.

### GetTotalAmountCurrencyOk

`func (o *CreditNoteItemObjectFee) GetTotalAmountCurrencyOk() (*Currency, bool)`

GetTotalAmountCurrencyOk returns a tuple with the TotalAmountCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountCurrency

`func (o *CreditNoteItemObjectFee) SetTotalAmountCurrency(v Currency)`

SetTotalAmountCurrency sets TotalAmountCurrency field to given value.


### GetEventsCount

`func (o *CreditNoteItemObjectFee) GetEventsCount() int32`

GetEventsCount returns the EventsCount field if non-nil, zero value otherwise.

### GetEventsCountOk

`func (o *CreditNoteItemObjectFee) GetEventsCountOk() (*int32, bool)`

GetEventsCountOk returns a tuple with the EventsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsCount

`func (o *CreditNoteItemObjectFee) SetEventsCount(v int32)`

SetEventsCount sets EventsCount field to given value.

### HasEventsCount

`func (o *CreditNoteItemObjectFee) HasEventsCount() bool`

HasEventsCount returns a boolean if a field has been set.

### GetPayInAdvance

`func (o *CreditNoteItemObjectFee) GetPayInAdvance() bool`

GetPayInAdvance returns the PayInAdvance field if non-nil, zero value otherwise.

### GetPayInAdvanceOk

`func (o *CreditNoteItemObjectFee) GetPayInAdvanceOk() (*bool, bool)`

GetPayInAdvanceOk returns a tuple with the PayInAdvance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayInAdvance

`func (o *CreditNoteItemObjectFee) SetPayInAdvance(v bool)`

SetPayInAdvance sets PayInAdvance field to given value.


### GetInvoiceable

`func (o *CreditNoteItemObjectFee) GetInvoiceable() bool`

GetInvoiceable returns the Invoiceable field if non-nil, zero value otherwise.

### GetInvoiceableOk

`func (o *CreditNoteItemObjectFee) GetInvoiceableOk() (*bool, bool)`

GetInvoiceableOk returns a tuple with the Invoiceable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceable

`func (o *CreditNoteItemObjectFee) SetInvoiceable(v bool)`

SetInvoiceable sets Invoiceable field to given value.


### GetFromDate

`func (o *CreditNoteItemObjectFee) GetFromDate() time.Time`

GetFromDate returns the FromDate field if non-nil, zero value otherwise.

### GetFromDateOk

`func (o *CreditNoteItemObjectFee) GetFromDateOk() (*time.Time, bool)`

GetFromDateOk returns a tuple with the FromDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDate

`func (o *CreditNoteItemObjectFee) SetFromDate(v time.Time)`

SetFromDate sets FromDate field to given value.

### HasFromDate

`func (o *CreditNoteItemObjectFee) HasFromDate() bool`

HasFromDate returns a boolean if a field has been set.

### SetFromDateNil

`func (o *CreditNoteItemObjectFee) SetFromDateNil(b bool)`

 SetFromDateNil sets the value for FromDate to be an explicit nil

### UnsetFromDate
`func (o *CreditNoteItemObjectFee) UnsetFromDate()`

UnsetFromDate ensures that no value is present for FromDate, not even an explicit nil
### GetToDate

`func (o *CreditNoteItemObjectFee) GetToDate() time.Time`

GetToDate returns the ToDate field if non-nil, zero value otherwise.

### GetToDateOk

`func (o *CreditNoteItemObjectFee) GetToDateOk() (*time.Time, bool)`

GetToDateOk returns a tuple with the ToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToDate

`func (o *CreditNoteItemObjectFee) SetToDate(v time.Time)`

SetToDate sets ToDate field to given value.

### HasToDate

`func (o *CreditNoteItemObjectFee) HasToDate() bool`

HasToDate returns a boolean if a field has been set.

### SetToDateNil

`func (o *CreditNoteItemObjectFee) SetToDateNil(b bool)`

 SetToDateNil sets the value for ToDate to be an explicit nil

### UnsetToDate
`func (o *CreditNoteItemObjectFee) UnsetToDate()`

UnsetToDate ensures that no value is present for ToDate, not even an explicit nil
### GetPaymentStatus

`func (o *CreditNoteItemObjectFee) GetPaymentStatus() string`

GetPaymentStatus returns the PaymentStatus field if non-nil, zero value otherwise.

### GetPaymentStatusOk

`func (o *CreditNoteItemObjectFee) GetPaymentStatusOk() (*string, bool)`

GetPaymentStatusOk returns a tuple with the PaymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentStatus

`func (o *CreditNoteItemObjectFee) SetPaymentStatus(v string)`

SetPaymentStatus sets PaymentStatus field to given value.


### GetCreatedAt

`func (o *CreditNoteItemObjectFee) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreditNoteItemObjectFee) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreditNoteItemObjectFee) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CreditNoteItemObjectFee) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *CreditNoteItemObjectFee) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *CreditNoteItemObjectFee) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetSucceededAt

`func (o *CreditNoteItemObjectFee) GetSucceededAt() time.Time`

GetSucceededAt returns the SucceededAt field if non-nil, zero value otherwise.

### GetSucceededAtOk

`func (o *CreditNoteItemObjectFee) GetSucceededAtOk() (*time.Time, bool)`

GetSucceededAtOk returns a tuple with the SucceededAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSucceededAt

`func (o *CreditNoteItemObjectFee) SetSucceededAt(v time.Time)`

SetSucceededAt sets SucceededAt field to given value.

### HasSucceededAt

`func (o *CreditNoteItemObjectFee) HasSucceededAt() bool`

HasSucceededAt returns a boolean if a field has been set.

### SetSucceededAtNil

`func (o *CreditNoteItemObjectFee) SetSucceededAtNil(b bool)`

 SetSucceededAtNil sets the value for SucceededAt to be an explicit nil

### UnsetSucceededAt
`func (o *CreditNoteItemObjectFee) UnsetSucceededAt()`

UnsetSucceededAt ensures that no value is present for SucceededAt, not even an explicit nil
### GetFailedAt

`func (o *CreditNoteItemObjectFee) GetFailedAt() time.Time`

GetFailedAt returns the FailedAt field if non-nil, zero value otherwise.

### GetFailedAtOk

`func (o *CreditNoteItemObjectFee) GetFailedAtOk() (*time.Time, bool)`

GetFailedAtOk returns a tuple with the FailedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedAt

`func (o *CreditNoteItemObjectFee) SetFailedAt(v time.Time)`

SetFailedAt sets FailedAt field to given value.

### HasFailedAt

`func (o *CreditNoteItemObjectFee) HasFailedAt() bool`

HasFailedAt returns a boolean if a field has been set.

### SetFailedAtNil

`func (o *CreditNoteItemObjectFee) SetFailedAtNil(b bool)`

 SetFailedAtNil sets the value for FailedAt to be an explicit nil

### UnsetFailedAt
`func (o *CreditNoteItemObjectFee) UnsetFailedAt()`

UnsetFailedAt ensures that no value is present for FailedAt, not even an explicit nil
### GetRefundedAt

`func (o *CreditNoteItemObjectFee) GetRefundedAt() time.Time`

GetRefundedAt returns the RefundedAt field if non-nil, zero value otherwise.

### GetRefundedAtOk

`func (o *CreditNoteItemObjectFee) GetRefundedAtOk() (*time.Time, bool)`

GetRefundedAtOk returns a tuple with the RefundedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundedAt

`func (o *CreditNoteItemObjectFee) SetRefundedAt(v time.Time)`

SetRefundedAt sets RefundedAt field to given value.

### HasRefundedAt

`func (o *CreditNoteItemObjectFee) HasRefundedAt() bool`

HasRefundedAt returns a boolean if a field has been set.

### SetRefundedAtNil

`func (o *CreditNoteItemObjectFee) SetRefundedAtNil(b bool)`

 SetRefundedAtNil sets the value for RefundedAt to be an explicit nil

### UnsetRefundedAt
`func (o *CreditNoteItemObjectFee) UnsetRefundedAt()`

UnsetRefundedAt ensures that no value is present for RefundedAt, not even an explicit nil
### GetEventTransactionId

`func (o *CreditNoteItemObjectFee) GetEventTransactionId() string`

GetEventTransactionId returns the EventTransactionId field if non-nil, zero value otherwise.

### GetEventTransactionIdOk

`func (o *CreditNoteItemObjectFee) GetEventTransactionIdOk() (*string, bool)`

GetEventTransactionIdOk returns a tuple with the EventTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTransactionId

`func (o *CreditNoteItemObjectFee) SetEventTransactionId(v string)`

SetEventTransactionId sets EventTransactionId field to given value.

### HasEventTransactionId

`func (o *CreditNoteItemObjectFee) HasEventTransactionId() bool`

HasEventTransactionId returns a boolean if a field has been set.

### SetEventTransactionIdNil

`func (o *CreditNoteItemObjectFee) SetEventTransactionIdNil(b bool)`

 SetEventTransactionIdNil sets the value for EventTransactionId to be an explicit nil

### UnsetEventTransactionId
`func (o *CreditNoteItemObjectFee) UnsetEventTransactionId()`

UnsetEventTransactionId ensures that no value is present for EventTransactionId, not even an explicit nil
### GetItem

`func (o *CreditNoteItemObjectFee) GetItem() FeeObjectItem`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *CreditNoteItemObjectFee) GetItemOk() (*FeeObjectItem, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *CreditNoteItemObjectFee) SetItem(v FeeObjectItem)`

SetItem sets Item field to given value.


### GetAppliedTaxes

`func (o *CreditNoteItemObjectFee) GetAppliedTaxes() []FeeAppliedTaxObject`

GetAppliedTaxes returns the AppliedTaxes field if non-nil, zero value otherwise.

### GetAppliedTaxesOk

`func (o *CreditNoteItemObjectFee) GetAppliedTaxesOk() (*[]FeeAppliedTaxObject, bool)`

GetAppliedTaxesOk returns a tuple with the AppliedTaxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedTaxes

`func (o *CreditNoteItemObjectFee) SetAppliedTaxes(v []FeeAppliedTaxObject)`

SetAppliedTaxes sets AppliedTaxes field to given value.

### HasAppliedTaxes

`func (o *CreditNoteItemObjectFee) HasAppliedTaxes() bool`

HasAppliedTaxes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


