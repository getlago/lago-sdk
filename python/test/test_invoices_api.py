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

from lago_api.api.invoices_api import InvoicesApi


class TestInvoicesApi(unittest.TestCase):
    """InvoicesApi unit test stubs"""

    def setUp(self) -> None:
        self.api = InvoicesApi()

    def tearDown(self) -> None:
        pass

    def test_create_invoice(self) -> None:
        """Test case for create_invoice

        Create a one-off invoice
        """
        pass

    def test_download_invoice(self) -> None:
        """Test case for download_invoice

        Download an invoice PDF
        """
        pass

    def test_finalize_invoice(self) -> None:
        """Test case for finalize_invoice

        Finalize a draft invoice
        """
        pass

    def test_find_all_invoices(self) -> None:
        """Test case for find_all_invoices

        List all invoices
        """
        pass

    def test_find_invoice(self) -> None:
        """Test case for find_invoice

        Retrieve an invoice
        """
        pass

    def test_refresh_invoice(self) -> None:
        """Test case for refresh_invoice

        Refresh a draft invoice
        """
        pass

    def test_retry_payment(self) -> None:
        """Test case for retry_payment

        Retry an invoice payment
        """
        pass

    def test_update_invoice(self) -> None:
        """Test case for update_invoice

        Update an invoice
        """
        pass


if __name__ == '__main__':
    unittest.main()
