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

from lago_api.models.wallet_transactions import WalletTransactions

class TestWalletTransactions(unittest.TestCase):
    """WalletTransactions unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> WalletTransactions:
        """Test WalletTransactions
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `WalletTransactions`
        """
        model = WalletTransactions()
        if include_optional:
            return WalletTransactions(
                wallet_transactions = [
                    lago_api.models.wallet_transaction_object.WalletTransactionObject(
                        lago_id = '1a901a90-1a90-1a90-1a90-1a901a901a90', 
                        lago_wallet_id = '1a901a90-1a90-1a90-1a90-1a901a901a90', 
                        status = 'settled', 
                        transaction_type = 'inbound', 
                        amount = '10.0', 
                        credit_amount = '100.0', 
                        settled_at = '2022-04-29T08:59:51Z', 
                        created_at = '2022-04-29T08:59:51Z', )
                    ]
            )
        else:
            return WalletTransactions(
                wallet_transactions = [
                    lago_api.models.wallet_transaction_object.WalletTransactionObject(
                        lago_id = '1a901a90-1a90-1a90-1a90-1a901a901a90', 
                        lago_wallet_id = '1a901a90-1a90-1a90-1a90-1a901a901a90', 
                        status = 'settled', 
                        transaction_type = 'inbound', 
                        amount = '10.0', 
                        credit_amount = '100.0', 
                        settled_at = '2022-04-29T08:59:51Z', 
                        created_at = '2022-04-29T08:59:51Z', )
                    ],
        )
        """

    def testWalletTransactions(self):
        """Test WalletTransactions"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
