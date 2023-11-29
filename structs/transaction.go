package structs

type TransactionRequest struct {
	NetworkMap  NetworkMap `json:"networkmap"`
	RuleResult  RuleResult `json:"ruleResult"`
	Transaction Pain002    `json:"transaction"`
}

type RuleResult struct {
	Id         string `json:"id"`
	Cfg        string `json:"cfg"`
	Result     bool   `json:"result"`
	SubRuleRef string `json:"subRuleRef"`
	Reason     string `json:"reason"`
	Desc       string `json:"desc"`
	PrcgTm     int32  `json:"prcgTm"`
}

type GrpHdr struct {
	MsgId   string `json:"MsgId"`
	CreDtTm string `json:"CreDtTm"`
}

type Amt struct {
	Amt int16  `json:"Amt"`
	Ccy string `json:"Ccy"`
}

type Agt struct {
	FinInstnId FinInstnId `json:"FinInstnId"`
}

type FinInstnId struct {
	ClrSysMmbId ClrSysMmbId `json:"ClrSysMmbId"`
}

type ClrSysMmbId struct {
	MmbId string `json:"MmbId"`
}

type ChrgsInf struct {
	Amt Amt `json:"Amt"`
	Agt Agt `json:"Agt"`
}

type InstgAgt struct {
	FinInstnId FinInstnId `json:"FinInstnId"`
}

type InstdAgt struct {
	FinInstnId FinInstnId `json:"FinInstnId"`
}

type TxInfAndSts struct {
	TxSts           string     `json:"TxSts"`
	OrgnlEndToEndId string     `json:"OrgnlEndToEndId"`
	OrgnlInstrId    string     `json:"OrgnlInstnumber"`
	ChrgsInf        []ChrgsInf `json:"ChrgsInf"`
	AccptncDtTm     string     `json:"AccptncDtTm"`
	InstgAgt        InstgAgt   `json:"InstgAgt"`
	InstdAgt        InstdAgt   `json:"InstdAgt"`
}

type FIToFIPmtSts struct {
	GrpHdr      GrpHdr      `json:"GrpHdr"`
	TxInfAndSts TxInfAndSts `json:"TxInfAndSts"`
}

type Pain002 struct {
	TxTp         string       `json:"TxTp"`
	FIToFIPmtSts FIToFIPmtSts `json:"FIToFIPmtSts"`
}
