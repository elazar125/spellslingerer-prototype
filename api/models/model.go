package models

// Model represents a type that has basic get/create/update/delete functions
type Model interface {
	All() ([]Model, error)
	Lookup(string) (Model, error)
	Create() error
	Update() error
	Upsert() error
	Delete() error
}
