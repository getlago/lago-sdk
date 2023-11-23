=begin
#Lago API documentation

#Lago API allows your application to push customer information and metrics (events) from your application to the billing application.

The version of the OpenAPI document: 0.51.0-beta
Contact: tech@getlago.com
Generated by: https://openapi-generator.tech
OpenAPI Generator version: 7.2.0-SNAPSHOT

=end

require 'cgi'

module LagoAPI
  class CustomersApi
    attr_accessor :api_client

    def initialize(api_client = ApiClient.default)
      @api_client = api_client
    end
    # Create a customer
    # This endpoint creates a new customer.
    # @param customer_create_input [CustomerCreateInput] Customer payload
    # @param [Hash] opts the optional parameters
    # @return [Customer]
    def create_customer(customer_create_input, opts = {})
      data, _status_code, _headers = create_customer_with_http_info(customer_create_input, opts)
      data
    end

    # Create a customer
    # This endpoint creates a new customer.
    # @param customer_create_input [CustomerCreateInput] Customer payload
    # @param [Hash] opts the optional parameters
    # @return [Array<(Customer, Integer, Hash)>] Customer data, response status code and response headers
    def create_customer_with_http_info(customer_create_input, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: CustomersApi.create_customer ...'
      end
      # verify the required parameter 'customer_create_input' is set
      if @api_client.config.client_side_validation && customer_create_input.nil?
        fail ArgumentError, "Missing the required parameter 'customer_create_input' when calling CustomersApi.create_customer"
      end
      # resource path
      local_var_path = '/customers'

      # query parameters
      query_params = opts[:query_params] || {}

      # header parameters
      header_params = opts[:header_params] || {}
      # HTTP header 'Accept' (if needed)
      header_params['Accept'] = @api_client.select_header_accept(['application/json'])
      # HTTP header 'Content-Type'
      content_type = @api_client.select_header_content_type(['application/json'])
      if !content_type.nil?
          header_params['Content-Type'] = content_type
      end

      # form parameters
      form_params = opts[:form_params] || {}

      # http body (model)
      post_body = opts[:debug_body] || @api_client.object_to_http_body(customer_create_input)

      # return_type
      return_type = opts[:debug_return_type] || 'Customer'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['bearerAuth']

      new_options = opts.merge(
        :operation => :"CustomersApi.create_customer",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:POST, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: CustomersApi#create_customer\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end

    # Delete an applied coupon
    # This endpoint is used to delete a specific coupon that has been applied to a customer.
    # @param external_customer_id [String] The customer external unique identifier (provided by your own application)
    # @param applied_coupon_id [String] Unique identifier of the applied coupon, created by Lago.
    # @param [Hash] opts the optional parameters
    # @return [AppliedCoupon]
    def delete_applied_coupon(external_customer_id, applied_coupon_id, opts = {})
      data, _status_code, _headers = delete_applied_coupon_with_http_info(external_customer_id, applied_coupon_id, opts)
      data
    end

    # Delete an applied coupon
    # This endpoint is used to delete a specific coupon that has been applied to a customer.
    # @param external_customer_id [String] The customer external unique identifier (provided by your own application)
    # @param applied_coupon_id [String] Unique identifier of the applied coupon, created by Lago.
    # @param [Hash] opts the optional parameters
    # @return [Array<(AppliedCoupon, Integer, Hash)>] AppliedCoupon data, response status code and response headers
    def delete_applied_coupon_with_http_info(external_customer_id, applied_coupon_id, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: CustomersApi.delete_applied_coupon ...'
      end
      # verify the required parameter 'external_customer_id' is set
      if @api_client.config.client_side_validation && external_customer_id.nil?
        fail ArgumentError, "Missing the required parameter 'external_customer_id' when calling CustomersApi.delete_applied_coupon"
      end
      # verify the required parameter 'applied_coupon_id' is set
      if @api_client.config.client_side_validation && applied_coupon_id.nil?
        fail ArgumentError, "Missing the required parameter 'applied_coupon_id' when calling CustomersApi.delete_applied_coupon"
      end
      # resource path
      local_var_path = '/customers/{external_customer_id}/applied_coupons/{applied_coupon_id}'.sub('{' + 'external_customer_id' + '}', CGI.escape(external_customer_id.to_s)).sub('{' + 'applied_coupon_id' + '}', CGI.escape(applied_coupon_id.to_s))

      # query parameters
      query_params = opts[:query_params] || {}

      # header parameters
      header_params = opts[:header_params] || {}
      # HTTP header 'Accept' (if needed)
      header_params['Accept'] = @api_client.select_header_accept(['application/json'])

      # form parameters
      form_params = opts[:form_params] || {}

      # http body (model)
      post_body = opts[:debug_body]

      # return_type
      return_type = opts[:debug_return_type] || 'AppliedCoupon'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['bearerAuth']

      new_options = opts.merge(
        :operation => :"CustomersApi.delete_applied_coupon",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:DELETE, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: CustomersApi#delete_applied_coupon\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end

    # Delete a customer
    # This endpoint deletes an existing customer.
    # @param external_id [String] External ID of the existing customer
    # @param [Hash] opts the optional parameters
    # @return [Customer]
    def destroy_customer(external_id, opts = {})
      data, _status_code, _headers = destroy_customer_with_http_info(external_id, opts)
      data
    end

    # Delete a customer
    # This endpoint deletes an existing customer.
    # @param external_id [String] External ID of the existing customer
    # @param [Hash] opts the optional parameters
    # @return [Array<(Customer, Integer, Hash)>] Customer data, response status code and response headers
    def destroy_customer_with_http_info(external_id, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: CustomersApi.destroy_customer ...'
      end
      # verify the required parameter 'external_id' is set
      if @api_client.config.client_side_validation && external_id.nil?
        fail ArgumentError, "Missing the required parameter 'external_id' when calling CustomersApi.destroy_customer"
      end
      # resource path
      local_var_path = '/customers/{external_id}'.sub('{' + 'external_id' + '}', CGI.escape(external_id.to_s))

      # query parameters
      query_params = opts[:query_params] || {}

      # header parameters
      header_params = opts[:header_params] || {}
      # HTTP header 'Accept' (if needed)
      header_params['Accept'] = @api_client.select_header_accept(['application/json'])

      # form parameters
      form_params = opts[:form_params] || {}

      # http body (model)
      post_body = opts[:debug_body]

      # return_type
      return_type = opts[:debug_return_type] || 'Customer'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['bearerAuth']

      new_options = opts.merge(
        :operation => :"CustomersApi.destroy_customer",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:DELETE, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: CustomersApi#destroy_customer\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end

    # Retrieve customer past usage
    # This endpoint enables the retrieval of the usage-based billing data for a customer within past periods.
    # @param external_customer_id [String] The customer external unique identifier (provided by your own application).
    # @param external_subscription_id [String] The unique identifier of the subscription within your application.
    # @param [Hash] opts the optional parameters
    # @option opts [Integer] :page Page number.
    # @option opts [Integer] :per_page Number of records per page.
    # @option opts [String] :billable_metric_code Billable metric code filter to apply to the charge usage
    # @option opts [Integer] :periods_count Number of past billing period to returns in the result
    # @return [CustomerPastUsage]
    def find_all_customer_past_usage(external_customer_id, external_subscription_id, opts = {})
      data, _status_code, _headers = find_all_customer_past_usage_with_http_info(external_customer_id, external_subscription_id, opts)
      data
    end

    # Retrieve customer past usage
    # This endpoint enables the retrieval of the usage-based billing data for a customer within past periods.
    # @param external_customer_id [String] The customer external unique identifier (provided by your own application).
    # @param external_subscription_id [String] The unique identifier of the subscription within your application.
    # @param [Hash] opts the optional parameters
    # @option opts [Integer] :page Page number.
    # @option opts [Integer] :per_page Number of records per page.
    # @option opts [String] :billable_metric_code Billable metric code filter to apply to the charge usage
    # @option opts [Integer] :periods_count Number of past billing period to returns in the result
    # @return [Array<(CustomerPastUsage, Integer, Hash)>] CustomerPastUsage data, response status code and response headers
    def find_all_customer_past_usage_with_http_info(external_customer_id, external_subscription_id, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: CustomersApi.find_all_customer_past_usage ...'
      end
      # verify the required parameter 'external_customer_id' is set
      if @api_client.config.client_side_validation && external_customer_id.nil?
        fail ArgumentError, "Missing the required parameter 'external_customer_id' when calling CustomersApi.find_all_customer_past_usage"
      end
      # verify the required parameter 'external_subscription_id' is set
      if @api_client.config.client_side_validation && external_subscription_id.nil?
        fail ArgumentError, "Missing the required parameter 'external_subscription_id' when calling CustomersApi.find_all_customer_past_usage"
      end
      # resource path
      local_var_path = '/customers/{external_customer_id}/past_usage'.sub('{' + 'external_customer_id' + '}', CGI.escape(external_customer_id.to_s))

      # query parameters
      query_params = opts[:query_params] || {}
      query_params[:'external_subscription_id'] = external_subscription_id
      query_params[:'page'] = opts[:'page'] if !opts[:'page'].nil?
      query_params[:'per_page'] = opts[:'per_page'] if !opts[:'per_page'].nil?
      query_params[:'billable_metric_code'] = opts[:'billable_metric_code'] if !opts[:'billable_metric_code'].nil?
      query_params[:'periods_count'] = opts[:'periods_count'] if !opts[:'periods_count'].nil?

      # header parameters
      header_params = opts[:header_params] || {}
      # HTTP header 'Accept' (if needed)
      header_params['Accept'] = @api_client.select_header_accept(['application/json'])

      # form parameters
      form_params = opts[:form_params] || {}

      # http body (model)
      post_body = opts[:debug_body]

      # return_type
      return_type = opts[:debug_return_type] || 'CustomerPastUsage'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['bearerAuth']

      new_options = opts.merge(
        :operation => :"CustomersApi.find_all_customer_past_usage",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:GET, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: CustomersApi#find_all_customer_past_usage\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end

    # List all customers
    # This endpoint retrieves all existing customers.
    # @param [Hash] opts the optional parameters
    # @option opts [Integer] :page Page number.
    # @option opts [Integer] :per_page Number of records per page.
    # @return [CustomersPaginated]
    def find_all_customers(opts = {})
      data, _status_code, _headers = find_all_customers_with_http_info(opts)
      data
    end

    # List all customers
    # This endpoint retrieves all existing customers.
    # @param [Hash] opts the optional parameters
    # @option opts [Integer] :page Page number.
    # @option opts [Integer] :per_page Number of records per page.
    # @return [Array<(CustomersPaginated, Integer, Hash)>] CustomersPaginated data, response status code and response headers
    def find_all_customers_with_http_info(opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: CustomersApi.find_all_customers ...'
      end
      # resource path
      local_var_path = '/customers'

      # query parameters
      query_params = opts[:query_params] || {}
      query_params[:'page'] = opts[:'page'] if !opts[:'page'].nil?
      query_params[:'per_page'] = opts[:'per_page'] if !opts[:'per_page'].nil?

      # header parameters
      header_params = opts[:header_params] || {}
      # HTTP header 'Accept' (if needed)
      header_params['Accept'] = @api_client.select_header_accept(['application/json'])

      # form parameters
      form_params = opts[:form_params] || {}

      # http body (model)
      post_body = opts[:debug_body]

      # return_type
      return_type = opts[:debug_return_type] || 'CustomersPaginated'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['bearerAuth']

      new_options = opts.merge(
        :operation => :"CustomersApi.find_all_customers",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:GET, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: CustomersApi#find_all_customers\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end

    # Retrieve a customer
    # This endpoint retrieves an existing customer.
    # @param external_id [String] External ID of the existing customer
    # @param [Hash] opts the optional parameters
    # @return [Customer]
    def find_customer(external_id, opts = {})
      data, _status_code, _headers = find_customer_with_http_info(external_id, opts)
      data
    end

    # Retrieve a customer
    # This endpoint retrieves an existing customer.
    # @param external_id [String] External ID of the existing customer
    # @param [Hash] opts the optional parameters
    # @return [Array<(Customer, Integer, Hash)>] Customer data, response status code and response headers
    def find_customer_with_http_info(external_id, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: CustomersApi.find_customer ...'
      end
      # verify the required parameter 'external_id' is set
      if @api_client.config.client_side_validation && external_id.nil?
        fail ArgumentError, "Missing the required parameter 'external_id' when calling CustomersApi.find_customer"
      end
      # resource path
      local_var_path = '/customers/{external_id}'.sub('{' + 'external_id' + '}', CGI.escape(external_id.to_s))

      # query parameters
      query_params = opts[:query_params] || {}

      # header parameters
      header_params = opts[:header_params] || {}
      # HTTP header 'Accept' (if needed)
      header_params['Accept'] = @api_client.select_header_accept(['application/json'])

      # form parameters
      form_params = opts[:form_params] || {}

      # http body (model)
      post_body = opts[:debug_body]

      # return_type
      return_type = opts[:debug_return_type] || 'Customer'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['bearerAuth']

      new_options = opts.merge(
        :operation => :"CustomersApi.find_customer",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:GET, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: CustomersApi#find_customer\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end

    # Retrieve customer current usage
    # This endpoint enables the retrieval of the usage-based billing data for a customer within the current period.
    # @param external_customer_id [String] The customer external unique identifier (provided by your own application).
    # @param external_subscription_id [String] The unique identifier of the subscription within your application.
    # @param [Hash] opts the optional parameters
    # @return [CustomerUsage]
    def find_customer_current_usage(external_customer_id, external_subscription_id, opts = {})
      data, _status_code, _headers = find_customer_current_usage_with_http_info(external_customer_id, external_subscription_id, opts)
      data
    end

    # Retrieve customer current usage
    # This endpoint enables the retrieval of the usage-based billing data for a customer within the current period.
    # @param external_customer_id [String] The customer external unique identifier (provided by your own application).
    # @param external_subscription_id [String] The unique identifier of the subscription within your application.
    # @param [Hash] opts the optional parameters
    # @return [Array<(CustomerUsage, Integer, Hash)>] CustomerUsage data, response status code and response headers
    def find_customer_current_usage_with_http_info(external_customer_id, external_subscription_id, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: CustomersApi.find_customer_current_usage ...'
      end
      # verify the required parameter 'external_customer_id' is set
      if @api_client.config.client_side_validation && external_customer_id.nil?
        fail ArgumentError, "Missing the required parameter 'external_customer_id' when calling CustomersApi.find_customer_current_usage"
      end
      # verify the required parameter 'external_subscription_id' is set
      if @api_client.config.client_side_validation && external_subscription_id.nil?
        fail ArgumentError, "Missing the required parameter 'external_subscription_id' when calling CustomersApi.find_customer_current_usage"
      end
      # resource path
      local_var_path = '/customers/{external_customer_id}/current_usage'.sub('{' + 'external_customer_id' + '}', CGI.escape(external_customer_id.to_s))

      # query parameters
      query_params = opts[:query_params] || {}
      query_params[:'external_subscription_id'] = external_subscription_id

      # header parameters
      header_params = opts[:header_params] || {}
      # HTTP header 'Accept' (if needed)
      header_params['Accept'] = @api_client.select_header_accept(['application/json'])

      # form parameters
      form_params = opts[:form_params] || {}

      # http body (model)
      post_body = opts[:debug_body]

      # return_type
      return_type = opts[:debug_return_type] || 'CustomerUsage'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['bearerAuth']

      new_options = opts.merge(
        :operation => :"CustomersApi.find_customer_current_usage",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:GET, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: CustomersApi#find_customer_current_usage\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end

    # Generate a Customer Payment Provider Checkout URL
    # This endpoint regenerates the Payment Provider Checkout URL of a Customer.
    # @param external_customer_id [String] The customer external unique identifier (provided by your own application).
    # @param [Hash] opts the optional parameters
    # @return [GenerateCustomerCheckoutURL200Response]
    def generate_customer_checkout_url(external_customer_id, opts = {})
      data, _status_code, _headers = generate_customer_checkout_url_with_http_info(external_customer_id, opts)
      data
    end

    # Generate a Customer Payment Provider Checkout URL
    # This endpoint regenerates the Payment Provider Checkout URL of a Customer.
    # @param external_customer_id [String] The customer external unique identifier (provided by your own application).
    # @param [Hash] opts the optional parameters
    # @return [Array<(GenerateCustomerCheckoutURL200Response, Integer, Hash)>] GenerateCustomerCheckoutURL200Response data, response status code and response headers
    def generate_customer_checkout_url_with_http_info(external_customer_id, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: CustomersApi.generate_customer_checkout_url ...'
      end
      # verify the required parameter 'external_customer_id' is set
      if @api_client.config.client_side_validation && external_customer_id.nil?
        fail ArgumentError, "Missing the required parameter 'external_customer_id' when calling CustomersApi.generate_customer_checkout_url"
      end
      # resource path
      local_var_path = '/customers/{external_customer_id}/checkout_url'.sub('{' + 'external_customer_id' + '}', CGI.escape(external_customer_id.to_s))

      # query parameters
      query_params = opts[:query_params] || {}

      # header parameters
      header_params = opts[:header_params] || {}
      # HTTP header 'Accept' (if needed)
      header_params['Accept'] = @api_client.select_header_accept(['application/json'])

      # form parameters
      form_params = opts[:form_params] || {}

      # http body (model)
      post_body = opts[:debug_body]

      # return_type
      return_type = opts[:debug_return_type] || 'GenerateCustomerCheckoutURL200Response'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['bearerAuth']

      new_options = opts.merge(
        :operation => :"CustomersApi.generate_customer_checkout_url",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:POST, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: CustomersApi#generate_customer_checkout_url\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end

    # Get a customer portal URL
    # Retrieves an embeddable link for displaying a customer portal.  This endpoint allows you to fetch the URL that can be embedded to provide customers access to a dedicated portal
    # @param external_customer_id [String] External ID of the existing customer
    # @param [Hash] opts the optional parameters
    # @return [GetCustomerPortalUrl200Response]
    def get_customer_portal_url(external_customer_id, opts = {})
      data, _status_code, _headers = get_customer_portal_url_with_http_info(external_customer_id, opts)
      data
    end

    # Get a customer portal URL
    # Retrieves an embeddable link for displaying a customer portal.  This endpoint allows you to fetch the URL that can be embedded to provide customers access to a dedicated portal
    # @param external_customer_id [String] External ID of the existing customer
    # @param [Hash] opts the optional parameters
    # @return [Array<(GetCustomerPortalUrl200Response, Integer, Hash)>] GetCustomerPortalUrl200Response data, response status code and response headers
    def get_customer_portal_url_with_http_info(external_customer_id, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: CustomersApi.get_customer_portal_url ...'
      end
      # verify the required parameter 'external_customer_id' is set
      if @api_client.config.client_side_validation && external_customer_id.nil?
        fail ArgumentError, "Missing the required parameter 'external_customer_id' when calling CustomersApi.get_customer_portal_url"
      end
      # resource path
      local_var_path = '/customers/{external_customer_id}/portal_url'.sub('{' + 'external_customer_id' + '}', CGI.escape(external_customer_id.to_s))

      # query parameters
      query_params = opts[:query_params] || {}

      # header parameters
      header_params = opts[:header_params] || {}
      # HTTP header 'Accept' (if needed)
      header_params['Accept'] = @api_client.select_header_accept(['application/json'])

      # form parameters
      form_params = opts[:form_params] || {}

      # http body (model)
      post_body = opts[:debug_body]

      # return_type
      return_type = opts[:debug_return_type] || 'GetCustomerPortalUrl200Response'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['bearerAuth']

      new_options = opts.merge(
        :operation => :"CustomersApi.get_customer_portal_url",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:GET, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: CustomersApi#get_customer_portal_url\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end
  end
end