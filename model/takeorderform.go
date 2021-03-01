package model

type TakeOrderForm struct {
	SendAddress string		`form:"sendAddress" json:"sendAddress"`
	GetAddress string		`form:"getAddress" json:"getAddress"`
	Remarks string			`form:"remarks" json:"remarks"`
	Goods string			`form:"goods" json:"goods"`
	Weight string			`form:"weight" json:"weight"`
	Time string				`form:"time" json:"time"`
	Baojia int64			`form:"baojia" json:"baojia"`
}
