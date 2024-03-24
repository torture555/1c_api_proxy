package models

const RetryConnectSeconds = 3
const ConstCountFailedConnections = 1200 / RetryConnectSeconds // timeout seconds / retry connections

type ConfSQL struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Login    string `json:"login"`
	Password string `json:"password"`
	DBName   string `json:"DBName"`
}
