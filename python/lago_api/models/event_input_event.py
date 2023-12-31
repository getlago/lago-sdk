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
from pydantic import BaseModel, StrictStr
from pydantic import Field
from lago_api.models.event_input_event_timestamp import EventInputEventTimestamp
try:
    from typing import Self
except ImportError:
    from typing_extensions import Self

class EventInputEvent(BaseModel):
    """
    EventInputEvent
    """ # noqa: E501
    transaction_id: StrictStr = Field(description="This field represents a unique identifier for the event. It is crucial for ensuring idempotency, meaning that each event can be uniquely identified and processed without causing any unintended side effects.")
    external_customer_id: Optional[StrictStr] = Field(default=None, description="The customer external unique identifier (provided by your own application). This field is optional if you send the `external_subscription_id`, targeting a specific subscription.")
    external_subscription_id: Optional[StrictStr] = Field(default=None, description="The unique identifier of the subscription within your application. It is a mandatory field when the customer possesses multiple subscriptions or when the `external_customer_id` is not provided.")
    code: StrictStr = Field(description="The code that identifies a targeted billable metric. It is essential that this code matches the `code` property of one of your active billable metrics. If the provided code does not correspond to any active billable metric, it will be ignored during the process.")
    timestamp: Optional[EventInputEventTimestamp] = None
    properties: Optional[Dict[str, StrictStr]] = Field(default=None, description="This field represents additional properties associated with the event, which are utilized in the calculation of the final fee. This object becomes mandatory when the targeted billable metric employs a `sum_agg`, `max_agg`, or `unique_count_agg` aggregation method. However, when using a simple `count_agg`, this object is not required.")
    __properties: ClassVar[List[str]] = ["transaction_id", "external_customer_id", "external_subscription_id", "code", "timestamp", "properties"]

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
        """Create an instance of EventInputEvent from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of timestamp
        if self.timestamp:
            _dict['timestamp'] = self.timestamp.to_dict()
        return _dict

    @classmethod
    def from_dict(cls, obj: Dict) -> Self:
        """Create an instance of EventInputEvent from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "transaction_id": obj.get("transaction_id"),
            "external_customer_id": obj.get("external_customer_id"),
            "external_subscription_id": obj.get("external_subscription_id"),
            "code": obj.get("code"),
            "timestamp": EventInputEventTimestamp.from_dict(obj.get("timestamp")) if obj.get("timestamp") is not None else None,
            "properties": obj.get("properties")
        })
        return _obj


