package protocol

type Record struct {
	Name string
	Data string
}

type AppStatus struct {
	Name    string
	Num     int
	Records []Record
}

type SysStatus struct {
	Tasks string
	Cpu   string
	Mem   string
	Swap  string
}

type SysInfo struct {
	IP   string
	ID   string
	Name string
}

// Code 错误码； Msg 错误信息； Grade 错误级别
type Alarm struct {
	IsErr bool
	Code  int
	Msg   string
	Grade int
}

type MonitorData struct {
	Info   SysInfo
	Status SysStatus
	Num    int
	Apps   []AppStatus
	Arms   []Alarm
	Time   string
}
