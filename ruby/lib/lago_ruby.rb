=begin
#Lago API documentation

#Lago API allows your application to push customer information and metrics (events) from your application to the billing application.

The version of the OpenAPI document: 0.53.0-beta
Contact: tech@getlago.com
Generated by: https://openapi-generator.tech
OpenAPI Generator version: 7.4.0-SNAPSHOT

=end

# Common files
require 'lago_ruby/api_client'
require 'lago_ruby/api_error'
require 'lago_ruby/version'
require 'lago_ruby/configuration'

# Models
require 'lago_ruby/models/add_on'
require 'lago_ruby/models/add_on_base_input'
require 'lago_ruby/models/add_on_create_input'
require 'lago_ruby/models/add_on_create_input_add_on'
require 'lago_ruby/models/add_on_object'
require 'lago_ruby/models/add_on_update_input'
require 'lago_ruby/models/add_ons_paginated'
require 'lago_ruby/models/api_error_bad_request'
require 'lago_ruby/models/api_error_forbidden'
require 'lago_ruby/models/api_error_not_allowed'
require 'lago_ruby/models/api_error_not_found'
require 'lago_ruby/models/api_error_unauthorized'
require 'lago_ruby/models/api_error_unprocessable_entity'
require 'lago_ruby/models/applied_coupon'
require 'lago_ruby/models/applied_coupon_input'
require 'lago_ruby/models/applied_coupon_input_applied_coupon'
require 'lago_ruby/models/applied_coupon_object'
require 'lago_ruby/models/applied_coupon_object_extended'
require 'lago_ruby/models/applied_coupons_paginated'
require 'lago_ruby/models/base_applied_tax'
require 'lago_ruby/models/billable_metric'
require 'lago_ruby/models/billable_metric_base_input'
require 'lago_ruby/models/billable_metric_create_input'
require 'lago_ruby/models/billable_metric_create_input_billable_metric'
require 'lago_ruby/models/billable_metric_group'
require 'lago_ruby/models/billable_metric_group_values_inner'
require 'lago_ruby/models/billable_metric_group_values_inner_one_of'
require 'lago_ruby/models/billable_metric_object'
require 'lago_ruby/models/billable_metric_update_input'
require 'lago_ruby/models/billable_metrics_paginated'
require 'lago_ruby/models/charge_object'
require 'lago_ruby/models/charge_object_properties'
require 'lago_ruby/models/charge_properties'
require 'lago_ruby/models/charge_properties_graduated_percentage_ranges_inner'
require 'lago_ruby/models/charge_properties_graduated_ranges_inner'
require 'lago_ruby/models/charge_properties_volume_ranges_inner'
require 'lago_ruby/models/country'
require 'lago_ruby/models/coupon'
require 'lago_ruby/models/coupon_base_input'
require 'lago_ruby/models/coupon_base_input_applies_to'
require 'lago_ruby/models/coupon_create_input'
require 'lago_ruby/models/coupon_create_input_coupon'
require 'lago_ruby/models/coupon_object'
require 'lago_ruby/models/coupon_update_input'
require 'lago_ruby/models/coupons_paginated'
require 'lago_ruby/models/credit_note'
require 'lago_ruby/models/credit_note_applied_tax_object'
require 'lago_ruby/models/credit_note_create_input'
require 'lago_ruby/models/credit_note_create_input_credit_note'
require 'lago_ruby/models/credit_note_estimate_input'
require 'lago_ruby/models/credit_note_estimate_input_credit_note'
require 'lago_ruby/models/credit_note_estimate_input_credit_note_items_inner'
require 'lago_ruby/models/credit_note_estimated'
require 'lago_ruby/models/credit_note_estimated_estimated_credit_note'
require 'lago_ruby/models/credit_note_estimated_estimated_credit_note_applied_taxes_inner'
require 'lago_ruby/models/credit_note_estimated_estimated_credit_note_items_inner'
require 'lago_ruby/models/credit_note_item_object'
require 'lago_ruby/models/credit_note_item_object_fee'
require 'lago_ruby/models/credit_note_object'
require 'lago_ruby/models/credit_note_update_input'
require 'lago_ruby/models/credit_note_update_input_credit_note'
require 'lago_ruby/models/credit_notes'
require 'lago_ruby/models/credit_object'
require 'lago_ruby/models/credit_object_invoice'
require 'lago_ruby/models/credit_object_item'
require 'lago_ruby/models/currency'
require 'lago_ruby/models/customer'
require 'lago_ruby/models/customer_billing_configuration'
require 'lago_ruby/models/customer_charge_grouped_usage_object_inner'
require 'lago_ruby/models/customer_charge_groups_usage_object_inner'
require 'lago_ruby/models/customer_charge_usage_object'
require 'lago_ruby/models/customer_charge_usage_object_billable_metric'
require 'lago_ruby/models/customer_charge_usage_object_charge'
require 'lago_ruby/models/customer_create_input'
require 'lago_ruby/models/customer_create_input_customer'
require 'lago_ruby/models/customer_create_input_customer_metadata_inner'
require 'lago_ruby/models/customer_metadata'
require 'lago_ruby/models/customer_object'
require 'lago_ruby/models/customer_object_extended'
require 'lago_ruby/models/customer_past_usage'
require 'lago_ruby/models/customer_usage'
require 'lago_ruby/models/customer_usage_object'
require 'lago_ruby/models/customers_paginated'
require 'lago_ruby/models/event'
require 'lago_ruby/models/event_batch_input'
require 'lago_ruby/models/event_estimate_fees_input'
require 'lago_ruby/models/event_estimate_fees_input_event'
require 'lago_ruby/models/event_input'
require 'lago_ruby/models/event_input_event'
require 'lago_ruby/models/event_input_event_timestamp'
require 'lago_ruby/models/event_object'
require 'lago_ruby/models/event_object_properties'
require 'lago_ruby/models/fee'
require 'lago_ruby/models/fee_applied_tax_object'
require 'lago_ruby/models/fee_object'
require 'lago_ruby/models/fee_object_amount_details'
require 'lago_ruby/models/fee_object_amount_details_all_of_graduated_percentage_ranges_inner'
require 'lago_ruby/models/fee_object_amount_details_all_of_graduated_ranges_inner'
require 'lago_ruby/models/fee_object_amount_details_all_of_volume_ranges_inner'
require 'lago_ruby/models/fee_object_item'
require 'lago_ruby/models/fee_update_input'
require 'lago_ruby/models/fee_update_input_fee'
require 'lago_ruby/models/fees'
require 'lago_ruby/models/fees_paginated'
require 'lago_ruby/models/generate_customer_checkout_url200_response'
require 'lago_ruby/models/generate_customer_checkout_url200_response_customer'
require 'lago_ruby/models/get_customer_portal_url200_response'
require 'lago_ruby/models/get_customer_portal_url200_response_customer'
require 'lago_ruby/models/gross_revenue_object'
require 'lago_ruby/models/gross_revenues'
require 'lago_ruby/models/group_object'
require 'lago_ruby/models/group_properties_object'
require 'lago_ruby/models/group_properties_object_values'
require 'lago_ruby/models/groups_paginated'
require 'lago_ruby/models/invoice'
require 'lago_ruby/models/invoice_applied_tax_object'
require 'lago_ruby/models/invoice_collection_object'
require 'lago_ruby/models/invoice_collections'
require 'lago_ruby/models/invoice_metadata_object'
require 'lago_ruby/models/invoice_object'
require 'lago_ruby/models/invoice_object_customer'
require 'lago_ruby/models/invoice_object_extended'
require 'lago_ruby/models/invoice_one_off_create_input'
require 'lago_ruby/models/invoice_one_off_create_input_invoice'
require 'lago_ruby/models/invoice_one_off_create_input_invoice_fees_inner'
require 'lago_ruby/models/invoice_update_input'
require 'lago_ruby/models/invoice_update_input_invoice'
require 'lago_ruby/models/invoice_update_input_invoice_metadata_inner'
require 'lago_ruby/models/invoiced_usage_object'
require 'lago_ruby/models/invoiced_usages'
require 'lago_ruby/models/invoices_paginated'
require 'lago_ruby/models/mrr_object'
require 'lago_ruby/models/mrrs'
require 'lago_ruby/models/organization'
require 'lago_ruby/models/organization_billing_configuration'
require 'lago_ruby/models/organization_object'
require 'lago_ruby/models/organization_update_input'
require 'lago_ruby/models/organization_update_input_organization'
require 'lago_ruby/models/pagination_meta'
require 'lago_ruby/models/plan'
require 'lago_ruby/models/plan_create_input'
require 'lago_ruby/models/plan_create_input_plan'
require 'lago_ruby/models/plan_create_input_plan_charges_inner'
require 'lago_ruby/models/plan_create_input_plan_charges_inner_group_properties_inner'
require 'lago_ruby/models/plan_object'
require 'lago_ruby/models/plan_overrides_object'
require 'lago_ruby/models/plan_overrides_object_charges_inner'
require 'lago_ruby/models/plan_update_input'
require 'lago_ruby/models/plan_update_input_plan'
require 'lago_ruby/models/plan_update_input_plan_charges_inner'
require 'lago_ruby/models/plans_paginated'
require 'lago_ruby/models/subscription'
require 'lago_ruby/models/subscription_create_input'
require 'lago_ruby/models/subscription_create_input_subscription'
require 'lago_ruby/models/subscription_object'
require 'lago_ruby/models/subscription_object_extended'
require 'lago_ruby/models/subscription_update_input'
require 'lago_ruby/models/subscription_update_input_subscription'
require 'lago_ruby/models/subscriptions_paginated'
require 'lago_ruby/models/tax'
require 'lago_ruby/models/tax_base_input'
require 'lago_ruby/models/tax_create_input'
require 'lago_ruby/models/tax_create_input_tax'
require 'lago_ruby/models/tax_object'
require 'lago_ruby/models/tax_update_input'
require 'lago_ruby/models/taxes_paginated'
require 'lago_ruby/models/timezone'
require 'lago_ruby/models/wallet'
require 'lago_ruby/models/wallet_create_input'
require 'lago_ruby/models/wallet_create_input_wallet'
require 'lago_ruby/models/wallet_create_input_wallet_recurring_transaction_rules_inner'
require 'lago_ruby/models/wallet_object'
require 'lago_ruby/models/wallet_object_recurring_transaction_rules_inner'
require 'lago_ruby/models/wallet_transaction_create_input'
require 'lago_ruby/models/wallet_transaction_create_input_wallet_transaction'
require 'lago_ruby/models/wallet_transaction_object'
require 'lago_ruby/models/wallet_transactions'
require 'lago_ruby/models/wallet_transactions_paginated'
require 'lago_ruby/models/wallet_update_input'
require 'lago_ruby/models/wallet_update_input_wallet'
require 'lago_ruby/models/wallet_update_input_wallet_recurring_transaction_rules_inner'
require 'lago_ruby/models/wallets_paginated'
require 'lago_ruby/models/webhook_endpoint'
require 'lago_ruby/models/webhook_endpoint_create_input'
require 'lago_ruby/models/webhook_endpoint_create_input_webhook_endpoint'
require 'lago_ruby/models/webhook_endpoint_object'
require 'lago_ruby/models/webhook_endpoint_update_input'
require 'lago_ruby/models/webhook_endpoints_paginated'

# APIs
require 'lago_ruby/api/add_ons_api'
require 'lago_ruby/api/analytics_api'
require 'lago_ruby/api/billable_metrics_api'
require 'lago_ruby/api/coupons_api'
require 'lago_ruby/api/credit_notes_api'
require 'lago_ruby/api/customers_api'
require 'lago_ruby/api/events_api'
require 'lago_ruby/api/fees_api'
require 'lago_ruby/api/invoices_api'
require 'lago_ruby/api/organizations_api'
require 'lago_ruby/api/plans_api'
require 'lago_ruby/api/subscriptions_api'
require 'lago_ruby/api/taxes_api'
require 'lago_ruby/api/wallets_api'
require 'lago_ruby/api/webhook_endpoints_api'
require 'lago_ruby/api/webhooks_api'

module LagoAPI
  class << self
    # Customize default settings for the SDK using block.
    #   LagoAPI.configure do |config|
    #     config.username = "xxx"
    #     config.password = "xxx"
    #   end
    # If no block given, return the default Configuration object.
    def configure
      if block_given?
        yield(Configuration.default)
      else
        Configuration.default
      end
    end
  end
end
