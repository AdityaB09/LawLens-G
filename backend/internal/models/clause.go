package models

type Clause struct {
	ID         uint            `gorm:"primaryKey" json:"id"`
	ContractID uint            `json:"contractId"`
	OrderIndex int             `json:"orderIndex"`
	Heading    string          `json:"heading"`
	Text       string          `json:"text"`
	ClauseType string          `json:"clauseType"`
	RiskLevel  string          `json:"riskLevel"`
	RiskScore  float64         `json:"riskScore"`
	Triggers   []ClauseTrigger `json:"triggers"`
	Status     string          `gorm:"-" json:"status,omitempty"`     // used in diff
	RiskDelta  float64         `gorm:"-" json:"riskDelta,omitempty"`  // used in diff
}

func (Clause) TableName() string {
	return "clauses"
}
