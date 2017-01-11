import time
import sys
sys.path.append(r'haina.im/monitor/monitor_node/pythons/tools')
import tool

def mysqldProcess(hcmd, cmd):
	data = {}
	head = tool.getHead(hcmd)
	info = tool.getInfo(cmd)
	#if info == 'noexist':
	#	return 'noexist'
	#return info
	
	for i in range(len(head)):
		key = head[i]
		if key == 'USER' or key == '%CPU' or key == "%MEM" or key == 'PID' or key == 'RSS':
			data[key] = info[i]
	return (str(data))
	
	return ("{0},{1}".format(str(data),id(redis)))

name = 'ps -aux |grep -v grep| grep ' + sys.argv[1]
print(mysqldProcess('ps -aux |head -1', name))
#print(ss[0])
#print(ss[1])
#print(sys.argv[2])