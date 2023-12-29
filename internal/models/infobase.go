package models

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

type ServiceInfobase1C struct {
	Base        *Infobase
	PortService int
	Status      bool
}
