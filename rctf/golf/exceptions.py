#!/usr/bin/env python3
# Author: Aaron Esau <python@aaronesau.com>

'''
Exceptions for use by rCTF-golf
'''

class CTFNotStartedError(RuntimeError):
    def __init__(self, msg = 'The CTF has not started yet', *args, **kwargs):
        super().__init__(msg, *args, **kwargs)

class CTFConfigurationError(ValueError):
    def __init__(self, msg = 'The challenge is misconfigured', *args, **kwargs):
        super().__init__(msg, *args, **kwargs)

