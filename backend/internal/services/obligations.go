package services

import (
	"regexp"
	"strings"
	"time"

	"lawlens-g/internal/models"
)

var (
	deadlineRegex  = regexp.MustCompile(`within\s+(\d+)\s+days`)
	byDateRegex    = regexp.MustCompile(`no later than\s+([A-Za-z0-9 ,]+)`)
	retentionRegex = regexp.MustCompile(`retain\s+records?\s+for\s+(\d+)\s+years?`)
)

// Extract obligations from clauses using primitive heuristics.
func ExtractObligationsFromClauses(contractID uint, clauses []models.Clause) []models.Obligation {
	obligations := []models.Obligation{}
	now := time.Now()

	for _, cl := range clauses {
		lower := strings.ToLower(cl.Text)
		if !(strings.Contains(lower, "shall") || strings.Contains(lower, "must") || strings.Contains(lower, "will")) {
			continue
		}

		desc := strings.TrimSpace(cl.Text)
		var due *time.Time

		if m := deadlineRegex.FindStringSubmatch(lower); len(m) == 2 {
			days := parseIntSafe(m[1])
			if days > 0 {
				d := now.AddDate(0, 0, days)
				due = &d
			}
		} else if m := retentionRegex.FindStringSubmatch(lower); len(m) == 2 {
			years := parseIntSafe(m[1])
			if years > 0 {
				d := now.AddDate(years, 0, 0)
				due = &d
			}
		} else if m := byDateRegex.FindStringSubmatch(lower); len(m) == 2 {
			// For now just keep description; real date parsing would be more complex.
			_ = m // could be used for a free-text "Due by" note later.
		}

		obligations = append(obligations, models.Obligation{
			ContractID:  contractID,
			ClauseID:    cl.ID,
			Description: desc,
			DueDate:     due,
		})
	}

	return obligations
}

func parseIntSafe(s string) int {
	n := 0
	for _, ch := range s {
		if ch < '0' || ch > '9' {
			return 0
		}
		n = n*10 + int(ch-'0')
	}
	return n
}
