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

from lago_api.models.coupon_create_input_coupon import CouponCreateInputCoupon

class TestCouponCreateInputCoupon(unittest.TestCase):
    """CouponCreateInputCoupon unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> CouponCreateInputCoupon:
        """Test CouponCreateInputCoupon
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `CouponCreateInputCoupon`
        """
        model = CouponCreateInputCoupon()
        if include_optional:
            return CouponCreateInputCoupon(
                name = 'Startup Deal',
                code = 'startup_deal',
                description = 'I am a coupon description',
                coupon_type = 'fixed_amount',
                amount_cents = 5000,
                amount_currency = None,
                reusable = False,
                percentage_rate = '472888001528021798096225500850762',
                frequency = 'recurring',
                frequency_duration = 6,
                expiration = 'time_limit',
                expiration_at = '2022-08-08T23:59:59Z',
                applies_to = lago_api.models.coupon_base_input_applies_to.CouponBaseInput_applies_to(
                    plan_codes = ["startup_plan"], 
                    billable_metric_codes = [], )
            )
        else:
            return CouponCreateInputCoupon(
                name = 'Startup Deal',
                code = 'startup_deal',
                coupon_type = 'fixed_amount',
                frequency = 'recurring',
        )
        """

    def testCouponCreateInputCoupon(self):
        """Test CouponCreateInputCoupon"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()