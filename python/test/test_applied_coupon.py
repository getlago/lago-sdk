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

from lago_api.models.applied_coupon import AppliedCoupon

class TestAppliedCoupon(unittest.TestCase):
    """AppliedCoupon unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> AppliedCoupon:
        """Test AppliedCoupon
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `AppliedCoupon`
        """
        model = AppliedCoupon()
        if include_optional:
            return AppliedCoupon(
                applied_coupon = lago_api.models.applied_coupon_object.AppliedCouponObject(
                    lago_id = '1a901a90-1a90-1a90-1a90-1a901a901a90', 
                    lago_coupon_id = '1a901a90-1a90-1a90-1a90-1a901a901a90', 
                    coupon_code = 'startup_deal', 
                    coupon_name = 'Startup Deal', 
                    lago_customer_id = '1a901a90-1a90-1a90-1a90-1a901a901a90', 
                    external_customer_id = '5eb02857-a71e-4ea2-bcf9-57d3a41bc6ba', 
                    status = 'active', 
                    amount_cents = 2000, 
                    amount_cents_remaining = 50, 
                    amount_currency = null, 
                    percentage_rate = '472888001528021798096225500850762', 
                    frequency = 'recurring', 
                    frequency_duration = 3, 
                    frequency_duration_remaining = 1, 
                    expiration_at = '2022-04-29T08:59:51Z', 
                    created_at = '2022-04-29T08:59:51Z', 
                    terminated_at = '2022-04-29T08:59:51Z', )
            )
        else:
            return AppliedCoupon(
                applied_coupon = lago_api.models.applied_coupon_object.AppliedCouponObject(
                    lago_id = '1a901a90-1a90-1a90-1a90-1a901a901a90', 
                    lago_coupon_id = '1a901a90-1a90-1a90-1a90-1a901a901a90', 
                    coupon_code = 'startup_deal', 
                    coupon_name = 'Startup Deal', 
                    lago_customer_id = '1a901a90-1a90-1a90-1a90-1a901a901a90', 
                    external_customer_id = '5eb02857-a71e-4ea2-bcf9-57d3a41bc6ba', 
                    status = 'active', 
                    amount_cents = 2000, 
                    amount_cents_remaining = 50, 
                    amount_currency = null, 
                    percentage_rate = '472888001528021798096225500850762', 
                    frequency = 'recurring', 
                    frequency_duration = 3, 
                    frequency_duration_remaining = 1, 
                    expiration_at = '2022-04-29T08:59:51Z', 
                    created_at = '2022-04-29T08:59:51Z', 
                    terminated_at = '2022-04-29T08:59:51Z', ),
        )
        """

    def testAppliedCoupon(self):
        """Test AppliedCoupon"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
