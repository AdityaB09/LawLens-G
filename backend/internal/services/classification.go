package services

import (
	"strings"
)

type ClauseClassification struct {
	Type     string
	Triggers []string
}

// Super lightweight keyword-based classifier.
// You can later swap this for ONNX or a Python microservice.
func ClassifyClause(text string) ClauseClassification {
	lower := strings.ToLower(text)
	triggers := []string{}
	typ := "OTHER"

	check := func(keywords []string) bool {
		for _, k := range keywords {
			if strings.Contains(lower, k) {
				return true
			}
		}
		return false
	}

	if check([]string{"confidential", "non-disclosure", "nda"}) {
		typ = "CONFIDENTIALITY"
		triggers = append(triggers, "confidentiality_keywords")
	}
	if check([]string{"terminate", "termination", "term and termination"}) {
		typ = "TERMINATION"
		triggers = append(triggers, "termination_keywords")
	}
	if check([]string{"liability", "hold harmless", "indemnify", "indemnification"}) {
		typ = "LIABILITY"
		triggers = append(triggers, "liability_keywords")
	}
	if check([]string{"payment", "fees", "invoice", "amount due"}) {
		typ = "PAYMENT"
		triggers = append(triggers, "payment_keywords")
	}
	if check([]string{"intellectual property", "ip", "ownership", "license"}) {
		typ = "IP"
		triggers = append(triggers, "ip_keywords")
	}
	if check([]string{"non-compete", "non compete", "restrict", "solicit"}) {
		typ = "NON_COMPETE"
		triggers = append(triggers, "non_compete_keywords")
	}
	if check([]string{"arbitration", "dispute resolution"}) {
		typ = "ARBITRATION"
		triggers = append(triggers, "arbitration_keywords")
	}

	return ClauseClassification{Type: typ, Triggers: triggers}
}
