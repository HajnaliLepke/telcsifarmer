package main

type TelcsiFarmolas struct {
	// ID     uint
	MinPrice        uint     `form:"minprice" binding:"required"`
	MaxPrice        uint     `form:"maxprice" binding:"required"`
	ImportantPhones []string `form:"good_phones[]" binding:"-"`
	IsPhone         string   `form:"isphone" binding:"-"`
	NeutralPhones   []string `form:"neutral_phones[]" binding:"required"`
}
