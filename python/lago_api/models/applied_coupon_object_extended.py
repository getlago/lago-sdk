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
from pydantic import BaseModel, StrictInt, StrictStr, field_validator
from pydantic import Field
from typing_extensions import Annotated
from lago_api.models.credit_object import CreditObject
from lago_api.models.currency import Currency
try:
    from typing import Self
except ImportError:
    from typing_extensions import Self

class AppliedCouponObjectExtended(BaseModel):
    """
    AppliedCouponObjectExtended
    """ # noqa: E501
    lago_id: StrictStr = Field(description="Unique identifier of the applied coupon, created by Lago.")
    lago_coupon_id: StrictStr = Field(description="Unique identifier of the coupon, created by Lago.")
    coupon_code: StrictStr = Field(description="Unique code used to identify the coupon.")
    coupon_name: StrictStr = Field(description="The name of the coupon.")
    lago_customer_id: StrictStr = Field(description="Unique identifier of the customer, created by Lago.")
    external_customer_id: StrictStr = Field(description="The customer external unique identifier (provided by your own application)")
    status: StrictStr = Field(description="The status of the coupon. Can be either `active` or `terminated`.")
    amount_cents: Optional[StrictInt] = Field(default=None, description="The amount of the coupon in cents. This field is required only for coupon with `fixed_amount` type.")
    amount_cents_remaining: Optional[StrictInt] = Field(default=None, description="The remaining amount in cents for a `fixed_amount` coupon with a frequency set to `once`. This field indicates the remaining balance or value that can still be utilized from the coupon.")
    amount_currency: Optional[Any] = None
    percentage_rate: Optional[Annotated[str, Field(strict=True)]] = Field(default=None, description="The percentage rate of the coupon. This field is required only for coupons with a `percentage` coupon type.")
    frequency: StrictStr = Field(description="The type of frequency for the coupon. It can have three possible values: `once`, `recurring` or `forever`.  - If set to `once`, the coupon is applicable only for a single use. - If set to `recurring`, the coupon can be used multiple times for recurring billing periods. - If set to `forever`, the coupon has unlimited usage and can be applied indefinitely.")
    frequency_duration: Optional[StrictInt] = Field(default=None, description="Specifies the number of billing periods to which the coupon applies. This field is required only for coupons with a `recurring` frequency type")
    frequency_duration_remaining: Optional[StrictInt] = Field(default=None, description="The remaining number of billing periods to which the coupon is applicable. This field determines the remaining usage or availability of the coupon based on the remaining billing periods.")
    expiration_at: Optional[datetime] = Field(default=None, description="The date and time after which the coupon will stop applying to customer's invoices. Once the expiration date is reached, the coupon will no longer be applicable, and any further invoices generated for the customer will not include the coupon discount.")
    created_at: datetime = Field(description="The date and time when the coupon was assigned to a customer. It is expressed in UTC format according to the ISO 8601 datetime standard.")
    terminated_at: Optional[datetime] = Field(default=None, description="This field indicates the specific moment when the coupon amount is fully utilized or when the coupon is removed from the customer's coupon list. It is expressed in UTC format according to the ISO 8601 datetime standard.")
    credits: List[CreditObject]
    __properties: ClassVar[List[str]] = ["lago_id", "lago_coupon_id", "coupon_code", "coupon_name", "lago_customer_id", "external_customer_id", "status", "amount_cents", "amount_cents_remaining", "amount_currency", "percentage_rate", "frequency", "frequency_duration", "frequency_duration_remaining", "expiration_at", "created_at", "terminated_at", "credits"]

    @field_validator('status')
    def status_validate_enum(cls, value):
        """Validates the enum"""
        if value not in ('active', 'terminated'):
            raise ValueError("must be one of enum values ('active', 'terminated')")
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
        """Create an instance of AppliedCouponObjectExtended from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of each item in credits (list)
        _items = []
        if self.credits:
            for _item in self.credits:
                if _item:
                    _items.append(_item.to_dict())
            _dict['credits'] = _items
        # set to None if amount_cents (nullable) is None
        # and model_fields_set contains the field
        if self.amount_cents is None and "amount_cents" in self.model_fields_set:
            _dict['amount_cents'] = None

        # set to None if amount_cents_remaining (nullable) is None
        # and model_fields_set contains the field
        if self.amount_cents_remaining is None and "amount_cents_remaining" in self.model_fields_set:
            _dict['amount_cents_remaining'] = None

        # set to None if percentage_rate (nullable) is None
        # and model_fields_set contains the field
        if self.percentage_rate is None and "percentage_rate" in self.model_fields_set:
            _dict['percentage_rate'] = None

        # set to None if frequency_duration (nullable) is None
        # and model_fields_set contains the field
        if self.frequency_duration is None and "frequency_duration" in self.model_fields_set:
            _dict['frequency_duration'] = None

        # set to None if frequency_duration_remaining (nullable) is None
        # and model_fields_set contains the field
        if self.frequency_duration_remaining is None and "frequency_duration_remaining" in self.model_fields_set:
            _dict['frequency_duration_remaining'] = None

        # set to None if expiration_at (nullable) is None
        # and model_fields_set contains the field
        if self.expiration_at is None and "expiration_at" in self.model_fields_set:
            _dict['expiration_at'] = None

        # set to None if terminated_at (nullable) is None
        # and model_fields_set contains the field
        if self.terminated_at is None and "terminated_at" in self.model_fields_set:
            _dict['terminated_at'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Dict) -> Self:
        """Create an instance of AppliedCouponObjectExtended from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "lago_id": obj.get("lago_id"),
            "lago_coupon_id": obj.get("lago_coupon_id"),
            "coupon_code": obj.get("coupon_code"),
            "coupon_name": obj.get("coupon_name"),
            "lago_customer_id": obj.get("lago_customer_id"),
            "external_customer_id": obj.get("external_customer_id"),
            "status": obj.get("status"),
            "amount_cents": obj.get("amount_cents"),
            "amount_cents_remaining": obj.get("amount_cents_remaining"),
            "amount_currency": Currency.from_dict(obj.get("amount_currency")) if obj.get("amount_currency") is not None else None,
            "percentage_rate": obj.get("percentage_rate"),
            "frequency": obj.get("frequency"),
            "frequency_duration": obj.get("frequency_duration"),
            "frequency_duration_remaining": obj.get("frequency_duration_remaining"),
            "expiration_at": obj.get("expiration_at"),
            "created_at": obj.get("created_at"),
            "terminated_at": obj.get("terminated_at"),
            "credits": [CreditObject.from_dict(_item) for _item in obj.get("credits")] if obj.get("credits") is not None else None
        })
        return _obj


