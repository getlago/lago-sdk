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

from lago_api.models.invoices_paginated import InvoicesPaginated

class TestInvoicesPaginated(unittest.TestCase):
    """InvoicesPaginated unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> InvoicesPaginated:
        """Test InvoicesPaginated
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `InvoicesPaginated`
        """
        model = InvoicesPaginated()
        if include_optional:
            return InvoicesPaginated(
                invoices = [
                    lago_api.models.invoice_object.InvoiceObject(
                        lago_id = '1a901a90-1a90-1a90-1a90-1a901a901a90', 
                        sequential_id = 2, 
                        number = 'LAG-1234-001-002', 
                        issuing_date = 'Sat Apr 30 00:00:00 UTC 2022', 
                        payment_due_date = 'Sat Apr 30 00:00:00 UTC 2022', 
                        net_payment_term = 30, 
                        invoice_type = 'subscription', 
                        status = 'finalized', 
                        payment_status = 'succeeded', 
                        currency = null, 
                        fees_amount_cents = 100, 
                        coupons_amount_cents = 10, 
                        credit_notes_amount_cents = 10, 
                        sub_total_excluding_taxes_amount_cents = 100, 
                        taxes_amount_cents = 20, 
                        sub_total_including_taxes_amount_cents = 120, 
                        prepaid_credit_amount_cents = 0, 
                        total_amount_cents = 100, 
                        version_number = 3, 
                        file_url = 'https://getlago.com/invoice/file', 
                        customer = null, 
                        metadata = [
                            lago_api.models.invoice_metadata_object.InvoiceMetadataObject(
                                lago_id = '1a901a90-1a90-1a90-1a90-1a901a901a90', 
                                key = 'digital_ref_id', 
                                value = 'INV-0123456-98765', 
                                created_at = '2022-04-29T08:59:51Z', )
                            ], 
                        applied_taxes = [
                            lago_api.models.invoice_applied_tax_object.InvoiceAppliedTaxObject(
                                lago_invoice_id = '1a901a90-1a90-1a90-1a90-1a901a901a90', 
                                fees_amount_cents = 20000, )
                            ], )
                    ],
                meta = lago_api.models.pagination_meta.PaginationMeta(
                    current_page = 2, 
                    next_page = 3, 
                    prev_page = 1, 
                    total_pages = 4, 
                    total_count = 70, )
            )
        else:
            return InvoicesPaginated(
                invoices = [
                    lago_api.models.invoice_object.InvoiceObject(
                        lago_id = '1a901a90-1a90-1a90-1a90-1a901a901a90', 
                        sequential_id = 2, 
                        number = 'LAG-1234-001-002', 
                        issuing_date = 'Sat Apr 30 00:00:00 UTC 2022', 
                        payment_due_date = 'Sat Apr 30 00:00:00 UTC 2022', 
                        net_payment_term = 30, 
                        invoice_type = 'subscription', 
                        status = 'finalized', 
                        payment_status = 'succeeded', 
                        currency = null, 
                        fees_amount_cents = 100, 
                        coupons_amount_cents = 10, 
                        credit_notes_amount_cents = 10, 
                        sub_total_excluding_taxes_amount_cents = 100, 
                        taxes_amount_cents = 20, 
                        sub_total_including_taxes_amount_cents = 120, 
                        prepaid_credit_amount_cents = 0, 
                        total_amount_cents = 100, 
                        version_number = 3, 
                        file_url = 'https://getlago.com/invoice/file', 
                        customer = null, 
                        metadata = [
                            lago_api.models.invoice_metadata_object.InvoiceMetadataObject(
                                lago_id = '1a901a90-1a90-1a90-1a90-1a901a901a90', 
                                key = 'digital_ref_id', 
                                value = 'INV-0123456-98765', 
                                created_at = '2022-04-29T08:59:51Z', )
                            ], 
                        applied_taxes = [
                            lago_api.models.invoice_applied_tax_object.InvoiceAppliedTaxObject(
                                lago_invoice_id = '1a901a90-1a90-1a90-1a90-1a901a901a90', 
                                fees_amount_cents = 20000, )
                            ], )
                    ],
                meta = lago_api.models.pagination_meta.PaginationMeta(
                    current_page = 2, 
                    next_page = 3, 
                    prev_page = 1, 
                    total_pages = 4, 
                    total_count = 70, ),
        )
        """

    def testInvoicesPaginated(self):
        """Test InvoicesPaginated"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
