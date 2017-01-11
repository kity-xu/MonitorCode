#!usr/bin python3
# -*- coding:utf-8 -*-
import subprocess

def get_osystem():
	try:
		o = subprocess.check_output('uname -o', shell=True).decode()
		return 'Sys::' + o
	except CalledProcessError:
		return ''

def get_user():
	try:
		u = subprocess.check_output('uname -n', shell=True).decode()
		return 'User::' + u
	except CalledProcessError:
		return -1

def get_ip():
	try:
		u = subprocess.check_output("ifconfig ens33 | grep 'inet addr' | awk '{print $2}' | awk -F: '{print $2}'", shell=True).decode()
		return 'IP::' + u
	except CalledProcessError:
		return -1

def get_version():
	try:
		u = subprocess.check_output('uname -v', shell=True).decode()
		return 'Version::' + u
	except CalledProcessError:
		return -1

def get_platform():
	try:
		u = subprocess.check_output('uname -i', shell=True).decode()
		return 'Platform::' + u
	except CalledProcessError:
		return -1

def get_all():
	try:
		u = subprocess.check_output('uname -a', shell=True).decode()
		return 'All::' + u
	except CalledProcessError:
		return -1

print(get_osystem()  + get_ip())