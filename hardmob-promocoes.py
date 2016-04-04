#!/usr/bin/python3

from urllib.request import urlopen
import re
import subprocess
p = re.compile(r'stream-item-tweet.*</li>', re.IGNORECASE | re.DOTALL)
response = urlopen('https://twitter.com/hardmob_promo')
html = response.read().decode()
result = p.findall(html)
if (len(result) > 0):
    p = re.compile(r'TweetTextSize.*?>(.*)<a.*</p>', re.IGNORECASE)
    result = p.search(result[0])
    print(result.group(1))
    subprocess.call(["notify-send","--expire-time=30000", "HardMob", result.group(1)])
