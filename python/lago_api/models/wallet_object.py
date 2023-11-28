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
try:
    from typing import Self
except ImportError:
    from typing_extensions import Self

class WalletObject(BaseModel):
    """
    WalletObject
    """ # noqa: E501
    lago_id: StrictStr = Field(description="Unique identifier assigned to the wallet within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the wallet’s record within the Lago system.")
    lago_customer_id: StrictStr = Field(description="Unique identifier assigned to the customer within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the customer’s record within the Lago system.")
    external_customer_id: StrictStr = Field(description="The customer external unique identifier (provided by your own application)")
    status: StrictStr = Field(description="The status of the wallet. Possible values are `active` or `terminated`.")
    currency: Any
    name: Optional[StrictStr] = Field(default=None, description="The name of the wallet.")
    rate_amount: Annotated[str, Field(strict=True)] = Field(description="The rate of conversion between credits and the amount in the specified currency. It indicates the ratio or factor used to convert credits into the corresponding monetary value in the currency of the transaction.")
    credits_balance: Annotated[str, Field(strict=True)] = Field(description="The current wallet balance expressed in credits.")
    balance_cents: StrictInt = Field(description="The current wallet balance expressed in cents.")
    consumed_credits: Annotated[str, Field(strict=True)] = Field(description="The number of consumed credits.")
    created_at: datetime = Field(description="The date of the wallet creation, represented in ISO 8601 datetime format and expressed in Coordinated Universal Time (UTC).")
    expiration_at: Optional[datetime] = Field(default=None, description="The date and time that determines when the wallet will expire. It follows the ISO 8601 datetime format and is expressed in Coordinated Universal Time (UTC).")
    last_balance_sync_at: Optional[datetime] = Field(default=None, description="The date and time of the last balance top-up. It follows the ISO 8601 datetime format and is expressed in Coordinated Universal Time (UTC).")
    last_consumed_credit_at: Optional[datetime] = Field(default=None, description="The date and time of the last credits consumption. It follows the ISO 8601 datetime format and is expressed in Coordinated Universal Time (UTC).")
    terminated_at: Optional[datetime] = Field(default=None, description="The date of terminaison of the wallet. It follows the ISO 8601 datetime format and is expressed in Coordinated Universal Time (UTC).")
    __properties: ClassVar[List[str]] = ["lago_id", "lago_customer_id", "external_customer_id", "status", "currency", "name", "rate_amount", "credits_balance", "balance_cents", "consumed_credits", "created_at", "expiration_at", "last_balance_sync_at", "last_consumed_credit_at", "terminated_at"]

    @field_validator('status')
    def status_validate_enum(cls, value):
        """Validates the enum"""
        if value not in ('active', 'terminated'):
            raise ValueError("must be one of enum values ('active', 'terminated')")
        return value

    @field_validator('rate_amount')
    def rate_amount_validate_regular_expression(cls, value):
        """Validates the regular expression"""
        if not re.match(r"^[0-9]+.?[0-9]*$", value):
            raise ValueError(r"must validate the regular expression /^[0-9]+.?[0-9]*$/")
        return value

    @field_validator('credits_balance')
    def credits_balance_validate_regular_expression(cls, value):
        """Validates the regular expression"""
        if not re.match(r"^[0-9]+.?[0-9]*$", value):
            raise ValueError(r"must validate the regular expression /^[0-9]+.?[0-9]*$/")
        return value

    @field_validator('consumed_credits')
    def consumed_credits_validate_regular_expression(cls, value):
        """Validates the regular expression"""
        if not re.match(r"^[0-9]+.?[0-9]*$", value):
            raise ValueError(r"must validate the regular expression /^[0-9]+.?[0-9]*$/")
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
        """Create an instance of WalletObject from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of currency
        if self.currency:
            _dict['currency'] = self.currency.to_dict()
        # set to None if expiration_at (nullable) is None
        # and model_fields_set contains the field
        if self.expiration_at is None and "expiration_at" in self.model_fields_set:
            _dict['expiration_at'] = None

        # set to None if last_balance_sync_at (nullable) is None
        # and model_fields_set contains the field
        if self.last_balance_sync_at is None and "last_balance_sync_at" in self.model_fields_set:
            _dict['last_balance_sync_at'] = None

        # set to None if last_consumed_credit_at (nullable) is None
        # and model_fields_set contains the field
        if self.last_consumed_credit_at is None and "last_consumed_credit_at" in self.model_fields_set:
            _dict['last_consumed_credit_at'] = None

        # set to None if terminated_at (nullable) is None
        # and model_fields_set contains the field
        if self.terminated_at is None and "terminated_at" in self.model_fields_set:
            _dict['terminated_at'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Dict) -> Self:
        """Create an instance of WalletObject from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "lago_id": obj.get("lago_id"),
            "lago_customer_id": obj.get("lago_customer_id"),
            "external_customer_id": obj.get("external_customer_id"),
            "status": obj.get("status"),
            "currency": Currency.from_dict(obj.get("currency")) if obj.get("currency") is not None else None,
            "name": obj.get("name"),
            "rate_amount": obj.get("rate_amount"),
            "credits_balance": obj.get("credits_balance"),
            "balance_cents": obj.get("balance_cents"),
            "consumed_credits": obj.get("consumed_credits"),
            "created_at": obj.get("created_at"),
            "expiration_at": obj.get("expiration_at"),
            "last_balance_sync_at": obj.get("last_balance_sync_at"),
            "last_consumed_credit_at": obj.get("last_consumed_credit_at"),
            "terminated_at": obj.get("terminated_at")
        })
        return _obj

