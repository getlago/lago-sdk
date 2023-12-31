=begin
#Lago API documentation

#Lago API allows your application to push customer information and metrics (events) from your application to the billing application.

The version of the OpenAPI document: 0.51.0-beta
Contact: tech@getlago.com
Generated by: https://openapi-generator.tech
OpenAPI Generator version: 7.2.0-SNAPSHOT

=end

require 'date'
require 'time'

module LagoAPI
  class OrganizationUpdateInputOrganization
    # The URL of your newest updated webhook endpoint. This URL allows your organization to receive important messages, notifications, or data from the Lago system. By configuring your webhook endpoint to this URL, you can ensure that your organization stays informed and receives relevant information in a timely manner.
    attr_accessor :webhook_url

    attr_accessor :country

    attr_accessor :default_currency

    # The first line of your organization’s billing address.
    attr_accessor :address_line1

    # The second line of your organization’s billing address.
    attr_accessor :address_line2

    # The state of your organization’s billing address.
    attr_accessor :state

    # The zipcode of your organization’s billing address.
    attr_accessor :zipcode

    # The email address of your organization used to bill your customers.
    attr_accessor :email

    # The city of your organization’s billing address.
    attr_accessor :city

    # The legal name of your organization.
    attr_accessor :legal_name

    # The legal number of your organization.
    attr_accessor :legal_number

    # The net payment term, expressed in days, specifies the duration within which a customer is expected to remit payment after the invoice is finalized.
    attr_accessor :net_payment_term

    # The tax identification number of your organization.
    attr_accessor :tax_identification_number

    attr_accessor :timezone

    # Represents the email settings of the organization. It allows you to define which documents are sent by email. The field value determines the types of documents that trigger email notifications. Possible values for are `invoice.finalized` and `credit_note.created`. By configuring this field, you can specify whether invoices, credit notes, or both should be sent to recipients via email.
    attr_accessor :email_settings

    attr_accessor :billing_configuration

    class EnumAttributeValidator
      attr_reader :datatype
      attr_reader :allowable_values

      def initialize(datatype, allowable_values)
        @allowable_values = allowable_values.map do |value|
          case datatype.to_s
          when /Integer/i
            value.to_i
          when /Float/i
            value.to_f
          else
            value
          end
        end
      end

      def valid?(value)
        !value || allowable_values.include?(value)
      end
    end

    # Attribute mapping from ruby-style variable name to JSON key.
    def self.attribute_map
      {
        :'webhook_url' => :'webhook_url',
        :'country' => :'country',
        :'default_currency' => :'default_currency',
        :'address_line1' => :'address_line1',
        :'address_line2' => :'address_line2',
        :'state' => :'state',
        :'zipcode' => :'zipcode',
        :'email' => :'email',
        :'city' => :'city',
        :'legal_name' => :'legal_name',
        :'legal_number' => :'legal_number',
        :'net_payment_term' => :'net_payment_term',
        :'tax_identification_number' => :'tax_identification_number',
        :'timezone' => :'timezone',
        :'email_settings' => :'email_settings',
        :'billing_configuration' => :'billing_configuration'
      }
    end

    # Returns all the JSON keys this model knows about
    def self.acceptable_attributes
      attribute_map.values
    end

    # Attribute type mapping.
    def self.openapi_types
      {
        :'webhook_url' => :'String',
        :'country' => :'Country',
        :'default_currency' => :'Currency',
        :'address_line1' => :'String',
        :'address_line2' => :'String',
        :'state' => :'String',
        :'zipcode' => :'String',
        :'email' => :'String',
        :'city' => :'String',
        :'legal_name' => :'String',
        :'legal_number' => :'String',
        :'net_payment_term' => :'Integer',
        :'tax_identification_number' => :'String',
        :'timezone' => :'Timezone',
        :'email_settings' => :'Array<String>',
        :'billing_configuration' => :'OrganizationBillingConfiguration'
      }
    end

    # List of attributes with nullable: true
    def self.openapi_nullable
      Set.new([
        :'webhook_url',
        :'address_line1',
        :'address_line2',
        :'state',
        :'zipcode',
        :'email',
        :'city',
        :'legal_name',
        :'legal_number',
        :'tax_identification_number',
      ])
    end

    # Initializes the object
    # @param [Hash] attributes Model attributes in the form of hash
    def initialize(attributes = {})
      if (!attributes.is_a?(Hash))
        fail ArgumentError, "The input argument (attributes) must be a hash in `LagoAPI::OrganizationUpdateInputOrganization` initialize method"
      end

      # check to see if the attribute exists and convert string to symbol for hash key
      attributes = attributes.each_with_object({}) { |(k, v), h|
        if (!self.class.attribute_map.key?(k.to_sym))
          fail ArgumentError, "`#{k}` is not a valid attribute in `LagoAPI::OrganizationUpdateInputOrganization`. Please check the name to make sure it's valid. List of attributes: " + self.class.attribute_map.keys.inspect
        end
        h[k.to_sym] = v
      }

      if attributes.key?(:'webhook_url')
        self.webhook_url = attributes[:'webhook_url']
      end

      if attributes.key?(:'country')
        self.country = attributes[:'country']
      end

      if attributes.key?(:'default_currency')
        self.default_currency = attributes[:'default_currency']
      end

      if attributes.key?(:'address_line1')
        self.address_line1 = attributes[:'address_line1']
      end

      if attributes.key?(:'address_line2')
        self.address_line2 = attributes[:'address_line2']
      end

      if attributes.key?(:'state')
        self.state = attributes[:'state']
      end

      if attributes.key?(:'zipcode')
        self.zipcode = attributes[:'zipcode']
      end

      if attributes.key?(:'email')
        self.email = attributes[:'email']
      end

      if attributes.key?(:'city')
        self.city = attributes[:'city']
      end

      if attributes.key?(:'legal_name')
        self.legal_name = attributes[:'legal_name']
      end

      if attributes.key?(:'legal_number')
        self.legal_number = attributes[:'legal_number']
      end

      if attributes.key?(:'net_payment_term')
        self.net_payment_term = attributes[:'net_payment_term']
      end

      if attributes.key?(:'tax_identification_number')
        self.tax_identification_number = attributes[:'tax_identification_number']
      end

      if attributes.key?(:'timezone')
        self.timezone = attributes[:'timezone']
      end

      if attributes.key?(:'email_settings')
        if (value = attributes[:'email_settings']).is_a?(Array)
          self.email_settings = value
        end
      end

      if attributes.key?(:'billing_configuration')
        self.billing_configuration = attributes[:'billing_configuration']
      end
    end

    # Show invalid properties with the reasons. Usually used together with valid?
    # @return Array for valid properties with the reasons
    def list_invalid_properties
      warn '[DEPRECATED] the `list_invalid_properties` method is obsolete'
      invalid_properties = Array.new
      invalid_properties
    end

    # Check to see if the all the properties in the model are valid
    # @return true if the model is valid
    def valid?
      warn '[DEPRECATED] the `valid?` method is obsolete'
      true
    end

    # Checks equality by comparing each attribute.
    # @param [Object] Object to be compared
    def ==(o)
      return true if self.equal?(o)
      self.class == o.class &&
          webhook_url == o.webhook_url &&
          country == o.country &&
          default_currency == o.default_currency &&
          address_line1 == o.address_line1 &&
          address_line2 == o.address_line2 &&
          state == o.state &&
          zipcode == o.zipcode &&
          email == o.email &&
          city == o.city &&
          legal_name == o.legal_name &&
          legal_number == o.legal_number &&
          net_payment_term == o.net_payment_term &&
          tax_identification_number == o.tax_identification_number &&
          timezone == o.timezone &&
          email_settings == o.email_settings &&
          billing_configuration == o.billing_configuration
    end

    # @see the `==` method
    # @param [Object] Object to be compared
    def eql?(o)
      self == o
    end

    # Calculates hash code according to all attributes.
    # @return [Integer] Hash code
    def hash
      [webhook_url, country, default_currency, address_line1, address_line2, state, zipcode, email, city, legal_name, legal_number, net_payment_term, tax_identification_number, timezone, email_settings, billing_configuration].hash
    end

    # Builds the object from hash
    # @param [Hash] attributes Model attributes in the form of hash
    # @return [Object] Returns the model itself
    def self.build_from_hash(attributes)
      return nil unless attributes.is_a?(Hash)
      attributes = attributes.transform_keys(&:to_sym)
      transformed_hash = {}
      openapi_types.each_pair do |key, type|
        if attributes.key?(attribute_map[key]) && attributes[attribute_map[key]].nil?
          transformed_hash["#{key}"] = nil
        elsif type =~ /\AArray<(.*)>/i
          # check to ensure the input is an array given that the attribute
          # is documented as an array but the input is not
          if attributes[attribute_map[key]].is_a?(Array)
            transformed_hash["#{key}"] = attributes[attribute_map[key]].map { |v| _deserialize($1, v) }
          end
        elsif !attributes[attribute_map[key]].nil?
          transformed_hash["#{key}"] = _deserialize(type, attributes[attribute_map[key]])
        end
      end
      new(transformed_hash)
    end

    # Deserializes the data based on type
    # @param string type Data type
    # @param string value Value to be deserialized
    # @return [Object] Deserialized data
    def self._deserialize(type, value)
      case type.to_sym
      when :Time
        Time.parse(value)
      when :Date
        Date.parse(value)
      when :String
        value.to_s
      when :Integer
        value.to_i
      when :Float
        value.to_f
      when :Boolean
        if value.to_s =~ /\A(true|t|yes|y|1)\z/i
          true
        else
          false
        end
      when :Object
        # generic object (usually a Hash), return directly
        value
      when /\AArray<(?<inner_type>.+)>\z/
        inner_type = Regexp.last_match[:inner_type]
        value.map { |v| _deserialize(inner_type, v) }
      when /\AHash<(?<k_type>.+?), (?<v_type>.+)>\z/
        k_type = Regexp.last_match[:k_type]
        v_type = Regexp.last_match[:v_type]
        {}.tap do |hash|
          value.each do |k, v|
            hash[_deserialize(k_type, k)] = _deserialize(v_type, v)
          end
        end
      else # model
        # models (e.g. Pet) or oneOf
        klass = LagoAPI.const_get(type)
        klass.respond_to?(:openapi_any_of) || klass.respond_to?(:openapi_one_of) ? klass.build(value) : klass.build_from_hash(value)
      end
    end

    # Returns the string representation of the object
    # @return [String] String presentation of the object
    def to_s
      to_hash.to_s
    end

    # to_body is an alias to to_hash (backward compatibility)
    # @return [Hash] Returns the object in the form of hash
    def to_body
      to_hash
    end

    # Returns the object in the form of hash
    # @return [Hash] Returns the object in the form of hash
    def to_hash
      hash = {}
      self.class.attribute_map.each_pair do |attr, param|
        value = self.send(attr)
        if value.nil?
          is_nullable = self.class.openapi_nullable.include?(attr)
          next if !is_nullable || (is_nullable && !instance_variable_defined?(:"@#{attr}"))
        end

        hash[param] = _to_hash(value)
      end
      hash
    end

    # Outputs non-array value in the form of hash
    # For object, use to_hash. Otherwise, just return the value
    # @param [Object] value Any valid value
    # @return [Hash] Returns the value in the form of hash
    def _to_hash(value)
      if value.is_a?(Array)
        value.compact.map { |v| _to_hash(v) }
      elsif value.is_a?(Hash)
        {}.tap do |hash|
          value.each { |k, v| hash[k] = _to_hash(v) }
        end
      elsif value.respond_to? :to_hash
        value.to_hash
      else
        value
      end
    end

  end

end
