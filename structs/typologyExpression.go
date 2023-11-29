package structs

type TypologyExpression struct {
	TypologyName string       `json:"typology_name"`
	Id           string       `json:"id"`
	Cfg          string       `json:"cfg"`
	Rules        []RuleConfig `json:"rules"`
	Expression   Expression   `json:"expression"`
}

type RuleConfig struct {
	Id    string `json:"id"`
	Cfg   string `json:"cfg"`
	Ref   string `json:"ref"`
	True  string `json:"true"`
	False string `json:"false"`
}

type Expression struct {
	Operator string `json:"operator"`
	Terms    []Term `json:"terms"`
}

type Term struct {
	Id  string `json:"id"`
	Cfg string `json:"cfg"`
}
