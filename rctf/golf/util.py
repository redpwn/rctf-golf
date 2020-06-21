#!/usr/bin/env python3
# Author: Aaron Esau <python@aaronesau.com>

'''
This module helps with developing golf challenges when using the rCTF platform
'''

import datetime, os, warnings, json
import requests, urllib.parse
from typing import Callable, Union

from .exceptions import CTFNotStartedError, CTFConfigurationError


# TODO: pull the start_datetime from the API once we support that

'''
Determines the limit for a golf challenge based on a scoring function, whether or not the challenge is solved, and the current time.

:param rctf_base_url: the url to the base of the rCTF installation (e.g. https://2020.redpwn.net/)
:param rctf_challenge_id: the challenge name or ID to lookup
:param start_datetime: the unix timestamp (int) or datetime object representing the start of the CTF event
:param limit_function: a lambda function which takes the number of hours since the CTF start (rounded down) and returns the limit
:return: the current limit for the golfing challenge
'''
def calculate_limit(rctf_base_url: str, rctf_challenge_id: str, start_datetime: Union[int, datetime.datetime], limit_function: Callable[[int], Union[float, int]], debug: Union[bool, int] = os.environ.get('DEBUG')) -> Union[float, int]:
    # check configuration
    if isinstance(start_datetime, int):
        start_datetime = datetime.datetime.utcfromtimestamp(start_datetime)

    assert isinstance(start_datetime, datetime.datetime) == True
    current_datetime = datetime.datetime.utcnow()

    # calculate the number of hours to plug into formula
    hours_in_ctf = 0
    if debug != None: # explicit None to allow debug == 0
        if isinstance(debug, bool):
            debug = 24 * 3 # 3 days default
        
        hours_in_ctf = int(debug)
        warnings.warn('DEBUG mode is enabled. rCTF-golf will assume the challenge is %d hours into the CTF.' % hours_in_ctf)
    else:
        if current_datetime < start_datetime:
            raise CTFNotStartedError()

        # debug mode is disabled, calculate number of hours
        response = json.loads(requests.get('{rctf_base_url}/api/v1/challs/{rctf_challenge_id}/solves?limit=1&offset=0'.format(
            rctf_base_url = rctf_base_url.rstrip('/'),
            rctf_challenge_id = urllib.parse.quote(rctf_challenge_id)
        )).text)

        if response.get('kind') != 'goodChallengeSolves':
            raise CTFConfigurationError(response.get('message', 'The challenge is misconfigured'))

        solves = response['data']['solves']
        if not solves:
            # nobody has solved yet, so use current time delta
            solved_datetime = current_datetime
        else:
            # at least one solve. pull the latest's timestamp
            solved_timestamp = solves[0]['createdAt'] # there should always be one solve because of the limit
            solved_datetime = datetime.datetime.utcfromtimestamp(solved_timestamp // 1000) # keep in utc

        if solved_datetime < start_datetime:
            raise CTFConfigurationError('The server thinks the CTF has started, but the challenge believes it has not')

        hours_in_ctf = (solved_datetime - start_datetime).seconds // (60 * 60)

    assert hours_in_ctf >= 0
    return limit_function(int(hours_in_ctf)) # round down
