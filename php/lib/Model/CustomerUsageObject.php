<?php
/**
 * CustomerUsageObject
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
 * CustomerUsageObject Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<string, mixed>
 */
class CustomerUsageObject implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'CustomerUsageObject';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'from_datetime' => '\DateTime',
        'to_datetime' => '\DateTime',
        'issuing_date' => '\DateTime',
        'lago_invoice_id' => 'string',
        'currency' => 'Currency',
        'amount_cents' => 'int',
        'taxes_amount_cents' => 'int',
        'total_amount_cents' => 'int',
        'charges_usage' => '\OpenAPI\Client\Model\CustomerChargeUsageObject[]'
    ];

    /**
      * Array of property to format mappings. Used for (de)serialization
      *
      * @var string[]
      * @phpstan-var array<string, string|null>
      * @psalm-var array<string, string|null>
      */
    protected static $openAPIFormats = [
        'from_datetime' => 'date-time',
        'to_datetime' => 'date-time',
        'issuing_date' => 'date-time',
        'lago_invoice_id' => 'uuid',
        'currency' => null,
        'amount_cents' => null,
        'taxes_amount_cents' => null,
        'total_amount_cents' => null,
        'charges_usage' => null
    ];

    /**
      * Array of nullable properties. Used for (de)serialization
      *
      * @var boolean[]
      */
    protected static array $openAPINullables = [
        'from_datetime' => false,
        'to_datetime' => false,
        'issuing_date' => false,
        'lago_invoice_id' => true,
        'currency' => false,
        'amount_cents' => false,
        'taxes_amount_cents' => false,
        'total_amount_cents' => false,
        'charges_usage' => false
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
        'from_datetime' => 'from_datetime',
        'to_datetime' => 'to_datetime',
        'issuing_date' => 'issuing_date',
        'lago_invoice_id' => 'lago_invoice_id',
        'currency' => 'currency',
        'amount_cents' => 'amount_cents',
        'taxes_amount_cents' => 'taxes_amount_cents',
        'total_amount_cents' => 'total_amount_cents',
        'charges_usage' => 'charges_usage'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'from_datetime' => 'setFromDatetime',
        'to_datetime' => 'setToDatetime',
        'issuing_date' => 'setIssuingDate',
        'lago_invoice_id' => 'setLagoInvoiceId',
        'currency' => 'setCurrency',
        'amount_cents' => 'setAmountCents',
        'taxes_amount_cents' => 'setTaxesAmountCents',
        'total_amount_cents' => 'setTotalAmountCents',
        'charges_usage' => 'setChargesUsage'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'from_datetime' => 'getFromDatetime',
        'to_datetime' => 'getToDatetime',
        'issuing_date' => 'getIssuingDate',
        'lago_invoice_id' => 'getLagoInvoiceId',
        'currency' => 'getCurrency',
        'amount_cents' => 'getAmountCents',
        'taxes_amount_cents' => 'getTaxesAmountCents',
        'total_amount_cents' => 'getTotalAmountCents',
        'charges_usage' => 'getChargesUsage'
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
        $this->setIfExists('from_datetime', $data ?? [], null);
        $this->setIfExists('to_datetime', $data ?? [], null);
        $this->setIfExists('issuing_date', $data ?? [], null);
        $this->setIfExists('lago_invoice_id', $data ?? [], null);
        $this->setIfExists('currency', $data ?? [], null);
        $this->setIfExists('amount_cents', $data ?? [], null);
        $this->setIfExists('taxes_amount_cents', $data ?? [], null);
        $this->setIfExists('total_amount_cents', $data ?? [], null);
        $this->setIfExists('charges_usage', $data ?? [], null);
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

        if ($this->container['from_datetime'] === null) {
            $invalidProperties[] = "'from_datetime' can't be null";
        }
        if ($this->container['to_datetime'] === null) {
            $invalidProperties[] = "'to_datetime' can't be null";
        }
        if ($this->container['issuing_date'] === null) {
            $invalidProperties[] = "'issuing_date' can't be null";
        }
        if ($this->container['amount_cents'] === null) {
            $invalidProperties[] = "'amount_cents' can't be null";
        }
        if ($this->container['taxes_amount_cents'] === null) {
            $invalidProperties[] = "'taxes_amount_cents' can't be null";
        }
        if ($this->container['total_amount_cents'] === null) {
            $invalidProperties[] = "'total_amount_cents' can't be null";
        }
        if ($this->container['charges_usage'] === null) {
            $invalidProperties[] = "'charges_usage' can't be null";
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
     * Gets from_datetime
     *
     * @return \DateTime
     */
    public function getFromDatetime()
    {
        return $this->container['from_datetime'];
    }

    /**
     * Sets from_datetime
     *
     * @param \DateTime $from_datetime The lower bound of the billing period, expressed in the ISO 8601 datetime format in Coordinated Universal Time (UTC).
     *
     * @return self
     */
    public function setFromDatetime($from_datetime)
    {
        if (is_null($from_datetime)) {
            throw new \InvalidArgumentException('non-nullable from_datetime cannot be null');
        }
        $this->container['from_datetime'] = $from_datetime;

        return $this;
    }

    /**
     * Gets to_datetime
     *
     * @return \DateTime
     */
    public function getToDatetime()
    {
        return $this->container['to_datetime'];
    }

    /**
     * Sets to_datetime
     *
     * @param \DateTime $to_datetime The upper bound of the billing period, expressed in the ISO 8601 datetime format in Coordinated Universal Time (UTC).
     *
     * @return self
     */
    public function setToDatetime($to_datetime)
    {
        if (is_null($to_datetime)) {
            throw new \InvalidArgumentException('non-nullable to_datetime cannot be null');
        }
        $this->container['to_datetime'] = $to_datetime;

        return $this;
    }

    /**
     * Gets issuing_date
     *
     * @return \DateTime
     */
    public function getIssuingDate()
    {
        return $this->container['issuing_date'];
    }

    /**
     * Sets issuing_date
     *
     * @param \DateTime $issuing_date The date of creation of the invoice.
     *
     * @return self
     */
    public function setIssuingDate($issuing_date)
    {
        if (is_null($issuing_date)) {
            throw new \InvalidArgumentException('non-nullable issuing_date cannot be null');
        }
        $this->container['issuing_date'] = $issuing_date;

        return $this;
    }

    /**
     * Gets lago_invoice_id
     *
     * @return string|null
     */
    public function getLagoInvoiceId()
    {
        return $this->container['lago_invoice_id'];
    }

    /**
     * Sets lago_invoice_id
     *
     * @param string|null $lago_invoice_id A unique identifier associated with the invoice related to this particular usage record.
     *
     * @return self
     */
    public function setLagoInvoiceId($lago_invoice_id)
    {
        if (is_null($lago_invoice_id)) {
            array_push($this->openAPINullablesSetToNull, 'lago_invoice_id');
        } else {
            $nullablesSetToNull = $this->getOpenAPINullablesSetToNull();
            $index = array_search('lago_invoice_id', $nullablesSetToNull);
            if ($index !== FALSE) {
                unset($nullablesSetToNull[$index]);
                $this->setOpenAPINullablesSetToNull($nullablesSetToNull);
            }
        }
        $this->container['lago_invoice_id'] = $lago_invoice_id;

        return $this;
    }

    /**
     * Gets currency
     *
     * @return Currency|null
     */
    public function getCurrency()
    {
        return $this->container['currency'];
    }

    /**
     * Sets currency
     *
     * @param Currency|null $currency currency
     *
     * @return self
     */
    public function setCurrency($currency)
    {
        if (is_null($currency)) {
            throw new \InvalidArgumentException('non-nullable currency cannot be null');
        }
        $this->container['currency'] = $currency;

        return $this;
    }

    /**
     * Gets amount_cents
     *
     * @return int
     */
    public function getAmountCents()
    {
        return $this->container['amount_cents'];
    }

    /**
     * Sets amount_cents
     *
     * @param int $amount_cents The amount in cents, tax excluded.
     *
     * @return self
     */
    public function setAmountCents($amount_cents)
    {
        if (is_null($amount_cents)) {
            throw new \InvalidArgumentException('non-nullable amount_cents cannot be null');
        }
        $this->container['amount_cents'] = $amount_cents;

        return $this;
    }

    /**
     * Gets taxes_amount_cents
     *
     * @return int
     */
    public function getTaxesAmountCents()
    {
        return $this->container['taxes_amount_cents'];
    }

    /**
     * Sets taxes_amount_cents
     *
     * @param int $taxes_amount_cents The tax amount in cents.
     *
     * @return self
     */
    public function setTaxesAmountCents($taxes_amount_cents)
    {
        if (is_null($taxes_amount_cents)) {
            throw new \InvalidArgumentException('non-nullable taxes_amount_cents cannot be null');
        }
        $this->container['taxes_amount_cents'] = $taxes_amount_cents;

        return $this;
    }

    /**
     * Gets total_amount_cents
     *
     * @return int
     */
    public function getTotalAmountCents()
    {
        return $this->container['total_amount_cents'];
    }

    /**
     * Sets total_amount_cents
     *
     * @param int $total_amount_cents The total amount in cents, tax included.
     *
     * @return self
     */
    public function setTotalAmountCents($total_amount_cents)
    {
        if (is_null($total_amount_cents)) {
            throw new \InvalidArgumentException('non-nullable total_amount_cents cannot be null');
        }
        $this->container['total_amount_cents'] = $total_amount_cents;

        return $this;
    }

    /**
     * Gets charges_usage
     *
     * @return \OpenAPI\Client\Model\CustomerChargeUsageObject[]
     */
    public function getChargesUsage()
    {
        return $this->container['charges_usage'];
    }

    /**
     * Sets charges_usage
     *
     * @param \OpenAPI\Client\Model\CustomerChargeUsageObject[] $charges_usage Array of charges that comprise the current usage. It contains detailed information about individual charge items associated with the usage.
     *
     * @return self
     */
    public function setChargesUsage($charges_usage)
    {
        if (is_null($charges_usage)) {
            throw new \InvalidArgumentException('non-nullable charges_usage cannot be null');
        }
        $this->container['charges_usage'] = $charges_usage;

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


