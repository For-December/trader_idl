from setuptools import setup, find_packages

setup(
    name='scraping_go',
    version='0.1',
    packages=find_packages(),
    install_requires=[
        'thrift',
    ],
)