#!/usr/bin/env python

from distutils.core import setup
import setuptools

def readme():
    with open('README.md', 'r') as f:
        return f.read()

setup(name = 'rctf-golf',
    version = '1.0.4',
    description = 'a tool for developing golfing CTF challenges',
    long_description = readme(),
    long_description_content_type = 'text/markdown',
    author = 'Aaron Esau',
    author_email = 'aaron@redpwn.net',
    url = 'https://github.com/redpwn/rctf-golf',
    keywords = 'golf golfing ctf rctf ctf-platform',
    install_requires = ['requests'],
    packages = ['rctf.golf']
)
