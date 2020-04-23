package model

//Course Struct
type Course struct {
	CourseId int    `json:"cid" xml:"cid" form:"cid" query:"cid"`
	Name     string `json:"name" xml:"name" form:"name" query:"name"`
}
