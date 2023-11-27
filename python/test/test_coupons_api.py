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

from lago_api.api.coupons_api import CouponsApi


class TestCouponsApi(unittest.TestCase):
    """CouponsApi unit test stubs"""

    def setUp(self) -> None:
        self.api = CouponsApi()

    def tearDown(self) -> None:
        pass

    def test_apply_coupon(self) -> None:
        """Test case for apply_coupon

        Apply a coupon to a customer
        """
        pass

    def test_create_coupon(self) -> None:
        """Test case for create_coupon

        Create a coupon
        """
        pass

    def test_delete_applied_coupon(self) -> None:
        """Test case for delete_applied_coupon

        Delete an applied coupon
        """
        pass

    def test_destroy_coupon(self) -> None:
        """Test case for destroy_coupon

        Delete a coupon
        """
        pass

    def test_find_all_applied_coupons(self) -> None:
        """Test case for find_all_applied_coupons

        List all applied coupons
        """
        pass

    def test_find_all_coupons(self) -> None:
        """Test case for find_all_coupons

        List all coupons
        """
        pass

    def test_find_coupon(self) -> None:
        """Test case for find_coupon

        Retrieve a coupon
        """
        pass

    def test_update_coupon(self) -> None:
        """Test case for update_coupon

        Update a coupon
        """
        pass


if __name__ == '__main__':
    unittest.main()
