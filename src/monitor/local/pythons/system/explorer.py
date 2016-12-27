#!usr/bin python3
# -*- coding:utf-8 -*-
import subprocess
import string

def get_some():
	try:
		u = subprocess.check_output('top -b -d 1 -n 1 | head -5', shell=True).decode()
		return u
	except CalledProcessError:
		return 

def operation(strs):
	print(strs.split('\n')[2])

info = get_some()
operation(info)