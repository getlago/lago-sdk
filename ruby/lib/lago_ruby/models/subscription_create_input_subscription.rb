=begin
#Lago API documentation

#Lago API allows your application to push customer information and metrics (events) from your application to the billing application.

The version of the OpenAPI document: 0.53.0-beta
Contact: tech@getlago.com
Generated by: https://openapi-generator.tech
OpenAPI Generator version: 7.4.0-SNAPSHOT

=end

require 'date'
require 'time'

module LagoAPI
  class SubscriptionCreateInputSubscription
    # The customer external unique identifier (provided by your own application)
    attr_accessor :external_customer_id

    # The unique code representing the plan to be attached to the customer. This code must correspond to the `code` property of one of the active plans.
    attr_accessor :plan_code

    # The display name of the subscription on an invoice. This field allows for customization of the subscription's name for billing purposes, especially useful when a single customer has multiple subscriptions using the same plan.
    attr_accessor :name

    # The unique external identifier for the subscription. This identifier serves as an idempotency key, ensuring that each subscription is unique.
    attr_accessor :external_id

    # The billing time for the subscription, which can be set as either `anniversary` or `calendar`. If not explicitly provided, it will default to `calendar`. The billing time determines the timing of recurring billing cycles for the subscription. By specifying `anniversary`, the billing cycle will be based on the specific date the subscription started (billed fully), while `calendar` sets the billing cycle at the first day of the week/month/year (billed with proration).
    attr_accessor :billing_time

    # The effective end date of the subscription. If this field is set to null, the subscription will automatically renew. This date should be provided in ISO 8601 datetime format, and use Coordinated Universal Time (UTC).
    attr_accessor :ending_at

    # The start date for the subscription, allowing for the creation of subscriptions that can begin in the past or future. Please note that it cannot be used to update the start date of a pending subscription or schedule an upgrade/downgrade. The start_date should be provided in ISO 8601 datetime format and expressed in Coordinated Universal Time (UTC).
    attr_accessor :subscription_at

    attr_accessor :plan_overrides

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
        :'external_customer_id' => :'external_customer_id',
        :'plan_code' => :'plan_code',
        :'name' => :'name',
        :'external_id' => :'external_id',
        :'billing_time' => :'billing_time',
        :'ending_at' => :'ending_at',
        :'subscription_at' => :'subscription_at',
        :'plan_overrides' => :'plan_overrides'
      }
    end

    # Returns all the JSON keys this model knows about
    def self.acceptable_attributes
      attribute_map.values
    end

    # Attribute type mapping.
    def self.openapi_types
      {
        :'external_customer_id' => :'String',
        :'plan_code' => :'String',
        :'name' => :'String',
        :'external_id' => :'String',
        :'billing_time' => :'String',
        :'ending_at' => :'Time',
        :'subscription_at' => :'Time',
        :'plan_overrides' => :'PlanOverridesObject'
      }
    end

    # List of attributes with nullable: true
    def self.openapi_nullable
      Set.new([
      ])
    end

    # Initializes the object
    # @param [Hash] attributes Model attributes in the form of hash
    def initialize(attributes = {})
      if (!attributes.is_a?(Hash))
        fail ArgumentError, "The input argument (attributes) must be a hash in `LagoAPI::SubscriptionCreateInputSubscription` initialize method"
      end

      # check to see if the attribute exists and convert string to symbol for hash key
      attributes = attributes.each_with_object({}) { |(k, v), h|
        if (!self.class.attribute_map.key?(k.to_sym))
          fail ArgumentError, "`#{k}` is not a valid attribute in `LagoAPI::SubscriptionCreateInputSubscription`. Please check the name to make sure it's valid. List of attributes: " + self.class.attribute_map.keys.inspect
        end
        h[k.to_sym] = v
      }

      if attributes.key?(:'external_customer_id')
        self.external_customer_id = attributes[:'external_customer_id']
      else
        self.external_customer_id = nil
      end

      if attributes.key?(:'plan_code')
        self.plan_code = attributes[:'plan_code']
      else
        self.plan_code = nil
      end

      if attributes.key?(:'name')
        self.name = attributes[:'name']
      end

      if attributes.key?(:'external_id')
        self.external_id = attributes[:'external_id']
      else
        self.external_id = nil
      end

      if attributes.key?(:'billing_time')
        self.billing_time = attributes[:'billing_time']
      end

      if attributes.key?(:'ending_at')
        self.ending_at = attributes[:'ending_at']
      end

      if attributes.key?(:'subscription_at')
        self.subscription_at = attributes[:'subscription_at']
      end

      if attributes.key?(:'plan_overrides')
        self.plan_overrides = attributes[:'plan_overrides']
      end
    end

    # Show invalid properties with the reasons. Usually used together with valid?
    # @return Array for valid properties with the reasons
    def list_invalid_properties
      warn '[DEPRECATED] the `list_invalid_properties` method is obsolete'
      invalid_properties = Array.new
      if @external_customer_id.nil?
        invalid_properties.push('invalid value for "external_customer_id", external_customer_id cannot be nil.')
      end

      if @plan_code.nil?
        invalid_properties.push('invalid value for "plan_code", plan_code cannot be nil.')
      end

      if @external_id.nil?
        invalid_properties.push('invalid value for "external_id", external_id cannot be nil.')
      end

      invalid_properties
    end

    # Check to see if the all the properties in the model are valid
    # @return true if the model is valid
    def valid?
      warn '[DEPRECATED] the `valid?` method is obsolete'
      return false if @external_customer_id.nil?
      return false if @plan_code.nil?
      return false if @external_id.nil?
      billing_time_validator = EnumAttributeValidator.new('String', ["calendar", "anniversary"])
      return false unless billing_time_validator.valid?(@billing_time)
      true
    end

    # Custom attribute writer method checking allowed values (enum).
    # @param [Object] billing_time Object to be assigned
    def billing_time=(billing_time)
      validator = EnumAttributeValidator.new('String', ["calendar", "anniversary"])
      unless validator.valid?(billing_time)
        fail ArgumentError, "invalid value for \"billing_time\", must be one of #{validator.allowable_values}."
      end
      @billing_time = billing_time
    end

    # Checks equality by comparing each attribute.
    # @param [Object] Object to be compared
    def ==(o)
      return true if self.equal?(o)
      self.class == o.class &&
          external_customer_id == o.external_customer_id &&
          plan_code == o.plan_code &&
          name == o.name &&
          external_id == o.external_id &&
          billing_time == o.billing_time &&
          ending_at == o.ending_at &&
          subscription_at == o.subscription_at &&
          plan_overrides == o.plan_overrides
    end

    # @see the `==` method
    # @param [Object] Object to be compared
    def eql?(o)
      self == o
    end

    # Calculates hash code according to all attributes.
    # @return [Integer] Hash code
    def hash
      [external_customer_id, plan_code, name, external_id, billing_time, ending_at, subscription_at, plan_overrides].hash
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
