# About

rCTF-golf is a Python 3 library for developing CTF golfing challenges.

# Installation

## Automatic

```sh
pip3 install rctf-golf
```

## Manual

```sh
git clone https://github.com/redpwn/rCTF-golf.git
cd rCTF-golf.git
pip3 install requirements.txt
./setup.py install
```

# Usage

```python
from rctf import golf

rctf_host = 'https://staging.redpwn.net/' # just the base url of your rCTF installation
challenge_id = 'e0efc6e1-3b04-400a-9d72-a2e2ae02c1f6' # you can also use the challenge name
ctf_start = int(time.time() - 3*3600) # three hours ago
limit_function = lambda x : x

current_limit = golf.calculate_limit(rctf_host, challenge_id, ctf_start, limit_function)
```

## Debugging

When testing your challenge, you can set the `DEBUG` environmental variable to the number of hours after the CTF has started, and rCTF-golf will automatically bypass calculation logic.

**Note:** Take care to ensure users do not control the `DEBUG` environmental variable when running your Python script or they may be able to hijack the limit calculations.

# Documentation

[Click here to see the documentation](rctf/golf/util.py)

