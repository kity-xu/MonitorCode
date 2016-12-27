import time
import sys
sys.path.append(r'monitor/local/pythons/tools')
import tool

def mysqldProcess(hcmd, cmd):
	data = {}
	head = tool.getHead(hcmd)
	info = tool.getInfo(cmd)
	
	for i in range(len(head)):
		data[head[i]] = info[i]
	return (str(data))
	
	#return ("{0},{1}".format(str(data),id(redis)))

print(mysqldProcess('ps -aux |head -1', 'ps -aux | grep mysqld'))