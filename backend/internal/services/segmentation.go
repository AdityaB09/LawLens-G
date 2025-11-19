package services

import (
	"regexp"
	"strings"

	"lawlens-g/internal/models"
)

var headingRegex = regexp.MustCompile(`(?m)^\s*(\d+(\.\d+)*)\s+([A-Z][A-Za-z ]+)\s*$`)

// Simple clause segmentation: split on double newlines + detect headings if present.
func SegmentClauses(text string) []models.Clause {
	parts := strings.Split(text, "\n\n")

	clauses := make([]models.Clause, 0, len(parts))
	order := 0

	for _, block := range parts {
		block = strings.TrimSpace(block)
		if block == "" {
			continue
		}
		order++

		heading := ""
		lines := strings.Split(block, "\n")
		if len(lines) > 0 {
			if headingRegex.MatchString(lines[0]) {
				heading = strings.TrimSpace(lines[0])
				block = strings.TrimSpace(strings.Join(lines[1:], "\n"))
			}
		}

		clauses = append(clauses, models.Clause{
			OrderIndex: order,
			Heading:    heading,
			Text:       block,
		})
	}

	if len(clauses) == 0 && strings.TrimSpace(text) != "" {
		clauses = append(clauses, models.Clause{
			OrderIndex: 1,
			Heading:    "",
			Text:       strings.TrimSpace(text),
		})
	}

	return clauses
}
