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

from lago_api.models.group_properties_object import GroupPropertiesObject

class TestGroupPropertiesObject(unittest.TestCase):
    """GroupPropertiesObject unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> GroupPropertiesObject:
        """Test GroupPropertiesObject
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `GroupPropertiesObject`
        """
        model = GroupPropertiesObject()
        if include_optional:
            return GroupPropertiesObject(
                group_id = '1a901a90-1a90-1a90-1a90-1a901a901a90',
                invoice_display_name = 'AWS',
                values = None
            )
        else:
            return GroupPropertiesObject(
                group_id = '1a901a90-1a90-1a90-1a90-1a901a901a90',
                values = None,
        )
        """

    def testGroupPropertiesObject(self):
        """Test GroupPropertiesObject"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
