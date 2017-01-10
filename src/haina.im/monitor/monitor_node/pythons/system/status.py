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
	ss = info.split(':')
	return 'Tasks::' + ss[1]

def getCpu(info):
	ss = info.split(':')
	return 'Cpu::' + ss[1]
	
def getMem(info):
	ss = info.split(':')
	return 'Mem::' + ss[1]

def getSwap(info):
	ss = info.split(':')
	return 'Swap::' + ss[1]


info = get_some()

#print(info)
print(getTasks(info[1]))
print(getCpu(info[2]))
print(getMem(info[3]))
print(getSwap(info[4]))