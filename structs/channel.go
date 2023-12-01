package structs

type Channel struct {
	Id         string     `json:"id"`
	Host       string     `json:"host"`
	Cfg        string     `json:"cfg"`
	Typologies []Typology `json:"typologies"`
}
