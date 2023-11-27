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
from typing import Any, ClassVar, Dict, List, Optional, Union
from pydantic import BaseModel, StrictFloat, StrictInt, StrictStr
from pydantic import Field
try:
    from typing import Self
except ImportError:
    from typing_extensions import Self

class BaseAppliedTax(BaseModel):
    """
    BaseAppliedTax
    """ # noqa: E501
    lago_id: Optional[StrictStr] = Field(default=None, description="Unique identifier of the applied tax, created by Lago.")
    lago_tax_id: Optional[StrictStr] = Field(default=None, description="Unique identifier of the tax, created by Lago.")
    tax_name: Optional[StrictStr] = Field(default=None, description="Name of the tax.")
    tax_code: Optional[StrictStr] = Field(default=None, description="Unique code used to identify the tax associated with the API request.")
    tax_rate: Optional[Union[StrictFloat, StrictInt]] = Field(default=None, description="The percentage rate of the tax")
    tax_description: Optional[StrictStr] = Field(default=None, description="Internal description of the taxe")
    amount_cents: Optional[StrictInt] = Field(default=None, description="Amount of the tax")
    amount_currency: Optional[Any] = None
    created_at: Optional[datetime] = Field(default=None, description="The date and time when the applied tax was created. It is expressed in UTC format according to the ISO 8601 datetime standard. This field provides the timestamp for the exact moment when the applied tax was initially created.")
    __properties: ClassVar[List[str]] = ["lago_id", "lago_tax_id", "tax_name", "tax_code", "tax_rate", "tax_description", "amount_cents", "amount_currency", "created_at"]

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
        """Create an instance of BaseAppliedTax from a JSON string"""
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
        return _dict

    @classmethod
    def from_dict(cls, obj: Dict) -> Self:
        """Create an instance of BaseAppliedTax from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "lago_id": obj.get("lago_id"),
            "lago_tax_id": obj.get("lago_tax_id"),
            "tax_name": obj.get("tax_name"),
            "tax_code": obj.get("tax_code"),
            "tax_rate": obj.get("tax_rate"),
            "tax_description": obj.get("tax_description"),
            "amount_cents": obj.get("amount_cents"),
            "amount_currency": Currency.from_dict(obj.get("amount_currency")) if obj.get("amount_currency") is not None else None,
            "created_at": obj.get("created_at")
        })
        return _obj


