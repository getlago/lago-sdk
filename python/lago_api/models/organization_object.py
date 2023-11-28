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
from pydantic import BaseModel, StrictInt, StrictStr
from pydantic import Field
from lago_api.models.organization_billing_configuration import OrganizationBillingConfiguration
from lago_api.models.tax_object import TaxObject
try:
    from typing import Self
except ImportError:
    from typing_extensions import Self

class OrganizationObject(BaseModel):
    """
    OrganizationObject
    """ # noqa: E501
    lago_id: StrictStr = Field(description="Unique identifier assigned to the organization within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the organization's record within the Lago system")
    name: StrictStr = Field(description="The name of your organization.")
    created_at: Optional[datetime] = Field(description="The date of creation of your organization, represented in ISO 8601 datetime format and expressed in Coordinated Universal Time (UTC).")
    webhook_url: Optional[StrictStr] = Field(default=None, description="The URL of your newest updated webhook endpoint. This URL allows your organization to receive important messages, notifications, or data from the Lago system. By configuring your webhook endpoint to this URL, you can ensure that your organization stays informed and receives relevant information in a timely manner.")
    webhook_urls: Optional[List[StrictStr]] = Field(default=None, description="The array containing your webhooks URLs.")
    country: Optional[Any] = None
    default_currency: Optional[Any] = None
    address_line1: Optional[StrictStr] = Field(default=None, description="The first line of your organization’s billing address.")
    address_line2: Optional[StrictStr] = Field(default=None, description="The second line of your organization’s billing address.")
    state: Optional[StrictStr] = Field(default=None, description="The state of your organization’s billing address.")
    zipcode: Optional[StrictStr] = Field(default=None, description="The zipcode of your organization’s billing address.")
    email: Optional[StrictStr] = Field(default=None, description="The email address of your organization used to bill your customers.")
    city: Optional[StrictStr] = Field(default=None, description="The city of your organization’s billing address.")
    legal_name: Optional[StrictStr] = Field(default=None, description="The legal name of your organization.")
    legal_number: Optional[StrictStr] = Field(default=None, description="The legal number of your organization.")
    net_payment_term: Optional[StrictInt] = Field(default=None, description="The net payment term, expressed in days, specifies the duration within which a customer is expected to remit payment after the invoice is finalized.")
    tax_identification_number: Optional[StrictStr] = Field(default=None, description="The tax identification number of your organization.")
    timezone: Optional[Any] = None
    billing_configuration: OrganizationBillingConfiguration
    taxes: Optional[List[TaxObject]] = Field(default=None, description="List of default organization taxes")
    __properties: ClassVar[List[str]] = ["lago_id", "name", "created_at", "webhook_url", "webhook_urls", "country", "default_currency", "address_line1", "address_line2", "state", "zipcode", "email", "city", "legal_name", "legal_number", "net_payment_term", "tax_identification_number", "timezone", "billing_configuration", "taxes"]

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
        """Create an instance of OrganizationObject from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of country
        if self.country:
            _dict['country'] = self.country.to_dict()
        # override the default output from pydantic by calling `to_dict()` of default_currency
        if self.default_currency:
            _dict['default_currency'] = self.default_currency.to_dict()
        # override the default output from pydantic by calling `to_dict()` of timezone
        if self.timezone:
            _dict['timezone'] = self.timezone.to_dict()
        # override the default output from pydantic by calling `to_dict()` of billing_configuration
        if self.billing_configuration:
            _dict['billing_configuration'] = self.billing_configuration.to_dict()
        # override the default output from pydantic by calling `to_dict()` of each item in taxes (list)
        _items = []
        if self.taxes:
            for _item in self.taxes:
                if _item:
                    _items.append(_item.to_dict())
            _dict['taxes'] = _items
        # set to None if created_at (nullable) is None
        # and model_fields_set contains the field
        if self.created_at is None and "created_at" in self.model_fields_set:
            _dict['created_at'] = None

        # set to None if webhook_url (nullable) is None
        # and model_fields_set contains the field
        if self.webhook_url is None and "webhook_url" in self.model_fields_set:
            _dict['webhook_url'] = None

        # set to None if webhook_urls (nullable) is None
        # and model_fields_set contains the field
        if self.webhook_urls is None and "webhook_urls" in self.model_fields_set:
            _dict['webhook_urls'] = None

        # set to None if address_line1 (nullable) is None
        # and model_fields_set contains the field
        if self.address_line1 is None and "address_line1" in self.model_fields_set:
            _dict['address_line1'] = None

        # set to None if address_line2 (nullable) is None
        # and model_fields_set contains the field
        if self.address_line2 is None and "address_line2" in self.model_fields_set:
            _dict['address_line2'] = None

        # set to None if state (nullable) is None
        # and model_fields_set contains the field
        if self.state is None and "state" in self.model_fields_set:
            _dict['state'] = None

        # set to None if zipcode (nullable) is None
        # and model_fields_set contains the field
        if self.zipcode is None and "zipcode" in self.model_fields_set:
            _dict['zipcode'] = None

        # set to None if email (nullable) is None
        # and model_fields_set contains the field
        if self.email is None and "email" in self.model_fields_set:
            _dict['email'] = None

        # set to None if city (nullable) is None
        # and model_fields_set contains the field
        if self.city is None and "city" in self.model_fields_set:
            _dict['city'] = None

        # set to None if legal_name (nullable) is None
        # and model_fields_set contains the field
        if self.legal_name is None and "legal_name" in self.model_fields_set:
            _dict['legal_name'] = None

        # set to None if legal_number (nullable) is None
        # and model_fields_set contains the field
        if self.legal_number is None and "legal_number" in self.model_fields_set:
            _dict['legal_number'] = None

        # set to None if tax_identification_number (nullable) is None
        # and model_fields_set contains the field
        if self.tax_identification_number is None and "tax_identification_number" in self.model_fields_set:
            _dict['tax_identification_number'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Dict) -> Self:
        """Create an instance of OrganizationObject from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "lago_id": obj.get("lago_id"),
            "name": obj.get("name"),
            "created_at": obj.get("created_at"),
            "webhook_url": obj.get("webhook_url"),
            "webhook_urls": obj.get("webhook_urls"),
            "country": Country.from_dict(obj.get("country")) if obj.get("country") is not None else None,
            "default_currency": Currency.from_dict(obj.get("default_currency")) if obj.get("default_currency") is not None else None,
            "address_line1": obj.get("address_line1"),
            "address_line2": obj.get("address_line2"),
            "state": obj.get("state"),
            "zipcode": obj.get("zipcode"),
            "email": obj.get("email"),
            "city": obj.get("city"),
            "legal_name": obj.get("legal_name"),
            "legal_number": obj.get("legal_number"),
            "net_payment_term": obj.get("net_payment_term"),
            "tax_identification_number": obj.get("tax_identification_number"),
            "timezone": Timezone.from_dict(obj.get("timezone")) if obj.get("timezone") is not None else None,
            "billing_configuration": OrganizationBillingConfiguration.from_dict(obj.get("billing_configuration")) if obj.get("billing_configuration") is not None else None,
            "taxes": [TaxObject.from_dict(_item) for _item in obj.get("taxes")] if obj.get("taxes") is not None else None
        })
        return _obj

