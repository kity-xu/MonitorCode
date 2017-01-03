package protocol

type ResultRec struct {
	Name string
	Data string
}

type ResultApp struct {
	Name string
	Recs []ResultRec
}

type Explorer struct {
	Tasks string
	Cpu   string
	Mem   string
	Swap  string
}

type Osystem struct {
	Sys      string
	User     string
	IP       string
	Version  string
	Platform string
}

type MonitorData struct {
	Apps []ResultApp
	Exp  Explorer
	Osys Osystem
	Time string
}
