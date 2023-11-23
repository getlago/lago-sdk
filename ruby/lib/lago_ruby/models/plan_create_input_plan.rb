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
  class PlanCreateInputPlan
    # The name of the plan.
    attr_accessor :name

    # Specifies the name that will be displayed on an invoice. If no value is set for this field, the name of the plan will be used as the default display name.
    attr_accessor :invoice_display_name

    # The code of the plan. It serves as a unique identifier associated with a particular plan. The code is typically used for internal or system-level identification purposes, like assigning a subscription, for instance.
    attr_accessor :code

    # The interval used for recurring billing. It represents the frequency at which subscription billing occurs. The interval can be one of the following values: `yearly`, `quarterly`, `monthly`, or `weekly`.
    attr_accessor :interval

    # The description on the plan.
    attr_accessor :description

    # The base cost of the plan, excluding any applicable taxes, that is billed on a recurring basis. This value is defined at 0 if your plan is a pay-as-you-go plan.
    attr_accessor :amount_cents

    attr_accessor :amount_currency

    # The duration in days during which the base cost of the plan is offered for free.
    attr_accessor :trial_period

    # This field determines the billing timing for the plan. When set to `true`, the base cost of the plan is due at the beginning of each billing period. Conversely, when set to `false`, the base cost of the plan is due at the end of each billing period.
    attr_accessor :pay_in_advance

    # This field, when set to `true`, enables to invoice usage-based charges on monthly basis, even if the cadence of the plan is yearly. This allows customers to pay charges overage on a monthly basis. This can be set to true only if the plan’s interval is `yearly`.
    attr_accessor :bill_charges_monthly

    # List of unique code used to identify the taxes.
    attr_accessor :tax_codes

    # Additional usage-based charges for this plan.
    attr_accessor :charges

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
        :'name' => :'name',
        :'invoice_display_name' => :'invoice_display_name',
        :'code' => :'code',
        :'interval' => :'interval',
        :'description' => :'description',
        :'amount_cents' => :'amount_cents',
        :'amount_currency' => :'amount_currency',
        :'trial_period' => :'trial_period',
        :'pay_in_advance' => :'pay_in_advance',
        :'bill_charges_monthly' => :'bill_charges_monthly',
        :'tax_codes' => :'tax_codes',
        :'charges' => :'charges'
      }
    end

    # Returns all the JSON keys this model knows about
    def self.acceptable_attributes
      attribute_map.values
    end

    # Attribute type mapping.
    def self.openapi_types
      {
        :'name' => :'String',
        :'invoice_display_name' => :'String',
        :'code' => :'String',
        :'interval' => :'String',
        :'description' => :'String',
        :'amount_cents' => :'Integer',
        :'amount_currency' => :'Currency',
        :'trial_period' => :'Float',
        :'pay_in_advance' => :'Boolean',
        :'bill_charges_monthly' => :'Boolean',
        :'tax_codes' => :'Array<String>',
        :'charges' => :'Array<PlanCreateInputPlanChargesInner>'
      }
    end

    # List of attributes with nullable: true
    def self.openapi_nullable
      Set.new([
        :'bill_charges_monthly',
      ])
    end

    # Initializes the object
    # @param [Hash] attributes Model attributes in the form of hash
    def initialize(attributes = {})
      if (!attributes.is_a?(Hash))
        fail ArgumentError, "The input argument (attributes) must be a hash in `LagoAPI::PlanCreateInputPlan` initialize method"
      end

      # check to see if the attribute exists and convert string to symbol for hash key
      attributes = attributes.each_with_object({}) { |(k, v), h|
        if (!self.class.attribute_map.key?(k.to_sym))
          fail ArgumentError, "`#{k}` is not a valid attribute in `LagoAPI::PlanCreateInputPlan`. Please check the name to make sure it's valid. List of attributes: " + self.class.attribute_map.keys.inspect
        end
        h[k.to_sym] = v
      }

      if attributes.key?(:'name')
        self.name = attributes[:'name']
      end

      if attributes.key?(:'invoice_display_name')
        self.invoice_display_name = attributes[:'invoice_display_name']
      end

      if attributes.key?(:'code')
        self.code = attributes[:'code']
      end

      if attributes.key?(:'interval')
        self.interval = attributes[:'interval']
      end

      if attributes.key?(:'description')
        self.description = attributes[:'description']
      end

      if attributes.key?(:'amount_cents')
        self.amount_cents = attributes[:'amount_cents']
      end

      if attributes.key?(:'amount_currency')
        self.amount_currency = attributes[:'amount_currency']
      end

      if attributes.key?(:'trial_period')
        self.trial_period = attributes[:'trial_period']
      end

      if attributes.key?(:'pay_in_advance')
        self.pay_in_advance = attributes[:'pay_in_advance']
      end

      if attributes.key?(:'bill_charges_monthly')
        self.bill_charges_monthly = attributes[:'bill_charges_monthly']
      end

      if attributes.key?(:'tax_codes')
        if (value = attributes[:'tax_codes']).is_a?(Array)
          self.tax_codes = value
        end
      end

      if attributes.key?(:'charges')
        if (value = attributes[:'charges']).is_a?(Array)
          self.charges = value
        end
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
      interval_validator = EnumAttributeValidator.new('String', ["weekly", "monthly", "quarterly", "yearly"])
      return false unless interval_validator.valid?(@interval)
      true
    end

    # Custom attribute writer method checking allowed values (enum).
    # @param [Object] interval Object to be assigned
    def interval=(interval)
      validator = EnumAttributeValidator.new('String', ["weekly", "monthly", "quarterly", "yearly"])
      unless validator.valid?(interval)
        fail ArgumentError, "invalid value for \"interval\", must be one of #{validator.allowable_values}."
      end
      @interval = interval
    end

    # Checks equality by comparing each attribute.
    # @param [Object] Object to be compared
    def ==(o)
      return true if self.equal?(o)
      self.class == o.class &&
          name == o.name &&
          invoice_display_name == o.invoice_display_name &&
          code == o.code &&
          interval == o.interval &&
          description == o.description &&
          amount_cents == o.amount_cents &&
          amount_currency == o.amount_currency &&
          trial_period == o.trial_period &&
          pay_in_advance == o.pay_in_advance &&
          bill_charges_monthly == o.bill_charges_monthly &&
          tax_codes == o.tax_codes &&
          charges == o.charges
    end

    # @see the `==` method
    # @param [Object] Object to be compared
    def eql?(o)
      self == o
    end

    # Calculates hash code according to all attributes.
    # @return [Integer] Hash code
    def hash
      [name, invoice_display_name, code, interval, description, amount_cents, amount_currency, trial_period, pay_in_advance, bill_charges_monthly, tax_codes, charges].hash
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
