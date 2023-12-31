# coding: utf-8

"""
    Lago API documentation

    Lago API allows your application to push customer information and metrics (events) from your application to the billing application.

    The version of the OpenAPI document: 0.52.0-beta
    Contact: tech@getlago.com
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


import unittest
import datetime

from lago_api.models.credit_note_item_object import CreditNoteItemObject

class TestCreditNoteItemObject(unittest.TestCase):
    """CreditNoteItemObject unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> CreditNoteItemObject:
        """Test CreditNoteItemObject
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `CreditNoteItemObject`
        """
        model = CreditNoteItemObject()
        if include_optional:
            return CreditNoteItemObject(
                lago_id = '1a901a90-1a90-1a90-1a90-1a901a901a90',
                amount_cents = 100,
                amount_currency = None,
                fee = None
            )
        else:
            return CreditNoteItemObject(
                lago_id = '1a901a90-1a90-1a90-1a90-1a901a901a90',
                amount_cents = 100,
                amount_currency = None,
                fee = None,
        )
        """

    def testCreditNoteItemObject(self):
        """Test CreditNoteItemObject"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
