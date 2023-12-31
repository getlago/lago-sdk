# coding: utf-8

"""
    Lago API documentation

    Lago API allows your application to push customer information and metrics (events) from your application to the billing application.

    The version of the OpenAPI document: 0.52.0-beta
    Contact: tech@getlago.com
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


from setuptools import setup, find_packages  # noqa: H301

# To install the library, run the following
#
# python setup.py install
#
# prerequisite: setuptools
# http://pypi.python.org/pypi/setuptools
NAME = "lago_api"
VERSION = "1.0.0"
PYTHON_REQUIRES = ">=3.7"
REQUIRES = [
    "urllib3 >= 1.25.3, < 2.1.0",
    "python-dateutil",
    "pydantic >= 2",
    "typing-extensions >= 4.7.1",
]

setup(
    name=NAME,
    version=VERSION,
    description="Lago API documentation",
    author="OpenAPI Generator community",
    author_email="tech@getlago.com",
    url="https://github.com/getlago/sdk",
    keywords=["OpenAPI", "OpenAPI-Generator", "Lago API documentation"],
    install_requires=REQUIRES,
    packages=find_packages(exclude=["test", "tests"]),
    include_package_data=True,
    license="AGPLv3",
    long_description_content_type='text/markdown',
    long_description="""\
    Lago API allows your application to push customer information and metrics (events) from your application to the billing application.
    """,  # noqa: E501
    package_data={"lago_api": ["py.typed"]},
)
