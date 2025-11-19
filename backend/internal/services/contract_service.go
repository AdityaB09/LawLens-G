package services

import (
	"time"

	"lawlens-g/internal/models"

	"gorm.io/gorm"
)

func CreateContractWithAnalysis(db *gorm.DB, title, partyA, partyB, text string) (*models.Contract, error) {
	clauses := SegmentClauses(text)

	// Evaluate each clause
	for i := range clauses {
		cc := ClassifyClause(clauses[i].Text)
		risk := EvaluateRisk(clauses[i].Text, cc)
		clauses[i].ClauseType = cc.Type
		clauses[i].RiskLevel = risk.Level
		clauses[i].RiskScore = risk.Score
		clauses[i].Triggers = risk.Triggers
	}

	// Compute overall risk
	var avgScore float64
	for _, cl := range clauses {
		avgScore += cl.RiskScore
	}
	if len(clauses) > 0 {
		avgScore /= float64(len(clauses))
	}

	overallLevel := "LOW"
	switch {
	case avgScore >= 0.7:
		overallLevel = "HIGH"
	case avgScore >= 0.4:
		overallLevel = "MEDIUM"
	}

	contract := &models.Contract{
		Title:            title,
		PartyA:           partyA,
		PartyB:           partyB,
		UploadedAt:       time.Now(),
		OverallRiskLevel: overallLevel,
		OverallRiskScore: avgScore,
		Clauses:          clauses,
	}

	if err := db.Create(contract).Error; err != nil {
		return nil, err
	}

	// Extract obligations now that clauses have IDs
	var persistedClauses []models.Clause
	if err := db.Where("contract_id = ?", contract.ID).Find(&persistedClauses).Error; err == nil {
		obls := ExtractObligationsFromClauses(contract.ID, persistedClauses)
		if len(obls) > 0 {
			_ = db.Create(&obls).Error
		}
	}

	return contract, nil
}

func GetClausesForContract(db *gorm.DB, contractID uint) ([]models.Clause, error) {
	var clauses []models.Clause
	if err := db.Where("contract_id = ?", contractID).Order("order_index ASC").
		Preload("Triggers").
		Find(&clauses).Error; err != nil {
		return nil, err
	}
	return clauses, nil
}

func GetObligationsForContract(db *gorm.DB, contractID uint) ([]models.Obligation, error) {
	var obls []models.Obligation
	if err := db.Where("contract_id = ?", contractID).Order("id ASC").Find(&obls).Error; err != nil {
		return nil, err
	}
	return obls, nil
}

type RiskSummary struct {
	TotalClauses int               `json:"totalClauses"`
	ByLevel      map[string]int    `json:"byLevel"`
	ByType       map[string]int    `json:"byType"`
	MaxRisk      float64           `json:"maxRisk"`
}

func ComputeRiskSummary(db *gorm.DB, contractID uint) (*RiskSummary, error) {
	var clauses []models.Clause
	if err := db.Where("contract_id = ?", contractID).Find(&clauses).Error; err != nil {
		return nil, err
	}

	summary := &RiskSummary{
		TotalClauses: len(clauses),
		ByLevel:      map[string]int{"LOW": 0, "MEDIUM": 0, "HIGH": 0},
		ByType:       map[string]int{},
		MaxRisk:      0,
	}

	for _, cl := range clauses {
		summary.ByLevel[cl.RiskLevel]++
		summary.ByType[cl.ClauseType]++
		if cl.RiskScore > summary.MaxRisk {
			summary.MaxRisk = cl.RiskScore
		}
	}

	return summary, nil
}
