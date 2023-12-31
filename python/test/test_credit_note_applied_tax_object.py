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

from lago_api.models.credit_note_applied_tax_object import CreditNoteAppliedTaxObject

class TestCreditNoteAppliedTaxObject(unittest.TestCase):
    """CreditNoteAppliedTaxObject unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> CreditNoteAppliedTaxObject:
        """Test CreditNoteAppliedTaxObject
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `CreditNoteAppliedTaxObject`
        """
        model = CreditNoteAppliedTaxObject()
        if include_optional:
            return CreditNoteAppliedTaxObject(
                lago_credit_note_id = '1a901a90-1a90-1a90-1a90-1a901a901a90',
                base_amount_cents = 100,
                lago_id = '1a901a90-1a90-1a90-1a90-1a901a901a90',
                lago_tax_id = '1a901a90-1a90-1a90-1a90-1a901a901a90',
                tax_name = 'TVA',
                tax_code = 'french_standard_vat',
                tax_rate = 20,
                tax_description = 'French standard VAT',
                amount_cents = 2000,
                amount_currency = None,
                created_at = '2022-09-14T16:35:31Z'
            )
        else:
            return CreditNoteAppliedTaxObject(
        )
        """

    def testCreditNoteAppliedTaxObject(self):
        """Test CreditNoteAppliedTaxObject"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
