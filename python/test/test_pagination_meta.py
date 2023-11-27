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

from lago_api.models.pagination_meta import PaginationMeta

class TestPaginationMeta(unittest.TestCase):
    """PaginationMeta unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> PaginationMeta:
        """Test PaginationMeta
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `PaginationMeta`
        """
        model = PaginationMeta()
        if include_optional:
            return PaginationMeta(
                current_page = 2,
                next_page = 3,
                prev_page = 1,
                total_pages = 4,
                total_count = 70
            )
        else:
            return PaginationMeta(
                current_page = 2,
                total_pages = 4,
                total_count = 70,
        )
        """

    def testPaginationMeta(self):
        """Test PaginationMeta"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
