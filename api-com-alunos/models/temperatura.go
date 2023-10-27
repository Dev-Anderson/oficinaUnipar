package models

import "time"

type Temperatura struct {
	ID          uint      `gorm:"primaryKey; autoIncrement:true" `
	Temperatura string    `json:"temperatura"`
	Outro       string    `json:"outro"`
	Data        time.Time `json:"data"`
}
