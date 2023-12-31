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

from lago_api.models.mrr_object import MrrObject

class TestMrrObject(unittest.TestCase):
    """MrrObject unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> MrrObject:
        """Test MrrObject
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `MrrObject`
        """
        model = MrrObject()
        if include_optional:
            return MrrObject(
                month = '2023-11-01T00:00:00.000Z',
                amount_cents = 50000,
                currency = None
            )
        else:
            return MrrObject(
                month = '2023-11-01T00:00:00.000Z',
        )
        """

    def testMrrObject(self):
        """Test MrrObject"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
