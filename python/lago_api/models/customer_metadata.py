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
from typing import Any, ClassVar, Dict, List
from pydantic import BaseModel, StrictBool, StrictStr
from pydantic import Field
try:
    from typing import Self
except ImportError:
    from typing_extensions import Self

class CustomerMetadata(BaseModel):
    """
    Set of key-value pairs that you can attach to a customer. This can be useful for storing additional information about the customer in a structured format
    """ # noqa: E501
    lago_id: StrictStr = Field(description="A unique identifier for the customer metadata object in the Lago application. Can be used to update a key-value pair")
    key: StrictStr = Field(description="The metadata object key")
    value: StrictStr = Field(description="The metadata object value")
    display_in_invoice: StrictBool = Field(description="Determines whether the item or information should be displayed in the invoice. If set to true, the item or information will be included and visible in the generated invoice. If set to false, the item or information will be excluded and not displayed in the invoice.")
    created_at: datetime = Field(description="The date of the metadata object creation, represented in ISO 8601 datetime format and expressed in Coordinated Universal Time (UTC). The creation_date provides a standardized and internationally recognized timestamp for when the metadata object was created")
    __properties: ClassVar[List[str]] = ["lago_id", "key", "value", "display_in_invoice", "created_at"]

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
        """Create an instance of CustomerMetadata from a JSON string"""
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
        return _dict

    @classmethod
    def from_dict(cls, obj: Dict) -> Self:
        """Create an instance of CustomerMetadata from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "lago_id": obj.get("lago_id"),
            "key": obj.get("key"),
            "value": obj.get("value"),
            "display_in_invoice": obj.get("display_in_invoice"),
            "created_at": obj.get("created_at")
        })
        return _obj


