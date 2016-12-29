package protocol

import (
	"errors"
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

const (
	py_dirpath_app = share.PY_DIRPATH_APP
)

/**	@funcName：	shellCommand
*	@function:	启动单个脚本
*	@parameter:	脚本名
*	@return:	脚本输出数据
**/
func shellCommand(name, paras string) string {
	//l4g.Info("~~~~~~~~~~~~~~~~~paras:%s", paras)
	cmd := exec.Command("python", py_dirpath_app+name, paras)
	if out, err := cmd.CombinedOutput(); err == nil {
		//l4g.Info("---------%s:", string(out)) //out return value of call *.py
		return string(out)
	} else {
		l4g.Debug("running shellCommand error ....")
		return ""
	}
}

/**	@funcName：	startScript
*	@function:	启动该应用脚本
*	@parameter:	record appName
*	@return:	none
**/
func (this *MonitorCode) startScript(rec utils.Record) {
	t, _ := strconv.Atoi(rec.Timespan)
	for {
		info := shellCommand(rec.Provide, rec.Paras)
		if strings.EqualFold(info, "") {
			l4g.Debug("shellCommand return value is null...")
			return
		}
		l4g.Debug("最终的数据:%s", info) //最终的数据
		time.Sleep(time.Duration(t) * time.Second)
	}
}

/**	@funcName：	start
*	@function:	启动应用
*	@parameter:	appName
*	@return:	none
**/
func (this *MonitorCode) startApplication(app utils.Application) {
	l4g.Info("**************** start %s application ****************", app.Name)
	for _, rec := range app.Record {
		l4g.Info("--------------- start %s of %s ---------------", rec.Provide, app.Name)
		//执行该应用下的每个脚本
		this.startScript(rec)
		//go this.startScript(rec, app.Name)
	}
}

/**	@funcName:	checkPythonScripts
*	@function：	核实配置脚本
*	@parameter:	解析的配置文件结构Monitornode
*	@return :	py脚本名，nil（或是出错脚本，ee）
**/
func checkPythonScripts(node *utils.Monitornode) ([]string, error) {
	ee := errors.New("Unable to identify the script file")
	ds, err := utils.GetAppsBywalkdir(py_dirpath_app, "py") //遍历python/application 目录，查看已有的脚本
	if err != nil {
		return nil, err
	}
	//l4g.Debug(ds)

	cs := utils.GetPythonsByConfig(node) //读取配置文件查看已有的脚本
	//l4g.Debug(cs)

	var apps []string
	var noexit []string

	var iTag bool
	for _, v := range cs {
		iTag = false
		for _, v2 := range ds {
			if strings.EqualFold(v, v2) {
				apps = append(apps, v)
				iTag = true
				break
			} else {
				continue
			}
		}

		if !iTag { //意味着配置文件中有的本地不存在
			noexit = append(noexit, v)
		}
	}
	if !iTag {
		return noexit, ee
	}
	l4g.Info("apps :", apps)
	return apps, nil
}

/**	@funcName：	StartMonitor
*	@function:	启动监听服务
*	@parameter:	none
*	@return:	none
**/
func (this *MonitorCode) StartMonitor() {
	node, err := utils.ParseXml(share.GO_CONFIG_FILE)
	if err != nil {
		return //parse config.xml error
	}
	scripts, err := checkPythonScripts(node) //apps 是可执行的脚本名或是不可执行的脚本名Provide
	if err != nil {                          //一般函数调用错误
		if scripts != nil { //意味着配置文件中有的本地不存在
			l4g.Debug("%v：%v", err, scripts) //把本地不存在的脚本名返回
		}
		//return
	}

	apps := utils.GetAppsByConfig(node)
	if apps == nil {
		return
	}

	for _, app := range apps {
		//开启单个应用
		this.startApplication(app)
	}

	//record := utils.GetAppRecordsInfo(node, "mysqld") //一个application的所有record，第二个参数为应用名
	//l4g.Debug("mysqld fist name:%s", record[1].Provide)

	// i := 0
	// for ; i<len(scriptsscripts); i++ {
	// 	go start(scripts[i])
	// }

}
