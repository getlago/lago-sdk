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

from lago_api.models.api_error_unprocessable_entity import ApiErrorUnprocessableEntity

class TestApiErrorUnprocessableEntity(unittest.TestCase):
    """ApiErrorUnprocessableEntity unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> ApiErrorUnprocessableEntity:
        """Test ApiErrorUnprocessableEntity
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `ApiErrorUnprocessableEntity`
        """
        model = ApiErrorUnprocessableEntity()
        if include_optional:
            return ApiErrorUnprocessableEntity(
                status = 422,
                error = 'Unprocessable entity',
                code = 'validation_errors',
                error_details = None
            )
        else:
            return ApiErrorUnprocessableEntity(
                status = 422,
                error = 'Unprocessable entity',
                code = 'validation_errors',
                error_details = None,
        )
        """

    def testApiErrorUnprocessableEntity(self):
        """Test ApiErrorUnprocessableEntity"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
