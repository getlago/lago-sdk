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
  class AppliedCouponObject
    # Unique identifier of the applied coupon, created by Lago.
    attr_accessor :lago_id

    # Unique identifier of the coupon, created by Lago.
    attr_accessor :lago_coupon_id

    # Unique code used to identify the coupon.
    attr_accessor :coupon_code

    # The name of the coupon.
    attr_accessor :coupon_name

    # Unique identifier of the customer, created by Lago.
    attr_accessor :lago_customer_id

    # The customer external unique identifier (provided by your own application)
    attr_accessor :external_customer_id

    # The status of the coupon. Can be either `active` or `terminated`.
    attr_accessor :status

    # The amount of the coupon in cents. This field is required only for coupon with `fixed_amount` type.
    attr_accessor :amount_cents

    # The remaining amount in cents for a `fixed_amount` coupon with a frequency set to `once`. This field indicates the remaining balance or value that can still be utilized from the coupon.
    attr_accessor :amount_cents_remaining

    attr_accessor :amount_currency

    # The percentage rate of the coupon. This field is required only for coupons with a `percentage` coupon type.
    attr_accessor :percentage_rate

    # The type of frequency for the coupon. It can have three possible values: `once`, `recurring` or `forever`.  - If set to `once`, the coupon is applicable only for a single use. - If set to `recurring`, the coupon can be used multiple times for recurring billing periods. - If set to `forever`, the coupon has unlimited usage and can be applied indefinitely.
    attr_accessor :frequency

    # Specifies the number of billing periods to which the coupon applies. This field is required only for coupons with a `recurring` frequency type
    attr_accessor :frequency_duration

    # The remaining number of billing periods to which the coupon is applicable. This field determines the remaining usage or availability of the coupon based on the remaining billing periods.
    attr_accessor :frequency_duration_remaining

    # The date and time after which the coupon will stop applying to customer's invoices. Once the expiration date is reached, the coupon will no longer be applicable, and any further invoices generated for the customer will not include the coupon discount.
    attr_accessor :expiration_at

    # The date and time when the coupon was assigned to a customer. It is expressed in UTC format according to the ISO 8601 datetime standard.
    attr_accessor :created_at

    # This field indicates the specific moment when the coupon amount is fully utilized or when the coupon is removed from the customer's coupon list. It is expressed in UTC format according to the ISO 8601 datetime standard.
    attr_accessor :terminated_at

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
        :'lago_id' => :'lago_id',
        :'lago_coupon_id' => :'lago_coupon_id',
        :'coupon_code' => :'coupon_code',
        :'coupon_name' => :'coupon_name',
        :'lago_customer_id' => :'lago_customer_id',
        :'external_customer_id' => :'external_customer_id',
        :'status' => :'status',
        :'amount_cents' => :'amount_cents',
        :'amount_cents_remaining' => :'amount_cents_remaining',
        :'amount_currency' => :'amount_currency',
        :'percentage_rate' => :'percentage_rate',
        :'frequency' => :'frequency',
        :'frequency_duration' => :'frequency_duration',
        :'frequency_duration_remaining' => :'frequency_duration_remaining',
        :'expiration_at' => :'expiration_at',
        :'created_at' => :'created_at',
        :'terminated_at' => :'terminated_at'
      }
    end

    # Returns all the JSON keys this model knows about
    def self.acceptable_attributes
      attribute_map.values
    end

    # Attribute type mapping.
    def self.openapi_types
      {
        :'lago_id' => :'String',
        :'lago_coupon_id' => :'String',
        :'coupon_code' => :'String',
        :'coupon_name' => :'String',
        :'lago_customer_id' => :'String',
        :'external_customer_id' => :'String',
        :'status' => :'String',
        :'amount_cents' => :'Integer',
        :'amount_cents_remaining' => :'Integer',
        :'amount_currency' => :'Currency',
        :'percentage_rate' => :'String',
        :'frequency' => :'String',
        :'frequency_duration' => :'Integer',
        :'frequency_duration_remaining' => :'Integer',
        :'expiration_at' => :'Time',
        :'created_at' => :'Time',
        :'terminated_at' => :'Time'
      }
    end

    # List of attributes with nullable: true
    def self.openapi_nullable
      Set.new([
        :'amount_cents',
        :'amount_cents_remaining',
        :'percentage_rate',
        :'frequency_duration',
        :'frequency_duration_remaining',
        :'expiration_at',
        :'terminated_at'
      ])
    end

    # Initializes the object
    # @param [Hash] attributes Model attributes in the form of hash
    def initialize(attributes = {})
      if (!attributes.is_a?(Hash))
        fail ArgumentError, "The input argument (attributes) must be a hash in `LagoAPI::AppliedCouponObject` initialize method"
      end

      # check to see if the attribute exists and convert string to symbol for hash key
      attributes = attributes.each_with_object({}) { |(k, v), h|
        if (!self.class.attribute_map.key?(k.to_sym))
          fail ArgumentError, "`#{k}` is not a valid attribute in `LagoAPI::AppliedCouponObject`. Please check the name to make sure it's valid. List of attributes: " + self.class.attribute_map.keys.inspect
        end
        h[k.to_sym] = v
      }

      if attributes.key?(:'lago_id')
        self.lago_id = attributes[:'lago_id']
      else
        self.lago_id = nil
      end

      if attributes.key?(:'lago_coupon_id')
        self.lago_coupon_id = attributes[:'lago_coupon_id']
      else
        self.lago_coupon_id = nil
      end

      if attributes.key?(:'coupon_code')
        self.coupon_code = attributes[:'coupon_code']
      else
        self.coupon_code = nil
      end

      if attributes.key?(:'coupon_name')
        self.coupon_name = attributes[:'coupon_name']
      else
        self.coupon_name = nil
      end

      if attributes.key?(:'lago_customer_id')
        self.lago_customer_id = attributes[:'lago_customer_id']
      else
        self.lago_customer_id = nil
      end

      if attributes.key?(:'external_customer_id')
        self.external_customer_id = attributes[:'external_customer_id']
      else
        self.external_customer_id = nil
      end

      if attributes.key?(:'status')
        self.status = attributes[:'status']
      else
        self.status = nil
      end

      if attributes.key?(:'amount_cents')
        self.amount_cents = attributes[:'amount_cents']
      end

      if attributes.key?(:'amount_cents_remaining')
        self.amount_cents_remaining = attributes[:'amount_cents_remaining']
      end

      if attributes.key?(:'amount_currency')
        self.amount_currency = attributes[:'amount_currency']
      end

      if attributes.key?(:'percentage_rate')
        self.percentage_rate = attributes[:'percentage_rate']
      end

      if attributes.key?(:'frequency')
        self.frequency = attributes[:'frequency']
      else
        self.frequency = nil
      end

      if attributes.key?(:'frequency_duration')
        self.frequency_duration = attributes[:'frequency_duration']
      end

      if attributes.key?(:'frequency_duration_remaining')
        self.frequency_duration_remaining = attributes[:'frequency_duration_remaining']
      end

      if attributes.key?(:'expiration_at')
        self.expiration_at = attributes[:'expiration_at']
      end

      if attributes.key?(:'created_at')
        self.created_at = attributes[:'created_at']
      else
        self.created_at = nil
      end

      if attributes.key?(:'terminated_at')
        self.terminated_at = attributes[:'terminated_at']
      end
    end

    # Show invalid properties with the reasons. Usually used together with valid?
    # @return Array for valid properties with the reasons
    def list_invalid_properties
      warn '[DEPRECATED] the `list_invalid_properties` method is obsolete'
      invalid_properties = Array.new
      if @lago_id.nil?
        invalid_properties.push('invalid value for "lago_id", lago_id cannot be nil.')
      end

      if @lago_coupon_id.nil?
        invalid_properties.push('invalid value for "lago_coupon_id", lago_coupon_id cannot be nil.')
      end

      if @coupon_code.nil?
        invalid_properties.push('invalid value for "coupon_code", coupon_code cannot be nil.')
      end

      if @coupon_name.nil?
        invalid_properties.push('invalid value for "coupon_name", coupon_name cannot be nil.')
      end

      if @lago_customer_id.nil?
        invalid_properties.push('invalid value for "lago_customer_id", lago_customer_id cannot be nil.')
      end

      if @external_customer_id.nil?
        invalid_properties.push('invalid value for "external_customer_id", external_customer_id cannot be nil.')
      end

      if @status.nil?
        invalid_properties.push('invalid value for "status", status cannot be nil.')
      end

      pattern = Regexp.new(/^[0-9]+.?[0-9]*$/)
      if !@percentage_rate.nil? && @percentage_rate !~ pattern
        invalid_properties.push("invalid value for \"percentage_rate\", must conform to the pattern #{pattern}.")
      end

      if @frequency.nil?
        invalid_properties.push('invalid value for "frequency", frequency cannot be nil.')
      end

      if @created_at.nil?
        invalid_properties.push('invalid value for "created_at", created_at cannot be nil.')
      end

      invalid_properties
    end

    # Check to see if the all the properties in the model are valid
    # @return true if the model is valid
    def valid?
      warn '[DEPRECATED] the `valid?` method is obsolete'
      return false if @lago_id.nil?
      return false if @lago_coupon_id.nil?
      return false if @coupon_code.nil?
      return false if @coupon_name.nil?
      return false if @lago_customer_id.nil?
      return false if @external_customer_id.nil?
      return false if @status.nil?
      status_validator = EnumAttributeValidator.new('String', ["active", "terminated"])
      return false unless status_validator.valid?(@status)
      return false if !@percentage_rate.nil? && @percentage_rate !~ Regexp.new(/^[0-9]+.?[0-9]*$/)
      return false if @frequency.nil?
      frequency_validator = EnumAttributeValidator.new('String', ["once", "recurring"])
      return false unless frequency_validator.valid?(@frequency)
      return false if @created_at.nil?
      true
    end

    # Custom attribute writer method checking allowed values (enum).
    # @param [Object] status Object to be assigned
    def status=(status)
      validator = EnumAttributeValidator.new('String', ["active", "terminated"])
      unless validator.valid?(status)
        fail ArgumentError, "invalid value for \"status\", must be one of #{validator.allowable_values}."
      end
      @status = status
    end

    # Custom attribute writer method with validation
    # @param [Object] percentage_rate Value to be assigned
    def percentage_rate=(percentage_rate)
      pattern = Regexp.new(/^[0-9]+.?[0-9]*$/)
      if !percentage_rate.nil? && percentage_rate !~ pattern
        fail ArgumentError, "invalid value for \"percentage_rate\", must conform to the pattern #{pattern}."
      end

      @percentage_rate = percentage_rate
    end

    # Custom attribute writer method checking allowed values (enum).
    # @param [Object] frequency Object to be assigned
    def frequency=(frequency)
      validator = EnumAttributeValidator.new('String', ["once", "recurring"])
      unless validator.valid?(frequency)
        fail ArgumentError, "invalid value for \"frequency\", must be one of #{validator.allowable_values}."
      end
      @frequency = frequency
    end

    # Checks equality by comparing each attribute.
    # @param [Object] Object to be compared
    def ==(o)
      return true if self.equal?(o)
      self.class == o.class &&
          lago_id == o.lago_id &&
          lago_coupon_id == o.lago_coupon_id &&
          coupon_code == o.coupon_code &&
          coupon_name == o.coupon_name &&
          lago_customer_id == o.lago_customer_id &&
          external_customer_id == o.external_customer_id &&
          status == o.status &&
          amount_cents == o.amount_cents &&
          amount_cents_remaining == o.amount_cents_remaining &&
          amount_currency == o.amount_currency &&
          percentage_rate == o.percentage_rate &&
          frequency == o.frequency &&
          frequency_duration == o.frequency_duration &&
          frequency_duration_remaining == o.frequency_duration_remaining &&
          expiration_at == o.expiration_at &&
          created_at == o.created_at &&
          terminated_at == o.terminated_at
    end

    # @see the `==` method
    # @param [Object] Object to be compared
    def eql?(o)
      self == o
    end

    # Calculates hash code according to all attributes.
    # @return [Integer] Hash code
    def hash
      [lago_id, lago_coupon_id, coupon_code, coupon_name, lago_customer_id, external_customer_id, status, amount_cents, amount_cents_remaining, amount_currency, percentage_rate, frequency, frequency_duration, frequency_duration_remaining, expiration_at, created_at, terminated_at].hash
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
