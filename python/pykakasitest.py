#!/usr/bin/env python
# -*- coding: utf-8 -*-


from pykakasi import kakasi

kakasi = kakasi()
kakasi.setMode('H', 'a')
kakasi.setMode('K', 'a')
kakasi.setMode('J', 'a')
conv = kakasi.getConverter()
print(conv.do(u'本日は晴天なり'))
