package protocol

import (
	"errors"
	l4g "github.com/alecthomas/log4go"
	"haina.im/monitor/monitor_node/share"
	"haina.im/monitor/monitor_node/utils"
	"os/exec"
	"strings"
	"time"
)

type MonitorCode struct {
}

const (
	py_dirpath_app = share.PY_DIRPATH_APP
	config_file    = share.GO_CONFIG_FILE
)

func shellCommand(name string) string {
	cmd := exec.Command("python", py_dirpath_app+name)
	if out, err := cmd.CombinedOutput(); err == nil {
		l4g.Info("---------%s:", string(out)) //out return value of call *.py
		return string(out)
	} else {
		l4g.Debug("running shellCommand error ....")
		return ""
	}
}

func (this *MonitorCode) start(name string) {
	l4g.Info("----------start %s server..", name)
	for {
		info := shellCommand(name)
		if strings.EqualFold(info, "") {
			l4g.Debug("shellCommand return value is null...")
			return
		}
		l4g.Debug("最终的数据:%s", info) //最终的数据
		time.Sleep(1000)
	}
}

func checkPythonScripts() ([]string, error) {
	ee := errors.New("Unable to identify the script file")
	/**
	*	ds 与 cs 对比，以配置文件cs为主。如果本地不存在返回配置错误应用，程序结束
	*
	**/
	ds, err := utils.GetAppsBywalkdir(py_dirpath_app, "pY") //遍历python/application 目录，查看已有的脚本
	if err != nil {
		return nil, err
	}
	//l4g.Debug(ds)

	cs, err := utils.GetAppsByConfigfile(config_file) //读取配置文件查看已有的脚本
	if err != nil {
		return nil, err
	}
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
			return noexit, ee
		}
	}
	l4g.Info("apps :", apps)
	return apps, nil
}

func (this *MonitorCode) StartMonitor() {
	apps, err := checkPythonScripts()
	if err != nil { //一般函数调用错误
		if apps != nil { //意味着配置文件中有的本地不存在
			l4g.Debug("%v：%v", err, apps) //把本地不存在的脚本名返回
		}
		//return
	}

	// i := 0
	// for ; i<len(apps); i++ {
	// 	go start(apps[i])
	// }
	this.start(apps[0])
}
