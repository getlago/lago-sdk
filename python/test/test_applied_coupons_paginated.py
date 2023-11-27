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

from lago_api.models.applied_coupons_paginated import AppliedCouponsPaginated

class TestAppliedCouponsPaginated(unittest.TestCase):
    """AppliedCouponsPaginated unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> AppliedCouponsPaginated:
        """Test AppliedCouponsPaginated
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `AppliedCouponsPaginated`
        """
        model = AppliedCouponsPaginated()
        if include_optional:
            return AppliedCouponsPaginated(
                applied_coupons = [
                    null
                    ],
                meta = lago_api.models.pagination_meta.PaginationMeta(
                    current_page = 2, 
                    next_page = 3, 
                    prev_page = 1, 
                    total_pages = 4, 
                    total_count = 70, )
            )
        else:
            return AppliedCouponsPaginated(
                applied_coupons = [
                    null
                    ],
                meta = lago_api.models.pagination_meta.PaginationMeta(
                    current_page = 2, 
                    next_page = 3, 
                    prev_page = 1, 
                    total_pages = 4, 
                    total_count = 70, ),
        )
        """

    def testAppliedCouponsPaginated(self):
        """Test AppliedCouponsPaginated"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
