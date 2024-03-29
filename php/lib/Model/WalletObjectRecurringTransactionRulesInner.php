<?php
/**
 * WalletObjectRecurringTransactionRulesInner
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
 * WalletObjectRecurringTransactionRulesInner Class Doc Comment
 *
 * @category Class
 * @description Object that represents rule for wallet recurring transactions.
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<string, mixed>
 */
class WalletObjectRecurringTransactionRulesInner implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'WalletObject_recurring_transaction_rules_inner';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'lago_id' => 'string',
        'rule_type' => 'string',
        'interval' => 'string',
        'threshold_credits' => 'string',
        'paid_credits' => 'string',
        'granted_credits' => 'string',
        'created_at' => '\DateTime'
    ];

    /**
      * Array of property to format mappings. Used for (de)serialization
      *
      * @var string[]
      * @phpstan-var array<string, string|null>
      * @psalm-var array<string, string|null>
      */
    protected static $openAPIFormats = [
        'lago_id' => 'uuid',
        'rule_type' => null,
        'interval' => null,
        'threshold_credits' => null,
        'paid_credits' => null,
        'granted_credits' => null,
        'created_at' => 'date-time'
    ];

    /**
      * Array of nullable properties. Used for (de)serialization
      *
      * @var boolean[]
      */
    protected static array $openAPINullables = [
        'lago_id' => false,
        'rule_type' => false,
        'interval' => false,
        'threshold_credits' => false,
        'paid_credits' => false,
        'granted_credits' => false,
        'created_at' => false
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
        'lago_id' => 'lago_id',
        'rule_type' => 'rule_type',
        'interval' => 'interval',
        'threshold_credits' => 'threshold_credits',
        'paid_credits' => 'paid_credits',
        'granted_credits' => 'granted_credits',
        'created_at' => 'created_at'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'lago_id' => 'setLagoId',
        'rule_type' => 'setRuleType',
        'interval' => 'setInterval',
        'threshold_credits' => 'setThresholdCredits',
        'paid_credits' => 'setPaidCredits',
        'granted_credits' => 'setGrantedCredits',
        'created_at' => 'setCreatedAt'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'lago_id' => 'getLagoId',
        'rule_type' => 'getRuleType',
        'interval' => 'getInterval',
        'threshold_credits' => 'getThresholdCredits',
        'paid_credits' => 'getPaidCredits',
        'granted_credits' => 'getGrantedCredits',
        'created_at' => 'getCreatedAt'
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

    public const RULE_TYPE_INTERVAL = 'interval';
    public const RULE_TYPE_THRESHOLD = 'threshold';
    public const INTERVAL_WEEKLY = 'weekly';
    public const INTERVAL_MONTHLY = 'monthly';
    public const INTERVAL_QUARTERLY = 'quarterly';
    public const INTERVAL_YEARLY = 'yearly';

    /**
     * Gets allowable values of the enum
     *
     * @return string[]
     */
    public function getRuleTypeAllowableValues()
    {
        return [
            self::RULE_TYPE_INTERVAL,
            self::RULE_TYPE_THRESHOLD,
        ];
    }

    /**
     * Gets allowable values of the enum
     *
     * @return string[]
     */
    public function getIntervalAllowableValues()
    {
        return [
            self::INTERVAL_WEEKLY,
            self::INTERVAL_MONTHLY,
            self::INTERVAL_QUARTERLY,
            self::INTERVAL_YEARLY,
        ];
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
        $this->setIfExists('lago_id', $data ?? [], null);
        $this->setIfExists('rule_type', $data ?? [], null);
        $this->setIfExists('interval', $data ?? [], null);
        $this->setIfExists('threshold_credits', $data ?? [], null);
        $this->setIfExists('paid_credits', $data ?? [], null);
        $this->setIfExists('granted_credits', $data ?? [], null);
        $this->setIfExists('created_at', $data ?? [], null);
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

        if ($this->container['lago_id'] === null) {
            $invalidProperties[] = "'lago_id' can't be null";
        }
        if ($this->container['rule_type'] === null) {
            $invalidProperties[] = "'rule_type' can't be null";
        }
        $allowedValues = $this->getRuleTypeAllowableValues();
        if (!is_null($this->container['rule_type']) && !in_array($this->container['rule_type'], $allowedValues, true)) {
            $invalidProperties[] = sprintf(
                "invalid value '%s' for 'rule_type', must be one of '%s'",
                $this->container['rule_type'],
                implode("', '", $allowedValues)
            );
        }

        $allowedValues = $this->getIntervalAllowableValues();
        if (!is_null($this->container['interval']) && !in_array($this->container['interval'], $allowedValues, true)) {
            $invalidProperties[] = sprintf(
                "invalid value '%s' for 'interval', must be one of '%s'",
                $this->container['interval'],
                implode("', '", $allowedValues)
            );
        }

        if (!is_null($this->container['threshold_credits']) && !preg_match("/^[0-9]+.?[0-9]*$/", $this->container['threshold_credits'])) {
            $invalidProperties[] = "invalid value for 'threshold_credits', must be conform to the pattern /^[0-9]+.?[0-9]*$/.";
        }

        if ($this->container['paid_credits'] === null) {
            $invalidProperties[] = "'paid_credits' can't be null";
        }
        if (!preg_match("/^[0-9]+.?[0-9]*$/", $this->container['paid_credits'])) {
            $invalidProperties[] = "invalid value for 'paid_credits', must be conform to the pattern /^[0-9]+.?[0-9]*$/.";
        }

        if ($this->container['granted_credits'] === null) {
            $invalidProperties[] = "'granted_credits' can't be null";
        }
        if (!preg_match("/^[0-9]+.?[0-9]*$/", $this->container['granted_credits'])) {
            $invalidProperties[] = "invalid value for 'granted_credits', must be conform to the pattern /^[0-9]+.?[0-9]*$/.";
        }

        if ($this->container['created_at'] === null) {
            $invalidProperties[] = "'created_at' can't be null";
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
     * Gets lago_id
     *
     * @return string
     */
    public function getLagoId()
    {
        return $this->container['lago_id'];
    }

    /**
     * Sets lago_id
     *
     * @param string $lago_id A unique identifier for the recurring transaction rule in the Lago application. Can be used to update a recurring transaction rule.
     *
     * @return self
     */
    public function setLagoId($lago_id)
    {
        if (is_null($lago_id)) {
            throw new \InvalidArgumentException('non-nullable lago_id cannot be null');
        }
        $this->container['lago_id'] = $lago_id;

        return $this;
    }

    /**
     * Gets rule_type
     *
     * @return string
     */
    public function getRuleType()
    {
        return $this->container['rule_type'];
    }

    /**
     * Sets rule_type
     *
     * @param string $rule_type The rule type. Possible values are `interval` or `threshold`.
     *
     * @return self
     */
    public function setRuleType($rule_type)
    {
        if (is_null($rule_type)) {
            throw new \InvalidArgumentException('non-nullable rule_type cannot be null');
        }
        $allowedValues = $this->getRuleTypeAllowableValues();
        if (!in_array($rule_type, $allowedValues, true)) {
            throw new \InvalidArgumentException(
                sprintf(
                    "Invalid value '%s' for 'rule_type', must be one of '%s'",
                    $rule_type,
                    implode("', '", $allowedValues)
                )
            );
        }
        $this->container['rule_type'] = $rule_type;

        return $this;
    }

    /**
     * Gets interval
     *
     * @return string|null
     */
    public function getInterval()
    {
        return $this->container['interval'];
    }

    /**
     * Sets interval
     *
     * @param string|null $interval The interval used for recurring top-up. It represents the frequency at which automatic top-up occurs. The interval can be one of the following values: `weekly`, `monthly`, `quarterly` or `yearly`. Required only if rule type is set to `interval`.
     *
     * @return self
     */
    public function setInterval($interval)
    {
        if (is_null($interval)) {
            throw new \InvalidArgumentException('non-nullable interval cannot be null');
        }
        $allowedValues = $this->getIntervalAllowableValues();
        if (!in_array($interval, $allowedValues, true)) {
            throw new \InvalidArgumentException(
                sprintf(
                    "Invalid value '%s' for 'interval', must be one of '%s'",
                    $interval,
                    implode("', '", $allowedValues)
                )
            );
        }
        $this->container['interval'] = $interval;

        return $this;
    }

    /**
     * Gets threshold_credits
     *
     * @return string|null
     */
    public function getThresholdCredits()
    {
        return $this->container['threshold_credits'];
    }

    /**
     * Sets threshold_credits
     *
     * @param string|null $threshold_credits The threshold for recurring top-ups is the value at which an automatic top-up is triggered. For example, if this threshold is set at 10 credits, an automatic top-up will occur whenever the wallet balance falls to or below 10 credits. Required only when rule type is set to `threshold`.
     *
     * @return self
     */
    public function setThresholdCredits($threshold_credits)
    {
        if (is_null($threshold_credits)) {
            throw new \InvalidArgumentException('non-nullable threshold_credits cannot be null');
        }

        if ((!preg_match("/^[0-9]+.?[0-9]*$/", ObjectSerializer::toString($threshold_credits)))) {
            throw new \InvalidArgumentException("invalid value for \$threshold_credits when calling WalletObjectRecurringTransactionRulesInner., must conform to the pattern /^[0-9]+.?[0-9]*$/.");
        }

        $this->container['threshold_credits'] = $threshold_credits;

        return $this;
    }

    /**
     * Gets paid_credits
     *
     * @return string
     */
    public function getPaidCredits()
    {
        return $this->container['paid_credits'];
    }

    /**
     * Sets paid_credits
     *
     * @param string $paid_credits The number of paid credits. Required only if there is no granted credits.
     *
     * @return self
     */
    public function setPaidCredits($paid_credits)
    {
        if (is_null($paid_credits)) {
            throw new \InvalidArgumentException('non-nullable paid_credits cannot be null');
        }

        if ((!preg_match("/^[0-9]+.?[0-9]*$/", ObjectSerializer::toString($paid_credits)))) {
            throw new \InvalidArgumentException("invalid value for \$paid_credits when calling WalletObjectRecurringTransactionRulesInner., must conform to the pattern /^[0-9]+.?[0-9]*$/.");
        }

        $this->container['paid_credits'] = $paid_credits;

        return $this;
    }

    /**
     * Gets granted_credits
     *
     * @return string
     */
    public function getGrantedCredits()
    {
        return $this->container['granted_credits'];
    }

    /**
     * Sets granted_credits
     *
     * @param string $granted_credits The number of free granted credits. Required only if there is no paid credits.
     *
     * @return self
     */
    public function setGrantedCredits($granted_credits)
    {
        if (is_null($granted_credits)) {
            throw new \InvalidArgumentException('non-nullable granted_credits cannot be null');
        }

        if ((!preg_match("/^[0-9]+.?[0-9]*$/", ObjectSerializer::toString($granted_credits)))) {
            throw new \InvalidArgumentException("invalid value for \$granted_credits when calling WalletObjectRecurringTransactionRulesInner., must conform to the pattern /^[0-9]+.?[0-9]*$/.");
        }

        $this->container['granted_credits'] = $granted_credits;

        return $this;
    }

    /**
     * Gets created_at
     *
     * @return \DateTime
     */
    public function getCreatedAt()
    {
        return $this->container['created_at'];
    }

    /**
     * Sets created_at
     *
     * @param \DateTime $created_at The date of the metadata object creation, represented in ISO 8601 datetime format and expressed in Coordinated Universal Time (UTC). The creation_date provides a standardized and internationally recognized timestamp for when the metadata object was created
     *
     * @return self
     */
    public function setCreatedAt($created_at)
    {
        if (is_null($created_at)) {
            throw new \InvalidArgumentException('non-nullable created_at cannot be null');
        }
        $this->container['created_at'] = $created_at;

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


