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

from lago_api.models.add_ons_paginated import AddOnsPaginated

class TestAddOnsPaginated(unittest.TestCase):
    """AddOnsPaginated unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> AddOnsPaginated:
        """Test AddOnsPaginated
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `AddOnsPaginated`
        """
        model = AddOnsPaginated()
        if include_optional:
            return AddOnsPaginated(
                add_ons = [
                    lago_api.models.add_on_object.AddOnObject(
                        lago_id = '1a901a90-1a90-1a90-1a90-1a901a901a90', 
                        name = 'Setup Fee', 
                        invoice_display_name = 'Setup Fee (SF1)', 
                        code = 'setup_fee', 
                        amount_cents = 50000, 
                        amount_currency = null, 
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
                            ], )
                    ],
                meta = lago_api.models.pagination_meta.PaginationMeta(
                    current_page = 2, 
                    next_page = 3, 
                    prev_page = 1, 
                    total_pages = 4, 
                    total_count = 70, )
            )
        else:
            return AddOnsPaginated(
                add_ons = [
                    lago_api.models.add_on_object.AddOnObject(
                        lago_id = '1a901a90-1a90-1a90-1a90-1a901a901a90', 
                        name = 'Setup Fee', 
                        invoice_display_name = 'Setup Fee (SF1)', 
                        code = 'setup_fee', 
                        amount_cents = 50000, 
                        amount_currency = null, 
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
                            ], )
                    ],
                meta = lago_api.models.pagination_meta.PaginationMeta(
                    current_page = 2, 
                    next_page = 3, 
                    prev_page = 1, 
                    total_pages = 4, 
                    total_count = 70, ),
        )
        """

    def testAddOnsPaginated(self):
        """Test AddOnsPaginated"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()