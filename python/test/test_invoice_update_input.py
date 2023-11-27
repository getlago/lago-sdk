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

from lago_api.models.invoice_update_input import InvoiceUpdateInput

class TestInvoiceUpdateInput(unittest.TestCase):
    """InvoiceUpdateInput unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> InvoiceUpdateInput:
        """Test InvoiceUpdateInput
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `InvoiceUpdateInput`
        """
        model = InvoiceUpdateInput()
        if include_optional:
            return InvoiceUpdateInput(
                invoice = lago_api.models.invoice_update_input_invoice.InvoiceUpdateInput_invoice(
                    payment_status = 'succeeded', 
                    metadata = [
                        lago_api.models.invoice_update_input_invoice_metadata_inner.InvoiceUpdateInput_invoice_metadata_inner(
                            id = '1a901a90-1a90-1a90-1a90-1a901a901a90', 
                            key = 'digital_ref_id', 
                            value = 'INV-0123456-98765', )
                        ], )
            )
        else:
            return InvoiceUpdateInput(
                invoice = lago_api.models.invoice_update_input_invoice.InvoiceUpdateInput_invoice(
                    payment_status = 'succeeded', 
                    metadata = [
                        lago_api.models.invoice_update_input_invoice_metadata_inner.InvoiceUpdateInput_invoice_metadata_inner(
                            id = '1a901a90-1a90-1a90-1a90-1a901a901a90', 
                            key = 'digital_ref_id', 
                            value = 'INV-0123456-98765', )
                        ], ),
        )
        """

    def testInvoiceUpdateInput(self):
        """Test InvoiceUpdateInput"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
