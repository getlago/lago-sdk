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
from pydantic import BaseModel, StrictBool, StrictFloat, StrictInt, StrictStr, field_validator
from pydantic import Field
from lago_api.models.currency import Currency
from lago_api.models.fee_applied_tax_object import FeeAppliedTaxObject
from lago_api.models.fee_object_item import FeeObjectItem
try:
    from typing import Self
except ImportError:
    from typing_extensions import Self

class CreditNoteItemObjectFee(BaseModel):
    """
    CreditNoteItemObjectFee
    """ # noqa: E501
    lago_id: Optional[StrictStr] = Field(default=None, description="Unique identifier assigned to the fee within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the fee’s record within the Lago system.")
    lago_group_id: Optional[StrictStr] = Field(default=None, description="Unique identifier assigned to the group that the fee belongs to")
    lago_invoice_id: Optional[StrictStr] = Field(default=None, description="Unique identifier assigned to the invoice that the fee belongs to")
    lago_true_up_fee_id: Optional[StrictStr] = Field(default=None, description="Unique identifier assigned to the true-up fee when a minimum has been set to the charge. This identifier helps to distinguish and manage the true-up fee associated with the charge, which may be applicable when a minimum threshold or limit is set for the charge amount.")
    lago_true_up_parent_fee_id: Optional[StrictStr] = Field(default=None, description="Unique identifier assigned to the parent fee on which the true-up fee is assigned. This identifier establishes the relationship between the parent fee and the associated true-up fee.")
    lago_subscription_id: Optional[StrictStr] = Field(default=None, description="Unique identifier assigned to the subscription, created by Lago. This field is specifically displayed when the fee type is charge or subscription.")
    lago_customer_id: Optional[StrictStr] = Field(default=None, description="Unique identifier assigned to the customer, created by Lago. This field is specifically displayed when the fee type is charge or subscription.")
    external_customer_id: Optional[StrictStr] = Field(default=None, description="Unique identifier assigned to the customer in your application. This field is specifically displayed when the fee type is charge or subscription.")
    external_subscription_id: Optional[StrictStr] = Field(default=None, description="Unique identifier assigned to the subscription in your application. This field is specifically displayed when the fee type is charge or subscription.")
    invoice_display_name: Optional[StrictStr] = Field(default=None, description="Specifies the name that will be displayed on an invoice. If no value is set for this field, the name of the actual charge will be used as the default display name.")
    amount_cents: StrictInt = Field(description="The cost of this specific fee, excluding any applicable taxes.")
    amount_currency: Any
    taxes_amount_cents: StrictInt = Field(description="The cost of the tax associated with this specific fee.")
    taxes_rate: Union[StrictFloat, StrictInt] = Field(description="The tax rate associated with this specific fee.")
    units: StrictStr = Field(description="The number of units used to charge the customer. This field indicates the quantity or count of units consumed or utilized in the context of the charge. It helps in determining the basis for calculating the fee or cost associated with the usage of the service or product provided to the customer.")
    precise_unit_amount: StrictStr = Field(description="The unit amount of the fee per unit, with precision.")
    total_amount_cents: StrictInt = Field(description="The cost of this specific fee, including any applicable taxes.")
    total_amount_currency: Any
    events_count: Optional[StrictInt] = Field(default=None, description="The number of events that have been sent and used to charge the customer. This field indicates the count or quantity of events that have been processed and considered in the charging process.")
    pay_in_advance: StrictBool = Field(description="Flag that indicates whether the fee was paid in advance. It serves as a boolean value, where `true` represents that the fee was paid in advance (straightaway), and `false` indicates that the fee was not paid in arrears (at the end of the period).")
    invoiceable: StrictBool = Field(description="Flag that indicates whether the fee was included on the invoice. It serves as a boolean value, where `true` represents that the fee was included on the invoice, and `false` indicates that the fee was not included on the invoice.")
    from_date: Optional[datetime] = Field(default=None, description="The beginning date of the period that the fee covers. It is applicable only to `subscription` and `charge` fees. This field indicates the start date of the billing period or subscription period associated with the fee.")
    to_date: Optional[datetime] = Field(default=None, description="The ending date of the period that the fee covers. It is applicable only to `subscription` and `charge` fees. This field indicates the end date of the billing period or subscription period associated with the fee.")
    payment_status: StrictStr = Field(description="Indicates the payment status of the fee. It represents the current status of the payment associated with the fee. The possible values for this field are `pending`, `succeeded`, `failed` and `refunded`.")
    created_at: Optional[datetime] = Field(default=None, description="The date and time when the fee was created. It is provided in Coordinated Universal Time (UTC) format.")
    succeeded_at: Optional[datetime] = Field(default=None, description="The date and time when the payment for the fee was successfully processed. It is provided in Coordinated Universal Time (UTC) format.")
    failed_at: Optional[datetime] = Field(default=None, description="The date and time when the payment for the fee failed to process. It is provided in Coordinated Universal Time (UTC) format.")
    refunded_at: Optional[datetime] = Field(default=None, description="The date and time when the payment for the fee was refunded. It is provided in Coordinated Universal Time (UTC) format")
    event_transaction_id: Optional[StrictStr] = Field(default=None, description="Unique identifier assigned to the transaction. This field is specifically displayed when the fee type is `charge` and the payment for the fee is made in advance (`pay_in_advance` is set to `true`).")
    item: FeeObjectItem
    applied_taxes: Optional[List[FeeAppliedTaxObject]] = Field(default=None, description="List of fee applied taxes")
    __properties: ClassVar[List[str]] = ["lago_id", "lago_group_id", "lago_invoice_id", "lago_true_up_fee_id", "lago_true_up_parent_fee_id", "lago_subscription_id", "lago_customer_id", "external_customer_id", "external_subscription_id", "invoice_display_name", "amount_cents", "amount_currency", "taxes_amount_cents", "taxes_rate", "units", "precise_unit_amount", "total_amount_cents", "total_amount_currency", "events_count", "pay_in_advance", "invoiceable", "from_date", "to_date", "payment_status", "created_at", "succeeded_at", "failed_at", "refunded_at", "event_transaction_id", "item", "applied_taxes"]

    @field_validator('payment_status')
    def payment_status_validate_enum(cls, value):
        """Validates the enum"""
        if value not in ('pending', 'succeeded', 'failed', 'refunded'):
            raise ValueError("must be one of enum values ('pending', 'succeeded', 'failed', 'refunded')")
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
        """Create an instance of CreditNoteItemObjectFee from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of total_amount_currency
        if self.total_amount_currency:
            _dict['total_amount_currency'] = self.total_amount_currency.to_dict()
        # override the default output from pydantic by calling `to_dict()` of item
        if self.item:
            _dict['item'] = self.item.to_dict()
        # override the default output from pydantic by calling `to_dict()` of each item in applied_taxes (list)
        _items = []
        if self.applied_taxes:
            for _item in self.applied_taxes:
                if _item:
                    _items.append(_item.to_dict())
            _dict['applied_taxes'] = _items
        # set to None if lago_id (nullable) is None
        # and model_fields_set contains the field
        if self.lago_id is None and "lago_id" in self.model_fields_set:
            _dict['lago_id'] = None

        # set to None if lago_group_id (nullable) is None
        # and model_fields_set contains the field
        if self.lago_group_id is None and "lago_group_id" in self.model_fields_set:
            _dict['lago_group_id'] = None

        # set to None if lago_invoice_id (nullable) is None
        # and model_fields_set contains the field
        if self.lago_invoice_id is None and "lago_invoice_id" in self.model_fields_set:
            _dict['lago_invoice_id'] = None

        # set to None if lago_true_up_fee_id (nullable) is None
        # and model_fields_set contains the field
        if self.lago_true_up_fee_id is None and "lago_true_up_fee_id" in self.model_fields_set:
            _dict['lago_true_up_fee_id'] = None

        # set to None if lago_true_up_parent_fee_id (nullable) is None
        # and model_fields_set contains the field
        if self.lago_true_up_parent_fee_id is None and "lago_true_up_parent_fee_id" in self.model_fields_set:
            _dict['lago_true_up_parent_fee_id'] = None

        # set to None if lago_subscription_id (nullable) is None
        # and model_fields_set contains the field
        if self.lago_subscription_id is None and "lago_subscription_id" in self.model_fields_set:
            _dict['lago_subscription_id'] = None

        # set to None if lago_customer_id (nullable) is None
        # and model_fields_set contains the field
        if self.lago_customer_id is None and "lago_customer_id" in self.model_fields_set:
            _dict['lago_customer_id'] = None

        # set to None if external_customer_id (nullable) is None
        # and model_fields_set contains the field
        if self.external_customer_id is None and "external_customer_id" in self.model_fields_set:
            _dict['external_customer_id'] = None

        # set to None if external_subscription_id (nullable) is None
        # and model_fields_set contains the field
        if self.external_subscription_id is None and "external_subscription_id" in self.model_fields_set:
            _dict['external_subscription_id'] = None

        # set to None if from_date (nullable) is None
        # and model_fields_set contains the field
        if self.from_date is None and "from_date" in self.model_fields_set:
            _dict['from_date'] = None

        # set to None if to_date (nullable) is None
        # and model_fields_set contains the field
        if self.to_date is None and "to_date" in self.model_fields_set:
            _dict['to_date'] = None

        # set to None if created_at (nullable) is None
        # and model_fields_set contains the field
        if self.created_at is None and "created_at" in self.model_fields_set:
            _dict['created_at'] = None

        # set to None if succeeded_at (nullable) is None
        # and model_fields_set contains the field
        if self.succeeded_at is None and "succeeded_at" in self.model_fields_set:
            _dict['succeeded_at'] = None

        # set to None if failed_at (nullable) is None
        # and model_fields_set contains the field
        if self.failed_at is None and "failed_at" in self.model_fields_set:
            _dict['failed_at'] = None

        # set to None if refunded_at (nullable) is None
        # and model_fields_set contains the field
        if self.refunded_at is None and "refunded_at" in self.model_fields_set:
            _dict['refunded_at'] = None

        # set to None if event_transaction_id (nullable) is None
        # and model_fields_set contains the field
        if self.event_transaction_id is None and "event_transaction_id" in self.model_fields_set:
            _dict['event_transaction_id'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Dict) -> Self:
        """Create an instance of CreditNoteItemObjectFee from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "lago_id": obj.get("lago_id"),
            "lago_group_id": obj.get("lago_group_id"),
            "lago_invoice_id": obj.get("lago_invoice_id"),
            "lago_true_up_fee_id": obj.get("lago_true_up_fee_id"),
            "lago_true_up_parent_fee_id": obj.get("lago_true_up_parent_fee_id"),
            "lago_subscription_id": obj.get("lago_subscription_id"),
            "lago_customer_id": obj.get("lago_customer_id"),
            "external_customer_id": obj.get("external_customer_id"),
            "external_subscription_id": obj.get("external_subscription_id"),
            "invoice_display_name": obj.get("invoice_display_name"),
            "amount_cents": obj.get("amount_cents"),
            "amount_currency": Currency.from_dict(obj.get("amount_currency")) if obj.get("amount_currency") is not None else None,
            "taxes_amount_cents": obj.get("taxes_amount_cents"),
            "taxes_rate": obj.get("taxes_rate"),
            "units": obj.get("units"),
            "precise_unit_amount": obj.get("precise_unit_amount"),
            "total_amount_cents": obj.get("total_amount_cents"),
            "total_amount_currency": Currency.from_dict(obj.get("total_amount_currency")) if obj.get("total_amount_currency") is not None else None,
            "events_count": obj.get("events_count"),
            "pay_in_advance": obj.get("pay_in_advance"),
            "invoiceable": obj.get("invoiceable"),
            "from_date": obj.get("from_date"),
            "to_date": obj.get("to_date"),
            "payment_status": obj.get("payment_status"),
            "created_at": obj.get("created_at"),
            "succeeded_at": obj.get("succeeded_at"),
            "failed_at": obj.get("failed_at"),
            "refunded_at": obj.get("refunded_at"),
            "event_transaction_id": obj.get("event_transaction_id"),
            "item": FeeObjectItem.from_dict(obj.get("item")) if obj.get("item") is not None else None,
            "applied_taxes": [FeeAppliedTaxObject.from_dict(_item) for _item in obj.get("applied_taxes")] if obj.get("applied_taxes") is not None else None
        })
        return _obj

