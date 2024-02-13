<?php
/**
 * ChargeObject
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
 * ChargeObject Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<string, mixed>
 */
class ChargeObject implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'ChargeObject';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'lago_id' => 'string',
        'lago_billable_metric_id' => 'string',
        'billable_metric_code' => 'string',
        'invoice_display_name' => 'string',
        'created_at' => '\DateTime',
        'charge_model' => 'string',
        'pay_in_advance' => 'bool',
        'invoiceable' => 'bool',
        'prorated' => 'bool',
        'min_amount_cents' => 'int',
        'properties' => '\OpenAPI\Client\Model\ChargeObjectProperties',
        'group_properties' => '\OpenAPI\Client\Model\GroupPropertiesObject[]',
        'taxes' => '\OpenAPI\Client\Model\TaxObject[]'
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
        'lago_billable_metric_id' => 'uuid',
        'billable_metric_code' => null,
        'invoice_display_name' => null,
        'created_at' => 'date-time',
        'charge_model' => null,
        'pay_in_advance' => null,
        'invoiceable' => null,
        'prorated' => null,
        'min_amount_cents' => null,
        'properties' => null,
        'group_properties' => null,
        'taxes' => null
    ];

    /**
      * Array of nullable properties. Used for (de)serialization
      *
      * @var boolean[]
      */
    protected static array $openAPINullables = [
        'lago_id' => false,
        'lago_billable_metric_id' => false,
        'billable_metric_code' => false,
        'invoice_display_name' => false,
        'created_at' => false,
        'charge_model' => false,
        'pay_in_advance' => false,
        'invoiceable' => false,
        'prorated' => false,
        'min_amount_cents' => false,
        'properties' => false,
        'group_properties' => false,
        'taxes' => false
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
        'lago_billable_metric_id' => 'lago_billable_metric_id',
        'billable_metric_code' => 'billable_metric_code',
        'invoice_display_name' => 'invoice_display_name',
        'created_at' => 'created_at',
        'charge_model' => 'charge_model',
        'pay_in_advance' => 'pay_in_advance',
        'invoiceable' => 'invoiceable',
        'prorated' => 'prorated',
        'min_amount_cents' => 'min_amount_cents',
        'properties' => 'properties',
        'group_properties' => 'group_properties',
        'taxes' => 'taxes'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'lago_id' => 'setLagoId',
        'lago_billable_metric_id' => 'setLagoBillableMetricId',
        'billable_metric_code' => 'setBillableMetricCode',
        'invoice_display_name' => 'setInvoiceDisplayName',
        'created_at' => 'setCreatedAt',
        'charge_model' => 'setChargeModel',
        'pay_in_advance' => 'setPayInAdvance',
        'invoiceable' => 'setInvoiceable',
        'prorated' => 'setProrated',
        'min_amount_cents' => 'setMinAmountCents',
        'properties' => 'setProperties',
        'group_properties' => 'setGroupProperties',
        'taxes' => 'setTaxes'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'lago_id' => 'getLagoId',
        'lago_billable_metric_id' => 'getLagoBillableMetricId',
        'billable_metric_code' => 'getBillableMetricCode',
        'invoice_display_name' => 'getInvoiceDisplayName',
        'created_at' => 'getCreatedAt',
        'charge_model' => 'getChargeModel',
        'pay_in_advance' => 'getPayInAdvance',
        'invoiceable' => 'getInvoiceable',
        'prorated' => 'getProrated',
        'min_amount_cents' => 'getMinAmountCents',
        'properties' => 'getProperties',
        'group_properties' => 'getGroupProperties',
        'taxes' => 'getTaxes'
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

    public const CHARGE_MODEL_STANDARD = 'standard';
    public const CHARGE_MODEL_GRADUATED = 'graduated';
    public const CHARGE_MODEL_GRADUATED_PERCENTAGE = 'graduated_percentage';
    public const CHARGE_MODEL_PACKAGE = 'package';
    public const CHARGE_MODEL_PERCENTAGE = 'percentage';
    public const CHARGE_MODEL_VOLUME = 'volume';

    /**
     * Gets allowable values of the enum
     *
     * @return string[]
     */
    public function getChargeModelAllowableValues()
    {
        return [
            self::CHARGE_MODEL_STANDARD,
            self::CHARGE_MODEL_GRADUATED,
            self::CHARGE_MODEL_GRADUATED_PERCENTAGE,
            self::CHARGE_MODEL_PACKAGE,
            self::CHARGE_MODEL_PERCENTAGE,
            self::CHARGE_MODEL_VOLUME,
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
        $this->setIfExists('lago_billable_metric_id', $data ?? [], null);
        $this->setIfExists('billable_metric_code', $data ?? [], null);
        $this->setIfExists('invoice_display_name', $data ?? [], null);
        $this->setIfExists('created_at', $data ?? [], null);
        $this->setIfExists('charge_model', $data ?? [], null);
        $this->setIfExists('pay_in_advance', $data ?? [], null);
        $this->setIfExists('invoiceable', $data ?? [], null);
        $this->setIfExists('prorated', $data ?? [], null);
        $this->setIfExists('min_amount_cents', $data ?? [], null);
        $this->setIfExists('properties', $data ?? [], null);
        $this->setIfExists('group_properties', $data ?? [], null);
        $this->setIfExists('taxes', $data ?? [], null);
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
        if ($this->container['lago_billable_metric_id'] === null) {
            $invalidProperties[] = "'lago_billable_metric_id' can't be null";
        }
        if ($this->container['billable_metric_code'] === null) {
            $invalidProperties[] = "'billable_metric_code' can't be null";
        }
        if ($this->container['created_at'] === null) {
            $invalidProperties[] = "'created_at' can't be null";
        }
        if ($this->container['charge_model'] === null) {
            $invalidProperties[] = "'charge_model' can't be null";
        }
        $allowedValues = $this->getChargeModelAllowableValues();
        if (!is_null($this->container['charge_model']) && !in_array($this->container['charge_model'], $allowedValues, true)) {
            $invalidProperties[] = sprintf(
                "invalid value '%s' for 'charge_model', must be one of '%s'",
                $this->container['charge_model'],
                implode("', '", $allowedValues)
            );
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
     * @param string $lago_id Unique identifier of charge, created by Lago.
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
     * Gets lago_billable_metric_id
     *
     * @return string
     */
    public function getLagoBillableMetricId()
    {
        return $this->container['lago_billable_metric_id'];
    }

    /**
     * Sets lago_billable_metric_id
     *
     * @param string $lago_billable_metric_id Unique identifier of the billable metric created by Lago.
     *
     * @return self
     */
    public function setLagoBillableMetricId($lago_billable_metric_id)
    {
        if (is_null($lago_billable_metric_id)) {
            throw new \InvalidArgumentException('non-nullable lago_billable_metric_id cannot be null');
        }
        $this->container['lago_billable_metric_id'] = $lago_billable_metric_id;

        return $this;
    }

    /**
     * Gets billable_metric_code
     *
     * @return string
     */
    public function getBillableMetricCode()
    {
        return $this->container['billable_metric_code'];
    }

    /**
     * Sets billable_metric_code
     *
     * @param string $billable_metric_code Unique code identifying a billable metric.
     *
     * @return self
     */
    public function setBillableMetricCode($billable_metric_code)
    {
        if (is_null($billable_metric_code)) {
            throw new \InvalidArgumentException('non-nullable billable_metric_code cannot be null');
        }
        $this->container['billable_metric_code'] = $billable_metric_code;

        return $this;
    }

    /**
     * Gets invoice_display_name
     *
     * @return string|null
     */
    public function getInvoiceDisplayName()
    {
        return $this->container['invoice_display_name'];
    }

    /**
     * Sets invoice_display_name
     *
     * @param string|null $invoice_display_name Specifies the name that will be displayed on an invoice. If no value is set for this field, the name of the actual charge will be used as the default display name.
     *
     * @return self
     */
    public function setInvoiceDisplayName($invoice_display_name)
    {
        if (is_null($invoice_display_name)) {
            throw new \InvalidArgumentException('non-nullable invoice_display_name cannot be null');
        }
        $this->container['invoice_display_name'] = $invoice_display_name;

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
     * @param \DateTime $created_at The date and time when the charge was created. It is expressed in UTC format according to the ISO 8601 datetime standard.
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
     * Gets charge_model
     *
     * @return string
     */
    public function getChargeModel()
    {
        return $this->container['charge_model'];
    }

    /**
     * Sets charge_model
     *
     * @param string $charge_model Specifies the pricing model used for the calculation of the final fee. It can be `standard`, `graduated`, `graduated_percentage`, `package`, `percentage` or `volume`.
     *
     * @return self
     */
    public function setChargeModel($charge_model)
    {
        if (is_null($charge_model)) {
            throw new \InvalidArgumentException('non-nullable charge_model cannot be null');
        }
        $allowedValues = $this->getChargeModelAllowableValues();
        if (!in_array($charge_model, $allowedValues, true)) {
            throw new \InvalidArgumentException(
                sprintf(
                    "Invalid value '%s' for 'charge_model', must be one of '%s'",
                    $charge_model,
                    implode("', '", $allowedValues)
                )
            );
        }
        $this->container['charge_model'] = $charge_model;

        return $this;
    }

    /**
     * Gets pay_in_advance
     *
     * @return bool|null
     */
    public function getPayInAdvance()
    {
        return $this->container['pay_in_advance'];
    }

    /**
     * Sets pay_in_advance
     *
     * @param bool|null $pay_in_advance This field determines the billing timing for this specific usage-based charge. When set to `true`, the charge is due and invoiced immediately. Conversely, when set to `false`, the charge is due and invoiced at the end of each billing period.
     *
     * @return self
     */
    public function setPayInAdvance($pay_in_advance)
    {
        if (is_null($pay_in_advance)) {
            throw new \InvalidArgumentException('non-nullable pay_in_advance cannot be null');
        }
        $this->container['pay_in_advance'] = $pay_in_advance;

        return $this;
    }

    /**
     * Gets invoiceable
     *
     * @return bool|null
     */
    public function getInvoiceable()
    {
        return $this->container['invoiceable'];
    }

    /**
     * Sets invoiceable
     *
     * @param bool|null $invoiceable This field specifies whether the charge should be included in a proper invoice. If set to `false`, no invoice will be issued for this charge. You can only set it to `false` when `pay_in_advance` is `true`.
     *
     * @return self
     */
    public function setInvoiceable($invoiceable)
    {
        if (is_null($invoiceable)) {
            throw new \InvalidArgumentException('non-nullable invoiceable cannot be null');
        }
        $this->container['invoiceable'] = $invoiceable;

        return $this;
    }

    /**
     * Gets prorated
     *
     * @return bool|null
     */
    public function getProrated()
    {
        return $this->container['prorated'];
    }

    /**
     * Sets prorated
     *
     * @param bool|null $prorated Specifies whether a charge is prorated based on the remaining number of days in the billing period or billed fully.  - If set to `true`, the charge is prorated based on the remaining days in the current billing period. - If set to `false`, the charge is billed in full. - If not defined in the request, default value is `false`.
     *
     * @return self
     */
    public function setProrated($prorated)
    {
        if (is_null($prorated)) {
            throw new \InvalidArgumentException('non-nullable prorated cannot be null');
        }
        $this->container['prorated'] = $prorated;

        return $this;
    }

    /**
     * Gets min_amount_cents
     *
     * @return int|null
     */
    public function getMinAmountCents()
    {
        return $this->container['min_amount_cents'];
    }

    /**
     * Sets min_amount_cents
     *
     * @param int|null $min_amount_cents The minimum spending amount required for the charge, measured in cents and excluding any applicable taxes. It indicates the minimum amount that needs to be charged for each billing period.
     *
     * @return self
     */
    public function setMinAmountCents($min_amount_cents)
    {
        if (is_null($min_amount_cents)) {
            throw new \InvalidArgumentException('non-nullable min_amount_cents cannot be null');
        }
        $this->container['min_amount_cents'] = $min_amount_cents;

        return $this;
    }

    /**
     * Gets properties
     *
     * @return \OpenAPI\Client\Model\ChargeObjectProperties|null
     */
    public function getProperties()
    {
        return $this->container['properties'];
    }

    /**
     * Sets properties
     *
     * @param \OpenAPI\Client\Model\ChargeObjectProperties|null $properties properties
     *
     * @return self
     */
    public function setProperties($properties)
    {
        if (is_null($properties)) {
            throw new \InvalidArgumentException('non-nullable properties cannot be null');
        }
        $this->container['properties'] = $properties;

        return $this;
    }

    /**
     * Gets group_properties
     *
     * @return \OpenAPI\Client\Model\GroupPropertiesObject[]|null
     */
    public function getGroupProperties()
    {
        return $this->container['group_properties'];
    }

    /**
     * Sets group_properties
     *
     * @param \OpenAPI\Client\Model\GroupPropertiesObject[]|null $group_properties All charge information, sorted by groups.
     *
     * @return self
     */
    public function setGroupProperties($group_properties)
    {
        if (is_null($group_properties)) {
            throw new \InvalidArgumentException('non-nullable group_properties cannot be null');
        }
        $this->container['group_properties'] = $group_properties;

        return $this;
    }

    /**
     * Gets taxes
     *
     * @return \OpenAPI\Client\Model\TaxObject[]|null
     */
    public function getTaxes()
    {
        return $this->container['taxes'];
    }

    /**
     * Sets taxes
     *
     * @param \OpenAPI\Client\Model\TaxObject[]|null $taxes All taxes applied to the charge.
     *
     * @return self
     */
    public function setTaxes($taxes)
    {
        if (is_null($taxes)) {
            throw new \InvalidArgumentException('non-nullable taxes cannot be null');
        }
        $this->container['taxes'] = $taxes;

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

