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

from lago_api.models.customer_past_usage import CustomerPastUsage

class TestCustomerPastUsage(unittest.TestCase):
    """CustomerPastUsage unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> CustomerPastUsage:
        """Test CustomerPastUsage
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `CustomerPastUsage`
        """
        model = CustomerPastUsage()
        if include_optional:
            return CustomerPastUsage(
                usage_periods = [
                    lago_api.models.customer_usage.CustomerUsage(
                        customer_usage = lago_api.models.customer_usage_object.CustomerUsageObject(
                            from_datetime = '2022-07-01T00:00Z', 
                            to_datetime = '2022-07-31T23:59:59Z', 
                            issuing_date = '2022-08-01T23:59:59Z', 
                            lago_invoice_id = '1a901a90-1a90-1a90-1a90-1a901a901a90', 
                            currency = null, 
                            amount_cents = 123, 
                            taxes_amount_cents = 200, 
                            total_amount_cents = 123, 
                            charges_usage = [
                                lago_api.models.customer_charge_usage_object.CustomerChargeUsageObject(
                                    units = '1.0', 
                                    events_count = 10, 
                                    amount_cents = 123, 
                                    amount_currency = null, 
                                    charge = lago_api.models.customer_charge_usage_object_charge.CustomerChargeUsageObject_charge(
                                        lago_id = '1a901a90-1a90-1a90-1a90-1a901a901a90', 
                                        charge_model = 'graduated', 
                                        invoice_display_name = 'Setup', ), 
                                    billable_metric = lago_api.models.customer_charge_usage_object_billable_metric.CustomerChargeUsageObject_billable_metric(
                                        lago_id = '1a901a90-1a90-1a90-1a90-1a901a901a90', 
                                        name = 'Storage', 
                                        code = 'storage', 
                                        aggregation_type = 'sum_agg', ), 
                                    groups = [
                                        lago_api.models.customer_charge_usage_object_groups_inner.CustomerChargeUsageObject_groups_inner(
                                            lago_id = '1a901a90-1a90-1a90-1a90-1a901a901a90', 
                                            key = '', 
                                            value = 'europe', 
                                            units = '0.9', 
                                            events_count = 10, 
                                            amount_cents = 1000, )
                                        ], )
                                ], ), )
                    ],
                meta = lago_api.models.pagination_meta.PaginationMeta(
                    current_page = 2, 
                    next_page = 3, 
                    prev_page = 1, 
                    total_pages = 4, 
                    total_count = 70, )
            )
        else:
            return CustomerPastUsage(
                usage_periods = [
                    lago_api.models.customer_usage.CustomerUsage(
                        customer_usage = lago_api.models.customer_usage_object.CustomerUsageObject(
                            from_datetime = '2022-07-01T00:00Z', 
                            to_datetime = '2022-07-31T23:59:59Z', 
                            issuing_date = '2022-08-01T23:59:59Z', 
                            lago_invoice_id = '1a901a90-1a90-1a90-1a90-1a901a901a90', 
                            currency = null, 
                            amount_cents = 123, 
                            taxes_amount_cents = 200, 
                            total_amount_cents = 123, 
                            charges_usage = [
                                lago_api.models.customer_charge_usage_object.CustomerChargeUsageObject(
                                    units = '1.0', 
                                    events_count = 10, 
                                    amount_cents = 123, 
                                    amount_currency = null, 
                                    charge = lago_api.models.customer_charge_usage_object_charge.CustomerChargeUsageObject_charge(
                                        lago_id = '1a901a90-1a90-1a90-1a90-1a901a901a90', 
                                        charge_model = 'graduated', 
                                        invoice_display_name = 'Setup', ), 
                                    billable_metric = lago_api.models.customer_charge_usage_object_billable_metric.CustomerChargeUsageObject_billable_metric(
                                        lago_id = '1a901a90-1a90-1a90-1a90-1a901a901a90', 
                                        name = 'Storage', 
                                        code = 'storage', 
                                        aggregation_type = 'sum_agg', ), 
                                    groups = [
                                        lago_api.models.customer_charge_usage_object_groups_inner.CustomerChargeUsageObject_groups_inner(
                                            lago_id = '1a901a90-1a90-1a90-1a90-1a901a901a90', 
                                            key = '', 
                                            value = 'europe', 
                                            units = '0.9', 
                                            events_count = 10, 
                                            amount_cents = 1000, )
                                        ], )
                                ], ), )
                    ],
                meta = lago_api.models.pagination_meta.PaginationMeta(
                    current_page = 2, 
                    next_page = 3, 
                    prev_page = 1, 
                    total_pages = 4, 
                    total_count = 70, ),
        )
        """

    def testCustomerPastUsage(self):
        """Test CustomerPastUsage"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
