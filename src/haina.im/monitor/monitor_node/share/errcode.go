package share

const (
	PY_CALL_ERROR    = 10001 // 调用脚本错误
	PY_RETURN_NONE   = 10002 // 脚本返回为空
	PY_NOT_EXIST     = 10003 // 脚本不存在
	PY_PARSE_ERRROR  = 10004 // 解析脚本错误
	PY_WALKDIR_ERROR = 10005 //遍历脚本目录错误

	START_SYS_ERROR = 20001 //启动Osystem错误
	START_EXP_ERROR = 20002 //启动Explorer错误

	//websocket disconnected
	WEBSOCKET_DISCONNECTED = 30001 // 连接断开

	//explorer
	CPU_OUTOFRANGE = 40001 //cpu使用超出设定值
	MEM_OUTOFRANGE = 40002 //内存使用超出设定值
)
