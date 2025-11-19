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

	// Liability flags
	if strings.Contains(lower, "unlimited liability") || strings.Contains(lower, "without limitation") {
		addTrigger("unlimited liability", "UNLIMITED_LIABILITY", "HIGH", 0.6)
	}
	if strings.Contains(lower, "consequential") || strings.Contains(lower, "indirect") || strings.Contains(lower, "special or punitive") {
		// Absence of exclusion is risky
		if !strings.Contains(lower, "excluding") && !strings.Contains(lower, "shall not be liable for") {
			addTrigger("consequential damages not excluded", "CONSEQUENTIAL_NOT_EXCLUDED", "HIGH", 0.3)
		} else {
			addTrigger("consequential damages excluded", "CONSEQUENTIAL_EXCLUDED", "LOW", 0.05)
		}
	}
	if strings.Contains(lower, "aggregate liability") && strings.Contains(lower, "fees") {
		if strings.Contains(lower, "twelve (12) months") || strings.Contains(lower, "12 months") || strings.Contains(lower, "one (1) year") {
			addTrigger("liability capped to 12 months fees", "LIABILITY_CAP_12M", "LOW", 0.1)
		} else {
			addTrigger("liability cap present", "LIABILITY_CAP", "MEDIUM", 0.2)
		}
	}

	// Term / termination
	if classification.Type == "TERMINATION" {
		if !strings.Contains(lower, "for convenience") {
			addTrigger("no termination for convenience", "NO_TERMINATION_FOR_CONVENIENCE", "MEDIUM", 0.2)
		}
		if strings.Contains(lower, "auto-renew") || strings.Contains(lower, "automatically renew") {
			addTrigger("auto-renewal", "AUTO_RENEWAL", "MEDIUM", 0.15)
		}
	}

	// Data protection / security
	if strings.Contains(lower, "personal data") || strings.Contains(lower, "gdpr") || strings.Contains(lower, "data protection") {
		if strings.Contains(lower, "breach") && strings.Contains(lower, "notify") {
			addTrigger("breach notification defined", "BREACH_NOTICE", "LOW", 0.05)
		}
		if strings.Contains(lower, "without undue delay") || strings.Contains(lower, "no later than 72 hours") {
			addTrigger("timely breach notification", "BREACH_NOTICE_TIMELY", "LOW", 0.05)
		}
		if strings.Contains(lower, "processor") && !strings.Contains(lower, "subprocessor") {
			addTrigger("data processing obligations incomplete", "DATA_PROCESSING_GAP", "MEDIUM", 0.2)
		}
	}

	// Payment
	if strings.Contains(lower, "within 30 days") || strings.Contains(lower, "net 30") {
		addTrigger("30-days payment term", "PAYMENT_TERM_30D", "LOW", 0.05)
	}
	if strings.Contains(lower, "late fee") || strings.Contains(lower, "interest at the rate") {
		addTrigger("late payment interest", "LATE_FEE", "MEDIUM", 0.15)
	}

	// Indemnity
	if strings.Contains(lower, "indemnify") || strings.Contains(lower, "indemnification") {
		if strings.Contains(lower, "sole remedy") {
			addTrigger("indemnity sole remedy", "INDEMNITY_SOLE_REMEDY", "MEDIUM", 0.15)
		}
		// One-sided indemnity (only one party named)
		if strings.Contains(lower, "customer") && !strings.Contains(lower, "vendor") {
			addTrigger("one-sided indemnity in favor of customer", "ONE_SIDED_INDEMNITY_CUSTOMER", "MEDIUM", 0.2)
		}
		if strings.Contains(lower, "vendor") && !strings.Contains(lower, "customer") {
			addTrigger("one-sided indemnity in favor of vendor", "ONE_SIDED_INDEMNITY_VENDOR", "HIGH", 0.3)
		}
	}

	// Governing law / jurisdiction – not directly risky but useful signals
	if strings.Contains(lower, "governing law") {
		addTrigger("governing law specified", "GOVERNING_LAW", "LOW", 0.02)
	}
	if strings.Contains(lower, "exclusive jurisdiction") {
		addTrigger("exclusive jurisdiction", "EXCLUSIVE_JURISDICTION", "MEDIUM", 0.1)
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
