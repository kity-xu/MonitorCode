package utils

import (
	"encoding/xml"
	"errors"
	l4g "github.com/alecthomas/log4go"
	"haina.im/monitor/monitor_node/share"
	"io/ioutil"
	"strings"
)

/**
	enconding/xml:
	func Marshal(v interface{}) ([]byte, error)
 	func Unmarshal(data []byte, v interface{}) error
 	上面两个函数的文档（注释）很长。从文档知道，Marshal是将v代表的数据转为XML格式（生成XML）；而Unmarshal刚好相反，是解析XML，同时将解析的结果保存在v中。

 	func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error)
	 这个函数和Marshal的不同是，每个XML元素会增加前缀和缩进

**/

type Monitornode struct {
	Applications Applications `xml:"Applications"`
	System       System       `xml:"System"`
}
type System struct {
	Node []Node `xml:"Node"`
}
type Node struct {
	Name     string `xml:"Name,attr"`
	Timespan string `xml:"Timespan,attr"`
	Cpu      string `xml:"Cpu"`
	Task     string `xml:"Task"`
	Mem      string `xml:"Mem"`
}

type Applications struct {
	Timespan    int           `xml:"Timespan,attr"`
	Application []Application `xml:"Application"`
}

type Application struct {
	Record []Record `xml:"Record"`
	Name   string   `xml:"Name,attr"`
}
type Record struct {
	//Command		[]Command	`xml:"Command"`
	Name     string `xml:"Name,attr"`
	Provide  string `xml:"Provide,attr"`
	Timespan string `xml:"Timespan,attr"`
	Paras    string `xml:"Paras"`
}

type RecordInfo struct {
	Name     string
	Provide  string
	Timespan string
	Paras    string
}

func ParseXml(path string) (*Monitornode, error) {
	var result *Monitornode
	content, err := ioutil.ReadFile(path)
	if err != nil {
		l4g.Error(err)
		return result, err
	}

	e := xml.Unmarshal(content, &result)
	if e != nil {
		l4g.Error(e)
		return result, err
	}

	l4g.Info(result)
	return result, nil
}

/**	@funcName：	GetAppsByConfig
*	@function:	获取配置文件所有应用(Application)结构
*	@parameter:	Monitornode
*	@return:	应用结构的集合[]Application
**/
func GetAppsByConfig(node *Monitornode) Applications {
	var res Applications
	var ss []Application
	for _, app := range node.Applications.Application {
		ss = append(ss, app)
	}
	res.Application = ss
	//l4g.Info("application total: %d,  that respectively are: %v\n", n+1, res)
	return res
}

/**	@funcName：	GetPythonsByConfig
*	@function:	获取配置文件所有python脚本名
*	@parameter:	Monitornode
*	@return:	脚本的集合
**/
func GetPythonsByConfig(node *Monitornode) []string {
	var ss []string
	var apps Application
	var record Record
	for _, apps = range node.Applications.Application {
		//ss = append(ss, apps.Record.Provide)
		for _, record = range apps.Record {
			ss = append(ss, record.Provide)
		}
	}
	return ss
}

/**	@funcName：	GetRecordByName
*	@function:	根据Record name获取对应信息
*	@parameter:	1：Application		2：recName
*	@return:	RecordInfo
**/
func GetRecordByName(node *Application, recname string) RecordInfo {
	var info RecordInfo
	var rec Record
	for _, rec = range node.Record {
		if strings.EqualFold(rec.Name, recname) {
			info.Name = rec.Name
			info.Provide = rec.Provide
			info.Timespan = rec.Timespan
			info.Paras = rec.Paras
		}
	}
	return info
}

/**	@funcName：	GetRecordByProvide
*	@function:	根据provideu获取对应的record
*	@parameter:	1：Monitornode		2：provide
*	@return:	RecordInfo
**/
func GetRecordByProvide(node *Application, provide string) RecordInfo {
	var info RecordInfo
	var rec Record
	for _, rec = range node.Record {
		if strings.EqualFold(rec.Provide, provide) {
			info.Name = rec.Name
			info.Provide = rec.Provide
			info.Timespan = rec.Timespan
			info.Paras = rec.Paras
		}
	}
	return info
}

/**	@funcName:	checkPythonScripts
*	@function：	核实配置脚本
*	@parameter:	解析的配置文件结构Monitornode
*	@return :	py脚本名，nil（或是出错脚本，ee）
**/
func CheckPythonScripts(node *Monitornode) ([]string, error) {
	ee := errors.New("Unable to identify the script file")
	ds, err := GetAppsBywalkdir(share.PY_DIRPATH_APP, "py") //遍历python/application 目录，查看已有的脚本
	if err != nil {
		return nil, err
	}
	//l4g.Debug(ds)

	cs := GetPythonsByConfig(node) //读取配置文件查看已有的脚本
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
