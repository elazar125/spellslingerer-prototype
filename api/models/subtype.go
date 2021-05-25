package models

import "api/db"

// SubType is a short description of what kind of creature is on the card
type SubType struct {
	ID   int64
	Name string
}

// AllSubTypes enumerates all subtypes
func AllSubTypes() ([]SubType, error) {
	var allSubTypes []SubType
	err := db.GlobalSqlxDB.Select(&allSubTypes, "SELECT * FROM public.sub_types")
	return allSubTypes, err
}
