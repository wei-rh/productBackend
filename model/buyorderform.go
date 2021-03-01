package model

type BuyOrderForm struct {
	SendAddress string		`form:"sendAddress" json:"sendAddress"`
	GetAddress string		`form:"getAddress" json:"getAddress"`
	Remarks string			`form:"remarks" json:"remarks"`
	Context string			`form:"context" json:"context"`
	Weight string			`form:"weight" json:"weight"`
	Time string				`form:"time" json:"time"`
	Baojia int64			`form:"baojia" json:"baojia"`
}
