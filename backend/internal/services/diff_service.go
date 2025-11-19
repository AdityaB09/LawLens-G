package services

import (
	"errors"

	"lawlens-g/internal/models"

	"gorm.io/gorm"
)

type ContractDiff struct {
	ContractAID uint             `json:"contractAId"`
	ContractBID uint             `json:"contractBId"`
	ClausesA    []models.Clause  `json:"clausesA"`
	ClausesB    []models.Clause  `json:"clausesB"`
}

// Very naive diff: align by OrderIndex
func CompareContracts(db *gorm.DB, aID, bID uint) (*ContractDiff, error) {
	var aClauses, bClauses []models.Clause

	if err := db.Where("contract_id = ?", aID).Order("order_index ASC").Find(&aClauses).Error; err != nil {
		return nil, err
	}
	if err := db.Where("contract_id = ?", bID).Order("order_index ASC").Find(&bClauses).Error; err != nil {
		return nil, err
	}
	if len(aClauses) == 0 && len(bClauses) == 0 {
		return nil, errors.New("no clauses for either contract")
	}

	// Mark status & risk delta
	maxLen := len(aClauses)
	if len(bClauses) > maxLen {
		maxLen = len(bClauses)
	}

	for i := 0; i < maxLen; i++ {
		var a *models.Clause
		var b *models.Clause

		if i < len(aClauses) {
			a = &aClauses[i]
		}
		if i < len(bClauses) {
			b = &bClauses[i]
		}

		switch {
		case a != nil && b != nil:
			a.Status = "BASELINE"
			b.Status = "UNCHANGED"
			if a.RiskScore != b.RiskScore {
				b.Status = "MODIFIED"
				b.RiskDelta = b.RiskScore - a.RiskScore
			}
		case a != nil && b == nil:
			a.Status = "REMOVED"
		case a == nil && b != nil:
			b.Status = "ADDED"
		}
	}

	return &ContractDiff{
		ContractAID: aID,
		ContractBID: bID,
		ClausesA:    aClauses,
		ClausesB:    bClauses,
	}, nil
}
