#!/bin/env python
#-*- coding: utf-8 -*-

import threading

#
# return active thread count
#
ac_count = threading.active_count()
print(ac_count)

#
# return currnet thread object
#
cur_thread = threading.current_thread()
print(cur_thread)

#
# not supported python < 3.3
#
# id = threading.get_ident()
# print(id)

# mthread = threading.main_thread()
# print(mthread)

#
# return the thread stack size used when creating new threads
# 
stack_size = threading.stack_size()
print(stack_size)


