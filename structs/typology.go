package structs

type Typology struct {
	Id    string `json:"id"`
	Host  string `json:"host"`
	Cfg   string `json:"cfg"`
	Desc  string `json:"desc"`
	Rules []Rule `json:"rules"`
}
