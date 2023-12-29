package models

const RetryConnectSeconds = 20
const ConstCountFailedConnections = 12000 / RetryConnectSeconds // timeout seconds / retry connections

type ConfService1CAPI struct {
	MinPort int `json:"MinPort"`
	MaxPort int `json:"MaxPort"`
	Timeout int `json:"Timeout"`
}

type ConfSQL struct {
	TypeSQL  string `json:"typeSQL"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Login    string `json:"login"`
	Password string `json:"password"`
	DBname   string `json:"dbname"`
}

type ConfApp struct {
	Port int `json:"port"`
}
