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

type MonitorData struct {
	Apps   []AppStatus
	Num    int
	Status SysStatus
	Info   SysInfo
	Time   string
}
