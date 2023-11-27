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

from lago_api.models.billable_metric_update_input import BillableMetricUpdateInput

class TestBillableMetricUpdateInput(unittest.TestCase):
    """BillableMetricUpdateInput unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> BillableMetricUpdateInput:
        """Test BillableMetricUpdateInput
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `BillableMetricUpdateInput`
        """
        model = BillableMetricUpdateInput()
        if include_optional:
            return BillableMetricUpdateInput(
                billable_metric = lago_api.models.billable_metric_base_input.BillableMetricBaseInput(
                    name = 'Storage', 
                    code = 'storage', 
                    description = 'GB of storage used in my application', 
                    recurring = False, 
                    field_name = 'gb', 
                    aggregation_type = 'sum_agg', 
                    weighted_interval = 'seconds', 
                    group = lago_api.models.billable_metric_group.BillableMetricGroup(
                        key = 'region', 
                        values = ["us-east-1","us-east-2","eu-west-1"], ), )
            )
        else:
            return BillableMetricUpdateInput(
                billable_metric = lago_api.models.billable_metric_base_input.BillableMetricBaseInput(
                    name = 'Storage', 
                    code = 'storage', 
                    description = 'GB of storage used in my application', 
                    recurring = False, 
                    field_name = 'gb', 
                    aggregation_type = 'sum_agg', 
                    weighted_interval = 'seconds', 
                    group = lago_api.models.billable_metric_group.BillableMetricGroup(
                        key = 'region', 
                        values = ["us-east-1","us-east-2","eu-west-1"], ), ),
        )
        """

    def testBillableMetricUpdateInput(self):
        """Test BillableMetricUpdateInput"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
