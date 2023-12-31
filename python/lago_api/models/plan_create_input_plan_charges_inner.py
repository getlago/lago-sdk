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
from pydantic import BaseModel, StrictBool, StrictInt, StrictStr, field_validator
from pydantic import Field
from lago_api.models.charge_object_properties import ChargeObjectProperties
from lago_api.models.plan_create_input_plan_charges_inner_group_properties_inner import PlanCreateInputPlanChargesInnerGroupPropertiesInner
try:
    from typing import Self
except ImportError:
    from typing_extensions import Self

class PlanCreateInputPlanChargesInner(BaseModel):
    """
    PlanCreateInputPlanChargesInner
    """ # noqa: E501
    billable_metric_id: Optional[StrictStr] = Field(default=None, description="Unique identifier of the billable metric created by Lago.")
    charge_model: Optional[StrictStr] = Field(default=None, description="Specifies the pricing model used for the calculation of the final fee. It can be `standard`, `graduated`, `graduated_percentage` `package`, `percentage` or `volume`.")
    pay_in_advance: Optional[StrictBool] = Field(default=None, description="This field determines the billing timing for this specific usage-based charge. When set to `true`, the charge is due and invoiced immediately. Conversely, when set to false, the charge is due and invoiced at the end of each billing period.")
    invoiceable: Optional[StrictBool] = Field(default=None, description="This field specifies whether the charge should be included in a proper invoice. If set to false, no invoice will be issued for this charge. You can only set it to `false` when `pay_in_advance` is `true`.")
    invoice_display_name: Optional[StrictStr] = Field(default=None, description="Specifies the name that will be displayed on an invoice. If no value is set for this field, the name of the actual charge will be used as the default display name.")
    prorated: Optional[StrictBool] = Field(default=None, description="Specifies whether a charge is prorated based on the remaining number of days in the billing period or billed fully.  - If set to `true`, the charge is prorated based on the remaining days in the current billing period. - If set to `false`, the charge is billed in full. - If not defined in the request, default value is `false`.")
    min_amount_cents: Optional[StrictInt] = Field(default=None, description="The minimum spending amount required for the charge, measured in cents and excluding any applicable taxes. It indicates the minimum amount that needs to be charged for each billing period.")
    properties: Optional[ChargeObjectProperties] = None
    group_properties: Optional[List[PlanCreateInputPlanChargesInnerGroupPropertiesInner]] = Field(default=None, description="All charge information, sorted by groups.")
    tax_codes: Optional[List[StrictStr]] = Field(default=None, description="List of unique code used to identify the taxes.")
    __properties: ClassVar[List[str]] = ["billable_metric_id", "charge_model", "pay_in_advance", "invoiceable", "invoice_display_name", "prorated", "min_amount_cents", "properties", "group_properties", "tax_codes"]

    @field_validator('charge_model')
    def charge_model_validate_enum(cls, value):
        """Validates the enum"""
        if value is None:
            return value

        if value not in ('standard', 'graduated', 'graduated_percentage', 'package', 'percentage', 'volume'):
            raise ValueError("must be one of enum values ('standard', 'graduated', 'graduated_percentage', 'package', 'percentage', 'volume')")
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
        """Create an instance of PlanCreateInputPlanChargesInner from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of properties
        if self.properties:
            _dict['properties'] = self.properties.to_dict()
        # override the default output from pydantic by calling `to_dict()` of each item in group_properties (list)
        _items = []
        if self.group_properties:
            for _item in self.group_properties:
                if _item:
                    _items.append(_item.to_dict())
            _dict['group_properties'] = _items
        return _dict

    @classmethod
    def from_dict(cls, obj: Dict) -> Self:
        """Create an instance of PlanCreateInputPlanChargesInner from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "billable_metric_id": obj.get("billable_metric_id"),
            "charge_model": obj.get("charge_model"),
            "pay_in_advance": obj.get("pay_in_advance"),
            "invoiceable": obj.get("invoiceable"),
            "invoice_display_name": obj.get("invoice_display_name"),
            "prorated": obj.get("prorated"),
            "min_amount_cents": obj.get("min_amount_cents"),
            "properties": ChargeObjectProperties.from_dict(obj.get("properties")) if obj.get("properties") is not None else None,
            "group_properties": [PlanCreateInputPlanChargesInnerGroupPropertiesInner.from_dict(_item) for _item in obj.get("group_properties")] if obj.get("group_properties") is not None else None,
            "tax_codes": obj.get("tax_codes")
        })
        return _obj


