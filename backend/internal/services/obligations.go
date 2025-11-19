package services

import (
	"regexp"
	"strings"
	"time"

	"lawlens-g/internal/models"
)

var deadlineRegex = regexp.MustCompile(`within\s+(\d+)\s+days`)

// Extract obligations from clauses using primitive heuristics.
func ExtractObligationsFromClauses(contractID uint, clauses []models.Clause) []models.Obligation {
	obligations := []models.Obligation{}
	now := time.Now()

	for _, cl := range clauses {
		lower := strings.ToLower(cl.Text)
		if !(strings.Contains(lower, "shall") || strings.Contains(lower, "must") || strings.Contains(lower, "will")) {
			continue
		}

		desc := cl.Text
		var due *time.Time

		if m := deadlineRegex.FindStringSubmatch(lower); len(m) == 2 {
			// Ignore parse error safely
			daysStr := m[1]
			var days int
			for _, ch := range daysStr {
				if ch < '0' || ch > '9' {
					days = 0
					break
				}
				days = days*10 + int(ch-'0')
			}
			if days > 0 {
				d := now.AddDate(0, 0, days)
				due = &d
			}
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
