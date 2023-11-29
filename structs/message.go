package structs

type Message struct {
	Id       string    `json:"id"`
	Host     string    `json:"host"`
	Cfg      string    `json:"cfg"`
	TxTp     string    `json:"txTp"`
	Channels []Channel `json:"channels"`
}
