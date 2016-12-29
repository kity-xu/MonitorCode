package utils

import (
	"encoding/xml"
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

const (
	path_appconfig = share.GO_CONFIGS_APP
)

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
func GetAppsByConfig(node *Monitornode) []Application {
	var ss []Application
	var n int
	var apps Application
	for n, apps = range node.Applications.Application {
		ss = append(ss, apps)
	}

	l4g.Info("application total: %d,  that respectively are: %s\n", n+1, ss)
	return ss
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
