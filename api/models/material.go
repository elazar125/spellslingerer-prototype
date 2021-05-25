package models

import "api/db"

// Material is used to craft cards
type Material struct {
	ID   int64
	Name string
}

// AllMaterials enumerates all materials
func AllMaterials() ([]Material, error) {
	var allMaterials []Material
	err := db.GlobalSqlxDB.Select(&allMaterials, "SELECT * FROM public.materials")
	return allMaterials, err
}
