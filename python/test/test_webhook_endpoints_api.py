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

from lago_api.api.webhook_endpoints_api import WebhookEndpointsApi


class TestWebhookEndpointsApi(unittest.TestCase):
    """WebhookEndpointsApi unit test stubs"""

    def setUp(self) -> None:
        self.api = WebhookEndpointsApi()

    def tearDown(self) -> None:
        pass

    def test_create_webhook_endpoint(self) -> None:
        """Test case for create_webhook_endpoint

        Create a webhook_endpoint
        """
        pass

    def test_destroy_webhook_endpoint(self) -> None:
        """Test case for destroy_webhook_endpoint

        Delete a webhook endpoint
        """
        pass

    def test_find_all_webhook_endpoints(self) -> None:
        """Test case for find_all_webhook_endpoints

        List all webhook endpoints
        """
        pass

    def test_find_webhook_endpoint(self) -> None:
        """Test case for find_webhook_endpoint

        Retrieve a webhook endpoint
        """
        pass

    def test_update_webhook_endpoint(self) -> None:
        """Test case for update_webhook_endpoint

        Update a webhook endpoint
        """
        pass


if __name__ == '__main__':
    unittest.main()
