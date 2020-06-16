#!/usr/bin/env python

from distutils.core import setup
import setuptools

setup(name = 'rctf-golf',
    version = '1.0.0',
    description = 'a tool for developing golfing CTF challenges',
    author = 'Aaron Esau',
    author_email = 'aaron@redpwn.net',
    url = 'https://github.com/redpwn/rctf-golf',
    keywords = 'golf golfing ctf rctf ctf-platform',
    install_requires = ['requests'],
    packages = ['rctf.golf']
)
