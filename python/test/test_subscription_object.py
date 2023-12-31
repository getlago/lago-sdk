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

from lago_api.models.subscription_object import SubscriptionObject

class TestSubscriptionObject(unittest.TestCase):
    """SubscriptionObject unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> SubscriptionObject:
        """Test SubscriptionObject
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `SubscriptionObject`
        """
        model = SubscriptionObject()
        if include_optional:
            return SubscriptionObject(
                lago_id = '1a901a90-1a90-1a90-1a90-1a901a901a90',
                external_id = '5eb02857-a71e-4ea2-bcf9-57d3a41bc6ba',
                lago_customer_id = '1a901a90-1a90-1a90-1a90-1a901a901a90',
                external_customer_id = '5eb02857-a71e-4ea2-bcf9-57d3a41bc6ba',
                billing_time = 'anniversary',
                name = 'Repository A',
                plan_code = 'premium',
                status = 'active',
                created_at = '2022-08-08T00:00Z',
                canceled_at = '2022-09-14T16:35:31Z',
                started_at = '2022-08-08T00:00Z',
                ending_at = '2022-10-08T00:00Z',
                subscription_at = '2022-08-08T00:00Z',
                terminated_at = '2022-09-14T16:35:31Z',
                previous_plan_code = '',
                next_plan_code = '',
                downgrade_plan_date = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f')
            )
        else:
            return SubscriptionObject(
                lago_id = '1a901a90-1a90-1a90-1a90-1a901a901a90',
                external_id = '5eb02857-a71e-4ea2-bcf9-57d3a41bc6ba',
                lago_customer_id = '1a901a90-1a90-1a90-1a90-1a901a901a90',
                external_customer_id = '5eb02857-a71e-4ea2-bcf9-57d3a41bc6ba',
                billing_time = 'anniversary',
                plan_code = 'premium',
                status = 'active',
                created_at = '2022-08-08T00:00Z',
                subscription_at = '2022-08-08T00:00Z',
        )
        """

    def testSubscriptionObject(self):
        """Test SubscriptionObject"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
