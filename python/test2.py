#!/bin/env python
#-*- coding: utf-8 -*-

import subprocess
import threading

acount = threading.active_count()

print(acount)

output = subprocess.check_output("bash")
output2 = subprocess.check_output("sh")

print(output)
print(output2)
