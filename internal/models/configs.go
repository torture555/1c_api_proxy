package models

type Service1CAPI struct {
	MinPort int `json:"MinPort"`
	MaxPort int `json:"MaxPort"`
	Timeout int `json:"Timeout"`
}
