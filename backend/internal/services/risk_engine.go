package services

import (
	"strings"

	"lawlens-g/internal/models"
)

type RiskResult struct {
	Score    float64
	Level    string
	Triggers []models.ClauseTrigger
}

func EvaluateRisk(clauseText string, classification ClauseClassification) RiskResult {
	lower := strings.ToLower(clauseText)
	score := 0.1
	triggers := []models.ClauseTrigger{}

	addTrigger := func(phrase, ttype, severity string, delta float64) {
		triggers = append(triggers, models.ClauseTrigger{
			Phrase:   phrase,
			Type:     ttype,
			Severity: severity,
		})
		score += delta
	}

	if strings.Contains(lower, "unlimited liability") || strings.Contains(lower, "without limitation") {
		addTrigger("unlimited liability", "UNLIMITED_LIABILITY", "HIGH", 0.6)
	}

	if strings.Contains(lower, "aggregate liability") && strings.Contains(lower, "fees") {
		addTrigger("capped to fees", "LIABILITY_CAP", "MEDIUM", 0.2)
	}

	if classification.Type == "TERMINATION" && !strings.Contains(lower, "for convenience") {
		addTrigger("no termination for convenience", "NO_TERMINATION_FOR_CONVENIENCE", "MEDIUM", 0.2)
	}

	if strings.Contains(lower, "within 30 days") || strings.Contains(lower, "net 30") {
		addTrigger("30-days payment", "PAYMENT_TERM_30D", "LOW", 0.05)
	}

	if score > 1.0 {
		score = 1.0
	}

	level := "LOW"
	switch {
	case score >= 0.7:
		level = "HIGH"
	case score >= 0.4:
		level = "MEDIUM"
	}

	return RiskResult{
		Score:    score,
		Level:    level,
		Triggers: triggers,
	}
}
