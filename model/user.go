package model

import (
	"database/sql"
	"encoding/json"
)

//User Struct
type User struct {
	UserId   int        `json:"uid" xml:"uid" form:"uid" query:"uid"`
	Name     string     `json:"name" xml:"name" form:"name" query:"name"`
	Email    NullString `json:"email" xml:"email" form:"email" query:"email"`
	CourseId NullString `json:"cid" xml:"cid" form:"cid" query:"cid"`
}

type NullString struct {
	sql.NullString
}

// MarshalJSON for NullString
func (ns *NullString) MarshalJSON() ([]byte, error) {
	if ns.Valid {
		return json.Marshal(ns.String)
	}
	return []byte("null"), nil
}

// UnmarshalJSON for NullString
func (ns *NullString) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &ns.String)
	ns.Valid = (err == nil)
	return err
}
