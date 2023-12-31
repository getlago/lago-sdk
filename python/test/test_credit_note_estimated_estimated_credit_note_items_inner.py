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

from lago_api.models.credit_note_estimated_estimated_credit_note_items_inner import CreditNoteEstimatedEstimatedCreditNoteItemsInner

class TestCreditNoteEstimatedEstimatedCreditNoteItemsInner(unittest.TestCase):
    """CreditNoteEstimatedEstimatedCreditNoteItemsInner unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> CreditNoteEstimatedEstimatedCreditNoteItemsInner:
        """Test CreditNoteEstimatedEstimatedCreditNoteItemsInner
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `CreditNoteEstimatedEstimatedCreditNoteItemsInner`
        """
        model = CreditNoteEstimatedEstimatedCreditNoteItemsInner()
        if include_optional:
            return CreditNoteEstimatedEstimatedCreditNoteItemsInner(
                amount_cents = 100,
                lago_fee_id = '1a901a90-1a90-1a90-1a90-1a901a901a90'
            )
        else:
            return CreditNoteEstimatedEstimatedCreditNoteItemsInner(
                amount_cents = 100,
                lago_fee_id = '1a901a90-1a90-1a90-1a90-1a901a901a90',
        )
        """

    def testCreditNoteEstimatedEstimatedCreditNoteItemsInner(self):
        """Test CreditNoteEstimatedEstimatedCreditNoteItemsInner"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
