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

from lago_api.models.applied_coupon_object_extended import AppliedCouponObjectExtended

class TestAppliedCouponObjectExtended(unittest.TestCase):
    """AppliedCouponObjectExtended unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> AppliedCouponObjectExtended:
        """Test AppliedCouponObjectExtended
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `AppliedCouponObjectExtended`
        """
        model = AppliedCouponObjectExtended()
        if include_optional:
            return AppliedCouponObjectExtended(
                lago_id = '1a901a90-1a90-1a90-1a90-1a901a901a90',
                lago_coupon_id = '1a901a90-1a90-1a90-1a90-1a901a901a90',
                coupon_code = 'startup_deal',
                coupon_name = 'Startup Deal',
                lago_customer_id = '1a901a90-1a90-1a90-1a90-1a901a901a90',
                external_customer_id = '5eb02857-a71e-4ea2-bcf9-57d3a41bc6ba',
                status = 'active',
                amount_cents = 2000,
                amount_cents_remaining = 50,
                amount_currency = None,
                percentage_rate = '472888001528021798096225500850762',
                frequency = 'recurring',
                frequency_duration = 3,
                frequency_duration_remaining = 1,
                expiration_at = '2022-04-29T08:59:51Z',
                created_at = '2022-04-29T08:59:51Z',
                terminated_at = '2022-04-29T08:59:51Z',
                credits = [
                    lago_api.models.credit_object.CreditObject(
                        lago_id = '1a901a90-1a90-1a90-1a90-1a901a901a90', 
                        amount_cents = 1200, 
                        amount_currency = null, 
                        before_taxes = False, 
                        item = lago_api.models.credit_object_item.CreditObject_item(
                            lago_item_id = '1a901a90-1a90-1a90-1a90-1a901a901a90', 
                            type = 'coupon', 
                            code = 'startup_deal', 
                            name = 'Startup Deal', ), 
                        invoice = lago_api.models.credit_object_invoice.CreditObject_invoice(
                            lago_id = '1a901a90-1a90-1a90-1a90-1a901a901a90', 
                            payment_status = 'succeeded', ), )
                    ]
            )
        else:
            return AppliedCouponObjectExtended(
                lago_id = '1a901a90-1a90-1a90-1a90-1a901a901a90',
                lago_coupon_id = '1a901a90-1a90-1a90-1a90-1a901a901a90',
                coupon_code = 'startup_deal',
                coupon_name = 'Startup Deal',
                lago_customer_id = '1a901a90-1a90-1a90-1a90-1a901a901a90',
                external_customer_id = '5eb02857-a71e-4ea2-bcf9-57d3a41bc6ba',
                status = 'active',
                frequency = 'recurring',
                created_at = '2022-04-29T08:59:51Z',
                credits = [
                    lago_api.models.credit_object.CreditObject(
                        lago_id = '1a901a90-1a90-1a90-1a90-1a901a901a90', 
                        amount_cents = 1200, 
                        amount_currency = null, 
                        before_taxes = False, 
                        item = lago_api.models.credit_object_item.CreditObject_item(
                            lago_item_id = '1a901a90-1a90-1a90-1a90-1a901a901a90', 
                            type = 'coupon', 
                            code = 'startup_deal', 
                            name = 'Startup Deal', ), 
                        invoice = lago_api.models.credit_object_invoice.CreditObject_invoice(
                            lago_id = '1a901a90-1a90-1a90-1a90-1a901a901a90', 
                            payment_status = 'succeeded', ), )
                    ],
        )
        """

    def testAppliedCouponObjectExtended(self):
        """Test AppliedCouponObjectExtended"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
