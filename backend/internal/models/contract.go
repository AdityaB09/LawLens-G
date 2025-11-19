package models

import "time"

type Contract struct {
	ID               uint      `gorm:"primaryKey" json:"id"`
	Title            string    `json:"title"`
	PartyA           string    `json:"partyA"`
	PartyB           string    `json:"partyB"`
	UploadedAt       time.Time `json:"uploadedAt"`
	OverallRiskLevel string    `json:"overallRiskLevel"`
	OverallRiskScore float64   `json:"overallRiskScore"`

	Clauses     []Clause     `json:"clauses"`
	Obligations []Obligation `json:"obligations"`
}

func (Contract) TableName() string {
	return "contracts"
}
