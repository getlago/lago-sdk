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

from lago_api.models.generate_customer_checkout_url200_response import GenerateCustomerCheckoutURL200Response

class TestGenerateCustomerCheckoutURL200Response(unittest.TestCase):
    """GenerateCustomerCheckoutURL200Response unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> GenerateCustomerCheckoutURL200Response:
        """Test GenerateCustomerCheckoutURL200Response
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `GenerateCustomerCheckoutURL200Response`
        """
        model = GenerateCustomerCheckoutURL200Response()
        if include_optional:
            return GenerateCustomerCheckoutURL200Response(
                lago_customer_id = '1a901a90-1a90-1a90-1a90-1a901a901a90',
                external_customer_id = '5eb02857-a71e-4ea2-bcf9-57d3a41bc6ba',
                payment_provider = 'stripe',
                checkout_url = 'https://foo.bar'
            )
        else:
            return GenerateCustomerCheckoutURL200Response(
        )
        """

    def testGenerateCustomerCheckoutURL200Response(self):
        """Test GenerateCustomerCheckoutURL200Response"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
