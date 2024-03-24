package models

type Infobases struct {
	Bases []Infobase `json:"bases"`
}

type Infobase struct {
	Name     string `json:"name"`
	URL      string `json:"URL"`
	Login    string `json:"login"`
	Password string `json:"password"`
}
