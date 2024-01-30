# lagoapi

Lago API allows your application to push customer information and metrics (events) from your application to the billing application.


## Installation & Usage

### Requirements

PHP 7.4 and later.
Should also work with PHP 8.0.

### Composer

To install the bindings via [Composer](https://getcomposer.org/), add the following to `composer.json`:

```json
{
  "repositories": [
    {
      "type": "vcs",
      "url": "https://github.com/getlago/sdk/php.git"
    }
  ],
  "require": {
    "getlago/sdk/php": "*@dev"
  }
}
```

Then run `composer install`

### Manual Installation

Download the files and include `autoload.php`:

```php
<?php
require_once('/path/to/lagoapi/vendor/autoload.php');
```

## Getting Started

Please follow the [installation procedure](#installation--usage) and then run the following:

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');



// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\AddOnsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$add_on_create_input = new \OpenAPI\Client\Model\AddOnCreateInput(); // \OpenAPI\Client\Model\AddOnCreateInput | Add-on payload

try {
    $result = $apiInstance->createAddOn($add_on_create_input);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling AddOnsApi->createAddOn: ', $e->getMessage(), PHP_EOL;
}

```

## API Endpoints

All URIs are relative to *https://api.getlago.com/api/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AddOnsApi* | [**createAddOn**](docs/Api/AddOnsApi.md#createaddon) | **POST** /add_ons | Create an add-on
*AddOnsApi* | [**destroyAddOn**](docs/Api/AddOnsApi.md#destroyaddon) | **DELETE** /add_ons/{code} | Delete an add-on
*AddOnsApi* | [**findAddOn**](docs/Api/AddOnsApi.md#findaddon) | **GET** /add_ons/{code} | Retrieve an add-on
*AddOnsApi* | [**findAllAddOns**](docs/Api/AddOnsApi.md#findalladdons) | **GET** /add_ons | List all add-ons
*AddOnsApi* | [**updateAddOn**](docs/Api/AddOnsApi.md#updateaddon) | **PUT** /add_ons/{code} | Update an add-on
*AnalyticsApi* | [**findAllGrossRevenues**](docs/Api/AnalyticsApi.md#findallgrossrevenues) | **GET** /analytics/gross_revenue | List gross revenue
*AnalyticsApi* | [**findAllInvoiceCollections**](docs/Api/AnalyticsApi.md#findallinvoicecollections) | **GET** /analytics/invoice_collection | List of finalized invoices
*AnalyticsApi* | [**findAllInvoicedUsages**](docs/Api/AnalyticsApi.md#findallinvoicedusages) | **GET** /analytics/invoiced_usage | List usage revenue
*AnalyticsApi* | [**findAllMrrs**](docs/Api/AnalyticsApi.md#findallmrrs) | **GET** /analytics/mrr | List MRR
*BillableMetricsApi* | [**createBillableMetric**](docs/Api/BillableMetricsApi.md#createbillablemetric) | **POST** /billable_metrics | Create a billable metric
*BillableMetricsApi* | [**destroyBillableMetric**](docs/Api/BillableMetricsApi.md#destroybillablemetric) | **DELETE** /billable_metrics/{code} | Delete a billable metric
*BillableMetricsApi* | [**findAllBillableMetricGroups**](docs/Api/BillableMetricsApi.md#findallbillablemetricgroups) | **GET** /billable_metrics/{code}/groups | Find a billable metric&#39;s groups
*BillableMetricsApi* | [**findAllBillableMetrics**](docs/Api/BillableMetricsApi.md#findallbillablemetrics) | **GET** /billable_metrics | List all billable metrics
*BillableMetricsApi* | [**findBillableMetric**](docs/Api/BillableMetricsApi.md#findbillablemetric) | **GET** /billable_metrics/{code} | Retrieve a billable metric
*BillableMetricsApi* | [**updateBillableMetric**](docs/Api/BillableMetricsApi.md#updatebillablemetric) | **PUT** /billable_metrics/{code} | Update a billable metric
*CouponsApi* | [**applyCoupon**](docs/Api/CouponsApi.md#applycoupon) | **POST** /applied_coupons | Apply a coupon to a customer
*CouponsApi* | [**createCoupon**](docs/Api/CouponsApi.md#createcoupon) | **POST** /coupons | Create a coupon
*CouponsApi* | [**deleteAppliedCoupon**](docs/Api/CouponsApi.md#deleteappliedcoupon) | **DELETE** /customers/{external_customer_id}/applied_coupons/{applied_coupon_id} | Delete an applied coupon
*CouponsApi* | [**destroyCoupon**](docs/Api/CouponsApi.md#destroycoupon) | **DELETE** /coupons/{code} | Delete a coupon
*CouponsApi* | [**findAllAppliedCoupons**](docs/Api/CouponsApi.md#findallappliedcoupons) | **GET** /applied_coupons | List all applied coupons
*CouponsApi* | [**findAllCoupons**](docs/Api/CouponsApi.md#findallcoupons) | **GET** /coupons | List all coupons
*CouponsApi* | [**findCoupon**](docs/Api/CouponsApi.md#findcoupon) | **GET** /coupons/{code} | Retrieve a coupon
*CouponsApi* | [**updateCoupon**](docs/Api/CouponsApi.md#updatecoupon) | **PUT** /coupons/{code} | Update a coupon
*CreditNotesApi* | [**createCreditNote**](docs/Api/CreditNotesApi.md#createcreditnote) | **POST** /credit_notes | Create a credit note
*CreditNotesApi* | [**downloadCreditNote**](docs/Api/CreditNotesApi.md#downloadcreditnote) | **POST** /credit_notes/{lago_id}/download | Download a credit note PDF
*CreditNotesApi* | [**estimateCreditNote**](docs/Api/CreditNotesApi.md#estimatecreditnote) | **POST** /credit_notes/estimate | Estimate amounts for a new credit note
*CreditNotesApi* | [**findAllCreditNotes**](docs/Api/CreditNotesApi.md#findallcreditnotes) | **GET** /credit_notes | List all credit notes
*CreditNotesApi* | [**findCreditNote**](docs/Api/CreditNotesApi.md#findcreditnote) | **GET** /credit_notes/{lago_id} | Retrieve a credit note
*CreditNotesApi* | [**updateCreditNote**](docs/Api/CreditNotesApi.md#updatecreditnote) | **PUT** /credit_notes/{lago_id} | Update a credit note
*CreditNotesApi* | [**voidCreditNote**](docs/Api/CreditNotesApi.md#voidcreditnote) | **PUT** /credit_notes/{lago_id}/void | Void a credit note
*CustomersApi* | [**createCustomer**](docs/Api/CustomersApi.md#createcustomer) | **POST** /customers | Create a customer
*CustomersApi* | [**deleteAppliedCoupon**](docs/Api/CustomersApi.md#deleteappliedcoupon) | **DELETE** /customers/{external_customer_id}/applied_coupons/{applied_coupon_id} | Delete an applied coupon
*CustomersApi* | [**destroyCustomer**](docs/Api/CustomersApi.md#destroycustomer) | **DELETE** /customers/{external_id} | Delete a customer
*CustomersApi* | [**findAllCustomerPastUsage**](docs/Api/CustomersApi.md#findallcustomerpastusage) | **GET** /customers/{external_customer_id}/past_usage | Retrieve customer past usage
*CustomersApi* | [**findAllCustomers**](docs/Api/CustomersApi.md#findallcustomers) | **GET** /customers | List all customers
*CustomersApi* | [**findCustomer**](docs/Api/CustomersApi.md#findcustomer) | **GET** /customers/{external_id} | Retrieve a customer
*CustomersApi* | [**findCustomerCurrentUsage**](docs/Api/CustomersApi.md#findcustomercurrentusage) | **GET** /customers/{external_customer_id}/current_usage | Retrieve customer current usage
*CustomersApi* | [**generateCustomerCheckoutURL**](docs/Api/CustomersApi.md#generatecustomercheckouturl) | **POST** /customers/{external_customer_id}/checkout_url | Generate a Customer Payment Provider Checkout URL
*CustomersApi* | [**getCustomerPortalUrl**](docs/Api/CustomersApi.md#getcustomerportalurl) | **GET** /customers/{external_customer_id}/portal_url | Get a customer portal URL
*EventsApi* | [**createBatchEvents**](docs/Api/EventsApi.md#createbatchevents) | **POST** /events/batch | Batch multiple events
*EventsApi* | [**createEvent**](docs/Api/EventsApi.md#createevent) | **POST** /events | Send usage events
*EventsApi* | [**eventEstimateFees**](docs/Api/EventsApi.md#eventestimatefees) | **POST** /events/estimate_fees | Estimate fees for a pay in advance charge
*EventsApi* | [**findEvent**](docs/Api/EventsApi.md#findevent) | **GET** /events/{transaction_id} | Retrieve a specific event
*FeesApi* | [**findAllFees**](docs/Api/FeesApi.md#findallfees) | **GET** /fees | List all fees
*FeesApi* | [**findFee**](docs/Api/FeesApi.md#findfee) | **GET** /fees/{lago_id} | Retrieve a specific fee
*FeesApi* | [**updateFee**](docs/Api/FeesApi.md#updatefee) | **PUT** /fees/{lago_id} | Update a fee
*InvoicesApi* | [**createInvoice**](docs/Api/InvoicesApi.md#createinvoice) | **POST** /invoices | Create a one-off invoice
*InvoicesApi* | [**downloadInvoice**](docs/Api/InvoicesApi.md#downloadinvoice) | **POST** /invoices/{lago_id}/download | Download an invoice PDF
*InvoicesApi* | [**finalizeInvoice**](docs/Api/InvoicesApi.md#finalizeinvoice) | **PUT** /invoices/{lago_id}/finalize | Finalize a draft invoice
*InvoicesApi* | [**findAllInvoices**](docs/Api/InvoicesApi.md#findallinvoices) | **GET** /invoices | List all invoices
*InvoicesApi* | [**findInvoice**](docs/Api/InvoicesApi.md#findinvoice) | **GET** /invoices/{lago_id} | Retrieve an invoice
*InvoicesApi* | [**refreshInvoice**](docs/Api/InvoicesApi.md#refreshinvoice) | **PUT** /invoices/{lago_id}/refresh | Refresh a draft invoice
*InvoicesApi* | [**retryPayment**](docs/Api/InvoicesApi.md#retrypayment) | **POST** /invoices/{lago_id}/retry_payment | Retry an invoice payment
*InvoicesApi* | [**updateInvoice**](docs/Api/InvoicesApi.md#updateinvoice) | **PUT** /invoices/{lago_id} | Update an invoice
*InvoicesApi* | [**voidInvoice**](docs/Api/InvoicesApi.md#voidinvoice) | **POST** /invoices/{lago_id}/void | Void an invoice
*OrganizationsApi* | [**updateOrganization**](docs/Api/OrganizationsApi.md#updateorganization) | **PUT** /organizations | Update your organization
*PlansApi* | [**createPlan**](docs/Api/PlansApi.md#createplan) | **POST** /plans | Create a plan
*PlansApi* | [**destroyPlan**](docs/Api/PlansApi.md#destroyplan) | **DELETE** /plans/{code} | Delete a plan
*PlansApi* | [**findAllPlans**](docs/Api/PlansApi.md#findallplans) | **GET** /plans | List all plans
*PlansApi* | [**findPlan**](docs/Api/PlansApi.md#findplan) | **GET** /plans/{code} | Retrieve a plan
*PlansApi* | [**updatePlan**](docs/Api/PlansApi.md#updateplan) | **PUT** /plans/{code} | Update a plan
*SubscriptionsApi* | [**createSubscription**](docs/Api/SubscriptionsApi.md#createsubscription) | **POST** /subscriptions | Assign a plan to a customer
*SubscriptionsApi* | [**destroySubscription**](docs/Api/SubscriptionsApi.md#destroysubscription) | **DELETE** /subscriptions/{external_id} | Terminate a subscription
*SubscriptionsApi* | [**findAllSubscriptions**](docs/Api/SubscriptionsApi.md#findallsubscriptions) | **GET** /subscriptions | List all subscriptions
*SubscriptionsApi* | [**findSubscription**](docs/Api/SubscriptionsApi.md#findsubscription) | **GET** /subscriptions/{external_id} | Retrieve a subscription
*SubscriptionsApi* | [**updateSubscription**](docs/Api/SubscriptionsApi.md#updatesubscription) | **PUT** /subscriptions/{external_id} | Update a subscription
*TaxesApi* | [**createTax**](docs/Api/TaxesApi.md#createtax) | **POST** /taxes | Create a tax
*TaxesApi* | [**destroyTax**](docs/Api/TaxesApi.md#destroytax) | **DELETE** /taxes/{code} | Delete a tax
*TaxesApi* | [**findAllTaxes**](docs/Api/TaxesApi.md#findalltaxes) | **GET** /taxes | List all taxes
*TaxesApi* | [**findTax**](docs/Api/TaxesApi.md#findtax) | **GET** /taxes/{code} | Retrieve a Tax
*TaxesApi* | [**updateTax**](docs/Api/TaxesApi.md#updatetax) | **PUT** /taxes/{code} | Update a tax
*WalletsApi* | [**createWallet**](docs/Api/WalletsApi.md#createwallet) | **POST** /wallets | Create a wallet
*WalletsApi* | [**createWalletTransaction**](docs/Api/WalletsApi.md#createwallettransaction) | **POST** /wallet_transactions | Top up a wallet
*WalletsApi* | [**destroyWallet**](docs/Api/WalletsApi.md#destroywallet) | **DELETE** /wallets/{lago_id} | Terminate a wallet
*WalletsApi* | [**findAllWalletTransactions**](docs/Api/WalletsApi.md#findallwallettransactions) | **GET** /wallets/{lago_id}/wallet_transactions | List all wallet transactions
*WalletsApi* | [**findAllWallets**](docs/Api/WalletsApi.md#findallwallets) | **GET** /wallets | List all wallets
*WalletsApi* | [**findWallet**](docs/Api/WalletsApi.md#findwallet) | **GET** /wallets/{lago_id} | Retrieve a wallet
*WalletsApi* | [**updateWallet**](docs/Api/WalletsApi.md#updatewallet) | **PUT** /wallets/{lago_id} | Update a wallet
*WebhookEndpointsApi* | [**createWebhookEndpoint**](docs/Api/WebhookEndpointsApi.md#createwebhookendpoint) | **POST** /webhook_endpoints | Create a webhook_endpoint
*WebhookEndpointsApi* | [**destroyWebhookEndpoint**](docs/Api/WebhookEndpointsApi.md#destroywebhookendpoint) | **DELETE** /webhook_endpoints/{lago_id} | Delete a webhook endpoint
*WebhookEndpointsApi* | [**findAllWebhookEndpoints**](docs/Api/WebhookEndpointsApi.md#findallwebhookendpoints) | **GET** /webhook_endpoints | List all webhook endpoints
*WebhookEndpointsApi* | [**findWebhookEndpoint**](docs/Api/WebhookEndpointsApi.md#findwebhookendpoint) | **GET** /webhook_endpoints/{lago_id} | Retrieve a webhook endpoint
*WebhookEndpointsApi* | [**updateWebhookEndpoint**](docs/Api/WebhookEndpointsApi.md#updatewebhookendpoint) | **PUT** /webhook_endpoints/{lago_id} | Update a webhook endpoint
*WebhooksApi* | [**fetchPublicKey**](docs/Api/WebhooksApi.md#fetchpublickey) | **GET** /webhooks/public_key | Retrieve webhook public key

## Models

- [AddOn](docs/Model/AddOn.md)
- [AddOnBaseInput](docs/Model/AddOnBaseInput.md)
- [AddOnCreateInput](docs/Model/AddOnCreateInput.md)
- [AddOnCreateInputAddOn](docs/Model/AddOnCreateInputAddOn.md)
- [AddOnObject](docs/Model/AddOnObject.md)
- [AddOnUpdateInput](docs/Model/AddOnUpdateInput.md)
- [AddOnsPaginated](docs/Model/AddOnsPaginated.md)
- [ApiErrorBadRequest](docs/Model/ApiErrorBadRequest.md)
- [ApiErrorForbidden](docs/Model/ApiErrorForbidden.md)
- [ApiErrorNotAllowed](docs/Model/ApiErrorNotAllowed.md)
- [ApiErrorNotFound](docs/Model/ApiErrorNotFound.md)
- [ApiErrorUnauthorized](docs/Model/ApiErrorUnauthorized.md)
- [ApiErrorUnprocessableEntity](docs/Model/ApiErrorUnprocessableEntity.md)
- [AppliedCoupon](docs/Model/AppliedCoupon.md)
- [AppliedCouponInput](docs/Model/AppliedCouponInput.md)
- [AppliedCouponInputAppliedCoupon](docs/Model/AppliedCouponInputAppliedCoupon.md)
- [AppliedCouponObject](docs/Model/AppliedCouponObject.md)
- [AppliedCouponObjectExtended](docs/Model/AppliedCouponObjectExtended.md)
- [AppliedCouponsPaginated](docs/Model/AppliedCouponsPaginated.md)
- [BaseAppliedTax](docs/Model/BaseAppliedTax.md)
- [BillableMetric](docs/Model/BillableMetric.md)
- [BillableMetricBaseInput](docs/Model/BillableMetricBaseInput.md)
- [BillableMetricCreateInput](docs/Model/BillableMetricCreateInput.md)
- [BillableMetricCreateInputBillableMetric](docs/Model/BillableMetricCreateInputBillableMetric.md)
- [BillableMetricGroup](docs/Model/BillableMetricGroup.md)
- [BillableMetricGroupValuesInner](docs/Model/BillableMetricGroupValuesInner.md)
- [BillableMetricGroupValuesInnerOneOf](docs/Model/BillableMetricGroupValuesInnerOneOf.md)
- [BillableMetricObject](docs/Model/BillableMetricObject.md)
- [BillableMetricUpdateInput](docs/Model/BillableMetricUpdateInput.md)
- [BillableMetricsPaginated](docs/Model/BillableMetricsPaginated.md)
- [ChargeObject](docs/Model/ChargeObject.md)
- [ChargeObjectProperties](docs/Model/ChargeObjectProperties.md)
- [ChargeProperties](docs/Model/ChargeProperties.md)
- [ChargePropertiesGraduatedPercentageRangesInner](docs/Model/ChargePropertiesGraduatedPercentageRangesInner.md)
- [ChargePropertiesGraduatedRangesInner](docs/Model/ChargePropertiesGraduatedRangesInner.md)
- [ChargePropertiesVolumeRangesInner](docs/Model/ChargePropertiesVolumeRangesInner.md)
- [Country](docs/Model/Country.md)
- [Coupon](docs/Model/Coupon.md)
- [CouponBaseInput](docs/Model/CouponBaseInput.md)
- [CouponBaseInputAppliesTo](docs/Model/CouponBaseInputAppliesTo.md)
- [CouponCreateInput](docs/Model/CouponCreateInput.md)
- [CouponCreateInputCoupon](docs/Model/CouponCreateInputCoupon.md)
- [CouponObject](docs/Model/CouponObject.md)
- [CouponUpdateInput](docs/Model/CouponUpdateInput.md)
- [CouponsPaginated](docs/Model/CouponsPaginated.md)
- [CreditNote](docs/Model/CreditNote.md)
- [CreditNoteAppliedTaxObject](docs/Model/CreditNoteAppliedTaxObject.md)
- [CreditNoteCreateInput](docs/Model/CreditNoteCreateInput.md)
- [CreditNoteCreateInputCreditNote](docs/Model/CreditNoteCreateInputCreditNote.md)
- [CreditNoteEstimateInput](docs/Model/CreditNoteEstimateInput.md)
- [CreditNoteEstimateInputCreditNote](docs/Model/CreditNoteEstimateInputCreditNote.md)
- [CreditNoteEstimateInputCreditNoteItemsInner](docs/Model/CreditNoteEstimateInputCreditNoteItemsInner.md)
- [CreditNoteEstimated](docs/Model/CreditNoteEstimated.md)
- [CreditNoteEstimatedEstimatedCreditNote](docs/Model/CreditNoteEstimatedEstimatedCreditNote.md)
- [CreditNoteEstimatedEstimatedCreditNoteAppliedTaxesInner](docs/Model/CreditNoteEstimatedEstimatedCreditNoteAppliedTaxesInner.md)
- [CreditNoteEstimatedEstimatedCreditNoteItemsInner](docs/Model/CreditNoteEstimatedEstimatedCreditNoteItemsInner.md)
- [CreditNoteItemObject](docs/Model/CreditNoteItemObject.md)
- [CreditNoteItemObjectFee](docs/Model/CreditNoteItemObjectFee.md)
- [CreditNoteObject](docs/Model/CreditNoteObject.md)
- [CreditNoteUpdateInput](docs/Model/CreditNoteUpdateInput.md)
- [CreditNoteUpdateInputCreditNote](docs/Model/CreditNoteUpdateInputCreditNote.md)
- [CreditNotes](docs/Model/CreditNotes.md)
- [CreditObject](docs/Model/CreditObject.md)
- [CreditObjectInvoice](docs/Model/CreditObjectInvoice.md)
- [CreditObjectItem](docs/Model/CreditObjectItem.md)
- [Currency](docs/Model/Currency.md)
- [Customer](docs/Model/Customer.md)
- [CustomerBillingConfiguration](docs/Model/CustomerBillingConfiguration.md)
- [CustomerChargeUsageObject](docs/Model/CustomerChargeUsageObject.md)
- [CustomerChargeUsageObjectBillableMetric](docs/Model/CustomerChargeUsageObjectBillableMetric.md)
- [CustomerChargeUsageObjectCharge](docs/Model/CustomerChargeUsageObjectCharge.md)
- [CustomerChargeUsageObjectGroupsInner](docs/Model/CustomerChargeUsageObjectGroupsInner.md)
- [CustomerCreateInput](docs/Model/CustomerCreateInput.md)
- [CustomerCreateInputCustomer](docs/Model/CustomerCreateInputCustomer.md)
- [CustomerCreateInputCustomerMetadataInner](docs/Model/CustomerCreateInputCustomerMetadataInner.md)
- [CustomerMetadata](docs/Model/CustomerMetadata.md)
- [CustomerObject](docs/Model/CustomerObject.md)
- [CustomerObjectExtended](docs/Model/CustomerObjectExtended.md)
- [CustomerPastUsage](docs/Model/CustomerPastUsage.md)
- [CustomerUsage](docs/Model/CustomerUsage.md)
- [CustomerUsageObject](docs/Model/CustomerUsageObject.md)
- [CustomersPaginated](docs/Model/CustomersPaginated.md)
- [Event](docs/Model/Event.md)
- [EventBatchInput](docs/Model/EventBatchInput.md)
- [EventEstimateFeesInput](docs/Model/EventEstimateFeesInput.md)
- [EventEstimateFeesInputEvent](docs/Model/EventEstimateFeesInputEvent.md)
- [EventInput](docs/Model/EventInput.md)
- [EventInputEvent](docs/Model/EventInputEvent.md)
- [EventInputEventTimestamp](docs/Model/EventInputEventTimestamp.md)
- [EventObject](docs/Model/EventObject.md)
- [EventObjectProperties](docs/Model/EventObjectProperties.md)
- [Fee](docs/Model/Fee.md)
- [FeeAppliedTaxObject](docs/Model/FeeAppliedTaxObject.md)
- [FeeObject](docs/Model/FeeObject.md)
- [FeeObjectAmountDetails](docs/Model/FeeObjectAmountDetails.md)
- [FeeObjectAmountDetailsAllOfGraduatedPercentageRangesInner](docs/Model/FeeObjectAmountDetailsAllOfGraduatedPercentageRangesInner.md)
- [FeeObjectAmountDetailsAllOfGraduatedRangesInner](docs/Model/FeeObjectAmountDetailsAllOfGraduatedRangesInner.md)
- [FeeObjectAmountDetailsAllOfVolumeRangesInner](docs/Model/FeeObjectAmountDetailsAllOfVolumeRangesInner.md)
- [FeeObjectItem](docs/Model/FeeObjectItem.md)
- [FeeUpdateInput](docs/Model/FeeUpdateInput.md)
- [FeeUpdateInputFee](docs/Model/FeeUpdateInputFee.md)
- [Fees](docs/Model/Fees.md)
- [FeesPaginated](docs/Model/FeesPaginated.md)
- [GenerateCustomerCheckoutURL200Response](docs/Model/GenerateCustomerCheckoutURL200Response.md)
- [GetCustomerPortalUrl200Response](docs/Model/GetCustomerPortalUrl200Response.md)
- [GetCustomerPortalUrl200ResponseCustomer](docs/Model/GetCustomerPortalUrl200ResponseCustomer.md)
- [GrossRevenueObject](docs/Model/GrossRevenueObject.md)
- [GrossRevenues](docs/Model/GrossRevenues.md)
- [GroupObject](docs/Model/GroupObject.md)
- [GroupPropertiesObject](docs/Model/GroupPropertiesObject.md)
- [GroupPropertiesObjectValues](docs/Model/GroupPropertiesObjectValues.md)
- [GroupsPaginated](docs/Model/GroupsPaginated.md)
- [Invoice](docs/Model/Invoice.md)
- [InvoiceAppliedTaxObject](docs/Model/InvoiceAppliedTaxObject.md)
- [InvoiceCollectionObject](docs/Model/InvoiceCollectionObject.md)
- [InvoiceCollections](docs/Model/InvoiceCollections.md)
- [InvoiceMetadataObject](docs/Model/InvoiceMetadataObject.md)
- [InvoiceObject](docs/Model/InvoiceObject.md)
- [InvoiceObjectCustomer](docs/Model/InvoiceObjectCustomer.md)
- [InvoiceObjectExtended](docs/Model/InvoiceObjectExtended.md)
- [InvoiceOneOffCreateInput](docs/Model/InvoiceOneOffCreateInput.md)
- [InvoiceOneOffCreateInputInvoice](docs/Model/InvoiceOneOffCreateInputInvoice.md)
- [InvoiceOneOffCreateInputInvoiceFeesInner](docs/Model/InvoiceOneOffCreateInputInvoiceFeesInner.md)
- [InvoiceUpdateInput](docs/Model/InvoiceUpdateInput.md)
- [InvoiceUpdateInputInvoice](docs/Model/InvoiceUpdateInputInvoice.md)
- [InvoiceUpdateInputInvoiceMetadataInner](docs/Model/InvoiceUpdateInputInvoiceMetadataInner.md)
- [InvoicedUsageObject](docs/Model/InvoicedUsageObject.md)
- [InvoicedUsages](docs/Model/InvoicedUsages.md)
- [InvoicesPaginated](docs/Model/InvoicesPaginated.md)
- [MrrObject](docs/Model/MrrObject.md)
- [Mrrs](docs/Model/Mrrs.md)
- [Organization](docs/Model/Organization.md)
- [OrganizationBillingConfiguration](docs/Model/OrganizationBillingConfiguration.md)
- [OrganizationObject](docs/Model/OrganizationObject.md)
- [OrganizationUpdateInput](docs/Model/OrganizationUpdateInput.md)
- [OrganizationUpdateInputOrganization](docs/Model/OrganizationUpdateInputOrganization.md)
- [PaginationMeta](docs/Model/PaginationMeta.md)
- [Plan](docs/Model/Plan.md)
- [PlanCreateInput](docs/Model/PlanCreateInput.md)
- [PlanCreateInputPlan](docs/Model/PlanCreateInputPlan.md)
- [PlanCreateInputPlanChargesInner](docs/Model/PlanCreateInputPlanChargesInner.md)
- [PlanCreateInputPlanChargesInnerGroupPropertiesInner](docs/Model/PlanCreateInputPlanChargesInnerGroupPropertiesInner.md)
- [PlanObject](docs/Model/PlanObject.md)
- [PlanOverridesObject](docs/Model/PlanOverridesObject.md)
- [PlanOverridesObjectChargesInner](docs/Model/PlanOverridesObjectChargesInner.md)
- [PlanUpdateInput](docs/Model/PlanUpdateInput.md)
- [PlanUpdateInputPlan](docs/Model/PlanUpdateInputPlan.md)
- [PlanUpdateInputPlanChargesInner](docs/Model/PlanUpdateInputPlanChargesInner.md)
- [PlansPaginated](docs/Model/PlansPaginated.md)
- [Subscription](docs/Model/Subscription.md)
- [SubscriptionCreateInput](docs/Model/SubscriptionCreateInput.md)
- [SubscriptionCreateInputSubscription](docs/Model/SubscriptionCreateInputSubscription.md)
- [SubscriptionObject](docs/Model/SubscriptionObject.md)
- [SubscriptionObjectExtended](docs/Model/SubscriptionObjectExtended.md)
- [SubscriptionUpdateInput](docs/Model/SubscriptionUpdateInput.md)
- [SubscriptionUpdateInputSubscription](docs/Model/SubscriptionUpdateInputSubscription.md)
- [SubscriptionsPaginated](docs/Model/SubscriptionsPaginated.md)
- [Tax](docs/Model/Tax.md)
- [TaxBaseInput](docs/Model/TaxBaseInput.md)
- [TaxCreateInput](docs/Model/TaxCreateInput.md)
- [TaxCreateInputTax](docs/Model/TaxCreateInputTax.md)
- [TaxObject](docs/Model/TaxObject.md)
- [TaxUpdateInput](docs/Model/TaxUpdateInput.md)
- [TaxesPaginated](docs/Model/TaxesPaginated.md)
- [Timezone](docs/Model/Timezone.md)
- [Wallet](docs/Model/Wallet.md)
- [WalletCreateInput](docs/Model/WalletCreateInput.md)
- [WalletCreateInputWallet](docs/Model/WalletCreateInputWallet.md)
- [WalletCreateInputWalletRecurringTransactionRulesInner](docs/Model/WalletCreateInputWalletRecurringTransactionRulesInner.md)
- [WalletObject](docs/Model/WalletObject.md)
- [WalletObjectRecurringTransactionRulesInner](docs/Model/WalletObjectRecurringTransactionRulesInner.md)
- [WalletTransactionCreateInput](docs/Model/WalletTransactionCreateInput.md)
- [WalletTransactionCreateInputWalletTransaction](docs/Model/WalletTransactionCreateInputWalletTransaction.md)
- [WalletTransactionObject](docs/Model/WalletTransactionObject.md)
- [WalletTransactions](docs/Model/WalletTransactions.md)
- [WalletTransactionsPaginated](docs/Model/WalletTransactionsPaginated.md)
- [WalletUpdateInput](docs/Model/WalletUpdateInput.md)
- [WalletUpdateInputWallet](docs/Model/WalletUpdateInputWallet.md)
- [WalletUpdateInputWalletRecurringTransactionRulesInner](docs/Model/WalletUpdateInputWalletRecurringTransactionRulesInner.md)
- [WalletsPaginated](docs/Model/WalletsPaginated.md)
- [WebhookEndpoint](docs/Model/WebhookEndpoint.md)
- [WebhookEndpointCreateInput](docs/Model/WebhookEndpointCreateInput.md)
- [WebhookEndpointCreateInputWebhookEndpoint](docs/Model/WebhookEndpointCreateInputWebhookEndpoint.md)
- [WebhookEndpointObject](docs/Model/WebhookEndpointObject.md)
- [WebhookEndpointUpdateInput](docs/Model/WebhookEndpointUpdateInput.md)
- [WebhookEndpointsPaginated](docs/Model/WebhookEndpointsPaginated.md)

## Authorization

Authentication schemes defined for the API:
### bearerAuth

- **Type**: Bearer authentication

## Tests

To run the tests, use:

```bash
composer install
vendor/bin/phpunit
```

## Author

tech@getlago.com

## About this package

This PHP package is automatically generated by the [OpenAPI Generator](https://openapi-generator.tech) project:

- API version: `0.53.0-beta`
- Build package: `org.openapitools.codegen.languages.PhpClientCodegen`
