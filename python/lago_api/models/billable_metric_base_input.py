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
from pydantic import BaseModel, StrictBool, StrictStr, field_validator
from pydantic import Field
from lago_api.models.billable_metric_group import BillableMetricGroup
try:
    from typing import Self
except ImportError:
    from typing_extensions import Self

class BillableMetricBaseInput(BaseModel):
    """
    BillableMetricBaseInput
    """ # noqa: E501
    name: Optional[StrictStr] = Field(default=None, description="Name of the billable metric.")
    code: Optional[StrictStr] = Field(default=None, description="Unique code used to identify the billable metric associated with the API request. This code associates each event with the correct metric.")
    description: Optional[StrictStr] = Field(default=None, description="Internal description of the billable metric.")
    recurring: Optional[StrictBool] = Field(default=None, description="Defines if the billable metric is persisted billing period over billing period.  - If set to `true`: the accumulated number of units calculated from the previous billing period is persisted to the next billing period. - If set to `false`: the accumulated number of units is reset to 0 at the end of the billing period. - If not defined in the request, default value is `false`.")
    field_name: Optional[StrictStr] = Field(default=None, description="Property of the billable metric used for aggregating usage data. This field is not required for `count_agg`.")
    aggregation_type: Optional[StrictStr] = Field(default=None, description="Aggregation method used to compute usage for this billable metric.")
    weighted_interval: Optional[StrictStr] = Field(default=None, description="Parameter exclusively utilized in conjunction with the `weighted_sum` aggregation type. It serves to adjust the aggregation result by assigning weights and proration to the result based on time intervals. When this field is not provided, the default time interval is assumed to be in `seconds`.")
    group: Optional[BillableMetricGroup] = None
    __properties: ClassVar[List[str]] = ["name", "code", "description", "recurring", "field_name", "aggregation_type", "weighted_interval", "group"]

    @field_validator('aggregation_type')
    def aggregation_type_validate_enum(cls, value):
        """Validates the enum"""
        if value is None:
            return value

        if value not in ('count_agg', 'sum_agg', 'max_agg', 'unique_count_agg', 'weighted_sum_agg', 'latest_agg'):
            raise ValueError("must be one of enum values ('count_agg', 'sum_agg', 'max_agg', 'unique_count_agg', 'weighted_sum_agg', 'latest_agg')")
        return value

    @field_validator('weighted_interval')
    def weighted_interval_validate_enum(cls, value):
        """Validates the enum"""
        if value is None:
            return value

        if value not in ('seconds'):
            raise ValueError("must be one of enum values ('seconds')")
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
        """Create an instance of BillableMetricBaseInput from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of group
        if self.group:
            _dict['group'] = self.group.to_dict()
        # set to None if description (nullable) is None
        # and model_fields_set contains the field
        if self.description is None and "description" in self.model_fields_set:
            _dict['description'] = None

        # set to None if field_name (nullable) is None
        # and model_fields_set contains the field
        if self.field_name is None and "field_name" in self.model_fields_set:
            _dict['field_name'] = None

        # set to None if weighted_interval (nullable) is None
        # and model_fields_set contains the field
        if self.weighted_interval is None and "weighted_interval" in self.model_fields_set:
            _dict['weighted_interval'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Dict) -> Self:
        """Create an instance of BillableMetricBaseInput from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "name": obj.get("name"),
            "code": obj.get("code"),
            "description": obj.get("description"),
            "recurring": obj.get("recurring"),
            "field_name": obj.get("field_name"),
            "aggregation_type": obj.get("aggregation_type"),
            "weighted_interval": obj.get("weighted_interval"),
            "group": BillableMetricGroup.from_dict(obj.get("group")) if obj.get("group") is not None else None
        })
        return _obj


