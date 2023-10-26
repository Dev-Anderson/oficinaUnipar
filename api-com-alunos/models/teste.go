package models

type Teste struct {
	ID    uint `gorm: "primaryKey; autoIncrement:treu"`
	Texto string
}
