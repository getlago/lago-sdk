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
from pydantic import BaseModel, StrictStr, field_validator
from pydantic import Field
try:
    from typing import Self
except ImportError:
    from typing_extensions import Self

class WebhookEndpointObject(BaseModel):
    """
    WebhookEndpointObject
    """ # noqa: E501
    lago_id: StrictStr = Field(description="Unique identifier assigned to the wallet within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the webhook endpoint's record within the Lago system.")
    lago_organization_id: StrictStr = Field(description="Unique identifier assigned to the organization attached to the webhook endpoint within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the organization’s record within the Lago system.")
    webhook_url: StrictStr = Field(description="The name of the wallet.")
    signature_algo: Optional[StrictStr] = Field(default=None, description="The signature algo for the webhook.")
    created_at: datetime = Field(description="The date of the webhook endpoint creation, represented in ISO 8601 datetime format and expressed in Coordinated Universal Time (UTC).")
    __properties: ClassVar[List[str]] = ["lago_id", "lago_organization_id", "webhook_url", "signature_algo", "created_at"]

    @field_validator('signature_algo')
    def signature_algo_validate_enum(cls, value):
        """Validates the enum"""
        if value is None:
            return value

        if value not in ('jwt', 'hmac'):
            raise ValueError("must be one of enum values ('jwt', 'hmac')")
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
        """Create an instance of WebhookEndpointObject from a JSON string"""
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
        """Create an instance of WebhookEndpointObject from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "lago_id": obj.get("lago_id"),
            "lago_organization_id": obj.get("lago_organization_id"),
            "webhook_url": obj.get("webhook_url"),
            "signature_algo": obj.get("signature_algo"),
            "created_at": obj.get("created_at")
        })
        return _obj


