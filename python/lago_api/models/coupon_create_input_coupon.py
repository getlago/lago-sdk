# coding: utf-8

"""
    Lago API documentation

    Lago API allows your application to push customer information and metrics (events) from your application to the billing application.

    The version of the OpenAPI document: 0.52.0-beta
    Contact: tech@getlago.com
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


from __future__ import annotations
import pprint
import re  # noqa: F401
import json

from datetime import datetime
from typing import Any, ClassVar, Dict, List, Optional
from pydantic import BaseModel, StrictBool, StrictInt, StrictStr, field_validator
from pydantic import Field
from typing_extensions import Annotated
from lago_api.models.coupon_base_input_applies_to import CouponBaseInputAppliesTo
from lago_api.models.currency import Currency
try:
    from typing import Self
except ImportError:
    from typing_extensions import Self

class CouponCreateInputCoupon(BaseModel):
    """
    CouponCreateInputCoupon
    """ # noqa: E501
    name: StrictStr = Field(description="The name of the coupon.")
    code: StrictStr = Field(description="Unique code used to identify the coupon.")
    description: Optional[StrictStr] = Field(default=None, description="Description of the coupon.")
    coupon_type: StrictStr = Field(description="The type of the coupon. It can have two possible values: `fixed_amount` or `percentage`.  - If set to `fixed_amount`, the coupon represents a fixed amount discount. - If set to `percentage`, the coupon represents a percentage-based discount.")
    amount_cents: Optional[StrictInt] = Field(default=None, description="The amount of the coupon in cents. This field is required only for coupon with `fixed_amount` type.")
    amount_currency: Optional[Any] = None
    reusable: Optional[StrictBool] = Field(default=None, description="Indicates whether the coupon can be reused or not. If set to `true`, the coupon is reusable, meaning it can be applied multiple times to the same customer. If set to `false`, the coupon can only be used once and is not reusable. If not specified, this field is set to `true` by default.")
    percentage_rate: Optional[Annotated[str, Field(strict=True)]] = Field(default=None, description="The percentage rate of the coupon. This field is required only for coupons with a `percentage` coupon type.")
    frequency: StrictStr = Field(description="The type of frequency for the coupon. It can have three possible values: `once`, `recurring` or `forever`.  - If set to `once`, the coupon is applicable only for a single use. - If set to `recurring`, the coupon can be used multiple times for recurring billing periods. - If set to `forever`, the coupon has unlimited usage and can be applied indefinitely.")
    frequency_duration: Optional[StrictInt] = Field(default=None, description="Specifies the number of billing periods to which the coupon applies. This field is required only for coupons with a `recurring` frequency type")
    expiration: Optional[StrictStr] = Field(default=None, description="Specifies the type of expiration for the coupon. It can have two possible values: `time_limit` or `no_expiration`.  - If set to `time_limit`, the coupon has an expiration based on a specified time limit. - If set to `no_expiration`, the coupon does not have an expiration date and remains valid indefinitely.")
    expiration_at: Optional[datetime] = Field(default=None, description="The expiration date and time of the coupon. This field is required only for coupons with `expiration` set to `time_limit`. The expiration date and time should be specified in UTC format according to the ISO 8601 datetime standard. It indicates the exact moment when the coupon will expire and is no longer valid.")
    applies_to: Optional[CouponBaseInputAppliesTo] = None
    __properties: ClassVar[List[str]] = ["name", "code", "description", "coupon_type", "amount_cents", "amount_currency", "reusable", "percentage_rate", "frequency", "frequency_duration", "expiration", "expiration_at", "applies_to"]

    @field_validator('coupon_type')
    def coupon_type_validate_enum(cls, value):
        """Validates the enum"""
        if value not in ('fixed_amount', 'percentage'):
            raise ValueError("must be one of enum values ('fixed_amount', 'percentage')")
        return value

    @field_validator('percentage_rate')
    def percentage_rate_validate_regular_expression(cls, value):
        """Validates the regular expression"""
        if value is None:
            return value

        if not re.match(r"^[0-9]+.?[0-9]*$", value):
            raise ValueError(r"must validate the regular expression /^[0-9]+.?[0-9]*$/")
        return value

    @field_validator('frequency')
    def frequency_validate_enum(cls, value):
        """Validates the enum"""
        if value not in ('once', 'recurring'):
            raise ValueError("must be one of enum values ('once', 'recurring')")
        return value

    @field_validator('expiration')
    def expiration_validate_enum(cls, value):
        """Validates the enum"""
        if value is None:
            return value

        if value not in ('no_expiration', 'time_limit'):
            raise ValueError("must be one of enum values ('no_expiration', 'time_limit')")
        return value

    model_config = {
        "populate_by_name": True,
        "validate_assignment": True,
        "protected_namespaces": (),
    }


    def to_str(self) -> str:
        """Returns the string representation of the model using alias"""
        return pprint.pformat(self.model_dump(by_alias=True))

    def to_json(self) -> str:
        """Returns the JSON representation of the model using alias"""
        # TODO: pydantic v2: use .model_dump_json(by_alias=True, exclude_unset=True) instead
        return json.dumps(self.to_dict())

    @classmethod
    def from_json(cls, json_str: str) -> Self:
        """Create an instance of CouponCreateInputCoupon from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self) -> Dict[str, Any]:
        """Return the dictionary representation of the model using alias.

        This has the following differences from calling pydantic's
        `self.model_dump(by_alias=True)`:

        * `None` is only added to the output dict for nullable fields that
          were set at model initialization. Other fields with value `None`
          are ignored.
        """
        _dict = self.model_dump(
            by_alias=True,
            exclude={
            },
            exclude_none=True,
        )
        # override the default output from pydantic by calling `to_dict()` of amount_currency
        if self.amount_currency:
            _dict['amount_currency'] = self.amount_currency.to_dict()
        # override the default output from pydantic by calling `to_dict()` of applies_to
        if self.applies_to:
            _dict['applies_to'] = self.applies_to.to_dict()
        # set to None if description (nullable) is None
        # and model_fields_set contains the field
        if self.description is None and "description" in self.model_fields_set:
            _dict['description'] = None

        # set to None if amount_cents (nullable) is None
        # and model_fields_set contains the field
        if self.amount_cents is None and "amount_cents" in self.model_fields_set:
            _dict['amount_cents'] = None

        # set to None if percentage_rate (nullable) is None
        # and model_fields_set contains the field
        if self.percentage_rate is None and "percentage_rate" in self.model_fields_set:
            _dict['percentage_rate'] = None

        # set to None if frequency_duration (nullable) is None
        # and model_fields_set contains the field
        if self.frequency_duration is None and "frequency_duration" in self.model_fields_set:
            _dict['frequency_duration'] = None

        # set to None if expiration_at (nullable) is None
        # and model_fields_set contains the field
        if self.expiration_at is None and "expiration_at" in self.model_fields_set:
            _dict['expiration_at'] = None

        # set to None if applies_to (nullable) is None
        # and model_fields_set contains the field
        if self.applies_to is None and "applies_to" in self.model_fields_set:
            _dict['applies_to'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Dict) -> Self:
        """Create an instance of CouponCreateInputCoupon from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "name": obj.get("name"),
            "code": obj.get("code"),
            "description": obj.get("description"),
            "coupon_type": obj.get("coupon_type"),
            "amount_cents": obj.get("amount_cents"),
            "amount_currency": Currency.from_dict(obj.get("amount_currency")) if obj.get("amount_currency") is not None else None,
            "reusable": obj.get("reusable"),
            "percentage_rate": obj.get("percentage_rate"),
            "frequency": obj.get("frequency"),
            "frequency_duration": obj.get("frequency_duration"),
            "expiration": obj.get("expiration"),
            "expiration_at": obj.get("expiration_at"),
            "applies_to": CouponBaseInputAppliesTo.from_dict(obj.get("applies_to")) if obj.get("applies_to") is not None else None
        })
        return _obj


