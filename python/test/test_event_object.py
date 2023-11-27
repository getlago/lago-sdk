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

from lago_api.models.event_object import EventObject

class TestEventObject(unittest.TestCase):
    """EventObject unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> EventObject:
        """Test EventObject
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `EventObject`
        """
        model = EventObject()
        if include_optional:
            return EventObject(
                lago_id = '1a901a90-1a90-1a90-1a90-1a901a901a90',
                transaction_id = 'transaction_1234567890',
                lago_customer_id = '1a901a90-1a90-1a90-1a90-1a901a901a90',
                external_customer_id = '5eb02857-a71e-4ea2-bcf9-57d3a41bc6ba',
                code = 'storage',
                timestamp = '2022-04-29T08:59:51.123Z',
                properties = {"gb":10},
                lago_subscription_id = '1a901a90-1a90-1a90-1a90-1a901a901a90',
                external_subscription_id = 'sub_1234567890',
                created_at = '2022-04-29T08:59:51Z'
            )
        else:
            return EventObject(
                lago_id = '1a901a90-1a90-1a90-1a90-1a901a901a90',
                transaction_id = 'transaction_1234567890',
                lago_customer_id = '1a901a90-1a90-1a90-1a90-1a901a901a90',
                external_customer_id = '5eb02857-a71e-4ea2-bcf9-57d3a41bc6ba',
                code = 'storage',
                timestamp = '2022-04-29T08:59:51.123Z',
                lago_subscription_id = '1a901a90-1a90-1a90-1a90-1a901a901a90',
                external_subscription_id = 'sub_1234567890',
                created_at = '2022-04-29T08:59:51Z',
        )
        """

    def testEventObject(self):
        """Test EventObject"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
