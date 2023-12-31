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
from pydantic import BaseModel, StrictInt, StrictStr, field_validator
from pydantic import Field
from typing_extensions import Annotated
from lago_api.models.charge_properties_graduated_percentage_ranges_inner import ChargePropertiesGraduatedPercentageRangesInner
from lago_api.models.charge_properties_graduated_ranges_inner import ChargePropertiesGraduatedRangesInner
from lago_api.models.charge_properties_volume_ranges_inner import ChargePropertiesVolumeRangesInner
try:
    from typing import Self
except ImportError:
    from typing_extensions import Self

class ChargeObjectProperties(BaseModel):
    """
    ChargeObjectProperties
    """ # noqa: E501
    graduated_ranges: Optional[List[ChargePropertiesGraduatedRangesInner]] = Field(default=None, description="Graduated ranges, sorted from bottom to top tiers, used for a `graduated` charge model.")
    graduated_percentage_ranges: Optional[List[ChargePropertiesGraduatedPercentageRangesInner]] = Field(default=None, description="Graduated percentage ranges, sorted from bottom to top tiers, used for a `graduated_percentage` charge model.")
    amount: Optional[Annotated[str, Field(strict=True)]] = Field(default=None, description="- The unit price, excluding tax, for a `standard` charge model. It is expressed as a decimal value. - The amount, excluding tax, for a complete set of units in a `package` charge model. It is expressed as a decimal value.")
    free_units: Optional[StrictInt] = Field(default=None, description="The quantity of units that are provided free of charge for each billing period in a `package` charge model. This field specifies the number of units that customers can use without incurring any additional cost during each billing cycle.")
    package_size: Optional[StrictInt] = Field(default=None, description="The quantity of units included in each pack or set for a `package` charge model. It indicates the number of units that are bundled together as a single package or set within the pricing structure.")
    rate: Optional[Annotated[str, Field(strict=True)]] = Field(default=None, description="The percentage rate that is applied to the amount of each transaction for a `percentage` charge model. It is expressed as a decimal value.")
    fixed_amount: Optional[Annotated[str, Field(strict=True)]] = Field(default=None, description="The fixed fee that is applied to each transaction for a `percentage` charge model. It is expressed as a decimal value.")
    free_units_per_events: Optional[StrictInt] = Field(default=None, description="The count of transactions that are not impacted by the `percentage` rate and fixed fee in a percentage charge model. This field indicates the number of transactions that are exempt from the calculation of charges based on the specified percentage rate and fixed fee.")
    free_units_per_total_aggregation: Optional[Annotated[str, Field(strict=True)]] = Field(default=None, description="The transaction amount that is not impacted by the `percentage` rate and fixed fee in a percentage charge model. This field indicates the portion of the transaction amount that is exempt from the calculation of charges based on the specified percentage rate and fixed fee.")
    per_transaction_max_amount: Optional[StrictStr] = Field(default=None, description="Specifies the maximum allowable spending for a single transaction. Working as a transaction cap.")
    per_transaction_min_amount: Optional[StrictStr] = Field(default=None, description="Specifies the minimum allowable spending for a single transaction. Working as a transaction floor.")
    volume_ranges: Optional[List[ChargePropertiesVolumeRangesInner]] = Field(default=None, description="Volume ranges, sorted from bottom to top tiers, used for a `volume` charge model.")
    __properties: ClassVar[List[str]] = ["graduated_ranges", "graduated_percentage_ranges", "amount", "free_units", "package_size", "rate", "fixed_amount", "free_units_per_events", "free_units_per_total_aggregation", "per_transaction_max_amount", "per_transaction_min_amount", "volume_ranges"]

    @field_validator('amount')
    def amount_validate_regular_expression(cls, value):
        """Validates the regular expression"""
        if value is None:
            return value

        if not re.match(r"^[0-9]+.?[0-9]*$", value):
            raise ValueError(r"must validate the regular expression /^[0-9]+.?[0-9]*$/")
        return value

    @field_validator('rate')
    def rate_validate_regular_expression(cls, value):
        """Validates the regular expression"""
        if value is None:
            return value

        if not re.match(r"^[0-9]+.?[0-9]*$", value):
            raise ValueError(r"must validate the regular expression /^[0-9]+.?[0-9]*$/")
        return value

    @field_validator('fixed_amount')
    def fixed_amount_validate_regular_expression(cls, value):
        """Validates the regular expression"""
        if value is None:
            return value

        if not re.match(r"^[0-9]+.?[0-9]*$", value):
            raise ValueError(r"must validate the regular expression /^[0-9]+.?[0-9]*$/")
        return value

    @field_validator('free_units_per_total_aggregation')
    def free_units_per_total_aggregation_validate_regular_expression(cls, value):
        """Validates the regular expression"""
        if value is None:
            return value

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
        """Create an instance of ChargeObjectProperties from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of each item in graduated_ranges (list)
        _items = []
        if self.graduated_ranges:
            for _item in self.graduated_ranges:
                if _item:
                    _items.append(_item.to_dict())
            _dict['graduated_ranges'] = _items
        # override the default output from pydantic by calling `to_dict()` of each item in graduated_percentage_ranges (list)
        _items = []
        if self.graduated_percentage_ranges:
            for _item in self.graduated_percentage_ranges:
                if _item:
                    _items.append(_item.to_dict())
            _dict['graduated_percentage_ranges'] = _items
        # override the default output from pydantic by calling `to_dict()` of each item in volume_ranges (list)
        _items = []
        if self.volume_ranges:
            for _item in self.volume_ranges:
                if _item:
                    _items.append(_item.to_dict())
            _dict['volume_ranges'] = _items
        # set to None if free_units_per_events (nullable) is None
        # and model_fields_set contains the field
        if self.free_units_per_events is None and "free_units_per_events" in self.model_fields_set:
            _dict['free_units_per_events'] = None

        # set to None if free_units_per_total_aggregation (nullable) is None
        # and model_fields_set contains the field
        if self.free_units_per_total_aggregation is None and "free_units_per_total_aggregation" in self.model_fields_set:
            _dict['free_units_per_total_aggregation'] = None

        # set to None if per_transaction_max_amount (nullable) is None
        # and model_fields_set contains the field
        if self.per_transaction_max_amount is None and "per_transaction_max_amount" in self.model_fields_set:
            _dict['per_transaction_max_amount'] = None

        # set to None if per_transaction_min_amount (nullable) is None
        # and model_fields_set contains the field
        if self.per_transaction_min_amount is None and "per_transaction_min_amount" in self.model_fields_set:
            _dict['per_transaction_min_amount'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Dict) -> Self:
        """Create an instance of ChargeObjectProperties from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "graduated_ranges": [ChargePropertiesGraduatedRangesInner.from_dict(_item) for _item in obj.get("graduated_ranges")] if obj.get("graduated_ranges") is not None else None,
            "graduated_percentage_ranges": [ChargePropertiesGraduatedPercentageRangesInner.from_dict(_item) for _item in obj.get("graduated_percentage_ranges")] if obj.get("graduated_percentage_ranges") is not None else None,
            "amount": obj.get("amount"),
            "free_units": obj.get("free_units"),
            "package_size": obj.get("package_size"),
            "rate": obj.get("rate"),
            "fixed_amount": obj.get("fixed_amount"),
            "free_units_per_events": obj.get("free_units_per_events"),
            "free_units_per_total_aggregation": obj.get("free_units_per_total_aggregation"),
            "per_transaction_max_amount": obj.get("per_transaction_max_amount"),
            "per_transaction_min_amount": obj.get("per_transaction_min_amount"),
            "volume_ranges": [ChargePropertiesVolumeRangesInner.from_dict(_item) for _item in obj.get("volume_ranges")] if obj.get("volume_ranges") is not None else None
        })
        return _obj


