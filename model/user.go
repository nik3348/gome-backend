package model

//User Struct
type User struct {
	UserId string `json:"uid" xml:"uid" form:"uid" query:"uid"`
	Name   string `json:"name" xml:"name" form:"name" query:"name"`
	Email  string `json:"email" xml:"email" form:"email" query:"email"`
}
