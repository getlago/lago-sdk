=begin
#Lago API documentation

#Lago API allows your application to push customer information and metrics (events) from your application to the billing application.

The version of the OpenAPI document: 0.52.0-beta
Contact: tech@getlago.com
Generated by: https://openapi-generator.tech
OpenAPI Generator version: 7.2.0-SNAPSHOT

=end

require 'date'
require 'time'

module LagoAPI
  class CustomerChargeUsageObject
    # The number of units consumed by the customer for a specific charge item.
    attr_accessor :units

    # The quantity of usage events that have been recorded for a particular charge during the specified time period. These events may also be referred to as the number of transactions in some contexts.
    attr_accessor :events_count

    # The amount in cents, tax excluded, consumed by the customer for a specific charge item.
    attr_accessor :amount_cents

    attr_accessor :amount_currency

    attr_accessor :charge

    attr_accessor :billable_metric

    # Array of group object, representing multiple dimensions for a charge item.
    attr_accessor :groups

    # Attribute mapping from ruby-style variable name to JSON key.
    def self.attribute_map
      {
        :'units' => :'units',
        :'events_count' => :'events_count',
        :'amount_cents' => :'amount_cents',
        :'amount_currency' => :'amount_currency',
        :'charge' => :'charge',
        :'billable_metric' => :'billable_metric',
        :'groups' => :'groups'
      }
    end

    # Returns all the JSON keys this model knows about
    def self.acceptable_attributes
      attribute_map.values
    end

    # Attribute type mapping.
    def self.openapi_types
      {
        :'units' => :'String',
        :'events_count' => :'Integer',
        :'amount_cents' => :'Integer',
        :'amount_currency' => :'Currency',
        :'charge' => :'CustomerChargeUsageObjectCharge',
        :'billable_metric' => :'CustomerChargeUsageObjectBillableMetric',
        :'groups' => :'Array<CustomerChargeUsageObjectGroupsInner>'
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
        fail ArgumentError, "The input argument (attributes) must be a hash in `LagoAPI::CustomerChargeUsageObject` initialize method"
      end

      # check to see if the attribute exists and convert string to symbol for hash key
      attributes = attributes.each_with_object({}) { |(k, v), h|
        if (!self.class.attribute_map.key?(k.to_sym))
          fail ArgumentError, "`#{k}` is not a valid attribute in `LagoAPI::CustomerChargeUsageObject`. Please check the name to make sure it's valid. List of attributes: " + self.class.attribute_map.keys.inspect
        end
        h[k.to_sym] = v
      }

      if attributes.key?(:'units')
        self.units = attributes[:'units']
      else
        self.units = nil
      end

      if attributes.key?(:'events_count')
        self.events_count = attributes[:'events_count']
      else
        self.events_count = nil
      end

      if attributes.key?(:'amount_cents')
        self.amount_cents = attributes[:'amount_cents']
      else
        self.amount_cents = nil
      end

      if attributes.key?(:'amount_currency')
        self.amount_currency = attributes[:'amount_currency']
      else
        self.amount_currency = nil
      end

      if attributes.key?(:'charge')
        self.charge = attributes[:'charge']
      else
        self.charge = nil
      end

      if attributes.key?(:'billable_metric')
        self.billable_metric = attributes[:'billable_metric']
      else
        self.billable_metric = nil
      end

      if attributes.key?(:'groups')
        if (value = attributes[:'groups']).is_a?(Array)
          self.groups = value
        end
      else
        self.groups = nil
      end
    end

    # Show invalid properties with the reasons. Usually used together with valid?
    # @return Array for valid properties with the reasons
    def list_invalid_properties
      warn '[DEPRECATED] the `list_invalid_properties` method is obsolete'
      invalid_properties = Array.new
      if @units.nil?
        invalid_properties.push('invalid value for "units", units cannot be nil.')
      end

      pattern = Regexp.new(/^[0-9]+.?[0-9]*$/)
      if @units !~ pattern
        invalid_properties.push("invalid value for \"units\", must conform to the pattern #{pattern}.")
      end

      if @events_count.nil?
        invalid_properties.push('invalid value for "events_count", events_count cannot be nil.')
      end

      if @amount_cents.nil?
        invalid_properties.push('invalid value for "amount_cents", amount_cents cannot be nil.')
      end

      if @amount_currency.nil?
        invalid_properties.push('invalid value for "amount_currency", amount_currency cannot be nil.')
      end

      if @charge.nil?
        invalid_properties.push('invalid value for "charge", charge cannot be nil.')
      end

      if @billable_metric.nil?
        invalid_properties.push('invalid value for "billable_metric", billable_metric cannot be nil.')
      end

      if @groups.nil?
        invalid_properties.push('invalid value for "groups", groups cannot be nil.')
      end

      invalid_properties
    end

    # Check to see if the all the properties in the model are valid
    # @return true if the model is valid
    def valid?
      warn '[DEPRECATED] the `valid?` method is obsolete'
      return false if @units.nil?
      return false if @units !~ Regexp.new(/^[0-9]+.?[0-9]*$/)
      return false if @events_count.nil?
      return false if @amount_cents.nil?
      return false if @amount_currency.nil?
      return false if @charge.nil?
      return false if @billable_metric.nil?
      return false if @groups.nil?
      true
    end

    # Custom attribute writer method with validation
    # @param [Object] units Value to be assigned
    def units=(units)
      if units.nil?
        fail ArgumentError, 'units cannot be nil'
      end

      pattern = Regexp.new(/^[0-9]+.?[0-9]*$/)
      if units !~ pattern
        fail ArgumentError, "invalid value for \"units\", must conform to the pattern #{pattern}."
      end

      @units = units
    end

    # Checks equality by comparing each attribute.
    # @param [Object] Object to be compared
    def ==(o)
      return true if self.equal?(o)
      self.class == o.class &&
          units == o.units &&
          events_count == o.events_count &&
          amount_cents == o.amount_cents &&
          amount_currency == o.amount_currency &&
          charge == o.charge &&
          billable_metric == o.billable_metric &&
          groups == o.groups
    end

    # @see the `==` method
    # @param [Object] Object to be compared
    def eql?(o)
      self == o
    end

    # Calculates hash code according to all attributes.
    # @return [Integer] Hash code
    def hash
      [units, events_count, amount_cents, amount_currency, charge, billable_metric, groups].hash
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
