package structs

type NetworkMap struct {
	Active   bool      `json:"active"`
	Cfg      string    `json:"cfg"`
	Messages []Message `json:"messages"`
}
