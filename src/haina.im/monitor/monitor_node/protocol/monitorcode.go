package protocol

import (
	"encoding/json"
	"fmt"
	l4g "github.com/alecthomas/log4go"
	"haina.im/monitor/monitor_node/share"
	"haina.im/monitor/monitor_node/utils"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type MonitorCode struct {
}

/**	@funcName：	writeTube
*	@function:	错误返回
*	@parameter:	错误码
*	@return:	错误数据
**/
func writeTube(c ...interface{}) []byte {
	if len(c) == 3 {
		s := fmt.Sprintln("错误码：", c[0], "其他信息：", c[1], c[2])
		return []byte(s)
	}
	if len(c) == 2 {
		s := fmt.Sprintln("错误码：", c[0], "脚本：", c[1])
		return []byte(s)
	}
	if len(c) == 1 {
		s := fmt.Sprintln("错误码：", c[0])
		return []byte(s)
	}
	return nil
}

/**	@funcName：	shellCommand
*	@function:	启动单个脚本
*	@parameter:	脚本名
*	@return:	脚本输出数据
**/
func shellCommand(name, paras string, soc *SocketClient) string {
	//l4g.Info("~~~~~~~~~~~~~~~~~paras:%s", paras)
	cmd := exec.Command("python", share.PY_DIRPATH_APP+name, paras)
	if out, err := cmd.CombinedOutput(); err == nil {
		//l4g.Info("---------%s:", string(out)) //out return value of call *.py
		return string(out)
	} else {
		soc.Send(writeTube(share.PY_CALL_ERROR, share.PY_DIRPATH_APP+name))
		return ""
	}
}

/**	@funcName：	startScript
*	@function:	启动该应用脚本
*	@parameter:	record appName
*	@return:	none
**/
func (this *MonitorCode) startScript(rec utils.Record, soc *SocketClient) ResultRec {
	var r ResultRec
	info := shellCommand(rec.Provide, rec.Paras, soc)
	if strings.EqualFold(info, "") {
		soc.Send(writeTube(share.PY_RETURN_NONE, rec.Provide))
		//return
	}
	r.Name = rec.Name
	r.Data = info
	return r
}

/**	@funcName：	start
*	@function:	启动应用
*	@parameter:	appName
*	@return:	none
**/
func (this *MonitorCode) startApplication(app utils.Application, soc *SocketClient) ResultApp {
	var rs []ResultRec
	var a ResultApp
	//l4g.Info("**************** start %s application ****************", app.Name)
	for _, rec := range app.Record {
		//l4g.Info("--------------- start %s of %s ---------------", rec.Provide, app.Name)
		//执行该应用下的每个脚本
		resultRec := this.startScript(rec, soc)
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
func (this *MonitorCode) startExplorer(node *utils.Monitornode, soc *SocketClient) Explorer {
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
				css := strings.Split(strings.TrimLeft(strings.Split(strings.Split(ss[1], ":")[0], ",")[0], " "), " ")[0] //用户使用cpu
				cpu, _ := strconv.ParseFloat(css, 64)
				//l4g.Debug("-------------cpu:%v-------------max cpu:%v", cpu, float64(node.System.Node.Cpu)/100)
				if cpu/100 > float64(node.System.Node.Cpu)/100 {
					l4g.Debug("the cpu out range..........")
					soc.Send(writeTube(share.CPU_OUTOFRANGE))
				}
			}

			if strings.EqualFold(ss[0], "Mem") {
				exp.Mem = ss[1]

				total := strings.Split(strings.Split(ss[1], ":")[0], ",")[0]
				used := strings.Split(strings.Split(ss[1], ":")[0], ",")[2]

				to, _ := strconv.ParseFloat(strings.Split(strings.TrimLeft(total, " "), " ")[0], 64)
				us, _ := strconv.ParseFloat(strings.Split(strings.TrimLeft(used, " "), " ")[0], 64)

				//l4g.Debug("-----------to:%v---us:%v----", to, us)

				if us/to > float64(node.System.Node.Mem)/100 {
					l4g.Debug("the mem out range..........")
					soc.Send(writeTube(share.MEM_OUTOFRANGE))
				}
				//l4g.Info("---Mem:%v--%v", us/to, float64(node.System.Node.Mem)/100)
			}
			if strings.EqualFold(ss[0], "Swap") {
				exp.Swap = ss[1]
			}
		}
	} else {
		soc.Send(writeTube(share.START_EXP_ERROR))
	}
	return exp
}

/**	@funcName：	startSystem
*	@function:	启动系统应用
*	@parameter:	None
*	@return:	ResultSystem
**/
func (this *MonitorCode) startOsystem(soc *SocketClient) Osystem {
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
		soc.Send(writeTube(share.START_SYS_ERROR))
	}
	return sys
}

/**	@funcName:	Collection
*	@function：	数据采集
*	@parameter:	解析的配置文件结构Monitornode
*	@return :	MonitorData
**/
func (this *MonitorCode) Collection(node *utils.Monitornode, cc chan MonitorData, soc *SocketClient) {
	apps := utils.GetAppsByConfig(node)

	for {
		var res MonitorData
		var afs []ResultApp
		for _, app := range apps.Application {
			//开启单个应用
			af := this.startApplication(app, soc)
			afs = append(afs, af)
		}
		//获取系统资源
		res.Exp = this.startExplorer(node, soc)
		res.Osys = this.startOsystem(soc)
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
	soc := new(SocketClient)
	soc.WebsocketClient() //开启websocket连接

	node, err := utils.ParseXml(share.GO_CONFIG_FILE)
	if err != nil {
		l4g.Debug("错误码：%d", share.PY_PARSE_ERRROR)
		soc.Send(writeTube(share.PY_PARSE_ERRROR))
		return
	}
	scripts, err := utils.CheckPythonScripts(node) //apps 是可执行的脚本名或是不可执行的脚本名Provide
	if err != nil {                                //一般函数调用错误
		if scripts != nil { //意味着配置文件中有的本地不存在
			l4g.Debug("错误码：%s， 其他信息： %v：%v", share.PY_NOT_EXIST, err, scripts) //把本地不存在的脚本名返回
			soc.Send(writeTube(share.PY_NOT_EXIST, err, scripts))
			return
		}
		l4g.Debug("错误码：%d", share.PY_WALKDIR_ERROR)
		soc.Send(writeTube(share.PY_WALKDIR_ERROR))
		return
	}

	cc := make(chan MonitorData, 1)

	go this.Collection(node, cc, soc)

	for {
		//l4g.Info("********************DATA*******************%v", <-cc)
		context, err := json.Marshal(<-cc)
		if err != nil {
			l4g.Error(err)
		}

		soc.Send(context)
		time.Sleep(time.Duration(5) * time.Second)
		l4g.Info("recevied data is:%s", soc.Received)
	}
}
