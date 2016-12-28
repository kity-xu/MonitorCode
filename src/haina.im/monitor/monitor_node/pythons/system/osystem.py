#!usr/bin python3
# coding: utf-8
import subprocess

def get_osystem(self):
	try:
		o = subprocess.check_output('uname -o', shell=True).decode()
		return o
	except CalledProcessError:
		return ''

def get_user(self):
	try:
		u = subprocess.check_output('uname -n', shell=True).decode()
		return u
	except CalledProcessError:
		return -1

def get_ip(self):
	try:
		u = subprocess.check_output("ifconfig ens33 | grep 'inet addr' | awk '{print $2}' | awk -F: '{print $2}'", shell=True).decode()
		return u
	except CalledProcessError:
		return -1

def get_version(self):
	try:
		u = subprocess.check_output('uname -v', shell=True).decode()
		return u
	except CalledProcessError:
		return -1

def get_platform(self):
	try:
		u = subprocess.check_output('uname -i', shell=True).decode()
		return u
	except CalledProcessError:
		return -1

def get_all(self):
	try:
		u = subprocess.check_output('uname -a', shell=True).decode()
		return u
	except CalledProcessError:
		return -1