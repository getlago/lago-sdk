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

from lago_api.api.billable_metrics_api import BillableMetricsApi


class TestBillableMetricsApi(unittest.TestCase):
    """BillableMetricsApi unit test stubs"""

    def setUp(self) -> None:
        self.api = BillableMetricsApi()

    def tearDown(self) -> None:
        pass

    def test_create_billable_metric(self) -> None:
        """Test case for create_billable_metric

        Create a billable metric
        """
        pass

    def test_destroy_billable_metric(self) -> None:
        """Test case for destroy_billable_metric

        Delete a billable metric
        """
        pass

    def test_find_all_billable_metric_groups(self) -> None:
        """Test case for find_all_billable_metric_groups

        Find a billable metric's groups
        """
        pass

    def test_find_all_billable_metrics(self) -> None:
        """Test case for find_all_billable_metrics

        List all billable metrics
        """
        pass

    def test_find_billable_metric(self) -> None:
        """Test case for find_billable_metric

        Retrieve a billable metric
        """
        pass

    def test_update_billable_metric(self) -> None:
        """Test case for update_billable_metric

        Update a billable metric
        """
        pass


if __name__ == '__main__':
    unittest.main()
