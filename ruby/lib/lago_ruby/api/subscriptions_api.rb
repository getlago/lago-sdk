=begin
#Lago API documentation

#Lago API allows your application to push customer information and metrics (events) from your application to the billing application.

The version of the OpenAPI document: 0.53.0-beta
Contact: tech@getlago.com
Generated by: https://openapi-generator.tech
OpenAPI Generator version: 7.4.0-SNAPSHOT

=end

require 'cgi'

module LagoAPI
  class SubscriptionsApi
    attr_accessor :api_client

    def initialize(api_client = ApiClient.default)
      @api_client = api_client
    end
    # Assign a plan to a customer
    # This endpoint assigns a plan to a customer, creating or modifying a subscription. Ideal for initial subscriptions or plan changes (upgrades/downgrades).
    # @param subscription_create_input [SubscriptionCreateInput] Subscription payload
    # @param [Hash] opts the optional parameters
    # @return [Subscription]
    def create_subscription(subscription_create_input, opts = {})
      data, _status_code, _headers = create_subscription_with_http_info(subscription_create_input, opts)
      data
    end

    # Assign a plan to a customer
    # This endpoint assigns a plan to a customer, creating or modifying a subscription. Ideal for initial subscriptions or plan changes (upgrades/downgrades).
    # @param subscription_create_input [SubscriptionCreateInput] Subscription payload
    # @param [Hash] opts the optional parameters
    # @return [Array<(Subscription, Integer, Hash)>] Subscription data, response status code and response headers
    def create_subscription_with_http_info(subscription_create_input, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: SubscriptionsApi.create_subscription ...'
      end
      # verify the required parameter 'subscription_create_input' is set
      if @api_client.config.client_side_validation && subscription_create_input.nil?
        fail ArgumentError, "Missing the required parameter 'subscription_create_input' when calling SubscriptionsApi.create_subscription"
      end
      # resource path
      local_var_path = '/subscriptions'

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
      post_body = opts[:debug_body] || @api_client.object_to_http_body(subscription_create_input)

      # return_type
      return_type = opts[:debug_return_type] || 'Subscription'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['bearerAuth']

      new_options = opts.merge(
        :operation => :"SubscriptionsApi.create_subscription",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:POST, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: SubscriptionsApi#create_subscription\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end

    # Terminate a subscription
    # This endpoint allows you to terminate a subscription.
    # @param external_id [String] External ID of the existing subscription
    # @param [Hash] opts the optional parameters
    # @option opts [String] :status If the field is not defined, Lago will terminate only &#x60;active&#x60; subscriptions. However, if you wish to cancel a &#x60;pending&#x60; subscription, please ensure that you include &#x60;status&#x3D;pending&#x60; in your request.
    # @return [Subscription]
    def destroy_subscription(external_id, opts = {})
      data, _status_code, _headers = destroy_subscription_with_http_info(external_id, opts)
      data
    end

    # Terminate a subscription
    # This endpoint allows you to terminate a subscription.
    # @param external_id [String] External ID of the existing subscription
    # @param [Hash] opts the optional parameters
    # @option opts [String] :status If the field is not defined, Lago will terminate only &#x60;active&#x60; subscriptions. However, if you wish to cancel a &#x60;pending&#x60; subscription, please ensure that you include &#x60;status&#x3D;pending&#x60; in your request.
    # @return [Array<(Subscription, Integer, Hash)>] Subscription data, response status code and response headers
    def destroy_subscription_with_http_info(external_id, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: SubscriptionsApi.destroy_subscription ...'
      end
      # verify the required parameter 'external_id' is set
      if @api_client.config.client_side_validation && external_id.nil?
        fail ArgumentError, "Missing the required parameter 'external_id' when calling SubscriptionsApi.destroy_subscription"
      end
      # resource path
      local_var_path = '/subscriptions/{external_id}'.sub('{' + 'external_id' + '}', CGI.escape(external_id.to_s))

      # query parameters
      query_params = opts[:query_params] || {}
      query_params[:'status'] = opts[:'status'] if !opts[:'status'].nil?

      # header parameters
      header_params = opts[:header_params] || {}
      # HTTP header 'Accept' (if needed)
      header_params['Accept'] = @api_client.select_header_accept(['application/json'])

      # form parameters
      form_params = opts[:form_params] || {}

      # http body (model)
      post_body = opts[:debug_body]

      # return_type
      return_type = opts[:debug_return_type] || 'Subscription'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['bearerAuth']

      new_options = opts.merge(
        :operation => :"SubscriptionsApi.destroy_subscription",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:DELETE, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: SubscriptionsApi#destroy_subscription\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end

    # List all subscriptions
    # This endpoint retrieves all active subscriptions.
    # @param [Hash] opts the optional parameters
    # @option opts [Integer] :page Page number.
    # @option opts [Integer] :per_page Number of records per page.
    # @option opts [String] :external_customer_id The customer external unique identifier (provided by your own application)
    # @option opts [String] :plan_code The unique code representing the plan to be attached to the customer. This code must correspond to the code property of one of the active plans.
    # @option opts [Array<String>] :status If the field is not defined, Lago will return only &#x60;active&#x60; subscriptions. However, if you wish to fetch subscriptions by different status you can define them in a status[] query param. Available filter values: &#x60;pending&#x60;, &#x60;canceled&#x60;, &#x60;terminated&#x60;, &#x60;active&#x60;.
    # @return [SubscriptionsPaginated]
    def find_all_subscriptions(opts = {})
      data, _status_code, _headers = find_all_subscriptions_with_http_info(opts)
      data
    end

    # List all subscriptions
    # This endpoint retrieves all active subscriptions.
    # @param [Hash] opts the optional parameters
    # @option opts [Integer] :page Page number.
    # @option opts [Integer] :per_page Number of records per page.
    # @option opts [String] :external_customer_id The customer external unique identifier (provided by your own application)
    # @option opts [String] :plan_code The unique code representing the plan to be attached to the customer. This code must correspond to the code property of one of the active plans.
    # @option opts [Array<String>] :status If the field is not defined, Lago will return only &#x60;active&#x60; subscriptions. However, if you wish to fetch subscriptions by different status you can define them in a status[] query param. Available filter values: &#x60;pending&#x60;, &#x60;canceled&#x60;, &#x60;terminated&#x60;, &#x60;active&#x60;.
    # @return [Array<(SubscriptionsPaginated, Integer, Hash)>] SubscriptionsPaginated data, response status code and response headers
    def find_all_subscriptions_with_http_info(opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: SubscriptionsApi.find_all_subscriptions ...'
      end
      allowable_values = ["pending", "canceled", "terminated", "active"]
      if @api_client.config.client_side_validation && opts[:'status'] && !opts[:'status'].all? { |item| allowable_values.include?(item) }
        fail ArgumentError, "invalid value for \"status\", must include one of #{allowable_values}"
      end
      # resource path
      local_var_path = '/subscriptions'

      # query parameters
      query_params = opts[:query_params] || {}
      query_params[:'page'] = opts[:'page'] if !opts[:'page'].nil?
      query_params[:'per_page'] = opts[:'per_page'] if !opts[:'per_page'].nil?
      query_params[:'external_customer_id'] = opts[:'external_customer_id'] if !opts[:'external_customer_id'].nil?
      query_params[:'plan_code'] = opts[:'plan_code'] if !opts[:'plan_code'].nil?
      query_params[:'status'] = @api_client.build_collection_param(opts[:'status'], :multi) if !opts[:'status'].nil?

      # header parameters
      header_params = opts[:header_params] || {}
      # HTTP header 'Accept' (if needed)
      header_params['Accept'] = @api_client.select_header_accept(['application/json'])

      # form parameters
      form_params = opts[:form_params] || {}

      # http body (model)
      post_body = opts[:debug_body]

      # return_type
      return_type = opts[:debug_return_type] || 'SubscriptionsPaginated'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['bearerAuth']

      new_options = opts.merge(
        :operation => :"SubscriptionsApi.find_all_subscriptions",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:GET, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: SubscriptionsApi#find_all_subscriptions\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end

    # Retrieve a subscription
    # This endpoint retrieves a specific subscription.
    # @param external_id [String] External ID of the existing subscription
    # @param [Hash] opts the optional parameters
    # @return [Subscription]
    def find_subscription(external_id, opts = {})
      data, _status_code, _headers = find_subscription_with_http_info(external_id, opts)
      data
    end

    # Retrieve a subscription
    # This endpoint retrieves a specific subscription.
    # @param external_id [String] External ID of the existing subscription
    # @param [Hash] opts the optional parameters
    # @return [Array<(Subscription, Integer, Hash)>] Subscription data, response status code and response headers
    def find_subscription_with_http_info(external_id, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: SubscriptionsApi.find_subscription ...'
      end
      # verify the required parameter 'external_id' is set
      if @api_client.config.client_side_validation && external_id.nil?
        fail ArgumentError, "Missing the required parameter 'external_id' when calling SubscriptionsApi.find_subscription"
      end
      # resource path
      local_var_path = '/subscriptions/{external_id}'.sub('{' + 'external_id' + '}', CGI.escape(external_id.to_s))

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
      return_type = opts[:debug_return_type] || 'Subscription'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['bearerAuth']

      new_options = opts.merge(
        :operation => :"SubscriptionsApi.find_subscription",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:GET, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: SubscriptionsApi#find_subscription\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end

    # Update a subscription
    # This endpoint allows you to update a subscription.
    # @param external_id [String] External ID of the existing subscription
    # @param subscription_update_input [SubscriptionUpdateInput] Update an existing subscription
    # @param [Hash] opts the optional parameters
    # @return [Subscription]
    def update_subscription(external_id, subscription_update_input, opts = {})
      data, _status_code, _headers = update_subscription_with_http_info(external_id, subscription_update_input, opts)
      data
    end

    # Update a subscription
    # This endpoint allows you to update a subscription.
    # @param external_id [String] External ID of the existing subscription
    # @param subscription_update_input [SubscriptionUpdateInput] Update an existing subscription
    # @param [Hash] opts the optional parameters
    # @return [Array<(Subscription, Integer, Hash)>] Subscription data, response status code and response headers
    def update_subscription_with_http_info(external_id, subscription_update_input, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: SubscriptionsApi.update_subscription ...'
      end
      # verify the required parameter 'external_id' is set
      if @api_client.config.client_side_validation && external_id.nil?
        fail ArgumentError, "Missing the required parameter 'external_id' when calling SubscriptionsApi.update_subscription"
      end
      # verify the required parameter 'subscription_update_input' is set
      if @api_client.config.client_side_validation && subscription_update_input.nil?
        fail ArgumentError, "Missing the required parameter 'subscription_update_input' when calling SubscriptionsApi.update_subscription"
      end
      # resource path
      local_var_path = '/subscriptions/{external_id}'.sub('{' + 'external_id' + '}', CGI.escape(external_id.to_s))

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
      post_body = opts[:debug_body] || @api_client.object_to_http_body(subscription_update_input)

      # return_type
      return_type = opts[:debug_return_type] || 'Subscription'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['bearerAuth']

      new_options = opts.merge(
        :operation => :"SubscriptionsApi.update_subscription",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:PUT, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: SubscriptionsApi#update_subscription\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end
  end
end
