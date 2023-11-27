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


from typing import Any, ClassVar, Dict, List
from pydantic import BaseModel
from lago_api.models.pagination_meta import PaginationMeta
from lago_api.models.webhook_endpoint_object import WebhookEndpointObject
try:
    from typing import Self
except ImportError:
    from typing_extensions import Self

class WebhookEndpointsPaginated(BaseModel):
    """
    WebhookEndpointsPaginated
    """ # noqa: E501
    webhook_endpoints: List[WebhookEndpointObject]
    meta: PaginationMeta
    __properties: ClassVar[List[str]] = ["webhook_endpoints", "meta"]

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
        """Create an instance of WebhookEndpointsPaginated from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of each item in webhook_endpoints (list)
        _items = []
        if self.webhook_endpoints:
            for _item in self.webhook_endpoints:
                if _item:
                    _items.append(_item.to_dict())
            _dict['webhook_endpoints'] = _items
        # override the default output from pydantic by calling `to_dict()` of meta
        if self.meta:
            _dict['meta'] = self.meta.to_dict()
        return _dict

    @classmethod
    def from_dict(cls, obj: Dict) -> Self:
        """Create an instance of WebhookEndpointsPaginated from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "webhook_endpoints": [WebhookEndpointObject.from_dict(_item) for _item in obj.get("webhook_endpoints")] if obj.get("webhook_endpoints") is not None else None,
            "meta": PaginationMeta.from_dict(obj.get("meta")) if obj.get("meta") is not None else None
        })
        return _obj


