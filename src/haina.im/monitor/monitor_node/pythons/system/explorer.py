#!usr/bin python3
# -*- coding:utf-8 -*-
import subprocess
import string

def get_some():
	try:
		u = subprocess.check_output('top -b -d 1 -n 1 | head -5', shell=True).decode().split('\n')
		return u
	except CalledProcessError:
		return 

def getTasks(info):
	return info

def getCpu(info):
	ss = info.split(':')
	return 'Cpu(%):' + ss[1]
	
def getMem(info):
	ss = info.split(':')
	return 'Mem(k):' + ss[1]

def getSwap(info):
	return 'Swap(k):' + ss[1]


info = get_some()

print(getTasks(info[1]))
print(getCpu(info[2]))
print(getMem(info[3]))
print(getSwap(info[4]))