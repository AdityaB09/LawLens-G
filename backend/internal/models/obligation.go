package models

import "time"

type Obligation struct {
	ID         uint       `gorm:"primaryKey" json:"id"`
	ContractID uint       `json:"contractId"`
	ClauseID   uint       `json:"clauseId"`
	Description string    `json:"description"`
	DueDate    *time.Time `json:"dueDate"`
}

func (Obligation) TableName() string {
	return "obligations"
}
