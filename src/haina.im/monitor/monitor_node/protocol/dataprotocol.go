package protocol

type ResultRec struct {
	Name string
	Data string
}

type ResultApp struct {
	Name string
	Recl int
	Recs []ResultRec
}

type SysStatus struct {
	Tasks string
	Cpu   string
	Mem   string
	Swap  string
}

type Osystem struct {
	Sys          string
	IP           string
	Id           string
	NodeDescribe string
	//User     string
	//Version  string
	//Platform string
}

type MonitorData struct {
	Apps  []ResultApp
	Appl  int
	Statu SysStatus
	Osys  Osystem
	Time  string
}
