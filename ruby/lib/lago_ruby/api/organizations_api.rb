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
  class OrganizationsApi
    attr_accessor :api_client

    def initialize(api_client = ApiClient.default)
      @api_client = api_client
    end
    # Update your organization
    # This endpoint is used to update your own organization's settings.
    # @param organization_update_input [OrganizationUpdateInput] Update an existing organization
    # @param [Hash] opts the optional parameters
    # @return [Organization]
    def update_organization(organization_update_input, opts = {})
      data, _status_code, _headers = update_organization_with_http_info(organization_update_input, opts)
      data
    end

    # Update your organization
    # This endpoint is used to update your own organization&#39;s settings.
    # @param organization_update_input [OrganizationUpdateInput] Update an existing organization
    # @param [Hash] opts the optional parameters
    # @return [Array<(Organization, Integer, Hash)>] Organization data, response status code and response headers
    def update_organization_with_http_info(organization_update_input, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: OrganizationsApi.update_organization ...'
      end
      # verify the required parameter 'organization_update_input' is set
      if @api_client.config.client_side_validation && organization_update_input.nil?
        fail ArgumentError, "Missing the required parameter 'organization_update_input' when calling OrganizationsApi.update_organization"
      end
      # resource path
      local_var_path = '/organizations'

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
      post_body = opts[:debug_body] || @api_client.object_to_http_body(organization_update_input)

      # return_type
      return_type = opts[:debug_return_type] || 'Organization'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['bearerAuth']

      new_options = opts.merge(
        :operation => :"OrganizationsApi.update_organization",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:PUT, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: OrganizationsApi#update_organization\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end
  end
end
