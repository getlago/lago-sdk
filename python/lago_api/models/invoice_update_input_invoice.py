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


from typing import Any, ClassVar, Dict, List, Optional
from pydantic import BaseModel, StrictStr, field_validator
from pydantic import Field
from lago_api.models.invoice_update_input_invoice_metadata_inner import InvoiceUpdateInputInvoiceMetadataInner
try:
    from typing import Self
except ImportError:
    from typing_extensions import Self

class InvoiceUpdateInputInvoice(BaseModel):
    """
    InvoiceUpdateInputInvoice
    """ # noqa: E501
    payment_status: Optional[StrictStr] = Field(default=None, description="The payment status of the invoice. Possible values are `pending`, `failed` or `succeeded`.")
    metadata: Optional[List[InvoiceUpdateInputInvoiceMetadataInner]] = None
    __properties: ClassVar[List[str]] = ["payment_status", "metadata"]

    @field_validator('payment_status')
    def payment_status_validate_enum(cls, value):
        """Validates the enum"""
        if value is None:
            return value

        if value not in ('pending', 'succeeded', 'failed'):
            raise ValueError("must be one of enum values ('pending', 'succeeded', 'failed')")
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
        """Create an instance of InvoiceUpdateInputInvoice from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of each item in metadata (list)
        _items = []
        if self.metadata:
            for _item in self.metadata:
                if _item:
                    _items.append(_item.to_dict())
            _dict['metadata'] = _items
        return _dict

    @classmethod
    def from_dict(cls, obj: Dict) -> Self:
        """Create an instance of InvoiceUpdateInputInvoice from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "payment_status": obj.get("payment_status"),
            "metadata": [InvoiceUpdateInputInvoiceMetadataInner.from_dict(_item) for _item in obj.get("metadata")] if obj.get("metadata") is not None else None
        })
        return _obj


