package models

type ClauseTrigger struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	ClauseID  uint   `json:"clauseId"`
	Phrase    string `json:"triggerPhrase"`
	Type      string `json:"triggerType"`
	Severity  string `json:"severity"`
}

func (ClauseTrigger) TableName() string {
	return "clause_triggers"
}
