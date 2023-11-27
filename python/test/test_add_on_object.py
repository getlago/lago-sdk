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

from lago_api.models.add_on_object import AddOnObject

class TestAddOnObject(unittest.TestCase):
    """AddOnObject unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> AddOnObject:
        """Test AddOnObject
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `AddOnObject`
        """
        model = AddOnObject()
        if include_optional:
            return AddOnObject(
                lago_id = '1a901a90-1a90-1a90-1a90-1a901a901a90',
                name = 'Setup Fee',
                invoice_display_name = 'Setup Fee (SF1)',
                code = 'setup_fee',
                amount_cents = 50000,
                amount_currency = None,
                description = 'Implementation fee for new customers.',
                created_at = '2022-04-29T08:59:51Z',
                taxes = [
                    lago_api.models.tax_object.TaxObject(
                        lago_id = '1a901a90-1a90-1a90-1a90-1a901a901a90', 
                        name = 'TVA', 
                        code = 'french_standard_vat', 
                        description = 'French standard VAT', 
                        rate = 20, 
                        applied_to_organization = True, 
                        add_ons_count = 0, 
                        charges_count = 0, 
                        customers_count = 0, 
                        plans_count = 0, 
                        created_at = '2023-07-06T14:35:58Z', )
                    ]
            )
        else:
            return AddOnObject(
                lago_id = '1a901a90-1a90-1a90-1a90-1a901a901a90',
                name = 'Setup Fee',
                code = 'setup_fee',
                amount_cents = 50000,
                amount_currency = None,
                created_at = '2022-04-29T08:59:51Z',
        )
        """

    def testAddOnObject(self):
        """Test AddOnObject"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
