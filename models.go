package main

type TelcsiFarmolas struct {
	// ID     uint
	MinPrice uint `form:"minprice"`
	MaxPrice uint `form:"maxprice"`
	Phone1 string `form:"first"`
	Phone1Important bool `form:"first_green"`
	Phone2 string `form:"second"`
	Phone2Important bool `form:"second_green"`
	Phone3 string `form:"third"`
	Phone3Important bool `form:"third_green"`
	Phone4 string `form:"fourth"`
	Phone4Important bool `form:"fourth_green"`
	Phone5 string `form:"fifth"`
	Phone5Important bool `form:"fifth_green"`
	Phone6 string `form:"sixth"`
	Phone6Important bool `form:"sixth_green"`
	Phone7 string `form:"seventh"`
	Phone7Important bool `form:"seventh_green"`
	Phone8 string `form:"eighth"`
	Phone8Important bool `form:"eighth_green"`
	Phone9 string `form:"ninth"`
	Phone9Important bool `form:"ninth_green"`
	Phone10 string `form:"tenth"`
	Phone10Important bool `form:"tenth_green"`
}
