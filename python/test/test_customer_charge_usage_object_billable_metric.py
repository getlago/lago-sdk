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

from lago_api.models.customer_charge_usage_object_billable_metric import CustomerChargeUsageObjectBillableMetric

class TestCustomerChargeUsageObjectBillableMetric(unittest.TestCase):
    """CustomerChargeUsageObjectBillableMetric unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> CustomerChargeUsageObjectBillableMetric:
        """Test CustomerChargeUsageObjectBillableMetric
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `CustomerChargeUsageObjectBillableMetric`
        """
        model = CustomerChargeUsageObjectBillableMetric()
        if include_optional:
            return CustomerChargeUsageObjectBillableMetric(
                lago_id = '1a901a90-1a90-1a90-1a90-1a901a901a90',
                name = 'Storage',
                code = 'storage',
                aggregation_type = 'sum_agg'
            )
        else:
            return CustomerChargeUsageObjectBillableMetric(
                lago_id = '1a901a90-1a90-1a90-1a90-1a901a901a90',
                name = 'Storage',
                code = 'storage',
                aggregation_type = 'sum_agg',
        )
        """

    def testCustomerChargeUsageObjectBillableMetric(self):
        """Test CustomerChargeUsageObjectBillableMetric"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
