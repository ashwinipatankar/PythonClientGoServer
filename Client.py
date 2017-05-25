#!/usr/bin/env python

import urllib2
import json
url = "http://127.0.0.1:8000"
dataRequest = urllib2.urlopen(url).read()
print dataRequest
j = json.loads(dataRequest)
print j
