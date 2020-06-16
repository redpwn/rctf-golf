# About

rCTF-golf is a Python 3 library for developing CTF golfing challenges.

# Installation

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

ctf_start = int(time.time() - 3*3600) # three hours ago
golf.calculate_limit('https://staging.redpwn.net/', 'e0efc6e1-3b04-400a-9d72-a2e2ae02c1f6', ctf_start, lambda x : x)
```

# Documentation

[Click here to see the documentation](rctf/golf/util.py)

