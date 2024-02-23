package models

const RetryConnectSeconds = 3
const ConstCountFailedConnections = 1200 / RetryConnectSeconds // timeout seconds / retry connections

type ConfSQL struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type ConfApp struct {
	Port int `json:"port"`
}

type Infobases struct {
	Bases []Infobase `json:"bases"`
}

type Infobase struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	URL      string `json:"URL"`
	Login    string `json:"login"`
	Password string `json:"password"`
}
