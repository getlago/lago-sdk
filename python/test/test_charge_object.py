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

from lago_api.models.charge_object import ChargeObject

class TestChargeObject(unittest.TestCase):
    """ChargeObject unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> ChargeObject:
        """Test ChargeObject
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `ChargeObject`
        """
        model = ChargeObject()
        if include_optional:
            return ChargeObject(
                lago_id = '1a901a90-1a90-1a90-1a90-1a901a901a90',
                lago_billable_metric_id = '1a901a90-1a90-1a90-1a90-1a901a901a90',
                billable_metric_code = 'requests',
                invoice_display_name = 'Setup',
                created_at = '2022-09-14T16:35:31Z',
                charge_model = 'standard',
                pay_in_advance = True,
                invoiceable = True,
                prorated = False,
                min_amount_cents = 1200,
                properties = None,
                group_properties = [
                    lago_api.models.group_properties_object.GroupPropertiesObject(
                        group_id = '1a901a90-1a90-1a90-1a90-1a901a901a90', 
                        invoice_display_name = 'AWS', 
                        values = null, )
                    ],
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
            return ChargeObject(
                lago_id = '1a901a90-1a90-1a90-1a90-1a901a901a90',
                lago_billable_metric_id = '1a901a90-1a90-1a90-1a90-1a901a901a90',
                billable_metric_code = 'requests',
                created_at = '2022-09-14T16:35:31Z',
                charge_model = 'standard',
        )
        """

    def testChargeObject(self):
        """Test ChargeObject"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
