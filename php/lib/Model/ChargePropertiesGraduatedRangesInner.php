<?php
/**
 * ChargePropertiesGraduatedRangesInner
 *
 * PHP version 7.4
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 */

/**
 * Lago API documentation
 *
 * Lago API allows your application to push customer information and metrics (events) from your application to the billing application.
 *
 * The version of the OpenAPI document: 0.53.0-beta
 * Contact: tech@getlago.com
 * Generated by: https://openapi-generator.tech
 * OpenAPI Generator version: 7.3.0-SNAPSHOT
 */

/**
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

namespace OpenAPI\Client\Model;

use \ArrayAccess;
use \OpenAPI\Client\ObjectSerializer;

/**
 * ChargePropertiesGraduatedRangesInner Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<string, mixed>
 */
class ChargePropertiesGraduatedRangesInner implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'ChargeProperties_graduated_ranges_inner';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'from_value' => 'int',
        'to_value' => 'int',
        'flat_amount' => 'string',
        'per_unit_amount' => 'string'
    ];

    /**
      * Array of property to format mappings. Used for (de)serialization
      *
      * @var string[]
      * @phpstan-var array<string, string|null>
      * @psalm-var array<string, string|null>
      */
    protected static $openAPIFormats = [
        'from_value' => null,
        'to_value' => null,
        'flat_amount' => null,
        'per_unit_amount' => null
    ];

    /**
      * Array of nullable properties. Used for (de)serialization
      *
      * @var boolean[]
      */
    protected static array $openAPINullables = [
        'from_value' => false,
        'to_value' => true,
        'flat_amount' => false,
        'per_unit_amount' => false
    ];

    /**
      * If a nullable field gets set to null, insert it here
      *
      * @var boolean[]
      */
    protected array $openAPINullablesSetToNull = [];

    /**
     * Array of property to type mappings. Used for (de)serialization
     *
     * @return array
     */
    public static function openAPITypes()
    {
        return self::$openAPITypes;
    }

    /**
     * Array of property to format mappings. Used for (de)serialization
     *
     * @return array
     */
    public static function openAPIFormats()
    {
        return self::$openAPIFormats;
    }

    /**
     * Array of nullable properties
     *
     * @return array
     */
    protected static function openAPINullables(): array
    {
        return self::$openAPINullables;
    }

    /**
     * Array of nullable field names deliberately set to null
     *
     * @return boolean[]
     */
    private function getOpenAPINullablesSetToNull(): array
    {
        return $this->openAPINullablesSetToNull;
    }

    /**
     * Setter - Array of nullable field names deliberately set to null
     *
     * @param boolean[] $openAPINullablesSetToNull
     */
    private function setOpenAPINullablesSetToNull(array $openAPINullablesSetToNull): void
    {
        $this->openAPINullablesSetToNull = $openAPINullablesSetToNull;
    }

    /**
     * Checks if a property is nullable
     *
     * @param string $property
     * @return bool
     */
    public static function isNullable(string $property): bool
    {
        return self::openAPINullables()[$property] ?? false;
    }

    /**
     * Checks if a nullable property is set to null.
     *
     * @param string $property
     * @return bool
     */
    public function isNullableSetToNull(string $property): bool
    {
        return in_array($property, $this->getOpenAPINullablesSetToNull(), true);
    }

    /**
     * Array of attributes where the key is the local name,
     * and the value is the original name
     *
     * @var string[]
     */
    protected static $attributeMap = [
        'from_value' => 'from_value',
        'to_value' => 'to_value',
        'flat_amount' => 'flat_amount',
        'per_unit_amount' => 'per_unit_amount'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'from_value' => 'setFromValue',
        'to_value' => 'setToValue',
        'flat_amount' => 'setFlatAmount',
        'per_unit_amount' => 'setPerUnitAmount'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'from_value' => 'getFromValue',
        'to_value' => 'getToValue',
        'flat_amount' => 'getFlatAmount',
        'per_unit_amount' => 'getPerUnitAmount'
    ];

    /**
     * Array of attributes where the key is the local name,
     * and the value is the original name
     *
     * @return array
     */
    public static function attributeMap()
    {
        return self::$attributeMap;
    }

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @return array
     */
    public static function setters()
    {
        return self::$setters;
    }

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @return array
     */
    public static function getters()
    {
        return self::$getters;
    }

    /**
     * The original name of the model.
     *
     * @return string
     */
    public function getModelName()
    {
        return self::$openAPIModelName;
    }


    /**
     * Associative array for storing property values
     *
     * @var mixed[]
     */
    protected $container = [];

    /**
     * Constructor
     *
     * @param mixed[] $data Associated array of property values
     *                      initializing the model
     */
    public function __construct(array $data = null)
    {
        $this->setIfExists('from_value', $data ?? [], null);
        $this->setIfExists('to_value', $data ?? [], null);
        $this->setIfExists('flat_amount', $data ?? [], null);
        $this->setIfExists('per_unit_amount', $data ?? [], null);
    }

    /**
    * Sets $this->container[$variableName] to the given data or to the given default Value; if $variableName
    * is nullable and its value is set to null in the $fields array, then mark it as "set to null" in the
    * $this->openAPINullablesSetToNull array
    *
    * @param string $variableName
    * @param array  $fields
    * @param mixed  $defaultValue
    */
    private function setIfExists(string $variableName, array $fields, $defaultValue): void
    {
        if (self::isNullable($variableName) && array_key_exists($variableName, $fields) && is_null($fields[$variableName])) {
            $this->openAPINullablesSetToNull[] = $variableName;
        }

        $this->container[$variableName] = $fields[$variableName] ?? $defaultValue;
    }

    /**
     * Show all the invalid properties with reasons.
     *
     * @return array invalid properties with reasons
     */
    public function listInvalidProperties()
    {
        $invalidProperties = [];

        if ($this->container['from_value'] === null) {
            $invalidProperties[] = "'from_value' can't be null";
        }
        if ($this->container['to_value'] === null) {
            $invalidProperties[] = "'to_value' can't be null";
        }
        if ($this->container['flat_amount'] === null) {
            $invalidProperties[] = "'flat_amount' can't be null";
        }
        if (!preg_match("/^[0-9]+.?[0-9]*$/", $this->container['flat_amount'])) {
            $invalidProperties[] = "invalid value for 'flat_amount', must be conform to the pattern /^[0-9]+.?[0-9]*$/.";
        }

        if ($this->container['per_unit_amount'] === null) {
            $invalidProperties[] = "'per_unit_amount' can't be null";
        }
        if (!preg_match("/^[0-9]+.?[0-9]*$/", $this->container['per_unit_amount'])) {
            $invalidProperties[] = "invalid value for 'per_unit_amount', must be conform to the pattern /^[0-9]+.?[0-9]*$/.";
        }

        return $invalidProperties;
    }

    /**
     * Validate all the properties in the model
     * return true if all passed
     *
     * @return bool True if all properties are valid
     */
    public function valid()
    {
        return count($this->listInvalidProperties()) === 0;
    }


    /**
     * Gets from_value
     *
     * @return int
     */
    public function getFromValue()
    {
        return $this->container['from_value'];
    }

    /**
     * Sets from_value
     *
     * @param int $from_value Specifies the lower value of a tier for a `graduated` charge model. It must be either 0 or the previous range's `to_value + 1` to maintain the proper sequence of values.
     *
     * @return self
     */
    public function setFromValue($from_value)
    {
        if (is_null($from_value)) {
            throw new \InvalidArgumentException('non-nullable from_value cannot be null');
        }
        $this->container['from_value'] = $from_value;

        return $this;
    }

    /**
     * Gets to_value
     *
     * @return int
     */
    public function getToValue()
    {
        return $this->container['to_value'];
    }

    /**
     * Sets to_value
     *
     * @param int $to_value Specifies the highest value of a tier for a `graduated` charge model. - This value must be higher than the from_value of the same tier. - This value must be null for the last tier.
     *
     * @return self
     */
    public function setToValue($to_value)
    {
        if (is_null($to_value)) {
            array_push($this->openAPINullablesSetToNull, 'to_value');
        } else {
            $nullablesSetToNull = $this->getOpenAPINullablesSetToNull();
            $index = array_search('to_value', $nullablesSetToNull);
            if ($index !== FALSE) {
                unset($nullablesSetToNull[$index]);
                $this->setOpenAPINullablesSetToNull($nullablesSetToNull);
            }
        }
        $this->container['to_value'] = $to_value;

        return $this;
    }

    /**
     * Gets flat_amount
     *
     * @return string
     */
    public function getFlatAmount()
    {
        return $this->container['flat_amount'];
    }

    /**
     * Sets flat_amount
     *
     * @param string $flat_amount The flat amount for a whole tier, excluding tax, for a `graduated` charge model. It is expressed as a decimal value.
     *
     * @return self
     */
    public function setFlatAmount($flat_amount)
    {
        if (is_null($flat_amount)) {
            throw new \InvalidArgumentException('non-nullable flat_amount cannot be null');
        }

        if ((!preg_match("/^[0-9]+.?[0-9]*$/", ObjectSerializer::toString($flat_amount)))) {
            throw new \InvalidArgumentException("invalid value for \$flat_amount when calling ChargePropertiesGraduatedRangesInner., must conform to the pattern /^[0-9]+.?[0-9]*$/.");
        }

        $this->container['flat_amount'] = $flat_amount;

        return $this;
    }

    /**
     * Gets per_unit_amount
     *
     * @return string
     */
    public function getPerUnitAmount()
    {
        return $this->container['per_unit_amount'];
    }

    /**
     * Sets per_unit_amount
     *
     * @param string $per_unit_amount The unit price, excluding tax, for a specific tier of a `graduated` charge model. It is expressed as a decimal value.
     *
     * @return self
     */
    public function setPerUnitAmount($per_unit_amount)
    {
        if (is_null($per_unit_amount)) {
            throw new \InvalidArgumentException('non-nullable per_unit_amount cannot be null');
        }

        if ((!preg_match("/^[0-9]+.?[0-9]*$/", ObjectSerializer::toString($per_unit_amount)))) {
            throw new \InvalidArgumentException("invalid value for \$per_unit_amount when calling ChargePropertiesGraduatedRangesInner., must conform to the pattern /^[0-9]+.?[0-9]*$/.");
        }

        $this->container['per_unit_amount'] = $per_unit_amount;

        return $this;
    }
    /**
     * Returns true if offset exists. False otherwise.
     *
     * @param integer $offset Offset
     *
     * @return boolean
     */
    public function offsetExists($offset): bool
    {
        return isset($this->container[$offset]);
    }

    /**
     * Gets offset.
     *
     * @param integer $offset Offset
     *
     * @return mixed|null
     */
    #[\ReturnTypeWillChange]
    public function offsetGet($offset)
    {
        return $this->container[$offset] ?? null;
    }

    /**
     * Sets value based on offset.
     *
     * @param int|null $offset Offset
     * @param mixed    $value  Value to be set
     *
     * @return void
     */
    public function offsetSet($offset, $value): void
    {
        if (is_null($offset)) {
            $this->container[] = $value;
        } else {
            $this->container[$offset] = $value;
        }
    }

    /**
     * Unsets offset.
     *
     * @param integer $offset Offset
     *
     * @return void
     */
    public function offsetUnset($offset): void
    {
        unset($this->container[$offset]);
    }

    /**
     * Serializes the object to a value that can be serialized natively by json_encode().
     * @link https://www.php.net/manual/en/jsonserializable.jsonserialize.php
     *
     * @return mixed Returns data which can be serialized by json_encode(), which is a value
     * of any type other than a resource.
     */
    #[\ReturnTypeWillChange]
    public function jsonSerialize()
    {
       return ObjectSerializer::sanitizeForSerialization($this);
    }

    /**
     * Gets the string presentation of the object
     *
     * @return string
     */
    public function __toString()
    {
        return json_encode(
            ObjectSerializer::sanitizeForSerialization($this),
            JSON_PRETTY_PRINT
        );
    }

    /**
     * Gets a header-safe presentation of the object
     *
     * @return string
     */
    public function toHeaderValue()
    {
        return json_encode(ObjectSerializer::sanitizeForSerialization($this));
    }
}


