#/usr/bin/python
#--* utf-8 --*
import subprocess

def getHead(cmd):
	try:
		output = subprocess.check_output(cmd, shell=True).decode().split('\n')[0]
		ss = ','.join(filter(lambda x: x, output.split(' '))).split(',')
		return ss
	except CalledProcessError:
		return "error"

def getInfo(cmd):
	try:
		output = subprocess.check_output(cmd, shell=True).decode().split('\n')[0]
		ss = ','.join(filter(lambda x: x, output.split(' '))).split(',')
		return ss
	except CalledProcessError:
		return "error"