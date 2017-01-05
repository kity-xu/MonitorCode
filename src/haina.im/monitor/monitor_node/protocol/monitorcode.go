package protocol

import (
	//"errors"
	l4g "github.com/alecthomas/log4go"
	"haina.im/monitor/monitor_node/share"
	"haina.im/monitor/monitor_node/utils"
	"os/exec"
	//"strconv"
	"strings"
	"time"
)

type MonitorCode struct {
}

/**	@funcName：	shellCommand
*	@function:	启动单个脚本
*	@parameter:	脚本名
*	@return:	脚本输出数据
**/
func shellCommand(name, paras string) string {
	//l4g.Info("~~~~~~~~~~~~~~~~~paras:%s", paras)
	cmd := exec.Command("python", share.PY_DIRPATH_APP+name, paras)
	if out, err := cmd.CombinedOutput(); err == nil {
		//l4g.Info("---------%s:", string(out)) //out return value of call *.py
		return string(out)
	} else {
		l4g.Debug("错误码:%d, 脚本：%s", share.PY_CALL_ERROR, share.PY_DIRPATH_APP+name)
		return ""
	}
}

/**	@funcName：	startScript
*	@function:	启动该应用脚本
*	@parameter:	record appName
*	@return:	none
**/
func (this *MonitorCode) startScript(rec utils.Record) ResultRec {
	var r ResultRec
	//t, _ := strconv.Atoi(rec.Timespan)
	//for {
	info := shellCommand(rec.Provide, rec.Paras)
	if strings.EqualFold(info, "") {
		l4g.Debug("错误码：%d, 脚本：%s", share.PY_RETURN_NONE, rec.Provide)
		//return
	}
	//l4g.Debug("最终的数据:%s", info) //最终的数据
	//utils.Writefile("haina.im/"+rec.Name+".txt", info)

	//time.Sleep(time.Duration(t) * time.Second)
	//}
	r.Name = rec.Name
	r.Data = info
	return r
}

/**	@funcName：	start
*	@function:	启动应用
*	@parameter:	appName
*	@return:	none
**/
func (this *MonitorCode) startApplication(app utils.Application) ResultApp {
	var rs []ResultRec
	var a ResultApp
	//l4g.Info("**************** start %s application ****************", app.Name)
	for _, rec := range app.Record {
		//l4g.Info("--------------- start %s of %s ---------------", rec.Provide, app.Name)
		//执行该应用下的每个脚本
		//this.startScript(rec)
		resultRec := this.startScript(rec)
		rs = append(rs, resultRec)
	}
	a.Name = app.Name
	a.Recs = rs
	return a
}

/**	@funcName：	startSystem
*	@function:	启动系统应用
*	@parameter:	None
*	@return:	ResultSystem
**/
func (this *MonitorCode) startExplorer() Explorer {
	var exp Explorer
	cmd := exec.Command("python", share.PY_DIRPATH_SYS+"explorer.py")
	if out, err := cmd.CombinedOutput(); err == nil {
		//l4g.Info("---------%s:", string(out)) //out return value of call *.py
		for _, value := range strings.Split(string(out), "\n") {
			//l4g.Info("---------values:%s", value)
			ss := strings.Split(value, "::")
			if strings.EqualFold(ss[0], "Tasks") {
				exp.Tasks = ss[1]
			}
			if strings.EqualFold(ss[0], "Cpu") {
				exp.Cpu = ss[1]
			}
			if strings.EqualFold(ss[0], "Mem") {
				exp.Mem = ss[1]
			}
			if strings.EqualFold(ss[0], "Swap") {
				exp.Swap = ss[1]
			}
		}
	} else {
		l4g.Debug("错误码：%d", share.START_EXP_ERROR)
	}
	return exp
}

/**	@funcName：	startSystem
*	@function:	启动系统应用
*	@parameter:	None
*	@return:	ResultSystem
**/
func (this *MonitorCode) startOsystem() Osystem {
	var sys Osystem
	cmd := exec.Command("python", share.PY_DIRPATH_SYS+"osystem.py")
	if out, err := cmd.CombinedOutput(); err == nil {
		//l4g.Info("---------%s:", string(out)) //out return value of call *.py
		for _, value := range strings.Split(string(out), "\n") {
			//l4g.Info("---------values:%s", value)
			ss := strings.Split(value, "::")
			if strings.EqualFold(ss[0], "Osystem") {
				sys.Sys = ss[1]
			}
			if strings.EqualFold(ss[0], "User") {
				sys.User = ss[1]
			}
			if strings.EqualFold(ss[0], "IP") {
				sys.IP = ss[1]
			}
			if strings.EqualFold(ss[0], "Version") {
				sys.Version = ss[1]
			}
			if strings.EqualFold(ss[0], "Platform") {
				sys.Platform = ss[1]
			}
		}
	} else {
		l4g.Debug("错误码：%d", share.START_SYS_ERROR)
	}
	return sys
}

/**	@funcName:	Collection
*	@function：	数据采集
*	@parameter:	解析的配置文件结构Monitornode
*	@return :	MonitorData
**/
func (this *MonitorCode) Collection(node *utils.Monitornode, cc chan MonitorData) {
	apps := utils.GetAppsByConfig(node)

	for {
		var res MonitorData
		var afs []ResultApp
		for _, app := range apps.Application {
			//开启单个应用
			af := this.startApplication(app)
			afs = append(afs, af)
		}
		//获取系统资源
		res.Exp = this.startExplorer()
		res.Osys = this.startOsystem()
		res.Apps = afs
		res.Time = time.Now().Format("2006-01-02 15:04:05")
		cc <- res

		time.Sleep(time.Duration(apps.Timespan) * time.Second)
	}
}

/**	@funcName：	StartMonitor
*	@function:	启动监听服务
*	@parameter:	none
*	@return:	none
**/
func (this *MonitorCode) StartMonitor() {
	node, err := utils.ParseXml(share.GO_CONFIG_FILE)
	if err != nil {
		l4g.Debug("错误码：%d", share.PY_PARSE_ERRROR)
		return
	}
	scripts, err := utils.CheckPythonScripts(node) //apps 是可执行的脚本名或是不可执行的脚本名Provide
	if err != nil {                                //一般函数调用错误
		if scripts != nil { //意味着配置文件中有的本地不存在
			l4g.Debug("错误码：%s， 其他信息： %v：%v", share.PY_NOT_EXIST, err, scripts) //把本地不存在的脚本名返回
			return
		}
		l4g.Debug("错误码：%d", share.PY_WALKDIR_ERROR)
		return
	}

	cc := make(chan MonitorData, 1)

	go this.Collection(node, cc)

	//for {
	//l4g.Info("********************DATA*******************%v", <-cc)
	WebsocketClient(<-cc)
	//time.Sleep(time.Duration(5) * time.Second)
	//l4g.Info("----jsons:%v", jsons)
	//}
}
